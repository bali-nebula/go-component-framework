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

func SaveClauseClass() SaveClauseClassLike {
	return saveClauseClass()
}

// Constructor Methods

func (c *saveClauseClass_) SaveClause(
	draft DraftLike,
	cited CitedLike,
) SaveClauseLike {
	if uti.IsUndefined(draft) {
		panic("The \"draft\" attribute is required by this class.")
	}
	if uti.IsUndefined(cited) {
		panic("The \"cited\" attribute is required by this class.")
	}
	var instance = &saveClause_{
		// Initialize the instance attributes.
		draft_: draft,
		cited_: cited,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *saveClause_) GetClass() SaveClauseClassLike {
	return saveClauseClass()
}

// Attribute Methods

func (v *saveClause_) GetDraft() DraftLike {
	return v.draft_
}

func (v *saveClause_) GetCited() CitedLike {
	return v.cited_
}

// PROTECTED INTERFACE

// Instance Structure

type saveClause_ struct {
	// Declare the instance attributes.
	draft_ DraftLike
	cited_ CitedLike
}

// Class Structure

type saveClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func saveClauseClass() *saveClauseClass_ {
	return saveClauseClassReference_
}

var saveClauseClassReference_ = &saveClauseClass_{
	// Initialize the class constants.
}
