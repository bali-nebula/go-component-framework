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

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
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
	additionalValue = ast.AdditionalValueClass().AdditionalValue(component)
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

func (v *parser_) parseAnnotatedStatement() (
	annotatedStatement ast.AnnotatedStatementLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single EmptyLine AnnotatedStatement.
	var emptyLine ast.EmptyLineLike
	emptyLine, token, ok = v.parseEmptyLine()
	if ok {
		// Found a single EmptyLine AnnotatedStatement.
		annotatedStatement = ast.AnnotatedStatementClass().AnnotatedStatement(emptyLine)
		return
	}

	// Attempt to parse a single AnnotationLine AnnotatedStatement.
	var annotationLine ast.AnnotationLineLike
	annotationLine, token, ok = v.parseAnnotationLine()
	if ok {
		// Found a single AnnotationLine AnnotatedStatement.
		annotatedStatement = ast.AnnotatedStatementClass().AnnotatedStatement(annotationLine)
		return
	}

	// Attempt to parse a single StatementLine AnnotatedStatement.
	var statementLine ast.StatementLineLike
	statementLine, token, ok = v.parseStatementLine()
	if ok {
		// Found a single StatementLine AnnotatedStatement.
		annotatedStatement = ast.AnnotatedStatementClass().AnnotatedStatement(statementLine)
		return
	}

	// This is not a single AnnotatedStatement rule.
	return
}

func (v *parser_) parseAnnotatedValue() (
	annotatedValue ast.AnnotatedValueLike,
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
		component,
		optionalNote,
	)
	return
}

func (v *parser_) parseAnnotationLine() (
	annotationLine ast.AnnotationLineLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single note AnnotationLine.
	var note string
	note, token, ok = v.parseToken(NoteToken)
	if ok {
		// Found a single note AnnotationLine.
		annotationLine = ast.AnnotationLineClass().AnnotationLine(note)
		return
	}

	// Attempt to parse a single comment AnnotationLine.
	var comment string
	comment, token, ok = v.parseToken(CommentToken)
	if ok {
		// Found a single comment AnnotationLine.
		annotationLine = ast.AnnotationLineClass().AnnotationLine(comment)
		return
	}

	// This is not a single AnnotationLine rule.
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

	// Attempt to parse a single Component rule.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
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
		component,
	)
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
	// Attempt to parse a single MultilineAttributes Collection.
	var multilineAttributes ast.MultilineAttributesLike
	multilineAttributes, token, ok = v.parseMultilineAttributes()
	if ok {
		// Found a single MultilineAttributes Collection.
		collection = ast.CollectionClass().Collection(multilineAttributes)
		return
	}

	// Attempt to parse a single MultilineValues Collection.
	var multilineValues ast.MultilineValuesLike
	multilineValues, token, ok = v.parseMultilineValues()
	if ok {
		// Found a single MultilineValues Collection.
		collection = ast.CollectionClass().Collection(multilineValues)
		return
	}

	// Attempt to parse a single NoAttributes Collection.
	var noAttributes ast.NoAttributesLike
	noAttributes, token, ok = v.parseNoAttributes()
	if ok {
		// Found a single NoAttributes Collection.
		collection = ast.CollectionClass().Collection(noAttributes)
		return
	}

	// Attempt to parse a single NoValues Collection.
	var noValues ast.NoValuesLike
	noValues, token, ok = v.parseNoValues()
	if ok {
		// Found a single NoValues Collection.
		collection = ast.CollectionClass().Collection(noValues)
		return
	}

	// Attempt to parse a single InclusiveAngles Collection.
	var inclusiveAngles ast.InclusiveAnglesLike
	inclusiveAngles, token, ok = v.parseInclusiveAngles()
	if ok {
		// Found a single InclusiveAngles Collection.
		collection = ast.CollectionClass().Collection(inclusiveAngles)
		return
	}

	// Attempt to parse a single InExclusiveAngles Collection.
	var inExclusiveAngles ast.InExclusiveAnglesLike
	inExclusiveAngles, token, ok = v.parseInExclusiveAngles()
	if ok {
		// Found a single InExclusiveAngles Collection.
		collection = ast.CollectionClass().Collection(inExclusiveAngles)
		return
	}

	// Attempt to parse a single ExInclusiveAngles Collection.
	var exInclusiveAngles ast.ExInclusiveAnglesLike
	exInclusiveAngles, token, ok = v.parseExInclusiveAngles()
	if ok {
		// Found a single ExInclusiveAngles Collection.
		collection = ast.CollectionClass().Collection(exInclusiveAngles)
		return
	}

	// Attempt to parse a single ExclusiveAngles Collection.
	var exclusiveAngles ast.ExclusiveAnglesLike
	exclusiveAngles, token, ok = v.parseExclusiveAngles()
	if ok {
		// Found a single ExclusiveAngles Collection.
		collection = ast.CollectionClass().Collection(exclusiveAngles)
		return
	}

	// Attempt to parse a single InclusiveCitations Collection.
	var inclusiveCitations ast.InclusiveCitationsLike
	inclusiveCitations, token, ok = v.parseInclusiveCitations()
	if ok {
		// Found a single InclusiveCitations Collection.
		collection = ast.CollectionClass().Collection(inclusiveCitations)
		return
	}

	// Attempt to parse a single InExclusiveCitations Collection.
	var inExclusiveCitations ast.InExclusiveCitationsLike
	inExclusiveCitations, token, ok = v.parseInExclusiveCitations()
	if ok {
		// Found a single InExclusiveCitations Collection.
		collection = ast.CollectionClass().Collection(inExclusiveCitations)
		return
	}

	// Attempt to parse a single ExInclusiveCitations Collection.
	var exInclusiveCitations ast.ExInclusiveCitationsLike
	exInclusiveCitations, token, ok = v.parseExInclusiveCitations()
	if ok {
		// Found a single ExInclusiveCitations Collection.
		collection = ast.CollectionClass().Collection(exInclusiveCitations)
		return
	}

	// Attempt to parse a single ExclusiveCitations Collection.
	var exclusiveCitations ast.ExclusiveCitationsLike
	exclusiveCitations, token, ok = v.parseExclusiveCitations()
	if ok {
		// Found a single ExclusiveCitations Collection.
		collection = ast.CollectionClass().Collection(exclusiveCitations)
		return
	}

	// Attempt to parse a single InclusiveDurations Collection.
	var inclusiveDurations ast.InclusiveDurationsLike
	inclusiveDurations, token, ok = v.parseInclusiveDurations()
	if ok {
		// Found a single InclusiveDurations Collection.
		collection = ast.CollectionClass().Collection(inclusiveDurations)
		return
	}

	// Attempt to parse a single InExclusiveDurations Collection.
	var inExclusiveDurations ast.InExclusiveDurationsLike
	inExclusiveDurations, token, ok = v.parseInExclusiveDurations()
	if ok {
		// Found a single InExclusiveDurations Collection.
		collection = ast.CollectionClass().Collection(inExclusiveDurations)
		return
	}

	// Attempt to parse a single ExInclusiveDurations Collection.
	var exInclusiveDurations ast.ExInclusiveDurationsLike
	exInclusiveDurations, token, ok = v.parseExInclusiveDurations()
	if ok {
		// Found a single ExInclusiveDurations Collection.
		collection = ast.CollectionClass().Collection(exInclusiveDurations)
		return
	}

	// Attempt to parse a single ExclusiveDurations Collection.
	var exclusiveDurations ast.ExclusiveDurationsLike
	exclusiveDurations, token, ok = v.parseExclusiveDurations()
	if ok {
		// Found a single ExclusiveDurations Collection.
		collection = ast.CollectionClass().Collection(exclusiveDurations)
		return
	}

	// Attempt to parse a single InclusiveMoments Collection.
	var inclusiveMoments ast.InclusiveMomentsLike
	inclusiveMoments, token, ok = v.parseInclusiveMoments()
	if ok {
		// Found a single InclusiveMoments Collection.
		collection = ast.CollectionClass().Collection(inclusiveMoments)
		return
	}

	// Attempt to parse a single InExclusiveMoments Collection.
	var inExclusiveMoments ast.InExclusiveMomentsLike
	inExclusiveMoments, token, ok = v.parseInExclusiveMoments()
	if ok {
		// Found a single InExclusiveMoments Collection.
		collection = ast.CollectionClass().Collection(inExclusiveMoments)
		return
	}

	// Attempt to parse a single ExInclusiveMoments Collection.
	var exInclusiveMoments ast.ExInclusiveMomentsLike
	exInclusiveMoments, token, ok = v.parseExInclusiveMoments()
	if ok {
		// Found a single ExInclusiveMoments Collection.
		collection = ast.CollectionClass().Collection(exInclusiveMoments)
		return
	}

	// Attempt to parse a single ExclusiveMoments Collection.
	var exclusiveMoments ast.ExclusiveMomentsLike
	exclusiveMoments, token, ok = v.parseExclusiveMoments()
	if ok {
		// Found a single ExclusiveMoments Collection.
		collection = ast.CollectionClass().Collection(exclusiveMoments)
		return
	}

	// Attempt to parse a single InclusiveNames Collection.
	var inclusiveNames ast.InclusiveNamesLike
	inclusiveNames, token, ok = v.parseInclusiveNames()
	if ok {
		// Found a single InclusiveNames Collection.
		collection = ast.CollectionClass().Collection(inclusiveNames)
		return
	}

	// Attempt to parse a single InExclusiveNames Collection.
	var inExclusiveNames ast.InExclusiveNamesLike
	inExclusiveNames, token, ok = v.parseInExclusiveNames()
	if ok {
		// Found a single InExclusiveNames Collection.
		collection = ast.CollectionClass().Collection(inExclusiveNames)
		return
	}

	// Attempt to parse a single ExInclusiveNames Collection.
	var exInclusiveNames ast.ExInclusiveNamesLike
	exInclusiveNames, token, ok = v.parseExInclusiveNames()
	if ok {
		// Found a single ExInclusiveNames Collection.
		collection = ast.CollectionClass().Collection(exInclusiveNames)
		return
	}

	// Attempt to parse a single ExclusiveNames Collection.
	var exclusiveNames ast.ExclusiveNamesLike
	exclusiveNames, token, ok = v.parseExclusiveNames()
	if ok {
		// Found a single ExclusiveNames Collection.
		collection = ast.CollectionClass().Collection(exclusiveNames)
		return
	}

	// Attempt to parse a single InclusiveNumbers Collection.
	var inclusiveNumbers ast.InclusiveNumbersLike
	inclusiveNumbers, token, ok = v.parseInclusiveNumbers()
	if ok {
		// Found a single InclusiveNumbers Collection.
		collection = ast.CollectionClass().Collection(inclusiveNumbers)
		return
	}

	// Attempt to parse a single InExclusiveNumbers Collection.
	var inExclusiveNumbers ast.InExclusiveNumbersLike
	inExclusiveNumbers, token, ok = v.parseInExclusiveNumbers()
	if ok {
		// Found a single InExclusiveNumbers Collection.
		collection = ast.CollectionClass().Collection(inExclusiveNumbers)
		return
	}

	// Attempt to parse a single ExInclusiveNumbers Collection.
	var exInclusiveNumbers ast.ExInclusiveNumbersLike
	exInclusiveNumbers, token, ok = v.parseExInclusiveNumbers()
	if ok {
		// Found a single ExInclusiveNumbers Collection.
		collection = ast.CollectionClass().Collection(exInclusiveNumbers)
		return
	}

	// Attempt to parse a single ExclusiveNumbers Collection.
	var exclusiveNumbers ast.ExclusiveNumbersLike
	exclusiveNumbers, token, ok = v.parseExclusiveNumbers()
	if ok {
		// Found a single ExclusiveNumbers Collection.
		collection = ast.CollectionClass().Collection(exclusiveNumbers)
		return
	}

	// Attempt to parse a single InclusivePercentages Collection.
	var inclusivePercentages ast.InclusivePercentagesLike
	inclusivePercentages, token, ok = v.parseInclusivePercentages()
	if ok {
		// Found a single InclusivePercentages Collection.
		collection = ast.CollectionClass().Collection(inclusivePercentages)
		return
	}

	// Attempt to parse a single InExclusivePercentages Collection.
	var inExclusivePercentages ast.InExclusivePercentagesLike
	inExclusivePercentages, token, ok = v.parseInExclusivePercentages()
	if ok {
		// Found a single InExclusivePercentages Collection.
		collection = ast.CollectionClass().Collection(inExclusivePercentages)
		return
	}

	// Attempt to parse a single ExInclusivePercentages Collection.
	var exInclusivePercentages ast.ExInclusivePercentagesLike
	exInclusivePercentages, token, ok = v.parseExInclusivePercentages()
	if ok {
		// Found a single ExInclusivePercentages Collection.
		collection = ast.CollectionClass().Collection(exInclusivePercentages)
		return
	}

	// Attempt to parse a single ExclusivePercentages Collection.
	var exclusivePercentages ast.ExclusivePercentagesLike
	exclusivePercentages, token, ok = v.parseExclusivePercentages()
	if ok {
		// Found a single ExclusivePercentages Collection.
		collection = ast.CollectionClass().Collection(exclusivePercentages)
		return
	}

	// Attempt to parse a single InclusiveProbabilities Collection.
	var inclusiveProbabilities ast.InclusiveProbabilitiesLike
	inclusiveProbabilities, token, ok = v.parseInclusiveProbabilities()
	if ok {
		// Found a single InclusiveProbabilities Collection.
		collection = ast.CollectionClass().Collection(inclusiveProbabilities)
		return
	}

	// Attempt to parse a single InExclusiveProbabilities Collection.
	var inExclusiveProbabilities ast.InExclusiveProbabilitiesLike
	inExclusiveProbabilities, token, ok = v.parseInExclusiveProbabilities()
	if ok {
		// Found a single InExclusiveProbabilities Collection.
		collection = ast.CollectionClass().Collection(inExclusiveProbabilities)
		return
	}

	// Attempt to parse a single ExInclusiveProbabilities Collection.
	var exInclusiveProbabilities ast.ExInclusiveProbabilitiesLike
	exInclusiveProbabilities, token, ok = v.parseExInclusiveProbabilities()
	if ok {
		// Found a single ExInclusiveProbabilities Collection.
		collection = ast.CollectionClass().Collection(exInclusiveProbabilities)
		return
	}

	// Attempt to parse a single ExclusiveProbabilities Collection.
	var exclusiveProbabilities ast.ExclusiveProbabilitiesLike
	exclusiveProbabilities, token, ok = v.parseExclusiveProbabilities()
	if ok {
		// Found a single ExclusiveProbabilities Collection.
		collection = ast.CollectionClass().Collection(exclusiveProbabilities)
		return
	}

	// Attempt to parse a single InclusiveQuotes Collection.
	var inclusiveQuotes ast.InclusiveQuotesLike
	inclusiveQuotes, token, ok = v.parseInclusiveQuotes()
	if ok {
		// Found a single InclusiveQuotes Collection.
		collection = ast.CollectionClass().Collection(inclusiveQuotes)
		return
	}

	// Attempt to parse a single InExclusiveQuotes Collection.
	var inExclusiveQuotes ast.InExclusiveQuotesLike
	inExclusiveQuotes, token, ok = v.parseInExclusiveQuotes()
	if ok {
		// Found a single InExclusiveQuotes Collection.
		collection = ast.CollectionClass().Collection(inExclusiveQuotes)
		return
	}

	// Attempt to parse a single ExInclusiveQuotes Collection.
	var exInclusiveQuotes ast.ExInclusiveQuotesLike
	exInclusiveQuotes, token, ok = v.parseExInclusiveQuotes()
	if ok {
		// Found a single ExInclusiveQuotes Collection.
		collection = ast.CollectionClass().Collection(exInclusiveQuotes)
		return
	}

	// Attempt to parse a single ExclusiveQuotes Collection.
	var exclusiveQuotes ast.ExclusiveQuotesLike
	exclusiveQuotes, token, ok = v.parseExclusiveQuotes()
	if ok {
		// Found a single ExclusiveQuotes Collection.
		collection = ast.CollectionClass().Collection(exclusiveQuotes)
		return
	}

	// Attempt to parse a single InclusiveSymbols Collection.
	var inclusiveSymbols ast.InclusiveSymbolsLike
	inclusiveSymbols, token, ok = v.parseInclusiveSymbols()
	if ok {
		// Found a single InclusiveSymbols Collection.
		collection = ast.CollectionClass().Collection(inclusiveSymbols)
		return
	}

	// Attempt to parse a single InExclusiveSymbols Collection.
	var inExclusiveSymbols ast.InExclusiveSymbolsLike
	inExclusiveSymbols, token, ok = v.parseInExclusiveSymbols()
	if ok {
		// Found a single InExclusiveSymbols Collection.
		collection = ast.CollectionClass().Collection(inExclusiveSymbols)
		return
	}

	// Attempt to parse a single ExInclusiveSymbols Collection.
	var exInclusiveSymbols ast.ExInclusiveSymbolsLike
	exInclusiveSymbols, token, ok = v.parseExInclusiveSymbols()
	if ok {
		// Found a single ExInclusiveSymbols Collection.
		collection = ast.CollectionClass().Collection(exInclusiveSymbols)
		return
	}

	// Attempt to parse a single ExclusiveSymbols Collection.
	var exclusiveSymbols ast.ExclusiveSymbolsLike
	exclusiveSymbols, token, ok = v.parseExclusiveSymbols()
	if ok {
		// Found a single ExclusiveSymbols Collection.
		collection = ast.CollectionClass().Collection(exclusiveSymbols)
		return
	}

	// Attempt to parse a single InclusiveVersions Collection.
	var inclusiveVersions ast.InclusiveVersionsLike
	inclusiveVersions, token, ok = v.parseInclusiveVersions()
	if ok {
		// Found a single InclusiveVersions Collection.
		collection = ast.CollectionClass().Collection(inclusiveVersions)
		return
	}

	// Attempt to parse a single InExclusiveVersions Collection.
	var inExclusiveVersions ast.InExclusiveVersionsLike
	inExclusiveVersions, token, ok = v.parseInExclusiveVersions()
	if ok {
		// Found a single InExclusiveVersions Collection.
		collection = ast.CollectionClass().Collection(inExclusiveVersions)
		return
	}

	// Attempt to parse a single ExInclusiveVersions Collection.
	var exInclusiveVersions ast.ExInclusiveVersionsLike
	exInclusiveVersions, token, ok = v.parseExInclusiveVersions()
	if ok {
		// Found a single ExInclusiveVersions Collection.
		collection = ast.CollectionClass().Collection(exInclusiveVersions)
		return
	}

	// Attempt to parse a single ExclusiveVersions Collection.
	var exclusiveVersions ast.ExclusiveVersionsLike
	exclusiveVersions, token, ok = v.parseExclusiveVersions()
	if ok {
		// Found a single ExclusiveVersions Collection.
		collection = ast.CollectionClass().Collection(exclusiveVersions)
		return
	}

	// Attempt to parse a single InlineAttributes Collection.
	var inlineAttributes ast.InlineAttributesLike
	inlineAttributes, token, ok = v.parseInlineAttributes()
	if ok {
		// Found a single InlineAttributes Collection.
		collection = ast.CollectionClass().Collection(inlineAttributes)
		return
	}

	// Attempt to parse a single InlineValues Collection.
	var inlineValues ast.InlineValuesLike
	inlineValues, token, ok = v.parseInlineValues()
	if ok {
		// Found a single InlineValues Collection.
		collection = ast.CollectionClass().Collection(inlineValues)
		return
	}

	// This is not a single Collection rule.
	return
}

func (v *parser_) parseComplement() (
	complement ast.ComplementLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "not" delimiter.
	_, token, ok = v.parseDelimiter("not")
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
	complement = ast.ComplementClass().Complement(logical)
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

	// Attempt to parse a single citation Element.
	var citation string
	citation, token, ok = v.parseToken(CitationToken)
	if ok {
		// Found a single citation Element.
		element = ast.ElementClass().Element(citation)
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

func (v *parser_) parseEmptyLine() (
	emptyLine ast.EmptyLineLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single EmptyLine rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$EmptyLine", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single EmptyLine rule.
	ok = true
	v.remove(tokens)
	emptyLine = ast.EmptyLineClass().EmptyLine(newline)
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

func (v *parser_) parseExInclusiveAngles() (
	exInclusiveAngles ast.ExInclusiveAnglesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single angle token.
	var angle1 string
	angle1, token, ok = v.parseToken(AngleToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single angle token.
	var angle2 string
	angle2, token, ok = v.parseToken(AngleToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveAngles rule.
	ok = true
	v.remove(tokens)
	exInclusiveAngles = ast.ExInclusiveAnglesClass().ExInclusiveAngles(
		angle1,
		angle2,
	)
	return
}

func (v *parser_) parseExInclusiveCitations() (
	exInclusiveCitations ast.ExInclusiveCitationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single citation token.
	var citation1 string
	citation1, token, ok = v.parseToken(CitationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single citation token.
	var citation2 string
	citation2, token, ok = v.parseToken(CitationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveCitations rule.
	ok = true
	v.remove(tokens)
	exInclusiveCitations = ast.ExInclusiveCitationsClass().ExInclusiveCitations(
		citation1,
		citation2,
	)
	return
}

func (v *parser_) parseExInclusiveDurations() (
	exInclusiveDurations ast.ExInclusiveDurationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single duration token.
	var duration1 string
	duration1, token, ok = v.parseToken(DurationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single duration token.
	var duration2 string
	duration2, token, ok = v.parseToken(DurationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveDurations rule.
	ok = true
	v.remove(tokens)
	exInclusiveDurations = ast.ExInclusiveDurationsClass().ExInclusiveDurations(
		duration1,
		duration2,
	)
	return
}

func (v *parser_) parseExInclusiveMoments() (
	exInclusiveMoments ast.ExInclusiveMomentsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single moment token.
	var moment1 string
	moment1, token, ok = v.parseToken(MomentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single moment token.
	var moment2 string
	moment2, token, ok = v.parseToken(MomentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveMoments rule.
	ok = true
	v.remove(tokens)
	exInclusiveMoments = ast.ExInclusiveMomentsClass().ExInclusiveMoments(
		moment1,
		moment2,
	)
	return
}

func (v *parser_) parseExInclusiveNames() (
	exInclusiveNames ast.ExInclusiveNamesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name1 string
	name1, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name2 string
	name2, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveNames rule.
	ok = true
	v.remove(tokens)
	exInclusiveNames = ast.ExInclusiveNamesClass().ExInclusiveNames(
		name1,
		name2,
	)
	return
}

func (v *parser_) parseExInclusiveNumbers() (
	exInclusiveNumbers ast.ExInclusiveNumbersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number1 string
	number1, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number2 string
	number2, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveNumbers rule.
	ok = true
	v.remove(tokens)
	exInclusiveNumbers = ast.ExInclusiveNumbersClass().ExInclusiveNumbers(
		number1,
		number2,
	)
	return
}

func (v *parser_) parseExInclusivePercentages() (
	exInclusivePercentages ast.ExInclusivePercentagesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single percentage token.
	var percentage1 string
	percentage1, token, ok = v.parseToken(PercentageToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single percentage token.
	var percentage2 string
	percentage2, token, ok = v.parseToken(PercentageToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusivePercentages rule.
	ok = true
	v.remove(tokens)
	exInclusivePercentages = ast.ExInclusivePercentagesClass().ExInclusivePercentages(
		percentage1,
		percentage2,
	)
	return
}

func (v *parser_) parseExInclusiveProbabilities() (
	exInclusiveProbabilities ast.ExInclusiveProbabilitiesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single probability token.
	var probability1 string
	probability1, token, ok = v.parseToken(ProbabilityToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single probability token.
	var probability2 string
	probability2, token, ok = v.parseToken(ProbabilityToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveProbabilities rule.
	ok = true
	v.remove(tokens)
	exInclusiveProbabilities = ast.ExInclusiveProbabilitiesClass().ExInclusiveProbabilities(
		probability1,
		probability2,
	)
	return
}

func (v *parser_) parseExInclusiveQuotes() (
	exInclusiveQuotes ast.ExInclusiveQuotesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single quote token.
	var quote1 string
	quote1, token, ok = v.parseToken(QuoteToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single quote token.
	var quote2 string
	quote2, token, ok = v.parseToken(QuoteToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveQuotes rule.
	ok = true
	v.remove(tokens)
	exInclusiveQuotes = ast.ExInclusiveQuotesClass().ExInclusiveQuotes(
		quote1,
		quote2,
	)
	return
}

func (v *parser_) parseExInclusiveSymbols() (
	exInclusiveSymbols ast.ExInclusiveSymbolsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol1 string
	symbol1, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol2 string
	symbol2, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveSymbols rule.
	ok = true
	v.remove(tokens)
	exInclusiveSymbols = ast.ExInclusiveSymbolsClass().ExInclusiveSymbols(
		symbol1,
		symbol2,
	)
	return
}

func (v *parser_) parseExInclusiveVersions() (
	exInclusiveVersions ast.ExInclusiveVersionsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single version token.
	var version1 string
	version1, token, ok = v.parseToken(VersionToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single version token.
	var version2 string
	version2, token, ok = v.parseToken(VersionToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExInclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExInclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExInclusiveVersions rule.
	ok = true
	v.remove(tokens)
	exInclusiveVersions = ast.ExInclusiveVersionsClass().ExInclusiveVersions(
		version1,
		version2,
	)
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

func (v *parser_) parseExclusiveAngles() (
	exclusiveAngles ast.ExclusiveAnglesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single angle token.
	var angle1 string
	angle1, token, ok = v.parseToken(AngleToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single angle token.
	var angle2 string
	angle2, token, ok = v.parseToken(AngleToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveAngles rule.
	ok = true
	v.remove(tokens)
	exclusiveAngles = ast.ExclusiveAnglesClass().ExclusiveAngles(
		angle1,
		angle2,
	)
	return
}

func (v *parser_) parseExclusiveCitations() (
	exclusiveCitations ast.ExclusiveCitationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single citation token.
	var citation1 string
	citation1, token, ok = v.parseToken(CitationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single citation token.
	var citation2 string
	citation2, token, ok = v.parseToken(CitationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveCitations rule.
	ok = true
	v.remove(tokens)
	exclusiveCitations = ast.ExclusiveCitationsClass().ExclusiveCitations(
		citation1,
		citation2,
	)
	return
}

func (v *parser_) parseExclusiveDurations() (
	exclusiveDurations ast.ExclusiveDurationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single duration token.
	var duration1 string
	duration1, token, ok = v.parseToken(DurationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single duration token.
	var duration2 string
	duration2, token, ok = v.parseToken(DurationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveDurations rule.
	ok = true
	v.remove(tokens)
	exclusiveDurations = ast.ExclusiveDurationsClass().ExclusiveDurations(
		duration1,
		duration2,
	)
	return
}

func (v *parser_) parseExclusiveMoments() (
	exclusiveMoments ast.ExclusiveMomentsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single moment token.
	var moment1 string
	moment1, token, ok = v.parseToken(MomentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single moment token.
	var moment2 string
	moment2, token, ok = v.parseToken(MomentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveMoments rule.
	ok = true
	v.remove(tokens)
	exclusiveMoments = ast.ExclusiveMomentsClass().ExclusiveMoments(
		moment1,
		moment2,
	)
	return
}

func (v *parser_) parseExclusiveNames() (
	exclusiveNames ast.ExclusiveNamesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name1 string
	name1, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name2 string
	name2, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveNames rule.
	ok = true
	v.remove(tokens)
	exclusiveNames = ast.ExclusiveNamesClass().ExclusiveNames(
		name1,
		name2,
	)
	return
}

func (v *parser_) parseExclusiveNumbers() (
	exclusiveNumbers ast.ExclusiveNumbersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number1 string
	number1, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number2 string
	number2, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveNumbers rule.
	ok = true
	v.remove(tokens)
	exclusiveNumbers = ast.ExclusiveNumbersClass().ExclusiveNumbers(
		number1,
		number2,
	)
	return
}

func (v *parser_) parseExclusivePercentages() (
	exclusivePercentages ast.ExclusivePercentagesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single percentage token.
	var percentage1 string
	percentage1, token, ok = v.parseToken(PercentageToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single percentage token.
	var percentage2 string
	percentage2, token, ok = v.parseToken(PercentageToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusivePercentages rule.
	ok = true
	v.remove(tokens)
	exclusivePercentages = ast.ExclusivePercentagesClass().ExclusivePercentages(
		percentage1,
		percentage2,
	)
	return
}

func (v *parser_) parseExclusiveProbabilities() (
	exclusiveProbabilities ast.ExclusiveProbabilitiesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single probability token.
	var probability1 string
	probability1, token, ok = v.parseToken(ProbabilityToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single probability token.
	var probability2 string
	probability2, token, ok = v.parseToken(ProbabilityToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveProbabilities rule.
	ok = true
	v.remove(tokens)
	exclusiveProbabilities = ast.ExclusiveProbabilitiesClass().ExclusiveProbabilities(
		probability1,
		probability2,
	)
	return
}

func (v *parser_) parseExclusiveQuotes() (
	exclusiveQuotes ast.ExclusiveQuotesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single quote token.
	var quote1 string
	quote1, token, ok = v.parseToken(QuoteToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single quote token.
	var quote2 string
	quote2, token, ok = v.parseToken(QuoteToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveQuotes rule.
	ok = true
	v.remove(tokens)
	exclusiveQuotes = ast.ExclusiveQuotesClass().ExclusiveQuotes(
		quote1,
		quote2,
	)
	return
}

func (v *parser_) parseExclusiveSymbols() (
	exclusiveSymbols ast.ExclusiveSymbolsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol1 string
	symbol1, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol2 string
	symbol2, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveSymbols rule.
	ok = true
	v.remove(tokens)
	exclusiveSymbols = ast.ExclusiveSymbolsClass().ExclusiveSymbols(
		symbol1,
		symbol2,
	)
	return
}

func (v *parser_) parseExclusiveVersions() (
	exclusiveVersions ast.ExclusiveVersionsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single version token.
	var version1 string
	version1, token, ok = v.parseToken(VersionToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single version token.
	var version2 string
	version2, token, ok = v.parseToken(VersionToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single ExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$ExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single ExclusiveVersions rule.
	ok = true
	v.remove(tokens)
	exclusiveVersions = ast.ExclusiveVersionsClass().ExclusiveVersions(
		version1,
		version2,
	)
	return
}

func (v *parser_) parseExpression() (
	expression ast.ExpressionLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single Subject rule.
	var subject ast.SubjectLike
	subject, token, ok = v.parseSubject()
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

	// Attempt to parse multiple Predicate rules.
	var predicates = fra.List[ast.PredicateLike]()
predicatesLoop:
	for count := 0; count < mat.MaxInt; count++ {
		var predicate ast.PredicateLike
		predicate, token, ok = v.parsePredicate()
		if !ok {
			switch {
			case count >= 0:
				break predicatesLoop
			case uti.IsDefined(tokens):
				// This is not multiple Predicate rules.
				v.putBack(tokens)
				return
			default:
				// Found a syntax error.
				var message = v.formatError("$Expression", token)
				message += "0 or more Predicate rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		predicates.AppendValue(predicate)
	}

	// Found a single Expression rule.
	ok = true
	v.remove(tokens)
	expression = ast.ExpressionClass().Expression(
		subject,
		predicates,
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

func (v *parser_) parseInExclusiveAngles() (
	inExclusiveAngles ast.InExclusiveAnglesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single angle token.
	var angle1 string
	angle1, token, ok = v.parseToken(AngleToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single angle token.
	var angle2 string
	angle2, token, ok = v.parseToken(AngleToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveAngles rule.
	ok = true
	v.remove(tokens)
	inExclusiveAngles = ast.InExclusiveAnglesClass().InExclusiveAngles(
		angle1,
		angle2,
	)
	return
}

func (v *parser_) parseInExclusiveCitations() (
	inExclusiveCitations ast.InExclusiveCitationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single citation token.
	var citation1 string
	citation1, token, ok = v.parseToken(CitationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single citation token.
	var citation2 string
	citation2, token, ok = v.parseToken(CitationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveCitations rule.
	ok = true
	v.remove(tokens)
	inExclusiveCitations = ast.InExclusiveCitationsClass().InExclusiveCitations(
		citation1,
		citation2,
	)
	return
}

func (v *parser_) parseInExclusiveDurations() (
	inExclusiveDurations ast.InExclusiveDurationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single duration token.
	var duration1 string
	duration1, token, ok = v.parseToken(DurationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single duration token.
	var duration2 string
	duration2, token, ok = v.parseToken(DurationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveDurations rule.
	ok = true
	v.remove(tokens)
	inExclusiveDurations = ast.InExclusiveDurationsClass().InExclusiveDurations(
		duration1,
		duration2,
	)
	return
}

func (v *parser_) parseInExclusiveMoments() (
	inExclusiveMoments ast.InExclusiveMomentsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single moment token.
	var moment1 string
	moment1, token, ok = v.parseToken(MomentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single moment token.
	var moment2 string
	moment2, token, ok = v.parseToken(MomentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveMoments rule.
	ok = true
	v.remove(tokens)
	inExclusiveMoments = ast.InExclusiveMomentsClass().InExclusiveMoments(
		moment1,
		moment2,
	)
	return
}

func (v *parser_) parseInExclusiveNames() (
	inExclusiveNames ast.InExclusiveNamesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name1 string
	name1, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name2 string
	name2, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveNames rule.
	ok = true
	v.remove(tokens)
	inExclusiveNames = ast.InExclusiveNamesClass().InExclusiveNames(
		name1,
		name2,
	)
	return
}

func (v *parser_) parseInExclusiveNumbers() (
	inExclusiveNumbers ast.InExclusiveNumbersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number1 string
	number1, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number2 string
	number2, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveNumbers rule.
	ok = true
	v.remove(tokens)
	inExclusiveNumbers = ast.InExclusiveNumbersClass().InExclusiveNumbers(
		number1,
		number2,
	)
	return
}

func (v *parser_) parseInExclusivePercentages() (
	inExclusivePercentages ast.InExclusivePercentagesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single percentage token.
	var percentage1 string
	percentage1, token, ok = v.parseToken(PercentageToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single percentage token.
	var percentage2 string
	percentage2, token, ok = v.parseToken(PercentageToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusivePercentages rule.
	ok = true
	v.remove(tokens)
	inExclusivePercentages = ast.InExclusivePercentagesClass().InExclusivePercentages(
		percentage1,
		percentage2,
	)
	return
}

func (v *parser_) parseInExclusiveProbabilities() (
	inExclusiveProbabilities ast.InExclusiveProbabilitiesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single probability token.
	var probability1 string
	probability1, token, ok = v.parseToken(ProbabilityToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single probability token.
	var probability2 string
	probability2, token, ok = v.parseToken(ProbabilityToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveProbabilities rule.
	ok = true
	v.remove(tokens)
	inExclusiveProbabilities = ast.InExclusiveProbabilitiesClass().InExclusiveProbabilities(
		probability1,
		probability2,
	)
	return
}

func (v *parser_) parseInExclusiveQuotes() (
	inExclusiveQuotes ast.InExclusiveQuotesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single quote token.
	var quote1 string
	quote1, token, ok = v.parseToken(QuoteToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single quote token.
	var quote2 string
	quote2, token, ok = v.parseToken(QuoteToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveQuotes rule.
	ok = true
	v.remove(tokens)
	inExclusiveQuotes = ast.InExclusiveQuotesClass().InExclusiveQuotes(
		quote1,
		quote2,
	)
	return
}

func (v *parser_) parseInExclusiveSymbols() (
	inExclusiveSymbols ast.InExclusiveSymbolsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol1 string
	symbol1, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol2 string
	symbol2, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveSymbols rule.
	ok = true
	v.remove(tokens)
	inExclusiveSymbols = ast.InExclusiveSymbolsClass().InExclusiveSymbols(
		symbol1,
		symbol2,
	)
	return
}

func (v *parser_) parseInExclusiveVersions() (
	inExclusiveVersions ast.InExclusiveVersionsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single version token.
	var version1 string
	version1, token, ok = v.parseToken(VersionToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single version token.
	var version2 string
	version2, token, ok = v.parseToken(VersionToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InExclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InExclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InExclusiveVersions rule.
	ok = true
	v.remove(tokens)
	inExclusiveVersions = ast.InExclusiveVersionsClass().InExclusiveVersions(
		version1,
		version2,
	)
	return
}

func (v *parser_) parseInclusiveAngles() (
	inclusiveAngles ast.InclusiveAnglesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single angle token.
	var angle1 string
	angle1, token, ok = v.parseToken(AngleToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single angle token.
	var angle2 string
	angle2, token, ok = v.parseToken(AngleToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveAngles rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveAngles", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveAngles rule.
	ok = true
	v.remove(tokens)
	inclusiveAngles = ast.InclusiveAnglesClass().InclusiveAngles(
		angle1,
		angle2,
	)
	return
}

func (v *parser_) parseInclusiveCitations() (
	inclusiveCitations ast.InclusiveCitationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single citation token.
	var citation1 string
	citation1, token, ok = v.parseToken(CitationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single citation token.
	var citation2 string
	citation2, token, ok = v.parseToken(CitationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveCitations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveCitations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveCitations rule.
	ok = true
	v.remove(tokens)
	inclusiveCitations = ast.InclusiveCitationsClass().InclusiveCitations(
		citation1,
		citation2,
	)
	return
}

func (v *parser_) parseInclusiveDurations() (
	inclusiveDurations ast.InclusiveDurationsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single duration token.
	var duration1 string
	duration1, token, ok = v.parseToken(DurationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single duration token.
	var duration2 string
	duration2, token, ok = v.parseToken(DurationToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveDurations rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveDurations", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveDurations rule.
	ok = true
	v.remove(tokens)
	inclusiveDurations = ast.InclusiveDurationsClass().InclusiveDurations(
		duration1,
		duration2,
	)
	return
}

func (v *parser_) parseInclusiveMoments() (
	inclusiveMoments ast.InclusiveMomentsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single moment token.
	var moment1 string
	moment1, token, ok = v.parseToken(MomentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single moment token.
	var moment2 string
	moment2, token, ok = v.parseToken(MomentToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveMoments rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveMoments", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveMoments rule.
	ok = true
	v.remove(tokens)
	inclusiveMoments = ast.InclusiveMomentsClass().InclusiveMoments(
		moment1,
		moment2,
	)
	return
}

func (v *parser_) parseInclusiveNames() (
	inclusiveNames ast.InclusiveNamesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name1 string
	name1, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single name token.
	var name2 string
	name2, token, ok = v.parseToken(NameToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNames rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNames", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveNames rule.
	ok = true
	v.remove(tokens)
	inclusiveNames = ast.InclusiveNamesClass().InclusiveNames(
		name1,
		name2,
	)
	return
}

func (v *parser_) parseInclusiveNumbers() (
	inclusiveNumbers ast.InclusiveNumbersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number1 string
	number1, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single number token.
	var number2 string
	number2, token, ok = v.parseToken(NumberToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveNumbers rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveNumbers", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveNumbers rule.
	ok = true
	v.remove(tokens)
	inclusiveNumbers = ast.InclusiveNumbersClass().InclusiveNumbers(
		number1,
		number2,
	)
	return
}

func (v *parser_) parseInclusivePercentages() (
	inclusivePercentages ast.InclusivePercentagesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single percentage token.
	var percentage1 string
	percentage1, token, ok = v.parseToken(PercentageToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single percentage token.
	var percentage2 string
	percentage2, token, ok = v.parseToken(PercentageToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusivePercentages rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusivePercentages", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusivePercentages rule.
	ok = true
	v.remove(tokens)
	inclusivePercentages = ast.InclusivePercentagesClass().InclusivePercentages(
		percentage1,
		percentage2,
	)
	return
}

func (v *parser_) parseInclusiveProbabilities() (
	inclusiveProbabilities ast.InclusiveProbabilitiesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single probability token.
	var probability1 string
	probability1, token, ok = v.parseToken(ProbabilityToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single probability token.
	var probability2 string
	probability2, token, ok = v.parseToken(ProbabilityToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveProbabilities rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveProbabilities", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveProbabilities rule.
	ok = true
	v.remove(tokens)
	inclusiveProbabilities = ast.InclusiveProbabilitiesClass().InclusiveProbabilities(
		probability1,
		probability2,
	)
	return
}

func (v *parser_) parseInclusiveQuotes() (
	inclusiveQuotes ast.InclusiveQuotesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single quote token.
	var quote1 string
	quote1, token, ok = v.parseToken(QuoteToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single quote token.
	var quote2 string
	quote2, token, ok = v.parseToken(QuoteToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveQuotes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveQuotes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveQuotes rule.
	ok = true
	v.remove(tokens)
	inclusiveQuotes = ast.InclusiveQuotesClass().InclusiveQuotes(
		quote1,
		quote2,
	)
	return
}

func (v *parser_) parseInclusiveSymbols() (
	inclusiveSymbols ast.InclusiveSymbolsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol1 string
	symbol1, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single symbol token.
	var symbol2 string
	symbol2, token, ok = v.parseToken(SymbolToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveSymbols rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveSymbols", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveSymbols rule.
	ok = true
	v.remove(tokens)
	inclusiveSymbols = ast.InclusiveSymbolsClass().InclusiveSymbols(
		symbol1,
		symbol2,
	)
	return
}

func (v *parser_) parseInclusiveVersions() (
	inclusiveVersions ast.InclusiveVersionsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single version token.
	var version1 string
	version1, token, ok = v.parseToken(VersionToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single ".." delimiter.
	_, token, ok = v.parseDelimiter("..")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single version token.
	var version2 string
	version2, token, ok = v.parseToken(VersionToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InclusiveVersions rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InclusiveVersions", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InclusiveVersions rule.
	ok = true
	v.remove(tokens)
	inclusiveVersions = ast.InclusiveVersionsClass().InclusiveVersions(
		version1,
		version2,
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

	// Attempt to parse a single Subcomponent Indirect.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Indirect.
		indirect = ast.IndirectClass().Indirect(subcomponent)
		return
	}

	// Attempt to parse a single Referent Indirect.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Indirect.
		indirect = ast.IndirectClass().Indirect(referent)
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

func (v *parser_) parseInlineAttributes() (
	inlineAttributes ast.InlineAttributesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineAttributes", token)
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
		// This is not a single InlineAttributes rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineAttributes", token)
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
				var message = v.formatError("$InlineAttributes", token)
				message += "0 or more AdditionalAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalAssociations.AppendValue(additionalAssociation)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InlineAttributes rule.
	ok = true
	v.remove(tokens)
	inlineAttributes = ast.InlineAttributesClass().InlineAttributes(
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

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineParameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineParameters", token)
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
		// This is not a single InlineParameters rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$InlineParameters", token)
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
				var message = v.formatError("$InlineParameters", token)
				message += "0 or more AdditionalAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		additionalAssociations.AppendValue(additionalAssociation)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineParameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineParameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InlineParameters rule.
	ok = true
	v.remove(tokens)
	inlineParameters = ast.InlineParametersClass().InlineParameters(
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

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineStatements", token)
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

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
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

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineValues", token)
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

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single InlineValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$InlineValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single InlineValues rule.
	ok = true
	v.remove(tokens)
	inlineValues = ast.InlineValuesClass().InlineValues(
		component,
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

	// Attempt to parse a single Subcomponent Logical.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Logical.
		logical = ast.LogicalClass().Logical(subcomponent)
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

	// Attempt to parse a single Referent Logical.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Logical.
		logical = ast.LogicalClass().Logical(referent)
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

func (v *parser_) parseMagnitude() (
	magnitude ast.MagnitudeLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single bar token.
	var bar1 string
	bar1, token, ok = v.parseToken(BarToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Magnitude rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Magnitude", token)
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
		// This is not a single Magnitude rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Magnitude", token)
		panic(message)
	}

	// Attempt to parse a single bar token.
	var bar2 string
	bar2, token, ok = v.parseToken(BarToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Magnitude rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Magnitude", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single Magnitude rule.
	ok = true
	v.remove(tokens)
	magnitude = ast.MagnitudeClass().Magnitude(
		bar1,
		numerical,
		bar2,
	)
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

func (v *parser_) parseMultilineAttributes() (
	multilineAttributes ast.MultilineAttributesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

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
				var message = v.formatError("$MultilineAttributes", token)
				message += "1 or more AnnotatedAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedAssociations.AppendValue(annotatedAssociation)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single MultilineAttributes rule.
	ok = true
	v.remove(tokens)
	multilineAttributes = ast.MultilineAttributesClass().MultilineAttributes(
		newline,
		annotatedAssociations,
	)
	return
}

func (v *parser_) parseMultilineParameters() (
	multilineParameters ast.MultilineParametersLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "(" delimiter.
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineParameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineParameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineParameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineParameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

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
				var message = v.formatError("$MultilineParameters", token)
				message += "1 or more AnnotatedAssociation rules are required."
				panic(message)
			}
		}
		// No additional put backs allowed at this point.
		tokens = nil
		annotatedAssociations.AppendValue(annotatedAssociation)
	}

	// Attempt to parse a single ")" delimiter.
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineParameters rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineParameters", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single MultilineParameters rule.
	ok = true
	v.remove(tokens)
	multilineParameters = ast.MultilineParametersClass().MultilineParameters(
		newline,
		annotatedAssociations,
	)
	return
}

func (v *parser_) parseMultilineStatements() (
	multilineStatements ast.MultilineStatementsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

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

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single MultilineStatements rule.
	ok = true
	v.remove(tokens)
	multilineStatements = ast.MultilineStatementsClass().MultilineStatements(
		newline,
		annotatedStatements,
	)
	return
}

func (v *parser_) parseMultilineValues() (
	multilineValues ast.MultilineValuesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single newline token.
	var newline string
	newline, token, ok = v.parseToken(NewlineToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

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

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single MultilineValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$MultilineValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single MultilineValues rule.
	ok = true
	v.remove(tokens)
	multilineValues = ast.MultilineValuesClass().MultilineValues(
		newline,
		annotatedValues,
	)
	return
}

func (v *parser_) parseNoAttributes() (
	noAttributes ast.NoAttributesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoAttributes", token)
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
			// This is not a single NoAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoAttributes rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoAttributes", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single NoAttributes rule.
	ok = true
	v.remove(tokens)
	noAttributes = ast.NoAttributesClass().NoAttributes(colon)
	return
}

func (v *parser_) parseNoStatements() (
	noStatements ast.NoStatementsLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "{" delimiter.
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "}" delimiter.
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoStatements rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoStatements", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single NoStatements rule.
	ok = true
	v.remove(tokens)
	noStatements = ast.NoStatementsClass().NoStatements()
	return
}

func (v *parser_) parseNoValues() (
	noValues ast.NoValuesLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single "[" delimiter.
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Attempt to parse a single "]" delimiter.
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single NoValues rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$NoValues", token)
			panic(message)
		}
	}
	if uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single NoValues rule.
	ok = true
	v.remove(tokens)
	noValues = ast.NoValuesClass().NoValues()
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

	// Attempt to parse a single Subcomponent Numerical.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Numerical.
		numerical = ast.NumericalClass().Numerical(subcomponent)
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

	// Attempt to parse a single Referent Numerical.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Numerical.
		numerical = ast.NumericalClass().Numerical(referent)
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

	// Attempt to parse a single Magnitude Numerical.
	var magnitude ast.MagnitudeLike
	magnitude, token, ok = v.parseMagnitude()
	if ok {
		// Found a single Magnitude Numerical.
		numerical = ast.NumericalClass().Numerical(magnitude)
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
	// Attempt to parse a single MultilineParameters Parameters.
	var multilineParameters ast.MultilineParametersLike
	multilineParameters, token, ok = v.parseMultilineParameters()
	if ok {
		// Found a single MultilineParameters Parameters.
		parameters = ast.ParametersClass().Parameters(multilineParameters)
		return
	}

	// Attempt to parse a single InlineParameters Parameters.
	var inlineParameters ast.InlineParametersLike
	inlineParameters, token, ok = v.parseInlineParameters()
	if ok {
		// Found a single InlineParameters Parameters.
		parameters = ast.ParametersClass().Parameters(inlineParameters)
		return
	}

	// This is not a single Parameters rule.
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

func (v *parser_) parsePredicate() (
	predicate ast.PredicateLike,
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
		// This is not a single Predicate rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Predicate", token)
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
		// This is not a single Predicate rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Predicate", token)
		panic(message)
	}

	// Found a single Predicate rule.
	ok = true
	v.remove(tokens)
	predicate = ast.PredicateClass().Predicate(
		operator,
		expression,
	)
	return
}

func (v *parser_) parseProcedure() (
	procedure ast.ProcedureLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single MultilineStatements Procedure.
	var multilineStatements ast.MultilineStatementsLike
	multilineStatements, token, ok = v.parseMultilineStatements()
	if ok {
		// Found a single MultilineStatements Procedure.
		procedure = ast.ProcedureClass().Procedure(multilineStatements)
		return
	}

	// Attempt to parse a single InlineStatements Procedure.
	var inlineStatements ast.InlineStatementsLike
	inlineStatements, token, ok = v.parseInlineStatements()
	if ok {
		// Found a single InlineStatements Procedure.
		procedure = ast.ProcedureClass().Procedure(inlineStatements)
		return
	}

	// Attempt to parse a single NoStatements Procedure.
	var noStatements ast.NoStatementsLike
	noStatements, token, ok = v.parseNoStatements()
	if ok {
		// Found a single NoStatements Procedure.
		procedure = ast.ProcedureClass().Procedure(noStatements)
		return
	}

	// This is not a single Procedure rule.
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

	// Attempt to parse a single Subcomponent Recipient.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Recipient.
		recipient = ast.RecipientClass().Recipient(subcomponent)
		return
	}

	// This is not a single Recipient rule.
	return
}

func (v *parser_) parseReferent() (
	referent ast.ReferentLike,
	token TokenLike,
	ok bool,
) {
	var tokens = fra.List[TokenLike]()

	// Attempt to parse a single snail token.
	var snail string
	snail, token, ok = v.parseToken(SnailToken)
	if !ok {
		if uti.IsDefined(tokens) {
			// This is not a single Referent rule.
			v.putBack(tokens)
			return
		} else {
			// Found a syntax error.
			var message = v.formatError("$Referent", token)
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
		// This is not a single Referent rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$Referent", token)
		panic(message)
	}

	// Found a single Referent rule.
	ok = true
	v.remove(tokens)
	referent = ast.ReferentClass().Referent(
		snail,
		indirect,
	)
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

func (v *parser_) parseStatementLine() (
	statementLine ast.StatementLineLike,
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
		// This is not a single StatementLine rule.
		v.putBack(tokens)
		return
	default:
		// Found a syntax error.
		var message = v.formatError("$StatementLine", token)
		panic(message)
	}

	// Attempt to parse an optional note token.
	var optionalNote string
	optionalNote, token, ok = v.parseToken(NoteToken)
	if ok && uti.IsDefined(tokens) {
		tokens.AppendValue(token)
	}

	// Found a single StatementLine rule.
	ok = true
	v.remove(tokens)
	statementLine = ast.StatementLineClass().StatementLine(
		statement,
		optionalNote,
	)
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

	// Attempt to parse a single identifier token.
	var identifier string
	identifier, token, ok = v.parseToken(IdentifierToken)
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
		identifier,
		indices,
	)
	return
}

func (v *parser_) parseSubject() (
	subject ast.SubjectLike,
	token TokenLike,
	ok bool,
) {
	// Attempt to parse a single Component Subject.
	var component ast.ComponentLike
	component, token, ok = v.parseComponent()
	if ok {
		// Found a single Component Subject.
		subject = ast.SubjectClass().Subject(component)
		return
	}

	// Attempt to parse a single Subcomponent Subject.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Subject.
		subject = ast.SubjectClass().Subject(subcomponent)
		return
	}

	// Attempt to parse a single Precedence Subject.
	var precedence ast.PrecedenceLike
	precedence, token, ok = v.parsePrecedence()
	if ok {
		// Found a single Precedence Subject.
		subject = ast.SubjectClass().Subject(precedence)
		return
	}

	// Attempt to parse a single Referent Subject.
	var referent ast.ReferentLike
	referent, token, ok = v.parseReferent()
	if ok {
		// Found a single Referent Subject.
		subject = ast.SubjectClass().Subject(referent)
		return
	}

	// Attempt to parse a single Complement Subject.
	var complement ast.ComplementLike
	complement, token, ok = v.parseComplement()
	if ok {
		// Found a single Complement Subject.
		subject = ast.SubjectClass().Subject(complement)
		return
	}

	// Attempt to parse a single Inversion Subject.
	var inversion ast.InversionLike
	inversion, token, ok = v.parseInversion()
	if ok {
		// Found a single Inversion Subject.
		subject = ast.SubjectClass().Subject(inversion)
		return
	}

	// Attempt to parse a single Magnitude Subject.
	var magnitude ast.MagnitudeLike
	magnitude, token, ok = v.parseMagnitude()
	if ok {
		// Found a single Magnitude Subject.
		subject = ast.SubjectClass().Subject(magnitude)
		return
	}

	// Attempt to parse a single Function Subject.
	var function ast.FunctionLike
	function, token, ok = v.parseFunction()
	if ok {
		// Found a single Function Subject.
		subject = ast.SubjectClass().Subject(function)
		return
	}

	// Attempt to parse a single Method Subject.
	var method ast.MethodLike
	method, token, ok = v.parseMethod()
	if ok {
		// Found a single Method Subject.
		subject = ast.SubjectClass().Subject(method)
		return
	}

	// Attempt to parse a single Variable Subject.
	var variable ast.VariableLike
	variable, token, ok = v.parseVariable()
	if ok {
		// Found a single Variable Subject.
		subject = ast.SubjectClass().Subject(variable)
		return
	}

	// This is not a single Subject rule.
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

	// Attempt to parse a single Subcomponent Target.
	var subcomponent ast.SubcomponentLike
	subcomponent, token, ok = v.parseSubcomponent()
	if ok {
		// Found a single Subcomponent Target.
		target = ast.TargetClass().Target(subcomponent)
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
			"$Parameters": `
  - MultilineParameters
  - InlineParameters`,
			"$MultilineParameters":   `"(" newline AnnotatedAssociation+ ")"`,
			"$AnnotatedAssociation":  `Association note?`,
			"$InlineParameters":      `"(" Association AdditionalAssociation* ")"`,
			"$AdditionalAssociation": `"," Association`,
			"$Association":           `symbol colon Component`,
			"$Element": `
  - angle
  - boolean
  - citation
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
			"$Collection": `
  - MultilineAttributes
  - MultilineValues
  - NoAttributes
  - NoValues
  - InclusiveAngles
  - InExclusiveAngles
  - ExInclusiveAngles
  - ExclusiveAngles
  - InclusiveCitations
  - InExclusiveCitations
  - ExInclusiveCitations
  - ExclusiveCitations
  - InclusiveDurations
  - InExclusiveDurations
  - ExInclusiveDurations
  - ExclusiveDurations
  - InclusiveMoments
  - InExclusiveMoments
  - ExInclusiveMoments
  - ExclusiveMoments
  - InclusiveNames
  - InExclusiveNames
  - ExInclusiveNames
  - ExclusiveNames
  - InclusiveNumbers
  - InExclusiveNumbers
  - ExInclusiveNumbers
  - ExclusiveNumbers
  - InclusivePercentages
  - InExclusivePercentages
  - ExInclusivePercentages
  - ExclusivePercentages
  - InclusiveProbabilities
  - InExclusiveProbabilities
  - ExInclusiveProbabilities
  - ExclusiveProbabilities
  - InclusiveQuotes
  - InExclusiveQuotes
  - ExInclusiveQuotes
  - ExclusiveQuotes
  - InclusiveSymbols
  - InExclusiveSymbols
  - ExInclusiveSymbols
  - ExclusiveSymbols
  - InclusiveVersions
  - InExclusiveVersions
  - ExInclusiveVersions
  - ExclusiveVersions
  - InlineAttributes
  - InlineValues`,
			"$MultilineAttributes":      `"[" newline AnnotatedAssociation+ "]"`,
			"$InlineAttributes":         `"[" Association AdditionalAssociation* "]"`,
			"$NoAttributes":             `"[" colon "]"`,
			"$MultilineValues":          `"[" newline AnnotatedValue+ "]"`,
			"$AnnotatedValue":           `Component note?`,
			"$InlineValues":             `"[" Component AdditionalValue* "]"`,
			"$AdditionalValue":          `"," Component`,
			"$NoValues":                 `"[" "]"`,
			"$InclusiveAngles":          `"[" angle ".." angle "]"`,
			"$InExclusiveAngles":        `"[" angle ".." angle ")"`,
			"$ExInclusiveAngles":        `"(" angle ".." angle "]"`,
			"$ExclusiveAngles":          `"(" angle ".." angle ")"`,
			"$InclusiveCitations":       `"[" citation ".." citation "]"`,
			"$InExclusiveCitations":     `"[" citation ".." citation ")"`,
			"$ExInclusiveCitations":     `"(" citation ".." citation "]"`,
			"$ExclusiveCitations":       `"(" citation ".." citation ")"`,
			"$InclusiveDurations":       `"[" duration ".." duration "]"`,
			"$InExclusiveDurations":     `"[" duration ".." duration ")"`,
			"$ExInclusiveDurations":     `"(" duration ".." duration "]"`,
			"$ExclusiveDurations":       `"(" duration ".." duration ")"`,
			"$InclusiveMoments":         `"[" moment ".." moment "]"`,
			"$InExclusiveMoments":       `"[" moment ".." moment ")"`,
			"$ExInclusiveMoments":       `"(" moment ".." moment "]"`,
			"$ExclusiveMoments":         `"(" moment ".." moment ")"`,
			"$InclusiveNames":           `"[" name ".." name "]"`,
			"$InExclusiveNames":         `"[" name ".." name ")"`,
			"$ExInclusiveNames":         `"(" name ".." name "]"`,
			"$ExclusiveNames":           `"(" name ".." name ")"`,
			"$InclusiveNumbers":         `"[" number ".." number "]"`,
			"$InExclusiveNumbers":       `"[" number ".." number ")"`,
			"$ExInclusiveNumbers":       `"(" number ".." number "]"`,
			"$ExclusiveNumbers":         `"(" number ".." number ")"`,
			"$InclusivePercentages":     `"[" percentage ".." percentage "]"`,
			"$InExclusivePercentages":   `"[" percentage ".." percentage ")"`,
			"$ExInclusivePercentages":   `"(" percentage ".." percentage "]"`,
			"$ExclusivePercentages":     `"(" percentage ".." percentage ")"`,
			"$InclusiveProbabilities":   `"[" probability ".." probability "]"`,
			"$InExclusiveProbabilities": `"[" probability ".." probability ")"`,
			"$ExInclusiveProbabilities": `"(" probability ".." probability "]"`,
			"$ExclusiveProbabilities":   `"(" probability ".." probability ")"`,
			"$InclusiveQuotes":          `"[" quote ".." quote "]"`,
			"$InExclusiveQuotes":        `"[" quote ".." quote ")"`,
			"$ExInclusiveQuotes":        `"(" quote ".." quote "]"`,
			"$ExclusiveQuotes":          `"(" quote ".." quote ")"`,
			"$InclusiveSymbols":         `"[" symbol ".." symbol "]"`,
			"$InExclusiveSymbols":       `"[" symbol ".." symbol ")"`,
			"$ExInclusiveSymbols":       `"(" symbol ".." symbol "]"`,
			"$ExclusiveSymbols":         `"(" symbol ".." symbol ")"`,
			"$InclusiveVersions":        `"[" version ".." version "]"`,
			"$InExclusiveVersions":      `"[" version ".." version ")"`,
			"$ExInclusiveVersions":      `"(" version ".." version "]"`,
			"$ExclusiveVersions":        `"(" version ".." version ")"`,
			"$Procedure": `
  - MultilineStatements
  - InlineStatements
  - NoStatements`,
			"$MultilineStatements": `"{" newline AnnotatedStatement+ "}"`,
			"$AnnotatedStatement": `
  - EmptyLine
  - AnnotationLine
  - StatementLine`,
			"$EmptyLine": `newline`,
			"$AnnotationLine": `
  - note
  - comment`,
			"$StatementLine":       `Statement note?`,
			"$InlineStatements":    `"{" Statement AdditionalStatement* "}"`,
			"$AdditionalStatement": `";" Statement`,
			"$NoStatements":        `"{" "}"`,
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
  - Subcomponent
  - Variable  ! This must be last since others also begin with an identifier.`,
			"$Function":           `identifier "(" Arguments? ")"`,
			"$Arguments":          `Argument AdditionalArgument*`,
			"$AdditionalArgument": `"," Argument`,
			"$Argument":           `identifier`,
			"$Method":             `identifier Threading identifier "(" Arguments? ")"`,
			"$Threading": `
  - dot
  - arrow`,
			"$Variable":       `identifier`,
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
			"$DoClause":       `"do" Invocation`,
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
  - Subcomponent`,
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
			"$Expression":     `Subject Predicate*`,
			"$Subject": `
  - Component
  - Subcomponent
  - Precedence
  - Referent
  - Complement
  - Inversion
  - Magnitude
  - Function
  - Method
  - Variable  ! This must be last since others also begin with an identifier.`,
			"$Subcomponent":    `identifier "[" Indices "]"`,
			"$Indices":         `Index AdditionalIndex*`,
			"$AdditionalIndex": `"," Index`,
			"$Index":           `Expression`,
			"$Predicate":       `Operator Expression`,
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
			"$Precedence": `"(" Expression ")"`,
			"$Referent":   `snail Indirect`,
			"$Indirect": `
  - Component
  - Subcomponent
  - Referent
  - Function
  - Method
  - Variable  ! This must be last since others also begin with an identifier.`,
			"$Complement": `"not" Logical`,
			"$Logical": `
  - Component
  - Subcomponent
  - Precedence
  - Referent
  - Complement
  - Function
  - Method
  - Variable  ! This must be last since others also begin with an identifier.`,
			"$Inversion": `Inverse Numerical`,
			"$Inverse": `
  - dash
  - slash
  - star`,
			"$Numerical": `
  - Component
  - Subcomponent
  - Precedence
  - Referent
  - Inversion
  - Magnitude
  - Function
  - Method
  - Variable  ! This must be last since others also begin with an identifier.`,
			"$Magnitude": `bar Numerical bar`,
		},
	),
}
