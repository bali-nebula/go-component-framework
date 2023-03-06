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
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	age "github.com/bali-nebula/go-component-framework/agents"
	bal "github.com/bali-nebula/go-component-framework/bali"
	doc "github.com/bali-nebula/go-component-framework/documents"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

const directory = "./"

func TestNotary(t *tes.T) {
	var module = age.SSMv1(directory)
	var notary = doc.Notary(module)

	var certificateV1 = notary.GenerateKey()
	fmt.Printf("certificate (v1): %v\n", string(bal.FormatDocument(certificateV1)))

	var citation = notary.GetCitation()
	fmt.Printf("citation: %v\n", string(bal.FormatDocument(citation)))

	var attributes = bal.Catalog(`[:]`)
	var type_ = bal.Component("/nebula/test/Document/v1.2.3")
	var tag = bal.NewTag()
	var version = bal.Version("v1.2")
	var permissions = bal.Moniker("/nebula/permissions/public/v1")
	var previous abs.CitationLike
	var document = doc.Document(attributes, type_, tag, version, permissions, previous)
	fmt.Printf("document: %v\n", string(bal.FormatDocument(document)))
	citation = notary.CiteDocument(document)
	ass.True(t, notary.CitationMatches(citation, document))

	var contract = notary.NotarizeDocument(document)
	fmt.Printf("contract: %v\n", string(bal.FormatDocument(contract)))
	ass.True(t, notary.SignatureMatches(contract, certificateV1.GetDocument().(abs.CertificateLike)))

	var certificateV2 = notary.RefreshKey()
	fmt.Printf("certificate (v2): %v\n", string(bal.FormatDocument(certificateV2)))

	var salt = bal.Binary(64)
	var credentials = notary.GenerateCredentials(salt)
	fmt.Printf("credentials: %v\n", string(bal.FormatDocument(credentials)))

	notary.ForgetKey()
}
