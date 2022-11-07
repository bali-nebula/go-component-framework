/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	col "github.com/craterdog-bali/go-bali-document-notation/collections"
	exp "github.com/craterdog-bali/go-bali-document-notation/expressions"
)

// This method attempts to parse the arguments within a call. It returns a
// list of the arguments and whether or not the arguments were successfully
// parsed.
func (v *parser) parseArguments() (abs.ListLike[abs.ExpressionLike], *Token, bool) {
	var ok bool
	var token *Token
	var argument abs.ExpressionLike
	var arguments = col.List[abs.ExpressionLike]()
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		// This is not an argument expression.
		return arguments, token, false
	}
	_, token, ok = v.parseDelimiter(")")
	if ok {
		// This is an empty argument expression.
		return arguments, token, true
	}
	argument, token, ok = v.parseExpression()
	for ok {
		arguments.AddValue(argument)
		// Every subsequent argument must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more arguments.
			break
		}
		argument, token, ok = v.parseExpression()
		if !ok {
			var message = v.formatError(token)
			message += generateGrammar("$expression",
				"$arguments")
			panic(message)
		}
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar(")",
			"$intrinsic",
			"$function",
			"$arguments")
		panic(message)
	}
	return arguments, token, true
}

// This method adds the canonical format for the specified sequence of arguments
// to the state of the formatter.
func (v *formatter) formatArguments(arguments abs.Sequential[abs.ExpressionLike]) {
	v.state.AppendString("(")
	var iterator = age.Iterator(arguments)
	if iterator.HasNext() {
		var argument = iterator.GetNext()
		v.formatExpression(argument)
	}
	for iterator.HasNext() {
		v.state.AppendString(", ")
		var argument = iterator.GetNext()
		v.formatExpression(argument)
	}
	v.state.AppendString(")")
}

// This method attempts to parse a arithmetic expression. It returns the
// arithmetic expression and whether or not the arithmetic expression was
// successfully parsed.
func (v *parser) parseArithmetic(first abs.ExpressionLike) (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.ExpressionLike
	var expression abs.ExpressionLike = first
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not an arithmetic expression.
		return expression, token, false
	}
	if operator < abs.PLUS || operator > abs.MODULO {
		// This is not an arithmetic expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	second, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$arithmetic")
		panic(message)
	}
	expression = exp.Arithmetic(first, operator, second)
	return expression, token, true
}

// This method adds the canonical format for the specified arithmetic expression
// to the state of the formatter.
func (v *formatter) formatArithmetic(arithmetic abs.ArithmeticLike) {
	var first = arithmetic.GetFirst()
	v.formatExpression(first)
	v.state.AppendString(" ")
	var operator = arithmetic.GetOperator()
	v.formatOperator(operator)
	v.state.AppendString(" ")
	var second = arithmetic.GetSecond()
	v.formatExpression(second)
}

// This method attempts to parse a chain expression. It returns the
// chain expression and whether or not the chain expression was
// successfully parsed.
func (v *parser) parseChaining(first abs.ExpressionLike) (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.ExpressionLike
	var expression abs.ExpressionLike = first
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not a chaining expression.
		return expression, token, false
	}
	if operator != abs.AMPERSAND {
		// This is not a chaining expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	second, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$chaining")
		panic(message)
	}
	expression = exp.Chaining(first, operator, second)
	return expression, token, true
}

// This method adds the canonical format for the specified chaining expression
// to the state of the formatter.
func (v *formatter) formatChaining(chaining abs.ChainingLike) {
	var first = chaining.GetFirst()
	v.formatExpression(first)
	v.state.AppendString(" ")
	var operator = chaining.GetOperator()
	v.formatOperator(operator)
	v.state.AppendString(" ")
	var second = chaining.GetSecond()
	v.formatExpression(second)
}

// This method attempts to parse a comparison expression. It returns the
// comparison expression and whether or not the comparison expression was
// successfully parsed.
func (v *parser) parseComparison(first abs.ExpressionLike) (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.ExpressionLike
	var expression abs.ExpressionLike = first
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not a comparison expression.
		return expression, token, false
	}
	if operator < abs.LESS || operator > abs.MATCHES {
		// This is not a comparison expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	second, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$comparison")
		panic(message)
	}
	expression = exp.Comparison(first, operator, second)
	return expression, token, true
}

// This method adds the canonical format for the specified comparison expression
// to the state of the formatter.
func (v *formatter) formatComparison(comparison abs.ComparisonLike) {
	var first = comparison.GetFirst()
	v.formatExpression(first)
	v.state.AppendString(" ")
	var operator = comparison.GetOperator()
	v.formatOperator(operator)
	v.state.AppendString(" ")
	var second = comparison.GetSecond()
	v.formatExpression(second)
}

// This method attempts to parse a complement expression. It returns the
// complement expression and whether or not the complement expression was
// successfully parsed.
func (v *parser) parseComplement() (abs.ComplementLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var logical abs.ExpressionLike
	var expression abs.ComplementLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not a complement expression.
		return expression, token, false
	}
	if operator != abs.NOT {
		// This is not a complement expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	logical, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$complement")
		panic(message)
	}
	expression = exp.Complement(operator, logical)
	return expression, token, true
}

// This method adds the canonical format for the specified complement expression
// to the state of the formatter.
func (v *formatter) formatComplement(complement abs.ComplementLike) {
	var operator = complement.GetOperator()
	v.formatOperator(operator)
	v.state.AppendString(" ")
	var expression = complement.GetExpression()
	v.formatExpression(expression)
}

// This method attempts to parse a dereference expression. It returns the
// dereference expression and whether or not the dereference expression was
// successfully parsed.
func (v *parser) parseDereference() (abs.DereferenceLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var reference abs.ExpressionLike
	var expression abs.DereferenceLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not a dereference expression.
		return expression, token, false
	}
	if operator != abs.AT {
		// This is not a dereference expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	reference, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$dereference")
		panic(message)
	}
	expression = exp.Dereference(operator, reference)
	return expression, token, true
}

// This method adds the canonical format for the specified dereference expression
// to the state of the formatter.
func (v *formatter) formatDereference(dereference abs.DereferenceLike) {
	var operator = dereference.GetOperator()
	v.formatOperator(operator)
	var expression = dereference.GetExpression()
	v.formatExpression(expression)
}

// This method attempts to parse a power expression. It returns the
// power expression and whether or not the power expression was
// successfully parsed.
func (v *parser) parseExponential(base abs.ExpressionLike) (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var exponent abs.ExpressionLike
	var expression abs.ExpressionLike = base
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not an exponential expression.
		return expression, token, false
	}
	if operator != abs.CARET {
		// This is not an exponential expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	exponent, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$exponential")
		panic(message)
	}
	expression = exp.Exponential(base, operator, exponent)
	return expression, token, true
}

// This method adds the canonical format for the specified exponential expression
// to the state of the formatter.
func (v *formatter) formatExponential(exponential abs.ExponentialLike) {
	var base = exponential.GetBase()
	v.formatExpression(base)
	var operator = exponential.GetOperator()
	v.formatOperator(operator)
	var exponent = exponential.GetExponent()
	v.formatExpression(exponent)
}

// This method attempts to parse an expression. It returns the expression and
// whether or not the expression was successfully parsed. The expressions are
// are checked in precedence order from highest to lowest precedence.
func (v *parser) parseExpression() (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var expression abs.ExpressionLike
	expression, token, ok = v.parseComponent()
	if !ok {
		expression, token, ok = v.parseIntrinsic()
	}
	if !ok {
		expression, token, ok = v.parsePrecedence()
	}
	if !ok {
		expression, token, ok = v.parseDereference()
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
	if !ok {
		expression, token, ok = v.parseVariable()
	}
	if ok {
		expression, token, ok = v.parseRecursive(expression)
	}
	return expression, token, ok
}

// This method adds the canonical format for the specified expression to the
// state of the formatter.
func (v *formatter) formatExpression(expression abs.ExpressionLike) {
	// NOTE: A type switch will only work if each case specifies a unique
	// interface. If two different interfaces define the same method signatures
	// they are indistinguishable as types. To get around this we have added as
	// necessary a unique "dummy" method to each interface to guarantee that it
	// is unique.
	switch value := expression.(type) {
	case abs.ComponentLike:
		v.formatComponent(value)
	case abs.IntrinsicLike:
		v.formatIntrinsic(value)
	case abs.VariableLike:
		v.formatVariable(value)
	case abs.PrecedenceLike:
		v.formatPrecedence(value)
	case abs.DereferenceLike:
		v.formatDereference(value)
	case abs.InvocationLike:
		v.formatInvocation(value)
	case abs.ItemLike:
		v.formatItem(value)
	case abs.ChainingLike:
		v.formatChaining(value)
	case abs.ExponentialLike:
		v.formatExponential(value)
	case abs.InversionLike:
		v.formatInversion(value)
	case abs.ArithmeticLike:
		v.formatArithmetic(value)
	case abs.MagnitudeLike:
		v.formatMagnitude(value)
	case abs.ComparisonLike:
		v.formatComparison(value)
	case abs.ComplementLike:
		v.formatComplement(value)
	case abs.LogicalLike:
		v.formatLogical(value)
	default:
		panic(fmt.Sprintf("An invalid expression (of type %T) was passed to the formatter: %v", value, value))
	}
}

// This method attempts to parse a function expression. It returns the
// function expression and whether or not the function expression was
// successfully parsed.
func (v *parser) parseIntrinsic() (abs.IntrinsicLike, *Token, bool) {
	var ok bool
	var token *Token
	var function string
	var arguments abs.ListLike[abs.ExpressionLike]
	var expression abs.IntrinsicLike
	function, token, ok = v.parseIdentifier()
	if !ok {
		// This is not a function expression.
		return expression, token, false
	}
	arguments, token, ok = v.parseArguments()
	if !ok {
		// This is not a function expression.
		v.backupOne() // Put back the identifier token.
		return expression, token, false
	}
	expression = exp.Intrinsic(function, arguments)
	return expression, token, true
}

// This method adds the canonical format for the specified intrinsic function
// expression to the state of the formatter.
func (v *formatter) formatIntrinsic(dereference abs.IntrinsicLike) {
	var function = dereference.GetFunction()
	v.state.AppendString(function)
	var arguments = dereference.GetArguments()
	v.formatArguments(arguments)
}

// This method attempts to parse a inversion expression. It returns the
// inversion expression and whether or not the inversion expression was
// successfully parsed.
func (v *parser) parseInversion() (abs.InversionLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var numeric abs.ExpressionLike
	var expression abs.InversionLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not an inversion expression.
		return expression, token, false
	}
	if operator < abs.PLUS || operator > abs.STAR {
		// This is not an inversion expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	numeric, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$inversion")
		panic(message)
	}
	expression = exp.Inversion(operator, numeric)
	return expression, token, true
}

// This method adds the canonical format for the specified inversion expression
// to the state of the formatter.
func (v *formatter) formatInversion(inversion abs.InversionLike) {
	var operator = inversion.GetOperator()
	v.formatOperator(operator)
	var expression = inversion.GetExpression()
	v.formatExpression(expression)
}

// This method attempts to parse a message expression. It returns the
// message expression and whether or not the message expression was
// successfully parsed.
func (v *parser) parseInvocation(target abs.ExpressionLike) (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var message string
	var arguments abs.ListLike[abs.ExpressionLike]
	var expression abs.ExpressionLike = target
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not an invocation expression.
		return expression, token, false
	}
	if operator < abs.DOT || operator > abs.ARROW {
		// This is not an invocation expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	message, token, ok = v.parseIdentifier()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$method",
			"$invocation",
			"$method",
			"$arguments")
		panic(message)
	}
	arguments, token, ok = v.parseArguments()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$invocation",
			"$method",
			"$arguments")
		panic(message)
	}
	expression = exp.Invocation(target, operator, message, arguments)
	return expression, token, true
}

// This method adds the canonical format for the specified invocation expression
// to the state of the formatter.
func (v *formatter) formatInvocation(invocation abs.InvocationLike) {
	var target = invocation.GetTarget()
	v.formatExpression(target)
	var operator = invocation.GetOperator()
	v.formatOperator(operator)
	var message = invocation.GetMessage()
	v.state.AppendString(message)
	var arguments = invocation.GetArguments()
	v.formatArguments(arguments)
}

// This method attempts to parse an item expression. It returns the
// item expression and whether or not the item expression was successfully
// parsed.
func (v *parser) parseItem(composite abs.ExpressionLike) (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var indices abs.ListLike[abs.ExpressionLike]
	var expression abs.ExpressionLike = composite
	indices, token, ok = v.parseIndices()
	if !ok {
		return expression, token, false
	}
	expression = exp.Item(composite, indices)
	return expression, token, true
}

// This method adds the canonical format for the specified item expression
// to the state of the formatter.
func (v *formatter) formatItem(item abs.ItemLike) {
	var composite = item.GetComposite()
	v.formatExpression(composite)
	var indices = item.GetIndices()
	v.formatIndices(indices)
}

// This method attempts to parse a logical expression. It returns the
// logical expression and whether or not the logical expression was
// successfully parsed.
func (v *parser) parseLogical(first abs.ExpressionLike) (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.ExpressionLike
	var expression abs.ExpressionLike = first
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not a logical expression.
		return expression, token, false
	}
	if operator < abs.NOT || operator > abs.XOR {
		// This is not a logical expression.
		v.backupOne() // Put back the operator token.
		return expression, token, false
	}
	second, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$logical")
		panic(message)
	}
	expression = exp.Logical(first, operator, second)
	return expression, token, true
}

// This method adds the canonical format for the specified logical expression
// to the state of the formatter.
func (v *formatter) formatLogical(logical abs.LogicalLike) {
	var first = logical.GetFirst()
	v.formatExpression(first)
	v.state.AppendString(" ")
	var operator = logical.GetOperator()
	v.formatOperator(operator)
	v.state.AppendString(" ")
	var second = logical.GetSecond()
	v.formatExpression(second)
}

// This method attempts to parse a magnitude expression. It returns the
// magnitude expression and whether or not the magnitude expression was
// successfully parsed.
func (v *parser) parseMagnitude() (abs.MagnitudeLike, *Token, bool) {
	var ok bool
	var token *Token
	var numeric abs.ExpressionLike
	var expression abs.MagnitudeLike
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		// This is not an magnitude expression.
		return expression, token, false
	}
	numeric, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$magnitude")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("|",
			"$magnitude")
		panic(message)
	}
	expression = exp.Magnitude(numeric)
	return expression, token, true
}

// This method adds the canonical format for the specified magnitude expression
// to the state of the formatter.
func (v *formatter) formatMagnitude(magnitude abs.MagnitudeLike) {
	v.state.AppendString("|")
	var expression = magnitude.GetExpression()
	v.formatExpression(expression)
	v.state.AppendString("|")
}

// This method attempts to parse a precedence expression. It returns the
// precedence expression and whether or not the precedence expression was
// successfully parsed.
func (v *parser) parsePrecedence() (abs.PrecedenceLike, *Token, bool) {
	var ok bool
	var token *Token
	var inner abs.ExpressionLike
	var expression abs.PrecedenceLike
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		// This is not an precedence expression.
		return expression, token, false
	}
	inner, token, ok = v.parseExpression()
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$precedence")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = v.formatError(token)
		message += generateGrammar("$expression",
			"$precedence")
		panic(message)
	}
	expression = exp.Precedence(inner)
	return expression, token, true
}

// This method adds the canonical format for the specified precedence expression
// to the state of the formatter.
func (v *formatter) formatPrecedence(precedence abs.PrecedenceLike) {
	v.state.AppendString("(")
	var expression = precedence.GetExpression()
	v.formatExpression(expression)
	v.state.AppendString(")")
}

// This method attempts to parse a left recursive expression. It returns
// the left recursive expression and whether or not the left recursive
// expression was successfully parsed.
func (v *parser) parseRecursive(expression abs.ExpressionLike) (abs.ExpressionLike, *Token, bool) {
	var ok bool
	var token *Token
	expression, token, ok = v.parseInvocation(expression)
	if !ok {
		expression, token, ok = v.parseItem(expression)
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
	if ok {
		expression, token, ok = v.parseRecursive(expression)
	}
	return expression, token, true
}

// This method attempts to parse the an operator. It returns the operator and
// whether or not the operator was successfully parsed.
func (v *parser) parseOperator() (abs.Operator, *Token, bool) {
	var token = v.nextToken()
	var operator abs.Operator
	if token.Type == TokenKeyword {
		switch token.Value {
		case "AND":
			operator = abs.AND
		case "IS":
			operator = abs.IS
		case "MATCHES":
			operator = abs.MATCHES
		case "NOT":
			operator = abs.NOT
		case "OR":
			operator = abs.OR
		case "SANS":
			operator = abs.SANS
		case "XOR":
			operator = abs.XOR
		default:
			// The token is not an operator.
			v.backupOne()
			return operator, token, false
		}
		return operator, token, true
	}
	if token.Type == TokenDelimiter {
		switch token.Value {
		case "&":
			operator = abs.AMPERSAND
		case "<-":
			operator = abs.ARROW
		case ":=":
			operator = abs.ASSIGN
		case "@":
			operator = abs.AT
		case "|":
			operator = abs.BAR
		case "^":
			operator = abs.CARET
		case "?=":
			operator = abs.DEFAULT
		case "-=":
			operator = abs.DIFFERENCE
		case ".":
			operator = abs.DOT
		case "=":
			operator = abs.EQUAL
		case "<":
			operator = abs.LESS
		case "-":
			operator = abs.MINUS
		case "//":
			operator = abs.MODULO
		case ">":
			operator = abs.MORE
		case "+":
			operator = abs.PLUS
		case "*=":
			operator = abs.PRODUCT
		case "/=":
			operator = abs.QUOTIENT
		case "/":
			operator = abs.SLASH
		case "*":
			operator = abs.STAR
		case "+=":
			operator = abs.SUM
		case "~":
			operator = abs.TILDA
		case "≠":
			operator = abs.UNEQUAL
		default:
			// The token is not an operator.
			v.backupOne()
			return operator, token, false
		}
		return operator, token, true
	}
	// The token is not an operator.
	v.backupOne()
	return operator, token, false
}

// This method adds the canonical format for the specified operator to the
// state of the formatter.
func (v *formatter) formatOperator(operator abs.Operator) {
	switch operator {
	case abs.AMPERSAND:
		v.state.AppendString("&")
	case abs.AND:
		v.state.AppendString("AND")
	case abs.ARROW:
		v.state.AppendString("<-")
	case abs.ASSIGN:
		v.state.AppendString(":=")
	case abs.AT:
		v.state.AppendString("@")
	case abs.BAR:
		v.state.AppendString("|")
	case abs.CARET:
		v.state.AppendString("^")
	case abs.DEFAULT:
		v.state.AppendString("?=")
	case abs.DIFFERENCE:
		v.state.AppendString("-=")
	case abs.DOT:
		v.state.AppendString(".")
	case abs.EQUAL:
		v.state.AppendString("=")
	case abs.IS:
		v.state.AppendString("IS")
	case abs.LESS:
		v.state.AppendString("<")
	case abs.MATCHES:
		v.state.AppendString("MATCHES")
	case abs.MINUS:
		v.state.AppendString("-")
	case abs.MODULO:
		v.state.AppendString("//")
	case abs.MORE:
		v.state.AppendString(">")
	case abs.NOT:
		v.state.AppendString("NOT")
	case abs.OR:
		v.state.AppendString("OR")
	case abs.PLUS:
		v.state.AppendString("+")
	case abs.PRODUCT:
		v.state.AppendString("*=")
	case abs.QUOTIENT:
		v.state.AppendString("/=")
	case abs.SANS:
		v.state.AppendString("SANS")
	case abs.SLASH:
		v.state.AppendString("/")
	case abs.STAR:
		v.state.AppendString("*")
	case abs.SUM:
		v.state.AppendString("+=")
	case abs.TILDA:
		v.state.AppendString("~")
	case abs.UNEQUAL:
		v.state.AppendString("≠")
	case abs.XOR:
		v.state.AppendString("XOR")
	default:
		var message = fmt.Sprintf("An unexpected operator token was received by the formatter: %v\n", operator)
		panic(message)
	}
}

// This method attempts to parse an identifier. It returns the identifier
// string and whether or not the identifier was successfully parsed.
func (v *parser) parseVariable() (abs.VariableLike, *Token, bool) {
	var variable abs.VariableLike
	var token = v.nextToken()
	if token.Type != TokenIdentifier {
		v.backupOne()
		return variable, token, false
	}
	variable = exp.Variable(token.Value)
	return variable, token, true
}

// This method adds the canonical format for the specified identifier to the
// state of the formatter.
func (v *formatter) formatVariable(variable abs.VariableLike) {
	var identifier = variable.GetIdentifier()
	v.state.AppendString(identifier)
}
