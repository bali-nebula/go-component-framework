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

// INVOCATION EXPRESSION IMPLEMENTATION

// This constructor creates a new invocation expression.
func Invocation(
	target abs.Expression,
	operator abs.Operator,
	method string,
	arguments abs.Sequential[abs.Expression],
) abs.InvocationLike {
	var v = &invocationExpression{}
	// Perform argument validation.
	v.SetTarget(target)
	v.SetOperator(operator)
	v.SetMethod(method)
	v.SetArguments(arguments)
	return v
}

// This type defines the structure and methods associated with an invocation
// expression.
type invocationExpression struct {
	target    abs.Expression
	operator  abs.Operator
	method    string
	arguments abs.Sequential[abs.Expression]
}

// This method determines whether or not this invocation expression is
// synchronous.
func (v *invocationExpression) IsSynchronous() bool {
	return v.operator == abs.ARROW
}

// This method returns the target expression for this invocation expression.
func (v *invocationExpression) GetTarget() abs.Expression {
	return v.target
}

// This method sets the target expression for this invocation expression.
func (v *invocationExpression) SetTarget(target abs.Expression) {
	if target == nil {
		panic("An invocation expression requires a target expression for the method.")
	}
	v.target = target
}

// This method returns the operator for this invocation expression.
func (v *invocationExpression) GetOperator() abs.Operator {
	return v.operator
}

// This method sets the target expression for this invocation expression.
func (v *invocationExpression) SetOperator(operator abs.Operator) {
	if operator < abs.DOT || operator > abs.ARROW {
		panic("An invocation expression requires a valid operator.")
	}
	v.operator = operator
}

// This method returns the method name for this invocation expression.
func (v *invocationExpression) GetMethod() string {
	return v.method
}

// This method sets the method name for this invocation expression.
func (v *invocationExpression) SetMethod(method string) {
	if len(method) == 0 {
		panic("An invocation expression requires a method name.")
	}
	v.method = method
}

// This method returns the list of arguments for this invocation expression.
func (v *invocationExpression) GetArguments() abs.Sequential[abs.Expression] {
	return v.arguments
}

// This method sets the list of arguments for this invocation expression.
func (v *invocationExpression) SetArguments(arguments abs.Sequential[abs.Expression]) {
	if arguments == nil {
		panic("An invocation expression requires an array (possibly empty) of arguments.")
	}
	v.arguments = arguments
}
