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

func StatementClass() StatementClassLike {
	return statementClass()
}

// Constructor Methods

func (c *statementClass_) Statement(
	mainClause MainClauseLike,
	optionalOnClause OnClauseLike,
) StatementLike {
	if uti.IsUndefined(mainClause) {
		panic("The \"mainClause\" attribute is required by this class.")
	}
	var instance = &statement_{
		// Initialize the instance attributes.
		mainClause_:       mainClause,
		optionalOnClause_: optionalOnClause,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *statement_) GetClass() StatementClassLike {
	return statementClass()
}

// Attribute Methods

func (v *statement_) GetMainClause() MainClauseLike {
	return v.mainClause_
}

func (v *statement_) GetOptionalOnClause() OnClauseLike {
	return v.optionalOnClause_
}

// PROTECTED INTERFACE

// Instance Structure

type statement_ struct {
	// Declare the instance attributes.
	mainClause_       MainClauseLike
	optionalOnClause_ OnClauseLike
}

// Class Structure

type statementClass_ struct {
	// Declare the class constants.
}

// Class Reference

func statementClass() *statementClass_ {
	return statementClassReference_
}

var statementClassReference_ = &statementClass_{
	// Initialize the class constants.
}
