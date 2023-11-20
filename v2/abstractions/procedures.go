/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// TYPE DEFINITIONS

type (
	Clause    any
	Recipient any
)

// INDIVIDUAL INTERFACES

type AcceptClauseLike interface {
	GetMessage() Expression
	SetMessage(message Expression)
}

type AttributeLike interface {
	GetVariable() string
	SetVariable(variable string)
	GetIndices() Sequential[Expression]
	SetIndices(indices Sequential[Expression])
}

type BlockLike interface {
	GetExpression() Expression
	SetExpression(expression Expression)
	GetProcedure() ProcedureLike
	SetProcedure(procedure ProcedureLike)
}

type BreakClauseLike interface {
}

type CheckoutClauseLike interface {
	GetRecipient() Recipient
	SetRecipient(recipient Recipient)
	GetLevel() Expression
	SetLevel(level Expression)
	GetName() Expression
	SetName(name Expression)
}

type ContinueClauseLike interface {
}

type DiscardClauseLike interface {
	GetDocument() Expression
	SetDocument(document Expression)
}

type IfClauseLike interface {
	GetBlock() BlockLike
	SetBlock(block BlockLike)
}

type LetClauseLike interface {
	HasRecipient() bool
	GetRecipient() (Recipient, Operator)
	SetRecipient(recipient Recipient, operator Operator)
	GetExpression() Expression
	SetExpression(expression Expression)
}

type NotarizeClauseLike interface {
	GetDocument() Expression
	SetDocument(document Expression)
	GetName() Expression
	SetName(name Expression)
}

type OnClauseLike interface {
	GetFailure() SymbolLike
	SetFailure(failure SymbolLike)
	GetBlocks() Sequential[BlockLike]
	SetBlocks(blocks Sequential[BlockLike])
}

type PostClauseLike interface {
	GetMessage() Expression
	SetMessage(message Expression)
	GetBag() Expression
	SetBag(bag Expression)
}

type PublishClauseLike interface {
	GetEvent() Expression
	SetEvent(event Expression)
}

type RejectClauseLike interface {
	GetMessage() Expression
	SetMessage(message Expression)
}

type RetrieveClauseLike interface {
	GetRecipient() Recipient
	SetRecipient(recipient Recipient)
	GetBag() Expression
	SetBag(bag Expression)
}

type ReturnClauseLike interface {
	GetResult() Expression
	SetResult(result Expression)
}

type SaveClauseLike interface {
	GetDocument() Expression
	SetDocument(document Expression)
	GetRecipient() Recipient
	SetRecipient(recipient Recipient)
}

type SelectClauseLike interface {
	GetTarget() Expression
	SetTarget(control Expression)
	GetBlocks() Sequential[BlockLike]
	SetBlocks(blocks Sequential[BlockLike])
}

type StatementLike interface {
	GetAnnotation() Annotation
	SetAnnotation(annotation Annotation)
	GetMainClause() Clause
	SetMainClause(mainClause Clause)
	GetOnClause() OnClauseLike
	SetOnClause(onClause OnClauseLike)
	GetNote() NoteLike
	SetNote(note NoteLike)
}

type ThrowClauseLike interface {
	GetException() Expression
	SetException(exception Expression)
}

type WhileClauseLike interface {
	GetBlock() BlockLike
	SetBlock(block BlockLike)
}

type WithClauseLike interface {
	GetItem() SymbolLike
	SetItem(item SymbolLike)
	GetBlock() BlockLike
	SetBlock(block BlockLike)
}

// CONSOLIDATED INTERFACES

type ProcedureLike interface {
	Sequential[StatementLike]
}
