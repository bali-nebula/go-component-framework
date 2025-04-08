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

func SelectClauseClass() SelectClauseClassLike {
	return selectClauseClass()
}

// Constructor Methods

func (c *selectClauseClass_) SelectClause(
	target TargetLike,
	matchHandlers col.Sequential[MatchHandlerLike],
) SelectClauseLike {
	if uti.IsUndefined(target) {
		panic("The \"target\" attribute is required by this class.")
	}
	if uti.IsUndefined(matchHandlers) {
		panic("The \"matchHandlers\" attribute is required by this class.")
	}
	var instance = &selectClause_{
		// Initialize the instance attributes.
		target_:        target,
		matchHandlers_: matchHandlers,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *selectClause_) GetClass() SelectClauseClassLike {
	return selectClauseClass()
}

// Attribute Methods

func (v *selectClause_) GetTarget() TargetLike {
	return v.target_
}

func (v *selectClause_) GetMatchHandlers() col.Sequential[MatchHandlerLike] {
	return v.matchHandlers_
}

// PROTECTED INTERFACE

// Instance Structure

type selectClause_ struct {
	// Declare the instance attributes.
	target_        TargetLike
	matchHandlers_ col.Sequential[MatchHandlerLike]
}

// Class Structure

type selectClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func selectClauseClass() *selectClauseClass_ {
	return selectClauseClassReference_
}

var selectClauseClassReference_ = &selectClauseClass_{
	// Initialize the class constants.
}
