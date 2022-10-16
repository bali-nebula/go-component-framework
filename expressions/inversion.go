/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package components

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// INVERSION EXPRESSION IMPLEMENTATION

// This constructor creates a new inversion expression.
func Inversion(expression any) abstractions.InversionLike {
	var v = &inversionExpression{}
	// Perform argument validation.
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with a inversion
// expression.
type inversionExpression struct {
	expression any
}

// This method returns the expression to be inversion by this inversion
// expression.
func (v *inversionExpression) GetExpression() any {
	return v.expression
}

// This method sets the expression to be inverted by this inversion
// expression to the specified value.
func (v *inversionExpression) SetExpression(expression any) {
	if expression == nil {
		panic("The expression to be inverted cannot be nil.")
	}
	v.expression = expression
}
