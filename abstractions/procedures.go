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

// ASSIGNMENT CONSTANTS

type Assignment int

const (
	_ Assignment = iota
	REGULAR
	DEFAULT
	TIMES
	DIVIDE
	PLUS
	MINUS
)

// PROCEDURE INTERFACES

type Clause any

type Recipient any

// This interface defines the methods supported by all accept-clause-like types.
type AcceptClauseLike interface {
	GetMessage() Expression
	SetMessage(message Expression)
}

// This interface defines the methods supported by all attribute-like types.
type AttributeLike interface {
	GetVariable() string
	SetVariable(variable string)
	GetIndex(index int) Expression
	SetIndex(index int, expression Expression)
	GetIndices() ListLike[Expression]
	SetIndices(indices ListLike[Expression])
}

// This interface defines the methods supported by all block-like types.
type BlockLike interface {
	GetExpression() Expression
	SetExpression(expression Expression)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
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
	GetCitation() Expression
	SetCitation(citation Expression)
}

// This interface defines the methods supported by all evaluate-clause-like types.
type EvaluateClauseLike interface {
	GetRecipient() (recipient Recipient, assignment Assignment)
	SetRecipient(recipient Recipient, assignment Assignment)
	GetExpression() Expression
	SetExpression(expression Expression)
}

// This interface defines the methods supported by all if-clause-like types.
type IfClauseLike interface {
	GetCondition() Expression
	SetCondition(condition Expression)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetStatements() ProcedureLike
	SetStatements(statements ProcedureLike)
}

// This interface defines the methods supported by all notarize-clause-like types.
type NotarizeClauseLike interface {
	GetDraft() Expression
	SetDraft(draft Expression)
	GetMoniker() Expression
	SetMoniker(moniker Expression)
}

// This interface defines the methods supported by all on-clause-like types.
type OnClauseLike interface {
	GetException() Symbolic
	SetException(exception Symbolic)
	GetHandler(index int) BlockLike
	SetHandler(index int, handler BlockLike)
	GetHandlers() ListLike[BlockLike]
	SetHandlers(blocks ListLike[BlockLike])
}

// This interface defines the methods supported by all post-clause-like types.
type PostClauseLike interface {
	GetMessage() Expression
	SetMessage(message Expression)
	GetBag() Expression
	SetBag(bag Expression)
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
	GetDraft() Expression
	SetDraft(draft Expression)
	GetRecipient() Recipient
	SetRecipient(recipient Recipient)
}

// This interface defines the methods supported by all select-clause-like types.
type SelectClauseLike interface {
	GetControl() Expression
	SetControl(control Expression)
	GetOption(index int) BlockLike
	SetOption(index int, option BlockLike)
	GetOptions() ListLike[BlockLike]
	SetOptions(blocks ListLike[BlockLike])
}

// This interface defines the methods supported by all statement-like types.
type StatementLike interface {
	GetAnnotation() string
	SetAnnotation(annotation string)
	GetMainClause() Clause
	SetMainClause(mainClause Clause)
	GetNote() string
	SetNote(note string)
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
	GetCondition() Expression
	SetCondition(condition Expression)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetStatements() ProcedureLike
	SetStatements(statements ProcedureLike)
}

// This interface defines the methods supported by all with-clause-like types.
type WithClauseLike interface {
	GetItem() Symbolic
	SetItem(exception Symbolic)
	GetSequence() Expression
	SetSequence(sequence Expression)
	GetStatement(index int) StatementLike
	SetStatement(index int, statement StatementLike)
	GetStatements() ProcedureLike
	SetStatements(statements ProcedureLike)
}
