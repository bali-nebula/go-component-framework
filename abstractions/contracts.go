/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// TYPE DEFINITIONS

type (
	Certificate any
	Citation    any
	Contract    any
	Credentials any
	Digest      any
	Document    any
	PublicKey   any
	Salt        any
	Signature   any
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all hardened security
// modules.
type Hardened interface {
	GenerateKeys() PublicKey
	DigestBytes(bytes []byte) Digest
	SignBytes(bytes []byte) Signature
	IsValid(publicKey PublicKey, signature Signature, bytes []byte) bool
	RotateKeys() PublicKey
	EraseKeys()
}

// This interface defines the methods supported by all prudent notary
// agents.
type Prudent interface {
	GenerateKey() Certificate
	ActivateKey(certificate Certificate) Citation
	GetCitation() Citation
	RefreshKey() Certificate
	ForgetKey()
}

// This interface defines the methods supported by all certified notary
// agents.
type Certified interface {
	GenerateCredentials(salt Salt) Credentials
	NotarizeDocument(document Document) Contract
	IsValid(contract Contract, certificate Certificate) bool
	CiteDocument(document Document) Citation
	CitationMatches(citation Citation, document Document) bool
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by
// security-module-like devices.
type SecurityModuleLike interface {
	Hardened
}

// This interface consolidates all the interfaces supported by notary-like
// agents.
type NotaryLike interface {
	Prudent
	Certified
}
