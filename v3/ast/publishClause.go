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

func PublishClauseClass() PublishClauseClassLike {
	return publishClauseClass()
}

// Constructor Methods

func (c *publishClauseClass_) PublishClause(
	event EventLike,
) PublishClauseLike {
	if uti.IsUndefined(event) {
		panic("The \"event\" attribute is required by this class.")
	}
	var instance = &publishClause_{
		// Initialize the instance attributes.
		event_: event,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *publishClause_) GetClass() PublishClauseClassLike {
	return publishClauseClass()
}

// Attribute Methods

func (v *publishClause_) GetEvent() EventLike {
	return v.event_
}

// PROTECTED INTERFACE

// Instance Structure

type publishClause_ struct {
	// Declare the instance attributes.
	event_ EventLike
}

// Class Structure

type publishClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func publishClauseClass() *publishClauseClass_ {
	return publishClauseClassReference_
}

var publishClauseClassReference_ = &publishClauseClass_{
	// Initialize the class constants.
}
