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

func WhileClauseClass() WhileClauseClassLike {
	return whileClauseClass()
}

// Constructor Methods

func (c *whileClauseClass_) WhileClause(
	condition ConditionLike,
	procedure ProcedureLike,
) WhileClauseLike {
	if uti.IsUndefined(condition) {
		panic("The \"condition\" attribute is required by this class.")
	}
	if uti.IsUndefined(procedure) {
		panic("The \"procedure\" attribute is required by this class.")
	}
	var instance = &whileClause_{
		// Initialize the instance attributes.
		condition_: condition,
		procedure_: procedure,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *whileClause_) GetClass() WhileClauseClassLike {
	return whileClauseClass()
}

// Attribute Methods

func (v *whileClause_) GetCondition() ConditionLike {
	return v.condition_
}

func (v *whileClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// PROTECTED INTERFACE

// Instance Structure

type whileClause_ struct {
	// Declare the instance attributes.
	condition_ ConditionLike
	procedure_ ProcedureLike
}

// Class Structure

type whileClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func whileClauseClass() *whileClauseClass_ {
	return whileClauseClassReference_
}

var whileClauseClassReference_ = &whileClauseClass_{
	// Initialize the class constants.
}
