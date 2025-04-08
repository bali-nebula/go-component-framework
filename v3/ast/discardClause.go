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
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func DiscardClauseClass() DiscardClauseClassLike {
	return discardClauseClass()
}

// Constructor Methods

func (c *discardClauseClass_) DiscardClause(
	draft DraftLike,
) DiscardClauseLike {
	if uti.IsUndefined(draft) {
		panic("The \"draft\" attribute is required by this class.")
	}
	var instance = &discardClause_{
		// Initialize the instance attributes.
		draft_: draft,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *discardClause_) GetClass() DiscardClauseClassLike {
	return discardClauseClass()
}

// Attribute Methods

func (v *discardClause_) GetDraft() DraftLike {
	return v.draft_
}

// PROTECTED INTERFACE

// Instance Structure

type discardClause_ struct {
	// Declare the instance attributes.
	draft_ DraftLike
}

// Class Structure

type discardClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func discardClauseClass() *discardClauseClass_ {
	return discardClauseClassReference_
}

var discardClauseClassReference_ = &discardClauseClass_{
	// Initialize the class constants.
}
