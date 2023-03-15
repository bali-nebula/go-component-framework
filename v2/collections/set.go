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

// This constructor creates a new empty component set.
func Set() abs.SetLike {
	var v = col.Set[abs.ComponentLike]()
	return &set{v}
}

// This constructor creates a new set from the specified sequence of components.
func SetFromSequence(sequence abs.Sequential[abs.ComponentLike]) abs.SetLike {
	var v = col.SetFromSequence[abs.ComponentLike](sequence)
	return &set{v}
}

// This type defines the structure and methods associated with a component set.
type set struct {
	components col.SetLike[abs.ComponentLike]
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this set is empty.
func (v *set) IsEmpty() bool {
	return v.components.IsEmpty()
}

// This method returns the number of components contained in this set.
func (v *set) GetSize() int {
	return v.components.GetSize()
}

// This method returns all the components in this set. The components retrieved are in
// the same order as they are in the set.
func (v *set) AsArray() []abs.ComponentLike {
	return v.components.AsArray()
}

// ACCESSIBLE INTERFACE

// This method retrieves from this set the component that is associated with the
// specified index.
func (v *set) GetValue(index int) abs.ComponentLike {
	return v.components.GetValue(index)
}

// This method retrieves from this set all components from the first index through
// the last index (inclusive).
func (v *set) GetValues(first int, last int) abs.Sequential[abs.ComponentLike] {
	return v.components.GetValues(first, last)
}

// SEARCHABLE INTERFACE

// This method returns the index of the FIRST occurrence of the specified component in
// this set, or zero if this set does not contain the component.
func (v *set) GetIndex(component abs.ComponentLike) int {
	return v.components.GetIndex(component)
}

// This method determines whether or not this set contains the specified component.
func (v *set) ContainsValue(component abs.ComponentLike) bool {
	return v.components.ContainsValue(component)
}

// This method determines whether or not this set contains ANY of the specified
// components.
func (v *set) ContainsAny(components abs.Sequential[abs.ComponentLike]) bool {
	return v.components.ContainsAny(components)
}

// This method determines whether or not this set contains ALL of the specified
// components.
func (v *set) ContainsAll(components abs.Sequential[abs.ComponentLike]) bool {
	return v.components.ContainsAll(components)
}

// FLEXIBLE INTERFACE

// This method adds the specified component to this set if it is not already a
// member of the set.
func (v *set) AddValue(component abs.ComponentLike) {
	v.components.AddValue(component)
}

// This method adds the specified components to this set if they are not already
// members of the set.
func (v *set) AddValues(components abs.Sequential[abs.ComponentLike]) {
	v.components.AddValues(components)
}

// This method removes the specified component from this set.
func (v *set) RemoveValue(component abs.ComponentLike) {
	v.components.RemoveValue(component)
}

// This method removes the specified components from this set.
func (v *set) RemoveValues(components abs.Sequential[abs.ComponentLike]) {
	v.components.RemoveValues(components)
}

// This method removes all components from this set.
func (v *set) RemoveAll() {
	v.components.RemoveAll()
}
