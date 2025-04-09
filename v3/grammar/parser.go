/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	fmt "fmt"
	ast "github.com/bali-nebula/go-component-framework/v3/ast"
	fra "github.com/craterdog/go-collection-framework/v5"
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
	mat "math"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func ParserClass() ParserClassLike {
	return parserClass()
}

// Constructor Methods

func (c *parserClass_) Parser() ParserLike {
	var instance = &parser_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parser_) GetClass() ParserClassLike {
	return parserClass()
}

func (v *parser_) ParseSource(
	source string,
) ast.DocumentLike {
	v.source_ = sts.ReplaceAll(source, "\t", "    ")
	v.tokens_ = fra.Queue[TokenLike]()
	v.next_ = fra.Stack[TokenLike]()

	// The scanner runs in a separate Go routine.
	ScannerClass().Scanner(v.source_, v.tokens_)

	// Attempt to parse the document.
	var document, token, ok = v.parseDocument()
	if !ok || !v.tokens_.IsEmpty() {
		var message = v.formatError("$Document", token)
		panic(message)
	}
	return document
}

// PROTECTED INTERFACE

// Private Methods

func (v *parser_) parseAcceptClause() (
	acceptClause ast.AcceptClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "accept" delimiter.
	_, token, ok = v.parseDelimiter("accept")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AcceptClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AcceptClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Message rule.
	var message ast.MessageLike
	message, token, ok = v.parseMessage()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AcceptClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AcceptClause", token)
		panic(message)
	}

	// Found a single AcceptClause rule.
	ok = true
	v.remove(tokens)
	acceptClause = ast.AcceptClauseClass().AcceptClause(message)
	return
}

func (v *parser_) parseAdditionalArgument() (
	additionalArgument ast.AdditionalArgumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalArgument rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalArgument", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Argument rule.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AdditionalArgument rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalArgument", token)
		panic(message)
	}

	// Found a single AdditionalArgument rule.
	ok = true
	v.remove(tokens)
	additionalArgument = ast.AdditionalArgumentClass().AdditionalArgument(argument)
	return
}

func (v *parser_) parseAdditionalAssociation() (
	additionalAssociation ast.AdditionalAssociationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalAssociation rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalAssociation", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Association rule.
	var association ast.AssociationLike
	association, token, ok = v.parseAssociation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AdditionalAssociation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalAssociation", token)
		panic(message)
	}

	// Found a single AdditionalAssociation rule.
	ok = true
	v.remove(tokens)
	additionalAssociation = ast.AdditionalAssociationClass().AdditionalAssociation(association)
	return
}

func (v *parser_) parseAdditionalIndex() (
	additionalIndex ast.AdditionalIndexLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalIndex rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalIndex", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Index rule.
	var index ast.IndexLike
	index, token, ok = v.parseIndex()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AdditionalIndex rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalIndex", token)
		panic(message)
	}

	// Found a single AdditionalIndex rule.
	ok = true
	v.remove(tokens)
	additionalIndex = ast.AdditionalIndexClass().AdditionalIndex(index)
	return
}

func (v *parser_) parseAdditionalParameter() (
	additionalParameter ast.AdditionalParameterLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalParameter rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalParameter", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Parameter rule.
	var parameter ast.ParameterLike
	parameter, token, ok = v.parseParameter()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AdditionalParameter rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalParameter", token)
		panic(message)
	}

	// Found a single AdditionalParameter rule.
	ok = true
	v.remove(tokens)
	additionalParameter = ast.AdditionalParameterClass().AdditionalParameter(parameter)
	return
}

func (v *parser_) parseAdditionalStatement() (
	additionalStatement ast.AdditionalStatementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ";" delimiter.
	_, token, ok = v.parseDelimiter(";")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalStatement rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalStatement", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Statement rule.
	var statement ast.StatementLike
	statement, token, ok = v.parseStatement()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AdditionalStatement rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalStatement", token)
		panic(message)
	}

	// Found a single AdditionalStatement rule.
	ok = true
	v.remove(tokens)
	additionalStatement = ast.AdditionalStatementClass().AdditionalStatement(statement)
	return
}

func (v *parser_) parseAdditionalValue() (
	additionalValue ast.AdditionalValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "," delimiter.
	_, token, ok = v.parseDelimiter(",")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AdditionalValue rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AdditionalValue", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Value rule.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AdditionalValue rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AdditionalValue", token)
		panic(message)
	}

	// Found a single AdditionalValue rule.
	ok = true
	v.remove(tokens)
	additionalValue = ast.AdditionalValueClass().AdditionalValue(value)
	return
}

func (v *parser_) parseAmplitude() (
	amplitude ast.AmplitudeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "|" delimiter.
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Amplitude rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Amplitude", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Amplitude rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Amplitude", token)
		panic(message)
	}

	// Attempt to parse a single "|" delimiter.
	_, token, ok = v.parseDelimiter("|")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Amplitude rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Amplitude", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Amplitude rule.
	ok = true
	v.remove(tokens)
	amplitude = ast.AmplitudeClass().Amplitude(expression)
	return
}

func (v *parser_) parseAnd() (
	and ast.AndLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "AND" delimiter.
	_, token, ok = v.parseDelimiter("AND")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single And rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$And", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single And rule.
	ok = true
	v.remove(tokens)
	and = ast.AndClass().And()
	return
}

func (v *parser_) parseAnnotatedAssociation() (
	annotatedAssociation ast.AnnotatedAssociationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Association rule.
	var association ast.AssociationLike
	association, token, ok = v.parseAssociation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AnnotatedAssociation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AnnotatedAssociation", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single AnnotatedAssociation rule.
	ok = true
	v.remove(tokens)
	annotatedAssociation = ast.AnnotatedAssociationClass().AnnotatedAssociation(
		association,
		optionalNote,
	)
	return
}

func (v *parser_) parseAnnotatedParameter() (
	annotatedParameter ast.AnnotatedParameterLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Parameter rule.
	var parameter ast.ParameterLike
	parameter, token, ok = v.parseParameter()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AnnotatedParameter rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AnnotatedParameter", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single AnnotatedParameter rule.
	ok = true
	v.remove(tokens)
	annotatedParameter = ast.AnnotatedParameterClass().AnnotatedParameter(
		parameter,
		optionalNote,
	)
	return
}

func (v *parser_) parseAnnotatedStatement() (
	annotatedStatement ast.AnnotatedStatementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Statement rule.
	var statement ast.StatementLike
	statement, token, ok = v.parseStatement()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AnnotatedStatement rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AnnotatedStatement", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single AnnotatedStatement rule.
	ok = true
	v.remove(tokens)
	annotatedStatement = ast.AnnotatedStatementClass().AnnotatedStatement(
		statement,
		optionalNote,
	)
	return
}

func (v *parser_) parseAnnotatedValue() (
	annotatedValue ast.AnnotatedValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Value rule.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AnnotatedValue rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AnnotatedValue", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single AnnotatedValue rule.
	ok = true
	v.remove(tokens)
	annotatedValue = ast.AnnotatedValueClass().AnnotatedValue(
		value,
		optionalNote,
	)
	return
}

func (v *parser_) parseArgument() (
	argument ast.ArgumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Argument rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Argument", token)
		panic(message)
	}

	// Found a single Argument rule.
	ok = true
	v.remove(tokens)
	argument = ast.ArgumentClass().Argument(expression)
	return
}

func (v *parser_) parseArguments() (
	arguments ast.ArgumentsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Argument rule.
	var argument ast.ArgumentLike
	argument, token, ok = v.parseArgument()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Arguments rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Arguments", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalArgument rules.
	var additionalArguments = fra.List[ast.AdditionalArgumentLike]()
additionalArgumentsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalArgument ast.AdditionalArgumentLike
		additionalArgument, token, ok = v.parseAdditionalArgument()
		if !ok {
			switch {
			case count >= 0:
				break additionalArgumentsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalArgument rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Arguments", token)
				message += "0 or more AdditionalArgument rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalArguments.AppendValue(additionalArgument)
	}

	// Found a single Arguments rule.
	ok = true
	v.remove(tokens)
	arguments = ast.ArgumentsClass().Arguments(
		argument,
		additionalArguments,
	)
	return
}

func (v *parser_) parseArithmetic() (
	arithmetic ast.ArithmeticLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression1 ast.ExpressionLike
	expression1, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Arithmetic rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Arithmetic", token)
		panic(message)
	}

	// Attempt to parse a single ArithmeticOperator rule.
	var arithmeticOperator ast.ArithmeticOperatorLike
	arithmeticOperator, token, ok = v.parseArithmeticOperator()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Arithmetic rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Arithmetic", token)
		panic(message)
	}

	// Attempt to parse a single Expression rule.
	var expression2 ast.ExpressionLike
	expression2, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Arithmetic rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Arithmetic", token)
		panic(message)
	}

	// Found a single Arithmetic rule.
	ok = true
	v.remove(tokens)
	arithmetic = ast.ArithmeticClass().Arithmetic(
		expression1,
		arithmeticOperator,
		expression2,
	)
	return
}

func (v *parser_) parseArithmeticOperator() (
	arithmeticOperator ast.ArithmeticOperatorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Times ArithmeticOperator.
	var times ast.TimesLike
	times, token, ok = v.parseTimes()
	if ok {
		// Found a single Times ArithmeticOperator.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(times)
		return
	}

	// Attempt to parse a single Divide ArithmeticOperator.
	var divide ast.DivideLike
	divide, token, ok = v.parseDivide()
	if ok {
		// Found a single Divide ArithmeticOperator.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(divide)
		return
	}

	// Attempt to parse a single Modulo ArithmeticOperator.
	var modulo ast.ModuloLike
	modulo, token, ok = v.parseModulo()
	if ok {
		// Found a single Modulo ArithmeticOperator.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(modulo)
		return
	}

	// Attempt to parse a single Plus ArithmeticOperator.
	var plus ast.PlusLike
	plus, token, ok = v.parsePlus()
	if ok {
		// Found a single Plus ArithmeticOperator.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(plus)
		return
	}

	// Attempt to parse a single Minus ArithmeticOperator.
	var minus ast.MinusLike
	minus, token, ok = v.parseMinus()
	if ok {
		// Found a single Minus ArithmeticOperator.
		arithmeticOperator = ast.ArithmeticOperatorClass().ArithmeticOperator(minus)
		return
	}

	// This is not a single ArithmeticOperator rule.
	return
}

func (v *parser_) parseArrow() (
	arrow ast.ArrowLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "<-" delimiter.
	_, token, ok = v.parseDelimiter("<-")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Arrow rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Arrow", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Arrow rule.
	ok = true
	v.remove(tokens)
	arrow = ast.ArrowClass().Arrow()
	return
}

func (v *parser_) parseAssignmentOperator() (
	assignmentOperator ast.AssignmentOperatorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single ColonEquals AssignmentOperator.
	var colonEquals ast.ColonEqualsLike
	colonEquals, token, ok = v.parseColonEquals()
	if ok {
		// Found a single ColonEquals AssignmentOperator.
		assignmentOperator = ast.AssignmentOperatorClass().AssignmentOperator(colonEquals)
		return
	}

	// Attempt to parse a single DefaultEquals AssignmentOperator.
	var defaultEquals ast.DefaultEqualsLike
	defaultEquals, token, ok = v.parseDefaultEquals()
	if ok {
		// Found a single DefaultEquals AssignmentOperator.
		assignmentOperator = ast.AssignmentOperatorClass().AssignmentOperator(defaultEquals)
		return
	}

	// Attempt to parse a single PlusEquals AssignmentOperator.
	var plusEquals ast.PlusEqualsLike
	plusEquals, token, ok = v.parsePlusEquals()
	if ok {
		// Found a single PlusEquals AssignmentOperator.
		assignmentOperator = ast.AssignmentOperatorClass().AssignmentOperator(plusEquals)
		return
	}

	// Attempt to parse a single MinusEquals AssignmentOperator.
	var minusEquals ast.MinusEqualsLike
	minusEquals, token, ok = v.parseMinusEquals()
	if ok {
		// Found a single MinusEquals AssignmentOperator.
		assignmentOperator = ast.AssignmentOperatorClass().AssignmentOperator(minusEquals)
		return
	}

	// Attempt to parse a single TimesEquals AssignmentOperator.
	var timesEquals ast.TimesEqualsLike
	timesEquals, token, ok = v.parseTimesEquals()
	if ok {
		// Found a single TimesEquals AssignmentOperator.
		assignmentOperator = ast.AssignmentOperatorClass().AssignmentOperator(timesEquals)
		return
	}

	// Attempt to parse a single DivideEquals AssignmentOperator.
	var divideEquals ast.DivideEqualsLike
	divideEquals, token, ok = v.parseDivideEquals()
	if ok {
		// Found a single DivideEquals AssignmentOperator.
		assignmentOperator = ast.AssignmentOperatorClass().AssignmentOperator(divideEquals)
		return
	}

	// This is not a single AssignmentOperator rule.
	return
}

func (v *parser_) parseAssociation() (
	association ast.AssociationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Key rule.
	var key ast.KeyLike
	key, token, ok = v.parseKey()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Association rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Association", token)
		panic(message)
	}

	// Attempt to parse a single ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Association rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Association", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Value rule.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Association rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Association", token)
		panic(message)
	}

	// Found a single Association rule.
	ok = true
	v.remove(tokens)
	association = ast.AssociationClass().Association(
		key,
		value,
	)
	return
}

func (v *parser_) parseAssociations() (
	associations ast.AssociationsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single InlineAssociations Associations.
	var inlineAssociations ast.InlineAssociationsLike
	inlineAssociations, token, ok = v.parseInlineAssociations()
	if ok {
		// Found a single InlineAssociations Associations.
		associations = ast.AssociationsClass().Associations(inlineAssociations)
		return
	}

	// Attempt to parse a single MultilineAssociations Associations.
	var multilineAssociations ast.MultilineAssociationsLike
	multilineAssociations, token, ok = v.parseMultilineAssociations()
	if ok {
		// Found a single MultilineAssociations Associations.
		associations = ast.AssociationsClass().Associations(multilineAssociations)
		return
	}

	// Attempt to parse a single NoAssociations Associations.
	var noAssociations ast.NoAssociationsLike
	noAssociations, token, ok = v.parseNoAssociations()
	if ok {
		// Found a single NoAssociations Associations.
		associations = ast.AssociationsClass().Associations(noAssociations)
		return
	}

	// This is not a single Associations rule.
	return
}

func (v *parser_) parseAtLevel() (
	atLevel ast.AtLevelLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "at" delimiter.
	_, token, ok = v.parseDelimiter("at")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AtLevel rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AtLevel", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "level" delimiter.
	_, token, ok = v.parseDelimiter("level")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AtLevel rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AtLevel", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single AtLevel rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$AtLevel", token)
		panic(message)
	}

	// Found a single AtLevel rule.
	ok = true
	v.remove(tokens)
	atLevel = ast.AtLevelClass().AtLevel(expression)
	return
}

func (v *parser_) parseAttribute() (
	attribute ast.AttributeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Variable rule.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Attribute rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Attribute", token)
		panic(message)
	}

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Attribute rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Attribute", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Indices rule.
	var indices ast.IndicesLike
	indices, token, ok = v.parseIndices()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Attribute rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Attribute", token)
		panic(message)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Attribute rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Attribute", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Attribute rule.
	ok = true
	v.remove(tokens)
	attribute = ast.AttributeClass().Attribute(
		variable,
		indices,
	)
	return
}

func (v *parser_) parseBag() (
	bag ast.BagLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Bag rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Bag", token)
		panic(message)
	}

	// Found a single Bag rule.
	ok = true
	v.remove(tokens)
	bag = ast.BagClass().Bag(expression)
	return
}

func (v *parser_) parseBreakClause() (
	breakClause ast.BreakClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "break" delimiter.
	_, token, ok = v.parseDelimiter("break")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single BreakClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$BreakClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "loop" delimiter.
	_, token, ok = v.parseDelimiter("loop")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single BreakClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$BreakClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single BreakClause rule.
	ok = true
	v.remove(tokens)
	breakClause = ast.BreakClauseClass().BreakClause()
	return
}

func (v *parser_) parseChaining() (
	chaining ast.ChainingLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression1 ast.ExpressionLike
	expression1, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Chaining rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Chaining", token)
		panic(message)
	}

	// Attempt to parse a single "&" delimiter.
	_, token, ok = v.parseDelimiter("&")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Chaining rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Chaining", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression2 ast.ExpressionLike
	expression2, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Chaining rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Chaining", token)
		panic(message)
	}

	// Found a single Chaining rule.
	ok = true
	v.remove(tokens)
	chaining = ast.ChainingClass().Chaining(
		expression1,
		expression2,
	)
	return
}

func (v *parser_) parseCheckoutClause() (
	checkoutClause ast.CheckoutClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "checkout" delimiter.
	_, token, ok = v.parseDelimiter("checkout")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single CheckoutClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$CheckoutClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Recipient rule.
	var recipient ast.RecipientLike
	recipient, token, ok = v.parseRecipient()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single CheckoutClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$CheckoutClause", token)
		panic(message)
	}

	// Attempt to parse an optional AtLevel rule.
	var optionalAtLevel ast.AtLevelLike
	optionalAtLevel, _, ok = v.parseAtLevel()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single "from" delimiter.
	_, token, ok = v.parseDelimiter("from")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single CheckoutClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$CheckoutClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Citation rule.
	var citation ast.CitationLike
	citation, token, ok = v.parseCitation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single CheckoutClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$CheckoutClause", token)
		panic(message)
	}

	// Found a single CheckoutClause rule.
	ok = true
	v.remove(tokens)
	checkoutClause = ast.CheckoutClauseClass().CheckoutClause(
		recipient,
		optionalAtLevel,
		citation,
	)
	return
}

func (v *parser_) parseCitation() (
	citation ast.CitationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Citation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Citation", token)
		panic(message)
	}

	// Found a single Citation rule.
	ok = true
	v.remove(tokens)
	citation = ast.CitationClass().Citation(expression)
	return
}

func (v *parser_) parseCollection() (
	collection ast.CollectionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Collection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Collection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Items rule.
	var optionalItems ast.ItemsLike
	optionalItems, _, ok = v.parseItems()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Collection rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Collection", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Collection rule.
	ok = true
	v.remove(tokens)
	collection = ast.CollectionClass().Collection(optionalItems)
	return
}

func (v *parser_) parseColonEquals() (
	colonEquals ast.ColonEqualsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ":=" delimiter.
	_, token, ok = v.parseDelimiter(":=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ColonEquals rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ColonEquals", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ColonEquals rule.
	ok = true
	v.remove(tokens)
	colonEquals = ast.ColonEqualsClass().ColonEquals()
	return
}

func (v *parser_) parseComparison() (
	comparison ast.ComparisonLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression1 ast.ExpressionLike
	expression1, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Comparison rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Comparison", token)
		panic(message)
	}

	// Attempt to parse a single ComparisonOperator rule.
	var comparisonOperator ast.ComparisonOperatorLike
	comparisonOperator, token, ok = v.parseComparisonOperator()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Comparison rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Comparison", token)
		panic(message)
	}

	// Attempt to parse a single Expression rule.
	var expression2 ast.ExpressionLike
	expression2, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Comparison rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Comparison", token)
		panic(message)
	}

	// Found a single Comparison rule.
	ok = true
	v.remove(tokens)
	comparison = ast.ComparisonClass().Comparison(
		expression1,
		comparisonOperator,
		expression2,
	)
	return
}

func (v *parser_) parseComparisonOperator() (
	comparisonOperator ast.ComparisonOperatorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Less ComparisonOperator.
	var less ast.LessLike
	less, token, ok = v.parseLess()
	if ok {
		// Found a single Less ComparisonOperator.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(less)
		return
	}

	// Attempt to parse a single Equals ComparisonOperator.
	var equals ast.EqualsLike
	equals, token, ok = v.parseEquals()
	if ok {
		// Found a single Equals ComparisonOperator.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(equals)
		return
	}

	// Attempt to parse a single More ComparisonOperator.
	var more ast.MoreLike
	more, token, ok = v.parseMore()
	if ok {
		// Found a single More ComparisonOperator.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(more)
		return
	}

	// Attempt to parse a single Is ComparisonOperator.
	var is ast.IsLike
	is, token, ok = v.parseIs()
	if ok {
		// Found a single Is ComparisonOperator.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(is)
		return
	}

	// Attempt to parse a single Matches ComparisonOperator.
	var matches ast.MatchesLike
	matches, token, ok = v.parseMatches()
	if ok {
		// Found a single Matches ComparisonOperator.
		comparisonOperator = ast.ComparisonOperatorClass().ComparisonOperator(matches)
		return
	}

	// This is not a single ComparisonOperator rule.
	return
}

func (v *parser_) parseComplement() (
	complement ast.ComplementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "NOT" delimiter.
	_, token, ok = v.parseDelimiter("NOT")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Complement rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Complement", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Complement rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Complement", token)
		panic(message)
	}

	// Found a single Complement rule.
	ok = true
	v.remove(tokens)
	complement = ast.ComplementClass().Complement(expression)
	return
}

func (v *parser_) parseComponent() (
	component ast.ComponentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Entity rule.
	var entity ast.EntityLike
	entity, token, ok = v.parseEntity()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Component rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Component", token)
		panic(message)
	}

	// Attempt to parse an optional Context rule.
	var optionalContext ast.ContextLike
	optionalContext, _, ok = v.parseContext()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Component rule.
	ok = true
	v.remove(tokens)
	component = ast.ComponentClass().Component(
		entity,
		optionalContext,
	)
	return
}

func (v *parser_) parseComposite() (
	composite ast.CompositeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Composite rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Composite", token)
		panic(message)
	}

	// Found a single Composite rule.
	ok = true
	v.remove(tokens)
	composite = ast.CompositeClass().Composite(expression)
	return
}

func (v *parser_) parseCondition() (
	condition ast.ConditionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Condition rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Condition", token)
		panic(message)
	}

	// Found a single Condition rule.
	ok = true
	v.remove(tokens)
	condition = ast.ConditionClass().Condition(expression)
	return
}

func (v *parser_) parseContext() (
	context ast.ContextLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Context rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Context", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Parameters rule.
	var parameters ast.ParametersLike
	parameters, token, ok = v.parseParameters()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Context rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Context", token)
		panic(message)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Context rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Context", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Context rule.
	ok = true
	v.remove(tokens)
	context = ast.ContextClass().Context(parameters)
	return
}

func (v *parser_) parseContinueClause() (
	continueClause ast.ContinueClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "continue" delimiter.
	_, token, ok = v.parseDelimiter("continue")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ContinueClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ContinueClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "loop" delimiter.
	_, token, ok = v.parseDelimiter("loop")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ContinueClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ContinueClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ContinueClause rule.
	ok = true
	v.remove(tokens)
	continueClause = ast.ContinueClauseClass().ContinueClause()
	return
}

func (v *parser_) parseDefaultEquals() (
	defaultEquals ast.DefaultEqualsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "?=" delimiter.
	_, token, ok = v.parseDelimiter("?=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DefaultEquals rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DefaultEquals", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single DefaultEquals rule.
	ok = true
	v.remove(tokens)
	defaultEquals = ast.DefaultEqualsClass().DefaultEquals()
	return
}

func (v *parser_) parseDereference() (
	dereference ast.DereferenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "@" delimiter.
	_, token, ok = v.parseDelimiter("@")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Dereference rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Dereference", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Dereference rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Dereference", token)
		panic(message)
	}

	// Found a single Dereference rule.
	ok = true
	v.remove(tokens)
	dereference = ast.DereferenceClass().Dereference(expression)
	return
}

func (v *parser_) parseDiscardClause() (
	discardClause ast.DiscardClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "discard" delimiter.
	_, token, ok = v.parseDelimiter("discard")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DiscardClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DiscardClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Draft rule.
	var draft ast.DraftLike
	draft, token, ok = v.parseDraft()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single DiscardClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$DiscardClause", token)
		panic(message)
	}

	// Found a single DiscardClause rule.
	ok = true
	v.remove(tokens)
	discardClause = ast.DiscardClauseClass().DiscardClause(draft)
	return
}

func (v *parser_) parseDivide() (
	divide ast.DivideLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "/" delimiter.
	_, token, ok = v.parseDelimiter("/")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Divide rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Divide", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Divide rule.
	ok = true
	v.remove(tokens)
	divide = ast.DivideClass().Divide()
	return
}

func (v *parser_) parseDivideEquals() (
	divideEquals ast.DivideEqualsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "/=" delimiter.
	_, token, ok = v.parseDelimiter("/=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DivideEquals rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DivideEquals", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single DivideEquals rule.
	ok = true
	v.remove(tokens)
	divideEquals = ast.DivideEqualsClass().DivideEquals()
	return
}

func (v *parser_) parseDoClause() (
	doClause ast.DoClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single DoClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$DoClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Operation rule.
	var operation ast.OperationLike
	operation, token, ok = v.parseOperation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single DoClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$DoClause", token)
		panic(message)
	}

	// Found a single DoClause rule.
	ok = true
	v.remove(tokens)
	doClause = ast.DoClauseClass().DoClause(operation)
	return
}

func (v *parser_) parseDocument() (
	document ast.DocumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single comment token.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Document rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Document", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Document rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Document", token)
		panic(message)
	}

	// Found a single Document rule.
	ok = true
	v.remove(tokens)
	document = ast.DocumentClass().Document(
		comment,
		component,
	)
	return
}

func (v *parser_) parseDot() (
	dot ast.DotLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "." delimiter.
	_, token, ok = v.parseDelimiter(".")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Dot rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Dot", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Dot rule.
	ok = true
	v.remove(tokens)
	dot = ast.DotClass().Dot()
	return
}

func (v *parser_) parseDraft() (
	draft ast.DraftLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Draft rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Draft", token)
		panic(message)
	}

	// Found a single Draft rule.
	ok = true
	v.remove(tokens)
	draft = ast.DraftClass().Draft(expression)
	return
}

func (v *parser_) parseElement() (
	element ast.ElementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single angle Element.
	var angle string
	angle, token, ok = v.parseToken(AngleToken)
	if ok {
		// Found a single angle Element.
		element = ast.ElementClass().Element(angle)
		return
	}

	// Attempt to parse a single boolean Element.
	var boolean string
	boolean, token, ok = v.parseToken(BooleanToken)
	if ok {
		// Found a single boolean Element.
		element = ast.ElementClass().Element(boolean)
		return
	}

	// Attempt to parse a single duration Element.
	var duration string
	duration, token, ok = v.parseToken(DurationToken)
	if ok {
		// Found a single duration Element.
		element = ast.ElementClass().Element(duration)
		return
	}

	// Attempt to parse a single moment Element.
	var moment string
	moment, token, ok = v.parseToken(MomentToken)
	if ok {
		// Found a single moment Element.
		element = ast.ElementClass().Element(moment)
		return
	}

	// Attempt to parse a single number Element.
	var number string
	number, token, ok = v.parseToken(NumberToken)
	if ok {
		// Found a single number Element.
		element = ast.ElementClass().Element(number)
		return
	}

	// Attempt to parse a single pattern Element.
	var pattern string
	pattern, token, ok = v.parseToken(PatternToken)
	if ok {
		// Found a single pattern Element.
		element = ast.ElementClass().Element(pattern)
		return
	}

	// Attempt to parse a single percentage Element.
	var percentage string
	percentage, token, ok = v.parseToken(PercentageToken)
	if ok {
		// Found a single percentage Element.
		element = ast.ElementClass().Element(percentage)
		return
	}

	// Attempt to parse a single probability Element.
	var probability string
	probability, token, ok = v.parseToken(ProbabilityToken)
	if ok {
		// Found a single probability Element.
		element = ast.ElementClass().Element(probability)
		return
	}

	// Attempt to parse a single resource Element.
	var resource string
	resource, token, ok = v.parseToken(ResourceToken)
	if ok {
		// Found a single resource Element.
		element = ast.ElementClass().Element(resource)
		return
	}

	// This is not a single Element rule.
	return
}

func (v *parser_) parseEntity() (
	entity ast.EntityLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Element Entity.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if ok {
		// Found a single Element Entity.
		entity = ast.EntityClass().Entity(element)
		return
	}

	// Attempt to parse a single String Entity.
	var string_ ast.StringLike
	string_, token, ok = v.parseString()
	if ok {
		// Found a single String Entity.
		entity = ast.EntityClass().Entity(string_)
		return
	}

	// Attempt to parse a single Range Entity.
	var range_ ast.RangeLike
	range_, token, ok = v.parseRange()
	if ok {
		// Found a single Range Entity.
		entity = ast.EntityClass().Entity(range_)
		return
	}

	// Attempt to parse a single Collection Entity.
	var collection ast.CollectionLike
	collection, token, ok = v.parseCollection()
	if ok {
		// Found a single Collection Entity.
		entity = ast.EntityClass().Entity(collection)
		return
	}

	// Attempt to parse a single Procedure Entity.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	if ok {
		// Found a single Procedure Entity.
		entity = ast.EntityClass().Entity(procedure)
		return
	}

	// This is not a single Entity rule.
	return
}

func (v *parser_) parseEquals() (
	equals ast.EqualsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "=" delimiter.
	_, token, ok = v.parseDelimiter("=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Equals rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Equals", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Equals rule.
	ok = true
	v.remove(tokens)
	equals = ast.EqualsClass().Equals()
	return
}

func (v *parser_) parseEvent() (
	event ast.EventLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Event rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Event", token)
		panic(message)
	}

	// Found a single Event rule.
	ok = true
	v.remove(tokens)
	event = ast.EventClass().Event(expression)
	return
}

func (v *parser_) parseException() (
	exception ast.ExceptionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Exception rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Exception", token)
		panic(message)
	}

	// Found a single Exception rule.
	ok = true
	v.remove(tokens)
	exception = ast.ExceptionClass().Exception(expression)
	return
}

func (v *parser_) parseExponential() (
	exponential ast.ExponentialLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression1 ast.ExpressionLike
	expression1, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Exponential rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Exponential", token)
		panic(message)
	}

	// Attempt to parse a single "^" delimiter.
	_, token, ok = v.parseDelimiter("^")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Exponential rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Exponential", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression2 ast.ExpressionLike
	expression2, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Exponential rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Exponential", token)
		panic(message)
	}

	// Found a single Exponential rule.
	ok = true
	v.remove(tokens)
	exponential = ast.ExponentialClass().Exponential(
		expression1,
		expression2,
	)
	return
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Expression.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Expression.
		expression = ast.ExpressionClass().Expression(component)
		return
	}

	// Attempt to parse a single Intrinsic Expression.
	var intrinsic ast.IntrinsicLike
	intrinsic, token, ok = v.parseIntrinsic()
	if ok {
		// Found a single Intrinsic Expression.
		expression = ast.ExpressionClass().Expression(intrinsic)
		return
	}

	// Attempt to parse a single Variable Expression.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Expression.
		expression = ast.ExpressionClass().Expression(variable)
		return
	}

	// Attempt to parse a single Precedence Expression.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Expression.
		expression = ast.ExpressionClass().Expression(precedence)
		return
	}

	// Attempt to parse a single Dereference Expression.
	var dereference ast.DereferenceLike
	dereference, token, ok = v.parseDereference()
	if ok {
		// Found a single Dereference Expression.
		expression = ast.ExpressionClass().Expression(dereference)
		return
	}

	// Attempt to parse a single Invocation Expression.
	var invocation ast.InvocationLike
	invocation, token, ok = v.parseInvocation()
	if ok {
		// Found a single Invocation Expression.
		expression = ast.ExpressionClass().Expression(invocation)
		return
	}

	// Attempt to parse a single Subcomponent Expression.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Expression.
		expression = ast.ExpressionClass().Expression(subcomponent)
		return
	}

	// Attempt to parse a single Chaining Expression.
	var chaining ast.ChainingLike
	chaining, token, ok = v.parseChaining()
	if ok {
		// Found a single Chaining Expression.
		expression = ast.ExpressionClass().Expression(chaining)
		return
	}

	// Attempt to parse a single Exponential Expression.
	var exponential ast.ExponentialLike
	exponential, token, ok = v.parseExponential()
	if ok {
		// Found a single Exponential Expression.
		expression = ast.ExpressionClass().Expression(exponential)
		return
	}

	// Attempt to parse a single Inversion Expression.
	var inversion ast.InversionLike
	inversion, token, ok = v.parseInversion()
	if ok {
		// Found a single Inversion Expression.
		expression = ast.ExpressionClass().Expression(inversion)
		return
	}

	// Attempt to parse a single Arithmetic Expression.
	var arithmetic ast.ArithmeticLike
	arithmetic, token, ok = v.parseArithmetic()
	if ok {
		// Found a single Arithmetic Expression.
		expression = ast.ExpressionClass().Expression(arithmetic)
		return
	}

	// Attempt to parse a single Amplitude Expression.
	var amplitude ast.AmplitudeLike
	amplitude, token, ok = v.parseAmplitude()
	if ok {
		// Found a single Amplitude Expression.
		expression = ast.ExpressionClass().Expression(amplitude)
		return
	}

	// Attempt to parse a single Comparison Expression.
	var comparison ast.ComparisonLike
	comparison, token, ok = v.parseComparison()
	if ok {
		// Found a single Comparison Expression.
		expression = ast.ExpressionClass().Expression(comparison)
		return
	}

	// Attempt to parse a single Complement Expression.
	var complement ast.ComplementLike
	complement, token, ok = v.parseComplement()
	if ok {
		// Found a single Complement Expression.
		expression = ast.ExpressionClass().Expression(complement)
		return
	}

	// Attempt to parse a single Logical Expression.
	var logical ast.LogicalLike
	logical, token, ok = v.parseLogical()
	if ok {
		// Found a single Logical Expression.
		expression = ast.ExpressionClass().Expression(logical)
		return
	}

	// This is not a single Expression rule.
	return
}

func (v *parser_) parseFailure() (
	failure ast.FailureLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Failure rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Failure", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Failure rule.
	ok = true
	v.remove(tokens)
	failure = ast.FailureClass().Failure(symbol)
	return
}

func (v *parser_) parseFlow() (
	flow ast.FlowLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single IfClause Flow.
	var ifClause ast.IfClauseLike
	ifClause, token, ok = v.parseIfClause()
	if ok {
		// Found a single IfClause Flow.
		flow = ast.FlowClass().Flow(ifClause)
		return
	}

	// Attempt to parse a single SelectClause Flow.
	var selectClause ast.SelectClauseLike
	selectClause, token, ok = v.parseSelectClause()
	if ok {
		// Found a single SelectClause Flow.
		flow = ast.FlowClass().Flow(selectClause)
		return
	}

	// Attempt to parse a single WhileClause Flow.
	var whileClause ast.WhileClauseLike
	whileClause, token, ok = v.parseWhileClause()
	if ok {
		// Found a single WhileClause Flow.
		flow = ast.FlowClass().Flow(whileClause)
		return
	}

	// Attempt to parse a single WithClause Flow.
	var withClause ast.WithClauseLike
	withClause, token, ok = v.parseWithClause()
	if ok {
		// Found a single WithClause Flow.
		flow = ast.FlowClass().Flow(withClause)
		return
	}

	// Attempt to parse a single ContinueClause Flow.
	var continueClause ast.ContinueClauseLike
	continueClause, token, ok = v.parseContinueClause()
	if ok {
		// Found a single ContinueClause Flow.
		flow = ast.FlowClass().Flow(continueClause)
		return
	}

	// Attempt to parse a single BreakClause Flow.
	var breakClause ast.BreakClauseLike
	breakClause, token, ok = v.parseBreakClause()
	if ok {
		// Found a single BreakClause Flow.
		flow = ast.FlowClass().Flow(breakClause)
		return
	}

	// Attempt to parse a single ReturnClause Flow.
	var returnClause ast.ReturnClauseLike
	returnClause, token, ok = v.parseReturnClause()
	if ok {
		// Found a single ReturnClause Flow.
		flow = ast.FlowClass().Flow(returnClause)
		return
	}

	// Attempt to parse a single ThrowClause Flow.
	var throwClause ast.ThrowClauseLike
	throwClause, token, ok = v.parseThrowClause()
	if ok {
		// Found a single ThrowClause Flow.
		flow = ast.FlowClass().Flow(throwClause)
		return
	}

	// This is not a single Flow rule.
	return
}

func (v *parser_) parseFunction() (
	function ast.FunctionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Function rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Function", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Function rule.
	ok = true
	v.remove(tokens)
	function = ast.FunctionClass().Function(identifier)
	return
}

func (v *parser_) parseIOR() (
	iOR ast.IORLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "IOR" delimiter.
	_, token, ok = v.parseDelimiter("IOR")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single IOR rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$IOR", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single IOR rule.
	ok = true
	v.remove(tokens)
	iOR = ast.IORClass().IOR()
	return
}

func (v *parser_) parseIfClause() (
	ifClause ast.IfClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "if" delimiter.
	_, token, ok = v.parseDelimiter("if")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single IfClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$IfClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Condition rule.
	var condition ast.ConditionLike
	condition, token, ok = v.parseCondition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single IfClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$IfClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single IfClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$IfClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single IfClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$IfClause", token)
		panic(message)
	}

	// Found a single IfClause rule.
	ok = true
	v.remove(tokens)
	ifClause = ast.IfClauseClass().IfClause(
		condition,
		procedure,
	)
	return
}

func (v *parser_) parseIndex() (
	index ast.IndexLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Index rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Index", token)
		panic(message)
	}

	// Found a single Index rule.
	ok = true
	v.remove(tokens)
	index = ast.IndexClass().Index(expression)
	return
}

func (v *parser_) parseIndices() (
	indices ast.IndicesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Index rule.
	var index ast.IndexLike
	index, token, ok = v.parseIndex()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Indices rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Indices", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalIndex rules.
	var additionalIndexes = fra.List[ast.AdditionalIndexLike]()
additionalIndexesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalIndex ast.AdditionalIndexLike
		additionalIndex, token, ok = v.parseAdditionalIndex()
		if !ok {
			switch {
			case count >= 0:
				break additionalIndexesLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalIndex rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Indices", token)
				message += "0 or more AdditionalIndex rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalIndexes.AppendValue(additionalIndex)
	}

	// Found a single Indices rule.
	ok = true
	v.remove(tokens)
	indices = ast.IndicesClass().Indices(
		index,
		additionalIndexes,
	)
	return
}

func (v *parser_) parseInduction() (
	induction ast.InductionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single DoClause Induction.
	var doClause ast.DoClauseLike
	doClause, token, ok = v.parseDoClause()
	if ok {
		// Found a single DoClause Induction.
		induction = ast.InductionClass().Induction(doClause)
		return
	}

	// Attempt to parse a single LetClause Induction.
	var letClause ast.LetClauseLike
	letClause, token, ok = v.parseLetClause()
	if ok {
		// Found a single LetClause Induction.
		induction = ast.InductionClass().Induction(letClause)
		return
	}

	// This is not a single Induction rule.
	return
}

func (v *parser_) parseInlineAssociations() (
	inlineAssociations ast.InlineAssociationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Association rule.
	var association ast.AssociationLike
	association, token, ok = v.parseAssociation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InlineAssociations rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineAssociations", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalAssociation rules.
	var additionalAssociations = fra.List[ast.AdditionalAssociationLike]()
additionalAssociationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalAssociation ast.AdditionalAssociationLike
		additionalAssociation, token, ok = v.parseAdditionalAssociation()
		if !ok {
			switch {
			case count >= 0:
				break additionalAssociationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalAssociation rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InlineAssociations", token)
				message += "0 or more AdditionalAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalAssociations.AppendValue(additionalAssociation)
	}

	// Found a single InlineAssociations rule.
	ok = true
	v.remove(tokens)
	inlineAssociations = ast.InlineAssociationsClass().InlineAssociations(
		association,
		additionalAssociations,
	)
	return
}

func (v *parser_) parseInlineParameters() (
	inlineParameters ast.InlineParametersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Parameter rule.
	var parameter ast.ParameterLike
	parameter, token, ok = v.parseParameter()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InlineParameters rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineParameters", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalParameter rules.
	var additionalParameters = fra.List[ast.AdditionalParameterLike]()
additionalParametersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalParameter ast.AdditionalParameterLike
		additionalParameter, token, ok = v.parseAdditionalParameter()
		if !ok {
			switch {
			case count >= 0:
				break additionalParametersLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalParameter rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InlineParameters", token)
				message += "0 or more AdditionalParameter rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalParameters.AppendValue(additionalParameter)
	}

	// Found a single InlineParameters rule.
	ok = true
	v.remove(tokens)
	inlineParameters = ast.InlineParametersClass().InlineParameters(
		parameter,
		additionalParameters,
	)
	return
}

func (v *parser_) parseInlineStatements() (
	inlineStatements ast.InlineStatementsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Statement rule.
	var statement ast.StatementLike
	statement, token, ok = v.parseStatement()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InlineStatements rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineStatements", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalStatement rules.
	var additionalStatements = fra.List[ast.AdditionalStatementLike]()
additionalStatementsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalStatement ast.AdditionalStatementLike
		additionalStatement, token, ok = v.parseAdditionalStatement()
		if !ok {
			switch {
			case count >= 0:
				break additionalStatementsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalStatement rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InlineStatements", token)
				message += "0 or more AdditionalStatement rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalStatements.AppendValue(additionalStatement)
	}

	// Found a single InlineStatements rule.
	ok = true
	v.remove(tokens)
	inlineStatements = ast.InlineStatementsClass().InlineStatements(
		statement,
		additionalStatements,
	)
	return
}

func (v *parser_) parseInlineValues() (
	inlineValues ast.InlineValuesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Value rule.
	var value ast.ValueLike
	value, token, ok = v.parseValue()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single InlineValues rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineValues", token)
		panic(message)
	}

	// Attempt to parse multiple AdditionalValue rules.
	var additionalValues = fra.List[ast.AdditionalValueLike]()
additionalValuesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var additionalValue ast.AdditionalValueLike
		additionalValue, token, ok = v.parseAdditionalValue()
		if !ok {
			switch {
			case count >= 0:
				break additionalValuesLoop
			case uti.IsDefined(tokens):
				// This is not multiple AdditionalValue rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$InlineValues", token)
				message += "0 or more AdditionalValue rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalValues.AppendValue(additionalValue)
	}

	// Found a single InlineValues rule.
	ok = true
	v.remove(tokens)
	inlineValues = ast.InlineValuesClass().InlineValues(
		value,
		additionalValues,
	)
	return
}

func (v *parser_) parseIntrinsic() (
	intrinsic ast.IntrinsicLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Function rule.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Intrinsic rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Intrinsic", token)
		panic(message)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Intrinsic rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Intrinsic", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Arguments rule.
	var optionalArguments ast.ArgumentsLike
	optionalArguments, _, ok = v.parseArguments()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Intrinsic rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Intrinsic", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Intrinsic rule.
	ok = true
	v.remove(tokens)
	intrinsic = ast.IntrinsicClass().Intrinsic(
		function,
		optionalArguments,
	)
	return
}

func (v *parser_) parseInversion() (
	inversion ast.InversionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single InversionOperator rule.
	var inversionOperator ast.InversionOperatorLike
	inversionOperator, token, ok = v.parseInversionOperator()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Inversion rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Inversion", token)
		panic(message)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Inversion rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Inversion", token)
		panic(message)
	}

	// Found a single Inversion rule.
	ok = true
	v.remove(tokens)
	inversion = ast.InversionClass().Inversion(
		inversionOperator,
		expression,
	)
	return
}

func (v *parser_) parseInversionOperator() (
	inversionOperator ast.InversionOperatorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Minus InversionOperator.
	var minus ast.MinusLike
	minus, token, ok = v.parseMinus()
	if ok {
		// Found a single Minus InversionOperator.
		inversionOperator = ast.InversionOperatorClass().InversionOperator(minus)
		return
	}

	// Attempt to parse a single Divide InversionOperator.
	var divide ast.DivideLike
	divide, token, ok = v.parseDivide()
	if ok {
		// Found a single Divide InversionOperator.
		inversionOperator = ast.InversionOperatorClass().InversionOperator(divide)
		return
	}

	// Attempt to parse a single Times InversionOperator.
	var times ast.TimesLike
	times, token, ok = v.parseTimes()
	if ok {
		// Found a single Times InversionOperator.
		inversionOperator = ast.InversionOperatorClass().InversionOperator(times)
		return
	}

	// This is not a single InversionOperator rule.
	return
}

func (v *parser_) parseInvocation() (
	invocation ast.InvocationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Target rule.
	var target ast.TargetLike
	target, token, ok = v.parseTarget()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Invocation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Invocation", token)
		panic(message)
	}

	// Attempt to parse a single InvocationOperator rule.
	var invocationOperator ast.InvocationOperatorLike
	invocationOperator, token, ok = v.parseInvocationOperator()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Invocation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Invocation", token)
		panic(message)
	}

	// Attempt to parse a single Method rule.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Invocation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Invocation", token)
		panic(message)
	}

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Invocation rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Invocation", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Arguments rule.
	var optionalArguments ast.ArgumentsLike
	optionalArguments, _, ok = v.parseArguments()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Invocation rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Invocation", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Invocation rule.
	ok = true
	v.remove(tokens)
	invocation = ast.InvocationClass().Invocation(
		target,
		invocationOperator,
		method,
		optionalArguments,
	)
	return
}

func (v *parser_) parseInvocationOperator() (
	invocationOperator ast.InvocationOperatorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Dot InvocationOperator.
	var dot ast.DotLike
	dot, token, ok = v.parseDot()
	if ok {
		// Found a single Dot InvocationOperator.
		invocationOperator = ast.InvocationOperatorClass().InvocationOperator(dot)
		return
	}

	// Attempt to parse a single Arrow InvocationOperator.
	var arrow ast.ArrowLike
	arrow, token, ok = v.parseArrow()
	if ok {
		// Found a single Arrow InvocationOperator.
		invocationOperator = ast.InvocationOperatorClass().InvocationOperator(arrow)
		return
	}

	// This is not a single InvocationOperator rule.
	return
}

func (v *parser_) parseIs() (
	is ast.IsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "IS" delimiter.
	_, token, ok = v.parseDelimiter("IS")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Is rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Is", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Is rule.
	ok = true
	v.remove(tokens)
	is = ast.IsClass().Is()
	return
}

func (v *parser_) parseItem() (
	item ast.ItemLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Item rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Item", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Item rule.
	ok = true
	v.remove(tokens)
	item = ast.ItemClass().Item(symbol)
	return
}

func (v *parser_) parseItems() (
	items ast.ItemsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Associations Items.
	var associations ast.AssociationsLike
	associations, token, ok = v.parseAssociations()
	if ok {
		// Found a single Associations Items.
		items = ast.ItemsClass().Items(associations)
		return
	}

	// Attempt to parse a single Values Items.
	var values ast.ValuesLike
	values, token, ok = v.parseValues()
	if ok {
		// Found a single Values Items.
		items = ast.ItemsClass().Items(values)
		return
	}

	// This is not a single Items rule.
	return
}

func (v *parser_) parseKey() (
	key ast.KeyLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Primitive rule.
	var primitive ast.PrimitiveLike
	primitive, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Key rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Key", token)
		panic(message)
	}

	// Found a single Key rule.
	ok = true
	v.remove(tokens)
	key = ast.KeyClass().Key(primitive)
	return
}

func (v *parser_) parseLabel() (
	label ast.LabelLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Label rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Label", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Label rule.
	ok = true
	v.remove(tokens)
	label = ast.LabelClass().Label(symbol)
	return
}

func (v *parser_) parseLeftBracket() (
	leftBracket ast.LeftBracketLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single LeftSquare LeftBracket.
	var leftSquare ast.LeftSquareLike
	leftSquare, token, ok = v.parseLeftSquare()
	if ok {
		// Found a single LeftSquare LeftBracket.
		leftBracket = ast.LeftBracketClass().LeftBracket(leftSquare)
		return
	}

	// Attempt to parse a single LeftRound LeftBracket.
	var leftRound ast.LeftRoundLike
	leftRound, token, ok = v.parseLeftRound()
	if ok {
		// Found a single LeftRound LeftBracket.
		leftBracket = ast.LeftBracketClass().LeftBracket(leftRound)
		return
	}

	// This is not a single LeftBracket rule.
	return
}

func (v *parser_) parseLeftRound() (
	leftRound ast.LeftRoundLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single LeftRound rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$LeftRound", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single LeftRound rule.
	ok = true
	v.remove(tokens)
	leftRound = ast.LeftRoundClass().LeftRound()
	return
}

func (v *parser_) parseLeftSquare() (
	leftSquare ast.LeftSquareLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single LeftSquare rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$LeftSquare", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single LeftSquare rule.
	ok = true
	v.remove(tokens)
	leftSquare = ast.LeftSquareClass().LeftSquare()
	return
}

func (v *parser_) parseLess() (
	less ast.LessLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "<" delimiter.
	_, token, ok = v.parseDelimiter("<")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Less rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Less", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Less rule.
	ok = true
	v.remove(tokens)
	less = ast.LessClass().Less()
	return
}

func (v *parser_) parseLetClause() (
	letClause ast.LetClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "let" delimiter.
	_, token, ok = v.parseDelimiter("let")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single LetClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$LetClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Recipient rule.
	var recipient ast.RecipientLike
	recipient, token, ok = v.parseRecipient()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single LetClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$LetClause", token)
		panic(message)
	}

	// Attempt to parse a single AssignmentOperator rule.
	var assignmentOperator ast.AssignmentOperatorLike
	assignmentOperator, token, ok = v.parseAssignmentOperator()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single LetClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$LetClause", token)
		panic(message)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single LetClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$LetClause", token)
		panic(message)
	}

	// Found a single LetClause rule.
	ok = true
	v.remove(tokens)
	letClause = ast.LetClauseClass().LetClause(
		recipient,
		assignmentOperator,
		expression,
	)
	return
}

func (v *parser_) parseLogical() (
	logical ast.LogicalLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression1 ast.ExpressionLike
	expression1, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Logical rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Logical", token)
		panic(message)
	}

	// Attempt to parse a single LogicalOperator rule.
	var logicalOperator ast.LogicalOperatorLike
	logicalOperator, token, ok = v.parseLogicalOperator()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Logical rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Logical", token)
		panic(message)
	}

	// Attempt to parse a single Expression rule.
	var expression2 ast.ExpressionLike
	expression2, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Logical rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Logical", token)
		panic(message)
	}

	// Found a single Logical rule.
	ok = true
	v.remove(tokens)
	logical = ast.LogicalClass().Logical(
		expression1,
		logicalOperator,
		expression2,
	)
	return
}

func (v *parser_) parseLogicalOperator() (
	logicalOperator ast.LogicalOperatorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single And LogicalOperator.
	var and ast.AndLike
	and, token, ok = v.parseAnd()
	if ok {
		// Found a single And LogicalOperator.
		logicalOperator = ast.LogicalOperatorClass().LogicalOperator(and)
		return
	}

	// Attempt to parse a single SAN LogicalOperator.
	var sAN ast.SANLike
	sAN, token, ok = v.parseSAN()
	if ok {
		// Found a single SAN LogicalOperator.
		logicalOperator = ast.LogicalOperatorClass().LogicalOperator(sAN)
		return
	}

	// Attempt to parse a single IOR LogicalOperator.
	var iOR ast.IORLike
	iOR, token, ok = v.parseIOR()
	if ok {
		// Found a single IOR LogicalOperator.
		logicalOperator = ast.LogicalOperatorClass().LogicalOperator(iOR)
		return
	}

	// Attempt to parse a single XOR LogicalOperator.
	var xOR ast.XORLike
	xOR, token, ok = v.parseXOR()
	if ok {
		// Found a single XOR LogicalOperator.
		logicalOperator = ast.LogicalOperatorClass().LogicalOperator(xOR)
		return
	}

	// This is not a single LogicalOperator rule.
	return
}

func (v *parser_) parseMainClause() (
	mainClause ast.MainClauseLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Flow MainClause.
	var flow ast.FlowLike
	flow, token, ok = v.parseFlow()
	if ok {
		// Found a single Flow MainClause.
		mainClause = ast.MainClauseClass().MainClause(flow)
		return
	}

	// Attempt to parse a single Induction MainClause.
	var induction ast.InductionLike
	induction, token, ok = v.parseInduction()
	if ok {
		// Found a single Induction MainClause.
		mainClause = ast.MainClauseClass().MainClause(induction)
		return
	}

	// Attempt to parse a single Messaging MainClause.
	var messaging ast.MessagingLike
	messaging, token, ok = v.parseMessaging()
	if ok {
		// Found a single Messaging MainClause.
		mainClause = ast.MainClauseClass().MainClause(messaging)
		return
	}

	// Attempt to parse a single Repository MainClause.
	var repository ast.RepositoryLike
	repository, token, ok = v.parseRepository()
	if ok {
		// Found a single Repository MainClause.
		mainClause = ast.MainClauseClass().MainClause(repository)
		return
	}

	// This is not a single MainClause rule.
	return
}

func (v *parser_) parseMatchHandler() (
	matchHandler ast.MatchHandlerLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "matching" delimiter.
	_, token, ok = v.parseDelimiter("matching")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MatchHandler rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MatchHandler", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Template rule.
	var template ast.TemplateLike
	template, token, ok = v.parseTemplate()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single MatchHandler rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$MatchHandler", token)
		panic(message)
	}

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MatchHandler rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MatchHandler", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single MatchHandler rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$MatchHandler", token)
		panic(message)
	}

	// Found a single MatchHandler rule.
	ok = true
	v.remove(tokens)
	matchHandler = ast.MatchHandlerClass().MatchHandler(
		template,
		procedure,
	)
	return
}

func (v *parser_) parseMatches() (
	matches ast.MatchesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "MATCHES" delimiter.
	_, token, ok = v.parseDelimiter("MATCHES")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Matches rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Matches", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Matches rule.
	ok = true
	v.remove(tokens)
	matches = ast.MatchesClass().Matches()
	return
}

func (v *parser_) parseMessage() (
	message ast.MessageLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Message rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Message", token)
		panic(message)
	}

	// Found a single Message rule.
	ok = true
	v.remove(tokens)
	message = ast.MessageClass().Message(expression)
	return
}

func (v *parser_) parseMessaging() (
	messaging ast.MessagingLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single PostClause Messaging.
	var postClause ast.PostClauseLike
	postClause, token, ok = v.parsePostClause()
	if ok {
		// Found a single PostClause Messaging.
		messaging = ast.MessagingClass().Messaging(postClause)
		return
	}

	// Attempt to parse a single RetrieveClause Messaging.
	var retrieveClause ast.RetrieveClauseLike
	retrieveClause, token, ok = v.parseRetrieveClause()
	if ok {
		// Found a single RetrieveClause Messaging.
		messaging = ast.MessagingClass().Messaging(retrieveClause)
		return
	}

	// Attempt to parse a single AcceptClause Messaging.
	var acceptClause ast.AcceptClauseLike
	acceptClause, token, ok = v.parseAcceptClause()
	if ok {
		// Found a single AcceptClause Messaging.
		messaging = ast.MessagingClass().Messaging(acceptClause)
		return
	}

	// Attempt to parse a single RejectClause Messaging.
	var rejectClause ast.RejectClauseLike
	rejectClause, token, ok = v.parseRejectClause()
	if ok {
		// Found a single RejectClause Messaging.
		messaging = ast.MessagingClass().Messaging(rejectClause)
		return
	}

	// Attempt to parse a single PublishClause Messaging.
	var publishClause ast.PublishClauseLike
	publishClause, token, ok = v.parsePublishClause()
	if ok {
		// Found a single PublishClause Messaging.
		messaging = ast.MessagingClass().Messaging(publishClause)
		return
	}

	// This is not a single Messaging rule.
	return
}

func (v *parser_) parseMethod() (
	method ast.MethodLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Method rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Method", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Method rule.
	ok = true
	v.remove(tokens)
	method = ast.MethodClass().Method(identifier)
	return
}

func (v *parser_) parseMinus() (
	minus ast.MinusLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "-" delimiter.
	_, token, ok = v.parseDelimiter("-")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Minus rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Minus", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Minus rule.
	ok = true
	v.remove(tokens)
	minus = ast.MinusClass().Minus()
	return
}

func (v *parser_) parseMinusEquals() (
	minusEquals ast.MinusEqualsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "-=" delimiter.
	_, token, ok = v.parseDelimiter("-=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MinusEquals rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MinusEquals", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single MinusEquals rule.
	ok = true
	v.remove(tokens)
	minusEquals = ast.MinusEqualsClass().MinusEquals()
	return
}

func (v *parser_) parseModulo() (
	modulo ast.ModuloLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "//" delimiter.
	_, token, ok = v.parseDelimiter("//")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Modulo rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Modulo", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Modulo rule.
	ok = true
	v.remove(tokens)
	modulo = ast.ModuloClass().Modulo()
	return
}

func (v *parser_) parseMore() (
	more ast.MoreLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ">" delimiter.
	_, token, ok = v.parseDelimiter(">")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single More rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$More", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single More rule.
	ok = true
	v.remove(tokens)
	more = ast.MoreClass().More()
	return
}

func (v *parser_) parseMultilineAssociations() (
	multilineAssociations ast.MultilineAssociationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple AnnotatedAssociation rules.
	var annotatedAssociations = fra.List[ast.AnnotatedAssociationLike]()
annotatedAssociationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var annotatedAssociation ast.AnnotatedAssociationLike
		annotatedAssociation, token, ok = v.parseAnnotatedAssociation()
		if !ok {
			switch {
			case count >= 1:
				break annotatedAssociationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AnnotatedAssociation rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$MultilineAssociations", token)
				message += "1 or more AnnotatedAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedAssociations.AppendValue(annotatedAssociation)
	}

	// Found a single MultilineAssociations rule.
	ok = true
	v.remove(tokens)
	multilineAssociations = ast.MultilineAssociationsClass().MultilineAssociations(annotatedAssociations)
	return
}

func (v *parser_) parseMultilineParameters() (
	multilineParameters ast.MultilineParametersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple AnnotatedParameter rules.
	var annotatedParameters = fra.List[ast.AnnotatedParameterLike]()
annotatedParametersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var annotatedParameter ast.AnnotatedParameterLike
		annotatedParameter, token, ok = v.parseAnnotatedParameter()
		if !ok {
			switch {
			case count >= 1:
				break annotatedParametersLoop
			case uti.IsDefined(tokens):
				// This is not multiple AnnotatedParameter rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$MultilineParameters", token)
				message += "1 or more AnnotatedParameter rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedParameters.AppendValue(annotatedParameter)
	}

	// Found a single MultilineParameters rule.
	ok = true
	v.remove(tokens)
	multilineParameters = ast.MultilineParametersClass().MultilineParameters(annotatedParameters)
	return
}

func (v *parser_) parseMultilineStatements() (
	multilineStatements ast.MultilineStatementsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple AnnotatedStatement rules.
	var annotatedStatements = fra.List[ast.AnnotatedStatementLike]()
annotatedStatementsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var annotatedStatement ast.AnnotatedStatementLike
		annotatedStatement, token, ok = v.parseAnnotatedStatement()
		if !ok {
			switch {
			case count >= 1:
				break annotatedStatementsLoop
			case uti.IsDefined(tokens):
				// This is not multiple AnnotatedStatement rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$MultilineStatements", token)
				message += "1 or more AnnotatedStatement rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedStatements.AppendValue(annotatedStatement)
	}

	// Found a single MultilineStatements rule.
	ok = true
	v.remove(tokens)
	multilineStatements = ast.MultilineStatementsClass().MultilineStatements(annotatedStatements)
	return
}

func (v *parser_) parseMultilineValues() (
	multilineValues ast.MultilineValuesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse multiple AnnotatedValue rules.
	var annotatedValues = fra.List[ast.AnnotatedValueLike]()
annotatedValuesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var annotatedValue ast.AnnotatedValueLike
		annotatedValue, token, ok = v.parseAnnotatedValue()
		if !ok {
			switch {
			case count >= 1:
				break annotatedValuesLoop
			case uti.IsDefined(tokens):
				// This is not multiple AnnotatedValue rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$MultilineValues", token)
				message += "1 or more AnnotatedValue rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedValues.AppendValue(annotatedValue)
	}

	// Found a single MultilineValues rule.
	ok = true
	v.remove(tokens)
	multilineValues = ast.MultilineValuesClass().MultilineValues(annotatedValues)
	return
}

func (v *parser_) parseNoAssociations() (
	noAssociations ast.NoAssociationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoAssociations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoAssociations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single NoAssociations rule.
	ok = true
	v.remove(tokens)
	noAssociations = ast.NoAssociationsClass().NoAssociations()
	return
}

func (v *parser_) parseNotarizeClause() (
	notarizeClause ast.NotarizeClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "notarize" delimiter.
	_, token, ok = v.parseDelimiter("notarize")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NotarizeClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NotarizeClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Draft rule.
	var draft ast.DraftLike
	draft, token, ok = v.parseDraft()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single NotarizeClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$NotarizeClause", token)
		panic(message)
	}

	// Attempt to parse a single "as" delimiter.
	_, token, ok = v.parseDelimiter("as")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NotarizeClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NotarizeClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Citation rule.
	var citation ast.CitationLike
	citation, token, ok = v.parseCitation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single NotarizeClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$NotarizeClause", token)
		panic(message)
	}

	// Found a single NotarizeClause rule.
	ok = true
	v.remove(tokens)
	notarizeClause = ast.NotarizeClauseClass().NotarizeClause(
		draft,
		citation,
	)
	return
}

func (v *parser_) parseOnClause() (
	onClause ast.OnClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "on" delimiter.
	_, token, ok = v.parseDelimiter("on")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single OnClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$OnClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Failure rule.
	var failure ast.FailureLike
	failure, token, ok = v.parseFailure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single OnClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$OnClause", token)
		panic(message)
	}

	// Attempt to parse multiple MatchHandler rules.
	var matchHandlers = fra.List[ast.MatchHandlerLike]()
matchHandlersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var matchHandler ast.MatchHandlerLike
		matchHandler, token, ok = v.parseMatchHandler()
		if !ok {
			switch {
			case count >= 1:
				break matchHandlersLoop
			case uti.IsDefined(tokens):
				// This is not multiple MatchHandler rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$OnClause", token)
				message += "1 or more MatchHandler rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		matchHandlers.AppendValue(matchHandler)
	}

	// Found a single OnClause rule.
	ok = true
	v.remove(tokens)
	onClause = ast.OnClauseClass().OnClause(
		failure,
		matchHandlers,
	)
	return
}

func (v *parser_) parseOperation() (
	operation ast.OperationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Intrinsic Operation.
	var intrinsic ast.IntrinsicLike
	intrinsic, token, ok = v.parseIntrinsic()
	if ok {
		// Found a single Intrinsic Operation.
		operation = ast.OperationClass().Operation(intrinsic)
		return
	}

	// Attempt to parse a single Invocation Operation.
	var invocation ast.InvocationLike
	invocation, token, ok = v.parseInvocation()
	if ok {
		// Found a single Invocation Operation.
		operation = ast.OperationClass().Operation(invocation)
		return
	}

	// This is not a single Operation rule.
	return
}

func (v *parser_) parseParameter() (
	parameter ast.ParameterLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Label rule.
	var label ast.LabelLike
	label, token, ok = v.parseLabel()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Parameter rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Parameter", token)
		panic(message)
	}

	// Attempt to parse a single ":" delimiter.
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameter rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameter", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Parameter rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Parameter", token)
		panic(message)
	}

	// Found a single Parameter rule.
	ok = true
	v.remove(tokens)
	parameter = ast.ParameterClass().Parameter(
		label,
		component,
	)
	return
}

func (v *parser_) parseParameters() (
	parameters ast.ParametersLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single InlineParameters Parameters.
	var inlineParameters ast.InlineParametersLike
	inlineParameters, token, ok = v.parseInlineParameters()
	if ok {
		// Found a single InlineParameters Parameters.
		parameters = ast.ParametersClass().Parameters(inlineParameters)
		return
	}

	// Attempt to parse a single MultilineParameters Parameters.
	var multilineParameters ast.MultilineParametersLike
	multilineParameters, token, ok = v.parseMultilineParameters()
	if ok {
		// Found a single MultilineParameters Parameters.
		parameters = ast.ParametersClass().Parameters(multilineParameters)
		return
	}

	// This is not a single Parameters rule.
	return
}

func (v *parser_) parsePlus() (
	plus ast.PlusLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "+" delimiter.
	_, token, ok = v.parseDelimiter("+")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Plus rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Plus", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Plus rule.
	ok = true
	v.remove(tokens)
	plus = ast.PlusClass().Plus()
	return
}

func (v *parser_) parsePlusEquals() (
	plusEquals ast.PlusEqualsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "+=" delimiter.
	_, token, ok = v.parseDelimiter("+=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PlusEquals rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PlusEquals", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single PlusEquals rule.
	ok = true
	v.remove(tokens)
	plusEquals = ast.PlusEqualsClass().PlusEquals()
	return
}

func (v *parser_) parsePostClause() (
	postClause ast.PostClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "post" delimiter.
	_, token, ok = v.parseDelimiter("post")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PostClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PostClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Message rule.
	var message ast.MessageLike
	message, token, ok = v.parseMessage()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single PostClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PostClause", token)
		panic(message)
	}

	// Attempt to parse a single "to" delimiter.
	_, token, ok = v.parseDelimiter("to")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PostClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PostClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Bag rule.
	var bag ast.BagLike
	bag, token, ok = v.parseBag()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single PostClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PostClause", token)
		panic(message)
	}

	// Found a single PostClause rule.
	ok = true
	v.remove(tokens)
	postClause = ast.PostClauseClass().PostClause(
		message,
		bag,
	)
	return
}

func (v *parser_) parsePrecedence() (
	precedence ast.PrecedenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Precedence rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Precedence", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Precedence rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Precedence", token)
		panic(message)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Precedence rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Precedence", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Precedence rule.
	ok = true
	v.remove(tokens)
	precedence = ast.PrecedenceClass().Precedence(expression)
	return
}

func (v *parser_) parsePrimitive() (
	primitive ast.PrimitiveLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Element Primitive.
	var element ast.ElementLike
	element, token, ok = v.parseElement()
	if ok {
		// Found a single Element Primitive.
		primitive = ast.PrimitiveClass().Primitive(element)
		return
	}

	// Attempt to parse a single String Primitive.
	var string_ ast.StringLike
	string_, token, ok = v.parseString()
	if ok {
		// Found a single String Primitive.
		primitive = ast.PrimitiveClass().Primitive(string_)
		return
	}

	// This is not a single Primitive rule.
	return
}

func (v *parser_) parseProcedure() (
	procedure ast.ProcedureLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Procedure rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Procedure", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse an optional Statements rule.
	var optionalStatements ast.StatementsLike
	optionalStatements, _, ok = v.parseStatements()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Procedure rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Procedure", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Procedure rule.
	ok = true
	v.remove(tokens)
	procedure = ast.ProcedureClass().Procedure(optionalStatements)
	return
}

func (v *parser_) parsePublishClause() (
	publishClause ast.PublishClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "publish" delimiter.
	_, token, ok = v.parseDelimiter("publish")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single PublishClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$PublishClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Event rule.
	var event ast.EventLike
	event, token, ok = v.parseEvent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single PublishClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$PublishClause", token)
		panic(message)
	}

	// Found a single PublishClause rule.
	ok = true
	v.remove(tokens)
	publishClause = ast.PublishClauseClass().PublishClause(event)
	return
}

func (v *parser_) parseRange() (
	range_ ast.RangeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single LeftBracket rule.
	var leftBracket ast.LeftBracketLike
	leftBracket, token, ok = v.parseLeftBracket()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Range rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Range", token)
		panic(message)
	}

	// Attempt to parse a single Primitive rule.
	var primitive1 ast.PrimitiveLike
	primitive1, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Range rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Range", token)
		panic(message)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Range rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Range", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Primitive rule.
	var primitive2 ast.PrimitiveLike
	primitive2, token, ok = v.parsePrimitive()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Range rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Range", token)
		panic(message)
	}

	// Attempt to parse a single RightBracket rule.
	var rightBracket ast.RightBracketLike
	rightBracket, token, ok = v.parseRightBracket()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Range rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Range", token)
		panic(message)
	}

	// Found a single Range rule.
	ok = true
	v.remove(tokens)
	range_ = ast.RangeClass().Range(
		leftBracket,
		primitive1,
		primitive2,
		rightBracket,
	)
	return
}

func (v *parser_) parseRecipient() (
	recipient ast.RecipientLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Label Recipient.
	var label ast.LabelLike
	label, token, ok = v.parseLabel()
	if ok {
		// Found a single Label Recipient.
		recipient = ast.RecipientClass().Recipient(label)
		return
	}

	// Attempt to parse a single Attribute Recipient.
	var attribute ast.AttributeLike
	attribute, token, ok = v.parseAttribute()
	if ok {
		// Found a single Attribute Recipient.
		recipient = ast.RecipientClass().Recipient(attribute)
		return
	}

	// This is not a single Recipient rule.
	return
}

func (v *parser_) parseRejectClause() (
	rejectClause ast.RejectClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "reject" delimiter.
	_, token, ok = v.parseDelimiter("reject")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RejectClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RejectClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Message rule.
	var message ast.MessageLike
	message, token, ok = v.parseMessage()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single RejectClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RejectClause", token)
		panic(message)
	}

	// Found a single RejectClause rule.
	ok = true
	v.remove(tokens)
	rejectClause = ast.RejectClauseClass().RejectClause(message)
	return
}

func (v *parser_) parseRepository() (
	repository ast.RepositoryLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single CheckoutClause Repository.
	var checkoutClause ast.CheckoutClauseLike
	checkoutClause, token, ok = v.parseCheckoutClause()
	if ok {
		// Found a single CheckoutClause Repository.
		repository = ast.RepositoryClass().Repository(checkoutClause)
		return
	}

	// Attempt to parse a single SaveClause Repository.
	var saveClause ast.SaveClauseLike
	saveClause, token, ok = v.parseSaveClause()
	if ok {
		// Found a single SaveClause Repository.
		repository = ast.RepositoryClass().Repository(saveClause)
		return
	}

	// Attempt to parse a single DiscardClause Repository.
	var discardClause ast.DiscardClauseLike
	discardClause, token, ok = v.parseDiscardClause()
	if ok {
		// Found a single DiscardClause Repository.
		repository = ast.RepositoryClass().Repository(discardClause)
		return
	}

	// Attempt to parse a single NotarizeClause Repository.
	var notarizeClause ast.NotarizeClauseLike
	notarizeClause, token, ok = v.parseNotarizeClause()
	if ok {
		// Found a single NotarizeClause Repository.
		repository = ast.RepositoryClass().Repository(notarizeClause)
		return
	}

	// This is not a single Repository rule.
	return
}

func (v *parser_) parseResult() (
	result ast.ResultLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Result rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Result", token)
		panic(message)
	}

	// Found a single Result rule.
	ok = true
	v.remove(tokens)
	result = ast.ResultClass().Result(expression)
	return
}

func (v *parser_) parseRetrieveClause() (
	retrieveClause ast.RetrieveClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "retrieve" delimiter.
	_, token, ok = v.parseDelimiter("retrieve")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RetrieveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RetrieveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Recipient rule.
	var recipient ast.RecipientLike
	recipient, token, ok = v.parseRecipient()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single RetrieveClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RetrieveClause", token)
		panic(message)
	}

	// Attempt to parse a single "from" delimiter.
	_, token, ok = v.parseDelimiter("from")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RetrieveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RetrieveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Bag rule.
	var bag ast.BagLike
	bag, token, ok = v.parseBag()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single RetrieveClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$RetrieveClause", token)
		panic(message)
	}

	// Found a single RetrieveClause rule.
	ok = true
	v.remove(tokens)
	retrieveClause = ast.RetrieveClauseClass().RetrieveClause(
		recipient,
		bag,
	)
	return
}

func (v *parser_) parseReturnClause() (
	returnClause ast.ReturnClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "return" delimiter.
	_, token, ok = v.parseDelimiter("return")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ReturnClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ReturnClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Result rule.
	var result ast.ResultLike
	result, token, ok = v.parseResult()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single ReturnClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ReturnClause", token)
		panic(message)
	}

	// Found a single ReturnClause rule.
	ok = true
	v.remove(tokens)
	returnClause = ast.ReturnClauseClass().ReturnClause(result)
	return
}

func (v *parser_) parseRightBracket() (
	rightBracket ast.RightBracketLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single RightSquare RightBracket.
	var rightSquare ast.RightSquareLike
	rightSquare, token, ok = v.parseRightSquare()
	if ok {
		// Found a single RightSquare RightBracket.
		rightBracket = ast.RightBracketClass().RightBracket(rightSquare)
		return
	}

	// Attempt to parse a single RightRound RightBracket.
	var rightRound ast.RightRoundLike
	rightRound, token, ok = v.parseRightRound()
	if ok {
		// Found a single RightRound RightBracket.
		rightBracket = ast.RightBracketClass().RightBracket(rightRound)
		return
	}

	// This is not a single RightBracket rule.
	return
}

func (v *parser_) parseRightRound() (
	rightRound ast.RightRoundLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RightRound rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RightRound", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single RightRound rule.
	ok = true
	v.remove(tokens)
	rightRound = ast.RightRoundClass().RightRound()
	return
}

func (v *parser_) parseRightSquare() (
	rightSquare ast.RightSquareLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single RightSquare rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$RightSquare", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single RightSquare rule.
	ok = true
	v.remove(tokens)
	rightSquare = ast.RightSquareClass().RightSquare()
	return
}

func (v *parser_) parseSAN() (
	sAN ast.SANLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "SAN" delimiter.
	_, token, ok = v.parseDelimiter("SAN")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SAN rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SAN", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single SAN rule.
	ok = true
	v.remove(tokens)
	sAN = ast.SANClass().SAN()
	return
}

func (v *parser_) parseSaveClause() (
	saveClause ast.SaveClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "save" delimiter.
	_, token, ok = v.parseDelimiter("save")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SaveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SaveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Draft rule.
	var draft ast.DraftLike
	draft, token, ok = v.parseDraft()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single SaveClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SaveClause", token)
		panic(message)
	}

	// Attempt to parse a single "as" delimiter.
	_, token, ok = v.parseDelimiter("as")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SaveClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SaveClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Citation rule.
	var citation ast.CitationLike
	citation, token, ok = v.parseCitation()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single SaveClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SaveClause", token)
		panic(message)
	}

	// Found a single SaveClause rule.
	ok = true
	v.remove(tokens)
	saveClause = ast.SaveClauseClass().SaveClause(
		draft,
		citation,
	)
	return
}

func (v *parser_) parseSelectClause() (
	selectClause ast.SelectClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "select" delimiter.
	_, token, ok = v.parseDelimiter("select")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single SelectClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$SelectClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Target rule.
	var target ast.TargetLike
	target, token, ok = v.parseTarget()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single SelectClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$SelectClause", token)
		panic(message)
	}

	// Attempt to parse multiple MatchHandler rules.
	var matchHandlers = fra.List[ast.MatchHandlerLike]()
matchHandlersLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var matchHandler ast.MatchHandlerLike
		matchHandler, token, ok = v.parseMatchHandler()
		if !ok {
			switch {
			case count >= 1:
				break matchHandlersLoop
			case uti.IsDefined(tokens):
				// This is not multiple MatchHandler rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$SelectClause", token)
				message += "1 or more MatchHandler rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		matchHandlers.AppendValue(matchHandler)
	}

	// Found a single SelectClause rule.
	ok = true
	v.remove(tokens)
	selectClause = ast.SelectClauseClass().SelectClause(
		target,
		matchHandlers,
	)
	return
}

func (v *parser_) parseSequence() (
	sequence ast.SequenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Sequence rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Sequence", token)
		panic(message)
	}

	// Found a single Sequence rule.
	ok = true
	v.remove(tokens)
	sequence = ast.SequenceClass().Sequence(expression)
	return
}

func (v *parser_) parseStatement() (
	statement ast.StatementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single MainClause rule.
	var mainClause ast.MainClauseLike
	mainClause, token, ok = v.parseMainClause()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Statement rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Statement", token)
		panic(message)
	}

	// Attempt to parse an optional OnClause rule.
	var optionalOnClause ast.OnClauseLike
	optionalOnClause, _, ok = v.parseOnClause()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Statement rule.
	ok = true
	v.remove(tokens)
	statement = ast.StatementClass().Statement(
		mainClause,
		optionalOnClause,
	)
	return
}

func (v *parser_) parseStatements() (
	statements ast.StatementsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single InlineStatements Statements.
	var inlineStatements ast.InlineStatementsLike
	inlineStatements, token, ok = v.parseInlineStatements()
	if ok {
		// Found a single InlineStatements Statements.
		statements = ast.StatementsClass().Statements(inlineStatements)
		return
	}

	// Attempt to parse a single MultilineStatements Statements.
	var multilineStatements ast.MultilineStatementsLike
	multilineStatements, token, ok = v.parseMultilineStatements()
	if ok {
		// Found a single MultilineStatements Statements.
		statements = ast.StatementsClass().Statements(multilineStatements)
		return
	}

	// This is not a single Statements rule.
	return
}

func (v *parser_) parseString() (
	string_ ast.StringLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single binary String.
	var binary string
	binary, token, ok = v.parseToken(BinaryToken)
	if ok {
		// Found a single binary String.
		string_ = ast.StringClass().String(binary)
		return
	}

	// Attempt to parse a single bytecode String.
	var bytecode string
	bytecode, token, ok = v.parseToken(BytecodeToken)
	if ok {
		// Found a single bytecode String.
		string_ = ast.StringClass().String(bytecode)
		return
	}

	// Attempt to parse a single name String.
	var name string
	name, token, ok = v.parseToken(NameToken)
	if ok {
		// Found a single name String.
		string_ = ast.StringClass().String(name)
		return
	}

	// Attempt to parse a single narrative String.
	var narrative string
	narrative, token, ok = v.parseToken(NarrativeToken)
	if ok {
		// Found a single narrative String.
		string_ = ast.StringClass().String(narrative)
		return
	}

	// Attempt to parse a single quote String.
	var quote string
	quote, token, ok = v.parseToken(QuoteToken)
	if ok {
		// Found a single quote String.
		string_ = ast.StringClass().String(quote)
		return
	}

	// Attempt to parse a single symbol String.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if ok {
		// Found a single symbol String.
		string_ = ast.StringClass().String(symbol)
		return
	}

	// Attempt to parse a single tag String.
	var tag string
	tag, token, ok = v.parseToken(TagToken)
	if ok {
		// Found a single tag String.
		string_ = ast.StringClass().String(tag)
		return
	}

	// Attempt to parse a single version String.
	var version string
	version, token, ok = v.parseToken(VersionToken)
	if ok {
		// Found a single version String.
		string_ = ast.StringClass().String(version)
		return
	}

	// This is not a single String rule.
	return
}

func (v *parser_) parseSubcomponent() (
	subcomponent ast.SubcomponentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Composite rule.
	var composite ast.CompositeLike
	composite, token, ok = v.parseComposite()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Subcomponent rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Subcomponent", token)
		panic(message)
	}

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Subcomponent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Subcomponent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Indices rule.
	var indices ast.IndicesLike
	indices, token, ok = v.parseIndices()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Subcomponent rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Subcomponent", token)
		panic(message)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Subcomponent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Subcomponent", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Subcomponent rule.
	ok = true
	v.remove(tokens)
	subcomponent = ast.SubcomponentClass().Subcomponent(
		composite,
		indices,
	)
	return
}

func (v *parser_) parseTarget() (
	target ast.TargetLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Target rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Target", token)
		panic(message)
	}

	// Found a single Target rule.
	ok = true
	v.remove(tokens)
	target = ast.TargetClass().Target(expression)
	return
}

func (v *parser_) parseTemplate() (
	template ast.TemplateLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Expression rule.
	var expression ast.ExpressionLike
	expression, token, ok = v.parseExpression()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Template rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Template", token)
		panic(message)
	}

	// Found a single Template rule.
	ok = true
	v.remove(tokens)
	template = ast.TemplateClass().Template(expression)
	return
}

func (v *parser_) parseThrowClause() (
	throwClause ast.ThrowClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "throw" delimiter.
	_, token, ok = v.parseDelimiter("throw")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ThrowClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ThrowClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Exception rule.
	var exception ast.ExceptionLike
	exception, token, ok = v.parseException()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single ThrowClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$ThrowClause", token)
		panic(message)
	}

	// Found a single ThrowClause rule.
	ok = true
	v.remove(tokens)
	throwClause = ast.ThrowClauseClass().ThrowClause(exception)
	return
}

func (v *parser_) parseTimes() (
	times ast.TimesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "*" delimiter.
	_, token, ok = v.parseDelimiter("*")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Times rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Times", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Times rule.
	ok = true
	v.remove(tokens)
	times = ast.TimesClass().Times()
	return
}

func (v *parser_) parseTimesEquals() (
	timesEquals ast.TimesEqualsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "*=" delimiter.
	_, token, ok = v.parseDelimiter("*=")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single TimesEquals rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$TimesEquals", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single TimesEquals rule.
	ok = true
	v.remove(tokens)
	timesEquals = ast.TimesEqualsClass().TimesEquals()
	return
}

func (v *parser_) parseValue() (
	value ast.ValueLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Value rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Value", token)
		panic(message)
	}

	// Found a single Value rule.
	ok = true
	v.remove(tokens)
	value = ast.ValueClass().Value(component)
	return
}

func (v *parser_) parseValues() (
	values ast.ValuesLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single InlineValues Values.
	var inlineValues ast.InlineValuesLike
	inlineValues, token, ok = v.parseInlineValues()
	if ok {
		// Found a single InlineValues Values.
		values = ast.ValuesClass().Values(inlineValues)
		return
	}

	// Attempt to parse a single MultilineValues Values.
	var multilineValues ast.MultilineValuesLike
	multilineValues, token, ok = v.parseMultilineValues()
	if ok {
		// Found a single MultilineValues Values.
		values = ast.ValuesClass().Values(multilineValues)
		return
	}

	// This is not a single Values rule.
	return
}

func (v *parser_) parseVariable() (
	variable ast.VariableLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Variable rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Variable", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Variable rule.
	ok = true
	v.remove(tokens)
	variable = ast.VariableClass().Variable(identifier)
	return
}

func (v *parser_) parseWhileClause() (
	whileClause ast.WhileClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "while" delimiter.
	_, token, ok = v.parseDelimiter("while")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WhileClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WhileClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Condition rule.
	var condition ast.ConditionLike
	condition, token, ok = v.parseCondition()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single WhileClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WhileClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WhileClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WhileClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single WhileClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WhileClause", token)
		panic(message)
	}

	// Found a single WhileClause rule.
	ok = true
	v.remove(tokens)
	whileClause = ast.WhileClauseClass().WhileClause(
		condition,
		procedure,
	)
	return
}

func (v *parser_) parseWithClause() (
	withClause ast.WithClauseLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "with" delimiter.
	_, token, ok = v.parseDelimiter("with")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "each" delimiter.
	_, token, ok = v.parseDelimiter("each")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Item rule.
	var item ast.ItemLike
	item, token, ok = v.parseItem()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single WithClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WithClause", token)
		panic(message)
	}

	// Attempt to parse a single "in" delimiter.
	_, token, ok = v.parseDelimiter("in")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Sequence rule.
	var sequence ast.SequenceLike
	sequence, token, ok = v.parseSequence()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single WithClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WithClause", token)
		panic(message)
	}

	// Attempt to parse a single "do" delimiter.
	_, token, ok = v.parseDelimiter("do")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single WithClause rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$WithClause", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Procedure rule.
	var procedure ast.ProcedureLike
	procedure, token, ok = v.parseProcedure()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single WithClause rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$WithClause", token)
		panic(message)
	}

	// Found a single WithClause rule.
	ok = true
	v.remove(tokens)
	withClause = ast.WithClauseClass().WithClause(
		item,
		sequence,
		procedure,
	)
	return
}

func (v *parser_) parseXOR() (
	xOR ast.XORLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "XOR" delimiter.
	_, token, ok = v.parseDelimiter("XOR")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single XOR rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$XOR", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single XOR rule.
	ok = true
	v.remove(tokens)
	xOR = ast.XORClass().XOR()
	return
}

func (v *parser_) parseDelimiter(
	expectedValue string,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single delimiter.
	value, token, ok = v.parseToken(DelimiterToken)
	if ok {
		if value == expectedValue {
			// Found the desired delimiter.
			return
		}
		v.next_.AddValue(token)
		ok = false
	}

	// This is not the desired delimiter.
	return
}

func (v *parser_) parseToken(
	tokenType TokenType,
) (
	value string,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a specific token type.
	var tokens = fra.List[TokenLike]()
	token = v.getNextToken()
	for token != nil {
		tokens.AppendValue(token)
		switch token.GetType() {
		case tokenType:
			// Found the desired token type.
			value = token.GetValue()
			ok = true
			return
		case SpaceToken, NewlineToken:
			// Ignore any unrequested whitespace.
			token = v.getNextToken()
		default:
			// This is not the desired token type.
			v.putBack(tokens)
			return
		}
	}

	// We are at the end-of-file marker.
	return
}

func (v *parser_) formatError(
	ruleName string,
	token TokenLike,
) string {
	// Format the error message.
	var message = fmt.Sprintf(
		"An unexpected token was received by the parser: %v\n",
		ScannerClass().FormatToken(token),
	)
	var line = token.GetLine()
	var lines = sts.Split(v.source_, "\n")

	// Append the source lines with the error in it.
	message += "\033[36m"
	for index := line - 3; index < line; index++ {
		if index > 1 {
			message += fmt.Sprintf("%04d: ", index) + string(lines[index-1]) + "\n"
		}
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	// Append an arrow pointing to the error.
	message += " \033[32m>>>─"
	var count uint
	for count < token.GetPosition() {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	// Append the following source line for context.
	if line < uint(len(lines)) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}
	message += "\033[0m\n"
	if uti.IsDefined(ruleName) {
		message += "Was expecting:\n"
		message += fmt.Sprintf(
			"  \033[32m%v: \033[33m%v\033[0m\n\n",
			ruleName,
			v.getDefinition(ruleName),
		)
	}
	return message
}

func (v *parser_) getDefinition(
	ruleName string,
) string {
	var syntax = parserClass().syntax_
	var definition = syntax.GetValue(ruleName)
	return definition
}

func (v *parser_) getNextToken() TokenLike {
	// Check for any read, but unprocessed tokens.
	if !v.next_.IsEmpty() {
		return v.next_.RemoveLast()
	}

	// Read a new token from the token stream.
	var token, ok = v.tokens_.RemoveFirst() // This will wait for a token.
	if !ok {
		// The token channel has been closed.
		return nil
	}

	// Check for an error token.
	if token.GetType() == ErrorToken {
		var message = v.formatError("", token)
		panic(message)
	}

	return token
}

func (v *parser_) putBack(
	tokens col.Sequential[TokenLike],
) {
	var iterator = tokens.GetIterator()
	for iterator.ToEnd(); iterator.HasPrevious(); {
		var token = iterator.GetPrevious()
		v.next_.AddValue(token)
	}
}

func (v *parser_) remove(
	tokens col.Sequential[TokenLike],
) {
	// NOTE: This method does nothing but must exist to satisfy the lint
	// check on the generated parser code.
}

// Instance Structure

type parser_ struct {
	// Declare the instance attributes.
	source_ string                   // The original source code.
	tokens_ col.QueueLike[TokenLike] // A queue of unread tokens from the scanner.
	next_   col.StackLike[TokenLike] // A stack of read, but unprocessed tokens.
}

// Class Structure

type parserClass_ struct {
	// Declare the class constants.
	syntax_ col.CatalogLike[string, string]
}

// Class Reference

func parserClass() *parserClass_ {
	return parserClassReference_
}

var parserClassReference_ = &parserClass_{
	// Initialize the class constants.
	syntax_: fra.CatalogFromMap[string, string](
		map[string]string{
			"$Document":  `comment Component`,
			"$Component": `Entity Context?`,
			"$Entity": `
  - Element
  - String
  - Range
  - Collection
  - Procedure`,
			"$Context": `"(" Parameters ")"`,
			"$Parameters": `
  - InlineParameters
  - MultilineParameters`,
			"$InlineParameters":    `Parameter AdditionalParameter*`,
			"$MultilineParameters": `AnnotatedParameter+`,
			"$AdditionalParameter": `"," Parameter`,
			"$AnnotatedParameter":  `Parameter note?`,
			"$Parameter":           `Label ":" Component`,
			"$Label":               `symbol`,
			"$Element": `
  - angle
  - boolean
  - duration
  - moment
  - number
  - pattern
  - percentage
  - probability
  - resource`,
			"$String": `
  - binary
  - bytecode
  - name
  - narrative
  - quote
  - symbol
  - tag
  - version`,
			"$Range": `LeftBracket Primitive ".." Primitive RightBracket`,
			"$LeftBracket": `
  - LeftSquare
  - LeftRound`,
			"$LeftSquare": `"["`,
			"$LeftRound":  `"("`,
			"$Primitive": `
  - Element
  - String`,
			"$RightBracket": `
  - RightSquare
  - RightRound`,
			"$RightSquare": `"]"`,
			"$RightRound":  `")"`,
			"$Collection":  `"[" Items? "]"`,
			"$Items": `
  - Associations
  - Values`,
			"$Associations": `
  - InlineAssociations
  - MultilineAssociations
  - NoAssociations`,
			"$InlineAssociations":    `Association AdditionalAssociation*`,
			"$MultilineAssociations": `AnnotatedAssociation+`,
			"$NoAssociations":        `":"`,
			"$AdditionalAssociation": `"," Association`,
			"$AnnotatedAssociation":  `Association note?`,
			"$Association":           `Key ":" Value`,
			"$Key":                   `Primitive`,
			"$Value":                 `Component`,
			"$Values": `
  - InlineValues
  - MultilineValues`,
			"$InlineValues":    `Value AdditionalValue*`,
			"$MultilineValues": `AnnotatedValue+`,
			"$AdditionalValue": `"," Value`,
			"$AnnotatedValue":  `Value note?`,
			"$Procedure":       `"{" Statements? "}"`,
			"$Statements": `
  - InlineStatements
  - MultilineStatements`,
			"$InlineStatements":    `Statement AdditionalStatement*`,
			"$MultilineStatements": `AnnotatedStatement+`,
			"$AdditionalStatement": `";" Statement`,
			"$AnnotatedStatement":  `Statement note?`,
			"$Statement":           `MainClause OnClause?`,
			"$MainClause": `
  - Flow
  - Induction
  - Messaging
  - Repository`,
			"$OnClause":     `"on" Failure MatchHandler+`,
			"$MatchHandler": `"matching" Template "do" Procedure`,
			"$Failure":      `symbol`,
			"$Template":     `Expression`,
			"$Flow": `
  - IfClause
  - SelectClause
  - WhileClause
  - WithClause
  - ContinueClause
  - BreakClause
  - ReturnClause
  - ThrowClause`,
			"$Induction": `
  - DoClause
  - LetClause`,
			"$Messaging": `
  - PostClause
  - RetrieveClause
  - AcceptClause
  - RejectClause
  - PublishClause`,
			"$Repository": `
  - CheckoutClause
  - SaveClause
  - DiscardClause
  - NotarizeClause`,
			"$IfClause":       `"if" Condition "do" Procedure`,
			"$Condition":      `Expression`,
			"$SelectClause":   `"select" Target MatchHandler+`,
			"$Target":         `Expression`,
			"$WhileClause":    `"while" Condition "do" Procedure`,
			"$WithClause":     `"with" "each" Item "in" Sequence "do" Procedure`,
			"$Item":           `symbol`,
			"$Sequence":       `Expression`,
			"$ContinueClause": `"continue" "loop"`,
			"$BreakClause":    `"break" "loop"`,
			"$ReturnClause":   `"return" Result`,
			"$Result":         `Expression`,
			"$ThrowClause":    `"throw" Exception`,
			"$Exception":      `Expression`,
			"$DoClause":       `"do" Operation`,
			"$Operation": `
  - Intrinsic
  - Invocation`,
			"$Intrinsic":  `Function "(" Arguments? ")"`,
			"$Function":   `identifier`,
			"$Invocation": `Target InvocationOperator Method "(" Arguments? ")"`,
			"$InvocationOperator": `
  - Dot
  - Arrow`,
			"$Dot":                `"."`,
			"$Arrow":              `"<-"`,
			"$Method":             `identifier`,
			"$Arguments":          `Argument AdditionalArgument*`,
			"$AdditionalArgument": `"," Argument`,
			"$Argument":           `Expression`,
			"$LetClause":          `"let" Recipient AssignmentOperator Expression`,
			"$AssignmentOperator": `
  - ColonEquals
  - DefaultEquals
  - PlusEquals
  - MinusEquals
  - TimesEquals
  - DivideEquals`,
			"$ColonEquals":   `":="`,
			"$DefaultEquals": `"?="`,
			"$PlusEquals":    `"+="`,
			"$MinusEquals":   `"-="`,
			"$TimesEquals":   `"*="`,
			"$DivideEquals":  `"/="`,
			"$Recipient": `
  - Label
  - Attribute`,
			"$Attribute":       `Variable "[" Indices "]"`,
			"$Variable":        `identifier`,
			"$Indices":         `Index AdditionalIndex*`,
			"$AdditionalIndex": `"," Index`,
			"$Index":           `Expression`,
			"$PostClause":      `"post" Message "to" Bag`,
			"$Message":         `Expression`,
			"$Bag":             `Expression`,
			"$RetrieveClause":  `"retrieve" Recipient "from" Bag`,
			"$AcceptClause":    `"accept" Message`,
			"$RejectClause":    `"reject" Message`,
			"$PublishClause":   `"publish" Event`,
			"$Event":           `Expression`,
			"$CheckoutClause":  `"checkout" Recipient AtLevel? "from" Citation`,
			"$AtLevel":         `"at" "level" Expression`,
			"$Citation":        `Expression`,
			"$SaveClause":      `"save" Draft "as" Citation`,
			"$Draft":           `Expression`,
			"$DiscardClause":   `"discard" Draft`,
			"$NotarizeClause":  `"notarize" Draft "as" Citation`,
			"$Expression": `
  - Component
  - Intrinsic
  - Variable
  - Precedence
  - Dereference
  - Invocation
  - Subcomponent
  - Chaining
  - Exponential
  - Inversion
  - Arithmetic
  - Amplitude
  - Comparison
  - Complement
  - Logical`,
			"$Precedence":   `"(" Expression ")"`,
			"$Dereference":  `"@" Expression`,
			"$Subcomponent": `Composite "[" Indices "]"`,
			"$Composite":    `Expression`,
			"$Chaining":     `Expression "&" Expression`,
			"$Exponential":  `Expression "^" Expression`,
			"$Inversion":    `InversionOperator Expression`,
			"$InversionOperator": `
  - Minus
  - Divide
  - Times`,
			"$Minus":      `"-"`,
			"$Divide":     `"/"`,
			"$Times":      `"*"`,
			"$Arithmetic": `Expression ArithmeticOperator Expression`,
			"$ArithmeticOperator": `
  - Times
  - Divide
  - Modulo
  - Plus
  - Minus`,
			"$Modulo":     `"//"`,
			"$Plus":       `"+"`,
			"$Amplitude":  `"|" Expression "|"`,
			"$Comparison": `Expression ComparisonOperator Expression`,
			"$ComparisonOperator": `
  - Less
  - Equals
  - More
  - Is
  - Matches`,
			"$Less":       `"<"`,
			"$Equals":     `"="`,
			"$More":       `">"`,
			"$Is":         `"IS"`,
			"$Matches":    `"MATCHES"`,
			"$Complement": `"NOT" Expression`,
			"$Logical":    `Expression LogicalOperator Expression`,
			"$LogicalOperator": `
  - And
  - SAN
  - IOR
  - XOR`,
			"$And": `"AND"`,
			"$SAN": `"SAN"`,
			"$IOR": `"IOR"`,
			"$XOR": `"XOR"`,
		},
	),
}
