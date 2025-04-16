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
	CitationToken
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
	ProcessCitation(
		citation string,
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
	PreprocessAnnotationLine(
		annotationLine ast.AnnotationLineLike,
	)
	ProcessAnnotationLineSlot(
		slot uint,
	)
	PostprocessAnnotationLine(
		annotationLine ast.AnnotationLineLike,
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
	PreprocessAtLevel(
		atLevel ast.AtLevelLike,
	)
	ProcessAtLevelSlot(
		slot uint,
	)
	PostprocessAtLevel(
		atLevel ast.AtLevelLike,
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
	PreprocessEmptyLine(
		emptyLine ast.EmptyLineLike,
	)
	ProcessEmptyLineSlot(
		slot uint,
	)
	PostprocessEmptyLine(
		emptyLine ast.EmptyLineLike,
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
	PreprocessExInclusiveAngles(
		exInclusiveAngles ast.ExInclusiveAnglesLike,
	)
	ProcessExInclusiveAnglesSlot(
		slot uint,
	)
	PostprocessExInclusiveAngles(
		exInclusiveAngles ast.ExInclusiveAnglesLike,
	)
	PreprocessExInclusiveCitations(
		exInclusiveCitations ast.ExInclusiveCitationsLike,
	)
	ProcessExInclusiveCitationsSlot(
		slot uint,
	)
	PostprocessExInclusiveCitations(
		exInclusiveCitations ast.ExInclusiveCitationsLike,
	)
	PreprocessExInclusiveDurations(
		exInclusiveDurations ast.ExInclusiveDurationsLike,
	)
	ProcessExInclusiveDurationsSlot(
		slot uint,
	)
	PostprocessExInclusiveDurations(
		exInclusiveDurations ast.ExInclusiveDurationsLike,
	)
	PreprocessExInclusiveMoments(
		exInclusiveMoments ast.ExInclusiveMomentsLike,
	)
	ProcessExInclusiveMomentsSlot(
		slot uint,
	)
	PostprocessExInclusiveMoments(
		exInclusiveMoments ast.ExInclusiveMomentsLike,
	)
	PreprocessExInclusiveNames(
		exInclusiveNames ast.ExInclusiveNamesLike,
	)
	ProcessExInclusiveNamesSlot(
		slot uint,
	)
	PostprocessExInclusiveNames(
		exInclusiveNames ast.ExInclusiveNamesLike,
	)
	PreprocessExInclusiveNumbers(
		exInclusiveNumbers ast.ExInclusiveNumbersLike,
	)
	ProcessExInclusiveNumbersSlot(
		slot uint,
	)
	PostprocessExInclusiveNumbers(
		exInclusiveNumbers ast.ExInclusiveNumbersLike,
	)
	PreprocessExInclusivePercentages(
		exInclusivePercentages ast.ExInclusivePercentagesLike,
	)
	ProcessExInclusivePercentagesSlot(
		slot uint,
	)
	PostprocessExInclusivePercentages(
		exInclusivePercentages ast.ExInclusivePercentagesLike,
	)
	PreprocessExInclusiveProbabilities(
		exInclusiveProbabilities ast.ExInclusiveProbabilitiesLike,
	)
	ProcessExInclusiveProbabilitiesSlot(
		slot uint,
	)
	PostprocessExInclusiveProbabilities(
		exInclusiveProbabilities ast.ExInclusiveProbabilitiesLike,
	)
	PreprocessExInclusiveQuotes(
		exInclusiveQuotes ast.ExInclusiveQuotesLike,
	)
	ProcessExInclusiveQuotesSlot(
		slot uint,
	)
	PostprocessExInclusiveQuotes(
		exInclusiveQuotes ast.ExInclusiveQuotesLike,
	)
	PreprocessExInclusiveSymbols(
		exInclusiveSymbols ast.ExInclusiveSymbolsLike,
	)
	ProcessExInclusiveSymbolsSlot(
		slot uint,
	)
	PostprocessExInclusiveSymbols(
		exInclusiveSymbols ast.ExInclusiveSymbolsLike,
	)
	PreprocessExInclusiveVersions(
		exInclusiveVersions ast.ExInclusiveVersionsLike,
	)
	ProcessExInclusiveVersionsSlot(
		slot uint,
	)
	PostprocessExInclusiveVersions(
		exInclusiveVersions ast.ExInclusiveVersionsLike,
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
	PreprocessExclusiveAngles(
		exclusiveAngles ast.ExclusiveAnglesLike,
	)
	ProcessExclusiveAnglesSlot(
		slot uint,
	)
	PostprocessExclusiveAngles(
		exclusiveAngles ast.ExclusiveAnglesLike,
	)
	PreprocessExclusiveCitations(
		exclusiveCitations ast.ExclusiveCitationsLike,
	)
	ProcessExclusiveCitationsSlot(
		slot uint,
	)
	PostprocessExclusiveCitations(
		exclusiveCitations ast.ExclusiveCitationsLike,
	)
	PreprocessExclusiveDurations(
		exclusiveDurations ast.ExclusiveDurationsLike,
	)
	ProcessExclusiveDurationsSlot(
		slot uint,
	)
	PostprocessExclusiveDurations(
		exclusiveDurations ast.ExclusiveDurationsLike,
	)
	PreprocessExclusiveMoments(
		exclusiveMoments ast.ExclusiveMomentsLike,
	)
	ProcessExclusiveMomentsSlot(
		slot uint,
	)
	PostprocessExclusiveMoments(
		exclusiveMoments ast.ExclusiveMomentsLike,
	)
	PreprocessExclusiveNames(
		exclusiveNames ast.ExclusiveNamesLike,
	)
	ProcessExclusiveNamesSlot(
		slot uint,
	)
	PostprocessExclusiveNames(
		exclusiveNames ast.ExclusiveNamesLike,
	)
	PreprocessExclusiveNumbers(
		exclusiveNumbers ast.ExclusiveNumbersLike,
	)
	ProcessExclusiveNumbersSlot(
		slot uint,
	)
	PostprocessExclusiveNumbers(
		exclusiveNumbers ast.ExclusiveNumbersLike,
	)
	PreprocessExclusivePercentages(
		exclusivePercentages ast.ExclusivePercentagesLike,
	)
	ProcessExclusivePercentagesSlot(
		slot uint,
	)
	PostprocessExclusivePercentages(
		exclusivePercentages ast.ExclusivePercentagesLike,
	)
	PreprocessExclusiveProbabilities(
		exclusiveProbabilities ast.ExclusiveProbabilitiesLike,
	)
	ProcessExclusiveProbabilitiesSlot(
		slot uint,
	)
	PostprocessExclusiveProbabilities(
		exclusiveProbabilities ast.ExclusiveProbabilitiesLike,
	)
	PreprocessExclusiveQuotes(
		exclusiveQuotes ast.ExclusiveQuotesLike,
	)
	ProcessExclusiveQuotesSlot(
		slot uint,
	)
	PostprocessExclusiveQuotes(
		exclusiveQuotes ast.ExclusiveQuotesLike,
	)
	PreprocessExclusiveSymbols(
		exclusiveSymbols ast.ExclusiveSymbolsLike,
	)
	ProcessExclusiveSymbolsSlot(
		slot uint,
	)
	PostprocessExclusiveSymbols(
		exclusiveSymbols ast.ExclusiveSymbolsLike,
	)
	PreprocessExclusiveVersions(
		exclusiveVersions ast.ExclusiveVersionsLike,
	)
	ProcessExclusiveVersionsSlot(
		slot uint,
	)
	PostprocessExclusiveVersions(
		exclusiveVersions ast.ExclusiveVersionsLike,
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
	PreprocessInExclusiveAngles(
		inExclusiveAngles ast.InExclusiveAnglesLike,
	)
	ProcessInExclusiveAnglesSlot(
		slot uint,
	)
	PostprocessInExclusiveAngles(
		inExclusiveAngles ast.InExclusiveAnglesLike,
	)
	PreprocessInExclusiveCitations(
		inExclusiveCitations ast.InExclusiveCitationsLike,
	)
	ProcessInExclusiveCitationsSlot(
		slot uint,
	)
	PostprocessInExclusiveCitations(
		inExclusiveCitations ast.InExclusiveCitationsLike,
	)
	PreprocessInExclusiveDurations(
		inExclusiveDurations ast.InExclusiveDurationsLike,
	)
	ProcessInExclusiveDurationsSlot(
		slot uint,
	)
	PostprocessInExclusiveDurations(
		inExclusiveDurations ast.InExclusiveDurationsLike,
	)
	PreprocessInExclusiveMoments(
		inExclusiveMoments ast.InExclusiveMomentsLike,
	)
	ProcessInExclusiveMomentsSlot(
		slot uint,
	)
	PostprocessInExclusiveMoments(
		inExclusiveMoments ast.InExclusiveMomentsLike,
	)
	PreprocessInExclusiveNames(
		inExclusiveNames ast.InExclusiveNamesLike,
	)
	ProcessInExclusiveNamesSlot(
		slot uint,
	)
	PostprocessInExclusiveNames(
		inExclusiveNames ast.InExclusiveNamesLike,
	)
	PreprocessInExclusiveNumbers(
		inExclusiveNumbers ast.InExclusiveNumbersLike,
	)
	ProcessInExclusiveNumbersSlot(
		slot uint,
	)
	PostprocessInExclusiveNumbers(
		inExclusiveNumbers ast.InExclusiveNumbersLike,
	)
	PreprocessInExclusivePercentages(
		inExclusivePercentages ast.InExclusivePercentagesLike,
	)
	ProcessInExclusivePercentagesSlot(
		slot uint,
	)
	PostprocessInExclusivePercentages(
		inExclusivePercentages ast.InExclusivePercentagesLike,
	)
	PreprocessInExclusiveProbabilities(
		inExclusiveProbabilities ast.InExclusiveProbabilitiesLike,
	)
	ProcessInExclusiveProbabilitiesSlot(
		slot uint,
	)
	PostprocessInExclusiveProbabilities(
		inExclusiveProbabilities ast.InExclusiveProbabilitiesLike,
	)
	PreprocessInExclusiveQuotes(
		inExclusiveQuotes ast.InExclusiveQuotesLike,
	)
	ProcessInExclusiveQuotesSlot(
		slot uint,
	)
	PostprocessInExclusiveQuotes(
		inExclusiveQuotes ast.InExclusiveQuotesLike,
	)
	PreprocessInExclusiveSymbols(
		inExclusiveSymbols ast.InExclusiveSymbolsLike,
	)
	ProcessInExclusiveSymbolsSlot(
		slot uint,
	)
	PostprocessInExclusiveSymbols(
		inExclusiveSymbols ast.InExclusiveSymbolsLike,
	)
	PreprocessInExclusiveVersions(
		inExclusiveVersions ast.InExclusiveVersionsLike,
	)
	ProcessInExclusiveVersionsSlot(
		slot uint,
	)
	PostprocessInExclusiveVersions(
		inExclusiveVersions ast.InExclusiveVersionsLike,
	)
	PreprocessInclusiveAngles(
		inclusiveAngles ast.InclusiveAnglesLike,
	)
	ProcessInclusiveAnglesSlot(
		slot uint,
	)
	PostprocessInclusiveAngles(
		inclusiveAngles ast.InclusiveAnglesLike,
	)
	PreprocessInclusiveCitations(
		inclusiveCitations ast.InclusiveCitationsLike,
	)
	ProcessInclusiveCitationsSlot(
		slot uint,
	)
	PostprocessInclusiveCitations(
		inclusiveCitations ast.InclusiveCitationsLike,
	)
	PreprocessInclusiveDurations(
		inclusiveDurations ast.InclusiveDurationsLike,
	)
	ProcessInclusiveDurationsSlot(
		slot uint,
	)
	PostprocessInclusiveDurations(
		inclusiveDurations ast.InclusiveDurationsLike,
	)
	PreprocessInclusiveMoments(
		inclusiveMoments ast.InclusiveMomentsLike,
	)
	ProcessInclusiveMomentsSlot(
		slot uint,
	)
	PostprocessInclusiveMoments(
		inclusiveMoments ast.InclusiveMomentsLike,
	)
	PreprocessInclusiveNames(
		inclusiveNames ast.InclusiveNamesLike,
	)
	ProcessInclusiveNamesSlot(
		slot uint,
	)
	PostprocessInclusiveNames(
		inclusiveNames ast.InclusiveNamesLike,
	)
	PreprocessInclusiveNumbers(
		inclusiveNumbers ast.InclusiveNumbersLike,
	)
	ProcessInclusiveNumbersSlot(
		slot uint,
	)
	PostprocessInclusiveNumbers(
		inclusiveNumbers ast.InclusiveNumbersLike,
	)
	PreprocessInclusivePercentages(
		inclusivePercentages ast.InclusivePercentagesLike,
	)
	ProcessInclusivePercentagesSlot(
		slot uint,
	)
	PostprocessInclusivePercentages(
		inclusivePercentages ast.InclusivePercentagesLike,
	)
	PreprocessInclusiveProbabilities(
		inclusiveProbabilities ast.InclusiveProbabilitiesLike,
	)
	ProcessInclusiveProbabilitiesSlot(
		slot uint,
	)
	PostprocessInclusiveProbabilities(
		inclusiveProbabilities ast.InclusiveProbabilitiesLike,
	)
	PreprocessInclusiveQuotes(
		inclusiveQuotes ast.InclusiveQuotesLike,
	)
	ProcessInclusiveQuotesSlot(
		slot uint,
	)
	PostprocessInclusiveQuotes(
		inclusiveQuotes ast.InclusiveQuotesLike,
	)
	PreprocessInclusiveSymbols(
		inclusiveSymbols ast.InclusiveSymbolsLike,
	)
	ProcessInclusiveSymbolsSlot(
		slot uint,
	)
	PostprocessInclusiveSymbols(
		inclusiveSymbols ast.InclusiveSymbolsLike,
	)
	PreprocessInclusiveVersions(
		inclusiveVersions ast.InclusiveVersionsLike,
	)
	ProcessInclusiveVersionsSlot(
		slot uint,
	)
	PostprocessInclusiveVersions(
		inclusiveVersions ast.InclusiveVersionsLike,
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
	PreprocessInlineAttributes(
		inlineAttributes ast.InlineAttributesLike,
	)
	ProcessInlineAttributesSlot(
		slot uint,
	)
	PostprocessInlineAttributes(
		inlineAttributes ast.InlineAttributesLike,
	)
	PreprocessInlineParameters(
		inlineParameters ast.InlineParametersLike,
	)
	ProcessInlineParametersSlot(
		slot uint,
	)
	PostprocessInlineParameters(
		inlineParameters ast.InlineParametersLike,
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
	PreprocessMultilineAttributes(
		multilineAttributes ast.MultilineAttributesLike,
	)
	ProcessMultilineAttributesSlot(
		slot uint,
	)
	PostprocessMultilineAttributes(
		multilineAttributes ast.MultilineAttributesLike,
	)
	PreprocessMultilineParameters(
		multilineParameters ast.MultilineParametersLike,
	)
	ProcessMultilineParametersSlot(
		slot uint,
	)
	PostprocessMultilineParameters(
		multilineParameters ast.MultilineParametersLike,
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
	PreprocessNoAttributes(
		noAttributes ast.NoAttributesLike,
	)
	ProcessNoAttributesSlot(
		slot uint,
	)
	PostprocessNoAttributes(
		noAttributes ast.NoAttributesLike,
	)
	PreprocessNoStatements(
		noStatements ast.NoStatementsLike,
	)
	ProcessNoStatementsSlot(
		slot uint,
	)
	PostprocessNoStatements(
		noStatements ast.NoStatementsLike,
	)
	PreprocessNoValues(
		noValues ast.NoValuesLike,
	)
	ProcessNoValuesSlot(
		slot uint,
	)
	PostprocessNoValues(
		noValues ast.NoValuesLike,
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
	PreprocessStatementLine(
		statementLine ast.StatementLineLike,
	)
	ProcessStatementLineSlot(
		slot uint,
	)
	PostprocessStatementLine(
		statementLine ast.StatementLineLike,
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
