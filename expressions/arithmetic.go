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
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// ARITHMETIC EXPRESSION IMPLEMENTATION

// This constructor creates a new arithmetic expression.
func Arithmetic(first abs.Expression, operator abs.Operator, second abs.Expression) abs.ArithmeticLike {
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
	first    abs.Expression
	operator abs.Operator
	second   abs.Expression
}

// This method is a dummy method that always returns true.
func (v *arithmeticExpression) IsArithmetic() bool {
	return true
}

// This method returns the first expression in this arithmetic expression.
func (v *arithmeticExpression) GetFirst() abs.Expression {
	return v.first
}

// This method sets the first expression in this arithmetic expression to the
// specified value.
func (v *arithmeticExpression) SetFirst(first abs.Expression) {
	if first == nil {
		panic("The first expression in an arithmetic expression cannot be nil.")
	}
	v.first = first
}

// This method returns the arithmetic operator in this arithmetic expression.
func (v *arithmeticExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the arithmetic operator in this arithmetic expression to the
// specified value.
func (v *arithmeticExpression) SetOperator(operator abs.Operator) {
	if operator < abs.PLUS || operator > abs.MODULO {
		panic(fmt.Sprintf("The operator in an arithmetic expression must be valid: %v", operator))
	}
	v.operator = operator
}

// This method returns the second expression in this arithmetic expression.
func (v *arithmeticExpression) GetSecond() abs.Expression {
	return v.second
}

// This method sets the second expression in this arithmetic expression to the
// specified value.
func (v *arithmeticExpression) SetSecond(second abs.Expression) {
	if second == nil {
		panic("The second expression in an arithmetic expression cannot be nil.")
	}
	v.second = second
}
