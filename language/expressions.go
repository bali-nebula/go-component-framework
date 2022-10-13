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
func (v *parser) parseArguments() ([]any, *Token, bool) {
	var ok bool
	var token *Token
	var argument any
	var arguments []any
	argument, token, ok = v.parseExpression()
	for ok {
		arguments = append(arguments, argument)
		// Every subsequent argument must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more arguments.
			break
		}
		argument, token, ok = v.parseExpression()
		if !ok {
			var message = fmt.Sprintf("Expected an argument after the ',' character but received:\n%v\n\n", token)
			message += generateGrammar(
				"$arguments")
			panic(message)
		}
	}
	return arguments, token, true
}

// This type defines the node structure associated with an expression that
// returns the result of an arithmetic operation on two values.
type ArithmeticExpression struct {
	Left     any
	Operator string
	Right    any
}

// This method attempts to parse a arithmetic expression. It returns the
// arithmetic expression and whether or not the arithmetic expression was
// successfully parsed.
func (v *parser) parseArithmeticExpression(left any) (*ArithmeticExpression, *Token, bool) {
	var ok bool
	var token *Token
	var operator string
	var right any
	var expression *ArithmeticExpression
	operator, token, ok = v.parseDelimiter("*")
	if !ok {
		operator, token, ok = v.parseDelimiter("/")
	}
	if !ok {
		operator, token, ok = v.parseDelimiter("//")
	}
	if !ok {
		operator, token, ok = v.parseDelimiter("+")
	}
	if !ok {
		operator, token, ok = v.parseDelimiter("-")
	}
	if !ok {
		// This is not an arithmetic expression.
		return expression, token, false
	}
	right, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '" + operator + "' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$arithmeticExpression")
		panic(message)
	}
	expression = &ArithmeticExpression{left, operator, right}
	return expression, token, true
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
func (v *parser) parseAttributeExpression(composite any) (*AttributeExpression, *Token, bool) {
	var ok bool
	var token *Token
	var indices []any
	var expression *AttributeExpression
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		// This is not an attribute expression.
		return expression, token, false
	}
	indices, token, ok = v.parseIndices()
	if !ok {
		var message = fmt.Sprintf("Expected indices following the '[' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$attributeExpression",
			"$indices")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = fmt.Sprintf("Expected a ']' character following the indices but received:\n%v\n\n", token)
		message += generateGrammar(
			"$attributeExpression",
			"$indices")
		panic(message)
	}
	expression = &AttributeExpression{composite, indices}
	return expression, token, true
}

// This type defines the node structure associated with an expression that
// returns the concatenation of two values.
type ChainExpression struct {
	Left  any
	Right any
}

// This method attempts to parse a chain expression. It returns the
// chain expression and whether or not the chain expression was
// successfully parsed.
func (v *parser) parseChainExpression(left any) (*ChainExpression, *Token, bool) {
	var ok bool
	var token *Token
	var right any
	var expression *ChainExpression
	_, token, ok = v.parseDelimiter("&")
	if !ok {
		// This is not a chain expression.
		return expression, token, false
	}
	right, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '&' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$chainExpression")
		panic(message)
	}
	expression = &ChainExpression{left, right}
	return expression, token, true
}

// This type defines the node structure associated with an expression that
// returns a comparison of two values.
type ComparisonExpression struct {
	Left     any
	Operator string
	Right    any
}

// This method attempts to parse a comparison expression. It returns the
// comparison expression and whether or not the comparison expression was
// successfully parsed.
func (v *parser) parseComparisonExpression(left any) (*ComparisonExpression, *Token, bool) {
	var ok bool
	var token *Token
	var operator string
	var right any
	var expression *ComparisonExpression
	operator, token, ok = v.parseDelimiter("<")
	if !ok {
		operator, token, ok = v.parseDelimiter("=")
	}
	if !ok {
		operator, token, ok = v.parseDelimiter(">")
	}
	if !ok {
		operator, token, ok = v.parseDelimiter("≠")
	}
	if !ok {
		operator, token, ok = v.parseKeyword("IS")
	}
	if !ok {
		operator, token, ok = v.parseKeyword("MATCHES")
	}
	if !ok {
		// This is not a comparison expression.
		return expression, token, false
	}
	right, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '" + operator + "' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$comparisonExpression")
		panic(message)
	}
	expression = &ComparisonExpression{left, operator, right}
	return expression, token, true
}

// This type defines the node structure associated with an expression that
// returns the logical complement of a value.
type ComplementExpression struct {
	Logical any
}

// This method attempts to parse a complement expression. It returns the
// complement expression and whether or not the complement expression was
// successfully parsed.
func (v *parser) parseComplementExpression() (*ComplementExpression, *Token, bool) {
	var ok bool
	var token *Token
	var logical any
	var expression *ComplementExpression
	_, token, ok = v.parseKeyword("NOT")
	if !ok {
		// This is not an complement expression.
		return expression, token, false
	}
	logical, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a logical expression following the 'NOT' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$complementExpression")
		panic(message)
	}
	expression = &ComplementExpression{logical}
	return expression, token, true
}

// This type defines the node structure associated with an expression made up of
// a single component.
type ComponentExpression struct {
	Component *Component
}

// This type defines the node structure associated with an expression that
// returns a value or its default value if not set.
type DefaultExpression struct {
	Value   any
	Default any
}

// This method attempts to parse a default expression. It returns the
// default expression and whether or not the default expression was
// successfully parsed.
func (v *parser) parseDefaultExpression(value any) (*DefaultExpression, *Token, bool) {
	var ok bool
	var token *Token
	var defaultValue any
	var expression *DefaultExpression
	_, token, ok = v.parseDelimiter("?")
	if !ok {
		// This is not a default expression.
		return expression, token, false
	}
	defaultValue, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a default expression following the '?' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$defaultExpression")
		panic(message)
	}
	expression = &DefaultExpression{value, defaultValue}
	return expression, token, true
}

// This type defines the node structure associated with an expression that
// returns the value that the expression references.
type DereferenceExpression struct {
	Expression any
}

// This method attempts to parse a dereference expression. It returns the
// dereference expression and whether or not the dereference expression was
// successfully parsed.
func (v *parser) parseDereferenceExpression() (*DereferenceExpression, *Token, bool) {
	var ok bool
	var token *Token
	var reference any
	var expression *DereferenceExpression
	_, token, ok = v.parseDelimiter("@")
	if !ok {
		// This is not an dereference expression.
		return expression, token, false
	}
	reference, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '@' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$dereferenceExpression")
		panic(message)
	}
	expression = &DereferenceExpression{reference}
	return expression, token, true
}

// This method attempts to parse an expression. It returns the expression and
// whether or not the expression was successfully parsed. The expressions are
// are checked in precedence order from highest to lowest precedence.
func (v *parser) parseExpression() (any, *Token, bool) {
	var ok bool
	var token *Token
	var expression any
	expression, token, ok = v.parseComponent()
	if !ok {
		// This must come before the parseIdentifier() for a variable.
		expression, token, ok = v.parseFunctionExpression()
	}
	if !ok {
		expression, token, ok = v.parseIdentifier()
	}
	if !ok {
		expression, token, ok = v.parsePrecedenceExpression()
	}
	if !ok {
		expression, token, ok = v.parseDereferenceExpression()
	}
	if !ok {
		expression, token, ok = v.parseRecursiveExpression()
	}
	if !ok {
		expression, token, ok = v.parseInversionExpression()
	}
	if !ok {
		expression, token, ok = v.parseMagnitudeExpression()
	}
	if !ok {
		expression, token, ok = v.parseComplementExpression()
	}
	return expression, token, ok
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
func (v *parser) parseFunctionExpression() (*FunctionExpression, *Token, bool) {
	var ok bool
	var token *Token
	var function string
	var arguments []any
	var expression *FunctionExpression
	function, token, ok = v.parseIdentifier()
	if !ok {
		// This is not an function expression.
		return expression, token, false
	}
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		// This is not an function expression.
		v.backupOne() // Put back the identifier token.
		return expression, token, false
	}
	arguments, token, ok = v.parseArguments()
	if !ok {
		var message = fmt.Sprintf("Expected arguments following the '(' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$functionExpression",
			"$function",
			"$arguments")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = fmt.Sprintf("Expected a ')' character following the arguments but received:\n%v\n\n", token)
		message += generateGrammar(
			"$functionExpression",
			"$function",
			"$arguments")
		panic(message)
	}
	expression = &FunctionExpression{function, arguments}
	return expression, token, true
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
func (v *parser) parseInversionExpression() (*InversionExpression, *Token, bool) {
	var ok bool
	var token *Token
	var operator string
	var numeric any
	var expression *InversionExpression
	operator, token, ok = v.parseDelimiter("-")
	if !ok {
		operator, token, ok = v.parseDelimiter("/")
	}
	if !ok {
		operator, token, ok = v.parseDelimiter("*")
	}
	if !ok {
		// This is not an inversion expression.
		return expression, token, false
	}
	numeric, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a numeric expression following the '" + operator + "' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$inversionExpression")
		panic(message)
	}
	expression = &InversionExpression{operator, numeric}
	return expression, token, true
}

// This method attempts to parse a left recursive expression. It returns
// the left recursive expression and whether or not the left recursive
// expression was successfully parsed.
func (v *parser) parseRecursiveExpression() (any, *Token, bool) {
	var ok bool
	var token *Token
	var expression any
	expression, token, ok = v.parseExpression()
	if !ok {
		// This is not a left recursive expression.
		return expression, token, false
	}
	expression, token, ok = v.parseMessageExpression(expression)
	if !ok {
		expression, token, ok = v.parseAttributeExpression(expression)
	}
	if !ok {
		expression, token, ok = v.parseChainExpression(expression)
	}
	if !ok {
		expression, token, ok = v.parsePowerExpression(expression)
	}
	if !ok {
		expression, token, ok = v.parseArithmeticExpression(expression)
	}
	if !ok {
		expression, token, ok = v.parseComparisonExpression(expression)
	}
	if !ok {
		expression, token, ok = v.parseLogicalExpression(expression)
	}
	if !ok {
		expression, token, ok = v.parseDefaultExpression(expression)
	}
	if !ok {
		var message = fmt.Sprintf("Expected a valid operator after the expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$messageExpression",
			"$attributeExpression",
			"$chainExpression",
			"$powerExpression",
			"$arithmeticExpression",
			"$comparisonExpression",
			"$logicalExpression",
			"$defaultExpression")
		panic(message)
	}
	return expression, token, true
}

// This type defines the node structure associated with an expression that
// returns the result of a logical operation on two values.
type LogicalExpression struct {
	Left     any
	Operator string
	Right    any
}

// This method attempts to parse a logical expression. It returns the
// logical expression and whether or not the logical expression was
// successfully parsed.
func (v *parser) parseLogicalExpression(left any) (*LogicalExpression, *Token, bool) {
	var ok bool
	var token *Token
	var operator string
	var right any
	var expression *LogicalExpression
	operator, token, ok = v.parseKeyword("AND")
	if !ok {
		operator, token, ok = v.parseKeyword("SANS")
	}
	if !ok {
		operator, token, ok = v.parseKeyword("XOR")
	}
	if !ok {
		operator, token, ok = v.parseKeyword("OR")
	}
	if !ok {
		// This is not a logical expression.
		return expression, token, false
	}
	right, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a logical expression following the '" + operator + "' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$logicalExpression")
		panic(message)
	}
	expression = &LogicalExpression{left, operator, right}
	return expression, token, true
}

// This type defines the node structure associated with an expression that
// returns the magnitude of a value.
type MagnitudeExpression struct {
	Numeric any
}

// This method attempts to parse a magnitude expression. It returns the
// magnitude expression and whether or not the magnitude expression was
// successfully parsed.
func (v *parser) parseMagnitudeExpression() (*MagnitudeExpression, *Token, bool) {
	var ok bool
	var token *Token
	var numeric any
	var expression *MagnitudeExpression
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		// This is not an magnitude expression.
		return expression, token, false
	}
	numeric, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a numeric expression following the '|' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$magnitudeExpression")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		var message = fmt.Sprintf("Expected a '|' operator following the numeric expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$magnitudeExpression")
		panic(message)
	}
	expression = &MagnitudeExpression{numeric}
	return expression, token, true
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

// This method attempts to parse a message expression. It returns the
// message expression and whether or not the message expression was
// successfully parsed.
func (v *parser) parseMessageExpression(target any) (*MessageExpression, *Token, bool) {
	var ok bool
	var token *Token
	var operator string
	var message string
	var arguments []any
	var expression *MessageExpression
	operator, token, ok = v.parseDelimiter(".")
	if !ok {
		operator, token, ok = v.parseDelimiter("<~")
	}
	if !ok {
		// This is not an message expression.
		return expression, token, false
	}
	message, token, ok = v.parseIdentifier()
	if !ok {
		var message = fmt.Sprintf("Expected a message identifier following the '" + operator + "' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$messageExpression",
			"$message",
			"$arguments")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		var message = fmt.Sprintf("Expected a '(' character following the message identifier but received:\n%v\n\n", token)
		message += generateGrammar(
			"$messageExpression",
			"$message",
			"$arguments")
		panic(message)
	}
	arguments, token, ok = v.parseArguments()
	if !ok {
		var message = fmt.Sprintf("Expected arguments following the '(' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$messageExpression",
			"$message",
			"$arguments")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = fmt.Sprintf("Expected a ')' character following the arguments but received:\n%v\n\n", token)
		message += generateGrammar(
			"$messageExpression",
			"$message",
			"$arguments")
		panic(message)
	}
	expression = &MessageExpression{target, operator, message, arguments}
	return expression, token, true
}

// This type defines the node structure associated with an expression that
// returns the power of a base value raised to an exponential value.
type PowerExpression struct {
	Base     any
	Exponent any
}

// This method attempts to parse a power expression. It returns the
// power expression and whether or not the power expression was
// successfully parsed.
func (v *parser) parsePowerExpression(base any) (*PowerExpression, *Token, bool) {
	var ok bool
	var token *Token
	var exponent any
	var expression *PowerExpression
	_, token, ok = v.parseDelimiter("^")
	if !ok {
		// This is not a power expression.
		return expression, token, false
	}
	exponent, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an exponent expression following the '^' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$powerExpression")
		panic(message)
	}
	expression = &PowerExpression{base, exponent}
	return expression, token, true
}

// This type defines the node structure associated with an expression that
// takes precdence over its surrounding expressions.
type PrecedenceExpression struct {
	Expression any
}

// This method attempts to parse a precedence expression. It returns the
// precedence expression and whether or not the precedence expression was
// successfully parsed.
func (v *parser) parsePrecedenceExpression() (*PrecedenceExpression, *Token, bool) {
	var ok bool
	var token *Token
	var inner any
	var expression *PrecedenceExpression
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		// This is not an precedence expression.
		return expression, token, false
	}
	inner, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '(' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$precedenceExpression")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = fmt.Sprintf("Expected a ')' character following the expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$precedenceExpression")
		panic(message)
	}
	expression = &PrecedenceExpression{inner}
	return expression, token, true
}
