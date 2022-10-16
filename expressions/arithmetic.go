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

// ARITHMETIC EXPRESSION IMPLEMENTATION

// This constructor creates a new arithmetic expression.
func Arithmetic(first any, operator abstractions.Operator, second any) abstractions.ArithmeticLike {
	var v = &arithmeticExpression{}
	// Perform argument validation.
	v.SetFirst(first)
	v.SetOperator(operator)
	v.SetSecond(second)
	return v
}

// This type defines the structure and methods associated with an arithmetic
// expression.
type arithmeticExpression struct {
	first    any
	operator abstractions.Operator
	second   any
}

// This method returns the first expression in this arithmetic expression.
func (v *arithmeticExpression) GetFirst() any {
	return v.first
}

// This method sets the first expression in this arithmetic expression to the
// specified value.
func (v *arithmeticExpression) SetFirst(first any) {
	if first == nil {
		panic("The first expression in an arithmetic expression cannot be nil.")
	}
	v.first = first
}

// This method returns the arithmetic operator in this arithmetic expression.
func (v *arithmeticExpression) GetOperator() abstractions.Operator {
	return v.operator
}

// This method sets the arithmetic operator in this arithmetic expression to the
// specified value.
func (v *arithmeticExpression) SetOperator(operator abstractions.Operator) {
	if operator < abstractions.PRODUCT || operator > abstractions.DIFFERENCE {
		panic("The operator in an arithmetic expression must be a valid arithmetic operation.")
	}
	v.operator = operator
}

// This method returns the second expression in this arithmetic expression.
func (v *arithmeticExpression) GetSecond() any {
	return v.second
}

// This method sets the second expression in this arithmetic expression to the
// specified value.
func (v *arithmeticExpression) SetSecond(second any) {
	if second == nil {
		panic("The second expression in an arithmetic expression cannot be nil.")
	}
	v.second = second
}
