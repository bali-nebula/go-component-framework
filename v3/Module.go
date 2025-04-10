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
	AdditionalParameterClassLike   = ast.AdditionalParameterClassLike
	AdditionalStatementClassLike   = ast.AdditionalStatementClassLike
	AdditionalValueClassLike       = ast.AdditionalValueClassLike
	AmplitudeClassLike             = ast.AmplitudeClassLike
	AndClassLike                   = ast.AndClassLike
	AnnotatedAssociationClassLike  = ast.AnnotatedAssociationClassLike
	AnnotatedParameterClassLike    = ast.AnnotatedParameterClassLike
	AnnotatedStatementClassLike    = ast.AnnotatedStatementClassLike
	AnnotatedValueClassLike        = ast.AnnotatedValueClassLike
	ArgumentClassLike              = ast.ArgumentClassLike
	ArgumentsClassLike             = ast.ArgumentsClassLike
	ArithmeticClassLike            = ast.ArithmeticClassLike
	ArithmeticOperatorClassLike    = ast.ArithmeticOperatorClassLike
	ArrowClassLike                 = ast.ArrowClassLike
	AssignmentOperatorClassLike    = ast.AssignmentOperatorClassLike
	AssociationClassLike           = ast.AssociationClassLike
	AssociationsClassLike          = ast.AssociationsClassLike
	AtLevelClassLike               = ast.AtLevelClassLike
	AttributeClassLike             = ast.AttributeClassLike
	BagClassLike                   = ast.BagClassLike
	BreakClauseClassLike           = ast.BreakClauseClassLike
	ChainingClassLike              = ast.ChainingClassLike
	CheckoutClauseClassLike        = ast.CheckoutClauseClassLike
	CitationClassLike              = ast.CitationClassLike
	CollectionClassLike            = ast.CollectionClassLike
	ColonEqualsClassLike           = ast.ColonEqualsClassLike
	ComparisonClassLike            = ast.ComparisonClassLike
	ComparisonOperatorClassLike    = ast.ComparisonOperatorClassLike
	ComplementClassLike            = ast.ComplementClassLike
	ComponentClassLike             = ast.ComponentClassLike
	CompositeClassLike             = ast.CompositeClassLike
	ConditionClassLike             = ast.ConditionClassLike
	ContextClassLike               = ast.ContextClassLike
	ContinueClauseClassLike        = ast.ContinueClauseClassLike
	DefaultEqualsClassLike         = ast.DefaultEqualsClassLike
	DereferenceClassLike           = ast.DereferenceClassLike
	DiscardClauseClassLike         = ast.DiscardClauseClassLike
	DivideClassLike                = ast.DivideClassLike
	DivideEqualsClassLike          = ast.DivideEqualsClassLike
	DoClauseClassLike              = ast.DoClauseClassLike
	DocumentClassLike              = ast.DocumentClassLike
	DotClassLike                   = ast.DotClassLike
	DraftClassLike                 = ast.DraftClassLike
	ElementClassLike               = ast.ElementClassLike
	EntityClassLike                = ast.EntityClassLike
	EqualsClassLike                = ast.EqualsClassLike
	EventClassLike                 = ast.EventClassLike
	ExceptionClassLike             = ast.ExceptionClassLike
	ExponentialClassLike           = ast.ExponentialClassLike
	ExpressionClassLike            = ast.ExpressionClassLike
	FailureClassLike               = ast.FailureClassLike
	FlowClassLike                  = ast.FlowClassLike
	FunctionClassLike              = ast.FunctionClassLike
	IORClassLike                   = ast.IORClassLike
	IfClauseClassLike              = ast.IfClauseClassLike
	IndexClassLike                 = ast.IndexClassLike
	IndicesClassLike               = ast.IndicesClassLike
	InductionClassLike             = ast.InductionClassLike
	InlineAssociationsClassLike    = ast.InlineAssociationsClassLike
	InlineParametersClassLike      = ast.InlineParametersClassLike
	InlineStatementsClassLike      = ast.InlineStatementsClassLike
	InlineValuesClassLike          = ast.InlineValuesClassLike
	IntrinsicClassLike             = ast.IntrinsicClassLike
	InversionClassLike             = ast.InversionClassLike
	InversionOperatorClassLike     = ast.InversionOperatorClassLike
	InvocationClassLike            = ast.InvocationClassLike
	InvocationOperatorClassLike    = ast.InvocationOperatorClassLike
	IsClassLike                    = ast.IsClassLike
	ItemClassLike                  = ast.ItemClassLike
	ItemsClassLike                 = ast.ItemsClassLike
	KeyClassLike                   = ast.KeyClassLike
	LabelClassLike                 = ast.LabelClassLike
	LeftBracketClassLike           = ast.LeftBracketClassLike
	LeftRoundClassLike             = ast.LeftRoundClassLike
	LeftSquareClassLike            = ast.LeftSquareClassLike
	LessClassLike                  = ast.LessClassLike
	LetClauseClassLike             = ast.LetClauseClassLike
	LogicalClassLike               = ast.LogicalClassLike
	LogicalOperatorClassLike       = ast.LogicalOperatorClassLike
	MainClauseClassLike            = ast.MainClauseClassLike
	MatchHandlerClassLike          = ast.MatchHandlerClassLike
	MatchesClassLike               = ast.MatchesClassLike
	MessageClassLike               = ast.MessageClassLike
	MessagingClassLike             = ast.MessagingClassLike
	MethodClassLike                = ast.MethodClassLike
	MinusClassLike                 = ast.MinusClassLike
	MinusEqualsClassLike           = ast.MinusEqualsClassLike
	ModuloClassLike                = ast.ModuloClassLike
	MoreClassLike                  = ast.MoreClassLike
	MultilineAssociationsClassLike = ast.MultilineAssociationsClassLike
	MultilineParametersClassLike   = ast.MultilineParametersClassLike
	MultilineStatementsClassLike   = ast.MultilineStatementsClassLike
	MultilineValuesClassLike       = ast.MultilineValuesClassLike
	NoAssociationsClassLike        = ast.NoAssociationsClassLike
	NotarizeClauseClassLike        = ast.NotarizeClauseClassLike
	OnClauseClassLike              = ast.OnClauseClassLike
	OperationClassLike             = ast.OperationClassLike
	ParameterClassLike             = ast.ParameterClassLike
	ParametersClassLike            = ast.ParametersClassLike
	PlusClassLike                  = ast.PlusClassLike
	PlusEqualsClassLike            = ast.PlusEqualsClassLike
	PostClauseClassLike            = ast.PostClauseClassLike
	PrecedenceClassLike            = ast.PrecedenceClassLike
	PrimitiveClassLike             = ast.PrimitiveClassLike
	ProcedureClassLike             = ast.ProcedureClassLike
	PublishClauseClassLike         = ast.PublishClauseClassLike
	RangeClassLike                 = ast.RangeClassLike
	RecipientClassLike             = ast.RecipientClassLike
	RejectClauseClassLike          = ast.RejectClauseClassLike
	RepositoryClassLike            = ast.RepositoryClassLike
	ResultClassLike                = ast.ResultClassLike
	RetrieveClauseClassLike        = ast.RetrieveClauseClassLike
	ReturnClauseClassLike          = ast.ReturnClauseClassLike
	RightBracketClassLike          = ast.RightBracketClassLike
	RightRoundClassLike            = ast.RightRoundClassLike
	RightSquareClassLike           = ast.RightSquareClassLike
	SANClassLike                   = ast.SANClassLike
	SaveClauseClassLike            = ast.SaveClauseClassLike
	SelectClauseClassLike          = ast.SelectClauseClassLike
	SequenceClassLike              = ast.SequenceClassLike
	StatementClassLike             = ast.StatementClassLike
	StatementsClassLike            = ast.StatementsClassLike
	StringClassLike                = ast.StringClassLike
	SubcomponentClassLike          = ast.SubcomponentClassLike
	TargetClassLike                = ast.TargetClassLike
	TemplateClassLike              = ast.TemplateClassLike
	ThrowClauseClassLike           = ast.ThrowClauseClassLike
	TimesClassLike                 = ast.TimesClassLike
	TimesEqualsClassLike           = ast.TimesEqualsClassLike
	ValueClassLike                 = ast.ValueClassLike
	ValuesClassLike                = ast.ValuesClassLike
	VariableClassLike              = ast.VariableClassLike
	WhileClauseClassLike           = ast.WhileClauseClassLike
	WithClauseClassLike            = ast.WithClauseClassLike
	XORClassLike                   = ast.XORClassLike
)

type (
	AcceptClauseLike          = ast.AcceptClauseLike
	AdditionalArgumentLike    = ast.AdditionalArgumentLike
	AdditionalAssociationLike = ast.AdditionalAssociationLike
	AdditionalIndexLike       = ast.AdditionalIndexLike
	AdditionalParameterLike   = ast.AdditionalParameterLike
	AdditionalStatementLike   = ast.AdditionalStatementLike
	AdditionalValueLike       = ast.AdditionalValueLike
	AmplitudeLike             = ast.AmplitudeLike
	AndLike                   = ast.AndLike
	AnnotatedAssociationLike  = ast.AnnotatedAssociationLike
	AnnotatedParameterLike    = ast.AnnotatedParameterLike
	AnnotatedStatementLike    = ast.AnnotatedStatementLike
	AnnotatedValueLike        = ast.AnnotatedValueLike
	ArgumentLike              = ast.ArgumentLike
	ArgumentsLike             = ast.ArgumentsLike
	ArithmeticLike            = ast.ArithmeticLike
	ArithmeticOperatorLike    = ast.ArithmeticOperatorLike
	ArrowLike                 = ast.ArrowLike
	AssignmentOperatorLike    = ast.AssignmentOperatorLike
	AssociationLike           = ast.AssociationLike
	AssociationsLike          = ast.AssociationsLike
	AtLevelLike               = ast.AtLevelLike
	AttributeLike             = ast.AttributeLike
	BagLike                   = ast.BagLike
	BreakClauseLike           = ast.BreakClauseLike
	ChainingLike              = ast.ChainingLike
	CheckoutClauseLike        = ast.CheckoutClauseLike
	CitationLike              = ast.CitationLike
	CollectionLike            = ast.CollectionLike
	ColonEqualsLike           = ast.ColonEqualsLike
	ComparisonLike            = ast.ComparisonLike
	ComparisonOperatorLike    = ast.ComparisonOperatorLike
	ComplementLike            = ast.ComplementLike
	ComponentLike             = ast.ComponentLike
	CompositeLike             = ast.CompositeLike
	ConditionLike             = ast.ConditionLike
	ContextLike               = ast.ContextLike
	ContinueClauseLike        = ast.ContinueClauseLike
	DefaultEqualsLike         = ast.DefaultEqualsLike
	DereferenceLike           = ast.DereferenceLike
	DiscardClauseLike         = ast.DiscardClauseLike
	DivideLike                = ast.DivideLike
	DivideEqualsLike          = ast.DivideEqualsLike
	DoClauseLike              = ast.DoClauseLike
	DocumentLike              = ast.DocumentLike
	DotLike                   = ast.DotLike
	DraftLike                 = ast.DraftLike
	ElementLike               = ast.ElementLike
	EntityLike                = ast.EntityLike
	EqualsLike                = ast.EqualsLike
	EventLike                 = ast.EventLike
	ExceptionLike             = ast.ExceptionLike
	ExponentialLike           = ast.ExponentialLike
	ExpressionLike            = ast.ExpressionLike
	FailureLike               = ast.FailureLike
	FlowLike                  = ast.FlowLike
	FunctionLike              = ast.FunctionLike
	IORLike                   = ast.IORLike
	IfClauseLike              = ast.IfClauseLike
	IndexLike                 = ast.IndexLike
	IndicesLike               = ast.IndicesLike
	InductionLike             = ast.InductionLike
	InlineAssociationsLike    = ast.InlineAssociationsLike
	InlineParametersLike      = ast.InlineParametersLike
	InlineStatementsLike      = ast.InlineStatementsLike
	InlineValuesLike          = ast.InlineValuesLike
	IntrinsicLike             = ast.IntrinsicLike
	InversionLike             = ast.InversionLike
	InversionOperatorLike     = ast.InversionOperatorLike
	InvocationLike            = ast.InvocationLike
	InvocationOperatorLike    = ast.InvocationOperatorLike
	IsLike                    = ast.IsLike
	ItemLike                  = ast.ItemLike
	ItemsLike                 = ast.ItemsLike
	KeyLike                   = ast.KeyLike
	LabelLike                 = ast.LabelLike
	LeftBracketLike           = ast.LeftBracketLike
	LeftRoundLike             = ast.LeftRoundLike
	LeftSquareLike            = ast.LeftSquareLike
	LessLike                  = ast.LessLike
	LetClauseLike             = ast.LetClauseLike
	LogicalLike               = ast.LogicalLike
	LogicalOperatorLike       = ast.LogicalOperatorLike
	MainClauseLike            = ast.MainClauseLike
	MatchHandlerLike          = ast.MatchHandlerLike
	MatchesLike               = ast.MatchesLike
	MessageLike               = ast.MessageLike
	MessagingLike             = ast.MessagingLike
	MethodLike                = ast.MethodLike
	MinusLike                 = ast.MinusLike
	MinusEqualsLike           = ast.MinusEqualsLike
	ModuloLike                = ast.ModuloLike
	MoreLike                  = ast.MoreLike
	MultilineAssociationsLike = ast.MultilineAssociationsLike
	MultilineParametersLike   = ast.MultilineParametersLike
	MultilineStatementsLike   = ast.MultilineStatementsLike
	MultilineValuesLike       = ast.MultilineValuesLike
	NoAssociationsLike        = ast.NoAssociationsLike
	NotarizeClauseLike        = ast.NotarizeClauseLike
	OnClauseLike              = ast.OnClauseLike
	OperationLike             = ast.OperationLike
	ParameterLike             = ast.ParameterLike
	ParametersLike            = ast.ParametersLike
	PlusLike                  = ast.PlusLike
	PlusEqualsLike            = ast.PlusEqualsLike
	PostClauseLike            = ast.PostClauseLike
	PrecedenceLike            = ast.PrecedenceLike
	PrimitiveLike             = ast.PrimitiveLike
	ProcedureLike             = ast.ProcedureLike
	PublishClauseLike         = ast.PublishClauseLike
	RangeLike                 = ast.RangeLike
	RecipientLike             = ast.RecipientLike
	RejectClauseLike          = ast.RejectClauseLike
	RepositoryLike            = ast.RepositoryLike
	ResultLike                = ast.ResultLike
	RetrieveClauseLike        = ast.RetrieveClauseLike
	ReturnClauseLike          = ast.ReturnClauseLike
	RightBracketLike          = ast.RightBracketLike
	RightRoundLike            = ast.RightRoundLike
	RightSquareLike           = ast.RightSquareLike
	SANLike                   = ast.SANLike
	SaveClauseLike            = ast.SaveClauseLike
	SelectClauseLike          = ast.SelectClauseLike
	SequenceLike              = ast.SequenceLike
	StatementLike             = ast.StatementLike
	StatementsLike            = ast.StatementsLike
	StringLike                = ast.StringLike
	SubcomponentLike          = ast.SubcomponentLike
	TargetLike                = ast.TargetLike
	TemplateLike              = ast.TemplateLike
	ThrowClauseLike           = ast.ThrowClauseLike
	TimesLike                 = ast.TimesLike
	TimesEqualsLike           = ast.TimesEqualsLike
	ValueLike                 = ast.ValueLike
	ValuesLike                = ast.ValuesLike
	VariableLike              = ast.VariableLike
	WhileClauseLike           = ast.WhileClauseLike
	WithClauseLike            = ast.WithClauseLike
	XORLike                   = ast.XORLike
)

// Grammar

type (
	TokenType = gra.TokenType
)

const (
	ErrorToken       = gra.ErrorToken
	AngleToken       = gra.AngleToken
	BinaryToken      = gra.BinaryToken
	BooleanToken     = gra.BooleanToken
	BytecodeToken    = gra.BytecodeToken
	CommentToken     = gra.CommentToken
	DelimiterToken   = gra.DelimiterToken
	DurationToken    = gra.DurationToken
	IdentifierToken  = gra.IdentifierToken
	MomentToken      = gra.MomentToken
	NameToken        = gra.NameToken
	NarrativeToken   = gra.NarrativeToken
	NewlineToken     = gra.NewlineToken
	NoteToken        = gra.NoteToken
	NumberToken      = gra.NumberToken
	PatternToken     = gra.PatternToken
	PercentageToken  = gra.PercentageToken
	ProbabilityToken = gra.ProbabilityToken
	QuoteToken       = gra.QuoteToken
	ResourceToken    = gra.ResourceToken
	SpaceToken       = gra.SpaceToken
	SymbolToken      = gra.SymbolToken
	TagToken         = gra.TagToken
	VersionToken     = gra.VersionToken
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

// Ast/AdditionalParameter

func AdditionalParameter(
	parameter ast.ParameterLike,
) ast.AdditionalParameterLike {
	return ast.AdditionalParameterClass().AdditionalParameter(
		parameter,
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

// Ast/Amplitude

func Amplitude(
	expression ast.ExpressionLike,
) ast.AmplitudeLike {
	return ast.AmplitudeClass().Amplitude(
		expression,
	)
}

// Ast/And

func And() ast.AndLike {
	return ast.AndClass().And()
}

// Ast/AnnotatedAssociation

func AnnotatedAssociation(
	association ast.AssociationLike,
	optionalNote string,
) ast.AnnotatedAssociationLike {
	return ast.AnnotatedAssociationClass().AnnotatedAssociation(
		association,
		optionalNote,
	)
}

// Ast/AnnotatedParameter

func AnnotatedParameter(
	parameter ast.ParameterLike,
	optionalNote string,
) ast.AnnotatedParameterLike {
	return ast.AnnotatedParameterClass().AnnotatedParameter(
		parameter,
		optionalNote,
	)
}

// Ast/AnnotatedStatement

func AnnotatedStatement(
	statement ast.StatementLike,
	optionalNote string,
) ast.AnnotatedStatementLike {
	return ast.AnnotatedStatementClass().AnnotatedStatement(
		statement,
		optionalNote,
	)
}

// Ast/AnnotatedValue

func AnnotatedValue(
	value ast.ValueLike,
	optionalNote string,
) ast.AnnotatedValueLike {
	return ast.AnnotatedValueClass().AnnotatedValue(
		value,
		optionalNote,
	)
}

// Ast/Argument

func Argument(
	expression ast.ExpressionLike,
) ast.ArgumentLike {
	return ast.ArgumentClass().Argument(
		expression,
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

// Ast/Arithmetic

func Arithmetic(
	expression1 ast.ExpressionLike,
	arithmeticOperator ast.ArithmeticOperatorLike,
	expression2 ast.ExpressionLike,
) ast.ArithmeticLike {
	return ast.ArithmeticClass().Arithmetic(
		expression1,
		arithmeticOperator,
		expression2,
	)
}

// Ast/ArithmeticOperator

func ArithmeticOperator(
	any_ any,
) ast.ArithmeticOperatorLike {
	return ast.ArithmeticOperatorClass().ArithmeticOperator(
		any_,
	)
}

// Ast/Arrow

func Arrow() ast.ArrowLike {
	return ast.ArrowClass().Arrow()
}

// Ast/AssignmentOperator

func AssignmentOperator(
	any_ any,
) ast.AssignmentOperatorLike {
	return ast.AssignmentOperatorClass().AssignmentOperator(
		any_,
	)
}

// Ast/Association

func Association(
	key ast.KeyLike,
	value ast.ValueLike,
) ast.AssociationLike {
	return ast.AssociationClass().Association(
		key,
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
	variable ast.VariableLike,
	indices ast.IndicesLike,
) ast.AttributeLike {
	return ast.AttributeClass().Attribute(
		variable,
		indices,
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

// Ast/Chaining

func Chaining(
	expression1 ast.ExpressionLike,
	expression2 ast.ExpressionLike,
) ast.ChainingLike {
	return ast.ChainingClass().Chaining(
		expression1,
		expression2,
	)
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

// Ast/ColonEquals

func ColonEquals() ast.ColonEqualsLike {
	return ast.ColonEqualsClass().ColonEquals()
}

// Ast/Comparison

func Comparison(
	expression1 ast.ExpressionLike,
	comparisonOperator ast.ComparisonOperatorLike,
	expression2 ast.ExpressionLike,
) ast.ComparisonLike {
	return ast.ComparisonClass().Comparison(
		expression1,
		comparisonOperator,
		expression2,
	)
}

// Ast/ComparisonOperator

func ComparisonOperator(
	any_ any,
) ast.ComparisonOperatorLike {
	return ast.ComparisonOperatorClass().ComparisonOperator(
		any_,
	)
}

// Ast/Complement

func Complement(
	expression ast.ExpressionLike,
) ast.ComplementLike {
	return ast.ComplementClass().Complement(
		expression,
	)
}

// Ast/Component

func Component(
	entity ast.EntityLike,
	optionalContext ast.ContextLike,
) ast.ComponentLike {
	return ast.ComponentClass().Component(
		entity,
		optionalContext,
	)
}

// Ast/Composite

func Composite(
	expression ast.ExpressionLike,
) ast.CompositeLike {
	return ast.CompositeClass().Composite(
		expression,
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

// Ast/Context

func Context(
	parameters ast.ParametersLike,
) ast.ContextLike {
	return ast.ContextClass().Context(
		parameters,
	)
}

// Ast/ContinueClause

func ContinueClause() ast.ContinueClauseLike {
	return ast.ContinueClauseClass().ContinueClause()
}

// Ast/DefaultEquals

func DefaultEquals() ast.DefaultEqualsLike {
	return ast.DefaultEqualsClass().DefaultEquals()
}

// Ast/Dereference

func Dereference(
	expression ast.ExpressionLike,
) ast.DereferenceLike {
	return ast.DereferenceClass().Dereference(
		expression,
	)
}

// Ast/DiscardClause

func DiscardClause(
	draft ast.DraftLike,
) ast.DiscardClauseLike {
	return ast.DiscardClauseClass().DiscardClause(
		draft,
	)
}

// Ast/Divide

func Divide() ast.DivideLike {
	return ast.DivideClass().Divide()
}

// Ast/DivideEquals

func DivideEquals() ast.DivideEqualsLike {
	return ast.DivideEqualsClass().DivideEquals()
}

// Ast/DoClause

func DoClause(
	operation ast.OperationLike,
) ast.DoClauseLike {
	return ast.DoClauseClass().DoClause(
		operation,
	)
}

// Ast/Document

func Document(
	comment string,
	component ast.ComponentLike,
) ast.DocumentLike {
	return ast.DocumentClass().Document(
		comment,
		component,
	)
}

// Ast/Dot

func Dot() ast.DotLike {
	return ast.DotClass().Dot()
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

// Ast/Equals

func Equals() ast.EqualsLike {
	return ast.EqualsClass().Equals()
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

// Ast/Exponential

func Exponential(
	expression1 ast.ExpressionLike,
	expression2 ast.ExpressionLike,
) ast.ExponentialLike {
	return ast.ExponentialClass().Exponential(
		expression1,
		expression2,
	)
}

// Ast/Expression

func Expression(
	any_ any,
) ast.ExpressionLike {
	return ast.ExpressionClass().Expression(
		any_,
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
) ast.FunctionLike {
	return ast.FunctionClass().Function(
		identifier,
	)
}

// Ast/IOR

func IOR() ast.IORLike {
	return ast.IORClass().IOR()
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

// Ast/InlineParameters

func InlineParameters(
	parameter ast.ParameterLike,
	additionalParameters col.Sequential[ast.AdditionalParameterLike],
) ast.InlineParametersLike {
	return ast.InlineParametersClass().InlineParameters(
		parameter,
		additionalParameters,
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

// Ast/Intrinsic

func Intrinsic(
	function ast.FunctionLike,
	optionalArguments ast.ArgumentsLike,
) ast.IntrinsicLike {
	return ast.IntrinsicClass().Intrinsic(
		function,
		optionalArguments,
	)
}

// Ast/Inversion

func Inversion(
	inversionOperator ast.InversionOperatorLike,
	expression ast.ExpressionLike,
) ast.InversionLike {
	return ast.InversionClass().Inversion(
		inversionOperator,
		expression,
	)
}

// Ast/InversionOperator

func InversionOperator(
	any_ any,
) ast.InversionOperatorLike {
	return ast.InversionOperatorClass().InversionOperator(
		any_,
	)
}

// Ast/Invocation

func Invocation(
	target ast.TargetLike,
	invocationOperator ast.InvocationOperatorLike,
	method ast.MethodLike,
	optionalArguments ast.ArgumentsLike,
) ast.InvocationLike {
	return ast.InvocationClass().Invocation(
		target,
		invocationOperator,
		method,
		optionalArguments,
	)
}

// Ast/InvocationOperator

func InvocationOperator(
	any_ any,
) ast.InvocationOperatorLike {
	return ast.InvocationOperatorClass().InvocationOperator(
		any_,
	)
}

// Ast/Is

func Is() ast.IsLike {
	return ast.IsClass().Is()
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

// Ast/Key

func Key(
	primitive ast.PrimitiveLike,
) ast.KeyLike {
	return ast.KeyClass().Key(
		primitive,
	)
}

// Ast/Label

func Label(
	symbol string,
) ast.LabelLike {
	return ast.LabelClass().Label(
		symbol,
	)
}

// Ast/LeftBracket

func LeftBracket(
	any_ any,
) ast.LeftBracketLike {
	return ast.LeftBracketClass().LeftBracket(
		any_,
	)
}

// Ast/LeftRound

func LeftRound() ast.LeftRoundLike {
	return ast.LeftRoundClass().LeftRound()
}

// Ast/LeftSquare

func LeftSquare() ast.LeftSquareLike {
	return ast.LeftSquareClass().LeftSquare()
}

// Ast/Less

func Less() ast.LessLike {
	return ast.LessClass().Less()
}

// Ast/LetClause

func LetClause(
	recipient ast.RecipientLike,
	assignmentOperator ast.AssignmentOperatorLike,
	expression ast.ExpressionLike,
) ast.LetClauseLike {
	return ast.LetClauseClass().LetClause(
		recipient,
		assignmentOperator,
		expression,
	)
}

// Ast/Logical

func Logical(
	expression1 ast.ExpressionLike,
	logicalOperator ast.LogicalOperatorLike,
	expression2 ast.ExpressionLike,
) ast.LogicalLike {
	return ast.LogicalClass().Logical(
		expression1,
		logicalOperator,
		expression2,
	)
}

// Ast/LogicalOperator

func LogicalOperator(
	any_ any,
) ast.LogicalOperatorLike {
	return ast.LogicalOperatorClass().LogicalOperator(
		any_,
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

// Ast/Matches

func Matches() ast.MatchesLike {
	return ast.MatchesClass().Matches()
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
	identifier string,
) ast.MethodLike {
	return ast.MethodClass().Method(
		identifier,
	)
}

// Ast/Minus

func Minus() ast.MinusLike {
	return ast.MinusClass().Minus()
}

// Ast/MinusEquals

func MinusEquals() ast.MinusEqualsLike {
	return ast.MinusEqualsClass().MinusEquals()
}

// Ast/Modulo

func Modulo() ast.ModuloLike {
	return ast.ModuloClass().Modulo()
}

// Ast/More

func More() ast.MoreLike {
	return ast.MoreClass().More()
}

// Ast/MultilineAssociations

func MultilineAssociations(
	annotatedAssociations col.Sequential[ast.AnnotatedAssociationLike],
) ast.MultilineAssociationsLike {
	return ast.MultilineAssociationsClass().MultilineAssociations(
		annotatedAssociations,
	)
}

// Ast/MultilineParameters

func MultilineParameters(
	annotatedParameters col.Sequential[ast.AnnotatedParameterLike],
) ast.MultilineParametersLike {
	return ast.MultilineParametersClass().MultilineParameters(
		annotatedParameters,
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

func NoAssociations() ast.NoAssociationsLike {
	return ast.NoAssociationsClass().NoAssociations()
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

// Ast/Operation

func Operation(
	any_ any,
) ast.OperationLike {
	return ast.OperationClass().Operation(
		any_,
	)
}

// Ast/Parameter

func Parameter(
	label ast.LabelLike,
	component ast.ComponentLike,
) ast.ParameterLike {
	return ast.ParameterClass().Parameter(
		label,
		component,
	)
}

// Ast/Parameters

func Parameters(
	any_ any,
) ast.ParametersLike {
	return ast.ParametersClass().Parameters(
		any_,
	)
}

// Ast/Plus

func Plus() ast.PlusLike {
	return ast.PlusClass().Plus()
}

// Ast/PlusEquals

func PlusEquals() ast.PlusEqualsLike {
	return ast.PlusEqualsClass().PlusEquals()
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
	leftBracket ast.LeftBracketLike,
	primitive1 ast.PrimitiveLike,
	primitive2 ast.PrimitiveLike,
	rightBracket ast.RightBracketLike,
) ast.RangeLike {
	return ast.RangeClass().Range(
		leftBracket,
		primitive1,
		primitive2,
		rightBracket,
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

// Ast/RightBracket

func RightBracket(
	any_ any,
) ast.RightBracketLike {
	return ast.RightBracketClass().RightBracket(
		any_,
	)
}

// Ast/RightRound

func RightRound() ast.RightRoundLike {
	return ast.RightRoundClass().RightRound()
}

// Ast/RightSquare

func RightSquare() ast.RightSquareLike {
	return ast.RightSquareClass().RightSquare()
}

// Ast/SAN

func SAN() ast.SANLike {
	return ast.SANClass().SAN()
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
	composite ast.CompositeLike,
	indices ast.IndicesLike,
) ast.SubcomponentLike {
	return ast.SubcomponentClass().Subcomponent(
		composite,
		indices,
	)
}

// Ast/Target

func Target(
	expression ast.ExpressionLike,
) ast.TargetLike {
	return ast.TargetClass().Target(
		expression,
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

// Ast/ThrowClause

func ThrowClause(
	exception ast.ExceptionLike,
) ast.ThrowClauseLike {
	return ast.ThrowClauseClass().ThrowClause(
		exception,
	)
}

// Ast/Times

func Times() ast.TimesLike {
	return ast.TimesClass().Times()
}

// Ast/TimesEquals

func TimesEquals() ast.TimesEqualsLike {
	return ast.TimesEqualsClass().TimesEquals()
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

// Ast/XOR

func XOR() ast.XORLike {
	return ast.XORClass().XOR()
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
