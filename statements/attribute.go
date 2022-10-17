/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package statements

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// ATTRIBUTE IMPLEMENTATION

// This constructor creates a new attribute.
func Attribute(variable string, indices abstractions.ListLike[any]) abstractions.AttributeLike {
	var v = &attribute{}
	// Perform argument validation.
	v.SetVariable(variable)
	v.SetIndices(indices)
	return v
}

// This type defines the structure and methods associated with an attribute.
type attribute struct {
	variable string
	indices  abstractions.ListLike[any]
}

// This method returns the variable for this attribute.
func (v *attribute) GetVariable() string {
	return v.variable
}

// This method sets the variable for this attribute.
func (v *attribute) SetVariable(variable string) {
	if len(variable) == 0 {
		panic("An attribute requires a variable.")
	}
	v.variable = variable
}

// This method returns the index at the specified index from this attribute.
func (v *attribute) GetIndex(index int) any {
	return v.indices.GetItem(index)
}

// This method sets the expression at the specified index for this attribute.
func (v *attribute) SetIndex(index int, expression any) {
	if expression == nil {
		panic("Each index expression for an attribute requires a value.")
	}
	v.indices.SetItem(index, expression)
}

// This method returns the list of indices for this attribute.
func (v *attribute) GetIndices() abstractions.ListLike[any] {
	return v.indices
}

// This method sets the list of indices for this attribute.
func (v *attribute) SetIndices(indices abstractions.ListLike[any]) {
	if indices == nil || indices.IsEmpty() {
		panic("An attribute requires at least one index.")
	}
	v.indices = indices
}
