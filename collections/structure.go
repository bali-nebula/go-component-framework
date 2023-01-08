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

// STRUCTURE IMPLEMENTATION

// This constructor creates a new empty structure of key-value pairs.
func Structure() abs.StructureLike {
	return col.Array[abs.AssociationLike]([]abs.AssociationLike{})
}

// This constructor creates a new structure of key-value pairs from the
// specified array.
func StructureFromArray(array []abs.AssociationLike) abs.StructureLike {
	return col.Array[abs.AssociationLike](array)
}
