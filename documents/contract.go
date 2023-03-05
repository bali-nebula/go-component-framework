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
	col "github.com/bali-nebula/go-component-framework/collections"
	com "github.com/bali-nebula/go-component-framework/components"
)

// CONTRACT INTERFACE

// This constructor creates a new contract.
func Contract(
	document abs.DocumentLike,
	account abs.TagLike,
	protocol abs.VersionLike,
	certificate abs.CitationLike,
) abs.ContractLike {
	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(documentAttribute, com.Component(document))
	attributes.SetValue(accountAttribute, com.Component(account))
	attributes.SetValue(protocolAttribute, com.Component(protocol))
	attributes.SetValue(certificateAttribute, com.Component(certificate))

	// Create a new context for the type.
	var context = com.Context()
	context.SetValue(typeAttribute, bal.ParseComponent("/nebula/types/Contract/v1"))

	// Create a new contract.
	return &contract{com.ComponentWithContext(attributes, context)}
}

// CONTRACT IMPLEMENTATION

type contract struct {
	abs.Encapsulated
}

func (v *contract) GetAccount() abs.TagLike {
	return v.GetContext().GetValue(accountAttribute).ExtractTag()
}

func (v *contract) GetCertificate() abs.CitationLike {
	return v.GetContext().GetValue(versionAttribute).(abs.CitationLike)
}

func (v *contract) GetDocument() abs.DocumentLike {
	return v.GetContext().GetValue(documentAttribute).(abs.DocumentLike)
}

func (v *contract) GetProtocol() abs.VersionLike {
	return v.GetContext().GetValue(protocolAttribute).ExtractVersion()
}

func (v *contract) GetSignature() abs.BinaryLike {
	return v.GetContext().GetValue(signatureAttribute).ExtractBinary()
}

func (v *contract) SetSignature(signature abs.BinaryLike) {
	v.GetContext().SetValue(signatureAttribute, bal.Component(signature))
}

func (v *contract) GetTimestamp() abs.MomentLike {
	return v.GetContext().GetValue(timestampAttribute).ExtractMoment()
}

func (v *contract) GetType() abs.MonikerLike {
	return v.GetContext().GetValue(typeAttribute).GetEntity().(abs.MonikerLike)
}
