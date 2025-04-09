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

func (v *visitor_) visitAdditionalParameter(
	additionalParameter ast.AdditionalParameterLike,
) {
	// Visit a single parameter rule.
	var parameter = additionalParameter.GetParameter()
	v.processor_.PreprocessParameter(parameter)
	v.visitParameter(parameter)
	v.processor_.PostprocessParameter(parameter)
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
	// Visit a single value rule.
	var value = additionalValue.GetValue()
	v.processor_.PreprocessValue(value)
	v.visitValue(value)
	v.processor_.PostprocessValue(value)
}

func (v *visitor_) visitAmplitude(
	amplitude ast.AmplitudeLike,
) {
	// Visit a single expression rule.
	var expression = amplitude.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
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

func (v *visitor_) visitAnnotatedParameter(
	annotatedParameter ast.AnnotatedParameterLike,
) {
	// Visit a single parameter rule.
	var parameter = annotatedParameter.GetParameter()
	v.processor_.PreprocessParameter(parameter)
	v.visitParameter(parameter)
	v.processor_.PostprocessParameter(parameter)

	// Visit slot 1 between references.
	v.processor_.ProcessAnnotatedParameterSlot(1)

	// Visit an optional note token.
	var optionalNote = annotatedParameter.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitAnnotatedStatement(
	annotatedStatement ast.AnnotatedStatementLike,
) {
	// Visit a single statement rule.
	var statement = annotatedStatement.GetStatement()
	v.processor_.PreprocessStatement(statement)
	v.visitStatement(statement)
	v.processor_.PostprocessStatement(statement)

	// Visit slot 1 between references.
	v.processor_.ProcessAnnotatedStatementSlot(1)

	// Visit an optional note token.
	var optionalNote = annotatedStatement.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitAnnotatedValue(
	annotatedValue ast.AnnotatedValueLike,
) {
	// Visit a single value rule.
	var value = annotatedValue.GetValue()
	v.processor_.PreprocessValue(value)
	v.visitValue(value)
	v.processor_.PostprocessValue(value)

	// Visit slot 1 between references.
	v.processor_.ProcessAnnotatedValueSlot(1)

	// Visit an optional note token.
	var optionalNote = annotatedValue.GetOptionalNote()
	if uti.IsDefined(optionalNote) {
		v.processor_.ProcessNote(optionalNote)
	}
}

func (v *visitor_) visitArgument(
	argument ast.ArgumentLike,
) {
	// Visit a single expression rule.
	var expression = argument.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
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

func (v *visitor_) visitArithmetic(
	arithmetic ast.ArithmeticLike,
) {
	// Visit a single expression rule.
	var expression1 = arithmetic.GetExpression1()
	v.processor_.PreprocessExpression(expression1)
	v.visitExpression(expression1)
	v.processor_.PostprocessExpression(expression1)

	// Visit slot 1 between references.
	v.processor_.ProcessArithmeticSlot(1)

	// Visit a single arithmeticOperator rule.
	var arithmeticOperator = arithmetic.GetArithmeticOperator()
	v.processor_.PreprocessArithmeticOperator(arithmeticOperator)
	v.visitArithmeticOperator(arithmeticOperator)
	v.processor_.PostprocessArithmeticOperator(arithmeticOperator)

	// Visit slot 2 between references.
	v.processor_.ProcessArithmeticSlot(2)

	// Visit a single expression rule.
	var expression2 = arithmetic.GetExpression2()
	v.processor_.PreprocessExpression(expression2)
	v.visitExpression(expression2)
	v.processor_.PostprocessExpression(expression2)
}

func (v *visitor_) visitArithmeticOperator(
	arithmeticOperator ast.ArithmeticOperatorLike,
) {
	// Visit the possible arithmeticOperator types.
	switch actual := arithmeticOperator.GetAny().(type) {
	case ast.TimesLike:
		v.processor_.PreprocessTimes(actual)
		v.visitTimes(actual)
		v.processor_.PostprocessTimes(actual)
	case ast.DivideLike:
		v.processor_.PreprocessDivide(actual)
		v.visitDivide(actual)
		v.processor_.PostprocessDivide(actual)
	case ast.ModuloLike:
		v.processor_.PreprocessModulo(actual)
		v.visitModulo(actual)
		v.processor_.PostprocessModulo(actual)
	case ast.PlusLike:
		v.processor_.PreprocessPlus(actual)
		v.visitPlus(actual)
		v.processor_.PostprocessPlus(actual)
	case ast.MinusLike:
		v.processor_.PreprocessMinus(actual)
		v.visitMinus(actual)
		v.processor_.PostprocessMinus(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitArrow(
	arrow ast.ArrowLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitAssignmentOperator(
	assignmentOperator ast.AssignmentOperatorLike,
) {
	// Visit the possible assignmentOperator types.
	switch actual := assignmentOperator.GetAny().(type) {
	case ast.ColonEqualsLike:
		v.processor_.PreprocessColonEquals(actual)
		v.visitColonEquals(actual)
		v.processor_.PostprocessColonEquals(actual)
	case ast.DefaultEqualsLike:
		v.processor_.PreprocessDefaultEquals(actual)
		v.visitDefaultEquals(actual)
		v.processor_.PostprocessDefaultEquals(actual)
	case ast.PlusEqualsLike:
		v.processor_.PreprocessPlusEquals(actual)
		v.visitPlusEquals(actual)
		v.processor_.PostprocessPlusEquals(actual)
	case ast.MinusEqualsLike:
		v.processor_.PreprocessMinusEquals(actual)
		v.visitMinusEquals(actual)
		v.processor_.PostprocessMinusEquals(actual)
	case ast.TimesEqualsLike:
		v.processor_.PreprocessTimesEquals(actual)
		v.visitTimesEquals(actual)
		v.processor_.PostprocessTimesEquals(actual)
	case ast.DivideEqualsLike:
		v.processor_.PreprocessDivideEquals(actual)
		v.visitDivideEquals(actual)
		v.processor_.PostprocessDivideEquals(actual)
	case string:
		switch {
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
	// Visit a single key rule.
	var key = association.GetKey()
	v.processor_.PreprocessKey(key)
	v.visitKey(key)
	v.processor_.PostprocessKey(key)

	// Visit slot 1 between references.
	v.processor_.ProcessAssociationSlot(1)

	// Visit a single value rule.
	var value = association.GetValue()
	v.processor_.PreprocessValue(value)
	v.visitValue(value)
	v.processor_.PostprocessValue(value)
}

func (v *visitor_) visitAssociations(
	associations ast.AssociationsLike,
) {
	// Visit the possible associations types.
	switch actual := associations.GetAny().(type) {
	case ast.InlineAssociationsLike:
		v.processor_.PreprocessInlineAssociations(actual)
		v.visitInlineAssociations(actual)
		v.processor_.PostprocessInlineAssociations(actual)
	case ast.MultilineAssociationsLike:
		v.processor_.PreprocessMultilineAssociations(actual)
		v.visitMultilineAssociations(actual)
		v.processor_.PostprocessMultilineAssociations(actual)
	case ast.NoAssociationsLike:
		v.processor_.PreprocessNoAssociations(actual)
		v.visitNoAssociations(actual)
		v.processor_.PostprocessNoAssociations(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
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

func (v *visitor_) visitAttribute(
	attribute ast.AttributeLike,
) {
	// Visit a single variable rule.
	var variable = attribute.GetVariable()
	v.processor_.PreprocessVariable(variable)
	v.visitVariable(variable)
	v.processor_.PostprocessVariable(variable)

	// Visit slot 1 between references.
	v.processor_.ProcessAttributeSlot(1)

	// Visit a single indices rule.
	var indices = attribute.GetIndices()
	v.processor_.PreprocessIndices(indices)
	v.visitIndices(indices)
	v.processor_.PostprocessIndices(indices)
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

func (v *visitor_) visitChaining(
	chaining ast.ChainingLike,
) {
	// Visit a single expression rule.
	var expression1 = chaining.GetExpression1()
	v.processor_.PreprocessExpression(expression1)
	v.visitExpression(expression1)
	v.processor_.PostprocessExpression(expression1)

	// Visit slot 1 between references.
	v.processor_.ProcessChainingSlot(1)

	// Visit a single expression rule.
	var expression2 = chaining.GetExpression2()
	v.processor_.PreprocessExpression(expression2)
	v.visitExpression(expression2)
	v.processor_.PostprocessExpression(expression2)
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
	// Visit an optional items rule.
	var optionalItems = collection.GetOptionalItems()
	if uti.IsDefined(optionalItems) {
		v.processor_.PreprocessItems(optionalItems)
		v.visitItems(optionalItems)
		v.processor_.PostprocessItems(optionalItems)
	}
}

func (v *visitor_) visitColonEquals(
	colonEquals ast.ColonEqualsLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitComparison(
	comparison ast.ComparisonLike,
) {
	// Visit a single expression rule.
	var expression1 = comparison.GetExpression1()
	v.processor_.PreprocessExpression(expression1)
	v.visitExpression(expression1)
	v.processor_.PostprocessExpression(expression1)

	// Visit slot 1 between references.
	v.processor_.ProcessComparisonSlot(1)

	// Visit a single comparisonOperator rule.
	var comparisonOperator = comparison.GetComparisonOperator()
	v.processor_.PreprocessComparisonOperator(comparisonOperator)
	v.visitComparisonOperator(comparisonOperator)
	v.processor_.PostprocessComparisonOperator(comparisonOperator)

	// Visit slot 2 between references.
	v.processor_.ProcessComparisonSlot(2)

	// Visit a single expression rule.
	var expression2 = comparison.GetExpression2()
	v.processor_.PreprocessExpression(expression2)
	v.visitExpression(expression2)
	v.processor_.PostprocessExpression(expression2)
}

func (v *visitor_) visitComparisonOperator(
	comparisonOperator ast.ComparisonOperatorLike,
) {
	// Visit the possible comparisonOperator types.
	switch actual := comparisonOperator.GetAny().(type) {
	case ast.LessLike:
		v.processor_.PreprocessLess(actual)
		v.visitLess(actual)
		v.processor_.PostprocessLess(actual)
	case ast.EqualsLike:
		v.processor_.PreprocessEquals(actual)
		v.visitEquals(actual)
		v.processor_.PostprocessEquals(actual)
	case ast.MoreLike:
		v.processor_.PreprocessMore(actual)
		v.visitMore(actual)
		v.processor_.PostprocessMore(actual)
	case ast.IsLike:
		v.processor_.PreprocessIs(actual)
		v.visitIs(actual)
		v.processor_.PostprocessIs(actual)
	case ast.MatchesLike:
		v.processor_.PreprocessMatches(actual)
		v.visitMatches(actual)
		v.processor_.PostprocessMatches(actual)
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
	// Visit a single expression rule.
	var expression = complement.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
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

	// Visit an optional context rule.
	var optionalContext = component.GetOptionalContext()
	if uti.IsDefined(optionalContext) {
		v.processor_.PreprocessContext(optionalContext)
		v.visitContext(optionalContext)
		v.processor_.PostprocessContext(optionalContext)
	}
}

func (v *visitor_) visitComposite(
	composite ast.CompositeLike,
) {
	// Visit a single expression rule.
	var expression = composite.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
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

func (v *visitor_) visitContext(
	context ast.ContextLike,
) {
	// Visit a single parameters rule.
	var parameters = context.GetParameters()
	v.processor_.PreprocessParameters(parameters)
	v.visitParameters(parameters)
	v.processor_.PostprocessParameters(parameters)
}

func (v *visitor_) visitContinueClause(
	continueClause ast.ContinueClauseLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitDefaultEquals(
	defaultEquals ast.DefaultEqualsLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitDereference(
	dereference ast.DereferenceLike,
) {
	// Visit a single expression rule.
	var expression = dereference.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
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

func (v *visitor_) visitDivide(
	divide ast.DivideLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitDivideEquals(
	divideEquals ast.DivideEqualsLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitDoClause(
	doClause ast.DoClauseLike,
) {
	// Visit a single operation rule.
	var operation = doClause.GetOperation()
	v.processor_.PreprocessOperation(operation)
	v.visitOperation(operation)
	v.processor_.PostprocessOperation(operation)
}

func (v *visitor_) visitDocument(
	document ast.DocumentLike,
) {
	// Visit a single comment token.
	var comment = document.GetComment()
	v.processor_.ProcessComment(comment)

	// Visit slot 1 between references.
	v.processor_.ProcessDocumentSlot(1)

	// Visit a single component rule.
	var component = document.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)
}

func (v *visitor_) visitDot(
	dot ast.DotLike,
) {
	// This method does not need to process anything.
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
	case ast.RangeLike:
		v.processor_.PreprocessRange(actual)
		v.visitRange(actual)
		v.processor_.PostprocessRange(actual)
	case ast.CollectionLike:
		v.processor_.PreprocessCollection(actual)
		v.visitCollection(actual)
		v.processor_.PostprocessCollection(actual)
	case ast.ProcedureLike:
		v.processor_.PreprocessProcedure(actual)
		v.visitProcedure(actual)
		v.processor_.PostprocessProcedure(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitEquals(
	equals ast.EqualsLike,
) {
	// This method does not need to process anything.
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

func (v *visitor_) visitException(
	exception ast.ExceptionLike,
) {
	// Visit a single expression rule.
	var expression = exception.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitExponential(
	exponential ast.ExponentialLike,
) {
	// Visit a single expression rule.
	var expression1 = exponential.GetExpression1()
	v.processor_.PreprocessExpression(expression1)
	v.visitExpression(expression1)
	v.processor_.PostprocessExpression(expression1)

	// Visit slot 1 between references.
	v.processor_.ProcessExponentialSlot(1)

	// Visit a single expression rule.
	var expression2 = exponential.GetExpression2()
	v.processor_.PreprocessExpression(expression2)
	v.visitExpression(expression2)
	v.processor_.PostprocessExpression(expression2)
}

func (v *visitor_) visitExpression(
	expression ast.ExpressionLike,
) {
	// Visit the possible expression types.
	switch actual := expression.GetAny().(type) {
	case ast.ComponentLike:
		v.processor_.PreprocessComponent(actual)
		v.visitComponent(actual)
		v.processor_.PostprocessComponent(actual)
	case ast.IntrinsicLike:
		v.processor_.PreprocessIntrinsic(actual)
		v.visitIntrinsic(actual)
		v.processor_.PostprocessIntrinsic(actual)
	case ast.VariableLike:
		v.processor_.PreprocessVariable(actual)
		v.visitVariable(actual)
		v.processor_.PostprocessVariable(actual)
	case ast.PrecedenceLike:
		v.processor_.PreprocessPrecedence(actual)
		v.visitPrecedence(actual)
		v.processor_.PostprocessPrecedence(actual)
	case ast.DereferenceLike:
		v.processor_.PreprocessDereference(actual)
		v.visitDereference(actual)
		v.processor_.PostprocessDereference(actual)
	case ast.InvocationLike:
		v.processor_.PreprocessInvocation(actual)
		v.visitInvocation(actual)
		v.processor_.PostprocessInvocation(actual)
	case ast.SubcomponentLike:
		v.processor_.PreprocessSubcomponent(actual)
		v.visitSubcomponent(actual)
		v.processor_.PostprocessSubcomponent(actual)
	case ast.ChainingLike:
		v.processor_.PreprocessChaining(actual)
		v.visitChaining(actual)
		v.processor_.PostprocessChaining(actual)
	case ast.ExponentialLike:
		v.processor_.PreprocessExponential(actual)
		v.visitExponential(actual)
		v.processor_.PostprocessExponential(actual)
	case ast.InversionLike:
		v.processor_.PreprocessInversion(actual)
		v.visitInversion(actual)
		v.processor_.PostprocessInversion(actual)
	case ast.ArithmeticLike:
		v.processor_.PreprocessArithmetic(actual)
		v.visitArithmetic(actual)
		v.processor_.PostprocessArithmetic(actual)
	case ast.AmplitudeLike:
		v.processor_.PreprocessAmplitude(actual)
		v.visitAmplitude(actual)
		v.processor_.PostprocessAmplitude(actual)
	case ast.ComparisonLike:
		v.processor_.PreprocessComparison(actual)
		v.visitComparison(actual)
		v.processor_.PostprocessComparison(actual)
	case ast.ComplementLike:
		v.processor_.PreprocessComplement(actual)
		v.visitComplement(actual)
		v.processor_.PostprocessComplement(actual)
	case ast.LogicalLike:
		v.processor_.PreprocessLogical(actual)
		v.visitLogical(actual)
		v.processor_.PostprocessLogical(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
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
}

func (v *visitor_) visitIOR(
	iOR ast.IORLike,
) {
	// This method does not need to process anything.
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

func (v *visitor_) visitInlineAssociations(
	inlineAssociations ast.InlineAssociationsLike,
) {
	// Visit a single association rule.
	var association = inlineAssociations.GetAssociation()
	v.processor_.PreprocessAssociation(association)
	v.visitAssociation(association)
	v.processor_.PostprocessAssociation(association)

	// Visit slot 1 between references.
	v.processor_.ProcessInlineAssociationsSlot(1)

	// Visit each additionalAssociation rule.
	var additionalAssociationIndex uint
	var additionalAssociations = inlineAssociations.GetAdditionalAssociations().GetIterator()
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
	// Visit a single parameter rule.
	var parameter = inlineParameters.GetParameter()
	v.processor_.PreprocessParameter(parameter)
	v.visitParameter(parameter)
	v.processor_.PostprocessParameter(parameter)

	// Visit slot 1 between references.
	v.processor_.ProcessInlineParametersSlot(1)

	// Visit each additionalParameter rule.
	var additionalParameterIndex uint
	var additionalParameters = inlineParameters.GetAdditionalParameters().GetIterator()
	var additionalParametersSize = uint(additionalParameters.GetSize())
	for additionalParameters.HasNext() {
		additionalParameterIndex++
		var additionalParameter = additionalParameters.GetNext()
		v.processor_.PreprocessAdditionalParameter(
			additionalParameter,
			additionalParameterIndex,
			additionalParametersSize,
		)
		v.visitAdditionalParameter(additionalParameter)
		v.processor_.PostprocessAdditionalParameter(
			additionalParameter,
			additionalParameterIndex,
			additionalParametersSize,
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
	// Visit a single value rule.
	var value = inlineValues.GetValue()
	v.processor_.PreprocessValue(value)
	v.visitValue(value)
	v.processor_.PostprocessValue(value)

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

func (v *visitor_) visitIntrinsic(
	intrinsic ast.IntrinsicLike,
) {
	// Visit a single function rule.
	var function = intrinsic.GetFunction()
	v.processor_.PreprocessFunction(function)
	v.visitFunction(function)
	v.processor_.PostprocessFunction(function)

	// Visit slot 1 between references.
	v.processor_.ProcessIntrinsicSlot(1)

	// Visit an optional arguments rule.
	var optionalArguments = intrinsic.GetOptionalArguments()
	if uti.IsDefined(optionalArguments) {
		v.processor_.PreprocessArguments(optionalArguments)
		v.visitArguments(optionalArguments)
		v.processor_.PostprocessArguments(optionalArguments)
	}
}

func (v *visitor_) visitInversion(
	inversion ast.InversionLike,
) {
	// Visit a single inversionOperator rule.
	var inversionOperator = inversion.GetInversionOperator()
	v.processor_.PreprocessInversionOperator(inversionOperator)
	v.visitInversionOperator(inversionOperator)
	v.processor_.PostprocessInversionOperator(inversionOperator)

	// Visit slot 1 between references.
	v.processor_.ProcessInversionSlot(1)

	// Visit a single expression rule.
	var expression = inversion.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
}

func (v *visitor_) visitInversionOperator(
	inversionOperator ast.InversionOperatorLike,
) {
	// Visit the possible inversionOperator types.
	switch actual := inversionOperator.GetAny().(type) {
	case ast.MinusLike:
		v.processor_.PreprocessMinus(actual)
		v.visitMinus(actual)
		v.processor_.PostprocessMinus(actual)
	case ast.DivideLike:
		v.processor_.PreprocessDivide(actual)
		v.visitDivide(actual)
		v.processor_.PostprocessDivide(actual)
	case ast.TimesLike:
		v.processor_.PreprocessTimes(actual)
		v.visitTimes(actual)
		v.processor_.PostprocessTimes(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitInvocation(
	invocation ast.InvocationLike,
) {
	// Visit a single target rule.
	var target = invocation.GetTarget()
	v.processor_.PreprocessTarget(target)
	v.visitTarget(target)
	v.processor_.PostprocessTarget(target)

	// Visit slot 1 between references.
	v.processor_.ProcessInvocationSlot(1)

	// Visit a single invocationOperator rule.
	var invocationOperator = invocation.GetInvocationOperator()
	v.processor_.PreprocessInvocationOperator(invocationOperator)
	v.visitInvocationOperator(invocationOperator)
	v.processor_.PostprocessInvocationOperator(invocationOperator)

	// Visit slot 2 between references.
	v.processor_.ProcessInvocationSlot(2)

	// Visit a single method rule.
	var method = invocation.GetMethod()
	v.processor_.PreprocessMethod(method)
	v.visitMethod(method)
	v.processor_.PostprocessMethod(method)

	// Visit slot 3 between references.
	v.processor_.ProcessInvocationSlot(3)

	// Visit an optional arguments rule.
	var optionalArguments = invocation.GetOptionalArguments()
	if uti.IsDefined(optionalArguments) {
		v.processor_.PreprocessArguments(optionalArguments)
		v.visitArguments(optionalArguments)
		v.processor_.PostprocessArguments(optionalArguments)
	}
}

func (v *visitor_) visitInvocationOperator(
	invocationOperator ast.InvocationOperatorLike,
) {
	// Visit the possible invocationOperator types.
	switch actual := invocationOperator.GetAny().(type) {
	case ast.DotLike:
		v.processor_.PreprocessDot(actual)
		v.visitDot(actual)
		v.processor_.PostprocessDot(actual)
	case ast.ArrowLike:
		v.processor_.PreprocessArrow(actual)
		v.visitArrow(actual)
		v.processor_.PostprocessArrow(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
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

func (v *visitor_) visitItems(
	items ast.ItemsLike,
) {
	// Visit the possible items types.
	switch actual := items.GetAny().(type) {
	case ast.AssociationsLike:
		v.processor_.PreprocessAssociations(actual)
		v.visitAssociations(actual)
		v.processor_.PostprocessAssociations(actual)
	case ast.ValuesLike:
		v.processor_.PreprocessValues(actual)
		v.visitValues(actual)
		v.processor_.PostprocessValues(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitKey(
	key ast.KeyLike,
) {
	// Visit a single primitive rule.
	var primitive = key.GetPrimitive()
	v.processor_.PreprocessPrimitive(primitive)
	v.visitPrimitive(primitive)
	v.processor_.PostprocessPrimitive(primitive)
}

func (v *visitor_) visitLabel(
	label ast.LabelLike,
) {
	// Visit a single symbol token.
	var symbol = label.GetSymbol()
	v.processor_.ProcessSymbol(symbol)
}

func (v *visitor_) visitLeftBracket(
	leftBracket ast.LeftBracketLike,
) {
	// Visit the possible leftBracket types.
	switch actual := leftBracket.GetAny().(type) {
	case ast.LeftSquareLike:
		v.processor_.PreprocessLeftSquare(actual)
		v.visitLeftSquare(actual)
		v.processor_.PostprocessLeftSquare(actual)
	case ast.LeftRoundLike:
		v.processor_.PreprocessLeftRound(actual)
		v.visitLeftRound(actual)
		v.processor_.PostprocessLeftRound(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitLeftRound(
	leftRound ast.LeftRoundLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitLeftSquare(
	leftSquare ast.LeftSquareLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitLess(
	less ast.LessLike,
) {
	// This method does not need to process anything.
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

	// Visit a single assignmentOperator rule.
	var assignmentOperator = letClause.GetAssignmentOperator()
	v.processor_.PreprocessAssignmentOperator(assignmentOperator)
	v.visitAssignmentOperator(assignmentOperator)
	v.processor_.PostprocessAssignmentOperator(assignmentOperator)

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
	// Visit a single expression rule.
	var expression1 = logical.GetExpression1()
	v.processor_.PreprocessExpression(expression1)
	v.visitExpression(expression1)
	v.processor_.PostprocessExpression(expression1)

	// Visit slot 1 between references.
	v.processor_.ProcessLogicalSlot(1)

	// Visit a single logicalOperator rule.
	var logicalOperator = logical.GetLogicalOperator()
	v.processor_.PreprocessLogicalOperator(logicalOperator)
	v.visitLogicalOperator(logicalOperator)
	v.processor_.PostprocessLogicalOperator(logicalOperator)

	// Visit slot 2 between references.
	v.processor_.ProcessLogicalSlot(2)

	// Visit a single expression rule.
	var expression2 = logical.GetExpression2()
	v.processor_.PreprocessExpression(expression2)
	v.visitExpression(expression2)
	v.processor_.PostprocessExpression(expression2)
}

func (v *visitor_) visitLogicalOperator(
	logicalOperator ast.LogicalOperatorLike,
) {
	// Visit the possible logicalOperator types.
	switch actual := logicalOperator.GetAny().(type) {
	case ast.AndLike:
		v.processor_.PreprocessAnd(actual)
		v.visitAnd(actual)
		v.processor_.PostprocessAnd(actual)
	case ast.SANLike:
		v.processor_.PreprocessSAN(actual)
		v.visitSAN(actual)
		v.processor_.PostprocessSAN(actual)
	case ast.IORLike:
		v.processor_.PreprocessIOR(actual)
		v.visitIOR(actual)
		v.processor_.PostprocessIOR(actual)
	case ast.XORLike:
		v.processor_.PreprocessXOR(actual)
		v.visitXOR(actual)
		v.processor_.PostprocessXOR(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
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
	var identifier = method.GetIdentifier()
	v.processor_.ProcessIdentifier(identifier)
}

func (v *visitor_) visitMinus(
	minus ast.MinusLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitMinusEquals(
	minusEquals ast.MinusEqualsLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitModulo(
	modulo ast.ModuloLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitMore(
	more ast.MoreLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitMultilineAssociations(
	multilineAssociations ast.MultilineAssociationsLike,
) {
	// Visit each annotatedAssociation rule.
	var annotatedAssociationIndex uint
	var annotatedAssociations = multilineAssociations.GetAnnotatedAssociations().GetIterator()
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
	// Visit each annotatedParameter rule.
	var annotatedParameterIndex uint
	var annotatedParameters = multilineParameters.GetAnnotatedParameters().GetIterator()
	var annotatedParametersSize = uint(annotatedParameters.GetSize())
	for annotatedParameters.HasNext() {
		annotatedParameterIndex++
		var annotatedParameter = annotatedParameters.GetNext()
		v.processor_.PreprocessAnnotatedParameter(
			annotatedParameter,
			annotatedParameterIndex,
			annotatedParametersSize,
		)
		v.visitAnnotatedParameter(annotatedParameter)
		v.processor_.PostprocessAnnotatedParameter(
			annotatedParameter,
			annotatedParameterIndex,
			annotatedParametersSize,
		)
	}
}

func (v *visitor_) visitMultilineStatements(
	multilineStatements ast.MultilineStatementsLike,
) {
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

func (v *visitor_) visitNoAssociations(
	noAssociations ast.NoAssociationsLike,
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

func (v *visitor_) visitOperation(
	operation ast.OperationLike,
) {
	// Visit the possible operation types.
	switch actual := operation.GetAny().(type) {
	case ast.IntrinsicLike:
		v.processor_.PreprocessIntrinsic(actual)
		v.visitIntrinsic(actual)
		v.processor_.PostprocessIntrinsic(actual)
	case ast.InvocationLike:
		v.processor_.PreprocessInvocation(actual)
		v.visitInvocation(actual)
		v.processor_.PostprocessInvocation(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitParameter(
	parameter ast.ParameterLike,
) {
	// Visit a single label rule.
	var label = parameter.GetLabel()
	v.processor_.PreprocessLabel(label)
	v.visitLabel(label)
	v.processor_.PostprocessLabel(label)

	// Visit slot 1 between references.
	v.processor_.ProcessParameterSlot(1)

	// Visit a single component rule.
	var component = parameter.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)
}

func (v *visitor_) visitParameters(
	parameters ast.ParametersLike,
) {
	// Visit the possible parameters types.
	switch actual := parameters.GetAny().(type) {
	case ast.InlineParametersLike:
		v.processor_.PreprocessInlineParameters(actual)
		v.visitInlineParameters(actual)
		v.processor_.PostprocessInlineParameters(actual)
	case ast.MultilineParametersLike:
		v.processor_.PreprocessMultilineParameters(actual)
		v.visitMultilineParameters(actual)
		v.processor_.PostprocessMultilineParameters(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitPlus(
	plus ast.PlusLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitPlusEquals(
	plusEquals ast.PlusEqualsLike,
) {
	// This method does not need to process anything.
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

func (v *visitor_) visitPrimitive(
	primitive ast.PrimitiveLike,
) {
	// Visit the possible primitive types.
	switch actual := primitive.GetAny().(type) {
	case ast.ElementLike:
		v.processor_.PreprocessElement(actual)
		v.visitElement(actual)
		v.processor_.PostprocessElement(actual)
	case ast.StringLike:
		v.processor_.PreprocessString(actual)
		v.visitString(actual)
		v.processor_.PostprocessString(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitProcedure(
	procedure ast.ProcedureLike,
) {
	// Visit an optional statements rule.
	var optionalStatements = procedure.GetOptionalStatements()
	if uti.IsDefined(optionalStatements) {
		v.processor_.PreprocessStatements(optionalStatements)
		v.visitStatements(optionalStatements)
		v.processor_.PostprocessStatements(optionalStatements)
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

func (v *visitor_) visitRange(
	range_ ast.RangeLike,
) {
	// Visit a single leftBracket rule.
	var leftBracket = range_.GetLeftBracket()
	v.processor_.PreprocessLeftBracket(leftBracket)
	v.visitLeftBracket(leftBracket)
	v.processor_.PostprocessLeftBracket(leftBracket)

	// Visit slot 1 between references.
	v.processor_.ProcessRangeSlot(1)

	// Visit a single primitive rule.
	var primitive1 = range_.GetPrimitive1()
	v.processor_.PreprocessPrimitive(primitive1)
	v.visitPrimitive(primitive1)
	v.processor_.PostprocessPrimitive(primitive1)

	// Visit slot 2 between references.
	v.processor_.ProcessRangeSlot(2)

	// Visit a single primitive rule.
	var primitive2 = range_.GetPrimitive2()
	v.processor_.PreprocessPrimitive(primitive2)
	v.visitPrimitive(primitive2)
	v.processor_.PostprocessPrimitive(primitive2)

	// Visit slot 3 between references.
	v.processor_.ProcessRangeSlot(3)

	// Visit a single rightBracket rule.
	var rightBracket = range_.GetRightBracket()
	v.processor_.PreprocessRightBracket(rightBracket)
	v.visitRightBracket(rightBracket)
	v.processor_.PostprocessRightBracket(rightBracket)
}

func (v *visitor_) visitRecipient(
	recipient ast.RecipientLike,
) {
	// Visit the possible recipient types.
	switch actual := recipient.GetAny().(type) {
	case ast.LabelLike:
		v.processor_.PreprocessLabel(actual)
		v.visitLabel(actual)
		v.processor_.PostprocessLabel(actual)
	case ast.AttributeLike:
		v.processor_.PreprocessAttribute(actual)
		v.visitAttribute(actual)
		v.processor_.PostprocessAttribute(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
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

func (v *visitor_) visitRightBracket(
	rightBracket ast.RightBracketLike,
) {
	// Visit the possible rightBracket types.
	switch actual := rightBracket.GetAny().(type) {
	case ast.RightSquareLike:
		v.processor_.PreprocessRightSquare(actual)
		v.visitRightSquare(actual)
		v.processor_.PostprocessRightSquare(actual)
	case ast.RightRoundLike:
		v.processor_.PreprocessRightRound(actual)
		v.visitRightRound(actual)
		v.processor_.PostprocessRightRound(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
}

func (v *visitor_) visitRightRound(
	rightRound ast.RightRoundLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitRightSquare(
	rightSquare ast.RightSquareLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitSAN(
	sAN ast.SANLike,
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

func (v *visitor_) visitStatements(
	statements ast.StatementsLike,
) {
	// Visit the possible statements types.
	switch actual := statements.GetAny().(type) {
	case ast.InlineStatementsLike:
		v.processor_.PreprocessInlineStatements(actual)
		v.visitInlineStatements(actual)
		v.processor_.PostprocessInlineStatements(actual)
	case ast.MultilineStatementsLike:
		v.processor_.PreprocessMultilineStatements(actual)
		v.visitMultilineStatements(actual)
		v.processor_.PostprocessMultilineStatements(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
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
	// Visit a single composite rule.
	var composite = subcomponent.GetComposite()
	v.processor_.PreprocessComposite(composite)
	v.visitComposite(composite)
	v.processor_.PostprocessComposite(composite)

	// Visit slot 1 between references.
	v.processor_.ProcessSubcomponentSlot(1)

	// Visit a single indices rule.
	var indices = subcomponent.GetIndices()
	v.processor_.PreprocessIndices(indices)
	v.visitIndices(indices)
	v.processor_.PostprocessIndices(indices)
}

func (v *visitor_) visitTarget(
	target ast.TargetLike,
) {
	// Visit a single expression rule.
	var expression = target.GetExpression()
	v.processor_.PreprocessExpression(expression)
	v.visitExpression(expression)
	v.processor_.PostprocessExpression(expression)
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

func (v *visitor_) visitThrowClause(
	throwClause ast.ThrowClauseLike,
) {
	// Visit a single exception rule.
	var exception = throwClause.GetException()
	v.processor_.PreprocessException(exception)
	v.visitException(exception)
	v.processor_.PostprocessException(exception)
}

func (v *visitor_) visitTimes(
	times ast.TimesLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitTimesEquals(
	timesEquals ast.TimesEqualsLike,
) {
	// This method does not need to process anything.
}

func (v *visitor_) visitValue(
	value ast.ValueLike,
) {
	// Visit a single component rule.
	var component = value.GetComponent()
	v.processor_.PreprocessComponent(component)
	v.visitComponent(component)
	v.processor_.PostprocessComponent(component)
}

func (v *visitor_) visitValues(
	values ast.ValuesLike,
) {
	// Visit the possible values types.
	switch actual := values.GetAny().(type) {
	case ast.InlineValuesLike:
		v.processor_.PreprocessInlineValues(actual)
		v.visitInlineValues(actual)
		v.processor_.PostprocessInlineValues(actual)
	case ast.MultilineValuesLike:
		v.processor_.PreprocessMultilineValues(actual)
		v.visitMultilineValues(actual)
		v.processor_.PostprocessMultilineValues(actual)
	case string:
		switch {
		default:
			panic(fmt.Sprintf("Invalid token: %v", actual))
		}
	default:
		panic(fmt.Sprintf("Invalid rule type: %T", actual))
	}
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

func (v *visitor_) visitXOR(
	xOR ast.XORLike,
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
