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
	com "github.com/bali-nebula/go-component-framework/components"
)

// DOCUMENT INTERFACE

// This constructor creates a new document.
func Document(
	attributes abs.CatalogLike,
	type_ abs.ComponentLike,
	tag abs.TagLike,
	version abs.VersionLike,
	permissions abs.MonikerLike,
	previous abs.CitationLike,
) abs.DocumentLike {

	// Create a new context.
	var context = com.Context()
	context.SetValue(typeAttribute, type_)
	context.SetValue(tagAttribute, com.Component(tag))
	context.SetValue(versionAttribute, com.Component(version))
	context.SetValue(permissionsAttribute, com.Component(permissions))
	context.SetValue(previousAttribute, com.Component(previous))

	// Create a new document.
	return &document{com.ComponentWithContext(attributes, context)}
}

// DOCUMENT IMPLEMENTATION

type document struct {
	abs.Encapsulated
}

func (v *document) GetPermissions() abs.MonikerLike {
	return v.GetContext().GetValue(permissionsAttribute).ExtractMoniker()
}

func (v *document) GetPrevious() abs.CitationLike {
	return v.GetContext().GetValue(previousAttribute).ExtractCatalog().(abs.CitationLike)
}

func (v *document) GetTag() abs.TagLike {
	return v.GetContext().GetValue(tagAttribute).ExtractTag()
}

func (v *document) GetType() abs.MonikerLike {
	return v.GetContext().GetValue(typeAttribute).GetEntity().(abs.MonikerLike)
}

func (v *document) GetVersion() abs.VersionLike {
	return v.GetContext().GetValue(versionAttribute).ExtractVersion()
}
