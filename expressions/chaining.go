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

// CHAINING EXPRESSION IMPLEMENTATION

// This constructor creates a new chaining expression.
func Chaining(first abs.Expression, operator abs.Operator, second abs.Expression) abs.ChainingLike {
	var v = &chainingExpression{}
	// Perform argument validation.
	v.SetFirst(first)
	v.SetOperator(operator)
	v.SetSecond(second)
	return v
}

// This type defines the structure and methods associated with a chaining
// expression.
type chainingExpression struct {
	first    abs.Expression
	operator abs.Operator
	second   abs.Expression
}

// This method returns the first expression in this chaining expression.
func (v *chainingExpression) GetFirst() abs.Expression {
	return v.first
}

// This method sets the first expression in this chaining expression to the
// specified value.
func (v *chainingExpression) SetFirst(first abs.Expression) {
	if first == nil {
		panic("The first expression in a chaining expression cannot be nil.")
	}
	v.first = first
}

// This method returns the comparison operator in this chaining expression.
func (v *chainingExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the comparison operator in this chaining expression to the
// specified value.
func (v *chainingExpression) SetOperator(operator abs.Operator) {
	if operator != abs.AMPERSAND {
		panic("The operator in a comparison expression must be a valid chaining operator.")
	}
	v.operator = operator
}

// This method returns the second expression in this chaining expression.
func (v *chainingExpression) GetSecond() abs.Expression {
	return v.second
}

// This method sets the second expression in this chaining expression to the
// specified value.
func (v *chainingExpression) SetSecond(second abs.Expression) {
	if second == nil {
		panic("The second expression in a chaining expression cannot be nil.")
	}
	v.second = second
}
