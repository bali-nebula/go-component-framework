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

type ExpressionLike any

// This interface defines the methods supported by all arithmetic-like expressions.
type ArithmeticLike interface {
	GetFirst() ExpressionLike
	SetFirst(first ExpressionLike)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() ExpressionLike
	SetSecond(second ExpressionLike)
}

// This interface defines the methods supported by all chaining-like expressions.
type ChainingLike interface {
	GetFirst() ExpressionLike
	SetFirst(first ExpressionLike)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() ExpressionLike
	SetSecond(second ExpressionLike)
}

// This interface defines the methods supported by all comparison-like expressions.
type ComparisonLike interface {
	GetFirst() ExpressionLike
	SetFirst(first ExpressionLike)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() ExpressionLike
	SetSecond(second ExpressionLike)
}

// This interface defines the methods supported by all complement-like expressions.
type ComplementLike interface {
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExpression() ExpressionLike
	SetExpression(expression ExpressionLike)
}

// This interface defines the methods supported by all dereference-like expressions.
type DereferenceLike interface {
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExpression() ExpressionLike
	SetExpression(expression ExpressionLike)
}

// This interface defines the methods supported by all exponential-like expressions.
type ExponentialLike interface {
	GetBase() ExpressionLike
	SetBase(base ExpressionLike)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExponent() ExpressionLike
	SetExponent(exponent ExpressionLike)
}

// This interface defines the methods supported by all intrinsic-like expressions.
type IntrinsicLike interface {
	GetFunction() string
	SetFunction(function string)
	GetArgument(index int) ExpressionLike           // Ordinal based indexing.
	SetArgument(index int, argument ExpressionLike) // Ordinal based indexing.
	GetArguments() ListLike[ExpressionLike]
	SetArguments(arguments ListLike[ExpressionLike])
}

// This interface defines the methods supported by all inversion-like expressions.
type InversionLike interface {
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExpression() ExpressionLike
	SetExpression(expression ExpressionLike)
}

// This interface defines the methods supported by all invocation-like expressions.
type InvocationLike interface {
	IsSynchronous() bool
	GetTarget() ExpressionLike
	SetTarget(target ExpressionLike)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetMessage() string
	SetMessage(message string)
	GetArgument(index int) ExpressionLike           // Ordinal based indexing.
	SetArgument(index int, argument ExpressionLike) // Ordinal based indexing.
	GetArguments() ListLike[ExpressionLike]
	SetArguments(arguments ListLike[ExpressionLike])
}

// This interface defines the methods supported by all logical-like expressions.
type LogicalLike interface {
	GetFirst() ExpressionLike
	SetFirst(first ExpressionLike)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() ExpressionLike
	SetSecond(second ExpressionLike)
}

// This interface defines the methods supported by all magnitude-like expressions.
type MagnitudeLike interface {
	GetExpression() ExpressionLike
	SetExpression(expression ExpressionLike)
}

// This interface defines the methods supported by all precedence-like expressions.
type PrecedenceLike interface {
	GetExpression() ExpressionLike
	SetExpression(expression ExpressionLike)
}

// This interface defines the methods supported by all item-like expressions.
type ItemLike interface {
	GetComposite() ExpressionLike
	SetComposite(composite ExpressionLike)
	GetIndex(index int) ExpressionLike
	SetIndex(index int, expression ExpressionLike)
	GetIndices() ListLike[ExpressionLike]
	SetIndices(indices ListLike[ExpressionLike])
}

// This interface defines the methods supported by all variable-like expressions.
type VariableLike interface {
	GetIdentifier() string
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
