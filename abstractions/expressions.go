/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// TYPE DEFINITIONS

type (
	Expression any
)

// INDIVIDUAL INTERFACES

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
	GetOperator() Operator
	SetOperator(operator Operator)
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
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExpression() Expression
	SetExpression(expression Expression)
}

// This interface defines the methods supported by all dereference-like expressions.
type DereferenceLike interface {
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExpression() Expression
	SetExpression(expression Expression)
}

// This interface defines the methods supported by all exponential-like expressions.
type ExponentialLike interface {
	GetBase() Expression
	SetBase(base Expression)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExponent() Expression
	SetExponent(exponent Expression)
}

// This interface defines the methods supported by all intrinsic-like expressions.
type IntrinsicLike interface {
	GetFunction() string
	SetFunction(function string)
	GetArguments() Sequential[Expression]
	SetArguments(arguments Sequential[Expression])
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
	GetTarget() Expression
	SetTarget(target Expression)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetMessage() string
	SetMessage(message string)
	GetArguments() Sequential[Expression]
	SetArguments(arguments Sequential[Expression])
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

// This interface defines the methods supported by all item-like expressions.
type SubcomponentLike interface {
	GetComposite() Expression
	SetComposite(composite Expression)
	GetIndices() Sequential[Expression]
	SetIndices(indices Sequential[Expression])
}

// This interface defines the methods supported by all variable-like expressions.
type VariableLike interface {
	GetIdentifier() string
	SetIdentifier(identifier string)
}

// CONSTANTS

type Operator int

const (
	_          Operator = iota // Invalid
	ASSIGN                     // :=
	DEFAULT                    // ?=
	SUM                        // +=
	DIFFERENCE                 // -=
	PRODUCT                    // *=
	QUOTIENT                   // /=
	DOT                        // .
	ARROW                      // <-
	AMPERSAND                  // &
	AT                         // @
	BAR                        // |
	TILDA                      // ~
	PLUS                       // +
	MINUS                      // -
	STAR                       // *
	SLASH                      // /
	MODULO                     // //
	CARET                      // ^
	LESS                       // <
	EQUAL                      // =
	UNEQUAL                    // ≠
	MORE                       // >
	IS                         // IS
	MATCHES                    // MATCHES
	NOT                        // NOT
	AND                        // AND
	SANS                       // SANS
	OR                         // OR
	XOR                        // XOR
)
