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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
)

// STACK IMPLEMENTATION

// This constructor creates a new empty stack with the default capacity.
// The default capacity is 16 values.
func Stack[T abs.ValueLike]() abs.StackLike[T] {
	return StackWithCapacity[T](0)
}

// This constructor creates a new empty stack with the specified capacity.
func StackWithCapacity[T abs.ValueLike](capacity int) abs.StackLike[T] {
	// Groom the arguments.
	if capacity < 1 {
		capacity = 16 // The default value.
	}

	// Return an empty stack.
	var values = List[T]()
	return &stack[T]{values, values, capacity}
}

// This type defines the structure and methods associated with a stack of
// values. A stack implements last-in-first-out semantics.
// This type is parameterized as follows:
//   - T is any type of value.
type stack[T abs.ValueLike] struct {
	// Note: The delegated methods don't see the real collection type.
	abs.Sequential[T]
	values   abs.ListLike[T]
	capacity int
}

// LIFO INTERFACE

// This method retrieves the capacity of this stack.
func (v *stack[T]) GetCapacity() int {
	return v.capacity
}

// This method adds the specified value to the top of this stack.
func (v *stack[T]) AddValue(value T) {
	if v.values.GetSize() == v.capacity {
		panic("Attempted to add an value onto a stack that has reached its capacity!")
	}
	v.values.AddValue(value)
}

// This method adds the specified values to the top of this stack.
func (v *stack[T]) AddValues(values abs.Sequential[T]) {
	var iterator = age.Iterator(values)
	for iterator.HasNext() {
		var value = iterator.GetNext()
		v.AddValue(value) // We must call this explicitly to get the capacity check.
	}
}

// This method retrieves from this stack the value that is on top of it.
func (v *stack[T]) GetTop() T {
	if v.values.IsEmpty() {
		panic("Attempted to retrieve the top of an empty stack!")
	}
	return v.values.GetValue(-1)
}

// This method removes from this stack the value that is on top of it.
func (v *stack[T]) RemoveTop() T {
	if v.values.IsEmpty() {
		panic("Attempted to remove the top of an empty stack!")
	}
	return v.values.RemoveValue(-1)
}

// This method removes all values from this stack.
func (v *stack[T]) RemoveAll() {
	v.values.RemoveAll()
}
