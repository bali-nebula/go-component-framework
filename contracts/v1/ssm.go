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
	age "github.com/bali-nebula/go-component-framework/agents"
	bal "github.com/bali-nebula/go-component-framework/bali"
	com "github.com/bali-nebula/go-component-framework/components"
	ele "github.com/bali-nebula/go-component-framework/elements"
	str "github.com/bali-nebula/go-component-framework/strings"
)

// CONSTANT DEFINITIONS

const (
	tagKey     = ele.Symbol("tag")
	publicKey  = ele.Symbol("publicKey")
	privateKey = ele.Symbol("privateKey")
	stateKey   = ele.Symbol("state")

	generateKeys = "generateKeys"
	signBytes    = "signBytes"
	rotateKeys   = "rotateKeys"
	keyless      = "keyless"
	loneKey      = "loneKey"
	twoKeys      = "twoKeys"
	invalid      = "invalid"
)

var events = []abs.Event{generateKeys, signBytes, rotateKeys}

var states = []abs.State{keyless, loneKey, twoKeys}

var table = [][]abs.State{
	{loneKey, invalid, invalid},
	{invalid, loneKey, twoKeys},
	{invalid, loneKey, invalid},
}

var protocol = bal.ParseEntity(`[
    $protocol: v1
    $digest: "sha512"
    $signature: "ed25519"
]`).(abs.CatalogLike)

// NOTARY IMPLEMENTATION

// This constructor creates a new software security module.
func SSM(directory string) abs.SecurityModuleLike {
	var filename = "SSMv1.bali"
	var state = abs.State(keyless)
	var configuration abs.CatalogLike
	var configurator = age.Configurator(directory, filename)
	if configurator.Exists() {
		configuration = loadConfiguration(configurator)
		state = abs.State(configuration.GetValue(stateKey).GetEntity().(str.Quote))
	} else {
		var tag = "#" + ele.TagOfSize(20) // Generate a new random tag.
		configuration = bal.ParseEntity(`[
    $tag: ` + string(tag) + `
    $state: "keyless"
]`).(abs.CatalogLike)
	}
	var controller = age.Controller(events, states, table, state)
	return &ssm{configuration, configurator, controller}
}

// This type defines the structure and methods associated with a software
// security module (SSM).
type ssm struct {
	configuration abs.CatalogLike
	configurator  abs.ConfiguratorLike
	controller    abs.ControllerLike
}

// SECURE INTERFACE

// This method generates a new public-private key pair and returns the public
// key.
func (v *ssm) GenerateKeys() abs.PublicKey {
	v.controller.TransitionState(generateKeys)
	var public, private, err = sig.GenerateKey(nil) // Uses crypto/rand.Reader.
	if err != nil {
		var message = fmt.Sprintf("Could not generate a new public-private keypair: %v.", err)
		panic(message)
	}
	v.configuration.SetValue(privateKey, com.Component(str.BinaryFromArray(private)))
	v.configuration.SetValue(publicKey, com.Component(str.BinaryFromArray(public)))
	storeConfiguration(v.configurator, v.configuration)
	return abs.PublicKey(public)
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
	var private = v.configuration.GetValue(privateKey).GetEntity().(str.Binary).AsArray()
	var signature = sig.Sign(private, bytes)
	return abs.Signature(signature)
}

// This method determines whether or not the specified digital signature is
// valid for the specified bytes using the specified public key.
func (v *ssm) IsValid(public abs.PublicKey, signature abs.Signature, bytes []byte) bool {
	var isValid = sig.Verify(sig.PublicKey(public), bytes, signature)
	return isValid
}

// This method replaces the existing public-key pair with a new one and returns
// the new public key.
func (v *ssm) RotateKeys() abs.PublicKey {
	v.controller.TransitionState(rotateKeys)
	var public, private, err = sig.GenerateKey(nil) // Uses crypto/rand.Reader.
	if err != nil {
		var message = fmt.Sprintf("Could not rotate the public-private keypair: %v.", err)
		panic(message)
	}
	v.configuration.SetValue(privateKey, com.Component(str.BinaryFromArray(private)))
	v.configuration.SetValue(publicKey, com.Component(str.BinaryFromArray(public)))
	storeConfiguration(v.configurator, v.configuration)
	return abs.PublicKey(public)
}

// This method erases the existing public-key pair.
func (v *ssm) EraseKeys() {
	v.configuration = nil
	v.configurator.Delete()
	v.controller.SetState(keyless)
}

// PRIVATE METHODS

func loadConfiguration(configurator abs.ConfiguratorLike) abs.CatalogLike {
	var document = configurator.Load()
	var component = bal.ParseDocument(document)
	var entity = component.GetEntity()
	var configuration = entity.(abs.CatalogLike)
	return configuration
}

func storeConfiguration(configurator abs.ConfiguratorLike, configuration abs.CatalogLike) {
	var document = []byte(bal.FormatEntity(configuration) + "\n")
	configurator.Store(document)
}
