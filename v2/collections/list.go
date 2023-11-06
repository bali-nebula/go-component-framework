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
// values.
type list struct {
	values col.ListLike[abs.ComponentLike]
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this list is empty.
func (v *list) IsEmpty() bool {
	return v.values.IsEmpty()
}

// This method returns the number of values contained in this list.
func (v *list) GetSize() int {
	return v.values.GetSize()
}

// This method returns all the values in this list. The values retrieved are in
// the same order as they are in the list.
func (v *list) AsArray() []abs.ComponentLike {
	return v.values.AsArray()
}

// ACCESSIBLE INTERFACE

// This method retrieves from this list the component that is associated with the
// specified index.
func (v *list) GetValue(index int) abs.ComponentLike {
	return v.values.GetValue(index)
}

// This method retrieves from this list all values from the first index through
// the last index (inclusive).
func (v *list) GetValues(first int, last int) abs.Sequential[abs.ComponentLike] {
	return v.values.GetValues(first, last)
}

// UPDATABLE INTERFACE

// This method sets the component in this list that is associated with the specified
// index to be the specified component.
func (v *list) SetValue(index int, value abs.ComponentLike) {
	v.values.SetValue(index, value)
}

// This method sets the values in this list starting with the specified index
// to the specified values.
func (v *list) SetValues(index int, values abs.Sequential[abs.ComponentLike]) {
	v.values.SetValues(index, values)
}

// SEARCHABLE INTERFACE

// This method returns the index of the FIRST occurrence of the specified component in
// this list, or zero if this list does not contain the component.
func (v *list) GetIndex(value abs.ComponentLike) int {
	return v.values.GetIndex(value)
}

// This method determines whether or not this list contains the specified component.
func (v *list) ContainsValue(value abs.ComponentLike) bool {
	return v.values.ContainsValue(value)
}

// This method determines whether or not this list contains ANY of the specified
// values.
func (v *list) ContainsAny(values abs.Sequential[abs.ComponentLike]) bool {
	return v.values.ContainsAny(values)
}

// This method determines whether or not this list contains ALL of the specified
// values.
func (v *list) ContainsAll(values abs.Sequential[abs.ComponentLike]) bool {
	return v.values.ContainsAll(values)
}

// MALLEABLE INTERFACE

// This method appends the specified component to the end of this list.
func (v *list) AddValue(value abs.ComponentLike) {
	v.values.AddValue(value)
}

// This method appends the specified values to the end of this list.
func (v *list) AddValues(values abs.Sequential[abs.ComponentLike]) {
	v.values.AddValues(values)
}

// This method inserts the specified component into this list in the specified
// slot between existing values.
func (v *list) InsertValue(slot int, value abs.ComponentLike) {
	v.values.InsertValue(slot, value)
}

// This method inserts the specified values into this list in the specified
// slot between existing values.
func (v *list) InsertValues(slot int, values abs.Sequential[abs.ComponentLike]) {
	v.values.InsertValues(slot, values)
}

// This method removes the component at the specified index from this list. The
// removed component is returned.
func (v *list) RemoveValue(index int) abs.ComponentLike {
	return v.values.RemoveValue(index)
}

// This method removes the values in the specified index range from this list.
// The removed values are returned.
func (v *list) RemoveValues(first int, last int) abs.Sequential[abs.ComponentLike] {
	return v.values.RemoveValues(first, last)
}

// This method removes all values from this list.
func (v *list) RemoveAll() {
	v.values.RemoveAll()
}

// SORTABLE INTERFACE

// This method sorts the values in this list using the natural rank function.
func (v *list) SortValues() {
	v.values.SortValues()
}

// This method reverses the order of all values in this list.
func (v *list) ReverseValues() {
	v.values.ReverseValues()
}

// This method pseudo-randomly shuffles the values in this list.
func (v *list) ShuffleValues() {
	v.values.ShuffleValues()
}
