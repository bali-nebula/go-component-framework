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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// SET IMPLEMENTATION

// This constructor creates a new empty value set.
func Set() abs.SetLike {
	var v = col.Set[abs.ComponentLike]()
	return &set{v}
}

// This constructor creates a new set from the specified sequence of values.
func SetFromSequence(sequence abs.Sequential[abs.ComponentLike]) abs.SetLike {
	var v = col.SetFromSequence[abs.ComponentLike](sequence)
	return &set{v}
}

// This type defines the structure and methods associated with a value set.
type set struct {
	values col.SetLike[abs.ComponentLike]
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this set is empty.
func (v *set) IsEmpty() bool {
	return v.values.IsEmpty()
}

// This method returns the number of values contained in this set.
func (v *set) GetSize() int {
	return v.values.GetSize()
}

// This method returns all the values in this set. The values retrieved are in
// the same order as they are in the set.
func (v *set) AsArray() []abs.ComponentLike {
	return v.values.AsArray()
}

// ACCESSIBLE INTERFACE

// This method retrieves from this set the value that is associated with the
// specified index.
func (v *set) GetValue(index int) abs.ComponentLike {
	return v.values.GetValue(index)
}

// This method retrieves from this set all values from the first index through
// the last index (inclusive).
func (v *set) GetValues(first int, last int) abs.Sequential[abs.ComponentLike] {
	return v.values.GetValues(first, last)
}

// SEARCHABLE INTERFACE

// This method returns the index of the FIRST occurrence of the specified value in
// this set, or zero if this set does not contain the value.
func (v *set) GetIndex(value abs.ComponentLike) int {
	return v.values.GetIndex(value)
}

// This method determines whether or not this set contains the specified value.
func (v *set) ContainsValue(value abs.ComponentLike) bool {
	return v.values.ContainsValue(value)
}

// This method determines whether or not this set contains ANY of the specified
// values.
func (v *set) ContainsAny(values abs.Sequential[abs.ComponentLike]) bool {
	return v.values.ContainsAny(values)
}

// This method determines whether or not this set contains ALL of the specified
// values.
func (v *set) ContainsAll(values abs.Sequential[abs.ComponentLike]) bool {
	return v.values.ContainsAll(values)
}

// FLEXIBLE INTERFACE

// This method adds the specified value to this set if it is not already a
// member of the set.
func (v *set) AddValue(value abs.ComponentLike) {
	v.values.AddValue(value)
}

// This method adds the specified values to this set if they are not already
// members of the set.
func (v *set) AddValues(values abs.Sequential[abs.ComponentLike]) {
	v.values.AddValues(values)
}

// This method removes the specified value from this set.
func (v *set) RemoveValue(value abs.ComponentLike) {
	v.values.RemoveValue(value)
}

// This method removes the specified values from this set.
func (v *set) RemoveValues(values abs.Sequential[abs.ComponentLike]) {
	v.values.RemoveValues(values)
}

// This method removes all values from this set.
func (v *set) RemoveAll() {
	v.values.RemoveAll()
}
