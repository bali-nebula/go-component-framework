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

func AdditionalStatementClass() AdditionalStatementClassLike {
	return additionalStatementClass()
}

// Constructor Methods

func (c *additionalStatementClass_) AdditionalStatement(
	statement StatementLike,
) AdditionalStatementLike {
	if uti.IsUndefined(statement) {
		panic("The \"statement\" attribute is required by this class.")
	}
	var instance = &additionalStatement_{
		// Initialize the instance attributes.
		statement_: statement,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalStatement_) GetClass() AdditionalStatementClassLike {
	return additionalStatementClass()
}

// Attribute Methods

func (v *additionalStatement_) GetStatement() StatementLike {
	return v.statement_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalStatement_ struct {
	// Declare the instance attributes.
	statement_ StatementLike
}

// Class Structure

type additionalStatementClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalStatementClass() *additionalStatementClass_ {
	return additionalStatementClassReference_
}

var additionalStatementClassReference_ = &additionalStatementClass_{
	// Initialize the class constants.
}
