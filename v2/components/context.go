/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package components

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// PARAMETER IMPLEMENTATION

// This constructor creates a new parameter with the specified key and value.
func Parameter(key abs.SymbolLike, value abs.ComponentLike) abs.ParameterLike {
	return &parameter{key, value}
}

// This type defines the structure and methods associated with a key-value pair.
// This structure is used by the context type to maintain its parameters.
type parameter struct {
	key   abs.SymbolLike
	value abs.ComponentLike
}

// BINDING INTERFACE

// This method returns the key for this parameter.
func (v *parameter) GetKey() abs.SymbolLike {
	return v.key
}

// This method returns the value for this parameter.
func (v *parameter) GetValue() abs.ComponentLike {
	return v.value
}

// This method sets the value of this parameter to a new value.
func (v *parameter) SetValue(value abs.ComponentLike) {
	v.value = value
}

// PARAMETER ITERATOR IMPLEMENTATION

// This constructor creates a new instance of an parameters iterator that can be
// used to traverse the parameters in the specified sequence.
func ParameterIterator(sequence abs.Sequential[abs.ParameterLike]) abs.ParameterIteratorLike {
	var v = col.Iterator[abs.ParameterLike](sequence)
	return &parameters{v}
}

// This type defines the structure and methods for an parameters iterator. The
// iterator operates on a sequence of parameters.
type parameters struct {
	abs.ParameterIteratorLike
}

// CONTEXT IMPLEMENTATION

// This constructor creates a new empty context.
func Context() abs.ContextLike {
	var v = col.Catalog[abs.SymbolLike, abs.ComponentLike]()
	return &context{v}
}

// This constructor creates a new context from the specified sequence of
// parameters.
func ContextFromSequence(sequence abs.Sequential[abs.ParameterLike]) abs.ContextLike {
	// The implementation should be the following:
	//     var v = col.CatalogFromSequence[abs.SymbolLike, abs.ComponentLike](sequence)
	// Alas, the Go compiler does not correctly recognize that the two result
	// types are identical so we have to do this explicitly:
	var v = col.Catalog[abs.SymbolLike, abs.ComponentLike]()
	var iterator = ParameterIterator(sequence)
	for iterator.HasNext() {
		var parameter = iterator.GetNext()
		var key = parameter.GetKey()
		var value = parameter.GetValue()
		v.SetValue(key, value)
	}
	return &context{v}
}

// This type defines the structure and methods associated with a context of
// key-value pair parameters.
type context struct {
	parameters col.CatalogLike[abs.SymbolLike, abs.ComponentLike]
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this context is empty.
func (v *context) IsEmpty() bool {
	return v.parameters.IsEmpty()
}

// This method returns the number of parameters contained in this context.
func (v *context) GetSize() int {
	return v.parameters.GetSize()
}

// This method returns all the parameters in this context. The parameters
// retrieved are in the same order as they are in the context.
func (v *context) AsArray() []abs.ParameterLike {
	// The implementation should be the following:
	//     return v.parameters.AsArray()
	// Alas, the Go compiler does not correctly recognize that the two result
	// types are identical so we have to do this explicitly:
	var array = make([]abs.ParameterLike, v.parameters.GetSize())
	var raw = v.parameters.AsArray()
	for index, parameter := range raw {
		array[index] = parameter
	}
	return array
}

// ASSOCIATIVE INTERFACE

// This method returns the keys for this context.
func (v *context) GetKeys() abs.Sequential[abs.SymbolLike] {
	return v.parameters.GetKeys()
}

// This method returns the values associated with the specified keys for this
// context. The values are returned in the same order as the keys in the
// context.
func (v *context) GetValues(keys abs.Sequential[abs.SymbolLike]) abs.Sequential[abs.ComponentLike] {
	return v.parameters.GetValues(keys)
}

// This method returns the value that is associated with the specified key in
// this context.
func (v *context) GetValue(key abs.SymbolLike) abs.ComponentLike {
	return v.parameters.GetValue(key)
}

// This method sets the value associated with the specified key to the
// specified value.
func (v *context) SetValue(key abs.SymbolLike, value abs.ComponentLike) {
	v.parameters.SetValue(key, value)
}

// This method removes the parameter associated with the specified key from the
// context and returns it.
func (v *context) RemoveValue(key abs.SymbolLike) abs.ComponentLike {
	return v.parameters.RemoveValue(key)
}

// This method removes the parameters associated with the specified keys from
// the context and returns the removed values.
func (v *context) RemoveValues(keys abs.Sequential[abs.SymbolLike]) abs.Sequential[abs.ComponentLike] {
	return v.parameters.RemoveValues(keys)
}

// This method removes all parameters from this context.
func (v *context) RemoveAll() {
	v.parameters.RemoveAll()
}
