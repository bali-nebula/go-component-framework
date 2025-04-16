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

func StatementLineClass() StatementLineClassLike {
	return statementLineClass()
}

// Constructor Methods

func (c *statementLineClass_) StatementLine(
	statement StatementLike,
	optionalNote string,
) StatementLineLike {
	if uti.IsUndefined(statement) {
		panic("The \"statement\" attribute is required by this class.")
	}
	var instance = &statementLine_{
		// Initialize the instance attributes.
		statement_:    statement,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *statementLine_) GetClass() StatementLineClassLike {
	return statementLineClass()
}

// Attribute Methods

func (v *statementLine_) GetStatement() StatementLike {
	return v.statement_
}

func (v *statementLine_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type statementLine_ struct {
	// Declare the instance attributes.
	statement_    StatementLike
	optionalNote_ string
}

// Class Structure

type statementLineClass_ struct {
	// Declare the class constants.
}

// Class Reference

func statementLineClass() *statementLineClass_ {
	return statementLineClassReference_
}

var statementLineClassReference_ = &statementLineClass_{
	// Initialize the class constants.
}
