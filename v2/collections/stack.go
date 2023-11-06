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

// STACK IMPLEMENTATION

// This constructor creates a new empty value stack with the default
// capacity.
func Stack() abs.StackLike {
	var v = col.Stack[abs.ComponentLike]()
	return &stack{v}
}

// This constructor creates a new empty value stack with the specified
// capacity.
func StackWithCapacity(capacity int) abs.StackLike {
	var v = col.StackWithCapacity[abs.ComponentLike](capacity)
	return &stack{v}
}

// This constructor creates a new value stack from the specified sequence.
func StackFromSequence(sequence abs.Sequential[abs.ComponentLike]) abs.StackLike {
	var v = col.StackWithCapacity[abs.ComponentLike](sequence.GetSize() * 2)
	var iterator = col.Iterator[abs.ComponentLike](sequence)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		v.AddValue(value)
	}
	return &stack{v}
}

// This type defines the structure and methods associated with a value
// stack.
type stack struct {
	values col.StackLike[abs.ComponentLike]
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this stack is empty.
func (v *stack) IsEmpty() bool {
	return v.values.IsEmpty()
}

// This method returns the number of values contained in this stack.
func (v *stack) GetSize() int {
	return v.values.GetSize()
}

// This method returns all the values in this stack. The values retrieved are in
// the same order as they are in the stack.
func (v *stack) AsArray() []abs.ComponentLike {
	return v.values.AsArray()
}

// LIFO INTERFACE

// This method retrieves the capacity of this stack.
func (v *stack) GetCapacity() int {
	return v.values.GetCapacity()
}

// This method appends the specified value to the end of this stack.
func (v *stack) AddValue(value abs.ComponentLike) {
	v.values.AddValue(value)
}

// This method retrieves from this stack the value that is on top of it.
func (v *stack) GetTop() abs.ComponentLike {
	return v.values.GetTop()
}

// This method removes from this stack the value that is on top of it.
func (v *stack) RemoveTop() abs.ComponentLike {
	return v.values.RemoveTop()
}

// This method removes all values from this stack.
func (v *stack) RemoveAll() {
	v.values.RemoveAll()
}
