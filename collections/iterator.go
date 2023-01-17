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
	iterator col.IteratorLike[abs.ComponentLike]
}

// This method returns the current slot between components that this iterator is
// currently locked into.
func (v *components) GetSlot() int {
	return v.iterator.GetSlot()
}

// This method moves this iterator to the specified slot between components.
func (v *components) ToSlot(slot int) {
	v.iterator.ToSlot(slot)
}

// This method moves this iterator to the slot before the first component.
func (v *components) ToStart() {
	v.iterator.ToStart()
}

// This method moves this iterator to the slot after the last component.
func (v *components) ToEnd() {
	v.iterator.ToEnd()
}

// This method determines whether or not there is a component before the current
// slot.
func (v *components) HasPrevious() bool {
	return v.iterator.HasPrevious()
}

// This method retrieves the component before the current slot.
func (v *components) GetPrevious() abs.ComponentLike {
	return v.iterator.GetPrevious()
}

// This method determines whether or not there is a component after the current
// slot.
func (v *components) HasNext() bool {
	return v.iterator.HasNext()
}

// This method retrieves the component after the current slot.
func (v *components) GetNext() abs.ComponentLike {
	return v.iterator.GetNext()
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
	iterator col.IteratorLike[abs.AssociationLike]
}

// This method returns the current slot between associations that this iterator is
// currently locked into.
func (v *associations) GetSlot() int {
	return v.iterator.GetSlot()
}

// This method moves this iterator to the specified slot between associations.
func (v *associations) ToSlot(slot int) {
	v.iterator.ToSlot(slot)
}

// This method moves this iterator to the slot before the first association.
func (v *associations) ToStart() {
	v.iterator.ToStart()
}

// This method moves this iterator to the slot after the last association.
func (v *associations) ToEnd() {
	v.iterator.ToEnd()
}

// This method determines whether or not there is an association before the current
// slot.
func (v *associations) HasPrevious() bool {
	return v.iterator.HasPrevious()
}

// This method retrieves the association before the current slot.
func (v *associations) GetPrevious() abs.AssociationLike {
	return v.iterator.GetPrevious()
}

// This method determines whether or not there is an association after the current
// slot.
func (v *associations) HasNext() bool {
	return v.iterator.HasNext()
}

// This method retrieves the association after the current slot.
func (v *associations) GetNext() abs.AssociationLike {
	return v.iterator.GetNext()
}
