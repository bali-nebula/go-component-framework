/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologiesâ„˘.  All Rights Reserved.   *
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

// STACK IMPLEMENTATION

// This constructor creates a new empty component stack with the default
// capacity.
func Stack() abs.StackLike {
	var v = col.Stack[abs.ComponentLike]()
	return &stack{v}
}

// This constructor creates a new empty component stack with the specified
// capacity.
func StackWithCapacity(capacity int) abs.StackLike {
	var v = col.StackWithCapacity[abs.ComponentLike](capacity)
	return &stack{v}
}

// This type defines the structure and methods associated with a component
// stack.
type stack struct {
	components col.StackLike[abs.ComponentLike]
}

// SEQUENTIAL INTERFACE

// This method determines whether or not this stack is empty.
func (v *stack) IsEmpty() bool {
	return v.components.IsEmpty()
}

// This method returns the number of components contained in this stack.
func (v *stack) GetSize() int {
	return v.components.GetSize()
}

// This method returns all the components in this stack. The components retrieved are in
// the same order as they are in the stack.
func (v *stack) AsArray() []abs.ComponentLike {
	return v.components.AsArray()
}

// EXPANDABLE INTERFACE

// This method appends the specified component to the end of this stack.
func (v *stack) AddValue(component abs.ComponentLike) {
	v.components.AddValue(component)
}

// This method appends the specified components to the end of this stack.
func (v *stack) AddValues(components abs.Sequential[abs.ComponentLike]) {
	v.components.AddValues(components)
}

// LIFO INTERFACE

// This method retrieves the capacity of this stack.
func (v *stack) GetCapacity() int {
	return v.components.GetCapacity()
}

// This method retrieves from this stack the value that is on top of it.
func (v *stack) GetTop() abs.ComponentLike {
	return v.components.GetTop()
}

// This method removes from this stack the value that is on top of it.
func (v *stack) RemoveTop() abs.ComponentLike {
	return v.components.RemoveTop()
}

// This method removes all values from this stack.
func (v *stack) RemoveAll() {
	v.components.RemoveAll()
}
