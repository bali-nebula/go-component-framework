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

	// Attempt to parse a single bar token.
	var bar1 string
	bar1, token, ok = v.parseToken(BarToken)
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

	// Attempt to parse a single Numerical rule.
	var numerical ast.NumericalLike
	numerical, token, ok = v.parseNumerical()
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

	// Attempt to parse a single bar token.
	var bar2 string
	bar2, token, ok = v.parseToken(BarToken)
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
	amplitude = ast.AmplitudeClass().Amplitude(
		bar1,
		numerical,
		bar2,
	)
	return
}

func (v *parser_) parseAnnotatedAssociation() (
	annotatedAssociation ast.AnnotatedAssociationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single dash token.
	var dash string
	dash, token, ok = v.parseToken(DashToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AnnotatedAssociation rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AnnotatedAssociation", token)
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
		dash,
		association,
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

	// Attempt to parse a single dash token.
	var dash string
	dash, token, ok = v.parseToken(DashToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AnnotatedStatement rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AnnotatedStatement", token)
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
		dash,
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

	// Attempt to parse a single dash token.
	var dash string
	dash, token, ok = v.parseToken(DashToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single AnnotatedValue rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$AnnotatedValue", token)
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
		dash,
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

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Argument rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Argument", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Argument rule.
	ok = true
	v.remove(tokens)
	argument = ast.ArgumentClass().Argument(identifier)
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

func (v *parser_) parseAssign() (
	assign ast.AssignLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single colonEqual Assign.
	var colonEqual string
	colonEqual, token, ok = v.parseToken(ColonEqualToken)
	if ok {
		// Found a single colonEqual Assign.
		assign = ast.AssignClass().Assign(colonEqual)
		return
	}

	// Attempt to parse a single defaultEqual Assign.
	var defaultEqual string
	defaultEqual, token, ok = v.parseToken(DefaultEqualToken)
	if ok {
		// Found a single defaultEqual Assign.
		assign = ast.AssignClass().Assign(defaultEqual)
		return
	}

	// Attempt to parse a single plusEqual Assign.
	var plusEqual string
	plusEqual, token, ok = v.parseToken(PlusEqualToken)
	if ok {
		// Found a single plusEqual Assign.
		assign = ast.AssignClass().Assign(plusEqual)
		return
	}

	// Attempt to parse a single dashEqual Assign.
	var dashEqual string
	dashEqual, token, ok = v.parseToken(DashEqualToken)
	if ok {
		// Found a single dashEqual Assign.
		assign = ast.AssignClass().Assign(dashEqual)
		return
	}

	// Attempt to parse a single starEqual Assign.
	var starEqual string
	starEqual, token, ok = v.parseToken(StarEqualToken)
	if ok {
		// Found a single starEqual Assign.
		assign = ast.AssignClass().Assign(starEqual)
		return
	}

	// Attempt to parse a single slashEqual Assign.
	var slashEqual string
	slashEqual, token, ok = v.parseToken(SlashEqualToken)
	if ok {
		// Found a single slashEqual Assign.
		assign = ast.AssignClass().Assign(slashEqual)
		return
	}

	// This is not a single Assign rule.
	return
}

func (v *parser_) parseAssociation() (
	association ast.AssociationLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single symbol token.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
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

	// Attempt to parse a single colon token.
	var colon string
	colon, token, ok = v.parseToken(ColonToken)
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
		symbol,
		colon,
		value,
	)
	return
}

func (v *parser_) parseAssociations() (
	associations ast.AssociationsLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single MultilineAssociations Associations.
	var multilineAssociations ast.MultilineAssociationsLike
	multilineAssociations, token, ok = v.parseMultilineAssociations()
	if ok {
		// Found a single MultilineAssociations Associations.
		associations = ast.AssociationsClass().Associations(multilineAssociations)
		return
	}

	// Attempt to parse a single InlineAssociations Associations.
	var inlineAssociations ast.InlineAssociationsLike
	inlineAssociations, token, ok = v.parseInlineAssociations()
	if ok {
		// Found a single InlineAssociations Associations.
		associations = ast.AssociationsClass().Associations(inlineAssociations)
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

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
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

	// Attempt to parse a single Subcomponent rule.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
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

	// Found a single Attribute rule.
	ok = true
	v.remove(tokens)
	attribute = ast.AttributeClass().Attribute(
		identifier,
		subcomponent,
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

func (v *parser_) parseComplement() (
	complement ast.ComplementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single not token.
	var not string
	not, token, ok = v.parseToken(NotToken)
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

	// Attempt to parse a single Logical rule.
	var logical ast.LogicalLike
	logical, token, ok = v.parseLogical()
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
	complement = ast.ComplementClass().Complement(
		not,
		logical,
	)
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

	// Attempt to parse an optional Parameters rule.
	var optionalParameters ast.ParametersLike
	optionalParameters, _, ok = v.parseParameters()
	if ok {
		// No additional put backs allowed at this point.
		tokens = nil
	}

	// Found a single Component rule.
	ok = true
	v.remove(tokens)
	component = ast.ComponentClass().Component(
		entity,
		optionalParameters,
	)
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

func (v *parser_) parseDereference() (
	dereference ast.DereferenceLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single snail token.
	var snail string
	snail, token, ok = v.parseToken(SnailToken)
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

	// Attempt to parse a single Indirect rule.
	var indirect ast.IndirectLike
	indirect, token, ok = v.parseIndirect()
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
	dereference = ast.DereferenceClass().Dereference(
		snail,
		indirect,
	)
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

	// Attempt to parse a single Invocation rule.
	var invocation ast.InvocationLike
	invocation, token, ok = v.parseInvocation()
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
	doClause = ast.DoClauseClass().DoClause(invocation)
	return
}

func (v *parser_) parseDocument() (
	document ast.DocumentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse an optional comment token.
	var optionalComment string
	optionalComment, token, ok = v.parseToken(CommentToken)
	if ok && uti.IsDefined(tokens) {
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
		optionalComment,
		component,
	)
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

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Unary rule.
	var unary ast.UnaryLike
	unary, token, ok = v.parseUnary()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Expression rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Expression", token)
		panic(message)
	}

	// Attempt to parse multiple Operation rules.
	var operations = fra.List[ast.OperationLike]()
operationsLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var operation ast.OperationLike
		operation, token, ok = v.parseOperation()
		if !ok {
			switch {
			case count >= 0:
				break operationsLoop
			case uti.IsDefined(tokens):
				// This is not multiple Operation rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Expression", token)
				message += "0 or more Operation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		operations.AppendValue(operation)
	}

	// Found a single Expression rule.
	ok = true
	v.remove(tokens)
	expression = ast.ExpressionClass().Expression(
		unary,
		operations,
	)
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

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
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
	function = ast.FunctionClass().Function(
		identifier,
		optionalArguments,
	)
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

func (v *parser_) parseInclusion() (
	inclusion ast.InclusionLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Left Inclusion.
	var left ast.LeftLike
	left, token, ok = v.parseLeft()
	if ok {
		// Found a single Left Inclusion.
		inclusion = ast.InclusionClass().Inclusion(left)
		return
	}

	// Attempt to parse a single Right Inclusion.
	var right ast.RightLike
	right, token, ok = v.parseRight()
	if ok {
		// Found a single Right Inclusion.
		inclusion = ast.InclusionClass().Inclusion(right)
		return
	}

	// This is not a single Inclusion rule.
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

func (v *parser_) parseIndirect() (
	indirect ast.IndirectLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Indirect.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Indirect.
		indirect = ast.IndirectClass().Indirect(component)
		return
	}

	// Attempt to parse a single Dereference Indirect.
	var dereference ast.DereferenceLike
	dereference, token, ok = v.parseDereference()
	if ok {
		// Found a single Dereference Indirect.
		indirect = ast.IndirectClass().Indirect(dereference)
		return
	}

	// Attempt to parse a single Function Indirect.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Indirect.
		indirect = ast.IndirectClass().Indirect(function)
		return
	}

	// Attempt to parse a single Method Indirect.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Indirect.
		indirect = ast.IndirectClass().Indirect(method)
		return
	}

	// Attempt to parse a single Attribute Indirect.
	var attribute ast.AttributeLike
	attribute, token, ok = v.parseAttribute()
	if ok {
		// Found a single Attribute Indirect.
		indirect = ast.IndirectClass().Indirect(attribute)
		return
	}

	// Attempt to parse a single Variable Indirect.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Indirect.
		indirect = ast.IndirectClass().Indirect(variable)
		return
	}

	// This is not a single Indirect rule.
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

func (v *parser_) parseInverse() (
	inverse ast.InverseLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single dash Inverse.
	var dash string
	dash, token, ok = v.parseToken(DashToken)
	if ok {
		// Found a single dash Inverse.
		inverse = ast.InverseClass().Inverse(dash)
		return
	}

	// Attempt to parse a single slash Inverse.
	var slash string
	slash, token, ok = v.parseToken(SlashToken)
	if ok {
		// Found a single slash Inverse.
		inverse = ast.InverseClass().Inverse(slash)
		return
	}

	// Attempt to parse a single star Inverse.
	var star string
	star, token, ok = v.parseToken(StarToken)
	if ok {
		// Found a single star Inverse.
		inverse = ast.InverseClass().Inverse(star)
		return
	}

	// This is not a single Inverse rule.
	return
}

func (v *parser_) parseInversion() (
	inversion ast.InversionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Inverse rule.
	var inverse ast.InverseLike
	inverse, token, ok = v.parseInverse()
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

	// Attempt to parse a single Numerical rule.
	var numerical ast.NumericalLike
	numerical, token, ok = v.parseNumerical()
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
		inverse,
		numerical,
	)
	return
}

func (v *parser_) parseInvocation() (
	invocation ast.InvocationLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Function Invocation.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Invocation.
		invocation = ast.InvocationClass().Invocation(function)
		return
	}

	// Attempt to parse a single Method Invocation.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Invocation.
		invocation = ast.InvocationClass().Invocation(method)
		return
	}

	// This is not a single Invocation rule.
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
	// Attempt to parse a single Range Items.
	var range_ ast.RangeLike
	range_, token, ok = v.parseRange()
	if ok {
		// Found a single Range Items.
		items = ast.ItemsClass().Items(range_)
		return
	}

	// Attempt to parse a single NoAssociations Items.
	var noAssociations ast.NoAssociationsLike
	noAssociations, token, ok = v.parseNoAssociations()
	if ok {
		// Found a single NoAssociations Items.
		items = ast.ItemsClass().Items(noAssociations)
		return
	}

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

func (v *parser_) parseLeft() (
	left ast.LeftLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Left rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Left", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Left rule.
	ok = true
	v.remove(tokens)
	left = ast.LeftClass().Left()
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

	// Attempt to parse a single Assign rule.
	var assign ast.AssignLike
	assign, token, ok = v.parseAssign()
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
		assign,
		expression,
	)
	return
}

func (v *parser_) parseLogical() (
	logical ast.LogicalLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Logical.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Logical.
		logical = ast.LogicalClass().Logical(component)
		return
	}

	// Attempt to parse a single Precedence Logical.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Logical.
		logical = ast.LogicalClass().Logical(precedence)
		return
	}

	// Attempt to parse a single Dereference Logical.
	var dereference ast.DereferenceLike
	dereference, token, ok = v.parseDereference()
	if ok {
		// Found a single Dereference Logical.
		logical = ast.LogicalClass().Logical(dereference)
		return
	}

	// Attempt to parse a single Complement Logical.
	var complement ast.ComplementLike
	complement, token, ok = v.parseComplement()
	if ok {
		// Found a single Complement Logical.
		logical = ast.LogicalClass().Logical(complement)
		return
	}

	// Attempt to parse a single Function Logical.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Logical.
		logical = ast.LogicalClass().Logical(function)
		return
	}

	// Attempt to parse a single Method Logical.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Logical.
		logical = ast.LogicalClass().Logical(method)
		return
	}

	// Attempt to parse a single Attribute Logical.
	var attribute ast.AttributeLike
	attribute, token, ok = v.parseAttribute()
	if ok {
		// Found a single Attribute Logical.
		logical = ast.LogicalClass().Logical(attribute)
		return
	}

	// Attempt to parse a single Variable Logical.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Logical.
		logical = ast.LogicalClass().Logical(variable)
		return
	}

	// This is not a single Logical rule.
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
	var identifier1 string
	identifier1, token, ok = v.parseToken(IdentifierToken)
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

	// Attempt to parse a single Threading rule.
	var threading ast.ThreadingLike
	threading, token, ok = v.parseThreading()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Method rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Method", token)
		panic(message)
	}

	// Attempt to parse a single identifier token.
	var identifier2 string
	identifier2, token, ok = v.parseToken(IdentifierToken)
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

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
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
	method = ast.MethodClass().Method(
		identifier1,
		threading,
		identifier2,
		optionalArguments,
	)
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

	// Attempt to parse a single colon token.
	var colon string
	colon, token, ok = v.parseToken(ColonToken)
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
	noAssociations = ast.NoAssociationsClass().NoAssociations(colon)
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

func (v *parser_) parseNumerical() (
	numerical ast.NumericalLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Numerical.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Numerical.
		numerical = ast.NumericalClass().Numerical(component)
		return
	}

	// Attempt to parse a single Precedence Numerical.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Numerical.
		numerical = ast.NumericalClass().Numerical(precedence)
		return
	}

	// Attempt to parse a single Dereference Numerical.
	var dereference ast.DereferenceLike
	dereference, token, ok = v.parseDereference()
	if ok {
		// Found a single Dereference Numerical.
		numerical = ast.NumericalClass().Numerical(dereference)
		return
	}

	// Attempt to parse a single Inversion Numerical.
	var inversion ast.InversionLike
	inversion, token, ok = v.parseInversion()
	if ok {
		// Found a single Inversion Numerical.
		numerical = ast.NumericalClass().Numerical(inversion)
		return
	}

	// Attempt to parse a single Amplitude Numerical.
	var amplitude ast.AmplitudeLike
	amplitude, token, ok = v.parseAmplitude()
	if ok {
		// Found a single Amplitude Numerical.
		numerical = ast.NumericalClass().Numerical(amplitude)
		return
	}

	// Attempt to parse a single Function Numerical.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Numerical.
		numerical = ast.NumericalClass().Numerical(function)
		return
	}

	// Attempt to parse a single Method Numerical.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Numerical.
		numerical = ast.NumericalClass().Numerical(method)
		return
	}

	// Attempt to parse a single Attribute Numerical.
	var attribute ast.AttributeLike
	attribute, token, ok = v.parseAttribute()
	if ok {
		// Found a single Attribute Numerical.
		numerical = ast.NumericalClass().Numerical(attribute)
		return
	}

	// Attempt to parse a single Variable Numerical.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Numerical.
		numerical = ast.NumericalClass().Numerical(variable)
		return
	}

	// This is not a single Numerical rule.
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
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Operator rule.
	var operator ast.OperatorLike
	operator, token, ok = v.parseOperator()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Operation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Operation", token)
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
		// This is not a single Operation rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Operation", token)
		panic(message)
	}

	// Found a single Operation rule.
	ok = true
	v.remove(tokens)
	operation = ast.OperationClass().Operation(
		operator,
		expression,
	)
	return
}

func (v *parser_) parseOperator() (
	operator ast.OperatorLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single plus Operator.
	var plus string
	plus, token, ok = v.parseToken(PlusToken)
	if ok {
		// Found a single plus Operator.
		operator = ast.OperatorClass().Operator(plus)
		return
	}

	// Attempt to parse a single dash Operator.
	var dash string
	dash, token, ok = v.parseToken(DashToken)
	if ok {
		// Found a single dash Operator.
		operator = ast.OperatorClass().Operator(dash)
		return
	}

	// Attempt to parse a single star Operator.
	var star string
	star, token, ok = v.parseToken(StarToken)
	if ok {
		// Found a single star Operator.
		operator = ast.OperatorClass().Operator(star)
		return
	}

	// Attempt to parse a single slash Operator.
	var slash string
	slash, token, ok = v.parseToken(SlashToken)
	if ok {
		// Found a single slash Operator.
		operator = ast.OperatorClass().Operator(slash)
		return
	}

	// Attempt to parse a single double Operator.
	var double string
	double, token, ok = v.parseToken(DoubleToken)
	if ok {
		// Found a single double Operator.
		operator = ast.OperatorClass().Operator(double)
		return
	}

	// Attempt to parse a single caret Operator.
	var caret string
	caret, token, ok = v.parseToken(CaretToken)
	if ok {
		// Found a single caret Operator.
		operator = ast.OperatorClass().Operator(caret)
		return
	}

	// Attempt to parse a single less Operator.
	var less string
	less, token, ok = v.parseToken(LessToken)
	if ok {
		// Found a single less Operator.
		operator = ast.OperatorClass().Operator(less)
		return
	}

	// Attempt to parse a single equal Operator.
	var equal string
	equal, token, ok = v.parseToken(EqualToken)
	if ok {
		// Found a single equal Operator.
		operator = ast.OperatorClass().Operator(equal)
		return
	}

	// Attempt to parse a single more Operator.
	var more string
	more, token, ok = v.parseToken(MoreToken)
	if ok {
		// Found a single more Operator.
		operator = ast.OperatorClass().Operator(more)
		return
	}

	// Attempt to parse a single is Operator.
	var is string
	is, token, ok = v.parseToken(IsToken)
	if ok {
		// Found a single is Operator.
		operator = ast.OperatorClass().Operator(is)
		return
	}

	// Attempt to parse a single matches Operator.
	var matches string
	matches, token, ok = v.parseToken(MatchesToken)
	if ok {
		// Found a single matches Operator.
		operator = ast.OperatorClass().Operator(matches)
		return
	}

	// Attempt to parse a single and Operator.
	var and string
	and, token, ok = v.parseToken(AndToken)
	if ok {
		// Found a single and Operator.
		operator = ast.OperatorClass().Operator(and)
		return
	}

	// Attempt to parse a single san Operator.
	var san string
	san, token, ok = v.parseToken(SanToken)
	if ok {
		// Found a single san Operator.
		operator = ast.OperatorClass().Operator(san)
		return
	}

	// Attempt to parse a single ior Operator.
	var ior string
	ior, token, ok = v.parseToken(IorToken)
	if ok {
		// Found a single ior Operator.
		operator = ast.OperatorClass().Operator(ior)
		return
	}

	// Attempt to parse a single xor Operator.
	var xor string
	xor, token, ok = v.parseToken(XorToken)
	if ok {
		// Found a single xor Operator.
		operator = ast.OperatorClass().Operator(xor)
		return
	}

	// Attempt to parse a single ampersand Operator.
	var ampersand string
	ampersand, token, ok = v.parseToken(AmpersandToken)
	if ok {
		// Found a single ampersand Operator.
		operator = ast.OperatorClass().Operator(ampersand)
		return
	}

	// This is not a single Operator rule.
	return
}

func (v *parser_) parseParameters() (
	parameters ast.ParametersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single Associations rule.
	var associations ast.AssociationsLike
	associations, token, ok = v.parseAssociations()
	switch {
	case ok:
		// No additional put backs allowed at this point.
		tokens = nil
	case uti.IsDefined(tokens):
		// This is not a single Parameters rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Parameters", token)
		panic(message)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Parameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Parameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Parameters rule.
	ok = true
	v.remove(tokens)
	parameters = ast.ParametersClass().Parameters(associations)
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

	// Attempt to parse a single Inclusion rule.
	var inclusion1 ast.InclusionLike
	inclusion1, token, ok = v.parseInclusion()
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

	// Attempt to parse a single Inclusion rule.
	var inclusion2 ast.InclusionLike
	inclusion2, token, ok = v.parseInclusion()
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
		inclusion1,
		primitive1,
		primitive2,
		inclusion2,
	)
	return
}

func (v *parser_) parseRecipient() (
	recipient ast.RecipientLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single symbol Recipient.
	var symbol string
	symbol, token, ok = v.parseToken(SymbolToken)
	if ok {
		// Found a single symbol Recipient.
		recipient = ast.RecipientClass().Recipient(symbol)
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

func (v *parser_) parseRight() (
	right ast.RightLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Right rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Right", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Right rule.
	ok = true
	v.remove(tokens)
	right = ast.RightClass().Right()
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
	// Attempt to parse a single MultilineStatements Statements.
	var multilineStatements ast.MultilineStatementsLike
	multilineStatements, token, ok = v.parseMultilineStatements()
	if ok {
		// Found a single MultilineStatements Statements.
		statements = ast.StatementsClass().Statements(multilineStatements)
		return
	}

	// Attempt to parse a single InlineStatements Statements.
	var inlineStatements ast.InlineStatementsLike
	inlineStatements, token, ok = v.parseInlineStatements()
	if ok {
		// Found a single InlineStatements Statements.
		statements = ast.StatementsClass().Statements(inlineStatements)
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
	subcomponent = ast.SubcomponentClass().Subcomponent(indices)
	return
}

func (v *parser_) parseTarget() (
	target ast.TargetLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Function Target.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Target.
		target = ast.TargetClass().Target(function)
		return
	}

	// Attempt to parse a single Method Target.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Target.
		target = ast.TargetClass().Target(method)
		return
	}

	// Attempt to parse a single Attribute Target.
	var attribute ast.AttributeLike
	attribute, token, ok = v.parseAttribute()
	if ok {
		// Found a single Attribute Target.
		target = ast.TargetClass().Target(attribute)
		return
	}

	// Attempt to parse a single Variable Target.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Target.
		target = ast.TargetClass().Target(variable)
		return
	}

	// This is not a single Target rule.
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

func (v *parser_) parseThreading() (
	threading ast.ThreadingLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single dot Threading.
	var dot string
	dot, token, ok = v.parseToken(DotToken)
	if ok {
		// Found a single dot Threading.
		threading = ast.ThreadingClass().Threading(dot)
		return
	}

	// Attempt to parse a single arrow Threading.
	var arrow string
	arrow, token, ok = v.parseToken(ArrowToken)
	if ok {
		// Found a single arrow Threading.
		threading = ast.ThreadingClass().Threading(arrow)
		return
	}

	// This is not a single Threading rule.
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

func (v *parser_) parseUnary() (
	unary ast.UnaryLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Unary.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Unary.
		unary = ast.UnaryClass().Unary(component)
		return
	}

	// Attempt to parse a single Precedence Unary.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Unary.
		unary = ast.UnaryClass().Unary(precedence)
		return
	}

	// Attempt to parse a single Dereference Unary.
	var dereference ast.DereferenceLike
	dereference, token, ok = v.parseDereference()
	if ok {
		// Found a single Dereference Unary.
		unary = ast.UnaryClass().Unary(dereference)
		return
	}

	// Attempt to parse a single Complement Unary.
	var complement ast.ComplementLike
	complement, token, ok = v.parseComplement()
	if ok {
		// Found a single Complement Unary.
		unary = ast.UnaryClass().Unary(complement)
		return
	}

	// Attempt to parse a single Inversion Unary.
	var inversion ast.InversionLike
	inversion, token, ok = v.parseInversion()
	if ok {
		// Found a single Inversion Unary.
		unary = ast.UnaryClass().Unary(inversion)
		return
	}

	// Attempt to parse a single Amplitude Unary.
	var amplitude ast.AmplitudeLike
	amplitude, token, ok = v.parseAmplitude()
	if ok {
		// Found a single Amplitude Unary.
		unary = ast.UnaryClass().Unary(amplitude)
		return
	}

	// Attempt to parse a single Function Unary.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Unary.
		unary = ast.UnaryClass().Unary(function)
		return
	}

	// Attempt to parse a single Method Unary.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Unary.
		unary = ast.UnaryClass().Unary(method)
		return
	}

	// Attempt to parse a single Attribute Unary.
	var attribute ast.AttributeLike
	attribute, token, ok = v.parseAttribute()
	if ok {
		// Found a single Attribute Unary.
		unary = ast.UnaryClass().Unary(attribute)
		return
	}

	// Attempt to parse a single Variable Unary.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Unary.
		unary = ast.UnaryClass().Unary(variable)
		return
	}

	// This is not a single Unary rule.
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
	// Attempt to parse a single MultilineValues Values.
	var multilineValues ast.MultilineValuesLike
	multilineValues, token, ok = v.parseMultilineValues()
	if ok {
		// Found a single MultilineValues Values.
		values = ast.ValuesClass().Values(multilineValues)
		return
	}

	// Attempt to parse a single InlineValues Values.
	var inlineValues ast.InlineValuesLike
	inlineValues, token, ok = v.parseInlineValues()
	if ok {
		// Found a single InlineValues Values.
		values = ast.ValuesClass().Values(inlineValues)
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
			"$Document":  `comment? Component`,
			"$Component": `Entity Parameters?`,
			"$Entity": `
  - Element
  - String
  - Collection
  - Procedure`,
			"$Parameters": `"(" Associations ")"`,
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
			"$Collection": `"[" Items? "]"`,
			"$Items": `
  - Range
  - NoAssociations
  - Associations
  - Values`,
			"$Range": `Inclusion Primitive ".." Primitive Inclusion`,
			"$Inclusion": `
  - Left
  - Right`,
			"$Left":  `"("`,
			"$Right": `")"`,
			"$Primitive": `
  - Element
  - String`,
			"$NoAssociations": `colon`,
			"$Associations": `
  - MultilineAssociations
  - InlineAssociations`,
			"$MultilineAssociations": `AnnotatedAssociation+`,
			"$AnnotatedAssociation":  `dash Association note?`,
			"$InlineAssociations":    `Association AdditionalAssociation*`,
			"$AdditionalAssociation": `"," Association`,
			"$Association":           `symbol colon Value`,
			"$Value":                 `Component`,
			"$Values": `
  - MultilineValues
  - InlineValues`,
			"$MultilineValues": `AnnotatedValue+`,
			"$AnnotatedValue":  `dash Value note?`,
			"$InlineValues":    `Value AdditionalValue*`,
			"$AdditionalValue": `"," Value`,
			"$Procedure":       `"{" Statements? "}"`,
			"$Statements": `
  - MultilineStatements
  - InlineStatements`,
			"$MultilineStatements": `AnnotatedStatement+`,
			"$AnnotatedStatement":  `dash Statement note?`,
			"$InlineStatements":    `Statement AdditionalStatement*`,
			"$AdditionalStatement": `";" Statement`,
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
			"$IfClause":     `"if" Condition "do" Procedure`,
			"$Condition":    `Expression`,
			"$SelectClause": `"select" Target MatchHandler+`,
			"$Target": `
  - Function
  - Method
  - Attribute
  - Variable  ! Must be last since the others also begin with an identifier.`,
			"$Function":           `identifier "(" Arguments? ")"`,
			"$Arguments":          `Argument AdditionalArgument*`,
			"$AdditionalArgument": `"," Argument`,
			"$Argument":           `identifier`,
			"$Method":             `identifier Threading identifier "(" Arguments? ")"`,
			"$Threading": `
  - dot
  - arrow`,
			"$Attribute":       `identifier Subcomponent`,
			"$Subcomponent":    `"[" Indices "]"`,
			"$Indices":         `Index AdditionalIndex*`,
			"$AdditionalIndex": `"," Index`,
			"$Index":           `Expression`,
			"$Variable":        `identifier`,
			"$WhileClause":     `"while" Condition "do" Procedure`,
			"$WithClause":      `"with" "each" Item "in" Sequence "do" Procedure`,
			"$Item":            `symbol`,
			"$Sequence":        `Expression`,
			"$ContinueClause":  `"continue" "loop"`,
			"$BreakClause":     `"break" "loop"`,
			"$ReturnClause":    `"return" Result`,
			"$Result":          `Expression`,
			"$ThrowClause":     `"throw" Exception`,
			"$Exception":       `Expression`,
			"$DoClause":        `"do" Invocation`,
			"$Invocation": `
  - Function
  - Method`,
			"$LetClause": `"let" Recipient Assign Expression`,
			"$Assign": `
  - colonEqual
  - defaultEqual
  - plusEqual
  - dashEqual
  - starEqual
  - slashEqual`,
			"$Recipient": `
  - symbol
  - Attribute`,
			"$PostClause":     `"post" Message "to" Bag`,
			"$Message":        `Expression`,
			"$Bag":            `Expression`,
			"$RetrieveClause": `"retrieve" Recipient "from" Bag`,
			"$AcceptClause":   `"accept" Message`,
			"$RejectClause":   `"reject" Message`,
			"$PublishClause":  `"publish" Event`,
			"$Event":          `Expression`,
			"$CheckoutClause": `"checkout" Recipient AtLevel? "from" Citation`,
			"$AtLevel":        `"at" "level" Expression`,
			"$Citation":       `Expression`,
			"$SaveClause":     `"save" Draft "as" Citation`,
			"$Draft":          `Expression`,
			"$DiscardClause":  `"discard" Draft`,
			"$NotarizeClause": `"notarize" Draft "as" Citation`,
			"$Expression":     `Unary Operation*`,
			"$Unary": `
  - Component
  - Precedence
  - Dereference
  - Complement
  - Inversion
  - Amplitude
  - Function
  - Method
  - Attribute
  - Variable  ! Must be last since the others also begin with an identifier.`,
			"$Operation": `Operator Expression`,
			"$Operator": `
  - plus
  - dash
  - star
  - slash
  - double
  - caret
  - less
  - equal
  - more
  - is
  - matches
  - and
  - san
  - ior
  - xor
  - ampersand`,
			"$Precedence":  `"(" Expression ")"`,
			"$Dereference": `snail Indirect`,
			"$Indirect": `
  - Component
  - Dereference
  - Function
  - Method
  - Attribute
  - Variable  ! Must be last since the others also begin with an identifier.`,
			"$Complement": `not Logical`,
			"$Logical": `
  - Component
  - Precedence
  - Dereference
  - Complement
  - Function
  - Method
  - Attribute
  - Variable  ! Must be last since the others also begin with an identifier.`,
			"$Inversion": `Inverse Numerical`,
			"$Inverse": `
  - dash
  - slash
  - star`,
			"$Numerical": `
  - Component
  - Precedence
  - Dereference
  - Inversion
  - Amplitude
  - Function
  - Method
  - Attribute
  - Variable  ! Must be last since the others also begin with an identifier.`,
			"$Amplitude": `bar Numerical bar`,
		},
	),
}
