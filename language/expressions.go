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

import (
	"fmt"
)

// This method attempts to parse the arguments within a call. It returns an
// array of the arguments and whether or not the arguments were successfully
// parsed.
func (v *parser) parseArguments() ([]any, bool) {
	var ok bool
	var argument any
	var arguments []any
	argument, ok = v.parseExpression()
	for ok {
		arguments = append(arguments, argument)
		// Every subsequent argument must be preceded by a ','.
		_, ok = v.parseDelimiter(",")
		if !ok {
			// No more arguments.
			break
		}
		argument, ok = v.parseExpression()
		if !ok {
			panic("Expected an argument after the ',' character.")
		}
	}
	return arguments, true
}

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

// This method attempts to parse an attribute expression. It returns the
// attribute expression and whether or not the attribute expression was
// successfully parsed.
func (v *parser) parseAttributeExpression() (*AttributeExpression, bool) {
	var ok bool
	var bad *token
	var composite any
	var indices []any
	var expression *AttributeExpression
	composite, ok = v.parseExpression()
	if !ok {
		// This is not an attribute expression.
		return expression, false
	}
	_, ok = v.parseDelimiter("[")
	if !ok {
		// This is not an attribute expression.
		// TODO: This wont work, call stateful expression parser here...
		v.backupOne() // Put back the variable token.
		return expression, false
	}
	indices, ok = v.parseIndices()
	if !ok {
		panic("Expected indices following the '[' character.")
	}
	bad, ok = v.parseDelimiter("]")
	if !ok {
		panic(fmt.Sprintf("Expected a ']' character following the indices and received: %v", *bad))
	}
	expression = &AttributeExpression{composite, indices}
	return expression, true
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

// This method attempts to parse a complement expression. It returns the
// complement expression and whether or not the complement expression was
// successfully parsed.
func (v *parser) parseComplementExpression() (*ComplementExpression, bool) {
	var ok bool
	var logical any
	var expression *ComplementExpression
	_, ok = v.parseKeyword("NOT")
	if !ok {
		// This is not an complement expression.
		return expression, false
	}
	logical, ok = v.parseExpression()
	if !ok {
		panic("Expected a logical expression following the 'NOT' operator.")
	}
	expression = &ComplementExpression{logical}
	return expression, true
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

// This method attempts to parse a dereference expression. It returns the
// dereference expression and whether or not the dereference expression was
// successfully parsed.
func (v *parser) parseDereferenceExpression() (*DereferenceExpression, bool) {
	var ok bool
	var reference any
	var expression *DereferenceExpression
	_, ok = v.parseDelimiter("@")
	if !ok {
		// This is not an dereference expression.
		return expression, false
	}
	reference, ok = v.parseExpression()
	if !ok {
		panic("Expected an expression following the '@' character.")
	}
	expression = &DereferenceExpression{reference}
	return expression, true
}

// This method attempts to parse an expression. It returns the expression and
// whether or not the expression was successfully parsed. The expressions are
// are checked in precedence order from highest to lowest precedence.
func (v *parser) parseExpression() (any, bool) {
	var ok bool
	var expression any
	expression, ok = v.parseComponent()
	if !ok {
		expression, ok = v.parseFunctionExpression()
	}
	if !ok {
		expression, ok = v.parseIdentifier()
	}
	if !ok {
		expression, ok = v.parsePrecedenceExpression()
	}
	if !ok {
		expression, ok = v.parseDereferenceExpression()
	}
	/*
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

// This method attempts to parse a function expression. It returns the
// function expression and whether or not the function expression was
// successfully parsed.
func (v *parser) parseFunctionExpression() (*FunctionExpression, bool) {
	var ok bool
	var bad *token
	var function string
	var arguments []any
	var expression *FunctionExpression
	function, ok = v.parseIdentifier()
	if !ok {
		// This is not an function expression.
		return expression, false
	}
	_, ok = v.parseDelimiter("(")
	if !ok {
		// This is not an function expression.
		v.backupOne() // Put back the identifier token.
		return expression, false
	}
	arguments, ok = v.parseArguments()
	if !ok {
		panic("Expected arguments following the '(' character.")
	}
	bad, ok = v.parseDelimiter(")")
	if !ok {
		panic(fmt.Sprintf("Expected a ')' character following the arguments and received: %v", *bad))
	}
	expression = &FunctionExpression{function, arguments}
	return expression, true
}

// This type defines the node structure associated with an expression that
// returns the inverse of a value.
type InversionExpression struct {
	Operator string
	Numeric  any
}

// This method attempts to parse a inversion expression. It returns the
// inversion expression and whether or not the inversion expression was
// successfully parsed.
func (v *parser) parseInversionExpression() (*InversionExpression, bool) {
	var ok bool
	var t *token
	var operator string
	var numeric any
	var expression *InversionExpression
	t, ok = v.parseDelimiter("-")
	if !ok {
		t, ok = v.parseDelimiter("/")
		if !ok {
			t, ok = v.parseDelimiter("*")
			if !ok {
				// This is not an inversion expression.
				return expression, false
			}
		}
	}
	operator = t.val
	numeric, ok = v.parseExpression()
	if !ok {
		panic(fmt.Sprintf("Expected a numeric expression following the %q operator.", operator))
	}
	expression = &InversionExpression{operator, numeric}
	return expression, true
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

// This method attempts to parse a precedence expression. It returns the
// precedence expression and whether or not the precedence expression was
// successfully parsed.
func (v *parser) parsePrecedenceExpression() (*PrecedenceExpression, bool) {
	var ok bool
	var bad *token
	var inner any
	var expression *PrecedenceExpression
	_, ok = v.parseDelimiter("(")
	if !ok {
		// This is not an precedence expression.
		return expression, false
	}
	inner, ok = v.parseExpression()
	if !ok {
		panic("Expected an expression following the '(' character.")
	}
	bad, ok = v.parseDelimiter(")")
	if !ok {
		panic(fmt.Sprintf("Expected a ')' character following the expression and received: %v", *bad))
	}
	expression = &PrecedenceExpression{inner}
	return expression, true
}

// This type defines the node structure associated with an expression that
// returns the value of a variable.
type VariableExpression struct {
	Variable string
}
