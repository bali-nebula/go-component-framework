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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// ITEM EXPRESSION IMPLEMENTATION

// This constructor creates a new item expression.
func Item(composite abs.Expression, indices abs.Indices) abs.ItemLike {
	var v = &itemExpression{}
	// Perform argument validation.
	v.SetComposite(composite)
	v.SetIndices(indices)
	return v
}

// This type defines the structure and methods associated with an item
// expression.
type itemExpression struct {
	composite abs.Expression
	indices   abs.Indices
}

// This method is a dummy method that always returns true.
func (v *itemExpression) IsItem() bool {
	return true
}

// This method returns the composite for this item expression.
func (v *itemExpression) GetComposite() abs.Expression {
	return v.composite
}

// This method sets the composite for this item expression.
func (v *itemExpression) SetComposite(composite abs.Expression) {
	if composite == nil {
		panic("An item expression requires a composite expression.")
	}
	v.composite = composite
}

// This method returns the list of indices for this item expression.
func (v *itemExpression) GetIndices() abs.Indices {
	return v.indices
}

// This method sets the list of indices for this item expression.
func (v *itemExpression) SetIndices(indices abs.Indices) {
	if indices == nil || indices.IsEmpty() {
		panic("An item expression requires at least one index.")
	}
	v.indices = indices
}
