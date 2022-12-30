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
	col "github.com/craterdog/go-collection-framework"
)

// ASSOCIATION IMPLEMENTATION

// This constructor creates a new association with the specified key and value.
func Association(key abs.Primitive, value abs.ComponentLike) abs.AssociationLike {
	return col.Association[abs.Primitive, abs.ComponentLike](key, value)
}

// CATALOG IMPLEMENTATION

// This constructor creates a new empty catalog.
func Catalog() abs.CatalogLike {
	return col.Catalog[abs.Primitive, abs.ComponentLike]()
}

// This constructor creates a new catalog from the specified array of
// components.
func CatalogFromArray(array []abs.Binding) abs.CatalogLike {
	return col.CatalogFromArray[abs.Primitive, abs.ComponentLike](array)
}

// This constructor creates a new catalog from the specified sequence of
// components.
func CatalogFromSequence(sequence col.Sequential[abs.Binding]) abs.CatalogLike {
	return col.CatalogFromSequence[abs.Primitive, abs.ComponentLike](sequence)
}
