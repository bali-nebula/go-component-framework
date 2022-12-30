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

// COMPARISON EXPRESSION IMPLEMENTATION

// This constructor creates a new comparison expression.
func Comparison(first abs.Expression, operator abs.Operator, second abs.Expression) abs.ComparisonLike {
	var v = &comparisonExpression{}
	// Perform argument validation.
	v.SetFirst(first)
	v.SetOperator(operator)
	v.SetSecond(second)
	return v
}

// This type defines the structure and methods associated with a comparison
// expression.
type comparisonExpression struct {
	first    abs.Expression
	operator abs.Operator
	second   abs.Expression
}

// This method is a dummy method that always returns true.
func (v *comparisonExpression) IsComparison() bool {
	return true
}

// This method returns the first expression in this comparison expression.
func (v *comparisonExpression) GetFirst() abs.Expression {
	return v.first
}

// This method sets the first expression in this comparison expression to the
// specified value.
func (v *comparisonExpression) SetFirst(first abs.Expression) {
	if first == nil {
		panic("The first expression in a comparison expression cannot be nil.")
	}
	v.first = first
}

// This method returns the comparison operator in this comparison expression.
func (v *comparisonExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the comparison operator in this comparison expression to the
// specified value.
func (v *comparisonExpression) SetOperator(operator abs.Operator) {
	if operator < abs.LESS || operator > abs.MATCHES {
		panic("The operator in a comparison expression must be a valid comparison operator.")
	}
	v.operator = operator
}

// This method returns the second expression in this comparison expression.
func (v *comparisonExpression) GetSecond() abs.Expression {
	return v.second
}

// This method sets the second expression in this comparison expression to the
// specified value.
func (v *comparisonExpression) SetSecond(second abs.Expression) {
	if second == nil {
		panic("The second expression in a comparison expression cannot be nil.")
	}
	v.second = second
}
