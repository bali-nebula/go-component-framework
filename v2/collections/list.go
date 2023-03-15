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

// LIST IMPLEMENTATION

// This constructor creates a new empty list.
func List() abs.ListLike {
	var v = col.List[abs.ComponentLike]()
	return &list{v}
}

// This constructor creates a new list from the specified sequence. The list
// uses the natural compare function.
func ListFromSequence(sequence abs.Sequential[abs.ComponentLike]) abs.ListLike {
	var v = col.ListFromSequence[abs.ComponentLike](sequence)
	return &list{v}
}

// This type defines the structure and methods associated with a list of
// components.
type list struct {
	components col.ListLike[abs.ComponentLike]
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this list is empty.
func (v *list) IsEmpty() bool {
	return v.components.IsEmpty()
}

// This method returns the number of components contained in this list.
func (v *list) GetSize() int {
	return v.components.GetSize()
}

// This method returns all the components in this list. The components retrieved are in
// the same order as they are in the list.
func (v *list) AsArray() []abs.ComponentLike {
	return v.components.AsArray()
}

// ACCESSIBLE INTERFACE

// This method retrieves from this list the component that is associated with the
// specified index.
func (v *list) GetValue(index int) abs.ComponentLike {
	return v.components.GetValue(index)
}

// This method retrieves from this list all components from the first index through
// the last index (inclusive).
func (v *list) GetValues(first int, last int) abs.Sequential[abs.ComponentLike] {
	return v.components.GetValues(first, last)
}

// UPDATABLE INTERFACE

// This method sets the component in this list that is associated with the specified
// index to be the specified component.
func (v *list) SetValue(index int, component abs.ComponentLike) {
	v.components.SetValue(index, component)
}

// This method sets the components in this list starting with the specified index
// to the specified components.
func (v *list) SetValues(index int, components abs.Sequential[abs.ComponentLike]) {
	v.components.SetValues(index, components)
}

// SEARCHABLE INTERFACE

// This method returns the index of the FIRST occurrence of the specified component in
// this list, or zero if this list does not contain the component.
func (v *list) GetIndex(component abs.ComponentLike) int {
	return v.components.GetIndex(component)
}

// This method determines whether or not this list contains the specified component.
func (v *list) ContainsValue(component abs.ComponentLike) bool {
	return v.components.ContainsValue(component)
}

// This method determines whether or not this list contains ANY of the specified
// components.
func (v *list) ContainsAny(components abs.Sequential[abs.ComponentLike]) bool {
	return v.components.ContainsAny(components)
}

// This method determines whether or not this list contains ALL of the specified
// components.
func (v *list) ContainsAll(components abs.Sequential[abs.ComponentLike]) bool {
	return v.components.ContainsAll(components)
}

// MALLEABLE INTERFACE

// This method appends the specified component to the end of this list.
func (v *list) AddValue(component abs.ComponentLike) {
	v.components.AddValue(component)
}

// This method appends the specified components to the end of this list.
func (v *list) AddValues(components abs.Sequential[abs.ComponentLike]) {
	v.components.AddValues(components)
}

// This method inserts the specified component into this list in the specified
// slot between existing components.
func (v *list) InsertValue(slot int, component abs.ComponentLike) {
	v.components.InsertValue(slot, component)
}

// This method inserts the specified components into this list in the specified
// slot between existing components.
func (v *list) InsertValues(slot int, components abs.Sequential[abs.ComponentLike]) {
	v.components.InsertValues(slot, components)
}

// This method removes the component at the specified index from this list. The
// removed component is returned.
func (v *list) RemoveValue(index int) abs.ComponentLike {
	return v.components.RemoveValue(index)
}

// This method removes the components in the specified index range from this list.
// The removed components are returned.
func (v *list) RemoveValues(first int, last int) abs.Sequential[abs.ComponentLike] {
	return v.components.RemoveValues(first, last)
}

// This method removes all components from this list.
func (v *list) RemoveAll() {
	v.components.RemoveAll()
}

// SORTABLE INTERFACE

// This method sorts the components in this list using the natural rank function.
func (v *list) SortValues() {
	v.components.SortValues()
}

// This method reverses the order of all components in this list.
func (v *list) ReverseValues() {
	v.components.ReverseValues()
}

// This method pseudo-randomly shuffles the components in this list.
func (v *list) ShuffleValues() {
	v.components.ShuffleValues()
}
