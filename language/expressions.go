/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package language

// This type defines the node structure associated with an expression that
// returns the result of an arithmetic operation on two values.
type ArithmeticExpression struct {
	Left     any
	Operator string
	Right    any
}

// This type defines the node structure associated with an expression that
// returns an indexed attribute from within a composite component.
type AttributeExpression struct {
	Composite any
	Indices   []any
}

// This type defines the node structure associated with an expression that
// returns the concatenation of two values.
type ChainExpression struct {
	Left  any
	Right any
}

// This type defines the node structure associated with an expression that
// returns a comparison of two values.
type ComparisonExpression struct {
	Left     any
	Operator string
	Right    any
}

// This type defines the node structure associated with an expression that
// returns the logical complement of a value.
type ComplementExpression struct {
	Logical any
}

// This type defines the node structure associated with an expression made up of
// a single component.
type ComponentExpression struct {
	Component *Component
}

// This type defines the node structure associated with an expression that
// returns a value or its default value if not set.
type DefaultExpression struct {
	Left  any
	Right any
}

// This type defines the node structure associated with an expression that
// returns the value that the expression references.
type DereferenceExpression struct {
	Expression any
}

// This method attempts to parse an expression. It returns the expression and
// whether or not the expression was successfully parsed.
func (v *parser) parseExpression() (any, bool) {
	// TODO: Reorder these based on how often each type of expression occurs.
	var ok bool
	var expression any
	expression, ok = v.parseComponent()
	/*
		if !ok {
			expression, ok = v.parseVariable()
		}
		if !ok {
			expression, ok = v.parseFunctionExpression()
		}
		if !ok {
			expression, ok = v.parsePrecedenceExpression()
		}
		if !ok {
			expression, ok = v.parseDereferenceExpression()
		}
		if !ok {
			expression, ok = v.parseMessageExpression()
		}
		if !ok {
			expression, ok = v.parseAttributeExpression()
		}
		if !ok {
			expression, ok = v.parseChainExpression()
		}
		if !ok {
			expression, ok = v.parsePowerExpression()
		}
		if !ok {
			expression, ok = v.parseInversionExpression()
		}
		if !ok {
			expression, ok = v.parseArithmeticExpression()
		}
		if !ok {
			expression, ok = v.parseMagnitudeExpression()
		}
		if !ok {
			expression, ok = v.parseComparisonExpression()
		}
		if !ok {
			expression, ok = v.parseComplementExpression()
		}
		if !ok {
			expression, ok = v.parseLogicalExpression()
		}
		if !ok {
			expression, ok = v.parseDefaultExpression()
		}
	*/
	return expression, ok
}

// This type defines the node structure associated with an expression that
// returns the result of a function call.
type FunctionExpression struct {
	Function  string
	Arguments []any
}

// This type defines the node structure associated with an expression that
// returns the inverse of a value.
type InversionExpression struct {
	Operator string
	Numeric  any
}

// This type defines the node structure associated with an expression that
// returns the result of a logical operation on two values.
type LogicalExpression struct {
	Left     any
	Operator string
	Right    any
}

// This type defines the node structure associated with an expression that
// returns the magnitude of a value.
type MagnitudeExpression struct {
	Numeric any
}

// This type defines the node structure associated with an expression that
// sends a message to a target component. The message can be sent synchronously
// or asynchronously.
type MessageExpression struct {
	Target    any
	Operator  string
	Message   string
	Arguments []any
}

// This type defines the node structure associated with an expression that
// returns the power of a base value raised to an exponential value.
type PowerExpression struct {
	Base     any
	Exponent any
}

// This type defines the node structure associated with an expression that
// takes precdence over its surrounding expressions.
type PrecedenceExpression struct {
	Expression any
}

// This type defines the node structure associated with an expression that
// returns the value of a variable.
type VariableExpression struct {
	Variable string
}
