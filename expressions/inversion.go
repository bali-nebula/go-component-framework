/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
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

// INVERSION EXPRESSION IMPLEMENTATION

// This constructor creates a new inversion expression.
func Inversion(operator abs.Operator, expression abs.Expression) abs.InversionLike {
	var v = &inversionExpression{}
	// Perform argument validation.
	v.SetOperator(operator)
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with a inversion
// expression.
type inversionExpression struct {
	operator   abs.Operator
	expression abs.Expression
}

// This method returns the inversion operator in this inversion expression.
func (v *inversionExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the inversion operator in this inversion expression to the
// specified value.
func (v *inversionExpression) SetOperator(operator abs.Operator) {
	if operator < abs.MINUS || operator > abs.SLASH {
		panic("The operator in an inversion expression must be a valid inversion operator.")
	}
	v.operator = operator
}

// This method returns the expression to be operated on by this inversion
// expression.
func (v *inversionExpression) GetExpression() abs.Expression {
	return v.expression
}

// This method sets the expression to be operated on by this inversion
// expression to the specified value.
func (v *inversionExpression) SetExpression(expression abs.Expression) {
	if expression == nil {
		panic("The expression to be operated on cannot be nil.")
	}
	v.expression = expression
}
