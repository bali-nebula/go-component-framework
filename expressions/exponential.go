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

// EXPONENTIAL EXPRESSION IMPLEMENTATION

// This constructor creates a new exponential expression.
func Exponential(base abs.ExpressionLike, operator abs.Operator, exponent abs.ExpressionLike) abs.ExponentialLike {
	var v = &exponentialExpression{}
	// Perform argument validation.
	v.SetBase(base)
	v.SetOperator(operator)
	v.SetExponent(exponent)
	return v
}

// This type defines the structure and methods associated with a exponential
// expression.
type exponentialExpression struct {
	base     abs.ExpressionLike
	operator abs.Operator
	exponent abs.ExpressionLike
}

// This method is a dummy method that always returns true.
func (v *exponentialExpression) IsExponential() bool {
	return true
}

// This method returns the base expression in this exponential expression.
func (v *exponentialExpression) GetBase() abs.ExpressionLike {
	return v.base
}

// This method sets the base expression in this exponential expression to the
// specified value.
func (v *exponentialExpression) SetBase(base abs.ExpressionLike) {
	if base == nil {
		panic("The base expression in a exponential expression cannot be nil.")
	}
	v.base = base
}

// This method returns the comparison operator in this exponential expression.
func (v *exponentialExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the comparison operator in this exponential expression to the
// specified value.
func (v *exponentialExpression) SetOperator(operator abs.Operator) {
	if operator != abs.CARET {
		panic("The operator in an comparison expression must be a valid exponential operator.")
	}
	v.operator = operator
}

// This method returns the exponent expression in this exponential expression.
func (v *exponentialExpression) GetExponent() abs.ExpressionLike {
	return v.exponent
}

// This method sets the exponent expression in this exponential expression to
// the specified value.
func (v *exponentialExpression) SetExponent(exponent abs.ExpressionLike) {
	if exponent == nil {
		panic("The exponent expression in a exponential expression cannot be nil.")
	}
	v.exponent = exponent
}
