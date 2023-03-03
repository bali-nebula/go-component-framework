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

// SOFTWARE SECURITY MODULE (SSM) INTERFACE

// This constructor creates a new software security module.
func SSM(directory string) abs.SecurityModuleLike {
	var configurator abs.ConfiguratorLike
	var controller abs.ControllerLike
	var v = &ssm{} // Assume this is only a trusted security module.
	if len(directory) > 0 {
		// Nope, this is pretending to be a hardened security module!
		fmt.Println("WARNING: Using a SOFTWARE security module instead of a HARDWARE security module.")
		var filename = "SSMv1.bali"
		configurator = age.Configurator(directory, filename)
		controller = age.Controller(states)
		v = &ssm{configurator: configurator, controller: controller}
		if configurator.Exists() {
			v.readConfiguration()
		} else {
			v.createConfiguration()
		}
	}
	return v
}

// SOFTWARE SECURITY MODULE (SSM) IMPLEMENTATION

// This type defines the structure and methods associated with a software
// security module (SSM).
type ssm struct {
	tag          abs.TagLike
	publicKey    abs.BinaryLike
	privateKey   abs.BinaryLike
	previousKey  abs.BinaryLike
	configurator abs.ConfiguratorLike
	controller   abs.ControllerLike
}

// These constants define the possible states for the state machine.
const (
	invalid int = iota
	keyless
	loneKey
	twoKeys
)

// These constants define the possible events for the state machine.
const (
	events int = iota
	generateKeys
	signBytes
	rotateKeys
)

// This table defines the allowed transitions for the state machine.
var states = [][]int{
	{events, generateKeys, signBytes, rotateKeys},
	{keyless, loneKey,     invalid,   invalid},
	{loneKey, invalid,     loneKey,   twoKeys},
	{twoKeys, invalid,     loneKey,   invalid},
}

// This controller implements the finite state machine.
var controller = age.Controller(states)

// These constants define the possible attribute names for the configuration.
const (
	tagAttribute         = "tag"
	stateAttribute       = "state"
	publicKeyAttribute   = "publicKey"
	privateKeyAttribute  = "privateKey"
	previousKeyAttribute = "previousKey"
)

// This is the latest security protocol used.
var protocol = bal.ParseEntity(`[
    $protocol: v1
    $digest: "sha512"
    $signature: "ed25519"
]`).(abs.CatalogLike)

// TRUSTED INTERFACE

// This method retrieves the protocol version for this security module.
func (v *ssm) GetVersion() string {
	return "1"
}

// This method generates a digital digest of the specified bytes and returns
// the digital digest.
func (v *ssm) DigestBytes(bytes []byte) abs.Digest {
	var digest = dig.Sum512(bytes)
	return abs.Digest(digest[:])
}

// This method determines whether or not the specified digital signature is
// valid for the specified bytes using the specified public key.
func (v *ssm) IsValid(public abs.PublicKey, signature abs.Signature, bytes []byte) bool {
	var isValid = sig.Verify(sig.PublicKey(public), bytes, signature)
	return isValid
}

// HARDENED INTERFACE

// This method generates a new public-private key pair and returns the public
// key.
func (v *ssm) GenerateKeys() abs.PublicKey {
	if v.configurator == nil {
		panic("Attempted to generate keys on a non-hardened security module")
	}
	v.controller.TransitionState(generateKeys)
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

// This method digitally signs the specified bytes using the private key and
// returns the digital signature.
func (v *ssm) SignBytes(bytes []byte) abs.Signature {
	if v.configurator == nil {
		panic("Attempted to sign bytes on a non-hardened security module")
	}
	v.controller.TransitionState(signBytes)
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

// This method replaces the existing public-key pair with a new one and returns
// the new public key.
func (v *ssm) RotateKeys() abs.PublicKey {
	if v.configurator == nil {
		panic("Attempted to rotated keys on a non-hardened security module")
	}
	v.controller.TransitionState(rotateKeys)
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
	if v.configurator == nil {
		panic("Attempted to erase keys on a non-hardened security module")
	}
	v.controller.SetState(keyless)
	v.deleteConfiguration()
}

// PRIVATE METHODS

func (v *ssm) stateAsString() string {
	switch v.controller.GetState() {
	case keyless:
		return "keyless"
	case loneKey:
		return "loneKey"
	case twoKeys:
		return "twoKeys"
	default:
		return "invalid"
	}
}

func (v *ssm) setState(state string) {
	switch state {
	case "keyless":
		v.controller.SetState(keyless)
	case "loneKey":
		v.controller.SetState(loneKey)
	case "twoKeys":
		v.controller.SetState(twoKeys)
	default:
		v.controller.SetState(invalid)
	}
}

func (v *ssm) createConfiguration() {
	v.tag = ele.TagOfSize(20) // Generate a new random tag.
	var configuration = bal.ParseEntity(`[
    $tag: #` + v.tag.AsString() + `
    $state: "` + v.stateAsString() + `"
]`).(abs.CatalogLike)
	var document = []byte(bal.FormatEntity(configuration) + "\n")
	v.configurator.Store(document)
}

func (v *ssm) readConfiguration() {
	var document = v.configurator.Load()
	var component = bal.ParseDocument(document)
	var configuration = component.ExtractCatalog()
	v.tag = configuration.GetValue(ele.SymbolFromString(tagAttribute)).ExtractTag()
	var state = configuration.GetValue(ele.SymbolFromString(stateAttribute)).ExtractQuote()
	v.setState(state.AsString())
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
	var component = com.Component(v.tag)
	configuration.SetValue(ele.SymbolFromString(tagAttribute), component)
	component = com.Component(str.QuoteFromString(v.stateAsString()))
	configuration.SetValue(ele.SymbolFromString(stateAttribute), component)
	if v.publicKey != nil {
		component = com.Component(v.publicKey)
		configuration.SetValue(ele.SymbolFromString(publicKeyAttribute), component)
	}
	if v.privateKey != nil {
		component = com.Component(v.privateKey)
		configuration.SetValue(ele.SymbolFromString(privateKeyAttribute), component)
	}
	if v.previousKey != nil {
		component = com.Component(v.previousKey)
		configuration.SetValue(ele.SymbolFromString(previousKeyAttribute), component)
	}
	var document = []byte(bal.FormatEntity(configuration) + "\n")
	v.configurator.Store(document)
}

func (v *ssm) deleteConfiguration() {
	v.tag = nil
	v.publicKey = nil
	v.privateKey = nil
	v.previousKey = nil
	v.configurator.Delete()
}
