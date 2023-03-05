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

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all notarized components.
type Notarized interface {
	GetDocument() DocumentLike
	GetTimestamp() MomentLike
	GetAccount() TagLike
	GetProtocol() VersionLike
	GetCertificate() CitationLike
	GetSignature() BinaryLike
	SetSignature(signature BinaryLike)
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
	GetType() MonikerLike
}

// This interface defines the methods supported by all versioned components.
type Versioned interface {
	GetTag() TagLike
	GetVersion() VersionLike
	GetPrevious() CitationLike
}

// This interface defines the methods supported by all prudent notary
// agents.
type Prudent interface {
	GenerateKey() ContractLike
	GetCitation() CitationLike
	RefreshKey() ContractLike
	ForgetKey()
}

// This interface defines the methods supported by all certified notary
// agents.
type Certified interface {
	GenerateCredentials(salt BinaryLike) CredentialsLike
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

// This interface consolidates all the interfaces supported by notary-like
// agents.
type NotaryLike interface {
	Prudent
	Certified
}
