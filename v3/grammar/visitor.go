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
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func VisitorClass() VisitorClassLike {
	return visitorClass()
}

// Constructor Methods

func (c *visitorClass_) Visitor(
	processor Methodical,
) VisitorLike {
	if uti.IsUndefined(processor) {
		panic("The \"processor\" attribute is required by this class.")
	}
	var instance = &visitor_{
		// Initialize the instance attributes.
		processor_: processor,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *visitor_) GetClass() VisitorClassLike {
	return visitorClass()
}

func (v *visitor_) VisitDocument(
	document ast.DocumentLike,
) {
	v.processor_.PreprocessDocument(document)
	v.visitDocument(document)
	v.processor_.PostprocessDocument(document)
}

// PROTECTED INTERFACE

// Private Methods

func (v *visitor_) visitAcceptClause(
	acceptClause ast.AcceptClauseLike,
) {
	// Visit a single message rule.
	var message = acceptClause.GetMessage()
	v.processor_.PreprocessMessage(message)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(message)
}

func (v *visitor_) visitAdditionalArgument(
	additionalArgument ast.AdditionalArgumentLike,
) {
	// Visit a single argument rule.
	var argument = additionalArgument.GetArgument()
	v.processor_.PreprocessArgument(argument)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(argument)
}

func (v *visitor_) visitAdditionalAssociation(
	additionalAssociation ast.AdditionalAssociationLike,
) {
	// Visit a single association rule.
	var association = additionalAssociation.GetAssociation()
	v.processor_.PreprocessAssociation(association)
	v.visitAssociation(association)
	v.processor_.PostprocessAssociation(association)
}

func (v *visitor_) visitAdditionalIndex(
	additionalIndex ast.AdditionalIndexLike,
) {
	// Visit a single index rule.
	var index = additionalIndex.GetIndex()
	v.processor_.PreprocessIndex(index)
	v.visitIndex(index)
	v.processor_.PostprocessIndex(index)
}

func (v *visitor_) visitAdditionalStatement(
	additionalStatement ast.AdditionalStatementLike,
) {
	// Visit a single statement rule.
	var statement = additionalStatement.GetStatement()
	v.processor_.PreprocessStatement(statement)
	v.visitStatement(statement)
	v.processor_.PostprocessStatement(statement)
}

func (v *visitor_) visitAdditionalValue(
	additionalValue ast.AdditionalValueLike,
) {
	// Visit a single component rule.
	var component = additionalValue.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)
}

func (v *visitor_) visitAnd(
	and ast.AndLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitAnnotatedAssociation(
	annotatedAssociation ast.AnnotatedAssociationLike,
) {
	// Visit a single association rule.
	var association = annotatedAssociation.GetAssociation()
	v.processor_.PreprocessAssociation(association)
	v.visitAssociation(association)
	v.processor_.PostprocessAssociation(association)

	// Visit slot 1 between references.
	v.processor_.ProcessAnnotatedAssociationSlot(1)

	// Visit an optional note token.
	var optionalNote = annotatedAssociation.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitAnnotatedStatement(
	annotatedStatement ast.AnnotatedStatementLike,
) {
	// Visit the possible annotatedStatement types.
	switch actual := annotatedStatement.GetAny().(type) {
	case ast.EmptyLineLike:
		v.processor_.PreprocessEmptyLine(actual)
		v.visitEmptyLine(actual)
		v.processor_.PostprocessEmptyLine(actual)
	case ast.AnnotationLineLike:
		v.processor_.PreprocessAnnotationLine(actual)
		v.visitAnnotationLine(actual)
		v.processor_.PostprocessAnnotationLine(actual)
	case ast.StatementLineLike:
		v.processor_.PreprocessStatementLine(actual)
		v.visitStatementLine(actual)
		v.processor_.PostprocessStatementLine(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitAnnotatedValue(
	annotatedValue ast.AnnotatedValueLike,
) {
	// Visit a single component rule.
	var component = annotatedValue.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)

	// Visit slot 1 between references.
	v.processor_.ProcessAnnotatedValueSlot(1)

	// Visit an optional note token.
	var optionalNote = annotatedValue.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitAnnotationLine(
	annotationLine ast.AnnotationLineLike,
) {
	// Visit the possible annotationLine types.
	switch actual := annotationLine.GetAny().(type) {
	case string:
		switch {
		case ScannerClass().MatchesType(actual, NoteToken):
			v.processor_.ProcessNote(actual)
		case ScannerClass().MatchesType(actual, CommentToken):
			v.processor_.ProcessComment(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitArgument(
	argument ast.ArgumentLike,
) {
	// Visit a single identifier token.
	var identifier = argument.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
}

func (v *visitor_) visitArguments(
	arguments ast.ArgumentsLike,
) {
	// Visit a single argument rule.
	var argument = arguments.GetArgument()
	v.processor_.PreprocessArgument(argument)
	v.visitArgument(argument)
	v.processor_.PostprocessArgument(argument)

	// Visit slot 1 between references.
	v.processor_.ProcessArgumentsSlot(1)

	// Visit each additionalArgument rule.
	var additionalArgumentIndex uint
	var additionalArguments = arguments.GetAdditionalArguments().GetIterator()
	var additionalArgumentsSize = uint(additionalArguments.GetSize())
	for additionalArguments.HasNext() {
		additionalArgumentIndex++
		var additionalArgument = additionalArguments.GetNext()
		v.processor_.PreprocessAdditionalArgument(
			additionalArgument,
			additionalArgumentIndex,
			additionalArgumentsSize,
		)
		v.visitAdditionalArgument(additionalArgument)
		v.processor_.PostprocessAdditionalArgument(
			additionalArgument,
			additionalArgumentIndex,
			additionalArgumentsSize,
		)
	}
}

func (v *visitor_) visitAssign(
	assign ast.AssignLike,
) {
	// Visit the possible assign types.
	switch actual := assign.GetAny().(type) {
	case string:
		switch {
		case ScannerClass().MatchesType(actual, ColonEqualToken):
			v.processor_.ProcessColonEqual(actual)
		case ScannerClass().MatchesType(actual, DefaultEqualToken):
			v.processor_.ProcessDefaultEqual(actual)
		case ScannerClass().MatchesType(actual, PlusEqualToken):
			v.processor_.ProcessPlusEqual(actual)
		case ScannerClass().MatchesType(actual, DashEqualToken):
			v.processor_.ProcessDashEqual(actual)
		case ScannerClass().MatchesType(actual, StarEqualToken):
			v.processor_.ProcessStarEqual(actual)
		case ScannerClass().MatchesType(actual, SlashEqualToken):
			v.processor_.ProcessSlashEqual(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitAssociation(
	association ast.AssociationLike,
) {
	// Visit a single symbol token.
	var symbol = association.GetSymbol()
	v.processor_.ProcessSymbol(symbol)

	// Visit slot 1 between references.
	v.processor_.ProcessAssociationSlot(1)

	// Visit a single colon token.
	var colon = association.GetColon()
	v.processor_.ProcessColon(colon)

	// Visit slot 2 between references.
	v.processor_.ProcessAssociationSlot(2)

	// Visit a single component rule.
	var component = association.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)
}

func (v *visitor_) visitAtLevel(
	atLevel ast.AtLevelLike,
) {
	// Visit a single expression rule.
	var expression = atLevel.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitBag(
	bag ast.BagLike,
) {
	// Visit a single expression rule.
	var expression = bag.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitBreakClause(
	breakClause ast.BreakClauseLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitCheckoutClause(
	checkoutClause ast.CheckoutClauseLike,
) {
	// Visit a single recipient rule.
	var recipient = checkoutClause.GetRecipient()
	v.processor_.PreprocessRecipient(recipient)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(recipient)

	// Visit slot 1 between references.
	v.processor_.ProcessCheckoutClauseSlot(1)

	// Visit an optional atLevel rule.
	var optionalAtLevel = checkoutClause.GetOptionalAtLevel()
	if uti.IsDefined(optionalAtLevel) {
		v.processor_.PreprocessAtLevel(optionalAtLevel)
		v.visitAtLevel(optionalAtLevel)
		v.processor_.PostprocessAtLevel(optionalAtLevel)
	}

	// Visit slot 2 between references.
	v.processor_.ProcessCheckoutClauseSlot(2)

	// Visit a single citation rule.
	var citation = checkoutClause.GetCitation()
	v.processor_.PreprocessCitation(citation)
	v.visitCitation(citation)
	v.processor_.PostprocessCitation(citation)
}

func (v *visitor_) visitCitation(
	citation ast.CitationLike,
) {
	// Visit a single expression rule.
	var expression = citation.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitCollection(
	collection ast.CollectionLike,
) {
	// Visit the possible collection types.
	switch actual := collection.GetAny().(type) {
	case ast.MultilineAttributesLike:
		v.processor_.PreprocessMultilineAttributes(actual)
		v.visitMultilineAttributes(actual)
		v.processor_.PostprocessMultilineAttributes(actual)
	case ast.MultilineValuesLike:
		v.processor_.PreprocessMultilineValues(actual)
		v.visitMultilineValues(actual)
		v.processor_.PostprocessMultilineValues(actual)
	case ast.NoAttributesLike:
		v.processor_.PreprocessNoAttributes(actual)
		v.visitNoAttributes(actual)
		v.processor_.PostprocessNoAttributes(actual)
	case ast.NoValuesLike:
		v.processor_.PreprocessNoValues(actual)
		v.visitNoValues(actual)
		v.processor_.PostprocessNoValues(actual)
	case ast.InclusiveAnglesLike:
		v.processor_.PreprocessInclusiveAngles(actual)
		v.visitInclusiveAngles(actual)
		v.processor_.PostprocessInclusiveAngles(actual)
	case ast.InExclusiveAnglesLike:
		v.processor_.PreprocessInExclusiveAngles(actual)
		v.visitInExclusiveAngles(actual)
		v.processor_.PostprocessInExclusiveAngles(actual)
	case ast.ExInclusiveAnglesLike:
		v.processor_.PreprocessExInclusiveAngles(actual)
		v.visitExInclusiveAngles(actual)
		v.processor_.PostprocessExInclusiveAngles(actual)
	case ast.ExclusiveAnglesLike:
		v.processor_.PreprocessExclusiveAngles(actual)
		v.visitExclusiveAngles(actual)
		v.processor_.PostprocessExclusiveAngles(actual)
	case ast.InclusiveCitationsLike:
		v.processor_.PreprocessInclusiveCitations(actual)
		v.visitInclusiveCitations(actual)
		v.processor_.PostprocessInclusiveCitations(actual)
	case ast.InExclusiveCitationsLike:
		v.processor_.PreprocessInExclusiveCitations(actual)
		v.visitInExclusiveCitations(actual)
		v.processor_.PostprocessInExclusiveCitations(actual)
	case ast.ExInclusiveCitationsLike:
		v.processor_.PreprocessExInclusiveCitations(actual)
		v.visitExInclusiveCitations(actual)
		v.processor_.PostprocessExInclusiveCitations(actual)
	case ast.ExclusiveCitationsLike:
		v.processor_.PreprocessExclusiveCitations(actual)
		v.visitExclusiveCitations(actual)
		v.processor_.PostprocessExclusiveCitations(actual)
	case ast.InclusiveDurationsLike:
		v.processor_.PreprocessInclusiveDurations(actual)
		v.visitInclusiveDurations(actual)
		v.processor_.PostprocessInclusiveDurations(actual)
	case ast.InExclusiveDurationsLike:
		v.processor_.PreprocessInExclusiveDurations(actual)
		v.visitInExclusiveDurations(actual)
		v.processor_.PostprocessInExclusiveDurations(actual)
	case ast.ExInclusiveDurationsLike:
		v.processor_.PreprocessExInclusiveDurations(actual)
		v.visitExInclusiveDurations(actual)
		v.processor_.PostprocessExInclusiveDurations(actual)
	case ast.ExclusiveDurationsLike:
		v.processor_.PreprocessExclusiveDurations(actual)
		v.visitExclusiveDurations(actual)
		v.processor_.PostprocessExclusiveDurations(actual)
	case ast.InclusiveMomentsLike:
		v.processor_.PreprocessInclusiveMoments(actual)
		v.visitInclusiveMoments(actual)
		v.processor_.PostprocessInclusiveMoments(actual)
	case ast.InExclusiveMomentsLike:
		v.processor_.PreprocessInExclusiveMoments(actual)
		v.visitInExclusiveMoments(actual)
		v.processor_.PostprocessInExclusiveMoments(actual)
	case ast.ExInclusiveMomentsLike:
		v.processor_.PreprocessExInclusiveMoments(actual)
		v.visitExInclusiveMoments(actual)
		v.processor_.PostprocessExInclusiveMoments(actual)
	case ast.ExclusiveMomentsLike:
		v.processor_.PreprocessExclusiveMoments(actual)
		v.visitExclusiveMoments(actual)
		v.processor_.PostprocessExclusiveMoments(actual)
	case ast.InclusiveNamesLike:
		v.processor_.PreprocessInclusiveNames(actual)
		v.visitInclusiveNames(actual)
		v.processor_.PostprocessInclusiveNames(actual)
	case ast.InExclusiveNamesLike:
		v.processor_.PreprocessInExclusiveNames(actual)
		v.visitInExclusiveNames(actual)
		v.processor_.PostprocessInExclusiveNames(actual)
	case ast.ExInclusiveNamesLike:
		v.processor_.PreprocessExInclusiveNames(actual)
		v.visitExInclusiveNames(actual)
		v.processor_.PostprocessExInclusiveNames(actual)
	case ast.ExclusiveNamesLike:
		v.processor_.PreprocessExclusiveNames(actual)
		v.visitExclusiveNames(actual)
		v.processor_.PostprocessExclusiveNames(actual)
	case ast.InclusiveNumbersLike:
		v.processor_.PreprocessInclusiveNumbers(actual)
		v.visitInclusiveNumbers(actual)
		v.processor_.PostprocessInclusiveNumbers(actual)
	case ast.InExclusiveNumbersLike:
		v.processor_.PreprocessInExclusiveNumbers(actual)
		v.visitInExclusiveNumbers(actual)
		v.processor_.PostprocessInExclusiveNumbers(actual)
	case ast.ExInclusiveNumbersLike:
		v.processor_.PreprocessExInclusiveNumbers(actual)
		v.visitExInclusiveNumbers(actual)
		v.processor_.PostprocessExInclusiveNumbers(actual)
	case ast.ExclusiveNumbersLike:
		v.processor_.PreprocessExclusiveNumbers(actual)
		v.visitExclusiveNumbers(actual)
		v.processor_.PostprocessExclusiveNumbers(actual)
	case ast.InclusivePercentagesLike:
		v.processor_.PreprocessInclusivePercentages(actual)
		v.visitInclusivePercentages(actual)
		v.processor_.PostprocessInclusivePercentages(actual)
	case ast.InExclusivePercentagesLike:
		v.processor_.PreprocessInExclusivePercentages(actual)
		v.visitInExclusivePercentages(actual)
		v.processor_.PostprocessInExclusivePercentages(actual)
	case ast.ExInclusivePercentagesLike:
		v.processor_.PreprocessExInclusivePercentages(actual)
		v.visitExInclusivePercentages(actual)
		v.processor_.PostprocessExInclusivePercentages(actual)
	case ast.ExclusivePercentagesLike:
		v.processor_.PreprocessExclusivePercentages(actual)
		v.visitExclusivePercentages(actual)
		v.processor_.PostprocessExclusivePercentages(actual)
	case ast.InclusiveProbabilitiesLike:
		v.processor_.PreprocessInclusiveProbabilities(actual)
		v.visitInclusiveProbabilities(actual)
		v.processor_.PostprocessInclusiveProbabilities(actual)
	case ast.InExclusiveProbabilitiesLike:
		v.processor_.PreprocessInExclusiveProbabilities(actual)
		v.visitInExclusiveProbabilities(actual)
		v.processor_.PostprocessInExclusiveProbabilities(actual)
	case ast.ExInclusiveProbabilitiesLike:
		v.processor_.PreprocessExInclusiveProbabilities(actual)
		v.visitExInclusiveProbabilities(actual)
		v.processor_.PostprocessExInclusiveProbabilities(actual)
	case ast.ExclusiveProbabilitiesLike:
		v.processor_.PreprocessExclusiveProbabilities(actual)
		v.visitExclusiveProbabilities(actual)
		v.processor_.PostprocessExclusiveProbabilities(actual)
	case ast.InclusiveQuotesLike:
		v.processor_.PreprocessInclusiveQuotes(actual)
		v.visitInclusiveQuotes(actual)
		v.processor_.PostprocessInclusiveQuotes(actual)
	case ast.InExclusiveQuotesLike:
		v.processor_.PreprocessInExclusiveQuotes(actual)
		v.visitInExclusiveQuotes(actual)
		v.processor_.PostprocessInExclusiveQuotes(actual)
	case ast.ExInclusiveQuotesLike:
		v.processor_.PreprocessExInclusiveQuotes(actual)
		v.visitExInclusiveQuotes(actual)
		v.processor_.PostprocessExInclusiveQuotes(actual)
	case ast.ExclusiveQuotesLike:
		v.processor_.PreprocessExclusiveQuotes(actual)
		v.visitExclusiveQuotes(actual)
		v.processor_.PostprocessExclusiveQuotes(actual)
	case ast.InclusiveSymbolsLike:
		v.processor_.PreprocessInclusiveSymbols(actual)
		v.visitInclusiveSymbols(actual)
		v.processor_.PostprocessInclusiveSymbols(actual)
	case ast.InExclusiveSymbolsLike:
		v.processor_.PreprocessInExclusiveSymbols(actual)
		v.visitInExclusiveSymbols(actual)
		v.processor_.PostprocessInExclusiveSymbols(actual)
	case ast.ExInclusiveSymbolsLike:
		v.processor_.PreprocessExInclusiveSymbols(actual)
		v.visitExInclusiveSymbols(actual)
		v.processor_.PostprocessExInclusiveSymbols(actual)
	case ast.ExclusiveSymbolsLike:
		v.processor_.PreprocessExclusiveSymbols(actual)
		v.visitExclusiveSymbols(actual)
		v.processor_.PostprocessExclusiveSymbols(actual)
	case ast.InclusiveVersionsLike:
		v.processor_.PreprocessInclusiveVersions(actual)
		v.visitInclusiveVersions(actual)
		v.processor_.PostprocessInclusiveVersions(actual)
	case ast.InExclusiveVersionsLike:
		v.processor_.PreprocessInExclusiveVersions(actual)
		v.visitInExclusiveVersions(actual)
		v.processor_.PostprocessInExclusiveVersions(actual)
	case ast.ExInclusiveVersionsLike:
		v.processor_.PreprocessExInclusiveVersions(actual)
		v.visitExInclusiveVersions(actual)
		v.processor_.PostprocessExInclusiveVersions(actual)
	case ast.ExclusiveVersionsLike:
		v.processor_.PreprocessExclusiveVersions(actual)
		v.visitExclusiveVersions(actual)
		v.processor_.PostprocessExclusiveVersions(actual)
	case ast.InlineAttributesLike:
		v.processor_.PreprocessInlineAttributes(actual)
		v.visitInlineAttributes(actual)
		v.processor_.PostprocessInlineAttributes(actual)
	case ast.InlineValuesLike:
		v.processor_.PreprocessInlineValues(actual)
		v.visitInlineValues(actual)
		v.processor_.PostprocessInlineValues(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitComplement(
	complement ast.ComplementLike,
) {
	// Visit a single logical rule.
	var logical = complement.GetLogical()
	v.processor_.PreprocessLogical(logical)
	v.visitLogical(logical)
	v.processor_.PostprocessLogical(logical)
}

func (v *visitor_) visitComponent(
	component ast.ComponentLike,
) {
	// Visit a single entity rule.
	var entity = component.GetEntity()
	v.processor_.PreprocessEntity(entity)
	v.visitEntity(entity)
	v.processor_.PostprocessEntity(entity)

	// Visit slot 1 between references.
	v.processor_.ProcessComponentSlot(1)

	// Visit an optional parameters rule.
	var optionalParameters = component.GetOptionalParameters()
	if uti.IsDefined(optionalParameters) {
		v.processor_.PreprocessParameters(optionalParameters)
		v.visitParameters(optionalParameters)
		v.processor_.PostprocessParameters(optionalParameters)
	}
}

func (v *visitor_) visitCondition(
	condition ast.ConditionLike,
) {
	// Visit a single expression rule.
	var expression = condition.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitContinueClause(
	continueClause ast.ContinueClauseLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitDiscardClause(
	discardClause ast.DiscardClauseLike,
) {
	// Visit a single draft rule.
	var draft = discardClause.GetDraft()
	v.processor_.PreprocessDraft(draft)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(draft)
}

func (v *visitor_) visitDoClause(
	doClause ast.DoClauseLike,
) {
	// Visit a single invocation rule.
	var invocation = doClause.GetInvocation()
	v.processor_.PreprocessInvocation(invocation)
	v.visitInvocation(invocation)
	v.processor_.PostprocessInvocation(invocation)
}

func (v *visitor_) visitDocument(
	document ast.DocumentLike,
) {
	// Visit an optional comment token.
	var optionalComment = document.GetOptionalComment()
	if uti.IsDefined(optionalComment) {
		v.processor_.ProcessComment(optionalComment)
	}

	// Visit slot 1 between references.
	v.processor_.ProcessDocumentSlot(1)

	// Visit a single component rule.
	var component = document.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)
}

func (v *visitor_) visitDraft(
	draft ast.DraftLike,
) {
	// Visit a single expression rule.
	var expression = draft.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitElement(
	element ast.ElementLike,
) {
	// Visit the possible element types.
	switch actual := element.GetAny().(type) {
	case string:
		switch {
		case ScannerClass().MatchesType(actual, AngleToken):
			v.processor_.ProcessAngle(actual)
		case ScannerClass().MatchesType(actual, BooleanToken):
			v.processor_.ProcessBoolean(actual)
		case ScannerClass().MatchesType(actual, CitationToken):
			v.processor_.ProcessCitation(actual)
		case ScannerClass().MatchesType(actual, DurationToken):
			v.processor_.ProcessDuration(actual)
		case ScannerClass().MatchesType(actual, MomentToken):
			v.processor_.ProcessMoment(actual)
		case ScannerClass().MatchesType(actual, NumberToken):
			v.processor_.ProcessNumber(actual)
		case ScannerClass().MatchesType(actual, PatternToken):
			v.processor_.ProcessPattern(actual)
		case ScannerClass().MatchesType(actual, PercentageToken):
			v.processor_.ProcessPercentage(actual)
		case ScannerClass().MatchesType(actual, ProbabilityToken):
			v.processor_.ProcessProbability(actual)
		case ScannerClass().MatchesType(actual, ResourceToken):
			v.processor_.ProcessResource(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitEmptyLine(
	emptyLine ast.EmptyLineLike,
) {
	// Visit a single newline token.
	var newline = emptyLine.GetNewline()
	v.processor_.ProcessNewline(newline)
}

func (v *visitor_) visitEntity(
	entity ast.EntityLike,
) {
	// Visit the possible entity types.
	switch actual := entity.GetAny().(type) {
	case ast.ElementLike:
		v.processor_.PreprocessElement(actual)
		v.visitElement(actual)
		v.processor_.PostprocessElement(actual)
	case ast.StringLike:
		v.processor_.PreprocessString(actual)
		v.visitString(actual)
		v.processor_.PostprocessString(actual)
	case ast.ProcedureLike:
		v.processor_.PreprocessProcedure(actual)
		v.visitProcedure(actual)
		v.processor_.PostprocessProcedure(actual)
	case ast.CollectionLike:
		v.processor_.PreprocessCollection(actual)
		v.visitCollection(actual)
		v.processor_.PostprocessCollection(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitEvent(
	event ast.EventLike,
) {
	// Visit a single expression rule.
	var expression = event.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitExInclusiveAngles(
	exInclusiveAngles ast.ExInclusiveAnglesLike,
) {
	// Visit a single angle token.
	var angle1 = exInclusiveAngles.GetAngle1()
	v.processor_.ProcessAngle(angle1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveAnglesSlot(1)

	// Visit a single angle token.
	var angle2 = exInclusiveAngles.GetAngle2()
	v.processor_.ProcessAngle(angle2)
}

func (v *visitor_) visitExInclusiveCitations(
	exInclusiveCitations ast.ExInclusiveCitationsLike,
) {
	// Visit a single citation token.
	var citation1 = exInclusiveCitations.GetCitation1()
	v.processor_.ProcessCitation(citation1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveCitationsSlot(1)

	// Visit a single citation token.
	var citation2 = exInclusiveCitations.GetCitation2()
	v.processor_.ProcessCitation(citation2)
}

func (v *visitor_) visitExInclusiveDurations(
	exInclusiveDurations ast.ExInclusiveDurationsLike,
) {
	// Visit a single duration token.
	var duration1 = exInclusiveDurations.GetDuration1()
	v.processor_.ProcessDuration(duration1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveDurationsSlot(1)

	// Visit a single duration token.
	var duration2 = exInclusiveDurations.GetDuration2()
	v.processor_.ProcessDuration(duration2)
}

func (v *visitor_) visitExInclusiveMoments(
	exInclusiveMoments ast.ExInclusiveMomentsLike,
) {
	// Visit a single moment token.
	var moment1 = exInclusiveMoments.GetMoment1()
	v.processor_.ProcessMoment(moment1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveMomentsSlot(1)

	// Visit a single moment token.
	var moment2 = exInclusiveMoments.GetMoment2()
	v.processor_.ProcessMoment(moment2)
}

func (v *visitor_) visitExInclusiveNames(
	exInclusiveNames ast.ExInclusiveNamesLike,
) {
	// Visit a single name token.
	var name1 = exInclusiveNames.GetName1()
	v.processor_.ProcessName(name1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveNamesSlot(1)

	// Visit a single name token.
	var name2 = exInclusiveNames.GetName2()
	v.processor_.ProcessName(name2)
}

func (v *visitor_) visitExInclusiveNumbers(
	exInclusiveNumbers ast.ExInclusiveNumbersLike,
) {
	// Visit a single number token.
	var number1 = exInclusiveNumbers.GetNumber1()
	v.processor_.ProcessNumber(number1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveNumbersSlot(1)

	// Visit a single number token.
	var number2 = exInclusiveNumbers.GetNumber2()
	v.processor_.ProcessNumber(number2)
}

func (v *visitor_) visitExInclusivePercentages(
	exInclusivePercentages ast.ExInclusivePercentagesLike,
) {
	// Visit a single percentage token.
	var percentage1 = exInclusivePercentages.GetPercentage1()
	v.processor_.ProcessPercentage(percentage1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusivePercentagesSlot(1)

	// Visit a single percentage token.
	var percentage2 = exInclusivePercentages.GetPercentage2()
	v.processor_.ProcessPercentage(percentage2)
}

func (v *visitor_) visitExInclusiveProbabilities(
	exInclusiveProbabilities ast.ExInclusiveProbabilitiesLike,
) {
	// Visit a single probability token.
	var probability1 = exInclusiveProbabilities.GetProbability1()
	v.processor_.ProcessProbability(probability1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveProbabilitiesSlot(1)

	// Visit a single probability token.
	var probability2 = exInclusiveProbabilities.GetProbability2()
	v.processor_.ProcessProbability(probability2)
}

func (v *visitor_) visitExInclusiveQuotes(
	exInclusiveQuotes ast.ExInclusiveQuotesLike,
) {
	// Visit a single quote token.
	var quote1 = exInclusiveQuotes.GetQuote1()
	v.processor_.ProcessQuote(quote1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveQuotesSlot(1)

	// Visit a single quote token.
	var quote2 = exInclusiveQuotes.GetQuote2()
	v.processor_.ProcessQuote(quote2)
}

func (v *visitor_) visitExInclusiveSymbols(
	exInclusiveSymbols ast.ExInclusiveSymbolsLike,
) {
	// Visit a single symbol token.
	var symbol1 = exInclusiveSymbols.GetSymbol1()
	v.processor_.ProcessSymbol(symbol1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveSymbolsSlot(1)

	// Visit a single symbol token.
	var symbol2 = exInclusiveSymbols.GetSymbol2()
	v.processor_.ProcessSymbol(symbol2)
}

func (v *visitor_) visitExInclusiveVersions(
	exInclusiveVersions ast.ExInclusiveVersionsLike,
) {
	// Visit a single version token.
	var version1 = exInclusiveVersions.GetVersion1()
	v.processor_.ProcessVersion(version1)

	// Visit slot 1 between references.
	v.processor_.ProcessExInclusiveVersionsSlot(1)

	// Visit a single version token.
	var version2 = exInclusiveVersions.GetVersion2()
	v.processor_.ProcessVersion(version2)
}

func (v *visitor_) visitException(
	exception ast.ExceptionLike,
) {
	// Visit a single expression rule.
	var expression = exception.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitExclusiveAngles(
	exclusiveAngles ast.ExclusiveAnglesLike,
) {
	// Visit a single angle token.
	var angle1 = exclusiveAngles.GetAngle1()
	v.processor_.ProcessAngle(angle1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveAnglesSlot(1)

	// Visit a single angle token.
	var angle2 = exclusiveAngles.GetAngle2()
	v.processor_.ProcessAngle(angle2)
}

func (v *visitor_) visitExclusiveCitations(
	exclusiveCitations ast.ExclusiveCitationsLike,
) {
	// Visit a single citation token.
	var citation1 = exclusiveCitations.GetCitation1()
	v.processor_.ProcessCitation(citation1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveCitationsSlot(1)

	// Visit a single citation token.
	var citation2 = exclusiveCitations.GetCitation2()
	v.processor_.ProcessCitation(citation2)
}

func (v *visitor_) visitExclusiveDurations(
	exclusiveDurations ast.ExclusiveDurationsLike,
) {
	// Visit a single duration token.
	var duration1 = exclusiveDurations.GetDuration1()
	v.processor_.ProcessDuration(duration1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveDurationsSlot(1)

	// Visit a single duration token.
	var duration2 = exclusiveDurations.GetDuration2()
	v.processor_.ProcessDuration(duration2)
}

func (v *visitor_) visitExclusiveMoments(
	exclusiveMoments ast.ExclusiveMomentsLike,
) {
	// Visit a single moment token.
	var moment1 = exclusiveMoments.GetMoment1()
	v.processor_.ProcessMoment(moment1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveMomentsSlot(1)

	// Visit a single moment token.
	var moment2 = exclusiveMoments.GetMoment2()
	v.processor_.ProcessMoment(moment2)
}

func (v *visitor_) visitExclusiveNames(
	exclusiveNames ast.ExclusiveNamesLike,
) {
	// Visit a single name token.
	var name1 = exclusiveNames.GetName1()
	v.processor_.ProcessName(name1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveNamesSlot(1)

	// Visit a single name token.
	var name2 = exclusiveNames.GetName2()
	v.processor_.ProcessName(name2)
}

func (v *visitor_) visitExclusiveNumbers(
	exclusiveNumbers ast.ExclusiveNumbersLike,
) {
	// Visit a single number token.
	var number1 = exclusiveNumbers.GetNumber1()
	v.processor_.ProcessNumber(number1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveNumbersSlot(1)

	// Visit a single number token.
	var number2 = exclusiveNumbers.GetNumber2()
	v.processor_.ProcessNumber(number2)
}

func (v *visitor_) visitExclusivePercentages(
	exclusivePercentages ast.ExclusivePercentagesLike,
) {
	// Visit a single percentage token.
	var percentage1 = exclusivePercentages.GetPercentage1()
	v.processor_.ProcessPercentage(percentage1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusivePercentagesSlot(1)

	// Visit a single percentage token.
	var percentage2 = exclusivePercentages.GetPercentage2()
	v.processor_.ProcessPercentage(percentage2)
}

func (v *visitor_) visitExclusiveProbabilities(
	exclusiveProbabilities ast.ExclusiveProbabilitiesLike,
) {
	// Visit a single probability token.
	var probability1 = exclusiveProbabilities.GetProbability1()
	v.processor_.ProcessProbability(probability1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveProbabilitiesSlot(1)

	// Visit a single probability token.
	var probability2 = exclusiveProbabilities.GetProbability2()
	v.processor_.ProcessProbability(probability2)
}

func (v *visitor_) visitExclusiveQuotes(
	exclusiveQuotes ast.ExclusiveQuotesLike,
) {
	// Visit a single quote token.
	var quote1 = exclusiveQuotes.GetQuote1()
	v.processor_.ProcessQuote(quote1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveQuotesSlot(1)

	// Visit a single quote token.
	var quote2 = exclusiveQuotes.GetQuote2()
	v.processor_.ProcessQuote(quote2)
}

func (v *visitor_) visitExclusiveSymbols(
	exclusiveSymbols ast.ExclusiveSymbolsLike,
) {
	// Visit a single symbol token.
	var symbol1 = exclusiveSymbols.GetSymbol1()
	v.processor_.ProcessSymbol(symbol1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveSymbolsSlot(1)

	// Visit a single symbol token.
	var symbol2 = exclusiveSymbols.GetSymbol2()
	v.processor_.ProcessSymbol(symbol2)
}

func (v *visitor_) visitExclusiveVersions(
	exclusiveVersions ast.ExclusiveVersionsLike,
) {
	// Visit a single version token.
	var version1 = exclusiveVersions.GetVersion1()
	v.processor_.ProcessVersion(version1)

	// Visit slot 1 between references.
	v.processor_.ProcessExclusiveVersionsSlot(1)

	// Visit a single version token.
	var version2 = exclusiveVersions.GetVersion2()
	v.processor_.ProcessVersion(version2)
}

func (v *visitor_) visitExpression(
	expression ast.ExpressionLike,
) {
	// Visit a single subject rule.
	var subject = expression.GetSubject()
	v.processor_.PreprocessSubject(subject)
	v.visitSubject(subject)
	v.processor_.PostprocessSubject(subject)

	// Visit slot 1 between references.
	v.processor_.ProcessExpressionSlot(1)

	// Visit each predicate rule.
	var predicateIndex uint
	var predicates = expression.GetPredicates().GetIterator()
	var predicatesSize = uint(predicates.GetSize())
	for predicates.HasNext() {
		predicateIndex++
		var predicate = predicates.GetNext()
		v.processor_.PreprocessPredicate(
			predicate,
			predicateIndex,
			predicatesSize,
		)
		v.visitPredicate(predicate)
		v.processor_.PostprocessPredicate(
			predicate,
			predicateIndex,
			predicatesSize,
		)
	}
}

func (v *visitor_) visitFailure(
	failure ast.FailureLike,
) {
	// Visit a single symbol token.
	var symbol = failure.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitFlow(
	flow ast.FlowLike,
) {
	// Visit the possible flow types.
	switch actual := flow.GetAny().(type) {
	case ast.IfClauseLike:
		v.processor_.PreprocessIfClause(actual)
		v.visitIfClause(actual)
		v.processor_.PostprocessIfClause(actual)
	case ast.SelectClauseLike:
		v.processor_.PreprocessSelectClause(actual)
		v.visitSelectClause(actual)
		v.processor_.PostprocessSelectClause(actual)
	case ast.WhileClauseLike:
		v.processor_.PreprocessWhileClause(actual)
		v.visitWhileClause(actual)
		v.processor_.PostprocessWhileClause(actual)
	case ast.WithClauseLike:
		v.processor_.PreprocessWithClause(actual)
		v.visitWithClause(actual)
		v.processor_.PostprocessWithClause(actual)
	case ast.ContinueClauseLike:
		v.processor_.PreprocessContinueClause(actual)
		v.visitContinueClause(actual)
		v.processor_.PostprocessContinueClause(actual)
	case ast.BreakClauseLike:
		v.processor_.PreprocessBreakClause(actual)
		v.visitBreakClause(actual)
		v.processor_.PostprocessBreakClause(actual)
	case ast.ReturnClauseLike:
		v.processor_.PreprocessReturnClause(actual)
		v.visitReturnClause(actual)
		v.processor_.PostprocessReturnClause(actual)
	case ast.ThrowClauseLike:
		v.processor_.PreprocessThrowClause(actual)
		v.visitThrowClause(actual)
		v.processor_.PostprocessThrowClause(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitFunction(
	function ast.FunctionLike,
) {
	// Visit a single identifier token.
	var identifier = function.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)

	// Visit slot 1 between references.
	v.processor_.ProcessFunctionSlot(1)

	// Visit an optional arguments rule.
	var optionalArguments = function.GetOptionalArguments()
	if uti.IsDefined(optionalArguments) {
		v.processor_.PreprocessArguments(optionalArguments)
		v.visitArguments(optionalArguments)
		v.processor_.PostprocessArguments(optionalArguments)
	}
}

func (v *visitor_) visitIfClause(
	ifClause ast.IfClauseLike,
) {
	// Visit a single condition rule.
	var condition = ifClause.GetCondition()
	v.processor_.PreprocessCondition(condition)
	v.visitCondition(condition)
	v.processor_.PostprocessCondition(condition)

	// Visit slot 1 between references.
	v.processor_.ProcessIfClauseSlot(1)

	// Visit a single procedure rule.
	var procedure = ifClause.GetProcedure()
	v.processor_.PreprocessProcedure(procedure)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(procedure)
}

func (v *visitor_) visitInExclusiveAngles(
	inExclusiveAngles ast.InExclusiveAnglesLike,
) {
	// Visit a single angle token.
	var angle1 = inExclusiveAngles.GetAngle1()
	v.processor_.ProcessAngle(angle1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveAnglesSlot(1)

	// Visit a single angle token.
	var angle2 = inExclusiveAngles.GetAngle2()
	v.processor_.ProcessAngle(angle2)
}

func (v *visitor_) visitInExclusiveCitations(
	inExclusiveCitations ast.InExclusiveCitationsLike,
) {
	// Visit a single citation token.
	var citation1 = inExclusiveCitations.GetCitation1()
	v.processor_.ProcessCitation(citation1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveCitationsSlot(1)

	// Visit a single citation token.
	var citation2 = inExclusiveCitations.GetCitation2()
	v.processor_.ProcessCitation(citation2)
}

func (v *visitor_) visitInExclusiveDurations(
	inExclusiveDurations ast.InExclusiveDurationsLike,
) {
	// Visit a single duration token.
	var duration1 = inExclusiveDurations.GetDuration1()
	v.processor_.ProcessDuration(duration1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveDurationsSlot(1)

	// Visit a single duration token.
	var duration2 = inExclusiveDurations.GetDuration2()
	v.processor_.ProcessDuration(duration2)
}

func (v *visitor_) visitInExclusiveMoments(
	inExclusiveMoments ast.InExclusiveMomentsLike,
) {
	// Visit a single moment token.
	var moment1 = inExclusiveMoments.GetMoment1()
	v.processor_.ProcessMoment(moment1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveMomentsSlot(1)

	// Visit a single moment token.
	var moment2 = inExclusiveMoments.GetMoment2()
	v.processor_.ProcessMoment(moment2)
}

func (v *visitor_) visitInExclusiveNames(
	inExclusiveNames ast.InExclusiveNamesLike,
) {
	// Visit a single name token.
	var name1 = inExclusiveNames.GetName1()
	v.processor_.ProcessName(name1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveNamesSlot(1)

	// Visit a single name token.
	var name2 = inExclusiveNames.GetName2()
	v.processor_.ProcessName(name2)
}

func (v *visitor_) visitInExclusiveNumbers(
	inExclusiveNumbers ast.InExclusiveNumbersLike,
) {
	// Visit a single number token.
	var number1 = inExclusiveNumbers.GetNumber1()
	v.processor_.ProcessNumber(number1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveNumbersSlot(1)

	// Visit a single number token.
	var number2 = inExclusiveNumbers.GetNumber2()
	v.processor_.ProcessNumber(number2)
}

func (v *visitor_) visitInExclusivePercentages(
	inExclusivePercentages ast.InExclusivePercentagesLike,
) {
	// Visit a single percentage token.
	var percentage1 = inExclusivePercentages.GetPercentage1()
	v.processor_.ProcessPercentage(percentage1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusivePercentagesSlot(1)

	// Visit a single percentage token.
	var percentage2 = inExclusivePercentages.GetPercentage2()
	v.processor_.ProcessPercentage(percentage2)
}

func (v *visitor_) visitInExclusiveProbabilities(
	inExclusiveProbabilities ast.InExclusiveProbabilitiesLike,
) {
	// Visit a single probability token.
	var probability1 = inExclusiveProbabilities.GetProbability1()
	v.processor_.ProcessProbability(probability1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveProbabilitiesSlot(1)

	// Visit a single probability token.
	var probability2 = inExclusiveProbabilities.GetProbability2()
	v.processor_.ProcessProbability(probability2)
}

func (v *visitor_) visitInExclusiveQuotes(
	inExclusiveQuotes ast.InExclusiveQuotesLike,
) {
	// Visit a single quote token.
	var quote1 = inExclusiveQuotes.GetQuote1()
	v.processor_.ProcessQuote(quote1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveQuotesSlot(1)

	// Visit a single quote token.
	var quote2 = inExclusiveQuotes.GetQuote2()
	v.processor_.ProcessQuote(quote2)
}

func (v *visitor_) visitInExclusiveSymbols(
	inExclusiveSymbols ast.InExclusiveSymbolsLike,
) {
	// Visit a single symbol token.
	var symbol1 = inExclusiveSymbols.GetSymbol1()
	v.processor_.ProcessSymbol(symbol1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveSymbolsSlot(1)

	// Visit a single symbol token.
	var symbol2 = inExclusiveSymbols.GetSymbol2()
	v.processor_.ProcessSymbol(symbol2)
}

func (v *visitor_) visitInExclusiveVersions(
	inExclusiveVersions ast.InExclusiveVersionsLike,
) {
	// Visit a single version token.
	var version1 = inExclusiveVersions.GetVersion1()
	v.processor_.ProcessVersion(version1)

	// Visit slot 1 between references.
	v.processor_.ProcessInExclusiveVersionsSlot(1)

	// Visit a single version token.
	var version2 = inExclusiveVersions.GetVersion2()
	v.processor_.ProcessVersion(version2)
}

func (v *visitor_) visitInclusiveAngles(
	inclusiveAngles ast.InclusiveAnglesLike,
) {
	// Visit a single angle token.
	var angle1 = inclusiveAngles.GetAngle1()
	v.processor_.ProcessAngle(angle1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveAnglesSlot(1)

	// Visit a single angle token.
	var angle2 = inclusiveAngles.GetAngle2()
	v.processor_.ProcessAngle(angle2)
}

func (v *visitor_) visitInclusiveCitations(
	inclusiveCitations ast.InclusiveCitationsLike,
) {
	// Visit a single citation token.
	var citation1 = inclusiveCitations.GetCitation1()
	v.processor_.ProcessCitation(citation1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveCitationsSlot(1)

	// Visit a single citation token.
	var citation2 = inclusiveCitations.GetCitation2()
	v.processor_.ProcessCitation(citation2)
}

func (v *visitor_) visitInclusiveDurations(
	inclusiveDurations ast.InclusiveDurationsLike,
) {
	// Visit a single duration token.
	var duration1 = inclusiveDurations.GetDuration1()
	v.processor_.ProcessDuration(duration1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveDurationsSlot(1)

	// Visit a single duration token.
	var duration2 = inclusiveDurations.GetDuration2()
	v.processor_.ProcessDuration(duration2)
}

func (v *visitor_) visitInclusiveMoments(
	inclusiveMoments ast.InclusiveMomentsLike,
) {
	// Visit a single moment token.
	var moment1 = inclusiveMoments.GetMoment1()
	v.processor_.ProcessMoment(moment1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveMomentsSlot(1)

	// Visit a single moment token.
	var moment2 = inclusiveMoments.GetMoment2()
	v.processor_.ProcessMoment(moment2)
}

func (v *visitor_) visitInclusiveNames(
	inclusiveNames ast.InclusiveNamesLike,
) {
	// Visit a single name token.
	var name1 = inclusiveNames.GetName1()
	v.processor_.ProcessName(name1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveNamesSlot(1)

	// Visit a single name token.
	var name2 = inclusiveNames.GetName2()
	v.processor_.ProcessName(name2)
}

func (v *visitor_) visitInclusiveNumbers(
	inclusiveNumbers ast.InclusiveNumbersLike,
) {
	// Visit a single number token.
	var number1 = inclusiveNumbers.GetNumber1()
	v.processor_.ProcessNumber(number1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveNumbersSlot(1)

	// Visit a single number token.
	var number2 = inclusiveNumbers.GetNumber2()
	v.processor_.ProcessNumber(number2)
}

func (v *visitor_) visitInclusivePercentages(
	inclusivePercentages ast.InclusivePercentagesLike,
) {
	// Visit a single percentage token.
	var percentage1 = inclusivePercentages.GetPercentage1()
	v.processor_.ProcessPercentage(percentage1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusivePercentagesSlot(1)

	// Visit a single percentage token.
	var percentage2 = inclusivePercentages.GetPercentage2()
	v.processor_.ProcessPercentage(percentage2)
}

func (v *visitor_) visitInclusiveProbabilities(
	inclusiveProbabilities ast.InclusiveProbabilitiesLike,
) {
	// Visit a single probability token.
	var probability1 = inclusiveProbabilities.GetProbability1()
	v.processor_.ProcessProbability(probability1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveProbabilitiesSlot(1)

	// Visit a single probability token.
	var probability2 = inclusiveProbabilities.GetProbability2()
	v.processor_.ProcessProbability(probability2)
}

func (v *visitor_) visitInclusiveQuotes(
	inclusiveQuotes ast.InclusiveQuotesLike,
) {
	// Visit a single quote token.
	var quote1 = inclusiveQuotes.GetQuote1()
	v.processor_.ProcessQuote(quote1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveQuotesSlot(1)

	// Visit a single quote token.
	var quote2 = inclusiveQuotes.GetQuote2()
	v.processor_.ProcessQuote(quote2)
}

func (v *visitor_) visitInclusiveSymbols(
	inclusiveSymbols ast.InclusiveSymbolsLike,
) {
	// Visit a single symbol token.
	var symbol1 = inclusiveSymbols.GetSymbol1()
	v.processor_.ProcessSymbol(symbol1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveSymbolsSlot(1)

	// Visit a single symbol token.
	var symbol2 = inclusiveSymbols.GetSymbol2()
	v.processor_.ProcessSymbol(symbol2)
}

func (v *visitor_) visitInclusiveVersions(
	inclusiveVersions ast.InclusiveVersionsLike,
) {
	// Visit a single version token.
	var version1 = inclusiveVersions.GetVersion1()
	v.processor_.ProcessVersion(version1)

	// Visit slot 1 between references.
	v.processor_.ProcessInclusiveVersionsSlot(1)

	// Visit a single version token.
	var version2 = inclusiveVersions.GetVersion2()
	v.processor_.ProcessVersion(version2)
}

func (v *visitor_) visitIndex(
	index ast.IndexLike,
) {
	// Visit a single expression rule.
	var expression = index.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitIndices(
	indices ast.IndicesLike,
) {
	// Visit a single index rule.
	var index = indices.GetIndex()
	v.processor_.PreprocessIndex(index)
	v.visitIndex(index)
	v.processor_.PostprocessIndex(index)

	// Visit slot 1 between references.
	v.processor_.ProcessIndicesSlot(1)

	// Visit each additionalIndex rule.
	var additionalIndexIndex uint
	var additionalIndexes = indices.GetAdditionalIndexes().GetIterator()
	var additionalIndexesSize = uint(additionalIndexes.GetSize())
	for additionalIndexes.HasNext() {
		additionalIndexIndex++
		var additionalIndex = additionalIndexes.GetNext()
		v.processor_.PreprocessAdditionalIndex(
			additionalIndex,
			additionalIndexIndex,
			additionalIndexesSize,
		)
		v.visitAdditionalIndex(additionalIndex)
		v.processor_.PostprocessAdditionalIndex(
			additionalIndex,
			additionalIndexIndex,
			additionalIndexesSize,
		)
	}
}

func (v *visitor_) visitIndirect(
	indirect ast.IndirectLike,
) {
	// Visit the possible indirect types.
	switch actual := indirect.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(actual)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(actual)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(actual)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(actual)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(actual)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(actual)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(actual)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(actual)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(actual)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(actual)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(actual)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitInduction(
	induction ast.InductionLike,
) {
	// Visit the possible induction types.
	switch actual := induction.GetAny().(type) {
	case ast.DoClauseLike:
		v.processor_.PreprocessDoClause(actual)
		v.visitDoClause(actual)
		v.processor_.PostprocessDoClause(actual)
	case ast.LetClauseLike:
		v.processor_.PreprocessLetClause(actual)
		v.visitLetClause(actual)
		v.processor_.PostprocessLetClause(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitInlineAttributes(
	inlineAttributes ast.InlineAttributesLike,
) {
	// Visit a single association rule.
	var association = inlineAttributes.GetAssociation()
	v.processor_.PreprocessAssociation(association)
	v.visitAssociation(association)
	v.processor_.PostprocessAssociation(association)

	// Visit slot 1 between references.
	v.processor_.ProcessInlineAttributesSlot(1)

	// Visit each additionalAssociation rule.
	var additionalAssociationIndex uint
	var additionalAssociations = inlineAttributes.GetAdditionalAssociations().GetIterator()
	var additionalAssociationsSize = uint(additionalAssociations.GetSize())
	for additionalAssociations.HasNext() {
		additionalAssociationIndex++
		var additionalAssociation = additionalAssociations.GetNext()
		v.processor_.PreprocessAdditionalAssociation(
			additionalAssociation,
			additionalAssociationIndex,
			additionalAssociationsSize,
		)
		v.visitAdditionalAssociation(additionalAssociation)
		v.processor_.PostprocessAdditionalAssociation(
			additionalAssociation,
			additionalAssociationIndex,
			additionalAssociationsSize,
		)
	}
}

func (v *visitor_) visitInlineParameters(
	inlineParameters ast.InlineParametersLike,
) {
	// Visit a single association rule.
	var association = inlineParameters.GetAssociation()
	v.processor_.PreprocessAssociation(association)
	v.visitAssociation(association)
	v.processor_.PostprocessAssociation(association)

	// Visit slot 1 between references.
	v.processor_.ProcessInlineParametersSlot(1)

	// Visit each additionalAssociation rule.
	var additionalAssociationIndex uint
	var additionalAssociations = inlineParameters.GetAdditionalAssociations().GetIterator()
	var additionalAssociationsSize = uint(additionalAssociations.GetSize())
	for additionalAssociations.HasNext() {
		additionalAssociationIndex++
		var additionalAssociation = additionalAssociations.GetNext()
		v.processor_.PreprocessAdditionalAssociation(
			additionalAssociation,
			additionalAssociationIndex,
			additionalAssociationsSize,
		)
		v.visitAdditionalAssociation(additionalAssociation)
		v.processor_.PostprocessAdditionalAssociation(
			additionalAssociation,
			additionalAssociationIndex,
			additionalAssociationsSize,
		)
	}
}

func (v *visitor_) visitInlineStatements(
	inlineStatements ast.InlineStatementsLike,
) {
	// Visit a single statement rule.
	var statement = inlineStatements.GetStatement()
	v.processor_.PreprocessStatement(statement)
	v.visitStatement(statement)
	v.processor_.PostprocessStatement(statement)

	// Visit slot 1 between references.
	v.processor_.ProcessInlineStatementsSlot(1)

	// Visit each additionalStatement rule.
	var additionalStatementIndex uint
	var additionalStatements = inlineStatements.GetAdditionalStatements().GetIterator()
	var additionalStatementsSize = uint(additionalStatements.GetSize())
	for additionalStatements.HasNext() {
		additionalStatementIndex++
		var additionalStatement = additionalStatements.GetNext()
		v.processor_.PreprocessAdditionalStatement(
			additionalStatement,
			additionalStatementIndex,
			additionalStatementsSize,
		)
		v.visitAdditionalStatement(additionalStatement)
		v.processor_.PostprocessAdditionalStatement(
			additionalStatement,
			additionalStatementIndex,
			additionalStatementsSize,
		)
	}
}

func (v *visitor_) visitInlineValues(
	inlineValues ast.InlineValuesLike,
) {
	// Visit a single component rule.
	var component = inlineValues.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)

	// Visit slot 1 between references.
	v.processor_.ProcessInlineValuesSlot(1)

	// Visit each additionalValue rule.
	var additionalValueIndex uint
	var additionalValues = inlineValues.GetAdditionalValues().GetIterator()
	var additionalValuesSize = uint(additionalValues.GetSize())
	for additionalValues.HasNext() {
		additionalValueIndex++
		var additionalValue = additionalValues.GetNext()
		v.processor_.PreprocessAdditionalValue(
			additionalValue,
			additionalValueIndex,
			additionalValuesSize,
		)
		v.visitAdditionalValue(additionalValue)
		v.processor_.PostprocessAdditionalValue(
			additionalValue,
			additionalValueIndex,
			additionalValuesSize,
		)
	}
}

func (v *visitor_) visitInverse(
	inverse ast.InverseLike,
) {
	// Visit the possible inverse types.
	switch actual := inverse.GetAny().(type) {
	case string:
		switch {
		case ScannerClass().MatchesType(actual, DashToken):
			v.processor_.ProcessDash(actual)
		case ScannerClass().MatchesType(actual, SlashToken):
			v.processor_.ProcessSlash(actual)
		case ScannerClass().MatchesType(actual, StarToken):
			v.processor_.ProcessStar(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitInversion(
	inversion ast.InversionLike,
) {
	// Visit a single inverse rule.
	var inverse = inversion.GetInverse()
	v.processor_.PreprocessInverse(inverse)
	v.visitInverse(inverse)
	v.processor_.PostprocessInverse(inverse)

	// Visit slot 1 between references.
	v.processor_.ProcessInversionSlot(1)

	// Visit a single numerical rule.
	var numerical = inversion.GetNumerical()
	v.processor_.PreprocessNumerical(numerical)
	v.visitNumerical(numerical)
	v.processor_.PostprocessNumerical(numerical)
}

func (v *visitor_) visitInvocation(
	invocation ast.InvocationLike,
) {
	// Visit the possible invocation types.
	switch actual := invocation.GetAny().(type) {
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(actual)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(actual)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(actual)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitIor(
	ior ast.IorLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitIs(
	is ast.IsLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitItem(
	item ast.ItemLike,
) {
	// Visit a single symbol token.
	var symbol = item.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitLetClause(
	letClause ast.LetClauseLike,
) {
	// Visit a single recipient rule.
	var recipient = letClause.GetRecipient()
	v.processor_.PreprocessRecipient(recipient)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(recipient)

	// Visit slot 1 between references.
	v.processor_.ProcessLetClauseSlot(1)

	// Visit a single assign rule.
	var assign = letClause.GetAssign()
	v.processor_.PreprocessAssign(assign)
	v.visitAssign(assign)
	v.processor_.PostprocessAssign(assign)

	// Visit slot 2 between references.
	v.processor_.ProcessLetClauseSlot(2)

	// Visit a single expression rule.
	var expression = letClause.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitLogical(
	logical ast.LogicalLike,
) {
	// Visit the possible logical types.
	switch actual := logical.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(actual)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(actual)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(actual)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(actual)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(actual)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(actual)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(actual)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(actual)
	case ast.ComplementLike:
		v.processor_.PreprocessComplement(actual)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(actual)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(actual)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(actual)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(actual)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(actual)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(actual)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitMagnitude(
	magnitude ast.MagnitudeLike,
) {
	// Visit a single bar token.
	var bar1 = magnitude.GetBar1()
	v.processor_.ProcessBar(bar1)

	// Visit slot 1 between references.
	v.processor_.ProcessMagnitudeSlot(1)

	// Visit a single numerical rule.
	var numerical = magnitude.GetNumerical()
	v.processor_.PreprocessNumerical(numerical)
	v.visitNumerical(numerical)
	v.processor_.PostprocessNumerical(numerical)

	// Visit slot 2 between references.
	v.processor_.ProcessMagnitudeSlot(2)

	// Visit a single bar token.
	var bar2 = magnitude.GetBar2()
	v.processor_.ProcessBar(bar2)
}

func (v *visitor_) visitMainClause(
	mainClause ast.MainClauseLike,
) {
	// Visit the possible mainClause types.
	switch actual := mainClause.GetAny().(type) {
	case ast.FlowLike:
		v.processor_.PreprocessFlow(actual)
		v.visitFlow(actual)
		v.processor_.PostprocessFlow(actual)
	case ast.InductionLike:
		v.processor_.PreprocessInduction(actual)
		v.visitInduction(actual)
		v.processor_.PostprocessInduction(actual)
	case ast.MessagingLike:
		v.processor_.PreprocessMessaging(actual)
		v.visitMessaging(actual)
		v.processor_.PostprocessMessaging(actual)
	case ast.RepositoryLike:
		v.processor_.PreprocessRepository(actual)
		v.visitRepository(actual)
		v.processor_.PostprocessRepository(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitMatchHandler(
	matchHandler ast.MatchHandlerLike,
) {
	// Visit a single template rule.
	var template = matchHandler.GetTemplate()
	v.processor_.PreprocessTemplate(template)
	v.visitTemplate(template)
	v.processor_.PostprocessTemplate(template)

	// Visit slot 1 between references.
	v.processor_.ProcessMatchHandlerSlot(1)

	// Visit a single procedure rule.
	var procedure = matchHandler.GetProcedure()
	v.processor_.PreprocessProcedure(procedure)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(procedure)
}

func (v *visitor_) visitMatches(
	matches ast.MatchesLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitMessage(
	message ast.MessageLike,
) {
	// Visit a single expression rule.
	var expression = message.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitMessaging(
	messaging ast.MessagingLike,
) {
	// Visit the possible messaging types.
	switch actual := messaging.GetAny().(type) {
	case ast.PostClauseLike:
		v.processor_.PreprocessPostClause(actual)
		v.visitPostClause(actual)
		v.processor_.PostprocessPostClause(actual)
	case ast.RetrieveClauseLike:
		v.processor_.PreprocessRetrieveClause(actual)
		v.visitRetrieveClause(actual)
		v.processor_.PostprocessRetrieveClause(actual)
	case ast.AcceptClauseLike:
		v.processor_.PreprocessAcceptClause(actual)
		v.visitAcceptClause(actual)
		v.processor_.PostprocessAcceptClause(actual)
	case ast.RejectClauseLike:
		v.processor_.PreprocessRejectClause(actual)
		v.visitRejectClause(actual)
		v.processor_.PostprocessRejectClause(actual)
	case ast.PublishClauseLike:
		v.processor_.PreprocessPublishClause(actual)
		v.visitPublishClause(actual)
		v.processor_.PostprocessPublishClause(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitMethod(
	method ast.MethodLike,
) {
	// Visit a single identifier token.
	var identifier1 = method.GetIdentifier1()
	v.processor_.ProcessIdentifier(identifier1)

	// Visit slot 1 between references.
	v.processor_.ProcessMethodSlot(1)

	// Visit a single threading rule.
	var threading = method.GetThreading()
	v.processor_.PreprocessThreading(threading)
	v.visitThreading(threading)
	v.processor_.PostprocessThreading(threading)

	// Visit slot 2 between references.
	v.processor_.ProcessMethodSlot(2)

	// Visit a single identifier token.
	var identifier2 = method.GetIdentifier2()
	v.processor_.ProcessIdentifier(identifier2)

	// Visit slot 3 between references.
	v.processor_.ProcessMethodSlot(3)

	// Visit an optional arguments rule.
	var optionalArguments = method.GetOptionalArguments()
	if uti.IsDefined(optionalArguments) {
		v.processor_.PreprocessArguments(optionalArguments)
		v.visitArguments(optionalArguments)
		v.processor_.PostprocessArguments(optionalArguments)
	}
}

func (v *visitor_) visitMultilineAttributes(
	multilineAttributes ast.MultilineAttributesLike,
) {
	// Visit a single newline token.
	var newline = multilineAttributes.GetNewline()
	v.processor_.ProcessNewline(newline)

	// Visit slot 1 between references.
	v.processor_.ProcessMultilineAttributesSlot(1)

	// Visit each annotatedAssociation rule.
	var annotatedAssociationIndex uint
	var annotatedAssociations = multilineAttributes.GetAnnotatedAssociations().GetIterator()
	var annotatedAssociationsSize = uint(annotatedAssociations.GetSize())
	for annotatedAssociations.HasNext() {
		annotatedAssociationIndex++
		var annotatedAssociation = annotatedAssociations.GetNext()
		v.processor_.PreprocessAnnotatedAssociation(
			annotatedAssociation,
			annotatedAssociationIndex,
			annotatedAssociationsSize,
		)
		v.visitAnnotatedAssociation(annotatedAssociation)
		v.processor_.PostprocessAnnotatedAssociation(
			annotatedAssociation,
			annotatedAssociationIndex,
			annotatedAssociationsSize,
		)
	}
}

func (v *visitor_) visitMultilineParameters(
	multilineParameters ast.MultilineParametersLike,
) {
	// Visit a single newline token.
	var newline = multilineParameters.GetNewline()
	v.processor_.ProcessNewline(newline)

	// Visit slot 1 between references.
	v.processor_.ProcessMultilineParametersSlot(1)

	// Visit each annotatedAssociation rule.
	var annotatedAssociationIndex uint
	var annotatedAssociations = multilineParameters.GetAnnotatedAssociations().GetIterator()
	var annotatedAssociationsSize = uint(annotatedAssociations.GetSize())
	for annotatedAssociations.HasNext() {
		annotatedAssociationIndex++
		var annotatedAssociation = annotatedAssociations.GetNext()
		v.processor_.PreprocessAnnotatedAssociation(
			annotatedAssociation,
			annotatedAssociationIndex,
			annotatedAssociationsSize,
		)
		v.visitAnnotatedAssociation(annotatedAssociation)
		v.processor_.PostprocessAnnotatedAssociation(
			annotatedAssociation,
			annotatedAssociationIndex,
			annotatedAssociationsSize,
		)
	}
}

func (v *visitor_) visitMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
) {
	// Visit a single newline token.
	var newline = multilineStatements.GetNewline()
	v.processor_.ProcessNewline(newline)

	// Visit slot 1 between references.
	v.processor_.ProcessMultilineStatementsSlot(1)

	// Visit each annotatedStatement rule.
	var annotatedStatementIndex uint
	var annotatedStatements = multilineStatements.GetAnnotatedStatements().GetIterator()
	var annotatedStatementsSize = uint(annotatedStatements.GetSize())
	for annotatedStatements.HasNext() {
		annotatedStatementIndex++
		var annotatedStatement = annotatedStatements.GetNext()
		v.processor_.PreprocessAnnotatedStatement(
			annotatedStatement,
			annotatedStatementIndex,
			annotatedStatementsSize,
		)
		v.visitAnnotatedStatement(annotatedStatement)
		v.processor_.PostprocessAnnotatedStatement(
			annotatedStatement,
			annotatedStatementIndex,
			annotatedStatementsSize,
		)
	}
}

func (v *visitor_) visitMultilineValues(
	multilineValues ast.MultilineValuesLike,
) {
	// Visit a single newline token.
	var newline = multilineValues.GetNewline()
	v.processor_.ProcessNewline(newline)

	// Visit slot 1 between references.
	v.processor_.ProcessMultilineValuesSlot(1)

	// Visit each annotatedValue rule.
	var annotatedValueIndex uint
	var annotatedValues = multilineValues.GetAnnotatedValues().GetIterator()
	var annotatedValuesSize = uint(annotatedValues.GetSize())
	for annotatedValues.HasNext() {
		annotatedValueIndex++
		var annotatedValue = annotatedValues.GetNext()
		v.processor_.PreprocessAnnotatedValue(
			annotatedValue,
			annotatedValueIndex,
			annotatedValuesSize,
		)
		v.visitAnnotatedValue(annotatedValue)
		v.processor_.PostprocessAnnotatedValue(
			annotatedValue,
			annotatedValueIndex,
			annotatedValuesSize,
		)
	}
}

func (v *visitor_) visitNoAttributes(
	noAttributes ast.NoAttributesLike,
) {
	// Visit a single colon token.
	var colon = noAttributes.GetColon()
	v.processor_.ProcessColon(colon)
}

func (v *visitor_) visitNoStatements(
	noStatements ast.NoStatementsLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitNoValues(
	noValues ast.NoValuesLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitNotarizeClause(
	notarizeClause ast.NotarizeClauseLike,
) {
	// Visit a single draft rule.
	var draft = notarizeClause.GetDraft()
	v.processor_.PreprocessDraft(draft)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(draft)

	// Visit slot 1 between references.
	v.processor_.ProcessNotarizeClauseSlot(1)

	// Visit a single citation rule.
	var citation = notarizeClause.GetCitation()
	v.processor_.PreprocessCitation(citation)
	v.visitCitation(citation)
	v.processor_.PostprocessCitation(citation)
}

func (v *visitor_) visitNumerical(
	numerical ast.NumericalLike,
) {
	// Visit the possible numerical types.
	switch actual := numerical.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(actual)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(actual)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(actual)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(actual)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(actual)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(actual)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(actual)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(actual)
	case ast.InversionLike:
		v.processor_.PreprocessInversion(actual)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(actual)
	case ast.MagnitudeLike:
		v.processor_.PreprocessMagnitude(actual)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(actual)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(actual)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(actual)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(actual)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(actual)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(actual)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitOnClause(
	onClause ast.OnClauseLike,
) {
	// Visit a single failure rule.
	var failure = onClause.GetFailure()
	v.processor_.PreprocessFailure(failure)
	v.visitFailure(failure)
	v.processor_.PostprocessFailure(failure)

	// Visit slot 1 between references.
	v.processor_.ProcessOnClauseSlot(1)

	// Visit each matchHandler rule.
	var matchHandlerIndex uint
	var matchHandlers = onClause.GetMatchHandlers().GetIterator()
	var matchHandlersSize = uint(matchHandlers.GetSize())
	for matchHandlers.HasNext() {
		matchHandlerIndex++
		var matchHandler = matchHandlers.GetNext()
		v.processor_.PreprocessMatchHandler(
			matchHandler,
			matchHandlerIndex,
			matchHandlersSize,
		)
		v.visitMatchHandler(matchHandler)
		v.processor_.PostprocessMatchHandler(
			matchHandler,
			matchHandlerIndex,
			matchHandlersSize,
		)
	}
}

func (v *visitor_) visitOperator(
	operator ast.OperatorLike,
) {
	// Visit the possible operator types.
	switch actual := operator.GetAny().(type) {
	case ast.IsLike:
		v.processor_.PreprocessIs(actual)
		v.visitIs(actual)
		v.processor_.PostprocessIs(actual)
	case ast.MatchesLike:
		v.processor_.PreprocessMatches(actual)
		v.visitMatches(actual)
		v.processor_.PostprocessMatches(actual)
	case ast.AndLike:
		v.processor_.PreprocessAnd(actual)
		v.visitAnd(actual)
		v.processor_.PostprocessAnd(actual)
	case ast.SanLike:
		v.processor_.PreprocessSan(actual)
		v.visitSan(actual)
		v.processor_.PostprocessSan(actual)
	case ast.IorLike:
		v.processor_.PreprocessIor(actual)
		v.visitIor(actual)
		v.processor_.PostprocessIor(actual)
	case ast.XorLike:
		v.processor_.PreprocessXor(actual)
		v.visitXor(actual)
		v.processor_.PostprocessXor(actual)
	case string:
		switch {
		case ScannerClass().MatchesType(actual, PlusToken):
			v.processor_.ProcessPlus(actual)
		case ScannerClass().MatchesType(actual, DashToken):
			v.processor_.ProcessDash(actual)
		case ScannerClass().MatchesType(actual, StarToken):
			v.processor_.ProcessStar(actual)
		case ScannerClass().MatchesType(actual, SlashToken):
			v.processor_.ProcessSlash(actual)
		case ScannerClass().MatchesType(actual, DoubleToken):
			v.processor_.ProcessDouble(actual)
		case ScannerClass().MatchesType(actual, CaretToken):
			v.processor_.ProcessCaret(actual)
		case ScannerClass().MatchesType(actual, LessToken):
			v.processor_.ProcessLess(actual)
		case ScannerClass().MatchesType(actual, EqualToken):
			v.processor_.ProcessEqual(actual)
		case ScannerClass().MatchesType(actual, MoreToken):
			v.processor_.ProcessMore(actual)
		case ScannerClass().MatchesType(actual, AmpersandToken):
			v.processor_.ProcessAmpersand(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitParameters(
	parameters ast.ParametersLike,
) {
	// Visit the possible parameters types.
	switch actual := parameters.GetAny().(type) {
	case ast.MultilineParametersLike:
		v.processor_.PreprocessMultilineParameters(actual)
		v.visitMultilineParameters(actual)
		v.processor_.PostprocessMultilineParameters(actual)
	case ast.InlineParametersLike:
		v.processor_.PreprocessInlineParameters(actual)
		v.visitInlineParameters(actual)
		v.processor_.PostprocessInlineParameters(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitPostClause(
	postClause ast.PostClauseLike,
) {
	// Visit a single message rule.
	var message = postClause.GetMessage()
	v.processor_.PreprocessMessage(message)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(message)

	// Visit slot 1 between references.
	v.processor_.ProcessPostClauseSlot(1)

	// Visit a single bag rule.
	var bag = postClause.GetBag()
	v.processor_.PreprocessBag(bag)
	v.visitBag(bag)
	v.processor_.PostprocessBag(bag)
}

func (v *visitor_) visitPrecedence(
	precedence ast.PrecedenceLike,
) {
	// Visit a single expression rule.
	var expression = precedence.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitPredicate(
	predicate ast.PredicateLike,
) {
	// Visit a single operator rule.
	var operator = predicate.GetOperator()
	v.processor_.PreprocessOperator(operator)
	v.visitOperator(operator)
	v.processor_.PostprocessOperator(operator)

	// Visit slot 1 between references.
	v.processor_.ProcessPredicateSlot(1)

	// Visit a single expression rule.
	var expression = predicate.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitProcedure(
	procedure ast.ProcedureLike,
) {
	// Visit the possible procedure types.
	switch actual := procedure.GetAny().(type) {
	case ast.MultilineStatementsLike:
		v.processor_.PreprocessMultilineStatements(actual)
		v.visitMultilineStatements(actual)
		v.processor_.PostprocessMultilineStatements(actual)
	case ast.InlineStatementsLike:
		v.processor_.PreprocessInlineStatements(actual)
		v.visitInlineStatements(actual)
		v.processor_.PostprocessInlineStatements(actual)
	case ast.NoStatementsLike:
		v.processor_.PreprocessNoStatements(actual)
		v.visitNoStatements(actual)
		v.processor_.PostprocessNoStatements(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitPublishClause(
	publishClause ast.PublishClauseLike,
) {
	// Visit a single event rule.
	var event = publishClause.GetEvent()
	v.processor_.PreprocessEvent(event)
	v.visitEvent(event)
	v.processor_.PostprocessEvent(event)
}

func (v *visitor_) visitRecipient(
	recipient ast.RecipientLike,
) {
	// Visit the possible recipient types.
	switch actual := recipient.GetAny().(type) {
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(actual)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(actual)
	case string:
		switch {
		case ScannerClass().MatchesType(actual, SymbolToken):
			v.processor_.ProcessSymbol(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitReferent(
	referent ast.ReferentLike,
) {
	// Visit a single snail token.
	var snail = referent.GetSnail()
	v.processor_.ProcessSnail(snail)

	// Visit slot 1 between references.
	v.processor_.ProcessReferentSlot(1)

	// Visit a single indirect rule.
	var indirect = referent.GetIndirect()
	v.processor_.PreprocessIndirect(indirect)
	v.visitIndirect(indirect)
	v.processor_.PostprocessIndirect(indirect)
}

func (v *visitor_) visitRejectClause(
	rejectClause ast.RejectClauseLike,
) {
	// Visit a single message rule.
	var message = rejectClause.GetMessage()
	v.processor_.PreprocessMessage(message)
	v.visitMessage(message)
	v.processor_.PostprocessMessage(message)
}

func (v *visitor_) visitRepository(
	repository ast.RepositoryLike,
) {
	// Visit the possible repository types.
	switch actual := repository.GetAny().(type) {
	case ast.CheckoutClauseLike:
		v.processor_.PreprocessCheckoutClause(actual)
		v.visitCheckoutClause(actual)
		v.processor_.PostprocessCheckoutClause(actual)
	case ast.SaveClauseLike:
		v.processor_.PreprocessSaveClause(actual)
		v.visitSaveClause(actual)
		v.processor_.PostprocessSaveClause(actual)
	case ast.DiscardClauseLike:
		v.processor_.PreprocessDiscardClause(actual)
		v.visitDiscardClause(actual)
		v.processor_.PostprocessDiscardClause(actual)
	case ast.NotarizeClauseLike:
		v.processor_.PreprocessNotarizeClause(actual)
		v.visitNotarizeClause(actual)
		v.processor_.PostprocessNotarizeClause(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitResult(
	result ast.ResultLike,
) {
	// Visit a single expression rule.
	var expression = result.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitRetrieveClause(
	retrieveClause ast.RetrieveClauseLike,
) {
	// Visit a single recipient rule.
	var recipient = retrieveClause.GetRecipient()
	v.processor_.PreprocessRecipient(recipient)
	v.visitRecipient(recipient)
	v.processor_.PostprocessRecipient(recipient)

	// Visit slot 1 between references.
	v.processor_.ProcessRetrieveClauseSlot(1)

	// Visit a single bag rule.
	var bag = retrieveClause.GetBag()
	v.processor_.PreprocessBag(bag)
	v.visitBag(bag)
	v.processor_.PostprocessBag(bag)
}

func (v *visitor_) visitReturnClause(
	returnClause ast.ReturnClauseLike,
) {
	// Visit a single result rule.
	var result = returnClause.GetResult()
	v.processor_.PreprocessResult(result)
	v.visitResult(result)
	v.processor_.PostprocessResult(result)
}

func (v *visitor_) visitSan(
	san ast.SanLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitSaveClause(
	saveClause ast.SaveClauseLike,
) {
	// Visit a single draft rule.
	var draft = saveClause.GetDraft()
	v.processor_.PreprocessDraft(draft)
	v.visitDraft(draft)
	v.processor_.PostprocessDraft(draft)

	// Visit slot 1 between references.
	v.processor_.ProcessSaveClauseSlot(1)

	// Visit a single citation rule.
	var citation = saveClause.GetCitation()
	v.processor_.PreprocessCitation(citation)
	v.visitCitation(citation)
	v.processor_.PostprocessCitation(citation)
}

func (v *visitor_) visitSelectClause(
	selectClause ast.SelectClauseLike,
) {
	// Visit a single target rule.
	var target = selectClause.GetTarget()
	v.processor_.PreprocessTarget(target)
	v.visitTarget(target)
	v.processor_.PostprocessTarget(target)

	// Visit slot 1 between references.
	v.processor_.ProcessSelectClauseSlot(1)

	// Visit each matchHandler rule.
	var matchHandlerIndex uint
	var matchHandlers = selectClause.GetMatchHandlers().GetIterator()
	var matchHandlersSize = uint(matchHandlers.GetSize())
	for matchHandlers.HasNext() {
		matchHandlerIndex++
		var matchHandler = matchHandlers.GetNext()
		v.processor_.PreprocessMatchHandler(
			matchHandler,
			matchHandlerIndex,
			matchHandlersSize,
		)
		v.visitMatchHandler(matchHandler)
		v.processor_.PostprocessMatchHandler(
			matchHandler,
			matchHandlerIndex,
			matchHandlersSize,
		)
	}
}

func (v *visitor_) visitSequence(
	sequence ast.SequenceLike,
) {
	// Visit a single expression rule.
	var expression = sequence.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitStatement(
	statement ast.StatementLike,
) {
	// Visit a single mainClause rule.
	var mainClause = statement.GetMainClause()
	v.processor_.PreprocessMainClause(mainClause)
	v.visitMainClause(mainClause)
	v.processor_.PostprocessMainClause(mainClause)

	// Visit slot 1 between references.
	v.processor_.ProcessStatementSlot(1)

	// Visit an optional onClause rule.
	var optionalOnClause = statement.GetOptionalOnClause()
	if uti.IsDefined(optionalOnClause) {
		v.processor_.PreprocessOnClause(optionalOnClause)
		v.visitOnClause(optionalOnClause)
		v.processor_.PostprocessOnClause(optionalOnClause)
	}
}

func (v *visitor_) visitStatementLine(
	statementLine ast.StatementLineLike,
) {
	// Visit a single statement rule.
	var statement = statementLine.GetStatement()
	v.processor_.PreprocessStatement(statement)
	v.visitStatement(statement)
	v.processor_.PostprocessStatement(statement)

	// Visit slot 1 between references.
	v.processor_.ProcessStatementLineSlot(1)

	// Visit an optional note token.
	var optionalNote = statementLine.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitString(
	string_ ast.StringLike,
) {
	// Visit the possible string types.
	switch actual := string_.GetAny().(type) {
	case string:
		switch {
		case ScannerClass().MatchesType(actual, BinaryToken):
			v.processor_.ProcessBinary(actual)
		case ScannerClass().MatchesType(actual, BytecodeToken):
			v.processor_.ProcessBytecode(actual)
		case ScannerClass().MatchesType(actual, NameToken):
			v.processor_.ProcessName(actual)
		case ScannerClass().MatchesType(actual, NarrativeToken):
			v.processor_.ProcessNarrative(actual)
		case ScannerClass().MatchesType(actual, QuoteToken):
			v.processor_.ProcessQuote(actual)
		case ScannerClass().MatchesType(actual, SymbolToken):
			v.processor_.ProcessSymbol(actual)
		case ScannerClass().MatchesType(actual, TagToken):
			v.processor_.ProcessTag(actual)
		case ScannerClass().MatchesType(actual, VersionToken):
			v.processor_.ProcessVersion(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitSubcomponent(
	subcomponent ast.SubcomponentLike,
) {
	// Visit a single identifier token.
	var identifier = subcomponent.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)

	// Visit slot 1 between references.
	v.processor_.ProcessSubcomponentSlot(1)

	// Visit a single indices rule.
	var indices = subcomponent.GetIndices()
	v.processor_.PreprocessIndices(indices)
	v.visitIndices(indices)
	v.processor_.PostprocessIndices(indices)
}

func (v *visitor_) visitSubject(
	subject ast.SubjectLike,
) {
	// Visit the possible subject types.
	switch actual := subject.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(actual)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(actual)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(actual)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(actual)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(actual)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(actual)
	case ast.ReferentLike:
		v.processor_.PreprocessReferent(actual)
		v.visitReferent(actual)
		v.processor_.PostprocessReferent(actual)
	case ast.ComplementLike:
		v.processor_.PreprocessComplement(actual)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(actual)
	case ast.InversionLike:
		v.processor_.PreprocessInversion(actual)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(actual)
	case ast.MagnitudeLike:
		v.processor_.PreprocessMagnitude(actual)
		v.visitMagnitude(actual)
		v.processor_.PostprocessMagnitude(actual)
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(actual)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(actual)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(actual)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(actual)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(actual)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitTarget(
	target ast.TargetLike,
) {
	// Visit the possible target types.
	switch actual := target.GetAny().(type) {
	case ast.FunctionLike:
		v.processor_.PreprocessFunction(actual)
		v.visitFunction(actual)
		v.processor_.PostprocessFunction(actual)
	case ast.MethodLike:
		v.processor_.PreprocessMethod(actual)
		v.visitMethod(actual)
		v.processor_.PostprocessMethod(actual)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(actual)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(actual)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(actual)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitTemplate(
	template ast.TemplateLike,
) {
	// Visit a single expression rule.
	var expression = template.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitThreading(
	threading ast.ThreadingLike,
) {
	// Visit the possible threading types.
	switch actual := threading.GetAny().(type) {
	case string:
		switch {
		case ScannerClass().MatchesType(actual, DotToken):
			v.processor_.ProcessDot(actual)
		case ScannerClass().MatchesType(actual, ArrowToken):
			v.processor_.ProcessArrow(actual)
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitThrowClause(
	throwClause ast.ThrowClauseLike,
) {
	// Visit a single exception rule.
	var exception = throwClause.GetException()
	v.processor_.PreprocessException(exception)
	v.visitException(exception)
	v.processor_.PostprocessException(exception)
}

func (v *visitor_) visitVariable(
	variable ast.VariableLike,
) {
	// Visit a single identifier token.
	var identifier = variable.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
}

func (v *visitor_) visitWhileClause(
	whileClause ast.WhileClauseLike,
) {
	// Visit a single condition rule.
	var condition = whileClause.GetCondition()
	v.processor_.PreprocessCondition(condition)
	v.visitCondition(condition)
	v.processor_.PostprocessCondition(condition)

	// Visit slot 1 between references.
	v.processor_.ProcessWhileClauseSlot(1)

	// Visit a single procedure rule.
	var procedure = whileClause.GetProcedure()
	v.processor_.PreprocessProcedure(procedure)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(procedure)
}

func (v *visitor_) visitWithClause(
	withClause ast.WithClauseLike,
) {
	// Visit a single item rule.
	var item = withClause.GetItem()
	v.processor_.PreprocessItem(item)
	v.visitItem(item)
	v.processor_.PostprocessItem(item)

	// Visit slot 1 between references.
	v.processor_.ProcessWithClauseSlot(1)

	// Visit a single sequence rule.
	var sequence = withClause.GetSequence()
	v.processor_.PreprocessSequence(sequence)
	v.visitSequence(sequence)
	v.processor_.PostprocessSequence(sequence)

	// Visit slot 2 between references.
	v.processor_.ProcessWithClauseSlot(2)

	// Visit a single procedure rule.
	var procedure = withClause.GetProcedure()
	v.processor_.PreprocessProcedure(procedure)
	v.visitProcedure(procedure)
	v.processor_.PostprocessProcedure(procedure)
}

func (v *visitor_) visitXor(
	xor ast.XorLike,
) {
	// This method does not need to process anything.
}

// Instance Structure

type visitor_ struct {
	// Declare the instance attributes.
	processor_ Methodical
}

// Class Structure

type visitorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func visitorClass() *visitorClass_ {
	return visitorClassReference_
}

var visitorClassReference_ = &visitorClass_{
	// Initialize the class constants.
}
