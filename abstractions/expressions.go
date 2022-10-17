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

// OPERATOR CONSTANTS

type Operator int

const (
	INVALID Operator = iota
	AT
	CHAIN
	POWER
	INVERSE
	RECIPROCAL
	CONJUGATE
	PRODUCT
	QUOTIENT
	REMAINDER
	SUM
	DIFFERENCE
	LESS
	EQUAL
	MORE
	UNEQUAL
	IS
	MATCHES
	NOT
	AND
	SANS
	XOR
	OR
)

// EXPRESSION INTERFACES

// This interface defines the methods supported by all arithmetic-like types.
type ArithmeticLike interface {
	GetFirst() any
	SetFirst(first any)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() any
	SetSecond(second any)
}

// This interface defines the methods supported by all chaining-like types.
type ChainingLike interface {
	GetFirst() any
	SetFirst(first any)
	GetSecond() any
	SetSecond(second any)
}

// This interface defines the methods supported by all comparison-like types.
type ComparisonLike interface {
	GetFirst() any
	SetFirst(first any)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() any
	SetSecond(second any)
}

// This interface defines the methods supported by all complement-like types.
type ComplementLike interface {
	GetExpression() any
	SetExpression(expression any)
}

// This interface defines the methods supported by all dereference-like types.
type DereferenceLike interface {
	GetExpression() any
	SetExpression(expression any)
}

// This interface defines the methods supported by all exponential-like types.
type ExponentialLike interface {
	GetBase() any
	SetBase(base any)
	GetExponent() any
	SetExponent(exponent any)
}

// This interface defines the methods supported by all intrinsic-like types.
type IntrinsicLike interface {
	GetFunction() string
	SetFunction(function string)
	GetArgument(index int) any           // Ordinal based indexing.
	SetArgument(index int, argument any) // Ordinal based indexing.
	GetArguments() ListLike[any]
	SetArguments(arguments ListLike[any])
}

// This interface defines the methods supported by all inversion-like types.
type InversionLike interface {
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExpression() any
	SetExpression(expression any)
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
	GetArguments() ListLike[any]
	SetArguments(arguments ListLike[any])
}

// This interface defines the methods supported by all logical-like types.
type LogicalLike interface {
	GetFirst() any
	SetFirst(first any)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() any
	SetSecond(second any)
}

// This interface defines the methods supported by all magnitude-like types.
type MagnitudeLike interface {
	GetExpression() any
	SetExpression(expression any)
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
	GetIndices() ListLike[any]
	SetIndices(indices ListLike[any])
}
