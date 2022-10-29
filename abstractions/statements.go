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

// STATEMENT INTERFACES

// This interface defines the methods supported by all accept-clause-like types.
type AcceptClauseLike interface {
	GetMessage() any
	SetMessage(message any)
}

// This interface defines the methods supported by all attribute-like types.
type AttributeLike interface {
	GetVariable() string
	SetVariable(variable string)
	GetIndex(index int) any
	SetIndex(index int, expression any)
	GetIndices() ListLike[any]
	SetIndices(indices ListLike[any])
}

// This interface defines the methods supported by all block-like types.
type BlockLike interface {
	GetPattern() any
	SetPattern(pattern any)
	GetStatement(index int) any
	SetStatement(index int, statement any)
	GetStatements() ListLike[any]
	SetStatements(statements ListLike[any])
}

// This interface defines the methods supported by all break-clause-like types.
type BreakClauseLike interface {
}

// This interface defines the methods supported by all checkout-clause-like types.
type CheckoutClauseLike interface {
	GetRecipient() any
	SetRecipient(recipient any)
	GetLevel() any
	SetLevel(level any)
	GetMoniker() any
	SetMoniker(moniker any)
}

// This interface defines the methods supported by all continue-clause-like types.
type ContinueClauseLike interface {
}

// This interface defines the methods supported by all discard-clause-like types.
type DiscardClauseLike interface {
	GetCitation() any
	SetCitation(citation any)
}

// This interface defines the methods supported by all evaluate-clause-like types.
type EvaluateClauseLike interface {
	GetRecipient() (recipient any, assignment Assignment)
	SetRecipient(recipient any, assignment Assignment)
	GetExpression() any
	SetExpression(expression any)
}

// This interface defines the methods supported by all if-clause-like types.
type IfClauseLike interface {
	GetCondition() any
	SetCondition(condition any)
	GetStatement(index int) any
	SetStatement(index int, statement any)
	GetStatements() ListLike[any]
	SetStatements(statements ListLike[any])
}

// This interface defines the methods supported by all notarize-clause-like types.
type NotarizeClauseLike interface {
	GetDraft() any
	SetDraft(draft any)
	GetMoniker() any
	SetMoniker(moniker any)
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
	GetMessage() any
	SetMessage(message any)
	GetBag() any
	SetBag(bag any)
}

// This interface consolidates all the interfaces supported by procedure-like
// types.
type ProcedureLike interface {
	Sequential[any]
	Indexed[any]
	Updatable[any]
	Malleable[any]
}

// This interface defines the methods supported by all publish-clause-like types.
type PublishClauseLike interface {
	GetEvent() any
	SetEvent(event any)
}

// This interface defines the methods supported by all reject-clause-like types.
type RejectClauseLike interface {
	GetMessage() any
	SetMessage(message any)
}

// This interface defines the methods supported by all retrieve-clause-like types.
type RetrieveClauseLike interface {
	GetRecipient() any
	SetRecipient(recipient any)
	GetBag() any
	SetBag(bag any)
}

// This interface defines the methods supported by all return-clause-like types.
type ReturnClauseLike interface {
	GetResult() any
	SetResult(result any)
}

// This interface defines the methods supported by all save-clause-like types.
type SaveClauseLike interface {
	GetDraft() any
	SetDraft(draft any)
	GetCitation() any
	SetCitation(citation any)
}

// This interface defines the methods supported by all select-clause-like types.
type SelectClauseLike interface {
	GetControl() any
	SetControl(control any)
	GetOption(index int) BlockLike
	SetOption(index int, option BlockLike)
	GetOptions() ListLike[BlockLike]
	SetOptions(blocks ListLike[BlockLike])
}

// This interface defines the methods supported by all statement-like types.
type StatementLike interface {
	GetMainClause() any
	SetMainClause(mainClause any)
	GetOnClause() OnClauseLike
	SetOnClause(onClause OnClauseLike)
}

// This interface defines the methods supported by all throw-clause-like types.
type ThrowClauseLike interface {
	GetException() any
	SetException(exception any)
}

// This interface defines the methods supported by all while-clause-like types.
type WhileClauseLike interface {
	GetCondition() any
	SetCondition(condition any)
	GetStatement(index int) any
	SetStatement(index int, statement any)
	GetStatements() ListLike[any]
	SetStatements(statements ListLike[any])
}

// This interface defines the methods supported by all with-clause-like types.
type WithClauseLike interface {
	GetItem() Symbolic
	SetItem(exception Symbolic)
	GetSequence() any
	SetSequence(sequence any)
	GetStatement(index int) any
	SetStatement(index int, statement any)
	GetStatements() ListLike[any]
	SetStatements(statements ListLike[any])
}
