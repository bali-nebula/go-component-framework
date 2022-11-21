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
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
)

// DEREFERENCE EXPRESSION IMPLEMENTATION

// This constructor creates a new dereference expression.
func Dereference(operator abs.Operator, expression abs.Expression) abs.DereferenceLike {
	var v = &dereferenceExpression{}
	// Perform argument validation.
	v.SetOperator(operator)
	v.SetExpression(expression)
	return v
}

// This type defines the structure and methods associated with a dereference
// expression.
type dereferenceExpression struct {
	operator   abs.Operator
	expression abs.Expression
}

// This method is a dummy method that always returns true.
func (v *dereferenceExpression) IsDereference() bool {
	return true
}

// This method returns the dereference operator in this dereference expression.
func (v *dereferenceExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the dereference operator in this dereference expression to the
// specified value.
func (v *dereferenceExpression) SetOperator(operator abs.Operator) {
	if operator != abs.AT {
		panic("The operator in an dereference expression must be a valid dereference operator.")
	}
	v.operator = operator
}

// This method returns the expression to be operated on by this dereference
// expression.
func (v *dereferenceExpression) GetExpression() abs.Expression {
	return v.expression
}

// This method sets the expression to be operated on by this dereference
// expression to the specified value.
func (v *dereferenceExpression) SetExpression(expression abs.Expression) {
	if expression == nil {
		panic("The expression to be operated on cannot be nil.")
	}
	v.expression = expression
}
