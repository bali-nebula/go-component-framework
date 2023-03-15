/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	exp "github.com/bali-nebula/go-component-framework/v2/expressions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// This method attempts to parse the arguments within a call. It returns a
// list of the arguments and whether or not the arguments were successfully
// parsed.
func (v *parser) parseArguments() (abs.Sequential[abs.Expression], *Token, bool) {
	var ok bool
	var token *Token
	var argument abs.Expression
	var arguments = col.List[abs.Expression]()
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
func (v *formatter) formatArguments(arguments abs.Sequential[abs.Expression]) {
	v.AppendString("(")
	var iterator = col.Iterator[abs.Expression](arguments)
	if iterator.HasNext() {
		var argument = iterator.GetNext()
		v.formatExpression(argument)
	}
	for iterator.HasNext() {
		v.AppendString(", ")
		var argument = iterator.GetNext()
		v.formatExpression(argument)
	}
	v.AppendString(")")
}

// This method attempts to parse a arithmetic expression. It returns the
// arithmetic expression and whether or not the arithmetic expression was
// successfully parsed.
func (v *parser) parseArithmetic(first abs.Expression) (abs.Expression, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.Expression
	var expression abs.Expression = first
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
	v.AppendString(" ")
	var operator = arithmetic.GetOperator()
	v.formatOperator(operator)
	v.AppendString(" ")
	var second = arithmetic.GetSecond()
	v.formatExpression(second)
}

// This method attempts to parse a chain expression. It returns the
// chain expression and whether or not the chain expression was
// successfully parsed.
func (v *parser) parseChaining(first abs.Expression) (abs.Expression, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.Expression
	var expression abs.Expression = first
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
	v.AppendString(" ")
	var operator = chaining.GetOperator()
	v.formatOperator(operator)
	v.AppendString(" ")
	var second = chaining.GetSecond()
	v.formatExpression(second)
}

// This method attempts to parse a comparison expression. It returns the
// comparison expression and whether or not the comparison expression was
// successfully parsed.
func (v *parser) parseComparison(first abs.Expression) (abs.Expression, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.Expression
	var expression abs.Expression = first
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
	v.AppendString(" ")
	var operator = comparison.GetOperator()
	v.formatOperator(operator)
	v.AppendString(" ")
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
	var logical abs.Expression
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
	v.AppendString(" ")
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
	var reference abs.Expression
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
func (v *parser) parseExponential(base abs.Expression) (abs.Expression, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var exponent abs.Expression
	var expression abs.Expression = base
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
	v.AppendString(" ")
	var operator = exponential.GetOperator()
	v.formatOperator(operator)
	v.AppendString(" ")
	var exponent = exponential.GetExponent()
	v.formatExpression(exponent)
}

// This method attempts to parse an expression. It returns the expression and
// whether or not the expression was successfully parsed. The expressions are
// are checked in precedence order from highest to lowest precedence.
func (v *parser) parseExpression() (abs.Expression, *Token, bool) {
	var ok bool
	var token *Token
	var expression abs.Expression
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
		expression, token = v.parseRecursive(expression)
	}
	return expression, token, ok
}

// This method adds the canonical format for the specified expression to the
// state of the formatter.
func (v *formatter) formatExpression(expression abs.Expression) {
	switch exp.GetType(expression) {
	case "ComponentExpression":
		var value = expression.(abs.ComponentLike)
		v.formatComponent(value)
	case "IntrinsicExpression":
		var value = expression.(abs.IntrinsicLike)
		v.formatIntrinsic(value)
	case "VariableExpression":
		var value = expression.(abs.VariableLike)
		v.formatVariable(value)
	case "PrecedenceExpression":
		var value = expression.(abs.PrecedenceLike)
		v.formatPrecedence(value)
	case "DereferenceExpression":
		var value = expression.(abs.DereferenceLike)
		v.formatDereference(value)
	case "InvocationExpression":
		var value = expression.(abs.InvocationLike)
		v.formatInvocation(value)
	case "SubcomponentExpression":
		var value = expression.(abs.SubcomponentLike)
		v.formatSubcomponent(value)
	case "ChainingExpression":
		var value = expression.(abs.ChainingLike)
		v.formatChaining(value)
	case "ExponentialExpression":
		var value = expression.(abs.ExponentialLike)
		v.formatExponential(value)
	case "InversionExpression":
		var value = expression.(abs.InversionLike)
		v.formatInversion(value)
	case "ArithmeticExpression":
		var value = expression.(abs.ArithmeticLike)
		v.formatArithmetic(value)
	case "MagnitudeExpression":
		var value = expression.(abs.MagnitudeLike)
		v.formatMagnitude(value)
	case "ComparisonExpression":
		var value = expression.(abs.ComparisonLike)
		v.formatComparison(value)
	case "ComplementExpression":
		var value = expression.(abs.ComplementLike)
		v.formatComplement(value)
	case "LogicalExpression":
		var value = expression.(abs.LogicalLike)
		v.formatLogical(value)
	default:
		var message = fmt.Sprintf("An invalid expression type was passed to the formatter: %v", expression)
		panic(message)
	}
}

// This method attempts to parse a function expression. It returns the
// function expression and whether or not the function expression was
// successfully parsed.
func (v *parser) parseIntrinsic() (abs.IntrinsicLike, *Token, bool) {
	var ok bool
	var token *Token
	var function string
	var arguments abs.Sequential[abs.Expression]
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
	v.AppendString(function)
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
	var numeric abs.Expression
	var expression abs.InversionLike
	operator, token, ok = v.parseOperator()
	if !ok {
		// This is not an inversion expression.
		return expression, token, false
	}
	if operator < abs.MINUS || operator > abs.SLASH {
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
func (v *parser) parseInvocation(target abs.Expression) (abs.Expression, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var message string
	var arguments abs.Sequential[abs.Expression]
	var expression abs.Expression = target
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
	v.AppendString(message)
	var arguments = invocation.GetArguments()
	v.formatArguments(arguments)
}

// This method attempts to parse an subcomponent expression. It returns the
// subcomponent expression and whether or not the subcomponent expression was
// successfully parsed.
func (v *parser) parseSubcomponent(composite abs.Expression) (abs.Expression, *Token, bool) {
	var ok bool
	var token *Token
	var indices abs.Sequential[abs.Expression]
	var expression abs.Expression = composite
	indices, token, ok = v.parseIndices()
	if !ok {
		return expression, token, false
	}
	expression = exp.Subcomponent(composite, indices)
	return expression, token, true
}

// This method adds the canonical format for the specified subcomponent
// expression to the state of the formatter.
func (v *formatter) formatSubcomponent(subcomponent abs.SubcomponentLike) {
	var composite = subcomponent.GetComposite()
	v.formatExpression(composite)
	var indices = subcomponent.GetIndices()
	v.formatIndices(indices)
}

// This method attempts to parse a logical expression. It returns the
// logical expression and whether or not the logical expression was
// successfully parsed.
func (v *parser) parseLogical(first abs.Expression) (abs.Expression, *Token, bool) {
	var ok bool
	var token *Token
	var operator abs.Operator
	var second abs.Expression
	var expression abs.Expression = first
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
	v.AppendString(" ")
	var operator = logical.GetOperator()
	v.formatOperator(operator)
	v.AppendString(" ")
	var second = logical.GetSecond()
	v.formatExpression(second)
}

// This method attempts to parse a magnitude expression. It returns the
// magnitude expression and whether or not the magnitude expression was
// successfully parsed.
func (v *parser) parseMagnitude() (abs.MagnitudeLike, *Token, bool) {
	var ok bool
	var token *Token
	var numeric abs.Expression
	var expression abs.MagnitudeLike
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		// This is not a magnitude expression.
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
	v.AppendString("|")
	var expression = magnitude.GetExpression()
	v.formatExpression(expression)
	v.AppendString("|")
}

// This method attempts to parse a precedence expression. It returns the
// precedence expression and whether or not the precedence expression was
// successfully parsed.
func (v *parser) parsePrecedence() (abs.PrecedenceLike, *Token, bool) {
	var ok bool
	var token *Token
	var inner abs.Expression
	var expression abs.PrecedenceLike
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		// This is not a precedence expression.
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
	v.AppendString("(")
	var expression = precedence.GetExpression()
	v.formatExpression(expression)
	v.AppendString(")")
}

// This method attempts to parse a recursive expression using the specified left
// expression. It returns the recursive expression.
func (v *parser) parseRecursive(left abs.Expression) (abs.Expression, *Token) {
	var ok bool
	var token *Token
	var expression abs.Expression
	expression, token, ok = v.parseInvocation(left)
	if !ok {
		expression, token, ok = v.parseSubcomponent(left)
	}
	if !ok {
		expression, token, ok = v.parseChaining(left)
	}
	if !ok {
		expression, token, ok = v.parseExponential(left)
	}
	if !ok {
		expression, token, ok = v.parseArithmetic(left)
	}
	if !ok {
		expression, token, ok = v.parseComparison(left)
	}
	if !ok {
		expression, token, ok = v.parseLogical(left)
	}
	if ok {
		expression, token = v.parseRecursive(expression)
	}
	return expression, token
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
		v.AppendString("&")
	case abs.AND:
		v.AppendString("AND")
	case abs.ARROW:
		v.AppendString("<-")
	case abs.ASSIGN:
		v.AppendString(":=")
	case abs.AT:
		v.AppendString("@")
	case abs.BAR:
		v.AppendString("|")
	case abs.CARET:
		v.AppendString("^")
	case abs.DEFAULT:
		v.AppendString("?=")
	case abs.DIFFERENCE:
		v.AppendString("-=")
	case abs.DOT:
		v.AppendString(".")
	case abs.EQUAL:
		v.AppendString("=")
	case abs.IS:
		v.AppendString("IS")
	case abs.LESS:
		v.AppendString("<")
	case abs.MATCHES:
		v.AppendString("MATCHES")
	case abs.MINUS:
		v.AppendString("-")
	case abs.MODULO:
		v.AppendString("//")
	case abs.MORE:
		v.AppendString(">")
	case abs.NOT:
		v.AppendString("NOT")
	case abs.OR:
		v.AppendString("OR")
	case abs.PLUS:
		v.AppendString("+")
	case abs.PRODUCT:
		v.AppendString("*=")
	case abs.QUOTIENT:
		v.AppendString("/=")
	case abs.SANS:
		v.AppendString("SANS")
	case abs.SLASH:
		v.AppendString("/")
	case abs.STAR:
		v.AppendString("*")
	case abs.SUM:
		v.AppendString("+=")
	case abs.TILDA:
		v.AppendString("~")
	case abs.UNEQUAL:
		v.AppendString("≠")
	case abs.XOR:
		v.AppendString("XOR")
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
	v.AppendString(identifier)
}
