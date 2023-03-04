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
	bal "github.com/bali-nebula/go-component-framework/bali"
	col "github.com/bali-nebula/go-component-framework/collections"
	com "github.com/bali-nebula/go-component-framework/components"
)

// CERTIFICATE INTERFACE

// This constructor creates a new certificate.
func Certificate(
	key abs.PublicKey,
	algorithms abs.CatalogLike,
	tag abs.TagLike,
	version abs.VersionLike,
	previous abs.CitationLike,
) abs.CertificateLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(keyAttribute, com.Component(key))
	attributes.SetValue(algorithmsAttribute, com.Component(algorithms))

	// Create a new context.
	var context = com.Context()
	context.SetValue(typeAttribute, bal.ParseComponent("/nebula/types/Certificate/v1"))
	context.SetValue(tagAttribute, com.Component(tag))
	context.SetValue(versionAttribute, com.Component(version))
	context.SetValue(permissionsAttribute, bal.ParseComponent("/nebula/permissions/public/v1"))
	context.SetValue(previousAttribute, com.Component(previous))

	// Create a new certificate.
	return &certificate{com.ComponentWithContext(attributes, context)}
}

// CERTIFICATE IMPLEMENTATION

type certificate struct {
	abs.Encapsulated
}

func (v *certificate) GetAlgorithms() abs.CatalogLike {
	return v.ExtractCatalog().GetValue(algorithmsAttribute).ExtractCatalog()
}

func (v *certificate) GetKey() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(keyAttribute).ExtractBinary()
}

func (v *certificate) GetPermissions() abs.MonikerLike {
	return v.GetContext().GetValue(permissionsAttribute).ExtractMoniker()
}

func (v *certificate) GetPrevious() abs.CitationLike {
	return v.GetContext().GetValue(previousAttribute).ExtractCatalog().(abs.CitationLike)
}

func (v *certificate) GetTag() abs.TagLike {
	return v.GetContext().GetValue(tagAttribute).ExtractTag()
}

func (v *certificate) GetType() abs.TypeLike {
	return v.GetContext().GetValue(typeAttribute).(abs.TypeLike)
}

func (v *certificate) GetVersion() abs.VersionLike {
	return v.GetContext().GetValue(versionAttribute).ExtractVersion()
}
