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
	REGULAR Assignment = iota
	DEFAULT
	TIMES
	DIVIDE
	PLUS
	MINUS
)

// STATEMENT INTERFACES

// This interface defines the methods supported by all accept-clause-like types.
type AcceptLike interface {
	GetMessage() any
	SetMessage(message any)
}

// This interface defines the methods supported by all attribute-like types.
type AttributeLike interface {
	GetIdentifier() string
	SetIdentifier(identifier string)
	GetIndex(index int) any
	SetIndex(index int, expression any)
	GetIndices() ListLike[any]
	SetIndices(indices ListLike[any])
}

// This interface defines the methods supported by all block-like types.
type BlockLike interface {
	GetExpression() any
	SetExpression(expression any)
	GetStatement(index int) any
	SetStatement(index int, statement any)
	GetStatements() ListLike[any]
	SetStatements(statements ListLike[any])
}

// This interface defines the methods supported by all break-clause-like types.
type BreakLike interface {
}

// This interface defines the methods supported by all checkout-clause-like types.
type CheckoutLike interface {
	GetRecipient() any
	SetRecipient(recipient any)
	GetLevel() any
	SetLevel(level any)
	GetMoniker() any
	SetMoniker(moniker any)
}

// This interface defines the methods supported by all discard-clause-like types.
type DiscardLike interface {
	GetCitation() any
	SetCitation(citation any)
}

// This interface defines the methods supported by all evaluate-clause-like types.
type EvaluateLike interface {
	GetRecipient() any
	SetRecipient(recipient any)
	GetAssignment() any
	SetAssignment(assignment any)
	GetExpression() any
	SetExpression(expression any)
}

// This interface defines the methods supported by all if-clause-like types.
type IfLike interface {
	GetBlock() BlockLike
	SetBlock(block BlockLike)
}

// This interface defines the methods supported by all notarize-clause-like types.
type NotarizeLike interface {
	GetDraft() any
	SetDraft(draft any)
	GetMoniker() any
	SetMoniker(moniker any)
}

// This interface defines the methods supported by all on-clause-like types.
type OnLike interface {
	GetException() Symbolic
	SetException(exception Symbolic)
	GetBlock(index int) BlockLike
	SetBlock(index int, block BlockLike)
	GetBlocks() ListLike[BlockLike]
	SetBlocks(blocks ListLike[BlockLike])
}
