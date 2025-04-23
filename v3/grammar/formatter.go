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
│ Updates to any section other than the Methodical Methods may be overwritten. │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package grammar

import (
	ast "github.com/bali-nebula/go-component-framework/v3/ast"
	sts "strings"
)

// CLASS INTERFACE

// Access Function

func FormatterClass() FormatterClassLike {
	return formatterClass()
}

// Constructor Methods

func (c *formatterClass_) Formatter() FormatterLike {
	var instance = &formatter_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *formatter_) GetClass() FormatterClassLike {
	return formatterClass()
}

func (v *formatter_) FormatDocument(document ast.DocumentLike) string {
	v.visitor_.VisitDocument(document)
	return v.getResult()
}

// Methodical Methods

func (v *formatter_) ProcessAmpersand(
	ampersand string,
) {
	v.appendString(ampersand)
}

func (v *formatter_) ProcessAngle(
	angle string,
) {
	v.appendString(angle)
}

func (v *formatter_) ProcessArrow(
	arrow string,
) {
	v.appendString(arrow)
}

func (v *formatter_) ProcessBinary(
	binary string,
) {
	v.appendString(binary)
}

func (v *formatter_) ProcessBoolean(
	boolean string,
) {
	v.appendString(boolean)
}

func (v *formatter_) ProcessBytecode(
	bytecode string,
) {
	v.appendString(bytecode)
}

func (v *formatter_) ProcessCaret(
	caret string,
) {
	v.appendString(caret)
}

func (v *formatter_) ProcessCitation(
	citation string,
) {
	v.appendString(citation)
}

func (v *formatter_) ProcessComment(
	comment string,
) {
	v.appendString(comment)
}

func (v *formatter_) ProcessDash(
	dash string,
) {
	v.appendString(dash)
}

func (v *formatter_) ProcessDot(
	dot string,
) {
	v.appendString(dot)
}

func (v *formatter_) ProcessDouble(
	double string,
) {
	v.appendString(double)
}

func (v *formatter_) ProcessDuration(
	duration string,
) {
	v.appendString(duration)
}

func (v *formatter_) ProcessEqual(
	equal string,
) {
	v.appendString(equal)
}

func (v *formatter_) PreprocessExclusion(
	exclusion ast.ExclusionLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessExclusiveRange(
	exclusiveRange ast.ExclusiveRangeLike,
) {
	v.appendString("(")
}

func (v *formatter_) ProcessExclusiveRangeSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString("..")
	}
}

func (v *formatter_) ProcessIdentifier(
	identifier string,
) {
	v.appendString(identifier)
}

func (v *formatter_) PreprocessInclusion(
	inclusion ast.InclusionLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessInclusiveRange(
	inclusiveRange ast.InclusiveRangeLike,
) {
	v.appendString("[")
}

func (v *formatter_) ProcessInclusiveRangeSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString("..")
	}
}

func (v *formatter_) ProcessLess(
	less string,
) {
	v.appendString(less)
}

func (v *formatter_) ProcessMoment(
	moment string,
) {
	v.appendString(moment)
}

func (v *formatter_) ProcessMore(
	more string,
) {
	v.appendString(more)
}

func (v *formatter_) ProcessName(
	name string,
) {
	v.appendString(name)
}

func (v *formatter_) ProcessNarrative(
	narrative string,
) {
	v.appendString(narrative)
}

func (v *formatter_) ProcessNote(
	note string,
) {
	v.appendString("  ")
	v.appendString(note)
}

func (v *formatter_) ProcessNumber(
	number string,
) {
	v.appendString(number)
}

func (v *formatter_) ProcessPattern(
	pattern string,
) {
	v.appendString(pattern)
}

func (v *formatter_) ProcessPercentage(
	percentage string,
) {
	v.appendString(percentage)
}

func (v *formatter_) ProcessPlus(
	plus string,
) {
	v.appendString(plus)
}

func (v *formatter_) ProcessProbability(
	probability string,
) {
	v.appendString(probability)
}

func (v *formatter_) ProcessQuote(
	quote string,
) {
	v.appendString(quote)
}

func (v *formatter_) ProcessResource(
	resource string,
) {
	v.appendString(resource)
}

func (v *formatter_) ProcessSlash(
	slash string,
) {
	v.appendString(slash)
}

func (v *formatter_) ProcessStar(
	star string,
) {
	v.appendString(star)
}

func (v *formatter_) ProcessSymbol(
	symbol string,
) {
	v.appendString(symbol)
}

func (v *formatter_) ProcessTag(
	tag string,
) {
	v.appendString(tag)
}

func (v *formatter_) ProcessVersion(
	version string,
) {
	v.appendString(version)
}

func (v *formatter_) PreprocessAcceptClause(
	acceptClause ast.AcceptClauseLike,
) {
	v.appendString("accept ")
}

func (v *formatter_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	v.appendString(", ")
}

func (v *formatter_) PreprocessAdditionalAssociation(
	additionalAssociation ast.AdditionalAssociationLike,
	index uint,
	size uint,
) {
	v.appendString(", ")
}

func (v *formatter_) PreprocessAdditionalIndex(
	additionalIndex ast.AdditionalIndexLike,
	index uint,
	size uint,
) {
	v.appendString(", ")
}

func (v *formatter_) PreprocessAdditionalStatement(
	additionalStatement ast.AdditionalStatementLike,
	index uint,
	size uint,
) {
	v.appendString("; ")
}

func (v *formatter_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	v.appendString(", ")
}

func (v *formatter_) PreprocessAnd(
	and ast.AndLike,
) {
	v.appendString("and")
}

func (v *formatter_) PreprocessAnnotatedAssociation(
	annotatedAssociation ast.AnnotatedAssociationLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessAnnotatedValue(
	annotatedValue ast.AnnotatedValueLike,
	index uint,
	size uint,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessAssociationSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(": ")
	}
}

func (v *formatter_) PreprocessAtLevel(
	atLevel ast.AtLevelLike,
) {
	v.appendString(" at level ")
}

func (v *formatter_) PreprocessBreakClause(
	breakClause ast.BreakClauseLike,
) {
	v.appendString("break loop")
}

func (v *formatter_) PreprocessCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	v.appendString("checkout ")
}

func (v *formatter_) ProcessCheckoutClauseSlot(
	slot uint,
) {
	switch slot {
	case 2:
		v.appendString(" from ")
	}
}

func (v *formatter_) PreprocessColonEqual(
	colonEqual ast.ColonEqualLike,
) {
	v.appendString(":=")
}

func (v *formatter_) PreprocessCommentLine(
	commentLine ast.CommentLineLike,
) {
	v.appendString("\n")
	v.appendNewline()
}

func (v *formatter_) PreprocessComplement(
	complement ast.ComplementLike,
) {
	v.appendString("not ")
}

func (v *formatter_) PreprocessContinueClause(
	continueClause ast.ContinueClauseLike,
) {
	v.appendString("continue loop")
}

func (v *formatter_) PreprocessDashEqual(
	dashEqual ast.DashEqualLike,
) {
	v.appendString("-=")
}

func (v *formatter_) PreprocessDefaultEqual(
	defaultEqual ast.DefaultEqualLike,
) {
	v.appendString("?=")
}

func (v *formatter_) PreprocessDiscardClause(
	discardClause ast.DiscardClauseLike,
) {
	v.appendString("discard ")
}

func (v *formatter_) PreprocessDoClause(
	doClause ast.DoClauseLike,
) {
	v.appendString("do ")
}

func (v *formatter_) PostprocessDocument(
	document ast.DocumentLike,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessFunctionSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString("(")
	}
}

func (v *formatter_) PostprocessFunction(
	function ast.FunctionLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessIfClause(
	ifClause ast.IfClauseLike,
) {
	v.appendString("if ")
}

func (v *formatter_) ProcessIfClauseSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" do ")
	}
}

func (v *formatter_) PreprocessInlineAttributes(
	inlineAttributes ast.InlineAttributesLike,
) {
	v.appendString("[")
}

func (v *formatter_) PostprocessInlineAttributes(
	inlineAttributes ast.InlineAttributesLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessInlineParameters(
	inlineParameters ast.InlineParametersLike,
) {
	v.appendString("(")
}

func (v *formatter_) PostprocessInlineParameters(
	inlineParameters ast.InlineParametersLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInlineStatements(
	inlineStatements ast.InlineStatementsLike,
) {
	v.appendString("{")
}

func (v *formatter_) PostprocessInlineStatements(
	inlineStatements ast.InlineStatementsLike,
) {
	v.appendString("}")
}

func (v *formatter_) PreprocessInlineValues(
	inlineValues ast.InlineValuesLike,
) {
	v.appendString("[")
}

func (v *formatter_) PostprocessInlineValues(
	inlineValues ast.InlineValuesLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessIor(
	ior ast.IorLike,
) {
	v.appendString("ior")
}

func (v *formatter_) PreprocessIs(
	is ast.IsLike,
) {
	v.appendString("is")
}

func (v *formatter_) PreprocessLetClause(
	letClause ast.LetClauseLike,
) {
	v.appendString("let ")
}

func (v *formatter_) ProcessLetClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessMagnitude(
	magnitude ast.MagnitudeLike,
) {
	v.appendString("|")
}

func (v *formatter_) PostprocessMagnitude(
	magnitude ast.MagnitudeLike,
) {
	v.appendString("|")
}

func (v *formatter_) PreprocessMatchHandler(
	matchHandler ast.MatchHandlerLike,
	index uint,
	size uint,
) {
	v.appendString(" matching ")
}

func (v *formatter_) ProcessMatchHandlerSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" do ")
	}
}

func (v *formatter_) PreprocessMatches(
	matches ast.MatchesLike,
) {
	v.appendString("matches")
}

func (v *formatter_) ProcessMethodSlot(
	slot uint,
) {
	switch slot {
	case 3:
		v.appendString("(")
	}
}

func (v *formatter_) PostprocessMethod(
	method ast.MethodLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessMultilineAttributes(
	multilineAttributes ast.MultilineAttributesLike,
) {
	v.appendString("[")
	v.depth_++
}

func (v *formatter_) PostprocessMultilineAttributes(
	multilineAttributes ast.MultilineAttributesLike,
) {
	v.depth_--
	v.appendNewline()
	v.appendString("]")
}

func (v *formatter_) PreprocessMultilineParameters(
	multilineParameters ast.MultilineParametersLike,
) {
	v.appendString("(")
	v.depth_++
}

func (v *formatter_) PostprocessMultilineParameters(
	multilineParameters ast.MultilineParametersLike,
) {
	v.depth_--
	v.appendNewline()
	v.appendString(")")
}

func (v *formatter_) PreprocessMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
) {
	v.appendString("{")
	v.depth_++
}

func (v *formatter_) PostprocessMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
) {
	v.depth_--
	v.appendNewline()
	v.appendString("}")
}

func (v *formatter_) PreprocessMultilineValues(
	multilineValues ast.MultilineValuesLike,
) {
	v.appendString("[")
	v.depth_++
}

func (v *formatter_) PostprocessMultilineValues(
	multilineValues ast.MultilineValuesLike,
) {
	v.depth_--
	v.appendNewline()
	v.appendString("]")
}

func (v *formatter_) PreprocessNoAttributes(
	noAttributes ast.NoAttributesLike,
) {
	v.appendString("[:]")
}

func (v *formatter_) PreprocessNoStatements(
	noStatements ast.NoStatementsLike,
) {
	v.appendString("{")
	v.depth_++
	v.appendNewline()
	v.depth_--
	v.appendString("}")
}

func (v *formatter_) PreprocessNoValues(
	noValues ast.NoValuesLike,
) {
	v.appendString("[ ]")
}

func (v *formatter_) PreprocessNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
) {
	v.appendString("notarize ")
}

func (v *formatter_) ProcessNotarizeClauseSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" as ")
	}
}

func (v *formatter_) PostprocessNotice(
	notice ast.NoticeLike,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessOnClause(
	onClause ast.OnClauseLike,
) {
	v.appendString(" on ")
}

func (v *formatter_) PreprocessPlusEqual(
	plusEqual ast.PlusEqualLike,
) {
	v.appendString("+=")
}

func (v *formatter_) PreprocessPostClause(
	postClause ast.PostClauseLike,
) {
	v.appendString("post ")
}

func (v *formatter_) ProcessPostClauseSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" to ")
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessPrecedence(
	precedence ast.PrecedenceLike,
) {
	v.appendString("(")
}

func (v *formatter_) PostprocessPrecedence(
	precedence ast.PrecedenceLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessPredicate(
	predicate ast.PredicateLike,
	index uint,
	size uint,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessPredicateSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessPublishClause(
	publishClause ast.PublishClauseLike,
) {
	v.appendString("publish ")
}

func (v *formatter_) PreprocessReferent(
	referent ast.ReferentLike,
) {
	v.appendString("@")
}

func (v *formatter_) PreprocessRejectClause(
	rejectClause ast.RejectClauseLike,
) {
	v.appendString("reject ")
}

func (v *formatter_) PreprocessRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
) {
	v.appendString("retrieve ")
}

func (v *formatter_) ProcessRetrieveClauseSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" from ")
	}
}

func (v *formatter_) PreprocessReturnClause(
	returnClause ast.ReturnClauseLike,
) {
	v.appendString("return ")
}

func (v *formatter_) PreprocessSan(
	san ast.SanLike,
) {
	v.appendString("san")
}

func (v *formatter_) PreprocessSaveClause(
	saveClause ast.SaveClauseLike,
) {
	v.appendString("save ")
}

func (v *formatter_) ProcessSaveClauseSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" as ")
	}
}

func (v *formatter_) PreprocessSelectClause(
	selectClause ast.SelectClauseLike,
) {
	v.appendString("select ")
}

func (v *formatter_) PreprocessSlashEqual(
	slashEqual ast.SlashEqualLike,
) {
	v.appendString("/=")
}

func (v *formatter_) PreprocessStarEqual(
	starEqual ast.StarEqualLike,
) {
	v.appendString("*=")
}

func (v *formatter_) PreprocessStatementLine(
	statementLine ast.StatementLineLike,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessSubcomponentSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString("[")
	}
}

func (v *formatter_) PostprocessSubcomponent(
	subcomponent ast.SubcomponentLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessThrowClause(
	throwClause ast.ThrowClauseLike,
) {
	v.appendString("throw ")
}

func (v *formatter_) PreprocessWhileClause(
	whileClause ast.WhileClauseLike,
) {
	v.appendString("while ")
}

func (v *formatter_) ProcessWhileClauseSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" do ")
	}
}

func (v *formatter_) PreprocessWithClause(
	withClause ast.WithClauseLike,
) {
	v.appendString("with each ")
}

func (v *formatter_) ProcessWithClauseSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(" in ")
	case 2:
		v.appendString(" do ")
	}
}

func (v *formatter_) PreprocessXor(
	xor ast.XorLike,
) {
	v.appendString("xor")
}

const _indentation = "    "

// PROTECTED INTERFACE

// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var level uint
	for ; level < v.depth_; level++ {
		newline += _indentation
	}
	v.appendString(newline)
}

func (v *formatter_) appendString(
	string_ string,
) {
	v.result_.WriteString(string_)
}

func (v *formatter_) getResult() string {
	var result = v.result_.String()
	v.result_.Reset()
	return result
}

// Instance Structure

type formatter_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike
	depth_   uint
	result_  sts.Builder

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type formatterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func formatterClass() *formatterClass_ {
	return formatterClassReference_
}

var formatterClassReference_ = &formatterClass_{
	// Initialize the class constants.
}
