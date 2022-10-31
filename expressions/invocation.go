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

// INVOCATION EXPRESSION IMPLEMENTATION

// This constructor creates a new asynchronous invocation expression.
func AsynchronousInvocation(
	target abs.ExpressionLike,
	message string,
	arguments abs.ListLike[abs.ExpressionLike],
) abs.InvocationLike {
	var v = &invocationExpression{}
	// Perform argument validation.
	v.SetSynchronous(false)
	v.SetTarget(target)
	v.SetMessage(message)
	v.SetArguments(arguments)
	return v
}

// This constructor creates a new invocation expression.
func Invocation(
	target abs.ExpressionLike,
	message string,
	arguments abs.ListLike[abs.ExpressionLike],
) abs.InvocationLike {
	var v = &invocationExpression{}
	// Perform argument validation.
	v.SetSynchronous(true)
	v.SetTarget(target)
	v.SetMessage(message)
	v.SetArguments(arguments)
	return v
}

// This type defines the structure and methods associated with an invocation
// expression.
type invocationExpression struct {
	isSynchronous bool
	target        abs.ExpressionLike
	message       string
	arguments     abs.ListLike[abs.ExpressionLike]
}

// This method determines whether or not this invocation expression is
// synchronous.
func (v *invocationExpression) IsSynchronous() bool {
	return v.isSynchronous
}

// This method sets whether or not this invocation expression is synchronous.
func (v *invocationExpression) SetSynchronous(isSynchronous bool) {
	v.isSynchronous = isSynchronous
}

// This method returns the target expression for this invocation expression.
func (v *invocationExpression) GetTarget() abs.ExpressionLike {
	return v.target
}

// This method sets the target expression for this invocation expression.
func (v *invocationExpression) SetTarget(target abs.ExpressionLike) {
	if target == nil {
		panic("An invocation expression requires a target expression for the message.")
	}
	v.target = target
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

// This method returns the argument at the specified index from this invocation
// expression.
func (v *invocationExpression) GetArgument(index int) abs.ExpressionLike {
	return v.arguments.GetItem(index)
}

// This method sets the argument at the specified index for this invocation
// expression.
func (v *invocationExpression) SetArgument(index int, argument abs.ExpressionLike) {
	if argument == nil {
		panic("Each argument for an invocation expression requires a value.")
	}
	v.arguments.SetItem(index, argument)
}

// This method returns the list of arguments for this invocation expression.
func (v *invocationExpression) GetArguments() abs.ListLike[abs.ExpressionLike] {
	return v.arguments
}

// This method sets the list of arguments for this invocation expression.
func (v *invocationExpression) SetArguments(arguments abs.ListLike[abs.ExpressionLike]) {
	if arguments == nil {
		panic("An invocation expression requires an array (possibly empty) of arguments.")
	}
	v.arguments = arguments
}
