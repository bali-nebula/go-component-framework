/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package documents

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	bal "github.com/bali-nebula/go-component-framework/bali"
	ele "github.com/bali-nebula/go-component-framework/elements"
	sm1 "github.com/bali-nebula/go-component-framework/security/v1"
	//sm2 "github.com/bali-nebula/go-component-framework/security/v2"
	//sm3 "github.com/bali-nebula/go-component-framework/security/v3"
	str "github.com/bali-nebula/go-component-framework/strings"
	col "github.com/craterdog/go-collection-framework/v2"
)

// NOTARY INTERFACE

// This constructor creates a new digital notary.
func Notary(hsm abs.SecurityModuleLike) abs.NotaryLike {
	// Initialize the modules catalog with all software security module versions.
	// The modules must be ordered with latest version first.
	var modules = col.Catalog[abs.VersionLike, abs.SecurityModuleLike]()
	//modules.SetValue(v3, sm3.SSM(""))
	//modules.SetValue(v2, sm2.SSM(""))
	modules.SetValue(v1, sm1.SSM(""))

	// Replace the corresponding SSM with the HSM (if provided)
	if hsm != nil {
		var version = str.VersionFromString(hsm.GetVersion())
		modules.SetValue(version, hsm)
	}
	return &notary{hsm, modules}
}

// NOTARY IMPLEMENTATION

// This type defines the structure and methods associated with a digital notary.
type notary struct {
	hsm     abs.SecurityModuleLike
	modules col.CatalogLike[abs.VersionLike, abs.SecurityModuleLike]
}

// These constants define the supported versions of the security protocol.
var (
	v1 = str.VersionFromString("1")
	v2 = str.VersionFromString("2")
	v3 = str.VersionFromString("3")
)

// This constant captures the algorithms used in this version of the protocol.
var algorithms = bal.Catalog(`[
    $digest: "SHA512"
    $signature: "ED25519"
]`)

// These constants define the attribute names for the standard attribues.
var (
	accountAttribute     = ele.SymbolFromString("account")
	algorithmsAttribute  = ele.SymbolFromString("algorithms")
	certificateAttribute = ele.SymbolFromString("certificate")
	digestAttribute      = ele.SymbolFromString("digest")
	documentAttribute    = ele.SymbolFromString("document")
	keyAttribute         = ele.SymbolFromString("key")
	permissionsAttribute = ele.SymbolFromString("permissions")
	previousAttribute    = ele.SymbolFromString("previous")
	protocolAttribute    = ele.SymbolFromString("protocol")
	saltAttribute        = ele.SymbolFromString("salt")
	signatureAttribute   = ele.SymbolFromString("signature")
	tagAttribute         = ele.SymbolFromString("tag")
	timestampAttribute   = ele.SymbolFromString("timestamp")
	typeAttribute        = ele.SymbolFromString("type")
	versionAttribute     = ele.SymbolFromString("version")
)

// PRUDENT INTERFACE

// This method generates a new private notary key and returns the corresponding
// public notary certificate.
func (v *notary) GenerateKey() abs.CertificateLike {
	var key = v.hsm.GenerateKeys()
	var tag = ele.TagOfSize(20)
	var version = v1
	var previous abs.CitationLike
	var certificate = Certificate(key, algorithms, tag, version, previous)
	return certificate
}

// This method activates a new private notary key and returns the a citation to
// the corresponding public notary certificate.
func (v *notary) ActivateKey(certificate abs.CertificateLike) abs.CitationLike {
	panic("Not yet implemented.")
}

// This method retrieves a citation to the public notary certificate for the
// current private notary key.
func (v *notary) GetCitation() abs.CitationLike {
	panic("Not yet implemented.")
}

// This method replaces an existing private notary key with a new one and
// returns the corresponding public notary certificate digitally signed by the
// old private notary key. The old private notary key is destroyed.
func (v *notary) RefreshKey() abs.CertificateLike {
	panic("Not yet implemented.")
}

// This method destroys any existing private notary key.
func (v *notary) ForgetKey() {
	panic("Not yet implemented.")
}

// CERTIFIED INTERFACE

// This method generates a new set of account credentials that can be used for
// authentication.
func (v *notary) GenerateCredentials(salt abs.Salt) abs.CredentialsLike {
	panic("Not yet implemented.")
}

// This method uses the current private notary key to notarized the specified
// document and returns the resulting contract.
func (v *notary) NotarizeDocument(document abs.DocumentLike) abs.ContractLike {
	panic("Not yet implemented.")
}

// This method determines whether or not the specified contract is valid using
// the specified public notary certificate.
func (v *notary) IsValid(contract abs.ContractLike, certificate abs.CertificateLike) bool {
	panic("Not yet implemented.")
}

// This method generates a citation to the specified document and returns that
// citation.
func (v *notary) CiteDocument(document abs.DocumentLike) abs.CitationLike {
	panic("Not yet implemented.")
}

// This method determines whether or not the specified document citation matches
// the specified document.
func (v *notary) CitationMatches(citation abs.CitationLike, document abs.DocumentLike) bool {
	panic("Not yet implemented.")
}

// PRIVATE METHODS

// This method creates a new document using the specified attributes and returns
// the new document.
func (v *notary) createDocument(
	type_ abs.MonikerLike,
	attributes abs.CatalogLike,
	tag abs.TagLike,
	version abs.VersionLike,
	permissions abs.MonikerLike,
	previous abs.CitationLike) abs.CatalogLike {
	panic("Not yet implemented.")
}

// This method creates a citation to the specified document and returns the
// citation.
func (v *notary) createCitation(document abs.DocumentLike) {
	panic("Not yet implemented.")
}

// This method creates a new digitally signed contract using the specified
// document and public notary certificate.
func (v *notary) createContract(document abs.DocumentLike, certificate abs.CertificateLike) {
	panic("Not yet implemented.")
}
