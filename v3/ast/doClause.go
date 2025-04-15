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

func DoClauseClass() DoClauseClassLike {
	return doClauseClass()
}

// Constructor Methods

func (c *doClauseClass_) DoClause(
	invocation InvocationLike,
) DoClauseLike {
	if uti.IsUndefined(invocation) {
		panic("The \"invocation\" attribute is required by this class.")
	}
	var instance = &doClause_{
		// Initialize the instance attributes.
		invocation_: invocation,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *doClause_) GetClass() DoClauseClassLike {
	return doClauseClass()
}

// Attribute Methods

func (v *doClause_) GetInvocation() InvocationLike {
	return v.invocation_
}

// PROTECTED INTERFACE

// Instance Structure

type doClause_ struct {
	// Declare the instance attributes.
	invocation_ InvocationLike
}

// Class Structure

type doClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func doClauseClass() *doClauseClass_ {
	return doClauseClassReference_
}

var doClauseClassReference_ = &doClauseClass_{
	// Initialize the class constants.
}
