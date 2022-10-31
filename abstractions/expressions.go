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
	_ Operator = iota
	AND
	AT
	CHAIN
	CONJUGATE
	DIFFERENCE
	EQUAL
	INVERSE
	IS
	LESS
	MATCHES
	MORE
	NOT
	OR
	POWER
	PRODUCT
	QUOTIENT
	RECIPROCAL
	REMAINDER
	SANS
	SUM
	UNEQUAL
	XOR
)

// EXPRESSION INTERFACES

type Expression any

// This interface defines the methods supported by all arithmetic-like expressions.
type ArithmeticLike interface {
	GetFirst() Expression
	SetFirst(first Expression)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() Expression
	SetSecond(second Expression)
}

// This interface defines the methods supported by all chaining-like expressions.
type ChainingLike interface {
	GetFirst() Expression
	SetFirst(first Expression)
	GetSecond() Expression
	SetSecond(second Expression)
}

// This interface defines the methods supported by all comparison-like expressions.
type ComparisonLike interface {
	GetFirst() Expression
	SetFirst(first Expression)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() Expression
	SetSecond(second Expression)
}

// This interface defines the methods supported by all complement-like expressions.
type ComplementLike interface {
	GetExpression() Expression
	SetExpression(expression Expression)
}

// This interface defines the methods supported by all dereference-like expressions.
type DereferenceLike interface {
	GetExpression() Expression
	SetExpression(expression Expression)
}

// This interface defines the methods supported by all exponential-like expressions.
type ExponentialLike interface {
	GetBase() Expression
	SetBase(base Expression)
	GetExponent() Expression
	SetExponent(exponent Expression)
}

// This interface defines the methods supported by all intrinsic-like expressions.
type IntrinsicLike interface {
	GetFunction() string
	SetFunction(function string)
	GetArgument(index int) Expression           // Ordinal based indexing.
	SetArgument(index int, argument Expression) // Ordinal based indexing.
	GetArguments() ListLike[Expression]
	SetArguments(arguments ListLike[Expression])
}

// This interface defines the methods supported by all inversion-like expressions.
type InversionLike interface {
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExpression() Expression
	SetExpression(expression Expression)
}

// This interface defines the methods supported by all invocation-like expressions.
type InvocationLike interface {
	IsSynchronous() bool
	SetSynchronous(isSynchronous bool)
	GetTarget() Expression
	SetTarget(target Expression)
	GetMessage() string
	SetMessage(message string)
	GetArgument(index int) Expression           // Ordinal based indexing.
	SetArgument(index int, argument Expression) // Ordinal based indexing.
	GetArguments() ListLike[Expression]
	SetArguments(arguments ListLike[Expression])
}

// This interface defines the methods supported by all logical-like expressions.
type LogicalLike interface {
	GetFirst() Expression
	SetFirst(first Expression)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() Expression
	SetSecond(second Expression)
}

// This interface defines the methods supported by all magnitude-like expressions.
type MagnitudeLike interface {
	GetExpression() Expression
	SetExpression(expression Expression)
}

// This interface defines the methods supported by all precedence-like expressions.
type PrecedenceLike interface {
	GetExpression() Expression
	SetExpression(expression Expression)
}

// This interface defines the methods supported by all value-like expressions.
type ValueLike interface {
	GetComposite() Expression
	SetComposite(composite Expression)
	GetIndex(index int) Expression
	SetIndex(index int, expression Expression)
	GetIndices() ListLike[Expression]
	SetIndices(indices ListLike[Expression])
}

// This interface defines the methods supported by all variable-like expressions.
type VariableLike interface {
	GetIdentifier() string
}
