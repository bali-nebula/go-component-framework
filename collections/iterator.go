/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
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
	col "github.com/craterdog/go-collection-framework/v2"
)

// COMPONENT ITERATOR IMPLEMENTATION

// This constructor creates a new instance of a components iterator that can be
// used to traverse the components in the specified sequence.
func ComponentIterator(sequence abs.Sequential[abs.ComponentLike]) abs.ComponentIteratorLike {
	var v = col.Iterator[abs.ComponentLike](sequence)
	return &components{v}
}

// This type defines the structure and methods for a components iterator. The
// iterator operates on a sequence of components.
type components struct {
	abs.ComponentIteratorLike
}

// ASSOCIATION ITERATOR IMPLEMENTATION

// This constructor creates a new instance of an associations iterator that can be
// used to traverse the associations in the specified sequence.
func AssociationIterator(sequence abs.Sequential[abs.AssociationLike]) abs.AssociationIteratorLike {
	var v = col.Iterator[abs.AssociationLike](sequence)
	return &associations{v}
}

// This type defines the structure and methods for an associations iterator. The
// iterator operates on a sequence of associations.
type associations struct {
	abs.AssociationIteratorLike
}
