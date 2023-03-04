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
	Digest     []byte
	PrivateKey []byte
	PublicKey  []byte
	Salt       []byte
	Signature  []byte
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all notarized components.
type Notarized interface {
	GetDocument() DocumentLike
	GetTimestamp() MomentLike
	GetAccount() TagLike
	GetProtocol() VersionLike
	GetCertificate() CitationLike
	GetSignature() BinaryLike
}

// This interface defines the methods supported by all published components.
type Published interface {
	GetAlgorithms() CatalogLike
	GetKey() BinaryLike
}

// This interface defines the methods supported by all referential components.
type Referential interface {
	GetTag() TagLike
	GetVersion() VersionLike
	GetProtocol() VersionLike
	GetDigest() BinaryLike
}

// This interface defines the methods supported by all restricted components.
type Restricted interface {
	GetPermissions() MonikerLike
}

// This interface defines the methods supported by all salted components.
type Salted interface {
	GetSalt() BinaryLike
}

// This interface defines the methods supported by all typed components.
type Typed interface {
	GetType() TypeLike
}

// This interface defines the methods supported by all versioned components.
type Versioned interface {
	GetTag() TagLike
	GetVersion() VersionLike
	GetPrevious() CitationLike
}

// This interface defines the methods supported by all trusted security
// modules.
type Trusted interface {
	GetVersion() string
	DigestBytes(bytes []byte) Digest
	IsValid(publicKey PublicKey, signature Signature, bytes []byte) bool
}

// This interface defines the methods supported by all hardened security
// modules.
type Hardened interface {
	GenerateKeys() PublicKey
	SignBytes(bytes []byte) Signature
	RotateKeys() PublicKey
	EraseKeys()
}

// This interface defines the methods supported by all prudent notary
// agents.
type Prudent interface {
	GenerateKey() CertificateLike
	ActivateKey(certificate CertificateLike) CitationLike
	GetCitation() CitationLike
	RefreshKey() CertificateLike
	ForgetKey()
}

// This interface defines the methods supported by all certified notary
// agents.
type Certified interface {
	GenerateCredentials(salt Salt) CredentialsLike
	NotarizeDocument(document DocumentLike) ContractLike
	IsValid(contract ContractLike, certificate CertificateLike) bool
	CiteDocument(document DocumentLike) CitationLike
	CitationMatches(citation CitationLike, document DocumentLike) bool
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by type-like
// components.
type TypeLike interface {
	Encapsulated
	Typed
}

// This interface consolidates all the interfaces supported by citation-like
// components.
type CitationLike interface {
	Encapsulated
	Referential
	Typed
}

// This interface consolidates all the interfaces supported by contract-like
// components.
type ContractLike interface {
	Encapsulated
	Notarized
	Typed
}

// This interface consolidates all the interfaces supported by document-like
// components.
type DocumentLike interface {
	Encapsulated
	Typed
	Restricted
	Versioned
}

// This interface consolidates all the interfaces supported by certificate-like
// components.
type CertificateLike interface {
	Encapsulated
	Published
	Typed
	Restricted
	Versioned
}

// This interface consolidates all the interfaces supported by credentials-like
// components.
type CredentialsLike interface {
	Encapsulated
	Salted
	Typed
	Restricted
	Versioned
}

// This interface consolidates all the interfaces supported by
// security-module-like devices.
type SecurityModuleLike interface {
	Trusted
	Hardened
}

// This interface consolidates all the interfaces supported by notary-like
// agents.
type NotaryLike interface {
	Prudent
	Certified
}
