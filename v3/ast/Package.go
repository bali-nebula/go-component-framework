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
│              This "Package.go" file was automatically generated.             │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘

Package "ast" provides the abstract syntax tree (AST) classes for this module
based on the "Syntax.cdsn" grammar for the module.  Each AST class manages the
attributes associated with its corresponding rule definition found in the
grammar.

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-component-framework/wiki

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-class-model/wiki

Additional concrete implementations of the classes declared by this package can
be developed and used seamlessly since the interface declarations only depend on
other interfaces and intrinsic types—and the class implementations only depend
on interfaces, not on each other.
*/
package ast

import (
	col "github.com/craterdog/go-collection-framework/v5/collection"
)

// TYPE DECLARATIONS

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
AcceptClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete accept-clause-like class.
*/
type AcceptClauseClassLike interface {
	// Constructor Methods
	AcceptClause(
		message MessageLike,
	) AcceptClauseLike
}

/*
AdditionalArgumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-argument-like class.
*/
type AdditionalArgumentClassLike interface {
	// Constructor Methods
	AdditionalArgument(
		argument ArgumentLike,
	) AdditionalArgumentLike
}

/*
AdditionalAssociationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-association-like class.
*/
type AdditionalAssociationClassLike interface {
	// Constructor Methods
	AdditionalAssociation(
		association AssociationLike,
	) AdditionalAssociationLike
}

/*
AdditionalIndexClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-index-like class.
*/
type AdditionalIndexClassLike interface {
	// Constructor Methods
	AdditionalIndex(
		index IndexLike,
	) AdditionalIndexLike
}

/*
AdditionalParameterClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-parameter-like class.
*/
type AdditionalParameterClassLike interface {
	// Constructor Methods
	AdditionalParameter(
		parameter ParameterLike,
	) AdditionalParameterLike
}

/*
AdditionalStatementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-statement-like class.
*/
type AdditionalStatementClassLike interface {
	// Constructor Methods
	AdditionalStatement(
		statement StatementLike,
	) AdditionalStatementLike
}

/*
AdditionalValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete additional-value-like class.
*/
type AdditionalValueClassLike interface {
	// Constructor Methods
	AdditionalValue(
		value ValueLike,
	) AdditionalValueLike
}

/*
AnnotatedAssociationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete annotated-association-like class.
*/
type AnnotatedAssociationClassLike interface {
	// Constructor Methods
	AnnotatedAssociation(
		association AssociationLike,
		optionalNote string,
	) AnnotatedAssociationLike
}

/*
AnnotatedParameterClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete annotated-parameter-like class.
*/
type AnnotatedParameterClassLike interface {
	// Constructor Methods
	AnnotatedParameter(
		parameter ParameterLike,
		optionalNote string,
	) AnnotatedParameterLike
}

/*
AnnotatedStatementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete annotated-statement-like class.
*/
type AnnotatedStatementClassLike interface {
	// Constructor Methods
	AnnotatedStatement(
		statement StatementLike,
		optionalNote string,
	) AnnotatedStatementLike
}

/*
AnnotatedValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete annotated-value-like class.
*/
type AnnotatedValueClassLike interface {
	// Constructor Methods
	AnnotatedValue(
		value ValueLike,
		optionalNote string,
	) AnnotatedValueLike
}

/*
ArgumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructor Methods
	Argument(
		expression ExpressionLike,
	) ArgumentLike
}

/*
ArgumentsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructor Methods
	Arguments(
		argument ArgumentLike,
		additionalArguments col.Sequential[AdditionalArgumentLike],
	) ArgumentsLike
}

/*
ArithmeticClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete arithmetic-like class.
*/
type ArithmeticClassLike interface {
	// Constructor Methods
	Arithmetic(
		expression1 ExpressionLike,
		arithmeticOperator string,
		expression2 ExpressionLike,
	) ArithmeticLike
}

/*
AssociationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete association-like class.
*/
type AssociationClassLike interface {
	// Constructor Methods
	Association(
		key KeyLike,
		value ValueLike,
	) AssociationLike
}

/*
AssociationsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete associations-like class.
*/
type AssociationsClassLike interface {
	// Constructor Methods
	Associations(
		any_ any,
	) AssociationsLike
}

/*
AtLevelClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete at-level-like class.
*/
type AtLevelClassLike interface {
	// Constructor Methods
	AtLevel(
		expression ExpressionLike,
	) AtLevelLike
}

/*
AttributeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete attribute-like class.
*/
type AttributeClassLike interface {
	// Constructor Methods
	Attribute(
		variable VariableLike,
		indices IndicesLike,
	) AttributeLike
}

/*
BagClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete bag-like class.
*/
type BagClassLike interface {
	// Constructor Methods
	Bag(
		expression ExpressionLike,
	) BagLike
}

/*
BreakClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete break-clause-like class.
*/
type BreakClauseClassLike interface {
	// Constructor Methods
	BreakClause() BreakClauseLike
}

/*
ChainingClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete chaining-like class.
*/
type ChainingClassLike interface {
	// Constructor Methods
	Chaining(
		expression1 ExpressionLike,
		expression2 ExpressionLike,
	) ChainingLike
}

/*
CheckoutClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete checkout-clause-like class.
*/
type CheckoutClauseClassLike interface {
	// Constructor Methods
	CheckoutClause(
		recipient RecipientLike,
		optionalAtLevel AtLevelLike,
		citation CitationLike,
	) CheckoutClauseLike
}

/*
CitationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete citation-like class.
*/
type CitationClassLike interface {
	// Constructor Methods
	Citation(
		expression ExpressionLike,
	) CitationLike
}

/*
CollectionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete collection-like class.
*/
type CollectionClassLike interface {
	// Constructor Methods
	Collection(
		items ItemsLike,
	) CollectionLike
}

/*
ComparisonClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete comparison-like class.
*/
type ComparisonClassLike interface {
	// Constructor Methods
	Comparison(
		expression1 ExpressionLike,
		comparisonOperator string,
		expression2 ExpressionLike,
	) ComparisonLike
}

/*
ComplementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete complement-like class.
*/
type ComplementClassLike interface {
	// Constructor Methods
	Complement(
		expression ExpressionLike,
	) ComplementLike
}

/*
ComponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructor Methods
	Component(
		entity EntityLike,
		optionalContext ContextLike,
	) ComponentLike
}

/*
CompositeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete composite-like class.
*/
type CompositeClassLike interface {
	// Constructor Methods
	Composite(
		expression ExpressionLike,
	) CompositeLike
}

/*
ConditionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete condition-like class.
*/
type ConditionClassLike interface {
	// Constructor Methods
	Condition(
		expression ExpressionLike,
	) ConditionLike
}

/*
ContextClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete context-like class.
*/
type ContextClassLike interface {
	// Constructor Methods
	Context(
		parameters ParametersLike,
	) ContextLike
}

/*
ContinueClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete continue-clause-like class.
*/
type ContinueClauseClassLike interface {
	// Constructor Methods
	ContinueClause() ContinueClauseLike
}

/*
DereferenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete dereference-like class.
*/
type DereferenceClassLike interface {
	// Constructor Methods
	Dereference(
		expression ExpressionLike,
	) DereferenceLike
}

/*
DiscardClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete discard-clause-like class.
*/
type DiscardClauseClassLike interface {
	// Constructor Methods
	DiscardClause(
		draft DraftLike,
	) DiscardClauseLike
}

/*
DoClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete do-clause-like class.
*/
type DoClauseClassLike interface {
	// Constructor Methods
	DoClause(
		operation OperationLike,
	) DoClauseLike
}

/*
DocumentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructor Methods
	Document(
		comment string,
		component ComponentLike,
	) DocumentLike
}

/*
DraftClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete draft-like class.
*/
type DraftClassLike interface {
	// Constructor Methods
	Draft(
		expression ExpressionLike,
	) DraftLike
}

/*
ElementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete element-like class.
*/
type ElementClassLike interface {
	// Constructor Methods
	Element(
		any_ any,
	) ElementLike
}

/*
EmptyAssociationsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete empty-associations-like class.
*/
type EmptyAssociationsClassLike interface {
	// Constructor Methods
	EmptyAssociations() EmptyAssociationsLike
}

/*
EmptyStatementsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete empty-statements-like class.
*/
type EmptyStatementsClassLike interface {
	// Constructor Methods
	EmptyStatements() EmptyStatementsLike
}

/*
EmptyValuesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete empty-values-like class.
*/
type EmptyValuesClassLike interface {
	// Constructor Methods
	EmptyValues() EmptyValuesLike
}

/*
EntityClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete entity-like class.
*/
type EntityClassLike interface {
	// Constructor Methods
	Entity(
		any_ any,
	) EntityLike
}

/*
EventClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete event-like class.
*/
type EventClassLike interface {
	// Constructor Methods
	Event(
		expression ExpressionLike,
	) EventLike
}

/*
ExceptionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete exception-like class.
*/
type ExceptionClassLike interface {
	// Constructor Methods
	Exception(
		expression ExpressionLike,
	) ExceptionLike
}

/*
ExponentialClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete exponential-like class.
*/
type ExponentialClassLike interface {
	// Constructor Methods
	Exponential(
		expression1 ExpressionLike,
		expression2 ExpressionLike,
	) ExponentialLike
}

/*
ExpressionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete expression-like class.
*/
type ExpressionClassLike interface {
	// Constructor Methods
	Expression(
		any_ any,
	) ExpressionLike
}

/*
FailureClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete failure-like class.
*/
type FailureClassLike interface {
	// Constructor Methods
	Failure(
		symbol string,
	) FailureLike
}

/*
FlowClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete flow-like class.
*/
type FlowClassLike interface {
	// Constructor Methods
	Flow(
		any_ any,
	) FlowLike
}

/*
FunctionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructor Methods
	Function(
		identifier string,
	) FunctionLike
}

/*
IfClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete if-clause-like class.
*/
type IfClauseClassLike interface {
	// Constructor Methods
	IfClause(
		condition ConditionLike,
		procedure ProcedureLike,
	) IfClauseLike
}

/*
IndexClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete index-like class.
*/
type IndexClassLike interface {
	// Constructor Methods
	Index(
		expression ExpressionLike,
	) IndexLike
}

/*
IndicesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete indices-like class.
*/
type IndicesClassLike interface {
	// Constructor Methods
	Indices(
		index IndexLike,
		additionalIndexes col.Sequential[AdditionalIndexLike],
	) IndicesLike
}

/*
InductionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete induction-like class.
*/
type InductionClassLike interface {
	// Constructor Methods
	Induction(
		any_ any,
	) InductionLike
}

/*
InlineAssociationsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-associations-like class.
*/
type InlineAssociationsClassLike interface {
	// Constructor Methods
	InlineAssociations(
		association AssociationLike,
		additionalAssociations col.Sequential[AdditionalAssociationLike],
	) InlineAssociationsLike
}

/*
InlineParametersClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-parameters-like class.
*/
type InlineParametersClassLike interface {
	// Constructor Methods
	InlineParameters(
		parameter ParameterLike,
		additionalParameters col.Sequential[AdditionalParameterLike],
	) InlineParametersLike
}

/*
InlineStatementsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-statements-like class.
*/
type InlineStatementsClassLike interface {
	// Constructor Methods
	InlineStatements(
		statement StatementLike,
		additionalStatements col.Sequential[AdditionalStatementLike],
	) InlineStatementsLike
}

/*
InlineValuesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inline-values-like class.
*/
type InlineValuesClassLike interface {
	// Constructor Methods
	InlineValues(
		value ValueLike,
		additionalValues col.Sequential[AdditionalValueLike],
	) InlineValuesLike
}

/*
IntrinsicClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete intrinsic-like class.
*/
type IntrinsicClassLike interface {
	// Constructor Methods
	Intrinsic(
		function FunctionLike,
		optionalArguments ArgumentsLike,
	) IntrinsicLike
}

/*
InversionClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete inversion-like class.
*/
type InversionClassLike interface {
	// Constructor Methods
	Inversion(
		inversionOperator string,
		expression ExpressionLike,
	) InversionLike
}

/*
InvocationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete invocation-like class.
*/
type InvocationClassLike interface {
	// Constructor Methods
	Invocation(
		target TargetLike,
		invocationOperator string,
		method MethodLike,
		optionalArguments ArgumentsLike,
	) InvocationLike
}

/*
ItemClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete item-like class.
*/
type ItemClassLike interface {
	// Constructor Methods
	Item(
		symbol string,
	) ItemLike
}

/*
ItemsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete items-like class.
*/
type ItemsClassLike interface {
	// Constructor Methods
	Items(
		any_ any,
	) ItemsLike
}

/*
KeyClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete key-like class.
*/
type KeyClassLike interface {
	// Constructor Methods
	Key(
		primitive PrimitiveLike,
	) KeyLike
}

/*
LabelClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete label-like class.
*/
type LabelClassLike interface {
	// Constructor Methods
	Label(
		symbol string,
	) LabelLike
}

/*
LetClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete let-clause-like class.
*/
type LetClauseClassLike interface {
	// Constructor Methods
	LetClause(
		recipient RecipientLike,
		assignmentOperator string,
		expression ExpressionLike,
	) LetClauseLike
}

/*
LogicalClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete logical-like class.
*/
type LogicalClassLike interface {
	// Constructor Methods
	Logical(
		expression1 ExpressionLike,
		logicalOperator string,
		expression2 ExpressionLike,
	) LogicalLike
}

/*
MagnitudeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete magnitude-like class.
*/
type MagnitudeClassLike interface {
	// Constructor Methods
	Magnitude(
		expression ExpressionLike,
	) MagnitudeLike
}

/*
MainClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete main-clause-like class.
*/
type MainClauseClassLike interface {
	// Constructor Methods
	MainClause(
		any_ any,
	) MainClauseLike
}

/*
MatchHandlerClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete match-handler-like class.
*/
type MatchHandlerClassLike interface {
	// Constructor Methods
	MatchHandler(
		template TemplateLike,
		procedure ProcedureLike,
	) MatchHandlerLike
}

/*
MessageClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete message-like class.
*/
type MessageClassLike interface {
	// Constructor Methods
	Message(
		expression ExpressionLike,
	) MessageLike
}

/*
MessagingClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete messaging-like class.
*/
type MessagingClassLike interface {
	// Constructor Methods
	Messaging(
		any_ any,
	) MessagingLike
}

/*
MethodClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete method-like class.
*/
type MethodClassLike interface {
	// Constructor Methods
	Method(
		identifier string,
	) MethodLike
}

/*
MultilineAssociationsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multiline-associations-like class.
*/
type MultilineAssociationsClassLike interface {
	// Constructor Methods
	MultilineAssociations(
		annotatedAssociations col.Sequential[AnnotatedAssociationLike],
	) MultilineAssociationsLike
}

/*
MultilineParametersClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multiline-parameters-like class.
*/
type MultilineParametersClassLike interface {
	// Constructor Methods
	MultilineParameters(
		annotatedParameters col.Sequential[AnnotatedParameterLike],
	) MultilineParametersLike
}

/*
MultilineStatementsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multiline-statements-like class.
*/
type MultilineStatementsClassLike interface {
	// Constructor Methods
	MultilineStatements(
		annotatedStatements col.Sequential[AnnotatedStatementLike],
	) MultilineStatementsLike
}

/*
MultilineValuesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete multiline-values-like class.
*/
type MultilineValuesClassLike interface {
	// Constructor Methods
	MultilineValues(
		annotatedValues col.Sequential[AnnotatedValueLike],
	) MultilineValuesLike
}

/*
NotarizeClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete notarize-clause-like class.
*/
type NotarizeClauseClassLike interface {
	// Constructor Methods
	NotarizeClause(
		draft DraftLike,
		citation CitationLike,
	) NotarizeClauseLike
}

/*
OnClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete on-clause-like class.
*/
type OnClauseClassLike interface {
	// Constructor Methods
	OnClause(
		failure FailureLike,
		matchHandlers col.Sequential[MatchHandlerLike],
	) OnClauseLike
}

/*
OperationClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete operation-like class.
*/
type OperationClassLike interface {
	// Constructor Methods
	Operation(
		any_ any,
	) OperationLike
}

/*
ParameterClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructor Methods
	Parameter(
		label LabelLike,
		component ComponentLike,
	) ParameterLike
}

/*
ParametersClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete parameters-like class.
*/
type ParametersClassLike interface {
	// Constructor Methods
	Parameters(
		any_ any,
	) ParametersLike
}

/*
PostClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete post-clause-like class.
*/
type PostClauseClassLike interface {
	// Constructor Methods
	PostClause(
		message MessageLike,
		bag BagLike,
	) PostClauseLike
}

/*
PrecedenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete precedence-like class.
*/
type PrecedenceClassLike interface {
	// Constructor Methods
	Precedence(
		expression ExpressionLike,
	) PrecedenceLike
}

/*
PrimitiveClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete primitive-like class.
*/
type PrimitiveClassLike interface {
	// Constructor Methods
	Primitive(
		any_ any,
	) PrimitiveLike
}

/*
ProcedureClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete procedure-like class.
*/
type ProcedureClassLike interface {
	// Constructor Methods
	Procedure(
		statements StatementsLike,
	) ProcedureLike
}

/*
PublishClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete publish-clause-like class.
*/
type PublishClauseClassLike interface {
	// Constructor Methods
	PublishClause(
		event EventLike,
	) PublishClauseLike
}

/*
RangeClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete range-like class.
*/
type RangeClassLike interface {
	// Constructor Methods
	Range(
		open string,
		primitive1 PrimitiveLike,
		primitive2 PrimitiveLike,
		close_ string,
	) RangeLike
}

/*
RecipientClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete recipient-like class.
*/
type RecipientClassLike interface {
	// Constructor Methods
	Recipient(
		any_ any,
	) RecipientLike
}

/*
RejectClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete reject-clause-like class.
*/
type RejectClauseClassLike interface {
	// Constructor Methods
	RejectClause(
		message MessageLike,
	) RejectClauseLike
}

/*
RepositoryClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete repository-like class.
*/
type RepositoryClassLike interface {
	// Constructor Methods
	Repository(
		any_ any,
	) RepositoryLike
}

/*
ResultClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete result-like class.
*/
type ResultClassLike interface {
	// Constructor Methods
	Result(
		expression ExpressionLike,
	) ResultLike
}

/*
RetrieveClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete retrieve-clause-like class.
*/
type RetrieveClauseClassLike interface {
	// Constructor Methods
	RetrieveClause(
		recipient RecipientLike,
		bag BagLike,
	) RetrieveClauseLike
}

/*
ReturnClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete return-clause-like class.
*/
type ReturnClauseClassLike interface {
	// Constructor Methods
	ReturnClause(
		result ResultLike,
	) ReturnClauseLike
}

/*
SaveClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete save-clause-like class.
*/
type SaveClauseClassLike interface {
	// Constructor Methods
	SaveClause(
		draft DraftLike,
		citation CitationLike,
	) SaveClauseLike
}

/*
SelectClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete select-clause-like class.
*/
type SelectClauseClassLike interface {
	// Constructor Methods
	SelectClause(
		target TargetLike,
		matchHandlers col.Sequential[MatchHandlerLike],
	) SelectClauseLike
}

/*
SequenceClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete sequence-like class.
*/
type SequenceClassLike interface {
	// Constructor Methods
	Sequence(
		expression ExpressionLike,
	) SequenceLike
}

/*
StatementClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete statement-like class.
*/
type StatementClassLike interface {
	// Constructor Methods
	Statement(
		mainClause MainClauseLike,
		optionalOnClause OnClauseLike,
	) StatementLike
}

/*
StatementsClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete statements-like class.
*/
type StatementsClassLike interface {
	// Constructor Methods
	Statements(
		any_ any,
	) StatementsLike
}

/*
StringClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete string-like class.
*/
type StringClassLike interface {
	// Constructor Methods
	String(
		any_ any,
	) StringLike
}

/*
SubcomponentClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete subcomponent-like class.
*/
type SubcomponentClassLike interface {
	// Constructor Methods
	Subcomponent(
		composite CompositeLike,
		indices IndicesLike,
	) SubcomponentLike
}

/*
TargetClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete target-like class.
*/
type TargetClassLike interface {
	// Constructor Methods
	Target(
		expression ExpressionLike,
	) TargetLike
}

/*
TemplateClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete template-like class.
*/
type TemplateClassLike interface {
	// Constructor Methods
	Template(
		expression ExpressionLike,
	) TemplateLike
}

/*
ThrowClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete throw-clause-like class.
*/
type ThrowClauseClassLike interface {
	// Constructor Methods
	ThrowClause(
		exception ExceptionLike,
	) ThrowClauseLike
}

/*
ValueClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete value-like class.
*/
type ValueClassLike interface {
	// Constructor Methods
	Value(
		component ComponentLike,
	) ValueLike
}

/*
ValuesClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete values-like class.
*/
type ValuesClassLike interface {
	// Constructor Methods
	Values(
		any_ any,
	) ValuesLike
}

/*
VariableClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete variable-like class.
*/
type VariableClassLike interface {
	// Constructor Methods
	Variable(
		identifier string,
	) VariableLike
}

/*
WhileClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete while-clause-like class.
*/
type WhileClauseClassLike interface {
	// Constructor Methods
	WhileClause(
		condition ConditionLike,
		procedure ProcedureLike,
	) WhileClauseLike
}

/*
WithClauseClassLike is a class interface that declares the
complete set of class constructors, constants and functions that must be
supported by each concrete with-clause-like class.
*/
type WithClauseClassLike interface {
	// Constructor Methods
	WithClause(
		item ItemLike,
		sequence SequenceLike,
		procedure ProcedureLike,
	) WithClauseLike
}

// INSTANCE DECLARATIONS

/*
AcceptClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete accept-clause-like class.
*/
type AcceptClauseLike interface {
	// Principal Methods
	GetClass() AcceptClauseClassLike

	// Attribute Methods
	GetMessage() MessageLike
}

/*
AdditionalArgumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-argument-like class.
*/
type AdditionalArgumentLike interface {
	// Principal Methods
	GetClass() AdditionalArgumentClassLike

	// Attribute Methods
	GetArgument() ArgumentLike
}

/*
AdditionalAssociationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-association-like class.
*/
type AdditionalAssociationLike interface {
	// Principal Methods
	GetClass() AdditionalAssociationClassLike

	// Attribute Methods
	GetAssociation() AssociationLike
}

/*
AdditionalIndexLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-index-like class.
*/
type AdditionalIndexLike interface {
	// Principal Methods
	GetClass() AdditionalIndexClassLike

	// Attribute Methods
	GetIndex() IndexLike
}

/*
AdditionalParameterLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-parameter-like class.
*/
type AdditionalParameterLike interface {
	// Principal Methods
	GetClass() AdditionalParameterClassLike

	// Attribute Methods
	GetParameter() ParameterLike
}

/*
AdditionalStatementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-statement-like class.
*/
type AdditionalStatementLike interface {
	// Principal Methods
	GetClass() AdditionalStatementClassLike

	// Attribute Methods
	GetStatement() StatementLike
}

/*
AdditionalValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete additional-value-like class.
*/
type AdditionalValueLike interface {
	// Principal Methods
	GetClass() AdditionalValueClassLike

	// Attribute Methods
	GetValue() ValueLike
}

/*
AnnotatedAssociationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete annotated-association-like class.
*/
type AnnotatedAssociationLike interface {
	// Principal Methods
	GetClass() AnnotatedAssociationClassLike

	// Attribute Methods
	GetAssociation() AssociationLike
	GetOptionalNote() string
}

/*
AnnotatedParameterLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete annotated-parameter-like class.
*/
type AnnotatedParameterLike interface {
	// Principal Methods
	GetClass() AnnotatedParameterClassLike

	// Attribute Methods
	GetParameter() ParameterLike
	GetOptionalNote() string
}

/*
AnnotatedStatementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete annotated-statement-like class.
*/
type AnnotatedStatementLike interface {
	// Principal Methods
	GetClass() AnnotatedStatementClassLike

	// Attribute Methods
	GetStatement() StatementLike
	GetOptionalNote() string
}

/*
AnnotatedValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete annotated-value-like class.
*/
type AnnotatedValueLike interface {
	// Principal Methods
	GetClass() AnnotatedValueClassLike

	// Attribute Methods
	GetValue() ValueLike
	GetOptionalNote() string
}

/*
ArgumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Principal Methods
	GetClass() ArgumentClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ArgumentsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Principal Methods
	GetClass() ArgumentsClassLike

	// Attribute Methods
	GetArgument() ArgumentLike
	GetAdditionalArguments() col.Sequential[AdditionalArgumentLike]
}

/*
ArithmeticLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete arithmetic-like class.
*/
type ArithmeticLike interface {
	// Principal Methods
	GetClass() ArithmeticClassLike

	// Attribute Methods
	GetExpression1() ExpressionLike
	GetArithmeticOperator() string
	GetExpression2() ExpressionLike
}

/*
AssociationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete association-like class.
*/
type AssociationLike interface {
	// Principal Methods
	GetClass() AssociationClassLike

	// Attribute Methods
	GetKey() KeyLike
	GetValue() ValueLike
}

/*
AssociationsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete associations-like class.
*/
type AssociationsLike interface {
	// Principal Methods
	GetClass() AssociationsClassLike

	// Attribute Methods
	GetAny() any
}

/*
AtLevelLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete at-level-like class.
*/
type AtLevelLike interface {
	// Principal Methods
	GetClass() AtLevelClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
AttributeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete attribute-like class.
*/
type AttributeLike interface {
	// Principal Methods
	GetClass() AttributeClassLike

	// Attribute Methods
	GetVariable() VariableLike
	GetIndices() IndicesLike
}

/*
BagLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete bag-like class.
*/
type BagLike interface {
	// Principal Methods
	GetClass() BagClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
BreakClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete break-clause-like class.
*/
type BreakClauseLike interface {
	// Principal Methods
	GetClass() BreakClauseClassLike
}

/*
ChainingLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete chaining-like class.
*/
type ChainingLike interface {
	// Principal Methods
	GetClass() ChainingClassLike

	// Attribute Methods
	GetExpression1() ExpressionLike
	GetExpression2() ExpressionLike
}

/*
CheckoutClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete checkout-clause-like class.
*/
type CheckoutClauseLike interface {
	// Principal Methods
	GetClass() CheckoutClauseClassLike

	// Attribute Methods
	GetRecipient() RecipientLike
	GetOptionalAtLevel() AtLevelLike
	GetCitation() CitationLike
}

/*
CitationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete citation-like class.
*/
type CitationLike interface {
	// Principal Methods
	GetClass() CitationClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
CollectionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete collection-like class.
*/
type CollectionLike interface {
	// Principal Methods
	GetClass() CollectionClassLike

	// Attribute Methods
	GetItems() ItemsLike
}

/*
ComparisonLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete comparison-like class.
*/
type ComparisonLike interface {
	// Principal Methods
	GetClass() ComparisonClassLike

	// Attribute Methods
	GetExpression1() ExpressionLike
	GetComparisonOperator() string
	GetExpression2() ExpressionLike
}

/*
ComplementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete complement-like class.
*/
type ComplementLike interface {
	// Principal Methods
	GetClass() ComplementClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ComponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Principal Methods
	GetClass() ComponentClassLike

	// Attribute Methods
	GetEntity() EntityLike
	GetOptionalContext() ContextLike
}

/*
CompositeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete composite-like class.
*/
type CompositeLike interface {
	// Principal Methods
	GetClass() CompositeClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ConditionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete condition-like class.
*/
type ConditionLike interface {
	// Principal Methods
	GetClass() ConditionClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ContextLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete context-like class.
*/
type ContextLike interface {
	// Principal Methods
	GetClass() ContextClassLike

	// Attribute Methods
	GetParameters() ParametersLike
}

/*
ContinueClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete continue-clause-like class.
*/
type ContinueClauseLike interface {
	// Principal Methods
	GetClass() ContinueClauseClassLike
}

/*
DereferenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete dereference-like class.
*/
type DereferenceLike interface {
	// Principal Methods
	GetClass() DereferenceClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
DiscardClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete discard-clause-like class.
*/
type DiscardClauseLike interface {
	// Principal Methods
	GetClass() DiscardClauseClassLike

	// Attribute Methods
	GetDraft() DraftLike
}

/*
DoClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete do-clause-like class.
*/
type DoClauseLike interface {
	// Principal Methods
	GetClass() DoClauseClassLike

	// Attribute Methods
	GetOperation() OperationLike
}

/*
DocumentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete document-like class.
*/
type DocumentLike interface {
	// Principal Methods
	GetClass() DocumentClassLike

	// Attribute Methods
	GetComment() string
	GetComponent() ComponentLike
}

/*
DraftLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete draft-like class.
*/
type DraftLike interface {
	// Principal Methods
	GetClass() DraftClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ElementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete element-like class.
*/
type ElementLike interface {
	// Principal Methods
	GetClass() ElementClassLike

	// Attribute Methods
	GetAny() any
}

/*
EmptyAssociationsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete empty-associations-like class.
*/
type EmptyAssociationsLike interface {
	// Principal Methods
	GetClass() EmptyAssociationsClassLike
}

/*
EmptyStatementsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete empty-statements-like class.
*/
type EmptyStatementsLike interface {
	// Principal Methods
	GetClass() EmptyStatementsClassLike
}

/*
EmptyValuesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete empty-values-like class.
*/
type EmptyValuesLike interface {
	// Principal Methods
	GetClass() EmptyValuesClassLike
}

/*
EntityLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete entity-like class.
*/
type EntityLike interface {
	// Principal Methods
	GetClass() EntityClassLike

	// Attribute Methods
	GetAny() any
}

/*
EventLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete event-like class.
*/
type EventLike interface {
	// Principal Methods
	GetClass() EventClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ExceptionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete exception-like class.
*/
type ExceptionLike interface {
	// Principal Methods
	GetClass() ExceptionClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ExponentialLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete exponential-like class.
*/
type ExponentialLike interface {
	// Principal Methods
	GetClass() ExponentialClassLike

	// Attribute Methods
	GetExpression1() ExpressionLike
	GetExpression2() ExpressionLike
}

/*
ExpressionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete expression-like class.
*/
type ExpressionLike interface {
	// Principal Methods
	GetClass() ExpressionClassLike

	// Attribute Methods
	GetAny() any
}

/*
FailureLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete failure-like class.
*/
type FailureLike interface {
	// Principal Methods
	GetClass() FailureClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
FlowLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete flow-like class.
*/
type FlowLike interface {
	// Principal Methods
	GetClass() FlowClassLike

	// Attribute Methods
	GetAny() any
}

/*
FunctionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete function-like class.
*/
type FunctionLike interface {
	// Principal Methods
	GetClass() FunctionClassLike

	// Attribute Methods
	GetIdentifier() string
}

/*
IfClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete if-clause-like class.
*/
type IfClauseLike interface {
	// Principal Methods
	GetClass() IfClauseClassLike

	// Attribute Methods
	GetCondition() ConditionLike
	GetProcedure() ProcedureLike
}

/*
IndexLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete index-like class.
*/
type IndexLike interface {
	// Principal Methods
	GetClass() IndexClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
IndicesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete indices-like class.
*/
type IndicesLike interface {
	// Principal Methods
	GetClass() IndicesClassLike

	// Attribute Methods
	GetIndex() IndexLike
	GetAdditionalIndexes() col.Sequential[AdditionalIndexLike]
}

/*
InductionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete induction-like class.
*/
type InductionLike interface {
	// Principal Methods
	GetClass() InductionClassLike

	// Attribute Methods
	GetAny() any
}

/*
InlineAssociationsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-associations-like class.
*/
type InlineAssociationsLike interface {
	// Principal Methods
	GetClass() InlineAssociationsClassLike

	// Attribute Methods
	GetAssociation() AssociationLike
	GetAdditionalAssociations() col.Sequential[AdditionalAssociationLike]
}

/*
InlineParametersLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-parameters-like class.
*/
type InlineParametersLike interface {
	// Principal Methods
	GetClass() InlineParametersClassLike

	// Attribute Methods
	GetParameter() ParameterLike
	GetAdditionalParameters() col.Sequential[AdditionalParameterLike]
}

/*
InlineStatementsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-statements-like class.
*/
type InlineStatementsLike interface {
	// Principal Methods
	GetClass() InlineStatementsClassLike

	// Attribute Methods
	GetStatement() StatementLike
	GetAdditionalStatements() col.Sequential[AdditionalStatementLike]
}

/*
InlineValuesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inline-values-like class.
*/
type InlineValuesLike interface {
	// Principal Methods
	GetClass() InlineValuesClassLike

	// Attribute Methods
	GetValue() ValueLike
	GetAdditionalValues() col.Sequential[AdditionalValueLike]
}

/*
IntrinsicLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete intrinsic-like class.
*/
type IntrinsicLike interface {
	// Principal Methods
	GetClass() IntrinsicClassLike

	// Attribute Methods
	GetFunction() FunctionLike
	GetOptionalArguments() ArgumentsLike
}

/*
InversionLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete inversion-like class.
*/
type InversionLike interface {
	// Principal Methods
	GetClass() InversionClassLike

	// Attribute Methods
	GetInversionOperator() string
	GetExpression() ExpressionLike
}

/*
InvocationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete invocation-like class.
*/
type InvocationLike interface {
	// Principal Methods
	GetClass() InvocationClassLike

	// Attribute Methods
	GetTarget() TargetLike
	GetInvocationOperator() string
	GetMethod() MethodLike
	GetOptionalArguments() ArgumentsLike
}

/*
ItemLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete item-like class.
*/
type ItemLike interface {
	// Principal Methods
	GetClass() ItemClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
ItemsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete items-like class.
*/
type ItemsLike interface {
	// Principal Methods
	GetClass() ItemsClassLike

	// Attribute Methods
	GetAny() any
}

/*
KeyLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete key-like class.
*/
type KeyLike interface {
	// Principal Methods
	GetClass() KeyClassLike

	// Attribute Methods
	GetPrimitive() PrimitiveLike
}

/*
LabelLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete label-like class.
*/
type LabelLike interface {
	// Principal Methods
	GetClass() LabelClassLike

	// Attribute Methods
	GetSymbol() string
}

/*
LetClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete let-clause-like class.
*/
type LetClauseLike interface {
	// Principal Methods
	GetClass() LetClauseClassLike

	// Attribute Methods
	GetRecipient() RecipientLike
	GetAssignmentOperator() string
	GetExpression() ExpressionLike
}

/*
LogicalLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete logical-like class.
*/
type LogicalLike interface {
	// Principal Methods
	GetClass() LogicalClassLike

	// Attribute Methods
	GetExpression1() ExpressionLike
	GetLogicalOperator() string
	GetExpression2() ExpressionLike
}

/*
MagnitudeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete magnitude-like class.
*/
type MagnitudeLike interface {
	// Principal Methods
	GetClass() MagnitudeClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
MainClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete main-clause-like class.
*/
type MainClauseLike interface {
	// Principal Methods
	GetClass() MainClauseClassLike

	// Attribute Methods
	GetAny() any
}

/*
MatchHandlerLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete match-handler-like class.
*/
type MatchHandlerLike interface {
	// Principal Methods
	GetClass() MatchHandlerClassLike

	// Attribute Methods
	GetTemplate() TemplateLike
	GetProcedure() ProcedureLike
}

/*
MessageLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete message-like class.
*/
type MessageLike interface {
	// Principal Methods
	GetClass() MessageClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
MessagingLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete messaging-like class.
*/
type MessagingLike interface {
	// Principal Methods
	GetClass() MessagingClassLike

	// Attribute Methods
	GetAny() any
}

/*
MethodLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete method-like class.
*/
type MethodLike interface {
	// Principal Methods
	GetClass() MethodClassLike

	// Attribute Methods
	GetIdentifier() string
}

/*
MultilineAssociationsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multiline-associations-like class.
*/
type MultilineAssociationsLike interface {
	// Principal Methods
	GetClass() MultilineAssociationsClassLike

	// Attribute Methods
	GetAnnotatedAssociations() col.Sequential[AnnotatedAssociationLike]
}

/*
MultilineParametersLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multiline-parameters-like class.
*/
type MultilineParametersLike interface {
	// Principal Methods
	GetClass() MultilineParametersClassLike

	// Attribute Methods
	GetAnnotatedParameters() col.Sequential[AnnotatedParameterLike]
}

/*
MultilineStatementsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multiline-statements-like class.
*/
type MultilineStatementsLike interface {
	// Principal Methods
	GetClass() MultilineStatementsClassLike

	// Attribute Methods
	GetAnnotatedStatements() col.Sequential[AnnotatedStatementLike]
}

/*
MultilineValuesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete multiline-values-like class.
*/
type MultilineValuesLike interface {
	// Principal Methods
	GetClass() MultilineValuesClassLike

	// Attribute Methods
	GetAnnotatedValues() col.Sequential[AnnotatedValueLike]
}

/*
NotarizeClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete notarize-clause-like class.
*/
type NotarizeClauseLike interface {
	// Principal Methods
	GetClass() NotarizeClauseClassLike

	// Attribute Methods
	GetDraft() DraftLike
	GetCitation() CitationLike
}

/*
OnClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete on-clause-like class.
*/
type OnClauseLike interface {
	// Principal Methods
	GetClass() OnClauseClassLike

	// Attribute Methods
	GetFailure() FailureLike
	GetMatchHandlers() col.Sequential[MatchHandlerLike]
}

/*
OperationLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete operation-like class.
*/
type OperationLike interface {
	// Principal Methods
	GetClass() OperationClassLike

	// Attribute Methods
	GetAny() any
}

/*
ParameterLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Principal Methods
	GetClass() ParameterClassLike

	// Attribute Methods
	GetLabel() LabelLike
	GetComponent() ComponentLike
}

/*
ParametersLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete parameters-like class.
*/
type ParametersLike interface {
	// Principal Methods
	GetClass() ParametersClassLike

	// Attribute Methods
	GetAny() any
}

/*
PostClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete post-clause-like class.
*/
type PostClauseLike interface {
	// Principal Methods
	GetClass() PostClauseClassLike

	// Attribute Methods
	GetMessage() MessageLike
	GetBag() BagLike
}

/*
PrecedenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete precedence-like class.
*/
type PrecedenceLike interface {
	// Principal Methods
	GetClass() PrecedenceClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
PrimitiveLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete primitive-like class.
*/
type PrimitiveLike interface {
	// Principal Methods
	GetClass() PrimitiveClassLike

	// Attribute Methods
	GetAny() any
}

/*
ProcedureLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete procedure-like class.
*/
type ProcedureLike interface {
	// Principal Methods
	GetClass() ProcedureClassLike

	// Attribute Methods
	GetStatements() StatementsLike
}

/*
PublishClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete publish-clause-like class.
*/
type PublishClauseLike interface {
	// Principal Methods
	GetClass() PublishClauseClassLike

	// Attribute Methods
	GetEvent() EventLike
}

/*
RangeLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete range-like class.
*/
type RangeLike interface {
	// Principal Methods
	GetClass() RangeClassLike

	// Attribute Methods
	GetOpen() string
	GetPrimitive1() PrimitiveLike
	GetPrimitive2() PrimitiveLike
	GetClose() string
}

/*
RecipientLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete recipient-like class.
*/
type RecipientLike interface {
	// Principal Methods
	GetClass() RecipientClassLike

	// Attribute Methods
	GetAny() any
}

/*
RejectClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete reject-clause-like class.
*/
type RejectClauseLike interface {
	// Principal Methods
	GetClass() RejectClauseClassLike

	// Attribute Methods
	GetMessage() MessageLike
}

/*
RepositoryLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete repository-like class.
*/
type RepositoryLike interface {
	// Principal Methods
	GetClass() RepositoryClassLike

	// Attribute Methods
	GetAny() any
}

/*
ResultLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete result-like class.
*/
type ResultLike interface {
	// Principal Methods
	GetClass() ResultClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
RetrieveClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete retrieve-clause-like class.
*/
type RetrieveClauseLike interface {
	// Principal Methods
	GetClass() RetrieveClauseClassLike

	// Attribute Methods
	GetRecipient() RecipientLike
	GetBag() BagLike
}

/*
ReturnClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete return-clause-like class.
*/
type ReturnClauseLike interface {
	// Principal Methods
	GetClass() ReturnClauseClassLike

	// Attribute Methods
	GetResult() ResultLike
}

/*
SaveClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete save-clause-like class.
*/
type SaveClauseLike interface {
	// Principal Methods
	GetClass() SaveClauseClassLike

	// Attribute Methods
	GetDraft() DraftLike
	GetCitation() CitationLike
}

/*
SelectClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete select-clause-like class.
*/
type SelectClauseLike interface {
	// Principal Methods
	GetClass() SelectClauseClassLike

	// Attribute Methods
	GetTarget() TargetLike
	GetMatchHandlers() col.Sequential[MatchHandlerLike]
}

/*
SequenceLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete sequence-like class.
*/
type SequenceLike interface {
	// Principal Methods
	GetClass() SequenceClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
StatementLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete statement-like class.
*/
type StatementLike interface {
	// Principal Methods
	GetClass() StatementClassLike

	// Attribute Methods
	GetMainClause() MainClauseLike
	GetOptionalOnClause() OnClauseLike
}

/*
StatementsLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete statements-like class.
*/
type StatementsLike interface {
	// Principal Methods
	GetClass() StatementsClassLike

	// Attribute Methods
	GetAny() any
}

/*
StringLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete string-like class.
*/
type StringLike interface {
	// Principal Methods
	GetClass() StringClassLike

	// Attribute Methods
	GetAny() any
}

/*
SubcomponentLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete subcomponent-like class.
*/
type SubcomponentLike interface {
	// Principal Methods
	GetClass() SubcomponentClassLike

	// Attribute Methods
	GetComposite() CompositeLike
	GetIndices() IndicesLike
}

/*
TargetLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete target-like class.
*/
type TargetLike interface {
	// Principal Methods
	GetClass() TargetClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
TemplateLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete template-like class.
*/
type TemplateLike interface {
	// Principal Methods
	GetClass() TemplateClassLike

	// Attribute Methods
	GetExpression() ExpressionLike
}

/*
ThrowClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete throw-clause-like class.
*/
type ThrowClauseLike interface {
	// Principal Methods
	GetClass() ThrowClauseClassLike

	// Attribute Methods
	GetException() ExceptionLike
}

/*
ValueLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete value-like class.
*/
type ValueLike interface {
	// Principal Methods
	GetClass() ValueClassLike

	// Attribute Methods
	GetComponent() ComponentLike
}

/*
ValuesLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete values-like class.
*/
type ValuesLike interface {
	// Principal Methods
	GetClass() ValuesClassLike

	// Attribute Methods
	GetAny() any
}

/*
VariableLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete variable-like class.
*/
type VariableLike interface {
	// Principal Methods
	GetClass() VariableClassLike

	// Attribute Methods
	GetIdentifier() string
}

/*
WhileClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete while-clause-like class.
*/
type WhileClauseLike interface {
	// Principal Methods
	GetClass() WhileClauseClassLike

	// Attribute Methods
	GetCondition() ConditionLike
	GetProcedure() ProcedureLike
}

/*
WithClauseLike is an instance interface that declares the
complete set of principal, attribute and aspect methods that must be supported
by each instance of a concrete with-clause-like class.
*/
type WithClauseLike interface {
	// Principal Methods
	GetClass() WithClauseClassLike

	// Attribute Methods
	GetItem() ItemLike
	GetSequence() SequenceLike
	GetProcedure() ProcedureLike
}

// ASPECT DECLARATIONS
