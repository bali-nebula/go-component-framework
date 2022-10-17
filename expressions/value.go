/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package expressions

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// VALUE EXPRESSION IMPLEMENTATION

// This constructor creates a new value expression.
func Value(composite any, indices abstractions.ListLike[any]) abstractions.ValueLike {
	var v = &valueExpression{}
	// Perform argument validation.
	v.SetComposite(composite)
	v.SetIndices(indices)
	return v
}

// This type defines the structure and methods associated with a value
// expression.
type valueExpression struct {
	composite any
	indices   abstractions.ListLike[any]
}

// This method returns the composite for this value expression.
func (v *valueExpression) GetComposite() any {
	return v.composite
}

// This method sets the composite for this value expression.
func (v *valueExpression) SetComposite(composite any) {
	if composite == nil {
		panic("A value expression requires a composite expression.")
	}
	v.composite = composite
}

// This method returns the index at the specified index from this value
// expression.
func (v *valueExpression) GetIndex(index int) any {
	return v.indices.GetItem(index)
}

// This method sets the expression at the specified index for this value
// expression.
func (v *valueExpression) SetIndex(index int, expression any) {
	if expression == nil {
		panic("Each index expression for a value expression requires a value.")
	}
	v.indices.SetItem(index, expression)
}

// This method returns the list of indices for this value expression.
func (v *valueExpression) GetIndices() abstractions.ListLike[any] {
	return v.indices
}

// This method sets the list of indices for this value expression.
func (v *valueExpression) SetIndices(indices abstractions.ListLike[any]) {
	if indices == nil || indices.IsEmpty() {
		panic("A value expression requires at least one index.")
	}
	v.indices = indices
}
