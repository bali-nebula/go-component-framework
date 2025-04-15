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
│              This "Module.go" file was automatically generated.              │
│      Updates to any part of this file—other than the Module Description      │
│             and the Global Functions sections may be overwritten.            │
└──────────────────────────────────────────────────────────────────────────────┘

Package "module" declares type aliases for the commonly used types declared in
the packages contained in this module.  It also provides constructors for each
commonly used class that is exported by the module.  Each constructor delegates
the actual construction process to its corresponding concrete class declared in
the corresponding package contained within this module.

For detailed documentation on this entire module refer to the wiki:
  - https://github.com/bali-nebula/go-component-framework/wiki
*/
package module

import (
	ast "github.com/bali-nebula/go-component-framework/v3/ast"
	gra "github.com/bali-nebula/go-component-framework/v3/grammar"
	col "github.com/craterdog/go-collection-framework/v5/collection"
)

// TYPE ALIASES

// Ast

type (
	AcceptClauseClassLike          = ast.AcceptClauseClassLike
	AdditionalArgumentClassLike    = ast.AdditionalArgumentClassLike
	AdditionalAssociationClassLike = ast.AdditionalAssociationClassLike
	AdditionalIndexClassLike       = ast.AdditionalIndexClassLike
	AdditionalStatementClassLike   = ast.AdditionalStatementClassLike
	AdditionalValueClassLike       = ast.AdditionalValueClassLike
	AnnotatedAssociationClassLike  = ast.AnnotatedAssociationClassLike
	AnnotatedStatementClassLike    = ast.AnnotatedStatementClassLike
	AnnotatedValueClassLike        = ast.AnnotatedValueClassLike
	ArgumentClassLike              = ast.ArgumentClassLike
	ArgumentsClassLike             = ast.ArgumentsClassLike
	AssignClassLike                = ast.AssignClassLike
	AssociationClassLike           = ast.AssociationClassLike
	AssociationsClassLike          = ast.AssociationsClassLike
	AtLevelClassLike               = ast.AtLevelClassLike
	AttributeClassLike             = ast.AttributeClassLike
	BagClassLike                   = ast.BagClassLike
	BreakClauseClassLike           = ast.BreakClauseClassLike
	CheckoutClauseClassLike        = ast.CheckoutClauseClassLike
	CitationClassLike              = ast.CitationClassLike
	CollectionClassLike            = ast.CollectionClassLike
	ComplementClassLike            = ast.ComplementClassLike
	ComponentClassLike             = ast.ComponentClassLike
	ConditionClassLike             = ast.ConditionClassLike
	ContinueClauseClassLike        = ast.ContinueClauseClassLike
	DiscardClauseClassLike         = ast.DiscardClauseClassLike
	DoClauseClassLike              = ast.DoClauseClassLike
	DocumentClassLike              = ast.DocumentClassLike
	DraftClassLike                 = ast.DraftClassLike
	ElementClassLike               = ast.ElementClassLike
	EntityClassLike                = ast.EntityClassLike
	EventClassLike                 = ast.EventClassLike
	ExceptionClassLike             = ast.ExceptionClassLike
	ExpressionClassLike            = ast.ExpressionClassLike
	FailureClassLike               = ast.FailureClassLike
	FlowClassLike                  = ast.FlowClassLike
	FunctionClassLike              = ast.FunctionClassLike
	IfClauseClassLike              = ast.IfClauseClassLike
	InclusionClassLike             = ast.InclusionClassLike
	IndexClassLike                 = ast.IndexClassLike
	IndicesClassLike               = ast.IndicesClassLike
	IndirectClassLike              = ast.IndirectClassLike
	InductionClassLike             = ast.InductionClassLike
	InlineAssociationsClassLike    = ast.InlineAssociationsClassLike
	InlineStatementsClassLike      = ast.InlineStatementsClassLike
	InlineValuesClassLike          = ast.InlineValuesClassLike
	InverseClassLike               = ast.InverseClassLike
	InversionClassLike             = ast.InversionClassLike
	InvocationClassLike            = ast.InvocationClassLike
	ItemClassLike                  = ast.ItemClassLike
	ItemsClassLike                 = ast.ItemsClassLike
	LeftClassLike                  = ast.LeftClassLike
	LetClauseClassLike             = ast.LetClauseClassLike
	LogicalClassLike               = ast.LogicalClassLike
	MagnitudeClassLike             = ast.MagnitudeClassLike
	MainClauseClassLike            = ast.MainClauseClassLike
	MatchHandlerClassLike          = ast.MatchHandlerClassLike
	MessageClassLike               = ast.MessageClassLike
	MessagingClassLike             = ast.MessagingClassLike
	MethodClassLike                = ast.MethodClassLike
	MultilineAssociationsClassLike = ast.MultilineAssociationsClassLike
	MultilineStatementsClassLike   = ast.MultilineStatementsClassLike
	MultilineValuesClassLike       = ast.MultilineValuesClassLike
	NoAssociationsClassLike        = ast.NoAssociationsClassLike
	NotarizeClauseClassLike        = ast.NotarizeClauseClassLike
	NumericalClassLike             = ast.NumericalClassLike
	OnClauseClassLike              = ast.OnClauseClassLike
	OperatorClassLike              = ast.OperatorClassLike
	ParametersClassLike            = ast.ParametersClassLike
	PostClauseClassLike            = ast.PostClauseClassLike
	PrecedenceClassLike            = ast.PrecedenceClassLike
	PredicateClassLike             = ast.PredicateClassLike
	PrimitiveClassLike             = ast.PrimitiveClassLike
	ProcedureClassLike             = ast.ProcedureClassLike
	PublishClauseClassLike         = ast.PublishClauseClassLike
	RangeClassLike                 = ast.RangeClassLike
	RecipientClassLike             = ast.RecipientClassLike
	ReferentClassLike              = ast.ReferentClassLike
	RejectClauseClassLike          = ast.RejectClauseClassLike
	RepositoryClassLike            = ast.RepositoryClassLike
	ResultClassLike                = ast.ResultClassLike
	RetrieveClauseClassLike        = ast.RetrieveClauseClassLike
	ReturnClauseClassLike          = ast.ReturnClauseClassLike
	RightClassLike                 = ast.RightClassLike
	SaveClauseClassLike            = ast.SaveClauseClassLike
	SelectClauseClassLike          = ast.SelectClauseClassLike
	SequenceClassLike              = ast.SequenceClassLike
	StatementClassLike             = ast.StatementClassLike
	StatementsClassLike            = ast.StatementsClassLike
	StringClassLike                = ast.StringClassLike
	SubcomponentClassLike          = ast.SubcomponentClassLike
	SubjectClassLike               = ast.SubjectClassLike
	TargetClassLike                = ast.TargetClassLike
	TemplateClassLike              = ast.TemplateClassLike
	ThreadingClassLike             = ast.ThreadingClassLike
	ThrowClauseClassLike           = ast.ThrowClauseClassLike
	ValueClassLike                 = ast.ValueClassLike
	ValuesClassLike                = ast.ValuesClassLike
	VariableClassLike              = ast.VariableClassLike
	WhileClauseClassLike           = ast.WhileClauseClassLike
	WithClauseClassLike            = ast.WithClauseClassLike
)

type (
	AcceptClauseLike          = ast.AcceptClauseLike
	AdditionalArgumentLike    = ast.AdditionalArgumentLike
	AdditionalAssociationLike = ast.AdditionalAssociationLike
	AdditionalIndexLike       = ast.AdditionalIndexLike
	AdditionalStatementLike   = ast.AdditionalStatementLike
	AdditionalValueLike       = ast.AdditionalValueLike
	AnnotatedAssociationLike  = ast.AnnotatedAssociationLike
	AnnotatedStatementLike    = ast.AnnotatedStatementLike
	AnnotatedValueLike        = ast.AnnotatedValueLike
	ArgumentLike              = ast.ArgumentLike
	ArgumentsLike             = ast.ArgumentsLike
	AssignLike                = ast.AssignLike
	AssociationLike           = ast.AssociationLike
	AssociationsLike          = ast.AssociationsLike
	AtLevelLike               = ast.AtLevelLike
	AttributeLike             = ast.AttributeLike
	BagLike                   = ast.BagLike
	BreakClauseLike           = ast.BreakClauseLike
	CheckoutClauseLike        = ast.CheckoutClauseLike
	CitationLike              = ast.CitationLike
	CollectionLike            = ast.CollectionLike
	ComplementLike            = ast.ComplementLike
	ComponentLike             = ast.ComponentLike
	ConditionLike             = ast.ConditionLike
	ContinueClauseLike        = ast.ContinueClauseLike
	DiscardClauseLike         = ast.DiscardClauseLike
	DoClauseLike              = ast.DoClauseLike
	DocumentLike              = ast.DocumentLike
	DraftLike                 = ast.DraftLike
	ElementLike               = ast.ElementLike
	EntityLike                = ast.EntityLike
	EventLike                 = ast.EventLike
	ExceptionLike             = ast.ExceptionLike
	ExpressionLike            = ast.ExpressionLike
	FailureLike               = ast.FailureLike
	FlowLike                  = ast.FlowLike
	FunctionLike              = ast.FunctionLike
	IfClauseLike              = ast.IfClauseLike
	InclusionLike             = ast.InclusionLike
	IndexLike                 = ast.IndexLike
	IndicesLike               = ast.IndicesLike
	IndirectLike              = ast.IndirectLike
	InductionLike             = ast.InductionLike
	InlineAssociationsLike    = ast.InlineAssociationsLike
	InlineStatementsLike      = ast.InlineStatementsLike
	InlineValuesLike          = ast.InlineValuesLike
	InverseLike               = ast.InverseLike
	InversionLike             = ast.InversionLike
	InvocationLike            = ast.InvocationLike
	ItemLike                  = ast.ItemLike
	ItemsLike                 = ast.ItemsLike
	LeftLike                  = ast.LeftLike
	LetClauseLike             = ast.LetClauseLike
	LogicalLike               = ast.LogicalLike
	MagnitudeLike             = ast.MagnitudeLike
	MainClauseLike            = ast.MainClauseLike
	MatchHandlerLike          = ast.MatchHandlerLike
	MessageLike               = ast.MessageLike
	MessagingLike             = ast.MessagingLike
	MethodLike                = ast.MethodLike
	MultilineAssociationsLike = ast.MultilineAssociationsLike
	MultilineStatementsLike   = ast.MultilineStatementsLike
	MultilineValuesLike       = ast.MultilineValuesLike
	NoAssociationsLike        = ast.NoAssociationsLike
	NotarizeClauseLike        = ast.NotarizeClauseLike
	NumericalLike             = ast.NumericalLike
	OnClauseLike              = ast.OnClauseLike
	OperatorLike              = ast.OperatorLike
	ParametersLike            = ast.ParametersLike
	PostClauseLike            = ast.PostClauseLike
	PrecedenceLike            = ast.PrecedenceLike
	PredicateLike             = ast.PredicateLike
	PrimitiveLike             = ast.PrimitiveLike
	ProcedureLike             = ast.ProcedureLike
	PublishClauseLike         = ast.PublishClauseLike
	RangeLike                 = ast.RangeLike
	RecipientLike             = ast.RecipientLike
	ReferentLike              = ast.ReferentLike
	RejectClauseLike          = ast.RejectClauseLike
	RepositoryLike            = ast.RepositoryLike
	ResultLike                = ast.ResultLike
	RetrieveClauseLike        = ast.RetrieveClauseLike
	ReturnClauseLike          = ast.ReturnClauseLike
	RightLike                 = ast.RightLike
	SaveClauseLike            = ast.SaveClauseLike
	SelectClauseLike          = ast.SelectClauseLike
	SequenceLike              = ast.SequenceLike
	StatementLike             = ast.StatementLike
	StatementsLike            = ast.StatementsLike
	StringLike                = ast.StringLike
	SubcomponentLike          = ast.SubcomponentLike
	SubjectLike               = ast.SubjectLike
	TargetLike                = ast.TargetLike
	TemplateLike              = ast.TemplateLike
	ThreadingLike             = ast.ThreadingLike
	ThrowClauseLike           = ast.ThrowClauseLike
	ValueLike                 = ast.ValueLike
	ValuesLike                = ast.ValuesLike
	VariableLike              = ast.VariableLike
	WhileClauseLike           = ast.WhileClauseLike
	WithClauseLike            = ast.WithClauseLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken        = gra.ErrorToken
	AmpersandToken    = gra.AmpersandToken
	AndToken          = gra.AndToken
	AngleToken        = gra.AngleToken
	ArrowToken        = gra.ArrowToken
	BarToken          = gra.BarToken
	BinaryToken       = gra.BinaryToken
	BooleanToken      = gra.BooleanToken
	BytecodeToken     = gra.BytecodeToken
	CaretToken        = gra.CaretToken
	ColonToken        = gra.ColonToken
	ColonEqualToken   = gra.ColonEqualToken
	CommentToken      = gra.CommentToken
	DashToken         = gra.DashToken
	DashEqualToken    = gra.DashEqualToken
	DefaultEqualToken = gra.DefaultEqualToken
	DelimiterToken    = gra.DelimiterToken
	DotToken          = gra.DotToken
	DoubleToken       = gra.DoubleToken
	DurationToken     = gra.DurationToken
	EqualToken        = gra.EqualToken
	IdentifierToken   = gra.IdentifierToken
	IorToken          = gra.IorToken
	IsToken           = gra.IsToken
	LessToken         = gra.LessToken
	MatchesToken      = gra.MatchesToken
	MomentToken       = gra.MomentToken
	MoreToken         = gra.MoreToken
	NameToken         = gra.NameToken
	NarrativeToken    = gra.NarrativeToken
	NewlineToken      = gra.NewlineToken
	NotToken          = gra.NotToken
	NoteToken         = gra.NoteToken
	NumberToken       = gra.NumberToken
	PatternToken      = gra.PatternToken
	PercentageToken   = gra.PercentageToken
	PlusToken         = gra.PlusToken
	PlusEqualToken    = gra.PlusEqualToken
	ProbabilityToken  = gra.ProbabilityToken
	QuoteToken        = gra.QuoteToken
	ResourceToken     = gra.ResourceToken
	SanToken          = gra.SanToken
	SlashToken        = gra.SlashToken
	SlashEqualToken   = gra.SlashEqualToken
	SnailToken        = gra.SnailToken
	SpaceToken        = gra.SpaceToken
	StarToken         = gra.StarToken
	StarEqualToken    = gra.StarEqualToken
	SymbolToken       = gra.SymbolToken
	TagToken          = gra.TagToken
	VersionToken      = gra.VersionToken
	XorToken          = gra.XorToken
	XyzToken          = gra.XyzToken
)

type (
	FormatterClassLike = gra.FormatterClassLike
	ParserClassLike    = gra.ParserClassLike
	ProcessorClassLike = gra.ProcessorClassLike
	ScannerClassLike   = gra.ScannerClassLike
	TokenClassLike     = gra.TokenClassLike
	ValidatorClassLike = gra.ValidatorClassLike
	VisitorClassLike   = gra.VisitorClassLike
)

type (
	FormatterLike = gra.FormatterLike
	ParserLike    = gra.ParserLike
	ProcessorLike = gra.ProcessorLike
	ScannerLike   = gra.ScannerLike
	TokenLike     = gra.TokenLike
	ValidatorLike = gra.ValidatorLike
	VisitorLike   = gra.VisitorLike
)

type (
	Methodical = gra.Methodical
)

// CLASS CONSTRUCTORS

// Ast/AcceptClause

func AcceptClause(
	message ast.MessageLike,
) ast.AcceptClauseLike {
	return ast.AcceptClauseClass().AcceptClause(
		message,
	)
}

// Ast/AdditionalArgument

func AdditionalArgument(
	argument ast.ArgumentLike,
) ast.AdditionalArgumentLike {
	return ast.AdditionalArgumentClass().AdditionalArgument(
		argument,
	)
}

// Ast/AdditionalAssociation

func AdditionalAssociation(
	association ast.AssociationLike,
) ast.AdditionalAssociationLike {
	return ast.AdditionalAssociationClass().AdditionalAssociation(
		association,
	)
}

// Ast/AdditionalIndex

func AdditionalIndex(
	index ast.IndexLike,
) ast.AdditionalIndexLike {
	return ast.AdditionalIndexClass().AdditionalIndex(
		index,
	)
}

// Ast/AdditionalStatement

func AdditionalStatement(
	statement ast.StatementLike,
) ast.AdditionalStatementLike {
	return ast.AdditionalStatementClass().AdditionalStatement(
		statement,
	)
}

// Ast/AdditionalValue

func AdditionalValue(
	value ast.ValueLike,
) ast.AdditionalValueLike {
	return ast.AdditionalValueClass().AdditionalValue(
		value,
	)
}

// Ast/AnnotatedAssociation

func AnnotatedAssociation(
	xyz string,
	association ast.AssociationLike,
	optionalNote string,
) ast.AnnotatedAssociationLike {
	return ast.AnnotatedAssociationClass().AnnotatedAssociation(
		xyz,
		association,
		optionalNote,
	)
}

// Ast/AnnotatedStatement

func AnnotatedStatement(
	xyz string,
	statement ast.StatementLike,
	optionalNote string,
) ast.AnnotatedStatementLike {
	return ast.AnnotatedStatementClass().AnnotatedStatement(
		xyz,
		statement,
		optionalNote,
	)
}

// Ast/AnnotatedValue

func AnnotatedValue(
	xyz string,
	value ast.ValueLike,
	optionalNote string,
) ast.AnnotatedValueLike {
	return ast.AnnotatedValueClass().AnnotatedValue(
		xyz,
		value,
		optionalNote,
	)
}

// Ast/Argument

func Argument(
	identifier string,
) ast.ArgumentLike {
	return ast.ArgumentClass().Argument(
		identifier,
	)
}

// Ast/Arguments

func Arguments(
	argument ast.ArgumentLike,
	additionalArguments col.Sequential[ast.AdditionalArgumentLike],
) ast.ArgumentsLike {
	return ast.ArgumentsClass().Arguments(
		argument,
		additionalArguments,
	)
}

// Ast/Assign

func Assign(
	any_ any,
) ast.AssignLike {
	return ast.AssignClass().Assign(
		any_,
	)
}

// Ast/Association

func Association(
	symbol string,
	colon string,
	value ast.ValueLike,
) ast.AssociationLike {
	return ast.AssociationClass().Association(
		symbol,
		colon,
		value,
	)
}

// Ast/Associations

func Associations(
	any_ any,
) ast.AssociationsLike {
	return ast.AssociationsClass().Associations(
		any_,
	)
}

// Ast/AtLevel

func AtLevel(
	expression ast.ExpressionLike,
) ast.AtLevelLike {
	return ast.AtLevelClass().AtLevel(
		expression,
	)
}

// Ast/Attribute

func Attribute(
	identifier string,
	subcomponent ast.SubcomponentLike,
) ast.AttributeLike {
	return ast.AttributeClass().Attribute(
		identifier,
		subcomponent,
	)
}

// Ast/Bag

func Bag(
	expression ast.ExpressionLike,
) ast.BagLike {
	return ast.BagClass().Bag(
		expression,
	)
}

// Ast/BreakClause

func BreakClause() ast.BreakClauseLike {
	return ast.BreakClauseClass().BreakClause()
}

// Ast/CheckoutClause

func CheckoutClause(
	recipient ast.RecipientLike,
	optionalAtLevel ast.AtLevelLike,
	citation ast.CitationLike,
) ast.CheckoutClauseLike {
	return ast.CheckoutClauseClass().CheckoutClause(
		recipient,
		optionalAtLevel,
		citation,
	)
}

// Ast/Citation

func Citation(
	expression ast.ExpressionLike,
) ast.CitationLike {
	return ast.CitationClass().Citation(
		expression,
	)
}

// Ast/Collection

func Collection(
	optionalItems ast.ItemsLike,
) ast.CollectionLike {
	return ast.CollectionClass().Collection(
		optionalItems,
	)
}

// Ast/Complement

func Complement(
	not string,
	logical ast.LogicalLike,
) ast.ComplementLike {
	return ast.ComplementClass().Complement(
		not,
		logical,
	)
}

// Ast/Component

func Component(
	entity ast.EntityLike,
	optionalParameters ast.ParametersLike,
) ast.ComponentLike {
	return ast.ComponentClass().Component(
		entity,
		optionalParameters,
	)
}

// Ast/Condition

func Condition(
	expression ast.ExpressionLike,
) ast.ConditionLike {
	return ast.ConditionClass().Condition(
		expression,
	)
}

// Ast/ContinueClause

func ContinueClause() ast.ContinueClauseLike {
	return ast.ContinueClauseClass().ContinueClause()
}

// Ast/DiscardClause

func DiscardClause(
	draft ast.DraftLike,
) ast.DiscardClauseLike {
	return ast.DiscardClauseClass().DiscardClause(
		draft,
	)
}

// Ast/DoClause

func DoClause(
	invocation ast.InvocationLike,
) ast.DoClauseLike {
	return ast.DoClauseClass().DoClause(
		invocation,
	)
}

// Ast/Document

func Document(
	optionalComment string,
	component ast.ComponentLike,
) ast.DocumentLike {
	return ast.DocumentClass().Document(
		optionalComment,
		component,
	)
}

// Ast/Draft

func Draft(
	expression ast.ExpressionLike,
) ast.DraftLike {
	return ast.DraftClass().Draft(
		expression,
	)
}

// Ast/Element

func Element(
	any_ any,
) ast.ElementLike {
	return ast.ElementClass().Element(
		any_,
	)
}

// Ast/Entity

func Entity(
	any_ any,
) ast.EntityLike {
	return ast.EntityClass().Entity(
		any_,
	)
}

// Ast/Event

func Event(
	expression ast.ExpressionLike,
) ast.EventLike {
	return ast.EventClass().Event(
		expression,
	)
}

// Ast/Exception

func Exception(
	expression ast.ExpressionLike,
) ast.ExceptionLike {
	return ast.ExceptionClass().Exception(
		expression,
	)
}

// Ast/Expression

func Expression(
	subject ast.SubjectLike,
	predicates col.Sequential[ast.PredicateLike],
) ast.ExpressionLike {
	return ast.ExpressionClass().Expression(
		subject,
		predicates,
	)
}

// Ast/Failure

func Failure(
	symbol string,
) ast.FailureLike {
	return ast.FailureClass().Failure(
		symbol,
	)
}

// Ast/Flow

func Flow(
	any_ any,
) ast.FlowLike {
	return ast.FlowClass().Flow(
		any_,
	)
}

// Ast/Function

func Function(
	identifier string,
	optionalArguments ast.ArgumentsLike,
) ast.FunctionLike {
	return ast.FunctionClass().Function(
		identifier,
		optionalArguments,
	)
}

// Ast/IfClause

func IfClause(
	condition ast.ConditionLike,
	procedure ast.ProcedureLike,
) ast.IfClauseLike {
	return ast.IfClauseClass().IfClause(
		condition,
		procedure,
	)
}

// Ast/Inclusion

func Inclusion(
	any_ any,
) ast.InclusionLike {
	return ast.InclusionClass().Inclusion(
		any_,
	)
}

// Ast/Index

func Index(
	expression ast.ExpressionLike,
) ast.IndexLike {
	return ast.IndexClass().Index(
		expression,
	)
}

// Ast/Indices

func Indices(
	index ast.IndexLike,
	additionalIndexes col.Sequential[ast.AdditionalIndexLike],
) ast.IndicesLike {
	return ast.IndicesClass().Indices(
		index,
		additionalIndexes,
	)
}

// Ast/Indirect

func Indirect(
	any_ any,
) ast.IndirectLike {
	return ast.IndirectClass().Indirect(
		any_,
	)
}

// Ast/Induction

func Induction(
	any_ any,
) ast.InductionLike {
	return ast.InductionClass().Induction(
		any_,
	)
}

// Ast/InlineAssociations

func InlineAssociations(
	association ast.AssociationLike,
	additionalAssociations col.Sequential[ast.AdditionalAssociationLike],
) ast.InlineAssociationsLike {
	return ast.InlineAssociationsClass().InlineAssociations(
		association,
		additionalAssociations,
	)
}

// Ast/InlineStatements

func InlineStatements(
	statement ast.StatementLike,
	additionalStatements col.Sequential[ast.AdditionalStatementLike],
) ast.InlineStatementsLike {
	return ast.InlineStatementsClass().InlineStatements(
		statement,
		additionalStatements,
	)
}

// Ast/InlineValues

func InlineValues(
	value ast.ValueLike,
	additionalValues col.Sequential[ast.AdditionalValueLike],
) ast.InlineValuesLike {
	return ast.InlineValuesClass().InlineValues(
		value,
		additionalValues,
	)
}

// Ast/Inverse

func Inverse(
	any_ any,
) ast.InverseLike {
	return ast.InverseClass().Inverse(
		any_,
	)
}

// Ast/Inversion

func Inversion(
	inverse ast.InverseLike,
	numerical ast.NumericalLike,
) ast.InversionLike {
	return ast.InversionClass().Inversion(
		inverse,
		numerical,
	)
}

// Ast/Invocation

func Invocation(
	any_ any,
) ast.InvocationLike {
	return ast.InvocationClass().Invocation(
		any_,
	)
}

// Ast/Item

func Item(
	symbol string,
) ast.ItemLike {
	return ast.ItemClass().Item(
		symbol,
	)
}

// Ast/Items

func Items(
	any_ any,
) ast.ItemsLike {
	return ast.ItemsClass().Items(
		any_,
	)
}

// Ast/Left

func Left() ast.LeftLike {
	return ast.LeftClass().Left()
}

// Ast/LetClause

func LetClause(
	recipient ast.RecipientLike,
	assign ast.AssignLike,
	expression ast.ExpressionLike,
) ast.LetClauseLike {
	return ast.LetClauseClass().LetClause(
		recipient,
		assign,
		expression,
	)
}

// Ast/Logical

func Logical(
	any_ any,
) ast.LogicalLike {
	return ast.LogicalClass().Logical(
		any_,
	)
}

// Ast/Magnitude

func Magnitude(
	bar1 string,
	numerical ast.NumericalLike,
	bar2 string,
) ast.MagnitudeLike {
	return ast.MagnitudeClass().Magnitude(
		bar1,
		numerical,
		bar2,
	)
}

// Ast/MainClause

func MainClause(
	any_ any,
) ast.MainClauseLike {
	return ast.MainClauseClass().MainClause(
		any_,
	)
}

// Ast/MatchHandler

func MatchHandler(
	template ast.TemplateLike,
	procedure ast.ProcedureLike,
) ast.MatchHandlerLike {
	return ast.MatchHandlerClass().MatchHandler(
		template,
		procedure,
	)
}

// Ast/Message

func Message(
	expression ast.ExpressionLike,
) ast.MessageLike {
	return ast.MessageClass().Message(
		expression,
	)
}

// Ast/Messaging

func Messaging(
	any_ any,
) ast.MessagingLike {
	return ast.MessagingClass().Messaging(
		any_,
	)
}

// Ast/Method

func Method(
	identifier1 string,
	threading ast.ThreadingLike,
	identifier2 string,
	optionalArguments ast.ArgumentsLike,
) ast.MethodLike {
	return ast.MethodClass().Method(
		identifier1,
		threading,
		identifier2,
		optionalArguments,
	)
}

// Ast/MultilineAssociations

func MultilineAssociations(
	annotatedAssociations col.Sequential[ast.AnnotatedAssociationLike],
) ast.MultilineAssociationsLike {
	return ast.MultilineAssociationsClass().MultilineAssociations(
		annotatedAssociations,
	)
}

// Ast/MultilineStatements

func MultilineStatements(
	annotatedStatements col.Sequential[ast.AnnotatedStatementLike],
) ast.MultilineStatementsLike {
	return ast.MultilineStatementsClass().MultilineStatements(
		annotatedStatements,
	)
}

// Ast/MultilineValues

func MultilineValues(
	annotatedValues col.Sequential[ast.AnnotatedValueLike],
) ast.MultilineValuesLike {
	return ast.MultilineValuesClass().MultilineValues(
		annotatedValues,
	)
}

// Ast/NoAssociations

func NoAssociations(
	colon string,
) ast.NoAssociationsLike {
	return ast.NoAssociationsClass().NoAssociations(
		colon,
	)
}

// Ast/NotarizeClause

func NotarizeClause(
	draft ast.DraftLike,
	citation ast.CitationLike,
) ast.NotarizeClauseLike {
	return ast.NotarizeClauseClass().NotarizeClause(
		draft,
		citation,
	)
}

// Ast/Numerical

func Numerical(
	any_ any,
) ast.NumericalLike {
	return ast.NumericalClass().Numerical(
		any_,
	)
}

// Ast/OnClause

func OnClause(
	failure ast.FailureLike,
	matchHandlers col.Sequential[ast.MatchHandlerLike],
) ast.OnClauseLike {
	return ast.OnClauseClass().OnClause(
		failure,
		matchHandlers,
	)
}

// Ast/Operator

func Operator(
	any_ any,
) ast.OperatorLike {
	return ast.OperatorClass().Operator(
		any_,
	)
}

// Ast/Parameters

func Parameters(
	associations ast.AssociationsLike,
) ast.ParametersLike {
	return ast.ParametersClass().Parameters(
		associations,
	)
}

// Ast/PostClause

func PostClause(
	message ast.MessageLike,
	bag ast.BagLike,
) ast.PostClauseLike {
	return ast.PostClauseClass().PostClause(
		message,
		bag,
	)
}

// Ast/Precedence

func Precedence(
	expression ast.ExpressionLike,
) ast.PrecedenceLike {
	return ast.PrecedenceClass().Precedence(
		expression,
	)
}

// Ast/Predicate

func Predicate(
	operator ast.OperatorLike,
	expression ast.ExpressionLike,
) ast.PredicateLike {
	return ast.PredicateClass().Predicate(
		operator,
		expression,
	)
}

// Ast/Primitive

func Primitive(
	any_ any,
) ast.PrimitiveLike {
	return ast.PrimitiveClass().Primitive(
		any_,
	)
}

// Ast/Procedure

func Procedure(
	optionalStatements ast.StatementsLike,
) ast.ProcedureLike {
	return ast.ProcedureClass().Procedure(
		optionalStatements,
	)
}

// Ast/PublishClause

func PublishClause(
	event ast.EventLike,
) ast.PublishClauseLike {
	return ast.PublishClauseClass().PublishClause(
		event,
	)
}

// Ast/Range

func Range(
	inclusion1 ast.InclusionLike,
	primitive1 ast.PrimitiveLike,
	primitive2 ast.PrimitiveLike,
	inclusion2 ast.InclusionLike,
) ast.RangeLike {
	return ast.RangeClass().Range(
		inclusion1,
		primitive1,
		primitive2,
		inclusion2,
	)
}

// Ast/Recipient

func Recipient(
	any_ any,
) ast.RecipientLike {
	return ast.RecipientClass().Recipient(
		any_,
	)
}

// Ast/Referent

func Referent(
	snail string,
	indirect ast.IndirectLike,
) ast.ReferentLike {
	return ast.ReferentClass().Referent(
		snail,
		indirect,
	)
}

// Ast/RejectClause

func RejectClause(
	message ast.MessageLike,
) ast.RejectClauseLike {
	return ast.RejectClauseClass().RejectClause(
		message,
	)
}

// Ast/Repository

func Repository(
	any_ any,
) ast.RepositoryLike {
	return ast.RepositoryClass().Repository(
		any_,
	)
}

// Ast/Result

func Result(
	expression ast.ExpressionLike,
) ast.ResultLike {
	return ast.ResultClass().Result(
		expression,
	)
}

// Ast/RetrieveClause

func RetrieveClause(
	recipient ast.RecipientLike,
	bag ast.BagLike,
) ast.RetrieveClauseLike {
	return ast.RetrieveClauseClass().RetrieveClause(
		recipient,
		bag,
	)
}

// Ast/ReturnClause

func ReturnClause(
	result ast.ResultLike,
) ast.ReturnClauseLike {
	return ast.ReturnClauseClass().ReturnClause(
		result,
	)
}

// Ast/Right

func Right() ast.RightLike {
	return ast.RightClass().Right()
}

// Ast/SaveClause

func SaveClause(
	draft ast.DraftLike,
	citation ast.CitationLike,
) ast.SaveClauseLike {
	return ast.SaveClauseClass().SaveClause(
		draft,
		citation,
	)
}

// Ast/SelectClause

func SelectClause(
	target ast.TargetLike,
	matchHandlers col.Sequential[ast.MatchHandlerLike],
) ast.SelectClauseLike {
	return ast.SelectClauseClass().SelectClause(
		target,
		matchHandlers,
	)
}

// Ast/Sequence

func Sequence(
	expression ast.ExpressionLike,
) ast.SequenceLike {
	return ast.SequenceClass().Sequence(
		expression,
	)
}

// Ast/Statement

func Statement(
	mainClause ast.MainClauseLike,
	optionalOnClause ast.OnClauseLike,
) ast.StatementLike {
	return ast.StatementClass().Statement(
		mainClause,
		optionalOnClause,
	)
}

// Ast/Statements

func Statements(
	any_ any,
) ast.StatementsLike {
	return ast.StatementsClass().Statements(
		any_,
	)
}

// Ast/String

func String(
	any_ any,
) ast.StringLike {
	return ast.StringClass().String(
		any_,
	)
}

// Ast/Subcomponent

func Subcomponent(
	indices ast.IndicesLike,
) ast.SubcomponentLike {
	return ast.SubcomponentClass().Subcomponent(
		indices,
	)
}

// Ast/Subject

func Subject(
	any_ any,
) ast.SubjectLike {
	return ast.SubjectClass().Subject(
		any_,
	)
}

// Ast/Target

func Target(
	any_ any,
) ast.TargetLike {
	return ast.TargetClass().Target(
		any_,
	)
}

// Ast/Template

func Template(
	expression ast.ExpressionLike,
) ast.TemplateLike {
	return ast.TemplateClass().Template(
		expression,
	)
}

// Ast/Threading

func Threading(
	any_ any,
) ast.ThreadingLike {
	return ast.ThreadingClass().Threading(
		any_,
	)
}

// Ast/ThrowClause

func ThrowClause(
	exception ast.ExceptionLike,
) ast.ThrowClauseLike {
	return ast.ThrowClauseClass().ThrowClause(
		exception,
	)
}

// Ast/Value

func Value(
	component ast.ComponentLike,
) ast.ValueLike {
	return ast.ValueClass().Value(
		component,
	)
}

// Ast/Values

func Values(
	any_ any,
) ast.ValuesLike {
	return ast.ValuesClass().Values(
		any_,
	)
}

// Ast/Variable

func Variable(
	identifier string,
) ast.VariableLike {
	return ast.VariableClass().Variable(
		identifier,
	)
}

// Ast/WhileClause

func WhileClause(
	condition ast.ConditionLike,
	procedure ast.ProcedureLike,
) ast.WhileClauseLike {
	return ast.WhileClauseClass().WhileClause(
		condition,
		procedure,
	)
}

// Ast/WithClause

func WithClause(
	item ast.ItemLike,
	sequence ast.SequenceLike,
	procedure ast.ProcedureLike,
) ast.WithClauseLike {
	return ast.WithClauseClass().WithClause(
		item,
		sequence,
		procedure,
	)
}

// Grammar/Formatter

func Formatter() gra.FormatterLike {
	return gra.FormatterClass().Formatter()
}

// Grammar/Parser

func Parser() gra.ParserLike {
	return gra.ParserClass().Parser()
}

// Grammar/Processor

func Processor() gra.ProcessorLike {
	return gra.ProcessorClass().Processor()
}

// Grammar/Scanner

func Scanner(
	source string,
	tokens col.QueueLike[gra.TokenLike],
) gra.ScannerLike {
	return gra.ScannerClass().Scanner(
		source,
		tokens,
	)
}

// Grammar/Token

func Token(
	line uint,
	position uint,
	type_ gra.TokenType,
	value string,
) gra.TokenLike {
	return gra.TokenClass().Token(
		line,
		position,
		type_,
		value,
	)
}

// Grammar/Validator

func Validator() gra.ValidatorLike {
	return gra.ValidatorClass().Validator()
}

// Grammar/Visitor

func Visitor(
	processor gra.Methodical,
) gra.VisitorLike {
	return gra.VisitorClass().Visitor(
		processor,
	)
}

// GLOBAL FUNCTIONS
