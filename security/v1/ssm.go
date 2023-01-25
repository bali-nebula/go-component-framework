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
	col "github.com/bali-nebula/go-component-framework/collections"
	com "github.com/bali-nebula/go-component-framework/components"
	ele "github.com/bali-nebula/go-component-framework/elements"
	str "github.com/bali-nebula/go-component-framework/strings"
)

// CONSTANT DEFINITIONS

const (
	tagAttribute         = "tag"
	stateAttribute       = "state"
	publicKeyAttribute   = "publicKey"
	privateKeyAttribute  = "privateKey"
	previousKeyAttribute = "previousKey"
	generateKeys         = "generateKeys"
	signBytes            = "signBytes"
	rotateKeys           = "rotateKeys"
	keyless              = "keyless"
	loneKey              = "loneKey"
	twoKeys              = "twoKeys"
	invalid              = "invalid"
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
	var configurator = age.Configurator(directory, filename)
	var controller = age.Controller(events, states, table, keyless)
	var v = &ssm{configurator: configurator, controller: controller}
	if configurator.Exists() {
		v.readConfiguration()
	} else {
		v.createConfiguration()
	}
	return v
}

// This type defines the structure and methods associated with a software
// security module (SSM).
type ssm struct {
	tag          abs.TagLike
	state        abs.State
	publicKey    abs.BinaryLike
	privateKey   abs.BinaryLike
	previousKey  abs.BinaryLike
	configurator abs.ConfiguratorLike
	controller   abs.ControllerLike
}

// SECURE INTERFACE

// This method generates a new public-private key pair and returns the public
// key.
func (v *ssm) GenerateKeys() abs.PublicKey {
	v.state = v.controller.TransitionState(generateKeys)
	var public, private, err = sig.GenerateKey(nil) // Uses crypto/rand.Reader.
	if err != nil {
		var message = fmt.Sprintf("Could not generate a new public-private keypair: %v.", err)
		panic(message)
	}
	v.privateKey = str.BinaryFromArray(private)
	v.publicKey = str.BinaryFromArray(public)
	v.updateConfiguration()
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
	v.state = v.controller.TransitionState(signBytes)
	var privateKey = v.previousKey
	if privateKey != nil {
		v.previousKey = nil
	} else {
		privateKey = v.privateKey
	}
	var private = privateKey.AsArray()
	var signature = sig.Sign(private, bytes)
	v.updateConfiguration()
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
	v.state = v.controller.TransitionState(rotateKeys)
	var public, private, err = sig.GenerateKey(nil) // Uses crypto/rand.Reader.
	if err != nil {
		var message = fmt.Sprintf("Could not rotate the public-private keypair: %v.", err)
		panic(message)
	}
	v.previousKey = v.privateKey
	v.privateKey = str.BinaryFromArray(private)
	v.publicKey = str.BinaryFromArray(public)
	v.updateConfiguration()
	return abs.PublicKey(public)
}

// This method erases the existing public-key pair.
func (v *ssm) EraseKeys() {
	v.controller.SetState(keyless)
	v.deleteConfiguration()
}

// PRIVATE METHODS

func (v *ssm) createConfiguration() {
	v.tag = ele.TagOfSize(20) // Generate a new random tag.
	v.state = keyless
	var configuration = bal.ParseEntity(`[
    $tag: #` + v.tag.AsString() + `
    $state: "` + string(v.state) + `"
]`).(abs.CatalogLike)
	var document = []byte(bal.FormatEntity(configuration) + "\n")
	v.configurator.Store(document)
}

func (v *ssm) readConfiguration() {
	var document = v.configurator.Load()
	var component = bal.ParseDocument(document)
	var configuration = component.ExtractCatalog()
	v.tag = configuration.GetValue(ele.SymbolFromString(tagAttribute)).ExtractTag()
	v.state = abs.State(configuration.GetValue(ele.SymbolFromString(stateAttribute)).ExtractQuote().AsString())
	component = configuration.GetValue(ele.SymbolFromString(publicKeyAttribute))
	if component != nil {
		v.publicKey = component.ExtractBinary()
	}
	component = configuration.GetValue(ele.SymbolFromString(privateKeyAttribute))
	if component != nil {
		v.privateKey = component.ExtractBinary()
	}
	component = configuration.GetValue(ele.SymbolFromString(previousKeyAttribute))
	if component != nil {
		v.previousKey = component.ExtractBinary()
	}
}

func (v *ssm) updateConfiguration() {
	var configuration = col.Catalog()
	configuration.SetValue(ele.SymbolFromString(tagAttribute), com.Component(v.tag))
	configuration.SetValue(ele.SymbolFromString(stateAttribute), com.Component((str.QuoteFromString(string(v.state)))))
	if v.publicKey != nil {
		configuration.SetValue(ele.SymbolFromString(publicKeyAttribute), com.Component(v.publicKey))
	}
	if v.privateKey != nil {
		configuration.SetValue(ele.SymbolFromString(privateKeyAttribute), com.Component(v.privateKey))
	}
	if v.previousKey != nil {
		configuration.SetValue(ele.SymbolFromString(previousKeyAttribute), com.Component(v.previousKey))
	}
	var document = []byte(bal.FormatEntity(configuration) + "\n")
	v.configurator.Store(document)
}

func (v *ssm) deleteConfiguration() {
	v.tag = nil
	v.state = invalid
	v.publicKey = nil
	v.privateKey = nil
	v.previousKey = nil
	v.configurator.Delete()
}
