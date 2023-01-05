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

// COMPLEMENT EXPRESSION IMPLEMENTATION

// This constructor creates a new complement expression.
func Complement(operator abs.Operator, expression abs.Expression) abs.ComplementLike {
	var v = &complementExpression{}
	// Perform argument validation.
	v.SetOperator(operator)
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with a complement
// expression.
type complementExpression struct {
	operator   abs.Operator
	expression abs.Expression
}

// This method returns the complement operator in this complement expression.
func (v *complementExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the complement operator in this complement expression to the
// specified value.
func (v *complementExpression) SetOperator(operator abs.Operator) {
	if operator != abs.NOT {
		panic("The operator in a complement expression must be a valid complement operator.")
	}
	v.operator = operator
}

// This method returns the expression to be operated on by this complement
// expression.
func (v *complementExpression) GetExpression() abs.Expression {
	return v.expression
}

// This method sets the expression to be operated on by this complement
// expression to the specified value.
func (v *complementExpression) SetExpression(expression abs.Expression) {
	if expression == nil {
		panic("The expression to be operated on cannot be nil.")
	}
	v.expression = expression
}
