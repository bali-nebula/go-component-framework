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

Package "grammar" provides the following grammar classes that operate on the
abstract syntax tree (AST) for this module:
  - Token captures the attributes associated with a parsed token.
  - Scanner is used to scan the source byte stream and recognize matching tokens.
  - Parser is used to process the token stream and generate the AST.
  - Validator is used to validate the semantics associated with an AST.
  - Formatter is used to format an AST back into a canonical version of its source.
  - Visitor walks the AST and calls processor methods for each node in the tree.
  - Processor provides empty processor methods to be inherited by the processors.

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
package grammar

import (
	ast "github.com/bali-nebula/go-component-framework/v3/ast"
	col "github.com/craterdog/go-collection-framework/v5/collection"
)

// TYPE DECLARATIONS

/*
TokenType is a constrained type representing any token type recognized by a
scanner.
*/
type TokenType uint8

const (
	ErrorToken TokenType = iota
	AmpersandToken
	AndToken
	AngleToken
	ArrowToken
	BarToken
	BinaryToken
	BooleanToken
	BytecodeToken
	CaretToken
	ColonToken
	ColonEqualToken
	CommentToken
	DashToken
	DashEqualToken
	DefaultEqualToken
	DelimiterToken
	DotToken
	DoubleToken
	DurationToken
	EqualToken
	IdentifierToken
	IorToken
	IsToken
	LessToken
	MatchesToken
	MomentToken
	MoreToken
	NameToken
	NarrativeToken
	NewlineToken
	NotToken
	NoteToken
	NumberToken
	PatternToken
	PercentageToken
	PlusToken
	PlusEqualToken
	ProbabilityToken
	QuoteToken
	ResourceToken
	SanToken
	SlashToken
	SlashEqualToken
	SnailToken
	SpaceToken
	StarToken
	StarEqualToken
	SymbolToken
	TagToken
	VersionToken
	XorToken
	XyzToken
)

// FUNCTIONAL DECLARATIONS

// CLASS DECLARATIONS

/*
FormatterClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete formatter-like class.
*/
type FormatterClassLike interface {
	// Constructor Methods
	Formatter() FormatterLike
}

/*
ParserClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete parser-like class.
*/
type ParserClassLike interface {
	// Constructor Methods
	Parser() ParserLike
}

/*
ProcessorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete processor-like class.
*/
type ProcessorClassLike interface {
	// Constructor Methods
	Processor() ProcessorLike
}

/*
ScannerClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete scanner-like class.  The following functions are supported:

FormatToken() returns a formatted string containing the attributes of the token.

FormatType() returns the string version of the token type.

MatchesType() determines whether or not a token value is of a specified type.
*/
type ScannerClassLike interface {
	// Constructor Methods
	Scanner(
		source string,
		tokens col.QueueLike[TokenLike],
	) ScannerLike

	// Function Methods
	FormatToken(
		token TokenLike,
	) string
	FormatType(
		tokenType TokenType,
	) string
	MatchesType(
		tokenValue string,
		tokenType TokenType,
	) bool
}

/*
TokenClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type TokenClassLike interface {
	// Constructor Methods
	Token(
		line uint,
		position uint,
		type_ TokenType,
		value string,
	) TokenLike
}

/*
ValidatorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete validator-like class.
*/
type ValidatorClassLike interface {
	// Constructor Methods
	Validator() ValidatorLike
}

/*
VisitorClassLike is a class interface that declares the complete set of
class constructors, constants and functions that must be supported by each
concrete visitor-like class.
*/
type VisitorClassLike interface {
	// Constructor Methods
	Visitor(
		processor Methodical,
	) VisitorLike
}

// INSTANCE DECLARATIONS

/*
FormatterLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete formatter-like class.
*/
type FormatterLike interface {
	// Principal Methods
	GetClass() FormatterClassLike
	FormatDocument(
		document ast.DocumentLike,
	) string

	// Aspect Interfaces
	Methodical
}

/*
ParserLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete parser-like class.
*/
type ParserLike interface {
	// Principal Methods
	GetClass() ParserClassLike
	ParseSource(
		source string,
	) ast.DocumentLike
}

/*
ProcessorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete processor-like class.
*/
type ProcessorLike interface {
	// Principal Methods
	GetClass() ProcessorClassLike

	// Aspect Interfaces
	Methodical
}

/*
ScannerLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete scanner-like class.
*/
type ScannerLike interface {
	// Principal Methods
	GetClass() ScannerClassLike
}

/*
TokenLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete token-like class.
*/
type TokenLike interface {
	// Principal Methods
	GetClass() TokenClassLike

	// Attribute Methods
	GetLine() uint
	GetPosition() uint
	GetType() TokenType
	GetValue() string
}

/*
ValidatorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete validator-like class.
*/
type ValidatorLike interface {
	// Principal Methods
	GetClass() ValidatorClassLike
	ValidateDocument(
		document ast.DocumentLike,
	)

	// Aspect Interfaces
	Methodical
}

/*
VisitorLike is an instance interface that declares the complete set of
principal, attribute and aspect methods that must be supported by each
instance of a concrete visitor-like class.
*/
type VisitorLike interface {
	// Principal Methods
	GetClass() VisitorClassLike
	VisitDocument(
		document ast.DocumentLike,
	)
}

// ASPECT DECLARATIONS

/*
Methodical declares the set of method signatures that must be supported by
all methodical processors.
*/
type Methodical interface {
	ProcessAmpersand(
		ampersand string,
	)
	ProcessAnd(
		and string,
	)
	ProcessAngle(
		angle string,
	)
	ProcessArrow(
		arrow string,
	)
	ProcessBar(
		bar string,
	)
	ProcessBinary(
		binary string,
	)
	ProcessBoolean(
		boolean string,
	)
	ProcessBytecode(
		bytecode string,
	)
	ProcessCaret(
		caret string,
	)
	ProcessColon(
		colon string,
	)
	ProcessColonEqual(
		colonEqual string,
	)
	ProcessComment(
		comment string,
	)
	ProcessDash(
		dash string,
	)
	ProcessDashEqual(
		dashEqual string,
	)
	ProcessDefaultEqual(
		defaultEqual string,
	)
	ProcessDot(
		dot string,
	)
	ProcessDouble(
		double string,
	)
	ProcessDuration(
		duration string,
	)
	ProcessEqual(
		equal string,
	)
	ProcessIdentifier(
		identifier string,
	)
	ProcessIor(
		ior string,
	)
	ProcessIs(
		is string,
	)
	ProcessLess(
		less string,
	)
	ProcessMatches(
		matches string,
	)
	ProcessMoment(
		moment string,
	)
	ProcessMore(
		more string,
	)
	ProcessName(
		name string,
	)
	ProcessNarrative(
		narrative string,
	)
	ProcessNewline(
		newline string,
	)
	ProcessNot(
		not string,
	)
	ProcessNote(
		note string,
	)
	ProcessNumber(
		number string,
	)
	ProcessPattern(
		pattern string,
	)
	ProcessPercentage(
		percentage string,
	)
	ProcessPlus(
		plus string,
	)
	ProcessPlusEqual(
		plusEqual string,
	)
	ProcessProbability(
		probability string,
	)
	ProcessQuote(
		quote string,
	)
	ProcessResource(
		resource string,
	)
	ProcessSan(
		san string,
	)
	ProcessSlash(
		slash string,
	)
	ProcessSlashEqual(
		slashEqual string,
	)
	ProcessSnail(
		snail string,
	)
	ProcessSpace(
		space string,
	)
	ProcessStar(
		star string,
	)
	ProcessStarEqual(
		starEqual string,
	)
	ProcessSymbol(
		symbol string,
	)
	ProcessTag(
		tag string,
	)
	ProcessVersion(
		version string,
	)
	ProcessXor(
		xor string,
	)
	ProcessXyz(
		xyz string,
	)
	PreprocessAcceptClause(
		acceptClause ast.AcceptClauseLike,
	)
	ProcessAcceptClauseSlot(
		slot uint,
	)
	PostprocessAcceptClause(
		acceptClause ast.AcceptClauseLike,
	)
	PreprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index uint,
		size uint,
	)
	ProcessAdditionalArgumentSlot(
		slot uint,
	)
	PostprocessAdditionalArgument(
		additionalArgument ast.AdditionalArgumentLike,
		index uint,
		size uint,
	)
	PreprocessAdditionalAssociation(
		additionalAssociation ast.AdditionalAssociationLike,
		index uint,
		size uint,
	)
	ProcessAdditionalAssociationSlot(
		slot uint,
	)
	PostprocessAdditionalAssociation(
		additionalAssociation ast.AdditionalAssociationLike,
		index uint,
		size uint,
	)
	PreprocessAdditionalIndex(
		additionalIndex ast.AdditionalIndexLike,
		index uint,
		size uint,
	)
	ProcessAdditionalIndexSlot(
		slot uint,
	)
	PostprocessAdditionalIndex(
		additionalIndex ast.AdditionalIndexLike,
		index uint,
		size uint,
	)
	PreprocessAdditionalStatement(
		additionalStatement ast.AdditionalStatementLike,
		index uint,
		size uint,
	)
	ProcessAdditionalStatementSlot(
		slot uint,
	)
	PostprocessAdditionalStatement(
		additionalStatement ast.AdditionalStatementLike,
		index uint,
		size uint,
	)
	PreprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index uint,
		size uint,
	)
	ProcessAdditionalValueSlot(
		slot uint,
	)
	PostprocessAdditionalValue(
		additionalValue ast.AdditionalValueLike,
		index uint,
		size uint,
	)
	PreprocessAnnotatedAssociation(
		annotatedAssociation ast.AnnotatedAssociationLike,
		index uint,
		size uint,
	)
	ProcessAnnotatedAssociationSlot(
		slot uint,
	)
	PostprocessAnnotatedAssociation(
		annotatedAssociation ast.AnnotatedAssociationLike,
		index uint,
		size uint,
	)
	PreprocessAnnotatedStatement(
		annotatedStatement ast.AnnotatedStatementLike,
		index uint,
		size uint,
	)
	ProcessAnnotatedStatementSlot(
		slot uint,
	)
	PostprocessAnnotatedStatement(
		annotatedStatement ast.AnnotatedStatementLike,
		index uint,
		size uint,
	)
	PreprocessAnnotatedValue(
		annotatedValue ast.AnnotatedValueLike,
		index uint,
		size uint,
	)
	ProcessAnnotatedValueSlot(
		slot uint,
	)
	PostprocessAnnotatedValue(
		annotatedValue ast.AnnotatedValueLike,
		index uint,
		size uint,
	)
	PreprocessArgument(
		argument ast.ArgumentLike,
	)
	ProcessArgumentSlot(
		slot uint,
	)
	PostprocessArgument(
		argument ast.ArgumentLike,
	)
	PreprocessArguments(
		arguments ast.ArgumentsLike,
	)
	ProcessArgumentsSlot(
		slot uint,
	)
	PostprocessArguments(
		arguments ast.ArgumentsLike,
	)
	PreprocessAssign(
		assign ast.AssignLike,
	)
	ProcessAssignSlot(
		slot uint,
	)
	PostprocessAssign(
		assign ast.AssignLike,
	)
	PreprocessAssociation(
		association ast.AssociationLike,
	)
	ProcessAssociationSlot(
		slot uint,
	)
	PostprocessAssociation(
		association ast.AssociationLike,
	)
	PreprocessAssociations(
		associations ast.AssociationsLike,
	)
	ProcessAssociationsSlot(
		slot uint,
	)
	PostprocessAssociations(
		associations ast.AssociationsLike,
	)
	PreprocessAtLevel(
		atLevel ast.AtLevelLike,
	)
	ProcessAtLevelSlot(
		slot uint,
	)
	PostprocessAtLevel(
		atLevel ast.AtLevelLike,
	)
	PreprocessAttribute(
		attribute ast.AttributeLike,
	)
	ProcessAttributeSlot(
		slot uint,
	)
	PostprocessAttribute(
		attribute ast.AttributeLike,
	)
	PreprocessBag(
		bag ast.BagLike,
	)
	ProcessBagSlot(
		slot uint,
	)
	PostprocessBag(
		bag ast.BagLike,
	)
	PreprocessBreakClause(
		breakClause ast.BreakClauseLike,
	)
	ProcessBreakClauseSlot(
		slot uint,
	)
	PostprocessBreakClause(
		breakClause ast.BreakClauseLike,
	)
	PreprocessCheckoutClause(
		checkoutClause ast.CheckoutClauseLike,
	)
	ProcessCheckoutClauseSlot(
		slot uint,
	)
	PostprocessCheckoutClause(
		checkoutClause ast.CheckoutClauseLike,
	)
	PreprocessCitation(
		citation ast.CitationLike,
	)
	ProcessCitationSlot(
		slot uint,
	)
	PostprocessCitation(
		citation ast.CitationLike,
	)
	PreprocessCollection(
		collection ast.CollectionLike,
	)
	ProcessCollectionSlot(
		slot uint,
	)
	PostprocessCollection(
		collection ast.CollectionLike,
	)
	PreprocessComplement(
		complement ast.ComplementLike,
	)
	ProcessComplementSlot(
		slot uint,
	)
	PostprocessComplement(
		complement ast.ComplementLike,
	)
	PreprocessComponent(
		component ast.ComponentLike,
	)
	ProcessComponentSlot(
		slot uint,
	)
	PostprocessComponent(
		component ast.ComponentLike,
	)
	PreprocessCondition(
		condition ast.ConditionLike,
	)
	ProcessConditionSlot(
		slot uint,
	)
	PostprocessCondition(
		condition ast.ConditionLike,
	)
	PreprocessContinueClause(
		continueClause ast.ContinueClauseLike,
	)
	ProcessContinueClauseSlot(
		slot uint,
	)
	PostprocessContinueClause(
		continueClause ast.ContinueClauseLike,
	)
	PreprocessDiscardClause(
		discardClause ast.DiscardClauseLike,
	)
	ProcessDiscardClauseSlot(
		slot uint,
	)
	PostprocessDiscardClause(
		discardClause ast.DiscardClauseLike,
	)
	PreprocessDoClause(
		doClause ast.DoClauseLike,
	)
	ProcessDoClauseSlot(
		slot uint,
	)
	PostprocessDoClause(
		doClause ast.DoClauseLike,
	)
	PreprocessDocument(
		document ast.DocumentLike,
	)
	ProcessDocumentSlot(
		slot uint,
	)
	PostprocessDocument(
		document ast.DocumentLike,
	)
	PreprocessDraft(
		draft ast.DraftLike,
	)
	ProcessDraftSlot(
		slot uint,
	)
	PostprocessDraft(
		draft ast.DraftLike,
	)
	PreprocessElement(
		element ast.ElementLike,
	)
	ProcessElementSlot(
		slot uint,
	)
	PostprocessElement(
		element ast.ElementLike,
	)
	PreprocessEntity(
		entity ast.EntityLike,
	)
	ProcessEntitySlot(
		slot uint,
	)
	PostprocessEntity(
		entity ast.EntityLike,
	)
	PreprocessEvent(
		event ast.EventLike,
	)
	ProcessEventSlot(
		slot uint,
	)
	PostprocessEvent(
		event ast.EventLike,
	)
	PreprocessException(
		exception ast.ExceptionLike,
	)
	ProcessExceptionSlot(
		slot uint,
	)
	PostprocessException(
		exception ast.ExceptionLike,
	)
	PreprocessExpression(
		expression ast.ExpressionLike,
	)
	ProcessExpressionSlot(
		slot uint,
	)
	PostprocessExpression(
		expression ast.ExpressionLike,
	)
	PreprocessFailure(
		failure ast.FailureLike,
	)
	ProcessFailureSlot(
		slot uint,
	)
	PostprocessFailure(
		failure ast.FailureLike,
	)
	PreprocessFlow(
		flow ast.FlowLike,
	)
	ProcessFlowSlot(
		slot uint,
	)
	PostprocessFlow(
		flow ast.FlowLike,
	)
	PreprocessFunction(
		function ast.FunctionLike,
	)
	ProcessFunctionSlot(
		slot uint,
	)
	PostprocessFunction(
		function ast.FunctionLike,
	)
	PreprocessIfClause(
		ifClause ast.IfClauseLike,
	)
	ProcessIfClauseSlot(
		slot uint,
	)
	PostprocessIfClause(
		ifClause ast.IfClauseLike,
	)
	PreprocessInclusion(
		inclusion ast.InclusionLike,
	)
	ProcessInclusionSlot(
		slot uint,
	)
	PostprocessInclusion(
		inclusion ast.InclusionLike,
	)
	PreprocessIndex(
		index ast.IndexLike,
	)
	ProcessIndexSlot(
		slot uint,
	)
	PostprocessIndex(
		index ast.IndexLike,
	)
	PreprocessIndices(
		indices ast.IndicesLike,
	)
	ProcessIndicesSlot(
		slot uint,
	)
	PostprocessIndices(
		indices ast.IndicesLike,
	)
	PreprocessIndirect(
		indirect ast.IndirectLike,
	)
	ProcessIndirectSlot(
		slot uint,
	)
	PostprocessIndirect(
		indirect ast.IndirectLike,
	)
	PreprocessInduction(
		induction ast.InductionLike,
	)
	ProcessInductionSlot(
		slot uint,
	)
	PostprocessInduction(
		induction ast.InductionLike,
	)
	PreprocessInlineAssociations(
		inlineAssociations ast.InlineAssociationsLike,
	)
	ProcessInlineAssociationsSlot(
		slot uint,
	)
	PostprocessInlineAssociations(
		inlineAssociations ast.InlineAssociationsLike,
	)
	PreprocessInlineStatements(
		inlineStatements ast.InlineStatementsLike,
	)
	ProcessInlineStatementsSlot(
		slot uint,
	)
	PostprocessInlineStatements(
		inlineStatements ast.InlineStatementsLike,
	)
	PreprocessInlineValues(
		inlineValues ast.InlineValuesLike,
	)
	ProcessInlineValuesSlot(
		slot uint,
	)
	PostprocessInlineValues(
		inlineValues ast.InlineValuesLike,
	)
	PreprocessInverse(
		inverse ast.InverseLike,
	)
	ProcessInverseSlot(
		slot uint,
	)
	PostprocessInverse(
		inverse ast.InverseLike,
	)
	PreprocessInversion(
		inversion ast.InversionLike,
	)
	ProcessInversionSlot(
		slot uint,
	)
	PostprocessInversion(
		inversion ast.InversionLike,
	)
	PreprocessInvocation(
		invocation ast.InvocationLike,
	)
	ProcessInvocationSlot(
		slot uint,
	)
	PostprocessInvocation(
		invocation ast.InvocationLike,
	)
	PreprocessItem(
		item ast.ItemLike,
	)
	ProcessItemSlot(
		slot uint,
	)
	PostprocessItem(
		item ast.ItemLike,
	)
	PreprocessItems(
		items ast.ItemsLike,
	)
	ProcessItemsSlot(
		slot uint,
	)
	PostprocessItems(
		items ast.ItemsLike,
	)
	PreprocessLeft(
		left ast.LeftLike,
	)
	ProcessLeftSlot(
		slot uint,
	)
	PostprocessLeft(
		left ast.LeftLike,
	)
	PreprocessLetClause(
		letClause ast.LetClauseLike,
	)
	ProcessLetClauseSlot(
		slot uint,
	)
	PostprocessLetClause(
		letClause ast.LetClauseLike,
	)
	PreprocessLogical(
		logical ast.LogicalLike,
	)
	ProcessLogicalSlot(
		slot uint,
	)
	PostprocessLogical(
		logical ast.LogicalLike,
	)
	PreprocessMagnitude(
		magnitude ast.MagnitudeLike,
	)
	ProcessMagnitudeSlot(
		slot uint,
	)
	PostprocessMagnitude(
		magnitude ast.MagnitudeLike,
	)
	PreprocessMainClause(
		mainClause ast.MainClauseLike,
	)
	ProcessMainClauseSlot(
		slot uint,
	)
	PostprocessMainClause(
		mainClause ast.MainClauseLike,
	)
	PreprocessMatchHandler(
		matchHandler ast.MatchHandlerLike,
		index uint,
		size uint,
	)
	ProcessMatchHandlerSlot(
		slot uint,
	)
	PostprocessMatchHandler(
		matchHandler ast.MatchHandlerLike,
		index uint,
		size uint,
	)
	PreprocessMessage(
		message ast.MessageLike,
	)
	ProcessMessageSlot(
		slot uint,
	)
	PostprocessMessage(
		message ast.MessageLike,
	)
	PreprocessMessaging(
		messaging ast.MessagingLike,
	)
	ProcessMessagingSlot(
		slot uint,
	)
	PostprocessMessaging(
		messaging ast.MessagingLike,
	)
	PreprocessMethod(
		method ast.MethodLike,
	)
	ProcessMethodSlot(
		slot uint,
	)
	PostprocessMethod(
		method ast.MethodLike,
	)
	PreprocessMultilineAssociations(
		multilineAssociations ast.MultilineAssociationsLike,
	)
	ProcessMultilineAssociationsSlot(
		slot uint,
	)
	PostprocessMultilineAssociations(
		multilineAssociations ast.MultilineAssociationsLike,
	)
	PreprocessMultilineStatements(
		multilineStatements ast.MultilineStatementsLike,
	)
	ProcessMultilineStatementsSlot(
		slot uint,
	)
	PostprocessMultilineStatements(
		multilineStatements ast.MultilineStatementsLike,
	)
	PreprocessMultilineValues(
		multilineValues ast.MultilineValuesLike,
	)
	ProcessMultilineValuesSlot(
		slot uint,
	)
	PostprocessMultilineValues(
		multilineValues ast.MultilineValuesLike,
	)
	PreprocessNoAssociations(
		noAssociations ast.NoAssociationsLike,
	)
	ProcessNoAssociationsSlot(
		slot uint,
	)
	PostprocessNoAssociations(
		noAssociations ast.NoAssociationsLike,
	)
	PreprocessNotarizeClause(
		notarizeClause ast.NotarizeClauseLike,
	)
	ProcessNotarizeClauseSlot(
		slot uint,
	)
	PostprocessNotarizeClause(
		notarizeClause ast.NotarizeClauseLike,
	)
	PreprocessNumerical(
		numerical ast.NumericalLike,
	)
	ProcessNumericalSlot(
		slot uint,
	)
	PostprocessNumerical(
		numerical ast.NumericalLike,
	)
	PreprocessOnClause(
		onClause ast.OnClauseLike,
	)
	ProcessOnClauseSlot(
		slot uint,
	)
	PostprocessOnClause(
		onClause ast.OnClauseLike,
	)
	PreprocessOperator(
		operator ast.OperatorLike,
	)
	ProcessOperatorSlot(
		slot uint,
	)
	PostprocessOperator(
		operator ast.OperatorLike,
	)
	PreprocessParameters(
		parameters ast.ParametersLike,
	)
	ProcessParametersSlot(
		slot uint,
	)
	PostprocessParameters(
		parameters ast.ParametersLike,
	)
	PreprocessPostClause(
		postClause ast.PostClauseLike,
	)
	ProcessPostClauseSlot(
		slot uint,
	)
	PostprocessPostClause(
		postClause ast.PostClauseLike,
	)
	PreprocessPrecedence(
		precedence ast.PrecedenceLike,
	)
	ProcessPrecedenceSlot(
		slot uint,
	)
	PostprocessPrecedence(
		precedence ast.PrecedenceLike,
	)
	PreprocessPredicate(
		predicate ast.PredicateLike,
		index uint,
		size uint,
	)
	ProcessPredicateSlot(
		slot uint,
	)
	PostprocessPredicate(
		predicate ast.PredicateLike,
		index uint,
		size uint,
	)
	PreprocessPrimitive(
		primitive ast.PrimitiveLike,
	)
	ProcessPrimitiveSlot(
		slot uint,
	)
	PostprocessPrimitive(
		primitive ast.PrimitiveLike,
	)
	PreprocessProcedure(
		procedure ast.ProcedureLike,
	)
	ProcessProcedureSlot(
		slot uint,
	)
	PostprocessProcedure(
		procedure ast.ProcedureLike,
	)
	PreprocessPublishClause(
		publishClause ast.PublishClauseLike,
	)
	ProcessPublishClauseSlot(
		slot uint,
	)
	PostprocessPublishClause(
		publishClause ast.PublishClauseLike,
	)
	PreprocessRange(
		range_ ast.RangeLike,
	)
	ProcessRangeSlot(
		slot uint,
	)
	PostprocessRange(
		range_ ast.RangeLike,
	)
	PreprocessRecipient(
		recipient ast.RecipientLike,
	)
	ProcessRecipientSlot(
		slot uint,
	)
	PostprocessRecipient(
		recipient ast.RecipientLike,
	)
	PreprocessReferent(
		referent ast.ReferentLike,
	)
	ProcessReferentSlot(
		slot uint,
	)
	PostprocessReferent(
		referent ast.ReferentLike,
	)
	PreprocessRejectClause(
		rejectClause ast.RejectClauseLike,
	)
	ProcessRejectClauseSlot(
		slot uint,
	)
	PostprocessRejectClause(
		rejectClause ast.RejectClauseLike,
	)
	PreprocessRepository(
		repository ast.RepositoryLike,
	)
	ProcessRepositorySlot(
		slot uint,
	)
	PostprocessRepository(
		repository ast.RepositoryLike,
	)
	PreprocessResult(
		result ast.ResultLike,
	)
	ProcessResultSlot(
		slot uint,
	)
	PostprocessResult(
		result ast.ResultLike,
	)
	PreprocessRetrieveClause(
		retrieveClause ast.RetrieveClauseLike,
	)
	ProcessRetrieveClauseSlot(
		slot uint,
	)
	PostprocessRetrieveClause(
		retrieveClause ast.RetrieveClauseLike,
	)
	PreprocessReturnClause(
		returnClause ast.ReturnClauseLike,
	)
	ProcessReturnClauseSlot(
		slot uint,
	)
	PostprocessReturnClause(
		returnClause ast.ReturnClauseLike,
	)
	PreprocessRight(
		right ast.RightLike,
	)
	ProcessRightSlot(
		slot uint,
	)
	PostprocessRight(
		right ast.RightLike,
	)
	PreprocessSaveClause(
		saveClause ast.SaveClauseLike,
	)
	ProcessSaveClauseSlot(
		slot uint,
	)
	PostprocessSaveClause(
		saveClause ast.SaveClauseLike,
	)
	PreprocessSelectClause(
		selectClause ast.SelectClauseLike,
	)
	ProcessSelectClauseSlot(
		slot uint,
	)
	PostprocessSelectClause(
		selectClause ast.SelectClauseLike,
	)
	PreprocessSequence(
		sequence ast.SequenceLike,
	)
	ProcessSequenceSlot(
		slot uint,
	)
	PostprocessSequence(
		sequence ast.SequenceLike,
	)
	PreprocessStatement(
		statement ast.StatementLike,
	)
	ProcessStatementSlot(
		slot uint,
	)
	PostprocessStatement(
		statement ast.StatementLike,
	)
	PreprocessStatements(
		statements ast.StatementsLike,
	)
	ProcessStatementsSlot(
		slot uint,
	)
	PostprocessStatements(
		statements ast.StatementsLike,
	)
	PreprocessString(
		string_ ast.StringLike,
	)
	ProcessStringSlot(
		slot uint,
	)
	PostprocessString(
		string_ ast.StringLike,
	)
	PreprocessSubcomponent(
		subcomponent ast.SubcomponentLike,
	)
	ProcessSubcomponentSlot(
		slot uint,
	)
	PostprocessSubcomponent(
		subcomponent ast.SubcomponentLike,
	)
	PreprocessSubject(
		subject ast.SubjectLike,
	)
	ProcessSubjectSlot(
		slot uint,
	)
	PostprocessSubject(
		subject ast.SubjectLike,
	)
	PreprocessTarget(
		target ast.TargetLike,
	)
	ProcessTargetSlot(
		slot uint,
	)
	PostprocessTarget(
		target ast.TargetLike,
	)
	PreprocessTemplate(
		template ast.TemplateLike,
	)
	ProcessTemplateSlot(
		slot uint,
	)
	PostprocessTemplate(
		template ast.TemplateLike,
	)
	PreprocessThreading(
		threading ast.ThreadingLike,
	)
	ProcessThreadingSlot(
		slot uint,
	)
	PostprocessThreading(
		threading ast.ThreadingLike,
	)
	PreprocessThrowClause(
		throwClause ast.ThrowClauseLike,
	)
	ProcessThrowClauseSlot(
		slot uint,
	)
	PostprocessThrowClause(
		throwClause ast.ThrowClauseLike,
	)
	PreprocessValue(
		value ast.ValueLike,
	)
	ProcessValueSlot(
		slot uint,
	)
	PostprocessValue(
		value ast.ValueLike,
	)
	PreprocessValues(
		values ast.ValuesLike,
	)
	ProcessValuesSlot(
		slot uint,
	)
	PostprocessValues(
		values ast.ValuesLike,
	)
	PreprocessVariable(
		variable ast.VariableLike,
	)
	ProcessVariableSlot(
		slot uint,
	)
	PostprocessVariable(
		variable ast.VariableLike,
	)
	PreprocessWhileClause(
		whileClause ast.WhileClauseLike,
	)
	ProcessWhileClauseSlot(
		slot uint,
	)
	PostprocessWhileClause(
		whileClause ast.WhileClauseLike,
	)
	PreprocessWithClause(
		withClause ast.WithClauseLike,
	)
	ProcessWithClauseSlot(
		slot uint,
	)
	PostprocessWithClause(
		withClause ast.WithClauseLike,
	)
}
