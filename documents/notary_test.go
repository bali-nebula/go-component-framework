/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package documents_test

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	age "github.com/bali-nebula/go-component-framework/agents"
	bal "github.com/bali-nebula/go-component-framework/bali"
	doc "github.com/bali-nebula/go-component-framework/documents"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	tes "testing"
)

const directory = "./"

func TestNotary(t *tes.T) {
	var module = age.SSMv1(directory)
	var notary = doc.Notary(module)

	var certificateV1 = notary.GenerateKey()
	osx.WriteFile("./examples/certificateV1.bali", bal.FormatDocument(certificateV1), 0600)

	var citation = notary.GetCitation()
	osx.WriteFile("./examples/citation.bali", bal.FormatDocument(citation), 0600)

	var attributes = bal.Catalog(`[:]`)
	var type_ = bal.Component("/nebula/examples/Document/v1.2.3")
	var tag = bal.NewTag()
	var version = bal.Version("v1.2")
	var permissions = bal.Moniker("/nebula/permissions/public/v1")
	var previous abs.CitationLike
	var document = doc.Document(attributes, type_, tag, version, permissions, previous)
	osx.WriteFile("./examples/document.bali", bal.FormatDocument(document), 0600)
	citation = notary.CiteDocument(document)
	ass.True(t, notary.CitationMatches(citation, document))

	var contract = notary.NotarizeDocument(document)
	osx.WriteFile("./examples/contract.bali", bal.FormatDocument(contract), 0600)
	ass.True(t, notary.SignatureMatches(contract, certificateV1.GetDocument().(abs.CertificateLike)))

	var certificateV2 = notary.RefreshKey()
	osx.WriteFile("./examples/certificateV2.bali", bal.FormatDocument(certificateV2), 0600)

	var salt = bal.Binary(64)
	var credentials = notary.GenerateCredentials(salt)
	osx.WriteFile("./examples/credentials.bali", bal.FormatDocument(credentials), 0600)

	notary.ForgetKey()
}
