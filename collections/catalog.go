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
	return &association{key, value}
}

// This type defines the structure and methods associated with a key-value pair.
// This structure is used by the catalog type to maintain its associations.
type association struct {
	key   abs.Primitive
	value abs.ComponentLike
}

// BINDING INTERFACE

// This method returns the key for this association.
func (v *association) GetKey() abs.Primitive {
	return v.key
}

// This method returns the value for this association.
func (v *association) GetValue() abs.ComponentLike {
	return v.value
}

// This method sets the value of this association to a new value.
func (v *association) SetValue(value abs.ComponentLike) {
	v.value = value
}

// CATALOG IMPLEMENTATION

// This constructor creates a new empty catalog.
func Catalog() abs.CatalogLike {
	var v = col.Catalog[abs.Primitive, abs.ComponentLike]()
	return &catalog{v}
}

// This constructor creates a new catalog from the specified sequence of
// associations.
func CatalogFromSequence(sequence abs.Sequential[abs.AssociationLike]) abs.CatalogLike {
	// The implementation should be the following:
	//     var v = col.CatalogFromSequence[abs.Primitive, abs.ComponentLike](sequence)
	// Alas, the Go compiler does not correctly recognize that the two result
	// types are identical so we have to do this explicitly:
	var v = col.Catalog[abs.Primitive, abs.ComponentLike]()
	var iterator = AssociationIterator(sequence)
	for iterator.HasNext() {
		var association = iterator.GetNext()
		var key = association.GetKey()
		var value = association.GetValue()
		v.SetValue(key, value)
	}
	return &catalog{v}
}

// This type defines the structure and methods associated with a catalog of
// key-value pair associations.
type catalog struct {
	associations col.CatalogLike[abs.Primitive, abs.ComponentLike]
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this catalog is empty.
func (v *catalog) IsEmpty() bool {
	return v.associations.IsEmpty()
}

// This method returns the number of associations contained in this catalog.
func (v *catalog) GetSize() int {
	return v.associations.GetSize()
}

// This method returns all the associations in this catalog. The associations
// retrieved are in the same order as they are in the catalog.
func (v *catalog) AsArray() []abs.AssociationLike {
	// The implementation should be the following:
	//     return v.associations.AsArray()
	// Alas, the Go compiler does not correctly recognize that the two result
	// types are identical so we have to do this explicitly:
	var array = make([]abs.AssociationLike, v.associations.GetSize())
	var raw = v.associations.AsArray()
	for index, association := range raw {
		array[index] = association
	}
	return array
}

// ASSOCIATIVE INTERFACE

// This method returns the keys for this catalog.
func (v *catalog) GetKeys() abs.Sequential[abs.Primitive] {
	return v.associations.GetKeys()
}

// This method returns the values associated with the specified keys for this
// catalog. The values are returned in the same order as the keys in the
// catalog.
func (v *catalog) GetValues(keys abs.Sequential[abs.Primitive]) abs.Sequential[abs.ComponentLike] {
	return v.associations.GetValues(keys)
}

// This method returns the value that is associated with the specified key in
// this catalog.
func (v *catalog) GetValue(key abs.Primitive) abs.ComponentLike {
	return v.associations.GetValue(key)
}

// This method sets the value associated with the specified key to the
// specified value.
func (v *catalog) SetValue(key abs.Primitive, value abs.ComponentLike) {
	v.associations.SetValue(key, value)
}

// This method removes the association associated with the specified key from the
// catalog and returns it.
func (v *catalog) RemoveValue(key abs.Primitive) abs.ComponentLike {
	return v.associations.RemoveValue(key)
}

// This method removes the associations associated with the specified keys from
// the catalog and returns the removed values.
func (v *catalog) RemoveValues(keys abs.Sequential[abs.Primitive]) abs.Sequential[abs.ComponentLike] {
	return v.associations.RemoveValues(keys)
}

// This method removes all associations from this catalog.
func (v *catalog) RemoveAll() {
	v.associations.RemoveAll()
}

// SORTABLE INTERFACE

// This method sorts this catalog using the canonical rank function to compare
// the keys.
func (v *catalog) SortValues() {
	v.associations.SortValues()
}

// This method reverses the order of all associations in this catalog.
func (v *catalog) ReverseValues() {
	v.associations.ReverseValues()
}

// This method pseudo-randomly shuffles the order of all associations in this
// catalog.
func (v *catalog) ShuffleValues() {
	v.associations.ShuffleValues()
}
