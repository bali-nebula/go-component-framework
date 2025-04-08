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

func RejectClauseClass() RejectClauseClassLike {
	return rejectClauseClass()
}

// Constructor Methods

func (c *rejectClauseClass_) RejectClause(
	message MessageLike,
) RejectClauseLike {
	if uti.IsUndefined(message) {
		panic("The \"message\" attribute is required by this class.")
	}
	var instance = &rejectClause_{
		// Initialize the instance attributes.
		message_: message,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *rejectClause_) GetClass() RejectClauseClassLike {
	return rejectClauseClass()
}

// Attribute Methods

func (v *rejectClause_) GetMessage() MessageLike {
	return v.message_
}

// PROTECTED INTERFACE

// Instance Structure

type rejectClause_ struct {
	// Declare the instance attributes.
	message_ MessageLike
}

// Class Structure

type rejectClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func rejectClauseClass() *rejectClauseClass_ {
	return rejectClauseClassReference_
}

var rejectClauseClassReference_ = &rejectClauseClass_{
	// Initialize the class constants.
}
