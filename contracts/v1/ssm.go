/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package v1

/////////////////////////////////////////////////////////////////////////////////
// This module should only be used for LOCAL TESTING or on a PHYSICALLY SECURE //
// device.  It CANNOT guarantee the protection of the private keys from people //
// and other processes that have access to the RAM and storage devices for the //
// device.                                                                     //
//                           YOU HAVE BEEN WARNED!!!                           //
/////////////////////////////////////////////////////////////////////////////////

import (
	sig "crypto/ed25519"
	dig "crypto/sha512"
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	bal "github.com/bali-nebula/go-component-framework/bali"
	col "github.com/bali-nebula/go-component-framework/collections"
	com "github.com/bali-nebula/go-component-framework/components"
	ele "github.com/bali-nebula/go-component-framework/elements"
	str "github.com/bali-nebula/go-component-framework/strings"
	uti "github.com/bali-nebula/go-component-framework/utilities"
)

// CONSTANT DEFINITIONS

const (
	protocol     = "v1"
	digest       = "sha512"
	signature    = "ed25519"
	generateKeys = "GenerateKeys"
	signBytes    = "SignBytes"
	rotateKeys   = "RotateKeys"
	keyless      = "Keyless"
	loneKey      = "LoneKey"
	twoKeys      = "TwoKeys"
	invalid      = "Invalid"
)

var events = []abs.Event{generateKeys, signBytes, rotateKeys}
var states = []abs.State{keyless, loneKey, twoKeys}
var table = [][]abs.State{
	{loneKey, invalid, invalid},
	{invalid, loneKey, twoKeys},
	{invalid, loneKey, invalid},
}

// NOTARY IMPLEMENTATION

// This constructor creates a new software security module.
func SSM(directory string) abs.SecurityModuleLike {
	var filename = "SSM" + protocol + ".bali"
	var configuration = col.Catalog()
	var configurator = uti.Configurator(directory, filename)
	var controller = uti.Controller(events, states, table, keyless)
	return &ssm{configuration, configurator, controller}
}

// This type defines the structure and methods associated with a software
// security module (SSM).
type ssm struct {
	configuration abs.CatalogLike
	configurator abs.ConfiguratorLike
	controller abs.ControllerLike
}

// SECURE INTERFACE

// This method generates a new public-private key pair and returns the public
// key.
func (v *ssm) GenerateKeys() abs.PublicKey {
	v.controller.TransitionState(generateKeys)
	var publicKey, privateKey, err = sig.GenerateKey(nil) // Uses crypto/rand.Reader.
	if err != nil {
		var message = fmt.Sprintf("Could not generate a new public-private keypair: %v.", err)
		panic(message)
	}
	v.configuration.SetValue(ele.Symbol("privateKey"), com.Component(str.BinaryFromArray(privateKey)))
	v.configuration.SetValue(ele.Symbol("publicKey"), com.Component(str.BinaryFromArray(publicKey)))
	v.configurator.Store(abs.Configuration(bal.FormatEntity(v.configuration)))
	return abs.PublicKey(publicKey)
}

// This method generates a digital digest of the specified bytes and returns
// the digital digest.
func (v *ssm) DigestBytes(bytes []byte) abs.Digest {
	var digest = dig.Sum512(bytes)
	return abs.Digest(digest[:])
}

// This method digitally signs the specified bytes using the private key and
// returns the digital signature.
func (v *ssm) SignBytes(bytes []byte) abs.Signature {
	v.controller.TransitionState(signBytes)
	var privateKey = v.configuration.GetValue(ele.Symbol("privateKey")).GetEntity().(str.Binary).AsArray()
	var signature = sig.Sign(privateKey, bytes)
	return abs.Signature(signature)
}

// This method determines whether or not the specified digital signature is
// valid for the specified bytes using the specified public key.
func (v *ssm) IsValid(publicKey abs.PublicKey, signature abs.Signature, bytes []byte) bool {
	var isValid = sig.Verify(sig.PublicKey(publicKey), bytes, signature)
	return isValid
}

// This method replaces the existing public-key pair with a new one and returns
// the new public key.
func (v *ssm) RotateKeys() abs.PublicKey {
	v.controller.TransitionState(rotateKeys)
	var publicKey, privateKey, err = sig.GenerateKey(nil) // Uses crypto/rand.Reader.
	if err != nil {
		var message = fmt.Sprintf("Could not rotate the public-private keypair: %v.", err)
		panic(message)
	}
	v.configuration.SetValue(ele.Symbol("privateKey"), com.Component(str.BinaryFromArray(privateKey)))
	v.configuration.SetValue(ele.Symbol("publicKey"), com.Component(str.BinaryFromArray(publicKey)))
	v.configurator.Store(abs.Configuration(bal.FormatEntity(v.configuration)))
	return abs.PublicKey(publicKey)
}

// This method erases the existing public-key pair.
func (v *ssm) EraseKeys() {
	v.configuration = nil
	v.configurator.Delete()
	v.controller.SetState(keyless)
}
