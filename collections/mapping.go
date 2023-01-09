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
)

// MAPPING IMPLEMENTATION

// This type defines the structure and methods associated with a mapping of
// key-value associations.
type Mapping []abs.AssociationLike

// SEQUENTIAL INTERFACE

// This method determines whether or not this mapping is empty.
func (v Mapping) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of values contained in this mapping.
func (v Mapping) GetSize() int {
	return len(v)
}

// This method returns all the values in this mapping. The values retrieved are in
// the same order as they are in the mapping.
func (v Mapping) AsArray() []abs.AssociationLike {
	var length = len(v)
	var array = make([]abs.AssociationLike, length)
	copy(array, v)
	return array
}
