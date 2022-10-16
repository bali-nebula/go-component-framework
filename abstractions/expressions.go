/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// EXPRESSION INTERFACES

// This interface defines the methods supported by all dereference-like types.
type DereferenceLike interface {
	GetExpression() any
	SetExpression(expression any)
}

// This interface defines the methods supported by all intrinsic-like types.
type IntrinsicLike interface {
	GetFunction() string
	SetFunction(function string)
	GetArgument(index int) any           // Ordinal based indexing.
	SetArgument(index int, argument any) // Ordinal based indexing.
	GetArguments() ArrayLike[any]
	SetArguments(arguments ArrayLike[any])
}

// This interface defines the methods supported by all invocation-like types.
type InvocationLike interface {
	IsSynchronous() bool
	SetSynchronous(isSynchronous bool)
	GetTarget() any
	SetTarget(target any)
	GetMessage() string
	SetMessage(message string)
	GetArgument(index int) any           // Ordinal based indexing.
	SetArgument(index int, argument any) // Ordinal based indexing.
	GetArguments() ArrayLike[any]
	SetArguments(arguments ArrayLike[any])
}

// This interface defines the methods supported by all precedence-like types.
type PrecedenceLike interface {
	GetExpression() any
	SetExpression(expression any)
}

// This interface defines the methods supported by all value-like types.
type ValueLike interface {
	GetComposite() any
	SetComposite(composite any)
	GetIndex(index int) any
	SetIndex(index int, expression any)
	GetIndices() ArrayLike[any]
	SetIndices(indices ArrayLike[any])
}
