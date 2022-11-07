/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// PROCEDURE INTERFACES

type ClauseLike any

type RecipientLike any

// This interface defines the methods supported by all accept-clause-like types.
type AcceptClauseLike interface {
	IsAcceptClause() bool
	GetMessage() ExpressionLike
	SetMessage(message ExpressionLike)
}

// This interface defines the methods supported by all attribute-like types.
type AttributeLike interface {
	GetVariable() string
	SetVariable(variable string)
	GetIndex(index int) ExpressionLike
	SetIndex(index int, expression ExpressionLike)
	GetIndices() Sequential[ExpressionLike]
	SetIndices(indices ListLike[ExpressionLike])
}

// This interface defines the methods supported by all block-like types.
type BlockLike interface {
	GetExpression() ExpressionLike
	SetExpression(expression ExpressionLike)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetProcedure() ProcedureLike
	SetProcedure(procedure ProcedureLike)
}

// This interface defines the methods supported by all break-clause-like types.
type BreakClauseLike interface {
	IsBreakClause() bool
}

// This interface defines the methods supported by all checkout-clause-like types.
type CheckoutClauseLike interface {
	IsCheckoutClause() bool
	GetRecipient() RecipientLike
	SetRecipient(recipient RecipientLike)
	GetLevel() ExpressionLike
	SetLevel(level ExpressionLike)
	GetMoniker() ExpressionLike
	SetMoniker(moniker ExpressionLike)
}

// This interface defines the methods supported by all continue-clause-like types.
type ContinueClauseLike interface {
	IsContinueClause() bool
}

// This interface defines the methods supported by all discard-clause-like types.
type DiscardClauseLike interface {
	IsDiscardClause() bool
	GetCitation() ExpressionLike
	SetCitation(citation ExpressionLike)
}

// This interface defines the methods supported by all evaluate-clause-like types.
type EvaluateClauseLike interface {
	IsEvaluateClause() bool
	HasRecipient() bool
	GetRecipient() (RecipientLike, Operator)
	SetRecipient(recipient RecipientLike, operator Operator)
	GetExpression() ExpressionLike
	SetExpression(expression ExpressionLike)
}

// This interface defines the methods supported by all if-clause-like types.
type IfClauseLike interface {
	IsIfClause() bool
	GetBlock() BlockLike
	SetBlock(block BlockLike)
	GetCondition() ExpressionLike
	SetCondition(condition ExpressionLike)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetProcedure() ProcedureLike
	SetProcedure(procedure ProcedureLike)
}

// This interface defines the methods supported by all notarize-clause-like types.
type NotarizeClauseLike interface {
	IsNotarizeClause() bool
	GetDocument() ExpressionLike
	SetDocument(document ExpressionLike)
	GetMoniker() ExpressionLike
	SetMoniker(moniker ExpressionLike)
}

// This interface defines the methods supported by all on-clause-like types.
type OnClauseLike interface {
	GetException() Symbolic
	SetException(exception Symbolic)
	GetBlock(index int) BlockLike
	SetBlock(index int, handler BlockLike)
	GetBlocks() Sequential[BlockLike]
	SetBlocks(blocks ListLike[BlockLike])
}

// This interface defines the methods supported by all post-clause-like types.
type PostClauseLike interface {
	IsPostClause() bool
	GetMessage() ExpressionLike
	SetMessage(message ExpressionLike)
	GetBag() ExpressionLike
	SetBag(bag ExpressionLike)
}

// This interface consolidates all the interfaces supported by procedure-like
// types.
type ProcedureLike interface {
	Sequential[StatementLike]
	Indexed[StatementLike]
	Updatable[StatementLike]
	Malleable[StatementLike]
}

// This interface defines the methods supported by all publish-clause-like types.
type PublishClauseLike interface {
	IsPublishClause() bool
	GetEvent() ExpressionLike
	SetEvent(event ExpressionLike)
}

// This interface defines the methods supported by all reject-clause-like types.
type RejectClauseLike interface {
	IsRejectClause() bool
	GetMessage() ExpressionLike
	SetMessage(message ExpressionLike)
}

// This interface defines the methods supported by all retrieve-clause-like types.
type RetrieveClauseLike interface {
	IsRetrieveClause() bool
	GetRecipient() RecipientLike
	SetRecipient(recipient RecipientLike)
	GetBag() ExpressionLike
	SetBag(bag ExpressionLike)
}

// This interface defines the methods supported by all return-clause-like types.
type ReturnClauseLike interface {
	IsReturnClause() bool
	GetResult() ExpressionLike
	SetResult(result ExpressionLike)
}

// This interface defines the methods supported by all save-clause-like types.
type SaveClauseLike interface {
	IsSaveClause() bool
	GetDocument() ExpressionLike
	SetDocument(document ExpressionLike)
	GetRecipient() RecipientLike
	SetRecipient(recipient RecipientLike)
}

// This interface defines the methods supported by all select-clause-like types.
type SelectClauseLike interface {
	IsSelectClause() bool
	GetTarget() ExpressionLike
	SetTarget(control ExpressionLike)
	GetBlock(index int) BlockLike
	SetBlock(index int, option BlockLike)
	GetBlocks() Sequential[BlockLike]
	SetBlocks(blocks ListLike[BlockLike])
}

// This interface defines the methods supported by all statement-like types.
type StatementLike interface {
	GetAnnotation() AnnotationLike
	SetAnnotation(annotation AnnotationLike)
	GetMainClause() ClauseLike
	SetMainClause(mainClause ClauseLike)
	GetNote() NoteLike
	SetNote(note NoteLike)
	GetOnClause() OnClauseLike
	SetOnClause(onClause OnClauseLike)
}

// This interface defines the methods supported by all throw-clause-like types.
type ThrowClauseLike interface {
	IsThrowClause() bool
	GetException() ExpressionLike
	SetException(exception ExpressionLike)
}

// This interface defines the methods supported by all while-clause-like types.
type WhileClauseLike interface {
	IsWhileClause() bool
	GetBlock() BlockLike
	SetBlock(block BlockLike)
	GetCondition() ExpressionLike
	SetCondition(condition ExpressionLike)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetProcedure() ProcedureLike
	SetProcedure(procedure ProcedureLike)
}

// This interface defines the methods supported by all with-clause-like types.
type WithClauseLike interface {
	IsWithClause() bool
	GetValue() Symbolic
	SetValue(exception Symbolic)
	GetBlock() BlockLike
	SetBlock(block BlockLike)
	GetSequence() ExpressionLike
	SetSequence(sequence ExpressionLike)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetProcedure() ProcedureLike
	SetProcedure(procedure ProcedureLike)
}
