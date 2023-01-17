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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// CONSTANT DEFINITIONS

const (
	protocol  = "v1"
	digest    = "sha512"
	signature = "ed25519"
)

// TYPE DEFINITIONS

type (
	Digest    any
	PublicKey any
	Signature any
)

// NOTARY IMPLEMENTATION

// This constructor creates a new software security module.
func SSM() abs.SecurityModuleLike {
	return &ssm{}
}

// This type defines the structure and methods associated with a software
// security module (SSM).
type ssm struct {
}

// SECURE INTERFACE

// This method generates a new public-private key pair and returns the public
// key.
func (v *ssm) GenerateKeys() PublicKey {
	panic("Not yet implemented.")
}

// This method generates a digital digest of the specified bytes and returns
// the digital digest.
func (v *ssm) DigestBytes(bytes []byte) Digest {
	panic("Not yet implemented.")
}

// This method digitally signs the specified bytes using the private key and
// returns the digital signature.
func (v *ssm) SignBytes(bytes []byte) Signature {
	panic("Not yet implemented.")
}

// This method determines whether or not the specified digital signature is
// valid for the specified bytes using the specified public key.
func (v *ssm) IsValid(publicKey PublicKey, signature Signature, bytes []byte) bool {
	panic("Not yet implemented.")
}

// This method replaces the existing public-key pair with a new one and returns
// the new public key.
func (v *ssm) RotateKeys() PublicKey {
	panic("Not yet implemented.")
}

// This method erases the existing public-key pair.
func (v *ssm) EraseKeys() {
	panic("Not yet implemented.")
}
