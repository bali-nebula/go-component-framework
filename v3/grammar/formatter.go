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

func (v *formatter_) ProcessBar(
	bar string,
) {
	v.appendString(bar)
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

func (v *formatter_) ProcessColon(
	colon string,
) {
	v.appendString(colon)
}

func (v *formatter_) ProcessColonEqual(
	colonEqual string,
) {
	v.appendString(colonEqual)
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

func (v *formatter_) ProcessDashEqual(
	dashEqual string,
) {
	v.appendString(dashEqual)
}

func (v *formatter_) ProcessDefaultEqual(
	defaultEqual string,
) {
	v.appendString(defaultEqual)
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

func (v *formatter_) ProcessIdentifier(
	identifier string,
) {
	v.appendString(identifier)
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

func (v *formatter_) ProcessPlusEqual(
	plusEqual string,
) {
	v.appendString(plusEqual)
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

func (v *formatter_) ProcessSlashEqual(
	slashEqual string,
) {
	v.appendString(slashEqual)
}

func (v *formatter_) ProcessSnail(
	snail string,
) {
	v.appendString(snail)
}

func (v *formatter_) ProcessSpace(
	space string,
) {
	v.appendString(space)
}

func (v *formatter_) ProcessStar(
	star string,
) {
	v.appendString(star)
}

func (v *formatter_) ProcessStarEqual(
	starEqual string,
) {
	v.appendString(starEqual)
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

func (v *formatter_) ProcessAcceptClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessAdditionalArgumentSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(", ")
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessAdditionalAssociationSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(", ")
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessAdditionalIndexSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(", ")
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessAdditionalStatementSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString("; ")
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessAdditionalValueSlot(
	slot uint,
) {
	switch slot {
	case 1:
		v.appendString(", ")
	default:
		v.appendString(" ")
	}
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

func (v *formatter_) PreprocessAnnotationLine(
	annotationLine ast.AnnotationLineLike,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessAssociation(
	association ast.AssociationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAssociationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAssociation(
	association ast.AssociationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAtLevel(
	atLevel ast.AtLevelLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAtLevelSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAtLevel(
	atLevel ast.AtLevelLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessBag(
	bag ast.BagLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessBagSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessBag(
	bag ast.BagLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessBreakClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	v.appendString(" ")
}

func (v *formatter_) ProcessCitationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessCitation(
	citation ast.CitationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessCollection(
	collection ast.CollectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessCollectionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessCollection(
	collection ast.CollectionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessComplement(
	complement ast.ComplementLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessComplementSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessComplement(
	complement ast.ComplementLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessComponent(
	component ast.ComponentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessComponentSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessComponent(
	component ast.ComponentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessCondition(
	condition ast.ConditionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessConditionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessCondition(
	condition ast.ConditionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessContinueClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessDiscardClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessDoClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessDocument(
	document ast.DocumentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessDocumentSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessDocument(
	document ast.DocumentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessDraft(
	draft ast.DraftLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessDraftSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessDraft(
	draft ast.DraftLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessElement(
	element ast.ElementLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessElementSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessElement(
	element ast.ElementLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessEmptyLine(
	emptyLine ast.EmptyLineLike,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessEntity(
	entity ast.EntityLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessEntitySlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessEntity(
	entity ast.EntityLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessEvent(
	event ast.EventLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessEventSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessEvent(
	event ast.EventLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInclusiveRange() {
	v.appendString("[")
}

func (v *formatter_) PostprocessInclusiveRange() {
	v.appendString("]")
}

func (v *formatter_) PreprocessExclusiveRange() {
	v.appendString("(")
}

func (v *formatter_) PostprocessExclusiveRange() {
	v.appendString(")")
}

func (v *formatter_) ProcessRangeSlot() {
	v.appendString("..")
}

func (v *formatter_) PreprocessExInclusiveAngles(
	exInclusiveAngles ast.ExInclusiveAnglesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveAnglesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveAngles(
	exInclusiveAngles ast.ExInclusiveAnglesLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusiveCitations(
	exInclusiveCitations ast.ExInclusiveCitationsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveCitationsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveCitations(
	exInclusiveCitations ast.ExInclusiveCitationsLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusiveDurations(
	exInclusiveDurations ast.ExInclusiveDurationsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveDurationsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveDurations(
	exInclusiveDurations ast.ExInclusiveDurationsLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusiveMoments(
	exInclusiveMoments ast.ExInclusiveMomentsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveMomentsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveMoments(
	exInclusiveMoments ast.ExInclusiveMomentsLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusiveNames(
	exInclusiveNames ast.ExInclusiveNamesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveNamesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveNames(
	exInclusiveNames ast.ExInclusiveNamesLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusiveNumbers(
	exInclusiveNumbers ast.ExInclusiveNumbersLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveNumbersSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveNumbers(
	exInclusiveNumbers ast.ExInclusiveNumbersLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusivePercentages(
	exInclusivePercentages ast.ExInclusivePercentagesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusivePercentagesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusivePercentages(
	exInclusivePercentages ast.ExInclusivePercentagesLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusiveProbabilities(
	exInclusiveProbabilities ast.ExInclusiveProbabilitiesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveProbabilitiesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveProbabilities(
	exInclusiveProbabilities ast.ExInclusiveProbabilitiesLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusiveQuotes(
	exInclusiveQuotes ast.ExInclusiveQuotesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveQuotesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveQuotes(
	exInclusiveQuotes ast.ExInclusiveQuotesLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusiveSymbols(
	exInclusiveSymbols ast.ExInclusiveSymbolsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveSymbolsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveSymbols(
	exInclusiveSymbols ast.ExInclusiveSymbolsLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessExInclusiveVersions(
	exInclusiveVersions ast.ExInclusiveVersionsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExInclusiveVersionsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExInclusiveVersions(
	exInclusiveVersions ast.ExInclusiveVersionsLike,
) {
	v.appendString("]")
}

func (v *formatter_) PreprocessException(
	exception ast.ExceptionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessExceptionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessException(
	exception ast.ExceptionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessExclusiveAngles(
	exclusiveAngles ast.ExclusiveAnglesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveAnglesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveAngles(
	exclusiveAngles ast.ExclusiveAnglesLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusiveCitations(
	exclusiveCitations ast.ExclusiveCitationsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveCitationsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveCitations(
	exclusiveCitations ast.ExclusiveCitationsLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusiveDurations(
	exclusiveDurations ast.ExclusiveDurationsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveDurationsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveDurations(
	exclusiveDurations ast.ExclusiveDurationsLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusiveMoments(
	exclusiveMoments ast.ExclusiveMomentsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveMomentsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveMoments(
	exclusiveMoments ast.ExclusiveMomentsLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusiveNames(
	exclusiveNames ast.ExclusiveNamesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveNamesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveNames(
	exclusiveNames ast.ExclusiveNamesLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusiveNumbers(
	exclusiveNumbers ast.ExclusiveNumbersLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveNumbersSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveNumbers(
	exclusiveNumbers ast.ExclusiveNumbersLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusivePercentages(
	exclusivePercentages ast.ExclusivePercentagesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusivePercentagesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusivePercentages(
	exclusivePercentages ast.ExclusivePercentagesLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusiveProbabilities(
	exclusiveProbabilities ast.ExclusiveProbabilitiesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveProbabilitiesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveProbabilities(
	exclusiveProbabilities ast.ExclusiveProbabilitiesLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusiveQuotes(
	exclusiveQuotes ast.ExclusiveQuotesLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveQuotesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveQuotes(
	exclusiveQuotes ast.ExclusiveQuotesLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusiveSymbols(
	exclusiveSymbols ast.ExclusiveSymbolsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveSymbolsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveSymbols(
	exclusiveSymbols ast.ExclusiveSymbolsLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExclusiveVersions(
	exclusiveVersions ast.ExclusiveVersionsLike,
) {
	v.PreprocessExclusiveRange()
}

func (v *formatter_) ProcessExclusiveVersionsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessExclusiveVersions(
	exclusiveVersions ast.ExclusiveVersionsLike,
) {
	v.PostprocessExclusiveRange()
}

func (v *formatter_) PreprocessExpression(
	expression ast.ExpressionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessExpressionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessExpression(
	expression ast.ExpressionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessFailure(
	failure ast.FailureLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessFailureSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessFailure(
	failure ast.FailureLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessFlow(
	flow ast.FlowLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessFlowSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessFlow(
	flow ast.FlowLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessFunction(
	function ast.FunctionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessFunctionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessFunction(
	function ast.FunctionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessIfClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessInExclusiveAngles(
	inExclusiveAngles ast.InExclusiveAnglesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveAnglesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveAngles(
	inExclusiveAngles ast.InExclusiveAnglesLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusiveCitations(
	inExclusiveCitations ast.InExclusiveCitationsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveCitationsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveCitations(
	inExclusiveCitations ast.InExclusiveCitationsLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusiveDurations(
	inExclusiveDurations ast.InExclusiveDurationsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveDurationsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveDurations(
	inExclusiveDurations ast.InExclusiveDurationsLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusiveMoments(
	inExclusiveMoments ast.InExclusiveMomentsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveMomentsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveMoments(
	inExclusiveMoments ast.InExclusiveMomentsLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusiveNames(
	inExclusiveNames ast.InExclusiveNamesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveNamesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveNames(
	inExclusiveNames ast.InExclusiveNamesLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusiveNumbers(
	inExclusiveNumbers ast.InExclusiveNumbersLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveNumbersSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveNumbers(
	inExclusiveNumbers ast.InExclusiveNumbersLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusivePercentages(
	inExclusivePercentages ast.InExclusivePercentagesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusivePercentagesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusivePercentages(
	inExclusivePercentages ast.InExclusivePercentagesLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusiveProbabilities(
	inExclusiveProbabilities ast.InExclusiveProbabilitiesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveProbabilitiesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveProbabilities(
	inExclusiveProbabilities ast.InExclusiveProbabilitiesLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusiveQuotes(
	inExclusiveQuotes ast.InExclusiveQuotesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveQuotesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveQuotes(
	inExclusiveQuotes ast.InExclusiveQuotesLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusiveSymbols(
	inExclusiveSymbols ast.InExclusiveSymbolsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveSymbolsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveSymbols(
	inExclusiveSymbols ast.InExclusiveSymbolsLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInExclusiveVersions(
	inExclusiveVersions ast.InExclusiveVersionsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInExclusiveVersionsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInExclusiveVersions(
	inExclusiveVersions ast.InExclusiveVersionsLike,
) {
	v.appendString(")")
}

func (v *formatter_) PreprocessInclusiveAngles(
	inclusiveAngles ast.InclusiveAnglesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveAnglesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveAngles(
	inclusiveAngles ast.InclusiveAnglesLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusiveCitations(
	inclusiveCitations ast.InclusiveCitationsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveCitationsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveCitations(
	inclusiveCitations ast.InclusiveCitationsLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusiveDurations(
	inclusiveDurations ast.InclusiveDurationsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveDurationsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveDurations(
	inclusiveDurations ast.InclusiveDurationsLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusiveMoments(
	inclusiveMoments ast.InclusiveMomentsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveMomentsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveMoments(
	inclusiveMoments ast.InclusiveMomentsLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusiveNames(
	inclusiveNames ast.InclusiveNamesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveNamesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveNames(
	inclusiveNames ast.InclusiveNamesLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusiveNumbers(
	inclusiveNumbers ast.InclusiveNumbersLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveNumbersSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveNumbers(
	inclusiveNumbers ast.InclusiveNumbersLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusivePercentages(
	inclusivePercentages ast.InclusivePercentagesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusivePercentagesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusivePercentages(
	inclusivePercentages ast.InclusivePercentagesLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusiveProbabilities(
	inclusiveProbabilities ast.InclusiveProbabilitiesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveProbabilitiesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveProbabilities(
	inclusiveProbabilities ast.InclusiveProbabilitiesLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusiveQuotes(
	inclusiveQuotes ast.InclusiveQuotesLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveQuotesSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveQuotes(
	inclusiveQuotes ast.InclusiveQuotesLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusiveSymbols(
	inclusiveSymbols ast.InclusiveSymbolsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveSymbolsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveSymbols(
	inclusiveSymbols ast.InclusiveSymbolsLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessInclusiveVersions(
	inclusiveVersions ast.InclusiveVersionsLike,
) {
	v.PreprocessInclusiveRange()
}

func (v *formatter_) ProcessInclusiveVersionsSlot(
	slot uint,
) {
	v.ProcessRangeSlot()
}

func (v *formatter_) PostprocessInclusiveVersions(
	inclusiveVersions ast.InclusiveVersionsLike,
) {
	v.PostprocessInclusiveRange()
}

func (v *formatter_) PreprocessIndex(
	index ast.IndexLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessIndexSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessIndex(
	index ast.IndexLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessIndices(
	indices ast.IndicesLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessIndicesSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessIndices(
	indices ast.IndicesLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessIndirect(
	indirect ast.IndirectLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessIndirectSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessIndirect(
	indirect ast.IndirectLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInduction(
	induction ast.InductionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInductionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInduction(
	induction ast.InductionLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessInverse(
	inverse ast.InverseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInverseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInverse(
	inverse ast.InverseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInversion(
	inversion ast.InversionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInversionSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInversion(
	inversion ast.InversionLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInvocation(
	invocation ast.InvocationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInvocationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInvocation(
	invocation ast.InvocationLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessItem(
	item ast.ItemLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessItemSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessItem(
	item ast.ItemLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessLetClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessLogical(
	logical ast.LogicalLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessLogicalSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessLogical(
	logical ast.LogicalLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMagnitude(
	magnitude ast.MagnitudeLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMagnitudeSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMagnitude(
	magnitude ast.MagnitudeLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMatchHandler(
	matchHandler ast.MatchHandlerLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMatchHandlerSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMatchHandler(
	matchHandler ast.MatchHandlerLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMatches(
	matches ast.MatchesLike,
) {
	v.appendString("matches")
}

func (v *formatter_) PreprocessMessage(
	message ast.MessageLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMessageSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMessage(
	message ast.MessageLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMessaging(
	messaging ast.MessagingLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMessagingSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMessaging(
	messaging ast.MessagingLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMethod(
	method ast.MethodLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMethodSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMethod(
	method ast.MethodLike,
) {
	// TBD - Add formatting of the delimited rule.
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
	v.appendString("]")
}

func (v *formatter_) PreprocessNoAttributes(
	noAttributes ast.NoAttributesLike,
) {
	v.appendString("[")
}

func (v *formatter_) PostprocessNoAttributes(
	noAttributes ast.NoAttributesLike,
) {
	v.appendString("]")
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

func (v *formatter_) ProcessNotarizeClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessNumerical(
	numerical ast.NumericalLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessNumericalSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessNumerical(
	numerical ast.NumericalLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessOnClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessPostClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessPrecedence(
	precedence ast.PrecedenceLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPrecedenceSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPrecedence(
	precedence ast.PrecedenceLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPredicate(
	predicate ast.PredicateLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPredicateSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPredicate(
	predicate ast.PredicateLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessProcedure(
	procedure ast.ProcedureLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessProcedureSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessProcedure(
	procedure ast.ProcedureLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPublishClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessRecipient(
	recipient ast.RecipientLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessRecipientSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessRecipient(
	recipient ast.RecipientLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessReferent(
	referent ast.ReferentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessReferentSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessReferent(
	referent ast.ReferentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessRejectClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessRepository(
	repository ast.RepositoryLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessRepositorySlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessRepository(
	repository ast.RepositoryLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessResult(
	result ast.ResultLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessResultSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessResult(
	result ast.ResultLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessRetrieveClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessReturnClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessSan(
	san ast.SanLike,
) {
	v.appendString("san")
}

func (v *formatter_) ProcessSaveClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessSelectClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessSequence(
	sequence ast.SequenceLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessSequenceSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessSequence(
	sequence ast.SequenceLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessStatement(
	statement ast.StatementLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessStatementSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessStatement(
	statement ast.StatementLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessStatementLine(
	statementLine ast.StatementLineLike,
) {
	v.appendNewline()
}

func (v *formatter_) PreprocessString(
	string_ ast.StringLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessStringSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessString(
	string_ ast.StringLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessSubcomponent(
	subcomponent ast.SubcomponentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessSubcomponentSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessSubcomponent(
	subcomponent ast.SubcomponentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessSubject(
	subject ast.SubjectLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessSubjectSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessSubject(
	subject ast.SubjectLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessTarget(
	target ast.TargetLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessTargetSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessTarget(
	target ast.TargetLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessTemplate(
	template ast.TemplateLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessTemplateSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessTemplate(
	template ast.TemplateLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessThreading(
	threading ast.ThreadingLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessThreadingSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessThreading(
	threading ast.ThreadingLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessThrowClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessVariable(
	variable ast.VariableLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessVariableSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessVariable(
	variable ast.VariableLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessWhileClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) ProcessWithClauseSlot(
	slot uint,
) {
	switch slot {
	default:
		v.appendString(" ")
	}
}

func (v *formatter_) PreprocessXor(
	xor ast.XorLike,
) {
	v.appendString("xor")
}

// PROTECTED INTERFACE

// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "    "
	var level uint
	for ; level < v.depth_; level++ {
		newline += indentation
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
