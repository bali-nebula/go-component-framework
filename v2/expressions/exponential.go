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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
)

// EXPONENTIAL EXPRESSION IMPLEMENTATION

// This constructor creates a new exponential expression.
func Exponential(base abs.Expression, operator abs.Operator, exponent abs.Expression) abs.BinaryOperationLike {
	var v = &exponentialExpression{}
	// Perform argument validation.
	v.SetFirst(base)
	v.SetOperator(operator)
	v.SetSecond(exponent)
	return v
}

// This type defines the structure and methods associated with a exponential
// expression.
type exponentialExpression struct {
	base     abs.Expression
	operator abs.Operator
	exponent abs.Expression
}

// This method returns the first expression in this exponential expression.
func (v *exponentialExpression) GetFirst() abs.Expression {
	return v.base
}

// This method sets the first expression in this exponential expression to the
// specified value.
func (v *exponentialExpression) SetFirst(base abs.Expression) {
	if base == nil {
		panic("The base expression in a exponential expression cannot be nil.")
	}
	v.base = base
}

// This method returns the exponential operator in this exponential expression.
func (v *exponentialExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the exponential operator in this exponential expression to the
// specified value.
func (v *exponentialExpression) SetOperator(operator abs.Operator) {
	if operator != abs.CARET {
		panic("The operator in a exponential expression must be a valid exponential operator.")
	}
	v.operator = operator
}

// This method returns the second expression in this exponential expression.
func (v *exponentialExpression) GetSecond() abs.Expression {
	return v.exponent
}

// This method sets the second expression in this exponential expression to
// the specified value.
func (v *exponentialExpression) SetSecond(exponent abs.Expression) {
	if exponent == nil {
		panic("The exponent expression in a exponential expression cannot be nil.")
	}
	v.exponent = exponent
}
