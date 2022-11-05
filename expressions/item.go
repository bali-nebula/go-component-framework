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
func Item(composite abs.ExpressionLike, indices abs.ListLike[abs.ExpressionLike]) abs.ItemLike {
	var v = &itemExpression{}
	// Perform argument validation.
	v.SetComposite(composite)
	v.SetIndices(indices)
	return v
}

// This type defines the structure and methods associated with an item
// expression.
type itemExpression struct {
	composite abs.ExpressionLike
	indices   abs.ListLike[abs.ExpressionLike]
}

// This method is a dummy method that always returns true.
func (v *itemExpression) IsItem() bool {
	return true
}

// This method returns the composite for this item expression.
func (v *itemExpression) GetComposite() abs.ExpressionLike {
	return v.composite
}

// This method sets the composite for this item expression.
func (v *itemExpression) SetComposite(composite abs.ExpressionLike) {
	if composite == nil {
		panic("An item expression requires a composite expression.")
	}
	v.composite = composite
}

// This method returns the index at the specified index from this item
// expression.
func (v *itemExpression) GetIndex(index int) abs.ExpressionLike {
	return v.indices.GetValue(index)
}

// This method sets the expression at the specified index for this item
// expression.
func (v *itemExpression) SetIndex(index int, expression abs.ExpressionLike) {
	if expression == nil {
		panic("Each index for an item expression requires an expression.")
	}
	v.indices.SetValue(index, expression)
}

// This method returns the list of indices for this item expression.
func (v *itemExpression) GetIndices() abs.ListLike[abs.ExpressionLike] {
	return v.indices
}

// This method sets the list of indices for this item expression.
func (v *itemExpression) SetIndices(indices abs.ListLike[abs.ExpressionLike]) {
	if indices == nil || indices.IsEmpty() {
		panic("An item expression requires at least one index.")
	}
	v.indices = indices
}
