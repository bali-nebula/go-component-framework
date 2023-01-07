/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	ent "github.com/bali-nebula/go-component-framework/entities"
	col "github.com/craterdog/go-collection-framework"
)

// ASSOCIATION IMPLEMENTATION

// This constructor creates a new association key-value pair.
func Association(key abs.Primitive, value abs.ComponentLike) abs.AssociationLike {
	return col.Association(key, value)
}

// CATALOG IMPLEMENTATION

// This constructor creates a new empty catalog.
func Catalog() abs.CatalogLike {
	return col.Catalog[abs.Primitive, abs.ComponentLike]().(abs.CatalogLike)
}

// This constructor creates a new catalog from the specified sequence of
// components.
func CatalogFromStructure(structure abs.StructureLike) abs.CatalogLike {
	var catalog = Catalog()
	var keys = structure.GetKeys()
	var iterator = ent.Iterator(keys)
	for iterator.HasNext() {
		var key = iterator.GetNext()
		var value = structure.GetValue(key)
		catalog.SetValue(key, value)
	}
	return catalog
}
