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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// LOGICAL EXPRESSION IMPLEMENTATION

// This constructor creates a new logical expression.
func Logical(first abs.Expression, operator abs.Operator, second abs.Expression) abs.LogicalLike {
	var v = &logicalExpression{}
	// Perform argument validation.
	v.SetFirst(first)
	v.SetOperator(operator)
	v.SetSecond(second)
	return v
}

// This type defines the structure and methods associated with a logical
// expression.
type logicalExpression struct {
	first    abs.Expression
	operator abs.Operator
	second   abs.Expression
}

// This method is a dummy method that always returns true.
func (v *logicalExpression) IsLogical() bool {
	return true
}

// This method returns the first expression in this logical expression.
func (v *logicalExpression) GetFirst() abs.Expression {
	return v.first
}

// This method sets the first expression in this logical expression to the
// specified value.
func (v *logicalExpression) SetFirst(first abs.Expression) {
	if first == nil {
		panic("The first expression in a logical expression cannot be nil.")
	}
	v.first = first
}

// This method returns the logical operator in this logical expression.
func (v *logicalExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the logical operator in this logical expression to the
// specified value.
func (v *logicalExpression) SetOperator(operator abs.Operator) {
	if operator < abs.AND || operator > abs.XOR {
		panic(fmt.Sprintf("The operator in a logical expression must be a valid logical operator: %v", operator))
	}
	v.operator = operator
}

// This method returns the second expression in this logical expression.
func (v *logicalExpression) GetSecond() abs.Expression {
	return v.second
}

// This method sets the second expression in this logical expression to the
// specified value.
func (v *logicalExpression) SetSecond(second abs.Expression) {
	if second == nil {
		panic("The second expression in a logical expression cannot be nil.")
	}
	v.second = second
}
