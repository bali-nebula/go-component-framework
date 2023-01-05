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

// INVOCATION EXPRESSION IMPLEMENTATION

// This constructor creates a new invocation expression.
func Invocation(
	target abs.Expression,
	operator abs.Operator,
	message string,
	arguments abs.Arguments,
) abs.InvocationLike {
	var v = &invocationExpression{}
	// Perform argument validation.
	v.SetTarget(target)
	v.SetOperator(operator)
	v.SetMessage(message)
	v.SetArguments(arguments)
	return v
}

// This type defines the structure and methods associated with an invocation
// expression.
type invocationExpression struct {
	target    abs.Expression
	operator  abs.Operator
	message   string
	arguments abs.Arguments
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
		panic("An invocation expression requires a target expression for the message.")
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

// This method returns the message name for this invocation expression.
func (v *invocationExpression) GetMessage() string {
	return v.message
}

// This method sets the message name for this invocation expression.
func (v *invocationExpression) SetMessage(message string) {
	if len(message) == 0 {
		panic("An invocation expression requires a message name.")
	}
	v.message = message
}

// This method returns the list of arguments for this invocation expression.
func (v *invocationExpression) GetArguments() abs.Arguments {
	return v.arguments
}

// This method sets the list of arguments for this invocation expression.
func (v *invocationExpression) SetArguments(arguments abs.Arguments) {
	if arguments == nil {
		panic("An invocation expression requires an array (possibly empty) of arguments.")
	}
	v.arguments = arguments
}
