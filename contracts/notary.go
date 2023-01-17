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
	sm1 "github.com/bali-nebula/go-component-framework/contracts/v1"
	//sm2 "github.com/bali-nebula/go-component-framework/contracts/v2"
	//sm3 "github.com/bali-nebula/go-component-framework/contracts/v3"
	ele "github.com/bali-nebula/go-component-framework/elements"
	col "github.com/craterdog/go-collection-framework/v2"
)

// CONSTANT DEFINITIONS

const (
	v1 = ele.Version("v1")
	v2 = ele.Version("v2")
	v3 = ele.Version("v3")
)

// TYPE DEFINITIONS

type (
	Document    any
	Citation    any
	Certificate any
	Contract    any
	Credentials any
)

// NOTARY IMPLEMENTATION

// This constructor creates a new digital notary.
func Notary() abs.NotaryLike {
	var protocols = col.Catalog[ele.Version, abs.SecurityModuleLike]()
	// Must be ordered with newest version first.
	//protocols.SetValue(v3, sm3.SSM())
	//protocols.SetValue(v2, sm2.SSM())
	protocols.SetValue(v1, sm1.SSM())
	return &notary{protocols}
}

// This type defines the structure and methods associated with a digital notary.
type notary struct {
	protocols col.CatalogLike[ele.Version, abs.SecurityModuleLike]
}

// AUTHORITATIVE INTERFACE

// This method generates a new private notary key and returns the corresponding
// public notary certificate.
func (v *notary) GenerateKey() Certificate {
	panic("Not yet implemented.")
}

// This method activates a new private notary key and returns the a citation to
// the corresponding public notary certificate.
func (v *notary) ActivateKey(certificate Certificate) Citation {
	panic("Not yet implemented.")
}

// This method retrieves a citation to the public notary certificate for the
// current private notary key.
func (v *notary) GetCitation() Citation {
	panic("Not yet implemented.")
}

// This method generates a new set of account credentials that can be used for
// authentication.
func (v *notary) GenerateCredentials(salt ele.Tag) Credentials {
	panic("Not yet implemented.")
}

// This method uses the current private notary key to notarized the specified
// document and returns the resulting contract.
func (v *notary) NotarizeDocument(document Document) Contract {
	panic("Not yet implemented.")
}

// This method determines whether or not the specified contract is valid using
// the specified public notary certificate.
func (v *notary) IsValid(contract Contract, certificate Certificate) bool {
	panic("Not yet implemented.")
}

// This method generates a citation to the specified document and returns that
// citation.
func (v *notary) CiteDocument(document Document) Citation {
	panic("Not yet implemented.")
}

// This method determines whether or not the specified document citation matches
// the specified document.
func (v *notary) CitationMatches(citation Citation, document Document) bool {
	panic("Not yet implemented.")
}

// This method replaces an existing private notary key with a new one and
// returns the corresponding public notary certificate digitally signed by the
// old private notary key. The old private notary key is destroyed.
func (v *notary) RefreshKey() Certificate {
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
	type_ ele.Moniker,
	attributes abs.Mapping,
	tag ele.Tag,
	version ele.Version,
	permissions abs.Mapping,
	previous Citation) {
	panic("Not yet implemented.")
}

// This method creates a citation to the specified document and returns the
// citation.
func (v *notary) createCitation(document Document) {
	panic("Not yet implemented.")
}

// This method creates a new certificate containing the specified public key
// and associated attributes.
func (v *notary) createCertificate(
	publicKey PublicKey,
	tag ele.Tag,
	version ele.Version,
	previous Citation) {
	panic("Not yet implemented.")
}

// This method creates a new digitally signed contract using the specified
// document and public notary certificate.
func (v *notary) createContract(document Document, certificate Certificate) {
	panic("Not yet implemented.")
}
