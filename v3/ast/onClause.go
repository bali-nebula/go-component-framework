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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func OnClauseClass() OnClauseClassLike {
	return onClauseClass()
}

// Constructor Methods

func (c *onClauseClass_) OnClause(
	failure FailureLike,
	matchHandlers col.Sequential[MatchHandlerLike],
) OnClauseLike {
	if uti.IsUndefined(failure) {
		panic("The \"failure\" attribute is required by this class.")
	}
	if uti.IsUndefined(matchHandlers) {
		panic("The \"matchHandlers\" attribute is required by this class.")
	}
	var instance = &onClause_{
		// Initialize the instance attributes.
		failure_:       failure,
		matchHandlers_: matchHandlers,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *onClause_) GetClass() OnClauseClassLike {
	return onClauseClass()
}

// Attribute Methods

func (v *onClause_) GetFailure() FailureLike {
	return v.failure_
}

func (v *onClause_) GetMatchHandlers() col.Sequential[MatchHandlerLike] {
	return v.matchHandlers_
}

// PROTECTED INTERFACE

// Instance Structure

type onClause_ struct {
	// Declare the instance attributes.
	failure_       FailureLike
	matchHandlers_ col.Sequential[MatchHandlerLike]
}

// Class Structure

type onClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func onClauseClass() *onClauseClass_ {
	return onClauseClassReference_
}

var onClauseClassReference_ = &onClauseClass_{
	// Initialize the class constants.
}
