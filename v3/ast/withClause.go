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

func WithClauseClass() WithClauseClassLike {
	return withClauseClass()
}

// Constructor Methods

func (c *withClauseClass_) WithClause(
	item ItemLike,
	sequence SequenceLike,
	procedure ProcedureLike,
) WithClauseLike {
	if uti.IsUndefined(item) {
		panic("The \"item\" attribute is required by this class.")
	}
	if uti.IsUndefined(sequence) {
		panic("The \"sequence\" attribute is required by this class.")
	}
	if uti.IsUndefined(procedure) {
		panic("The \"procedure\" attribute is required by this class.")
	}
	var instance = &withClause_{
		// Initialize the instance attributes.
		item_:      item,
		sequence_:  sequence,
		procedure_: procedure,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *withClause_) GetClass() WithClauseClassLike {
	return withClauseClass()
}

// Attribute Methods

func (v *withClause_) GetItem() ItemLike {
	return v.item_
}

func (v *withClause_) GetSequence() SequenceLike {
	return v.sequence_
}

func (v *withClause_) GetProcedure() ProcedureLike {
	return v.procedure_
}

// PROTECTED INTERFACE

// Instance Structure

type withClause_ struct {
	// Declare the instance attributes.
	item_      ItemLike
	sequence_  SequenceLike
	procedure_ ProcedureLike
}

// Class Structure

type withClauseClass_ struct {
	// Declare the class constants.
}

// Class Reference

func withClauseClass() *withClauseClass_ {
	return withClauseClassReference_
}

var withClauseClassReference_ = &withClauseClass_{
	// Initialize the class constants.
}
