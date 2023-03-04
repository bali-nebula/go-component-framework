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
	ele "github.com/bali-nebula/go-component-framework/elements"
)

// CREDENTIALS INTERFACE

// This constructor creates a new credentials.
func Credentials(
	salt abs.BinaryLike,
) abs.CredentialsLike {

	// Create a new catalog for the attributes.
	var attributes = col.Catalog()
	attributes.SetValue(saltAttribute, com.Component(salt))

	// Create a new context.
	var context = com.Context()
	context.SetValue(typeAttribute, bal.ParseComponent("/nebula/types/Credentials/v1"))
	context.SetValue(tagAttribute, com.Component(ele.TagOfSize(20)))
	context.SetValue(versionAttribute, com.Component(v1))
	context.SetValue(permissionsAttribute, bal.ParseComponent("/nebula/permissions/private/v1"))

	// Create a new credentials.
	return &credentials{com.ComponentWithContext(attributes, context)}
}

// CREDENTIALS IMPLEMENTATION

type credentials struct {
	abs.Encapsulated
}

func (v *credentials) GetSalt() abs.BinaryLike {
	return v.ExtractCatalog().GetValue(saltAttribute).ExtractBinary()
}

func (v *credentials) GetPermissions() abs.MonikerLike {
	return v.GetContext().GetValue(permissionsAttribute).ExtractMoniker()
}

func (v *credentials) GetPrevious() abs.CitationLike {
	return v.GetContext().GetValue(previousAttribute).ExtractCatalog().(abs.CitationLike)
}

func (v *credentials) GetTag() abs.TagLike {
	return v.GetContext().GetValue(tagAttribute).ExtractTag()
}

func (v *credentials) GetType() abs.TypeLike {
	return v.GetContext().GetValue(typeAttribute).(abs.TypeLike)
}

func (v *credentials) GetVersion() abs.VersionLike {
	return v.GetContext().GetValue(versionAttribute).ExtractVersion()
}
