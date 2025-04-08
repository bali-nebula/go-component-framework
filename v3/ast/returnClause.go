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

func ReturnClauseClass() ReturnClauseClassLike {
	return returnClauseClass()
}

// Constructor Methods

func (c *returnClauseClass_) ReturnClause(
	result ResultLike,
) ReturnClauseLike {
	if uti.IsUndefined(result) {
		panic("The \"result\" attribute is required by this class.")
	}
	var instance = &returnClause_{
		// Initialize the instance attributes.
		result_: result,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *returnClause_) GetClass() ReturnClauseClassLike {
	return returnClauseClass()
}

// Attribute Methods

func (v *returnClause_) GetResult() ResultLike {
	return v.result_
}

// PROTECTED INTERFACE

// Instance Structure

type returnClause_ struct {
	// Declare the instance attributes.
	result_ ResultLike
}

// Class Structure

type returnClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func returnClauseClass() *returnClauseClass_ {
	return returnClauseClassReference_
}

var returnClauseClassReference_ = &returnClauseClass_{
	// Initialize the class constants.
}
