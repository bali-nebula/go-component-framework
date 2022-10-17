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
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/expressions"
)

// This method attempts to parse the arguments within a call. It returns an
// array of the arguments and whether or not the arguments were successfully
// parsed.
func (v *parser) parseArguments() (abstractions.ListLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var argument any
	var arguments abstractions.ListLike[any]
	argument, token, ok = v.parseExpression()
	for ok {
		arguments.AddItem(argument)
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

// This method attempts to parse a arithmetic expression. It returns the
// arithmetic expression and whether or not the arithmetic expression was
// successfully parsed.
func (v *parser) parseArithmetic(left any) (abstractions.ArithmeticLike, *Token, bool) {
	var ok bool
	var token *Token
	var delimiter string
	var operator abstractions.Operator
	var right any
	var expression abstractions.ArithmeticLike
	delimiter, token, ok = v.parseDelimiter("*")
	if !ok {
		delimiter, token, ok = v.parseDelimiter("/")
	}
	if !ok {
		delimiter, token, ok = v.parseDelimiter("//")
	}
	if !ok {
		delimiter, token, ok = v.parseDelimiter("+")
	}
	if !ok {
		delimiter, token, ok = v.parseDelimiter("-")
	}
	if !ok {
		// This is not an arithmetic expression.
		return expression, token, false
	}
	right, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '"+delimiter+"' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$arithmetic")
		panic(message)
	}
	switch delimiter {
	case "*":
		operator = abstractions.PRODUCT
	case "/":
		operator = abstractions.QUOTIENT
	case "//":
		operator = abstractions.REMAINDER
	case "+":
		operator = abstractions.SUM
	case "-":
		operator = abstractions.DIFFERENCE
	}
	expression = expressions.Arithmetic(left, operator, right)
	return expression, token, true
}

// This method attempts to parse a chain expression. It returns the
// chain expression and whether or not the chain expression was
// successfully parsed.
func (v *parser) parseChaining(left any) (abstractions.ChainingLike, *Token, bool) {
	var ok bool
	var token *Token
	var right any
	var expression abstractions.ChainingLike
	_, token, ok = v.parseDelimiter("&")
	if !ok {
		// This is not a chain expression.
		return expression, token, false
	}
	right, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '&' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$chaining")
		panic(message)
	}
	expression = expressions.Chaining(left, right)
	return expression, token, true
}

// This method attempts to parse a comparison expression. It returns the
// comparison expression and whether or not the comparison expression was
// successfully parsed.
func (v *parser) parseComparison(left any) (abstractions.ComparisonLike, *Token, bool) {
	var ok bool
	var token *Token
	var delimiter string
	var operator abstractions.Operator
	var right any
	var expression abstractions.ComparisonLike
	delimiter, token, ok = v.parseDelimiter("<")
	if !ok {
		delimiter, token, ok = v.parseDelimiter("=")
	}
	if !ok {
		delimiter, token, ok = v.parseDelimiter(">")
	}
	if !ok {
		delimiter, token, ok = v.parseDelimiter("≠")
	}
	if !ok {
		delimiter, token, ok = v.parseKeyword("IS")
	}
	if !ok {
		delimiter, token, ok = v.parseKeyword("MATCHES")
	}
	if !ok {
		// This is not a comparison expression.
		return expression, token, false
	}
	right, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '"+delimiter+"' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$comparison")
		panic(message)
	}
	switch delimiter {
	case "<":
		operator = abstractions.LESS
	case "=":
		operator = abstractions.EQUAL
	case ">":
		operator = abstractions.MORE
	case "≠":
		operator = abstractions.UNEQUAL
	case "IS":
		operator = abstractions.IS
	case "MATCHES":
		operator = abstractions.MATCHES
	}
	expression = expressions.Comparison(left, operator, right)
	return expression, token, true
}

// This method attempts to parse a complement expression. It returns the
// complement expression and whether or not the complement expression was
// successfully parsed.
func (v *parser) parseComplement() (abstractions.ComplementLike, *Token, bool) {
	var ok bool
	var token *Token
	var logical any
	var expression abstractions.ComplementLike
	_, token, ok = v.parseKeyword("NOT")
	if !ok {
		// This is not an complement expression.
		return expression, token, false
	}
	logical, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a logical expression following the 'NOT' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$complement")
		panic(message)
	}
	expression = expressions.Complement(logical)
	return expression, token, true
}

// This method attempts to parse a dereference expression. It returns the
// dereference expression and whether or not the dereference expression was
// successfully parsed.
func (v *parser) parseDereference() (abstractions.DereferenceLike, *Token, bool) {
	var ok bool
	var token *Token
	var reference any
	var expression abstractions.DereferenceLike
	_, token, ok = v.parseDelimiter("@")
	if !ok {
		// This is not an dereference expression.
		return expression, token, false
	}
	reference, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '@' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$dereference")
		panic(message)
	}
	expression = expressions.Dereference(reference)
	return expression, token, true
}

// This method attempts to parse a power expression. It returns the
// power expression and whether or not the power expression was
// successfully parsed.
func (v *parser) parseExponential(base any) (abstractions.ExponentialLike, *Token, bool) {
	var ok bool
	var token *Token
	var exponent any
	var expression abstractions.ExponentialLike
	_, token, ok = v.parseDelimiter("^")
	if !ok {
		// This is not a power expression.
		return expression, token, false
	}
	exponent, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an exponent expression following the '^' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$exponential")
		panic(message)
	}
	expression = expressions.Exponential(base, exponent)
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
		expression, token, ok = v.parseIntrinsic()
	}
	if !ok {
		expression, token, ok = v.parseIdentifier()
	}
	if !ok {
		expression, token, ok = v.parsePrecedence()
	}
	if !ok {
		expression, token, ok = v.parseDereference()
	}
	if !ok {
		expression, token, ok = v.parseRecursive()
	}
	if !ok {
		expression, token, ok = v.parseInversion()
	}
	if !ok {
		expression, token, ok = v.parseMagnitude()
	}
	if !ok {
		expression, token, ok = v.parseComplement()
	}
	return expression, token, ok
}

// This method attempts to parse a sequence of indices. It returns an array of
// the indices and whether or not the indices were successfully parsed.
func (v *parser) parseIndices() (abstractions.ListLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var index any
	var indices abstractions.ListLike[any]
	index, token, ok = v.parseExpression()
	// There must be at least one index.
	if !ok {
		var message = fmt.Sprintf("Expected at least one index in the sequence of indices but received:\n%v\n\n", token)
		message += generateGrammar(
			"$indices")
		panic(message)
	}
	for {
		indices.AddItem(index)
		// Every subsequent index must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more indices.
			break
		}
		index, token, ok = v.parseExpression()
		if !ok {
			var message = fmt.Sprintf("Expected an index after the ',' character but received:\n%v\n\n", token)
			message += generateGrammar(
				"$indices")
			panic(message)
		}
	}
	return indices, token, true
}

// This method attempts to parse a function expression. It returns the
// function expression and whether or not the function expression was
// successfully parsed.
func (v *parser) parseIntrinsic() (abstractions.IntrinsicLike, *Token, bool) {
	var ok bool
	var token *Token
	var function string
	var arguments abstractions.ListLike[any]
	var expression abstractions.IntrinsicLike
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
			"$intrinsic",
			"$function",
			"$arguments")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = fmt.Sprintf("Expected a ')' character following the arguments but received:\n%v\n\n", token)
		message += generateGrammar(
			"$intrinsic",
			"$function",
			"$arguments")
		panic(message)
	}
	expression = expressions.Intrinsic(function, arguments)
	return expression, token, true
}

// This method attempts to parse a inversion expression. It returns the
// inversion expression and whether or not the inversion expression was
// successfully parsed.
func (v *parser) parseInversion() (abstractions.InversionLike, *Token, bool) {
	var ok bool
	var token *Token
	var delimiter string
	var operator abstractions.Operator
	var numeric any
	var expression abstractions.InversionLike
	delimiter, token, ok = v.parseDelimiter("-")
	if !ok {
		delimiter, token, ok = v.parseDelimiter("/")
	}
	if !ok {
		delimiter, token, ok = v.parseDelimiter("*")
	}
	if !ok {
		// This is not an inversion expression.
		return expression, token, false
	}
	numeric, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a numeric expression following the '"+delimiter+"' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$inversion")
		panic(message)
	}
	switch delimiter {
	case "-":
		operator = abstractions.INVERSE
	case "/":
		operator = abstractions.RECIPROCAL
	case "*":
		operator = abstractions.CONJUGATE
	}
	expression = expressions.Inversion(operator, numeric)
	return expression, token, true
}

// This method attempts to parse a message expression. It returns the
// message expression and whether or not the message expression was
// successfully parsed.
func (v *parser) parseInvocation(target any) (abstractions.InvocationLike, *Token, bool) {
	var ok bool
	var token *Token
	var delimiter string
	var message string
	var arguments abstractions.ListLike[any]
	var expression abstractions.InvocationLike
	delimiter, token, ok = v.parseDelimiter(".")
	if !ok {
		delimiter, token, ok = v.parseDelimiter("<~")
	}
	if !ok {
		// This is not an message expression.
		return expression, token, false
	}
	message, token, ok = v.parseIdentifier()
	if !ok {
		var message = fmt.Sprintf("Expected a message identifier following the '"+delimiter+"' delimiter but received:\n%v\n\n", token)
		message += generateGrammar(
			"$invocation",
			"$message",
			"$arguments")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		var message = fmt.Sprintf("Expected a '(' character following the message identifier but received:\n%v\n\n", token)
		message += generateGrammar(
			"$invocation",
			"$message",
			"$arguments")
		panic(message)
	}
	arguments, token, ok = v.parseArguments()
	if !ok {
		var message = fmt.Sprintf("Expected arguments following the '(' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$invocation",
			"$message",
			"$arguments")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = fmt.Sprintf("Expected a ')' character following the arguments but received:\n%v\n\n", token)
		message += generateGrammar(
			"$invocation",
			"$message",
			"$arguments")
		panic(message)
	}
	switch delimiter {
	case ".":
		expression = expressions.Invocation(target, message, arguments)
	case "<~":
		expression = expressions.AsynchronousInvocation(target, message, arguments)
	}
	return expression, token, true
}

// This method attempts to parse a logical expression. It returns the
// logical expression and whether or not the logical expression was
// successfully parsed.
func (v *parser) parseLogical(left any) (abstractions.LogicalLike, *Token, bool) {
	var ok bool
	var token *Token
	var delimiter string
	var operator abstractions.Operator
	var right any
	var expression abstractions.LogicalLike
	delimiter, token, ok = v.parseKeyword("AND")
	if !ok {
		delimiter, token, ok = v.parseKeyword("SANS")
	}
	if !ok {
		delimiter, token, ok = v.parseKeyword("XOR")
	}
	if !ok {
		delimiter, token, ok = v.parseKeyword("OR")
	}
	if !ok {
		// This is not a logical expression.
		return expression, token, false
	}
	right, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a logical expression following the '"+delimiter+"' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$logical")
		panic(message)
	}
	switch delimiter {
	case "AND":
		operator = abstractions.AND
	case "SANS":
		operator = abstractions.SANS
	case "XOR":
		operator = abstractions.XOR
	case "OR":
		operator = abstractions.OR
	}
	expression = expressions.Logical(left, operator, right)
	return expression, token, true
}

// This method attempts to parse a magnitude expression. It returns the
// magnitude expression and whether or not the magnitude expression was
// successfully parsed.
func (v *parser) parseMagnitude() (abstractions.MagnitudeLike, *Token, bool) {
	var ok bool
	var token *Token
	var numeric any
	var expression abstractions.MagnitudeLike
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		// This is not an magnitude expression.
		return expression, token, false
	}
	numeric, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a numeric expression following the '|' operator but received:\n%v\n\n", token)
		message += generateGrammar(
			"$magnitude")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		var message = fmt.Sprintf("Expected a '|' operator following the numeric expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$magnitude")
		panic(message)
	}
	expression = expressions.Magnitude(numeric)
	return expression, token, true
}

// This method attempts to parse a precedence expression. It returns the
// precedence expression and whether or not the precedence expression was
// successfully parsed.
func (v *parser) parsePrecedence() (abstractions.PrecedenceLike, *Token, bool) {
	var ok bool
	var token *Token
	var inner any
	var expression abstractions.PrecedenceLike
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		// This is not an precedence expression.
		return expression, token, false
	}
	inner, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression following the '(' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$precedence")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = fmt.Sprintf("Expected a ')' character following the expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$precedence")
		panic(message)
	}
	expression = expressions.Precedence(inner)
	return expression, token, true
}

// This method attempts to parse a left recursive expression. It returns
// the left recursive expression and whether or not the left recursive
// expression was successfully parsed.
func (v *parser) parseRecursive() (any, *Token, bool) {
	var ok bool
	var token *Token
	var expression any
	expression, token, ok = v.parseExpression()
	if !ok {
		// This is not a left recursive expression.
		return expression, token, false
	}
	expression, token, ok = v.parseInvocation(expression)
	if !ok {
		expression, token, ok = v.parseValue(expression)
	}
	if !ok {
		expression, token, ok = v.parseChaining(expression)
	}
	if !ok {
		expression, token, ok = v.parseExponential(expression)
	}
	if !ok {
		expression, token, ok = v.parseArithmetic(expression)
	}
	if !ok {
		expression, token, ok = v.parseComparison(expression)
	}
	if !ok {
		expression, token, ok = v.parseLogical(expression)
	}
	if !ok {
		var message = fmt.Sprintf("Expected a valid operator after the expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$invocation",
			"$value",
			"$chaining",
			"$exponential",
			"$arithmetic",
			"$comparison",
			"$logical")
		panic(message)
	}
	return expression, token, true
}

// This method attempts to parse an attribute expression. It returns the
// attribute expression and whether or not the attribute expression was
// successfully parsed.
func (v *parser) parseValue(composite any) (abstractions.ValueLike, *Token, bool) {
	var ok bool
	var token *Token
	var indices abstractions.ListLike[any]
	var expression abstractions.ValueLike
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		// This is not an attribute expression.
		return expression, token, false
	}
	indices, token, ok = v.parseIndices()
	if !ok {
		var message = fmt.Sprintf("Expected indices following the '[' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$value",
			"$indices")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = fmt.Sprintf("Expected a ']' character following the indices but received:\n%v\n\n", token)
		message += generateGrammar(
			"$value",
			"$indices")
		panic(message)
	}
	expression = expressions.Value(composite, indices)
	return expression, token, true
}
