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

func InlineStatementsClass() InlineStatementsClassLike {
	return inlineStatementsClass()
}

// Constructor Methods

func (c *inlineStatementsClass_) InlineStatements(
	statement StatementLike,
	additionalStatements col.Sequential[AdditionalStatementLike],
) InlineStatementsLike {
	if uti.IsUndefined(statement) {
		panic("The \"statement\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalStatements) {
		panic("The \"additionalStatements\" attribute is required by this class.")
	}
	var instance = &inlineStatements_{
		// Initialize the instance attributes.
		statement_:            statement,
		additionalStatements_: additionalStatements,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inlineStatements_) GetClass() InlineStatementsClassLike {
	return inlineStatementsClass()
}

// Attribute Methods

func (v *inlineStatements_) GetStatement() StatementLike {
	return v.statement_
}

func (v *inlineStatements_) GetAdditionalStatements() col.Sequential[AdditionalStatementLike] {
	return v.additionalStatements_
}

// PROTECTED INTERFACE

// Instance Structure

type inlineStatements_ struct {
	// Declare the instance attributes.
	statement_            StatementLike
	additionalStatements_ col.Sequential[AdditionalStatementLike]
}

// Class Structure

type inlineStatementsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineStatementsClass() *inlineStatementsClass_ {
	return inlineStatementsClassReference_
}

var inlineStatementsClassReference_ = &inlineStatementsClass_{
	// Initialize the class constants.
}
