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

func (v *formatter_) ProcessAngle(
	angle string,
) {
	v.appendString(angle)
}

func (v *formatter_) ProcessArithmeticOperator(
	arithmeticOperator string,
) {
	v.appendString(arithmeticOperator)
}

func (v *formatter_) ProcessAssignmentOperator(
	assignmentOperator string,
) {
	v.appendString(assignmentOperator)
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

func (v *formatter_) ProcessClose(
	close_ string,
) {
	v.appendString(close_)
}

func (v *formatter_) ProcessComment(
	comment string,
) {
	v.appendString(comment)
}

func (v *formatter_) ProcessComparisonOperator(
	comparisonOperator string,
) {
	v.appendString(comparisonOperator)
}

func (v *formatter_) ProcessDuration(
	duration string,
) {
	v.appendString(duration)
}

func (v *formatter_) ProcessIdentifier(
	identifier string,
) {
	v.appendString(identifier)
}

func (v *formatter_) ProcessInversionOperator(
	inversionOperator string,
) {
	v.appendString(inversionOperator)
}

func (v *formatter_) ProcessInvocationOperator(
	invocationOperator string,
) {
	v.appendString(invocationOperator)
}

func (v *formatter_) ProcessLogicalOperator(
	logicalOperator string,
) {
	v.appendString(logicalOperator)
}

func (v *formatter_) ProcessMoment(
	moment string,
) {
	v.appendString(moment)
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

func (v *formatter_) ProcessNewline(
	newline string,
) {
	v.appendNewline()
}

func (v *formatter_) ProcessNote(
	note string,
) {
	v.appendString(note)
}

func (v *formatter_) ProcessNumber(
	number string,
) {
	v.appendString(number)
}

func (v *formatter_) ProcessOpen(
	open string,
) {
	v.appendString(open)
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

func (v *formatter_) ProcessSpace(
	space string,
) {
	v.appendString(space)
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
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAcceptClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAcceptClause(
	acceptClause ast.AcceptClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAdditionalArgumentSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAdditionalAssociation(
	additionalAssociation ast.AdditionalAssociationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAdditionalAssociationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalAssociation(
	additionalAssociation ast.AdditionalAssociationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAdditionalIndex(
	additionalIndex ast.AdditionalIndexLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAdditionalIndexSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalIndex(
	additionalIndex ast.AdditionalIndexLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAdditionalParameter(
	additionalParameter ast.AdditionalParameterLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAdditionalParameterSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalParameter(
	additionalParameter ast.AdditionalParameterLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAdditionalStatement(
	additionalStatement ast.AdditionalStatementLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAdditionalStatementSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalStatement(
	additionalStatement ast.AdditionalStatementLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAdditionalValueSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAdditionalValue(
	additionalValue ast.AdditionalValueLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAnnotatedAssociation(
	annotatedAssociation ast.AnnotatedAssociationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAnnotatedAssociationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAnnotatedAssociation(
	annotatedAssociation ast.AnnotatedAssociationLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAnnotatedParameter(
	annotatedParameter ast.AnnotatedParameterLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAnnotatedParameterSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAnnotatedParameter(
	annotatedParameter ast.AnnotatedParameterLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAnnotatedStatement(
	annotatedStatement ast.AnnotatedStatementLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAnnotatedStatementSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAnnotatedStatement(
	annotatedStatement ast.AnnotatedStatementLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessAnnotatedValue(
	annotatedValue ast.AnnotatedValueLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAnnotatedValueSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAnnotatedValue(
	annotatedValue ast.AnnotatedValueLike,
	index uint,
	size uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessArgument(
	argument ast.ArgumentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessArgumentSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessArgument(
	argument ast.ArgumentLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessArguments(
	arguments ast.ArgumentsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessArgumentsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessArguments(
	arguments ast.ArgumentsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessArithmetic(
	arithmetic ast.ArithmeticLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessArithmeticSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessArithmetic(
	arithmetic ast.ArithmeticLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessAssociations(
	associations ast.AssociationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAssociationsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAssociations(
	associations ast.AssociationsLike,
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

func (v *formatter_) PreprocessAttribute(
	attribute ast.AttributeLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessAttributeSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessAttribute(
	attribute ast.AttributeLike,
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

func (v *formatter_) PreprocessBreakClause(
	breakClause ast.BreakClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessBreakClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessBreakClause(
	breakClause ast.BreakClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessChaining(
	chaining ast.ChainingLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessChainingSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessChaining(
	chaining ast.ChainingLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessCheckoutClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessCitation(
	citation ast.CitationLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessComparison(
	comparison ast.ComparisonLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessComparisonSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessComparison(
	comparison ast.ComparisonLike,
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

func (v *formatter_) PreprocessComposite(
	composite ast.CompositeLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessCompositeSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessComposite(
	composite ast.CompositeLike,
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

func (v *formatter_) PreprocessContext(
	context ast.ContextLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessContextSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessContext(
	context ast.ContextLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessContinueClause(
	continueClause ast.ContinueClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessContinueClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessContinueClause(
	continueClause ast.ContinueClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessDereference(
	dereference ast.DereferenceLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessDereferenceSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessDereference(
	dereference ast.DereferenceLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessDiscardClause(
	discardClause ast.DiscardClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessDiscardClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessDiscardClause(
	discardClause ast.DiscardClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessDoClause(
	doClause ast.DoClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessDoClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessDoClause(
	doClause ast.DoClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessEmptyAssociations(
	emptyAssociations ast.EmptyAssociationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessEmptyAssociationsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessEmptyAssociations(
	emptyAssociations ast.EmptyAssociationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessEmptyStatements(
	emptyStatements ast.EmptyStatementsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessEmptyStatementsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessEmptyStatements(
	emptyStatements ast.EmptyStatementsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessEmptyValues(
	emptyValues ast.EmptyValuesLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessEmptyValuesSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessEmptyValues(
	emptyValues ast.EmptyValuesLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessExponential(
	exponential ast.ExponentialLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessExponentialSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessExponential(
	exponential ast.ExponentialLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessIfClause(
	ifClause ast.IfClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessIfClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessIfClause(
	ifClause ast.IfClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessInlineAssociations(
	inlineAssociations ast.InlineAssociationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInlineAssociationsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInlineAssociations(
	inlineAssociations ast.InlineAssociationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInlineParameters(
	inlineParameters ast.InlineParametersLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInlineParametersSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInlineParameters(
	inlineParameters ast.InlineParametersLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInlineStatements(
	inlineStatements ast.InlineStatementsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInlineStatementsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInlineStatements(
	inlineStatements ast.InlineStatementsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessInlineValues(
	inlineValues ast.InlineValuesLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessInlineValuesSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessInlineValues(
	inlineValues ast.InlineValuesLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessIntrinsic(
	intrinsic ast.IntrinsicLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessIntrinsicSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessIntrinsic(
	intrinsic ast.IntrinsicLike,
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

func (v *formatter_) PreprocessItems(
	items ast.ItemsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessItemsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessItems(
	items ast.ItemsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessKey(
	key ast.KeyLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessKeySlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessKey(
	key ast.KeyLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessLabel(
	label ast.LabelLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessLabelSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessLabel(
	label ast.LabelLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessLetClause(
	letClause ast.LetClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessLetClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessLetClause(
	letClause ast.LetClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessMainClause(
	mainClause ast.MainClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMainClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMainClause(
	mainClause ast.MainClauseLike,
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

func (v *formatter_) PreprocessMultilineAssociations(
	multilineAssociations ast.MultilineAssociationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMultilineAssociationsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMultilineAssociations(
	multilineAssociations ast.MultilineAssociationsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMultilineParameters(
	multilineParameters ast.MultilineParametersLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMultilineParametersSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMultilineParameters(
	multilineParameters ast.MultilineParametersLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMultilineStatementsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessMultilineValues(
	multilineValues ast.MultilineValuesLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessMultilineValuesSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessMultilineValues(
	multilineValues ast.MultilineValuesLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessNotarizeClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessOnClause(
	onClause ast.OnClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessOnClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessOnClause(
	onClause ast.OnClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessOperation(
	operation ast.OperationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessOperationSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessOperation(
	operation ast.OperationLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessParameter(
	parameter ast.ParameterLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessParameterSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessParameter(
	parameter ast.ParameterLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessParameters(
	parameters ast.ParametersLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessParametersSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessParameters(
	parameters ast.ParametersLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessPostClause(
	postClause ast.PostClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPostClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPostClause(
	postClause ast.PostClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessPrimitive(
	primitive ast.PrimitiveLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPrimitiveSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPrimitive(
	primitive ast.PrimitiveLike,
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

func (v *formatter_) PreprocessPublishClause(
	publishClause ast.PublishClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessPublishClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessPublishClause(
	publishClause ast.PublishClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessRange(
	range_ ast.RangeLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessRangeSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessRange(
	range_ ast.RangeLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessRejectClause(
	rejectClause ast.RejectClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessRejectClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessRejectClause(
	rejectClause ast.RejectClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessRetrieveClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessReturnClause(
	returnClause ast.ReturnClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessReturnClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessReturnClause(
	returnClause ast.ReturnClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessSaveClause(
	saveClause ast.SaveClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessSaveClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessSaveClause(
	saveClause ast.SaveClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessSelectClause(
	selectClause ast.SelectClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessSelectClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessSelectClause(
	selectClause ast.SelectClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessStatements(
	statements ast.StatementsLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessStatementsSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessStatements(
	statements ast.StatementsLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessThrowClause(
	throwClause ast.ThrowClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessThrowClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessThrowClause(
	throwClause ast.ThrowClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessValue(
	value ast.ValueLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessValueSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessValue(
	value ast.ValueLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessValues(
	values ast.ValuesLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessValuesSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessValues(
	values ast.ValuesLike,
) {
	// TBD - Add formatting of the delimited rule.
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

func (v *formatter_) PreprocessWhileClause(
	whileClause ast.WhileClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessWhileClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessWhileClause(
	whileClause ast.WhileClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PreprocessWithClause(
	withClause ast.WithClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) ProcessWithClauseSlot(
	slot uint,
) {
	// TBD - Add formatting of the delimited rule.
}

func (v *formatter_) PostprocessWithClause(
	withClause ast.WithClauseLike,
) {
	// TBD - Add formatting of the delimited rule.
}

// PROTECTED INTERFACE

// Private Methods

func (v *formatter_) appendNewline() {
	var newline = "\n"
	var indentation = "\t"
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
