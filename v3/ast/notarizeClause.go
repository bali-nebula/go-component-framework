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

func NotarizeClauseClass() NotarizeClauseClassLike {
	return notarizeClauseClass()
}

// Constructor Methods

func (c *notarizeClauseClass_) NotarizeClause(
	draft DraftLike,
	citation CitationLike,
) NotarizeClauseLike {
	if uti.IsUndefined(draft) {
		panic("The \"draft\" attribute is required by this class.")
	}
	if uti.IsUndefined(citation) {
		panic("The \"citation\" attribute is required by this class.")
	}
	var instance = &notarizeClause_{
		// Initialize the instance attributes.
		draft_:    draft,
		citation_: citation,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *notarizeClause_) GetClass() NotarizeClauseClassLike {
	return notarizeClauseClass()
}

// Attribute Methods

func (v *notarizeClause_) GetDraft() DraftLike {
	return v.draft_
}

func (v *notarizeClause_) GetCitation() CitationLike {
	return v.citation_
}

// PROTECTED INTERFACE

// Instance Structure

type notarizeClause_ struct {
	// Declare the instance attributes.
	draft_    DraftLike
	citation_ CitationLike
}

// Class Structure

type notarizeClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func notarizeClauseClass() *notarizeClauseClass_ {
	return notarizeClauseClassReference_
}

var notarizeClauseClassReference_ = &notarizeClauseClass_{
	// Initialize the class constants.
}
