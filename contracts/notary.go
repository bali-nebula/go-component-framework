/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package contracts

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	sm1 "github.com/bali-nebula/go-component-framework/security/v1"
	//sm2 "github.com/bali-nebula/go-component-framework/security/v2"
	//sm3 "github.com/bali-nebula/go-component-framework/security/v3"
	str "github.com/bali-nebula/go-component-framework/strings"
	col "github.com/craterdog/go-collection-framework/v2"
)

// CONSTANT DEFINITIONS

const (
	protocolAttribute   = "protocol"
	tagAttribute        = "tag"
	versionAttribute    = "version"
	digestAttribute     = "digest"
	publicKeyAttribute  = "publicKey"
	privateKeyAttribute = "privateKey"
)

var (
	v1 = str.VersionFromString("1")
	v2 = str.VersionFromString("2")
	v3 = str.VersionFromString("3")
)

// NOTARY IMPLEMENTATION

// This constructor creates a new digital notary.
func Notary() abs.NotaryLike {
	var modules = col.Catalog[abs.VersionLike, abs.SecurityModuleLike]()
	// Must be ordered with latest version first.
	//modules.SetValue(v3, sm3.SSM())
	//modules.SetValue(v2, sm2.SSM())
	modules.SetValue(v1, sm1.SSM(""))
	return &notary{modules}
}

// This type defines the structure and methods associated with a digital notary.
type notary struct {
	modules col.CatalogLike[abs.VersionLike, abs.SecurityModuleLike]
}

// AUTHORITATIVE INTERFACE

// This method generates a new private notary key and returns the corresponding
// public notary certificate.
func (v *notary) GenerateKey() abs.Certificate {
	panic("Not yet implemented.")
}

// This method activates a new private notary key and returns the a citation to
// the corresponding public notary certificate.
func (v *notary) ActivateKey(certificate abs.Certificate) abs.Citation {
	panic("Not yet implemented.")
}

// This method retrieves a citation to the public notary certificate for the
// current private notary key.
func (v *notary) GetCitation() abs.Citation {
	panic("Not yet implemented.")
}

// This method generates a new set of account credentials that can be used for
// authentication.
func (v *notary) GenerateCredentials(salt abs.Salt) abs.Credentials {
	panic("Not yet implemented.")
}

// This method uses the current private notary key to notarized the specified
// document and returns the resulting contract.
func (v *notary) NotarizeDocument(document abs.Document) abs.Contract {
	panic("Not yet implemented.")
}

// This method determines whether or not the specified contract is valid using
// the specified public notary certificate.
func (v *notary) IsValid(contract abs.Contract, certificate abs.Certificate) bool {
	panic("Not yet implemented.")
}

// This method generates a citation to the specified document and returns that
// citation.
func (v *notary) CiteDocument(document abs.Document) abs.Citation {
	panic("Not yet implemented.")
}

// This method determines whether or not the specified document citation matches
// the specified document.
func (v *notary) CitationMatches(citation abs.Citation, document abs.Document) bool {
	panic("Not yet implemented.")
}

// This method replaces an existing private notary key with a new one and
// returns the corresponding public notary certificate digitally signed by the
// old private notary key. The old private notary key is destroyed.
func (v *notary) RefreshKey() abs.Certificate {
	panic("Not yet implemented.")
}

// This method destroys any existing private notary key.
func (v *notary) ForgetKey() {
	panic("Not yet implemented.")
}

// PRIVATE METHODS

// This method creates a new document using the specified attributes and returns
// the new document.
func (v *notary) createDocument(
	type_ abs.MonikerLike,
	attributes abs.MappingLike,
	tag abs.TagLike,
	version abs.VersionLike,
	permissions abs.MappingLike,
	previous abs.Citation) {
	panic("Not yet implemented.")
}

// This method creates a citation to the specified document and returns the
// citation.
func (v *notary) createCitation(document abs.Document) {
	panic("Not yet implemented.")
}

// This method creates a new certificate containing the specified public key
// and associated attributes.
func (v *notary) createCertificate(
	publicKey abs.PublicKey,
	tag abs.TagLike,
	version abs.VersionLike,
	previous abs.Citation) {
	panic("Not yet implemented.")
}

// This method creates a new digitally signed contract using the specified
// document and public notary certificate.
func (v *notary) createContract(document abs.Document, certificate abs.Certificate) {
	panic("Not yet implemented.")
}
