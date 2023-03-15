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

// This interface defines the methods supported by all accept-clause-like types.
type AcceptClauseLike interface {
	GetMessage() Expression
	SetMessage(message Expression)
}

// This interface defines the methods supported by all attribute-like types.
type AttributeLike interface {
	GetVariable() string
	SetVariable(variable string)
	GetIndices() Sequential[Expression]
	SetIndices(indices Sequential[Expression])
}

// This interface defines the methods supported by all block-like types.
type BlockLike interface {
	GetExpression() Expression
	SetExpression(expression Expression)
	GetProcedure() ProcedureLike
	SetProcedure(procedure ProcedureLike)
}

// This interface defines the methods supported by all break-clause-like types.
type BreakClauseLike interface {
}

// This interface defines the methods supported by all checkout-clause-like types.
type CheckoutClauseLike interface {
	GetRecipient() Recipient
	SetRecipient(recipient Recipient)
	GetLevel() Expression
	SetLevel(level Expression)
	GetMoniker() Expression
	SetMoniker(moniker Expression)
}

// This interface defines the methods supported by all continue-clause-like types.
type ContinueClauseLike interface {
}

// This interface defines the methods supported by all discard-clause-like types.
type DiscardClauseLike interface {
	GetDocument() Expression
	SetDocument(document Expression)
}

// This interface defines the methods supported by all if-clause-like types.
type IfClauseLike interface {
	GetBlock() BlockLike
	SetBlock(block BlockLike)
}

// This interface defines the methods supported by all let-clause-like types.
type LetClauseLike interface {
	HasRecipient() bool
	GetRecipient() (Recipient, Operator)
	SetRecipient(recipient Recipient, operator Operator)
	GetExpression() Expression
	SetExpression(expression Expression)
}

// This interface defines the methods supported by all notarize-clause-like types.
type NotarizeClauseLike interface {
	GetDocument() Expression
	SetDocument(document Expression)
	GetMoniker() Expression
	SetMoniker(moniker Expression)
}

// This interface defines the methods supported by all on-clause-like types.
type OnClauseLike interface {
	GetFailure() SymbolLike
	SetFailure(failure SymbolLike)
	GetBlocks() Sequential[BlockLike]
	SetBlocks(blocks Sequential[BlockLike])
}

// This interface defines the methods supported by all post-clause-like types.
type PostClauseLike interface {
	GetMessage() Expression
	SetMessage(message Expression)
	GetBag() Expression
	SetBag(bag Expression)
}

// This interface defines the methods supported by all publish-clause-like types.
type PublishClauseLike interface {
	GetEvent() Expression
	SetEvent(event Expression)
}

// This interface defines the methods supported by all reject-clause-like types.
type RejectClauseLike interface {
	GetMessage() Expression
	SetMessage(message Expression)
}

// This interface defines the methods supported by all retrieve-clause-like types.
type RetrieveClauseLike interface {
	GetRecipient() Recipient
	SetRecipient(recipient Recipient)
	GetBag() Expression
	SetBag(bag Expression)
}

// This interface defines the methods supported by all return-clause-like types.
type ReturnClauseLike interface {
	GetResult() Expression
	SetResult(result Expression)
}

// This interface defines the methods supported by all save-clause-like types.
type SaveClauseLike interface {
	GetDocument() Expression
	SetDocument(document Expression)
	GetRecipient() Recipient
	SetRecipient(recipient Recipient)
}

// This interface defines the methods supported by all select-clause-like types.
type SelectClauseLike interface {
	GetTarget() Expression
	SetTarget(control Expression)
	GetBlocks() Sequential[BlockLike]
	SetBlocks(blocks Sequential[BlockLike])
}

// This interface defines the methods supported by all statement-like types.
type StatementLike interface {
	GetAnnotation() Annotation
	SetAnnotation(annotation Annotation)
	GetMainClause() Clause
	SetMainClause(mainClause Clause)
	GetNote() NoteLike
	SetNote(note NoteLike)
	GetOnClause() OnClauseLike
	SetOnClause(onClause OnClauseLike)
}

// This interface defines the methods supported by all throw-clause-like types.
type ThrowClauseLike interface {
	GetException() Expression
	SetException(exception Expression)
}

// This interface defines the methods supported by all while-clause-like types.
type WhileClauseLike interface {
	GetBlock() BlockLike
	SetBlock(block BlockLike)
}

// This interface defines the methods supported by all with-clause-like types.
type WithClauseLike interface {
	GetItem() SymbolLike
	SetItem(item SymbolLike)
	GetBlock() BlockLike
	SetBlock(block BlockLike)
}

// CONSOLIDATED INTERFACES

// This interface consolidates all the interfaces supported by procedure-like
// entities.
type ProcedureLike interface {
	Sequential[StatementLike]
}
