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
	Operator   int
)

// CONSTANT DEFINITIONS

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
	MAGNITUDE                  // | |
	PRECEDENCE                 // ( )
)

// INDIVIDUAL INTERFACES

type BinaryOperationLike interface {
	GetFirst() Expression
	SetFirst(first Expression)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetSecond() Expression
	SetSecond(second Expression)
}

type IntrinsicLike interface {
	GetFunction() string
	SetFunction(function string)
	GetArguments() Sequential[Expression]
	SetArguments(arguments Sequential[Expression])
}

type InvocationLike interface {
	IsSynchronous() bool
	GetTarget() Expression
	SetTarget(target Expression)
	GetOperator() Operator
	SetOperator(operator Operator)
	GetMethod() string
	SetMethod(method string)
	GetArguments() Sequential[Expression]
	SetArguments(arguments Sequential[Expression])
}

type SubcomponentLike interface {
	GetComposite() Expression
	SetComposite(composite Expression)
	GetIndices() Sequential[Expression]
	SetIndices(indices Sequential[Expression])
}

type UnaryOperationLike interface {
	GetOperator() Operator
	SetOperator(operator Operator)
	GetExpression() Expression
	SetExpression(expression Expression)
}

type ValueLike interface {
	GetComponent() ComponentLike
	SetComponent(component ComponentLike)
}

type VariableLike interface {
	GetIdentifier() string
	SetIdentifier(identifier string)
}
