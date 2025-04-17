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
	fmt "fmt"
	ast "github.com/bali-nebula/go-component-framework/v3/ast"
)

// CLASS INTERFACE

// Access Function

func ValidatorClass() ValidatorClassLike {
	return validatorClass()
}

// Constructor Methods

func (c *validatorClass_) Validator() ValidatorLike {
	var instance = &validator_{
		// Initialize the instance attributes.

		// Initialize the inherited aspects.
		Methodical: ProcessorClass().Processor(),
	}
	instance.visitor_ = VisitorClass().Visitor(instance)
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *validator_) GetClass() ValidatorClassLike {
	return validatorClass()
}

func (v *validator_) ValidateDocument(
	document ast.DocumentLike,
) {
	v.visitor_.VisitDocument(document)
}

// Methodical Methods

func (v *validator_) ProcessAmpersand(
	ampersand string,
) {
	v.validateToken(ampersand, AmpersandToken)
}

func (v *validator_) ProcessAngle(
	angle string,
) {
	v.validateToken(angle, AngleToken)
}

func (v *validator_) ProcessArrow(
	arrow string,
) {
	v.validateToken(arrow, ArrowToken)
}

func (v *validator_) ProcessBar(
	bar string,
) {
	v.validateToken(bar, BarToken)
}

func (v *validator_) ProcessBinary(
	binary string,
) {
	v.validateToken(binary, BinaryToken)
}

func (v *validator_) ProcessBoolean(
	boolean string,
) {
	v.validateToken(boolean, BooleanToken)
}

func (v *validator_) ProcessBytecode(
	bytecode string,
) {
	v.validateToken(bytecode, BytecodeToken)
}

func (v *validator_) ProcessCaret(
	caret string,
) {
	v.validateToken(caret, CaretToken)
}

func (v *validator_) ProcessCitation(
	citation string,
) {
	v.validateToken(citation, CitationToken)
}

func (v *validator_) ProcessColon(
	colon string,
) {
	v.validateToken(colon, ColonToken)
}

func (v *validator_) ProcessColonEqual(
	colonEqual string,
) {
	v.validateToken(colonEqual, ColonEqualToken)
}

func (v *validator_) ProcessComment(
	comment string,
) {
	v.validateToken(comment, CommentToken)
}

func (v *validator_) ProcessDash(
	dash string,
) {
	v.validateToken(dash, DashToken)
}

func (v *validator_) ProcessDashEqual(
	dashEqual string,
) {
	v.validateToken(dashEqual, DashEqualToken)
}

func (v *validator_) ProcessDefaultEqual(
	defaultEqual string,
) {
	v.validateToken(defaultEqual, DefaultEqualToken)
}

func (v *validator_) ProcessDot(
	dot string,
) {
	v.validateToken(dot, DotToken)
}

func (v *validator_) ProcessDouble(
	double string,
) {
	v.validateToken(double, DoubleToken)
}

func (v *validator_) ProcessDuration(
	duration string,
) {
	v.validateToken(duration, DurationToken)
}

func (v *validator_) ProcessEqual(
	equal string,
) {
	v.validateToken(equal, EqualToken)
}

func (v *validator_) ProcessIdentifier(
	identifier string,
) {
	v.validateToken(identifier, IdentifierToken)
}

func (v *validator_) ProcessLess(
	less string,
) {
	v.validateToken(less, LessToken)
}

func (v *validator_) ProcessMoment(
	moment string,
) {
	v.validateToken(moment, MomentToken)
}

func (v *validator_) ProcessMore(
	more string,
) {
	v.validateToken(more, MoreToken)
}

func (v *validator_) ProcessName(
	name string,
) {
	v.validateToken(name, NameToken)
}

func (v *validator_) ProcessNarrative(
	narrative string,
) {
	v.validateToken(narrative, NarrativeToken)
}

func (v *validator_) ProcessNewline(
	newline string,
) {
	v.validateToken(newline, NewlineToken)
}

func (v *validator_) ProcessNote(
	note string,
) {
	v.validateToken(note, NoteToken)
}

func (v *validator_) ProcessNumber(
	number string,
) {
	v.validateToken(number, NumberToken)
}

func (v *validator_) ProcessPattern(
	pattern string,
) {
	v.validateToken(pattern, PatternToken)
}

func (v *validator_) ProcessPercentage(
	percentage string,
) {
	v.validateToken(percentage, PercentageToken)
}

func (v *validator_) ProcessPlus(
	plus string,
) {
	v.validateToken(plus, PlusToken)
}

func (v *validator_) ProcessPlusEqual(
	plusEqual string,
) {
	v.validateToken(plusEqual, PlusEqualToken)
}

func (v *validator_) ProcessProbability(
	probability string,
) {
	v.validateToken(probability, ProbabilityToken)
}

func (v *validator_) ProcessQuote(
	quote string,
) {
	v.validateToken(quote, QuoteToken)
}

func (v *validator_) ProcessResource(
	resource string,
) {
	v.validateToken(resource, ResourceToken)
}

func (v *validator_) ProcessSlash(
	slash string,
) {
	v.validateToken(slash, SlashToken)
}

func (v *validator_) ProcessSlashEqual(
	slashEqual string,
) {
	v.validateToken(slashEqual, SlashEqualToken)
}

func (v *validator_) ProcessSnail(
	snail string,
) {
	v.validateToken(snail, SnailToken)
}

func (v *validator_) ProcessSpace(
	space string,
) {
	v.validateToken(space, SpaceToken)
}

func (v *validator_) ProcessStar(
	star string,
) {
	v.validateToken(star, StarToken)
}

func (v *validator_) ProcessStarEqual(
	starEqual string,
) {
	v.validateToken(starEqual, StarEqualToken)
}

func (v *validator_) ProcessSymbol(
	symbol string,
) {
	v.validateToken(symbol, SymbolToken)
}

func (v *validator_) ProcessTag(
	tag string,
) {
	v.validateToken(tag, TagToken)
}

func (v *validator_) ProcessVersion(
	version string,
) {
	v.validateToken(version, VersionToken)
}

func (v *validator_) PreprocessAcceptClause(
	acceptClause ast.AcceptClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAcceptClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAcceptClause(
	acceptClause ast.AcceptClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAdditionalArgumentSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalAssociation(
	additionalAssociation ast.AdditionalAssociationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAdditionalAssociationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalAssociation(
	additionalAssociation ast.AdditionalAssociationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalIndex(
	additionalIndex ast.AdditionalIndexLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAdditionalIndexSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalIndex(
	additionalIndex ast.AdditionalIndexLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalStatement(
	additionalStatement ast.AdditionalStatementLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAdditionalStatementSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalStatement(
	additionalStatement ast.AdditionalStatementLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAdditionalValueSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAnd(
	and ast.AndLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAndSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAnd(
	and ast.AndLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAnnotatedAssociation(
	annotatedAssociation ast.AnnotatedAssociationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAnnotatedAssociationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAnnotatedAssociation(
	annotatedAssociation ast.AnnotatedAssociationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAnnotatedStatement(
	annotatedStatement ast.AnnotatedStatementLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAnnotatedStatementSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAnnotatedStatement(
	annotatedStatement ast.AnnotatedStatementLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAnnotatedValue(
	annotatedValue ast.AnnotatedValueLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAnnotatedValueSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAnnotatedValue(
	annotatedValue ast.AnnotatedValueLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAnnotationLine(
	annotationLine ast.AnnotationLineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAnnotationLineSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAnnotationLine(
	annotationLine ast.AnnotationLineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArgument(
	argument ast.ArgumentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessArgumentSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArgument(
	argument ast.ArgumentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArguments(
	arguments ast.ArgumentsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessArgumentsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArguments(
	arguments ast.ArgumentsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAssign(
	assign ast.AssignLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAssignSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAssign(
	assign ast.AssignLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAssociation(
	association ast.AssociationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAssociationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAssociation(
	association ast.AssociationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessAtLevel(
	atLevel ast.AtLevelLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAtLevelSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAtLevel(
	atLevel ast.AtLevelLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessBag(
	bag ast.BagLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessBagSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessBag(
	bag ast.BagLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessBreakClause(
	breakClause ast.BreakClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessBreakClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessBreakClause(
	breakClause ast.BreakClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessCheckoutClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCitation(
	citation ast.CitationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessCitationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCitation(
	citation ast.CitationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCollection(
	collection ast.CollectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessCollectionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCollection(
	collection ast.CollectionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessComplement(
	complement ast.ComplementLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessComplementSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessComplement(
	complement ast.ComplementLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessComponent(
	component ast.ComponentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessComponentSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessComponent(
	component ast.ComponentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessCondition(
	condition ast.ConditionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConditionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCondition(
	condition ast.ConditionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessContinueClause(
	continueClause ast.ContinueClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessContinueClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessContinueClause(
	continueClause ast.ContinueClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDiscardClause(
	discardClause ast.DiscardClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessDiscardClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDiscardClause(
	discardClause ast.DiscardClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDoClause(
	doClause ast.DoClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessDoClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDoClause(
	doClause ast.DoClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDocument(
	document ast.DocumentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessDocumentSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDocument(
	document ast.DocumentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessDraft(
	draft ast.DraftLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessDraftSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDraft(
	draft ast.DraftLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessElement(
	element ast.ElementLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessElementSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessElement(
	element ast.ElementLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessEmptyLine(
	emptyLine ast.EmptyLineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessEmptyLineSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessEmptyLine(
	emptyLine ast.EmptyLineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessEntity(
	entity ast.EntityLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessEntitySlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessEntity(
	entity ast.EntityLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessEvent(
	event ast.EventLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessEventSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessEvent(
	event ast.EventLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExInclusiveRange(
	exInclusiveRange ast.ExInclusiveRangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessExInclusiveRangeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExInclusiveRange(
	exInclusiveRange ast.ExInclusiveRangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessException(
	exception ast.ExceptionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessExceptionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessException(
	exception ast.ExceptionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExclusiveRange(
	exclusiveRange ast.ExclusiveRangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessExclusiveRangeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExclusiveRange(
	exclusiveRange ast.ExclusiveRangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessExpression(
	expression ast.ExpressionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessExpressionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExpression(
	expression ast.ExpressionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFailure(
	failure ast.FailureLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessFailureSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFailure(
	failure ast.FailureLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFlow(
	flow ast.FlowLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessFlowSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFlow(
	flow ast.FlowLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessFunction(
	function ast.FunctionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessFunctionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessFunction(
	function ast.FunctionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIfClause(
	ifClause ast.IfClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessIfClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessIfClause(
	ifClause ast.IfClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInExclusiveRange(
	inExclusiveRange ast.InExclusiveRangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInExclusiveRangeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInExclusiveRange(
	inExclusiveRange ast.InExclusiveRangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInclusiveRange(
	inclusiveRange ast.InclusiveRangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInclusiveRangeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInclusiveRange(
	inclusiveRange ast.InclusiveRangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIndex(
	index ast.IndexLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessIndexSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessIndex(
	index ast.IndexLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIndices(
	indices ast.IndicesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessIndicesSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessIndices(
	indices ast.IndicesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIndirect(
	indirect ast.IndirectLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessIndirectSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessIndirect(
	indirect ast.IndirectLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInduction(
	induction ast.InductionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInductionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInduction(
	induction ast.InductionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInlineAttributes(
	inlineAttributes ast.InlineAttributesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInlineAttributesSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInlineAttributes(
	inlineAttributes ast.InlineAttributesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInlineParameters(
	inlineParameters ast.InlineParametersLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInlineParametersSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInlineParameters(
	inlineParameters ast.InlineParametersLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInlineStatements(
	inlineStatements ast.InlineStatementsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInlineStatementsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInlineStatements(
	inlineStatements ast.InlineStatementsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInlineValues(
	inlineValues ast.InlineValuesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInlineValuesSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInlineValues(
	inlineValues ast.InlineValuesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInverse(
	inverse ast.InverseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInverseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInverse(
	inverse ast.InverseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInversion(
	inversion ast.InversionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInversionSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInversion(
	inversion ast.InversionLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInvocation(
	invocation ast.InvocationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInvocationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInvocation(
	invocation ast.InvocationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIor(
	ior ast.IorLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessIorSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessIor(
	ior ast.IorLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessIs(
	is ast.IsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessIsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessIs(
	is ast.IsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessItem(
	item ast.ItemLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessItemSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessItem(
	item ast.ItemLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLetClause(
	letClause ast.LetClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLetClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLetClause(
	letClause ast.LetClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLogical(
	logical ast.LogicalLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLogicalSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLogical(
	logical ast.LogicalLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMagnitude(
	magnitude ast.MagnitudeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMagnitudeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMagnitude(
	magnitude ast.MagnitudeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMainClause(
	mainClause ast.MainClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMainClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMainClause(
	mainClause ast.MainClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMatchHandler(
	matchHandler ast.MatchHandlerLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMatchHandlerSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMatchHandler(
	matchHandler ast.MatchHandlerLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMatches(
	matches ast.MatchesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMatchesSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMatches(
	matches ast.MatchesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMessage(
	message ast.MessageLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMessageSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMessage(
	message ast.MessageLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMessaging(
	messaging ast.MessagingLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMessagingSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMessaging(
	messaging ast.MessagingLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMethod(
	method ast.MethodLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMethodSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMethod(
	method ast.MethodLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultilineAttributes(
	multilineAttributes ast.MultilineAttributesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMultilineAttributesSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultilineAttributes(
	multilineAttributes ast.MultilineAttributesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultilineParameters(
	multilineParameters ast.MultilineParametersLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMultilineParametersSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultilineParameters(
	multilineParameters ast.MultilineParametersLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMultilineStatementsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessMultilineValues(
	multilineValues ast.MultilineValuesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMultilineValuesSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultilineValues(
	multilineValues ast.MultilineValuesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNoAttributes(
	noAttributes ast.NoAttributesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessNoAttributesSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNoAttributes(
	noAttributes ast.NoAttributesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNoStatements(
	noStatements ast.NoStatementsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessNoStatementsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNoStatements(
	noStatements ast.NoStatementsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNoValues(
	noValues ast.NoValuesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessNoValuesSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNoValues(
	noValues ast.NoValuesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessNotarizeClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessNumerical(
	numerical ast.NumericalLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessNumericalSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNumerical(
	numerical ast.NumericalLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessOnClause(
	onClause ast.OnClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessOnClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessOnClause(
	onClause ast.OnClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessOperator(
	operator ast.OperatorLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessOperatorSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessOperator(
	operator ast.OperatorLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessParameters(
	parameters ast.ParametersLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessParametersSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessParameters(
	parameters ast.ParametersLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPostClause(
	postClause ast.PostClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPostClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPostClause(
	postClause ast.PostClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrecedence(
	precedence ast.PrecedenceLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPrecedenceSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPrecedence(
	precedence ast.PrecedenceLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPredicate(
	predicate ast.PredicateLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPredicateSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPredicate(
	predicate ast.PredicateLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPrimitive(
	primitive ast.PrimitiveLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPrimitiveSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPrimitive(
	primitive ast.PrimitiveLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessProcedure(
	procedure ast.ProcedureLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessProcedureSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessProcedure(
	procedure ast.ProcedureLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessPublishClause(
	publishClause ast.PublishClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessPublishClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessPublishClause(
	publishClause ast.PublishClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRange(
	range_ ast.RangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRangeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRange(
	range_ ast.RangeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRecipient(
	recipient ast.RecipientLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRecipientSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRecipient(
	recipient ast.RecipientLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessReferent(
	referent ast.ReferentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessReferentSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessReferent(
	referent ast.ReferentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRejectClause(
	rejectClause ast.RejectClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRejectClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRejectClause(
	rejectClause ast.RejectClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRepository(
	repository ast.RepositoryLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRepositorySlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRepository(
	repository ast.RepositoryLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessResult(
	result ast.ResultLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessResultSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessResult(
	result ast.ResultLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRetrieveClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessReturnClause(
	returnClause ast.ReturnClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessReturnClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessReturnClause(
	returnClause ast.ReturnClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSan(
	san ast.SanLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessSanSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSan(
	san ast.SanLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSaveClause(
	saveClause ast.SaveClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessSaveClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSaveClause(
	saveClause ast.SaveClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSelectClause(
	selectClause ast.SelectClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessSelectClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSelectClause(
	selectClause ast.SelectClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSequence(
	sequence ast.SequenceLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessSequenceSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSequence(
	sequence ast.SequenceLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessStatement(
	statement ast.StatementLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessStatementSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessStatement(
	statement ast.StatementLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessStatementLine(
	statementLine ast.StatementLineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessStatementLineSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessStatementLine(
	statementLine ast.StatementLineLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessString(
	string_ ast.StringLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessStringSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessString(
	string_ ast.StringLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSubcomponent(
	subcomponent ast.SubcomponentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessSubcomponentSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSubcomponent(
	subcomponent ast.SubcomponentLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessSubject(
	subject ast.SubjectLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessSubjectSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessSubject(
	subject ast.SubjectLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTarget(
	target ast.TargetLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessTargetSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTarget(
	target ast.TargetLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessTemplate(
	template ast.TemplateLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessTemplateSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTemplate(
	template ast.TemplateLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessThreading(
	threading ast.ThreadingLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessThreadingSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessThreading(
	threading ast.ThreadingLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessThrowClause(
	throwClause ast.ThrowClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessThrowClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessThrowClause(
	throwClause ast.ThrowClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessVariable(
	variable ast.VariableLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessVariableSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessVariable(
	variable ast.VariableLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessWhileClause(
	whileClause ast.WhileClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessWhileClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessWhileClause(
	whileClause ast.WhileClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessWithClause(
	withClause ast.WithClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessWithClauseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessWithClause(
	withClause ast.WithClauseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessXor(
	xor ast.XorLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessXorSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessXor(
	xor ast.XorLike,
) {
	// TBD - Add any validation checks.
}

// PROTECTED INTERFACE

// Private Methods

func (v *validator_) validateToken(
	tokenValue string,
	tokenType TokenType,
) {
	var scannerClass = ScannerClass()
	if !scannerClass.MatchesType(tokenValue, tokenType) {
		var message = fmt.Sprintf(
			"The following token value is not of type %v: %v",
			scannerClass.FormatType(tokenType),
			tokenValue,
		)
		panic(message)
	}
}

// Instance Structure

type validator_ struct {
	// Declare the instance attributes.
	visitor_ VisitorLike

	// Declare the inherited aspects.
	Methodical
}

// Class Structure

type validatorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func validatorClass() *validatorClass_ {
	return validatorClassReference_
}

var validatorClassReference_ = &validatorClass_{
	// Initialize the class constants.
}
