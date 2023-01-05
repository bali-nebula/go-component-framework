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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// ITEM EXPRESSION IMPLEMENTATION

// This constructor creates a new subcomponent expression.
func Subcomponent(composite abs.Expression, indices abs.Indices) abs.SubcomponentLike {
	var v = &subcomponentExpression{}
	// Perform argument validation.
	v.SetComposite(composite)
	v.SetIndices(indices)
	return v
}

// This type defines the structure and methods associated with an subcomponent
// expression.
type subcomponentExpression struct {
	composite abs.Expression
	indices   abs.Indices
}

// This method returns the composite for this subcomponent expression.
func (v *subcomponentExpression) GetComposite() abs.Expression {
	return v.composite
}

// This method sets the composite for this subcomponent expression.
func (v *subcomponentExpression) SetComposite(composite abs.Expression) {
	if composite == nil {
		panic("An subcomponent expression requires a composite expression.")
	}
	v.composite = composite
}

// This method returns the list of indices for this subcomponent expression.
func (v *subcomponentExpression) GetIndices() abs.Indices {
	return v.indices
}

// This method sets the list of indices for this subcomponent expression.
func (v *subcomponentExpression) SetIndices(indices abs.Indices) {
	if indices == nil || indices.IsEmpty() {
		panic("An subcomponent expression requires at least one index.")
	}
	v.indices = indices
}
