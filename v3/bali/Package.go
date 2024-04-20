/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "bali" provides...

For detailed documentation on this package refer to the wiki:
  - <wiki>

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

Additional concrete implementations of the classes defined by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package bali

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// Types

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	AngleToken
	BinaryToken
	BooleanToken
	BytecodeToken
	CommentToken
	DelimiterToken
	DurationToken
	EOFToken
	EOLToken
	IdentifierToken
	MomentToken
	NameToken
	NarrativeToken
	NoteToken
	NumberToken
	PatternToken
	PercentageToken
	ProbabilityToken
	QuoteToken
	ResourceToken
	SpaceToken
	SymbolToken
	TagToken
	VersionToken
)

// Classes

/*
AcceptClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete acceptclause-like class.
*/
type AcceptClauseClassLike interface {
	// Constructors
	MakeWithMessage(message MessageLike) AcceptClauseLike
}

/*
AnnotationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete annotation-like class.
*/
type AnnotationClassLike interface {
	// Constructors
	MakeWithAttributes(
		note string,
		comment string,
	) AnnotationLike
}

/*
ArgumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete argument-like class.
*/
type ArgumentClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) ArgumentLike
}

/*
ArgumentsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete arguments-like class.
*/
type ArgumentsClassLike interface {
	// Constructors
	MakeWithArguments(arguments col.ListLike[ArgumentLike]) ArgumentsLike
}

/*
ArithmeticClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete arithmetic-like class.
*/
type ArithmeticClassLike interface {
	// Constructors
	MakeWithExpressions(expressions col.ListLike[ExpressionLike]) ArithmeticLike
}

/*
AssignmentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete assignment-like class.
*/
type AssignmentClassLike interface {
	// Constructors
	MakeWithLetClause(letClause LetClauseLike) AssignmentLike
}

/*
AssociationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete association-like class.
*/
type AssociationClassLike interface {
	// Constructors
	MakeWithAttributes(
		key KeyLike,
		value ValueLike,
	) AssociationLike
}

/*
AssociationsClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete associations-like class.
*/
type AssociationsClassLike interface {
	// Constructors
	MakeWithAssociations(associations col.ListLike[AssociationLike]) AssociationsLike
	MakeWithAttributes(
		association AssociationLike,
		note string,
	) AssociationsLike
}

/*
AttributeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete attribute-like class.
*/
type AttributeClassLike interface {
	// Constructors
	MakeWithAttributes(
		variable VariableLike,
		indices IndicesLike,
	) AttributeLike
}

/*
BagClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete bag-like class.
*/
type BagClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) BagLike
}

/*
BreakClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete breakclause-like class.
*/
type BreakClauseClassLike interface {
	// Constructors
	Make() BreakClauseLike
}

/*
ChainingClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete chaining-like class.
*/
type ChainingClassLike interface {
	// Constructors
	MakeWithExpressions(expressions col.ListLike[ExpressionLike]) ChainingLike
}

/*
CheckoutClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete checkoutclause-like class.
*/
type CheckoutClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		recipient RecipientLike,
		level LevelLike,
		citation CitationLike,
	) CheckoutClauseLike
}

/*
CitationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete citation-like class.
*/
type CitationClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) CitationLike
}

/*
CollectionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete collection-like class.
*/
type CollectionClassLike interface {
	// Constructors
	MakeWithAttributes(
		associations AssociationsLike,
		values ValuesLike,
	) CollectionLike
}

/*
ComparisonClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete comparison-like class.
*/
type ComparisonClassLike interface {
	// Constructors
	MakeWithExpressions(expressions col.ListLike[ExpressionLike]) ComparisonLike
}

/*
ComplementClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete complement-like class.
*/
type ComplementClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) ComplementLike
}

/*
ComponentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete component-like class.
*/
type ComponentClassLike interface {
	// Constructors
	MakeWithAttributes(
		entity EntityLike,
		context ContextLike,
	) ComponentLike
}

/*
CompositeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete composite-like class.
*/
type CompositeClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) CompositeLike
}

/*
ConditionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete condition-like class.
*/
type ConditionClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) ConditionLike
}

/*
ContextClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete context-like class.
*/
type ContextClassLike interface {
	// Constructors
	MakeWithParameters(parameters ParametersLike) ContextLike
}

/*
ContinueClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete continueclause-like class.
*/
type ContinueClauseClassLike interface {
	// Constructors
	Make() ContinueClauseLike
}

/*
DereferenceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete dereference-like class.
*/
type DereferenceClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) DereferenceLike
}

/*
DiscardClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete discardclause-like class.
*/
type DiscardClauseClassLike interface {
	// Constructors
	MakeWithDraft(draft DraftLike) DiscardClauseLike
}

/*
DocumentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete document-like class.
*/
type DocumentClassLike interface {
	// Constructors
	MakeWithAttributes(
		header HeaderLike,
		component ComponentLike,
	) DocumentLike
}

/*
DraftClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete draft-like class.
*/
type DraftClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) DraftLike
}

/*
ElementClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete element-like class.
*/
type ElementClassLike interface {
	// Constructors
	MakeWithAngle(angle string) ElementLike
	MakeWithBoolean(boolean string) ElementLike
	MakeWithDuration(duration string) ElementLike
	MakeWithMoment(moment string) ElementLike
	MakeWithNumber(number string) ElementLike
	MakeWithPattern(pattern string) ElementLike
	MakeWithPercentage(percentage string) ElementLike
	MakeWithProbability(probability string) ElementLike
	MakeWithResource(resource string) ElementLike
}

/*
EntityClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete entity-like class.
*/
type EntityClassLike interface {
	// Constructors
	MakeWithElement(element ElementLike) EntityLike
	MakeWithString(string_ StringLike) EntityLike
	MakeWithRange(range_ RangeLike) EntityLike
	MakeWithCollection(collection CollectionLike) EntityLike
	MakeWithProcedure(procedure ProcedureLike) EntityLike
}

/*
EventClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete event-like class.
*/
type EventClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) EventLike
}

/*
ExceptionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete exception-like class.
*/
type ExceptionClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) ExceptionLike
}

/*
ExponentialClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete exponential-like class.
*/
type ExponentialClassLike interface {
	// Constructors
	MakeWithExpressions(expressions col.ListLike[ExpressionLike]) ExponentialLike
}

/*
ExpressionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete expression-like class.
*/
type ExpressionClassLike interface {
	// Constructors
	MakeWithComponent(component ComponentLike) ExpressionLike
	MakeWithIntrinsic(intrinsic IntrinsicLike) ExpressionLike
	MakeWithVariable(variable VariableLike) ExpressionLike
	MakeWithPrecedence(precedence PrecedenceLike) ExpressionLike
	MakeWithDereference(dereference DereferenceLike) ExpressionLike
	MakeWithInvocation(invocation InvocationLike) ExpressionLike
	MakeWithSubcomponent(subcomponent SubcomponentLike) ExpressionLike
	MakeWithChaining(chaining ChainingLike) ExpressionLike
	MakeWithExponential(exponential ExponentialLike) ExpressionLike
	MakeWithInversion(inversion InversionLike) ExpressionLike
	MakeWithArithmetic(arithmetic ArithmeticLike) ExpressionLike
	MakeWithMagnitude(magnitude MagnitudeLike) ExpressionLike
	MakeWithComparison(comparison ComparisonLike) ExpressionLike
	MakeWithComplement(complement ComplementLike) ExpressionLike
	MakeWithLogical(logical LogicalLike) ExpressionLike
}

/*
FailureClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete failure-like class.
*/
type FailureClassLike interface {
	// Constructors
	MakeWithSymbol(symbol string) FailureLike
}

/*
FlowClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete flow-like class.
*/
type FlowClassLike interface {
	// Constructors
	MakeWithIfClause(ifClause IfClauseLike) FlowLike
	MakeWithSelectClause(selectClause SelectClauseLike) FlowLike
	MakeWithWhileClause(whileClause WhileClauseLike) FlowLike
	MakeWithWithClause(withClause WithClauseLike) FlowLike
	MakeWithContinueClause(continueClause ContinueClauseLike) FlowLike
	MakeWithBreakClause(breakClause BreakClauseLike) FlowLike
	MakeWithReturnClause(returnClause ReturnClauseLike) FlowLike
	MakeWithThrowClause(throwClause ThrowClauseLike) FlowLike
}

/*
FormatterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructors
	Make() FormatterLike
}

/*
FunctionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete function-like class.
*/
type FunctionClassLike interface {
	// Constructors
	MakeWithIdentifier(identifier string) FunctionLike
}

/*
HeaderClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete header-like class.
*/
type HeaderClassLike interface {
	// Constructors
	MakeWithComment(comment string) HeaderLike
}

/*
IfClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete ifclause-like class.
*/
type IfClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		condition ConditionLike,
		procedure ProcedureLike,
	) IfClauseLike
}

/*
IndexClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete index-like class.
*/
type IndexClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) IndexLike
}

/*
IndicesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete indices-like class.
*/
type IndicesClassLike interface {
	// Constructors
	MakeWithIndexs(indexs col.ListLike[IndexLike]) IndicesLike
}

/*
IntrinsicClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete intrinsic-like class.
*/
type IntrinsicClassLike interface {
	// Constructors
	MakeWithAttributes(
		function FunctionLike,
		arguments ArgumentsLike,
	) IntrinsicLike
}

/*
InversionClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete inversion-like class.
*/
type InversionClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) InversionLike
}

/*
InvocationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete invocation-like class.
*/
type InvocationClassLike interface {
	// Constructors
	MakeWithAttributes(
		target TargetLike,
		method MethodLike,
		arguments ArgumentsLike,
	) InvocationLike
}

/*
ItemClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete item-like class.
*/
type ItemClassLike interface {
	// Constructors
	MakeWithSymbol(symbol string) ItemLike
}

/*
KeyClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete key-like class.
*/
type KeyClassLike interface {
	// Constructors
	MakeWithPrimitive(primitive PrimitiveLike) KeyLike
}

/*
LetClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete letclause-like class.
*/
type LetClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		recipient RecipientLike,
		expression ExpressionLike,
	) LetClauseLike
}

/*
LevelClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete level-like class.
*/
type LevelClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) LevelLike
}

/*
LineClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete line-like class.
*/
type LineClassLike interface {
	// Constructors
	MakeWithAttributes(
		annotation AnnotationLike,
		statement StatementLike,
	) LineLike
}

/*
LinesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete lines-like class.
*/
type LinesClassLike interface {
	// Constructors
	MakeWithLines(lines col.ListLike[LineLike]) LinesLike
	MakeWithAttributes(
		line LineLike,
		note string,
	) LinesLike
}

/*
LogicalClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete logical-like class.
*/
type LogicalClassLike interface {
	// Constructors
	MakeWithExpressions(expressions col.ListLike[ExpressionLike]) LogicalLike
}

/*
MagnitudeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete magnitude-like class.
*/
type MagnitudeClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) MagnitudeLike
}

/*
MainClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete mainclause-like class.
*/
type MainClauseClassLike interface {
	// Constructors
	MakeWithFlow(flow FlowLike) MainClauseLike
	MakeWithAssignment(assignment AssignmentLike) MainClauseLike
	MakeWithMessaging(messaging MessagingLike) MainClauseLike
	MakeWithRepository(repository RepositoryLike) MainClauseLike
}

/*
MessageClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete message-like class.
*/
type MessageClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) MessageLike
}

/*
MessagingClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete messaging-like class.
*/
type MessagingClassLike interface {
	// Constructors
	MakeWithPostClause(postClause PostClauseLike) MessagingLike
	MakeWithRetrieveClause(retrieveClause RetrieveClauseLike) MessagingLike
	MakeWithAcceptClause(acceptClause AcceptClauseLike) MessagingLike
	MakeWithRejectClause(rejectClause RejectClauseLike) MessagingLike
	MakeWithPublishClause(publishClause PublishClauseLike) MessagingLike
}

/*
MethodClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete method-like class.
*/
type MethodClassLike interface {
	// Constructors
	MakeWithIdentifier(identifier string) MethodLike
}

/*
NotarizeClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete notarizeclause-like class.
*/
type NotarizeClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		draft DraftLike,
		citation CitationLike,
	) NotarizeClauseLike
}

/*
OnClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete onclause-like class.
*/
type OnClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		failure FailureLike,
		template TemplateLike,
		procedure ProcedureLike,
	) OnClauseLike
}

/*
ParameterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameter-like class.
*/
type ParameterClassLike interface {
	// Constructors
	MakeWithAttributes(
		symbol string,
		component ComponentLike,
	) ParameterLike
}

/*
ParametersClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parameters-like class.
*/
type ParametersClassLike interface {
	// Constructors
	MakeWithParameters(parameters col.ListLike[ParameterLike]) ParametersLike
	MakeWithAttributes(
		parameter ParameterLike,
		note string,
	) ParametersLike
}

/*
ParserClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructors
	Make() ParserLike
}

/*
PostClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete postclause-like class.
*/
type PostClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		message MessageLike,
		bag BagLike,
	) PostClauseLike
}

/*
PrecedenceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete precedence-like class.
*/
type PrecedenceClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) PrecedenceLike
}

/*
PrimitiveClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete primitive-like class.
*/
type PrimitiveClassLike interface {
	// Constructors
	MakeWithElement(element ElementLike) PrimitiveLike
	MakeWithString(string_ StringLike) PrimitiveLike
}

/*
ProcedureClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete procedure-like class.
*/
type ProcedureClassLike interface {
	// Constructors
	MakeWithLines(lines LinesLike) ProcedureLike
}

/*
PublishClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete publishclause-like class.
*/
type PublishClauseClassLike interface {
	// Constructors
	MakeWithEvent(event EventLike) PublishClauseLike
}

/*
RangeClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete range-like class.
*/
type RangeClassLike interface {
	// Constructors
	MakeWithPrimitives(primitives col.ListLike[PrimitiveLike]) RangeLike
}

/*
RecipientClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete recipient-like class.
*/
type RecipientClassLike interface {
	// Constructors
	MakeWithSymbol(symbol string) RecipientLike
	MakeWithAttribute(attribute AttributeLike) RecipientLike
}

/*
RejectClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete rejectclause-like class.
*/
type RejectClauseClassLike interface {
	// Constructors
	MakeWithMessage(message MessageLike) RejectClauseLike
}

/*
RepositoryClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete repository-like class.
*/
type RepositoryClassLike interface {
	// Constructors
	MakeWithCheckoutClause(checkoutClause CheckoutClauseLike) RepositoryLike
	MakeWithSaveClause(saveClause SaveClauseLike) RepositoryLike
	MakeWithDiscardClause(discardClause DiscardClauseLike) RepositoryLike
	MakeWithNotarizeClause(notarizeClause NotarizeClauseLike) RepositoryLike
}

/*
ResultClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete result-like class.
*/
type ResultClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) ResultLike
}

/*
RetrieveClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete retrieveclause-like class.
*/
type RetrieveClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		recipient RecipientLike,
		bag BagLike,
	) RetrieveClauseLike
}

/*
ReturnClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete returnclause-like class.
*/
type ReturnClauseClassLike interface {
	// Constructors
	MakeWithResult(result ResultLike) ReturnClauseLike
}

/*
SaveClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete saveclause-like class.
*/
type SaveClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		draft DraftLike,
		citation CitationLike,
	) SaveClauseLike
}

/*
ScannerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

MatchToken() a list of strings representing any matches found in the specified
text of the specified token type using the regular expression defined for that
token type.  If the regular expression contains submatch patterns the matching
substrings are returned as additional values in the list.
*/
type ScannerClassLike interface {
	// Constructors
	Make(
		source string,
		tokens col.QueueLike[TokenLike],
	) ScannerLike

	// Functions
	FormatToken(token TokenLike) string
	MatchToken(
		type_ TokenType,
		text string,
	) col.ListLike[string]
}

/*
SelectClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete selectclause-like class.
*/
type SelectClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		target TargetLike,
		template TemplateLike,
		procedure ProcedureLike,
	) SelectClauseLike
}

/*
SequenceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete sequence-like class.
*/
type SequenceClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) SequenceLike
}

/*
StatementClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete statement-like class.
*/
type StatementClassLike interface {
	// Constructors
	MakeWithAttributes(
		mainClause MainClauseLike,
		onClause OnClauseLike,
	) StatementLike
}

/*
StringClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete string-like class.
*/
type StringClassLike interface {
	// Constructors
	MakeWithBinary(binary string) StringLike
	MakeWithBytecode(bytecode string) StringLike
	MakeWithName(name string) StringLike
	MakeWithNarrative(narrative string) StringLike
	MakeWithQuote(quote string) StringLike
	MakeWithSymbol(symbol string) StringLike
	MakeWithTag(tag string) StringLike
	MakeWithVersion(version string) StringLike
}

/*
SubcomponentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete subcomponent-like class.
*/
type SubcomponentClassLike interface {
	// Constructors
	MakeWithAttributes(
		composite CompositeLike,
		indices IndicesLike,
	) SubcomponentLike
}

/*
TargetClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete target-like class.
*/
type TargetClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) TargetLike
}

/*
TemplateClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete template-like class.
*/
type TemplateClassLike interface {
	// Constructors
	MakeWithExpression(expression ExpressionLike) TemplateLike
}

/*
ThrowClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete throwclause-like class.
*/
type ThrowClauseClassLike interface {
	// Constructors
	MakeWithException(exception ExceptionLike) ThrowClauseLike
}

/*
TokenClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete token-like class.
*/
type TokenClassLike interface {
	// Constructors
	MakeWithAttributes(
		line int,
		position int,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructors
	Make() ValidatorLike
}

/*
ValueClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete value-like class.
*/
type ValueClassLike interface {
	// Constructors
	MakeWithComponent(component ComponentLike) ValueLike
}

/*
ValuesClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete values-like class.
*/
type ValuesClassLike interface {
	// Constructors
	MakeWithValues(values col.ListLike[ValueLike]) ValuesLike
	MakeWithAttributes(
		value ValueLike,
		note string,
	) ValuesLike
}

/*
VariableClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete variable-like class.
*/
type VariableClassLike interface {
	// Constructors
	MakeWithIdentifier(identifier string) VariableLike
}

/*
WhileClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete whileclause-like class.
*/
type WhileClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		condition ConditionLike,
		procedure ProcedureLike,
	) WhileClauseLike
}

/*
WithClauseClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
concrete withclause-like class.
*/
type WithClauseClassLike interface {
	// Constructors
	MakeWithAttributes(
		item ItemLike,
		sequence SequenceLike,
		procedure ProcedureLike,
	) WithClauseLike
}

// Instances

/*
AcceptClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete acceptclause-like class.
*/
type AcceptClauseLike interface {
	// Attributes
	GetMessage() MessageLike
}

/*
AnnotationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete annotation-like class.
*/
type AnnotationLike interface {
	// Attributes
	GetNote() string
	GetComment() string
}

/*
ArgumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete argument-like class.
*/
type ArgumentLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
ArgumentsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete arguments-like class.
*/
type ArgumentsLike interface {
	// Attributes
	GetArguments() col.ListLike[ArgumentLike]
}

/*
ArithmeticLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete arithmetic-like class.
*/
type ArithmeticLike interface {
	// Attributes
	GetExpressions() col.ListLike[ExpressionLike]
}

/*
AssignmentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete assignment-like class.
*/
type AssignmentLike interface {
	// Attributes
	GetLetClause() LetClauseLike
}

/*
AssociationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete association-like class.
*/
type AssociationLike interface {
	// Attributes
	GetKey() KeyLike
	GetValue() ValueLike
}

/*
AssociationsLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete associations-like class.
*/
type AssociationsLike interface {
	// Attributes
	GetAssociations() col.ListLike[AssociationLike]
	GetNote() string
}

/*
AttributeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete attribute-like class.
*/
type AttributeLike interface {
	// Attributes
	GetVariable() VariableLike
	GetIndices() IndicesLike
}

/*
BagLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete bag-like class.
*/
type BagLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
BreakClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete breakclause-like class.
*/
type BreakClauseLike interface {
}

/*
ChainingLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete chaining-like class.
*/
type ChainingLike interface {
	// Attributes
	GetExpressions() col.ListLike[ExpressionLike]
}

/*
CheckoutClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete checkoutclause-like class.
*/
type CheckoutClauseLike interface {
	// Attributes
	GetRecipient() RecipientLike
	GetLevel() LevelLike
	GetCitation() CitationLike
}

/*
CitationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete citation-like class.
*/
type CitationLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
CollectionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete collection-like class.
*/
type CollectionLike interface {
	// Attributes
	GetAssociations() AssociationsLike
	GetValues() ValuesLike
}

/*
ComparisonLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete comparison-like class.
*/
type ComparisonLike interface {
	// Attributes
	GetExpressions() col.ListLike[ExpressionLike]
}

/*
ComplementLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete complement-like class.
*/
type ComplementLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
ComponentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete component-like class.
*/
type ComponentLike interface {
	// Attributes
	GetEntity() EntityLike
	GetContext() ContextLike
}

/*
CompositeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete composite-like class.
*/
type CompositeLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
ConditionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete condition-like class.
*/
type ConditionLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
ContextLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete context-like class.
*/
type ContextLike interface {
	// Attributes
	GetParameters() ParametersLike
}

/*
ContinueClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete continueclause-like class.
*/
type ContinueClauseLike interface {
}

/*
DereferenceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete dereference-like class.
*/
type DereferenceLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
DiscardClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete discardclause-like class.
*/
type DiscardClauseLike interface {
	// Attributes
	GetDraft() DraftLike
}

/*
DocumentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete document-like class.
*/
type DocumentLike interface {
	// Attributes
	GetHeader() HeaderLike
	GetComponent() ComponentLike
}

/*
DraftLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete draft-like class.
*/
type DraftLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
ElementLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete element-like class.
*/
type ElementLike interface {
	// Attributes
	GetAngle() string
	GetBoolean() string
	GetDuration() string
	GetMoment() string
	GetNumber() string
	GetPattern() string
	GetPercentage() string
	GetProbability() string
	GetResource() string
}

/*
EntityLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete entity-like class.
*/
type EntityLike interface {
	// Attributes
	GetElement() ElementLike
	GetString() StringLike
	GetRange() RangeLike
	GetCollection() CollectionLike
	GetProcedure() ProcedureLike
}

/*
EventLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete event-like class.
*/
type EventLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
ExceptionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete exception-like class.
*/
type ExceptionLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
ExponentialLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete exponential-like class.
*/
type ExponentialLike interface {
	// Attributes
	GetExpressions() col.ListLike[ExpressionLike]
}

/*
ExpressionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete expression-like class.
*/
type ExpressionLike interface {
	// Attributes
	GetComponent() ComponentLike
	GetIntrinsic() IntrinsicLike
	GetVariable() VariableLike
	GetPrecedence() PrecedenceLike
	GetDereference() DereferenceLike
	GetInvocation() InvocationLike
	GetSubcomponent() SubcomponentLike
	GetChaining() ChainingLike
	GetExponential() ExponentialLike
	GetInversion() InversionLike
	GetArithmetic() ArithmeticLike
	GetMagnitude() MagnitudeLike
	GetComparison() ComparisonLike
	GetComplement() ComplementLike
	GetLogical() LogicalLike
}

/*
FailureLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete failure-like class.
*/
type FailureLike interface {
	// Attributes
	GetSymbol() string
}

/*
FlowLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete flow-like class.
*/
type FlowLike interface {
	// Attributes
	GetIfClause() IfClauseLike
	GetSelectClause() SelectClauseLike
	GetWhileClause() WhileClauseLike
	GetWithClause() WithClauseLike
	GetContinueClause() ContinueClauseLike
	GetBreakClause() BreakClauseLike
	GetReturnClause() ReturnClauseLike
	GetThrowClause() ThrowClauseLike
}

/*
FormatterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Methods
	FormatDocument(document DocumentLike) string
}

/*
FunctionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete function-like class.
*/
type FunctionLike interface {
	// Attributes
	GetIdentifier() string
}

/*
HeaderLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete header-like class.
*/
type HeaderLike interface {
	// Attributes
	GetComment() string
}

/*
IfClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete ifclause-like class.
*/
type IfClauseLike interface {
	// Attributes
	GetCondition() ConditionLike
	GetProcedure() ProcedureLike
}

/*
IndexLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete index-like class.
*/
type IndexLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
IndicesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete indices-like class.
*/
type IndicesLike interface {
	// Attributes
	GetIndexs() col.ListLike[IndexLike]
}

/*
IntrinsicLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete intrinsic-like class.
*/
type IntrinsicLike interface {
	// Attributes
	GetFunction() FunctionLike
	GetArguments() ArgumentsLike
}

/*
InversionLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete inversion-like class.
*/
type InversionLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
InvocationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete invocation-like class.
*/
type InvocationLike interface {
	// Attributes
	GetTarget() TargetLike
	GetMethod() MethodLike
	GetArguments() ArgumentsLike
}

/*
ItemLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete item-like class.
*/
type ItemLike interface {
	// Attributes
	GetSymbol() string
}

/*
KeyLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete key-like class.
*/
type KeyLike interface {
	// Attributes
	GetPrimitive() PrimitiveLike
}

/*
LetClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete letclause-like class.
*/
type LetClauseLike interface {
	// Attributes
	GetRecipient() RecipientLike
	GetExpression() ExpressionLike
}

/*
LevelLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete level-like class.
*/
type LevelLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
LineLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete line-like class.
*/
type LineLike interface {
	// Attributes
	GetAnnotation() AnnotationLike
	GetStatement() StatementLike
}

/*
LinesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete lines-like class.
*/
type LinesLike interface {
	// Attributes
	GetLines() col.ListLike[LineLike]
	GetNote() string
}

/*
LogicalLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete logical-like class.
*/
type LogicalLike interface {
	// Attributes
	GetExpressions() col.ListLike[ExpressionLike]
}

/*
MagnitudeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete magnitude-like class.
*/
type MagnitudeLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
MainClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete mainclause-like class.
*/
type MainClauseLike interface {
	// Attributes
	GetFlow() FlowLike
	GetAssignment() AssignmentLike
	GetMessaging() MessagingLike
	GetRepository() RepositoryLike
}

/*
MessageLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete message-like class.
*/
type MessageLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
MessagingLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete messaging-like class.
*/
type MessagingLike interface {
	// Attributes
	GetPostClause() PostClauseLike
	GetRetrieveClause() RetrieveClauseLike
	GetAcceptClause() AcceptClauseLike
	GetRejectClause() RejectClauseLike
	GetPublishClause() PublishClauseLike
}

/*
MethodLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete method-like class.
*/
type MethodLike interface {
	// Attributes
	GetIdentifier() string
}

/*
NotarizeClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete notarizeclause-like class.
*/
type NotarizeClauseLike interface {
	// Attributes
	GetDraft() DraftLike
	GetCitation() CitationLike
}

/*
OnClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete onclause-like class.
*/
type OnClauseLike interface {
	// Attributes
	GetFailure() FailureLike
	GetTemplate() TemplateLike
	GetProcedure() ProcedureLike
}

/*
ParameterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameter-like class.
*/
type ParameterLike interface {
	// Attributes
	GetSymbol() string
	GetComponent() ComponentLike
}

/*
ParametersLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parameters-like class.
*/
type ParametersLike interface {
	// Attributes
	GetParameters() col.ListLike[ParameterLike]
	GetNote() string
}

/*
ParserLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Methods
	ParseSource(source string) DocumentLike
}

/*
PostClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete postclause-like class.
*/
type PostClauseLike interface {
	// Attributes
	GetMessage() MessageLike
	GetBag() BagLike
}

/*
PrecedenceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete precedence-like class.
*/
type PrecedenceLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
PrimitiveLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete primitive-like class.
*/
type PrimitiveLike interface {
	// Attributes
	GetElement() ElementLike
	GetString() StringLike
}

/*
ProcedureLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete procedure-like class.
*/
type ProcedureLike interface {
	// Attributes
	GetLines() LinesLike
}

/*
PublishClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete publishclause-like class.
*/
type PublishClauseLike interface {
	// Attributes
	GetEvent() EventLike
}

/*
RangeLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete range-like class.
*/
type RangeLike interface {
	// Attributes
	GetPrimitives() col.ListLike[PrimitiveLike]
}

/*
RecipientLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete recipient-like class.
*/
type RecipientLike interface {
	// Attributes
	GetSymbol() string
	GetAttribute() AttributeLike
}

/*
RejectClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete rejectclause-like class.
*/
type RejectClauseLike interface {
	// Attributes
	GetMessage() MessageLike
}

/*
RepositoryLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete repository-like class.
*/
type RepositoryLike interface {
	// Attributes
	GetCheckoutClause() CheckoutClauseLike
	GetSaveClause() SaveClauseLike
	GetDiscardClause() DiscardClauseLike
	GetNotarizeClause() NotarizeClauseLike
}

/*
ResultLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete result-like class.
*/
type ResultLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
RetrieveClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete retrieveclause-like class.
*/
type RetrieveClauseLike interface {
	// Attributes
	GetRecipient() RecipientLike
	GetBag() BagLike
}

/*
ReturnClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete returnclause-like class.
*/
type ReturnClauseLike interface {
	// Attributes
	GetResult() ResultLike
}

/*
SaveClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete saveclause-like class.
*/
type SaveClauseLike interface {
	// Attributes
	GetDraft() DraftLike
	GetCitation() CitationLike
}

/*
ScannerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
}

/*
SelectClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete selectclause-like class.
*/
type SelectClauseLike interface {
	// Attributes
	GetTarget() TargetLike
	GetTemplate() TemplateLike
	GetProcedure() ProcedureLike
}

/*
SequenceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete sequence-like class.
*/
type SequenceLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
StatementLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete statement-like class.
*/
type StatementLike interface {
	// Attributes
	GetMainClause() MainClauseLike
	GetOnClause() OnClauseLike
}

/*
StringLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete string-like class.
*/
type StringLike interface {
	// Attributes
	GetBinary() string
	GetBytecode() string
	GetName() string
	GetNarrative() string
	GetQuote() string
	GetSymbol() string
	GetTag() string
	GetVersion() string
}

/*
SubcomponentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete subcomponent-like class.
*/
type SubcomponentLike interface {
	// Attributes
	GetComposite() CompositeLike
	GetIndices() IndicesLike
}

/*
TargetLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete target-like class.
*/
type TargetLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
TemplateLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete template-like class.
*/
type TemplateLike interface {
	// Attributes
	GetExpression() ExpressionLike
}

/*
ThrowClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete throwclause-like class.
*/
type ThrowClauseLike interface {
	// Attributes
	GetException() ExceptionLike
}

/*
TokenLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Attributes
	GetLine() int
	GetPosition() int
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Methods
	ValidateDocument(document DocumentLike)
}

/*
ValueLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete value-like class.
*/
type ValueLike interface {
	// Attributes
	GetComponent() ComponentLike
}

/*
ValuesLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete values-like class.
*/
type ValuesLike interface {
	// Attributes
	GetValues() col.ListLike[ValueLike]
	GetNote() string
}

/*
VariableLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete variable-like class.
*/
type VariableLike interface {
	// Attributes
	GetIdentifier() string
}

/*
WhileClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete whileclause-like class.
*/
type WhileClauseLike interface {
	// Attributes
	GetCondition() ConditionLike
	GetProcedure() ProcedureLike
}

/*
WithClauseLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a concrete withclause-like class.
*/
type WithClauseLike interface {
	// Attributes
	GetItem() ItemLike
	GetSequence() SequenceLike
	GetProcedure() ProcedureLike
}
