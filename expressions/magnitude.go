/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologiesâ„˘.  All Rights Reserved.   *
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

// MAGNITUDE EXPRESSION IMPLEMENTATION

// This constructor creates a new magnitude expression.
func Magnitude(expression abs.Expression) abs.MagnitudeLike {
	var v = &magnitudeExpression{}
	// Perform argument validation.
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with a magnitude
// expression.
type magnitudeExpression struct {
	expression abs.Expression
}

// This method returns the expression to be operated on by this magnitude
// expression.
func (v *magnitudeExpression) GetExpression() abs.Expression {
	return v.expression
}

// This method sets the expression to be operated on by this magnitude
// expression to the specified value.
func (v *magnitudeExpression) SetExpression(expression abs.Expression) {
	if expression == nil {
		panic("The expression to be operated on cannot be nil.")
	}
	v.expression = expression
}
