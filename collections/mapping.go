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

// This constructor creates a new empty mapping of key-value pairs.
func Mapping() abs.MappingLike {
	var array = make([]abs.AssociationLike, 0)
	return &mapping{array}
}

// This constructor creates a new mapping of key-value pairs from the specified
// array of associations.
func MappingFromSequence(sequence abs.Sequential[abs.AssociationLike]) abs.MappingLike {
	var array = sequence.AsArray()
	return &mapping{array}
}

// This type defines the structure and methods associated with a mapping of
// key-value pairs.
type mapping struct {
	array []abs.AssociationLike
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this mapping of key-value pairs is
// empty.
func (v *mapping) IsEmpty() bool {
	return len(v.array) == 0
}

// This method returns the number of key-value pairs maintained by this mapping.
func (v *mapping) GetSize() int {
	return len(v.array)
}

// This method returns an array of the key-value pairs maintained by this
// mapping. The order of the values retrieved is preserved in the array.
func (v *mapping) AsArray() []abs.AssociationLike {
	var length = len(v.array)
	var array = make([]abs.AssociationLike, length)
	copy(array, v.array)
	return array
}
