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

func (v *validator_) ProcessAnd(
	and string,
) {
	v.validateToken(and, AndToken)
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

func (v *validator_) ProcessIor(
	ior string,
) {
	v.validateToken(ior, IorToken)
}

func (v *validator_) ProcessIs(
	is string,
) {
	v.validateToken(is, IsToken)
}

func (v *validator_) ProcessLess(
	less string,
) {
	v.validateToken(less, LessToken)
}

func (v *validator_) ProcessMatches(
	matches string,
) {
	v.validateToken(matches, MatchesToken)
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

func (v *validator_) ProcessNot(
	not string,
) {
	v.validateToken(not, NotToken)
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

func (v *validator_) ProcessSan(
	san string,
) {
	v.validateToken(san, SanToken)
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

func (v *validator_) ProcessSlashSlash(
	slashSlash string,
) {
	v.validateToken(slashSlash, SlashSlashToken)
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

func (v *validator_) ProcessXor(
	xor string,
) {
	v.validateToken(xor, XorToken)
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

func (v *validator_) PreprocessAmplitude(
	amplitude ast.AmplitudeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAmplitudeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAmplitude(
	amplitude ast.AmplitudeLike,
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

func (v *validator_) PreprocessArithmetic(
	arithmetic ast.ArithmeticLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessArithmeticSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArithmetic(
	arithmetic ast.ArithmeticLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArithmeticOperation(
	arithmeticOperation ast.ArithmeticOperationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessArithmeticOperationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArithmeticOperation(
	arithmeticOperation ast.ArithmeticOperationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessArithmeticOperator(
	arithmeticOperator ast.ArithmeticOperatorLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessArithmeticOperatorSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessArithmeticOperator(
	arithmeticOperator ast.ArithmeticOperatorLike,
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

func (v *validator_) PreprocessAssociations(
	associations ast.AssociationsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAssociationsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAssociations(
	associations ast.AssociationsLike,
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

func (v *validator_) PreprocessAttribute(
	attribute ast.AttributeLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessAttributeSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessAttribute(
	attribute ast.AttributeLike,
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

func (v *validator_) PreprocessBase(
	base ast.BaseLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessBaseSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessBase(
	base ast.BaseLike,
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

func (v *validator_) PreprocessCompareOperator(
	compareOperator ast.CompareOperatorLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessCompareOperatorSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessCompareOperator(
	compareOperator ast.CompareOperatorLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessComparison(
	comparison ast.ComparisonLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessComparisonSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessComparison(
	comparison ast.ComparisonLike,
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

func (v *validator_) PreprocessConcatenation(
	concatenation ast.ConcatenationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConcatenationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConcatenation(
	concatenation ast.ConcatenationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessConcatenationOperation(
	concatenationOperation ast.ConcatenationOperationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessConcatenationOperationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessConcatenationOperation(
	concatenationOperation ast.ConcatenationOperationLike,
	index uint,
	size uint,
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

func (v *validator_) PreprocessDereference(
	dereference ast.DereferenceLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessDereferenceSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessDereference(
	dereference ast.DereferenceLike,
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

func (v *validator_) PreprocessExponential(
	exponential ast.ExponentialLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessExponentialSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessExponential(
	exponential ast.ExponentialLike,
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

func (v *validator_) PreprocessInference(
	inference ast.InferenceLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInferenceSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInference(
	inference ast.InferenceLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessInlineAssociations(
	inlineAssociations ast.InlineAssociationsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessInlineAssociationsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessInlineAssociations(
	inlineAssociations ast.InlineAssociationsLike,
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

func (v *validator_) PreprocessItems(
	items ast.ItemsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessItemsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessItems(
	items ast.ItemsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLabel(
	label ast.LabelLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLabelSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLabel(
	label ast.LabelLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLeftBracket(
	leftBracket ast.LeftBracketLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLeftBracketSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLeftBracket(
	leftBracket ast.LeftBracketLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLeftRound(
	leftRound ast.LeftRoundLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLeftRoundSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLeftRound(
	leftRound ast.LeftRoundLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessLeftSquare(
	leftSquare ast.LeftSquareLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLeftSquareSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLeftSquare(
	leftSquare ast.LeftSquareLike,
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

func (v *validator_) PreprocessLogicOperator(
	logicOperator ast.LogicOperatorLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLogicOperatorSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLogicOperator(
	logicOperator ast.LogicOperatorLike,
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

func (v *validator_) PreprocessLogicalOperation(
	logicalOperation ast.LogicalOperationLike,
	index uint,
	size uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessLogicalOperationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessLogicalOperation(
	logicalOperation ast.LogicalOperationLike,
	index uint,
	size uint,
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

func (v *validator_) PreprocessMultilineAssociations(
	multilineAssociations ast.MultilineAssociationsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessMultilineAssociationsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessMultilineAssociations(
	multilineAssociations ast.MultilineAssociationsLike,
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

func (v *validator_) PreprocessNoAssociations(
	noAssociations ast.NoAssociationsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessNoAssociationsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessNoAssociations(
	noAssociations ast.NoAssociationsLike,
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

func (v *validator_) PreprocessOperation(
	operation ast.OperationLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessOperationSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessOperation(
	operation ast.OperationLike,
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

func (v *validator_) PreprocessRightBracket(
	rightBracket ast.RightBracketLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRightBracketSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRightBracket(
	rightBracket ast.RightBracketLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRightRound(
	rightRound ast.RightRoundLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRightRoundSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRightRound(
	rightRound ast.RightRoundLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessRightSquare(
	rightSquare ast.RightSquareLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessRightSquareSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessRightSquare(
	rightSquare ast.RightSquareLike,
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

func (v *validator_) PreprocessStatements(
	statements ast.StatementsLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessStatementsSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessStatements(
	statements ast.StatementsLike,
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

func (v *validator_) PreprocessTextual(
	textual ast.TextualLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessTextualSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessTextual(
	textual ast.TextualLike,
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

func (v *validator_) PreprocessValue(
	value ast.ValueLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessValueSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessValue(
	value ast.ValueLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PreprocessValues(
	values ast.ValuesLike,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) ProcessValuesSlot(
	slot uint,
) {
	// TBD - Add any validation checks.
}

func (v *validator_) PostprocessValues(
	values ast.ValuesLike,
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
