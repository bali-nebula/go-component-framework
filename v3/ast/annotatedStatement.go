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

func AnnotatedStatementClass() AnnotatedStatementClassLike {
	return annotatedStatementClass()
}

// Constructor Methods

func (c *annotatedStatementClass_) AnnotatedStatement(
	xyz string,
	statement StatementLike,
	optionalNote string,
) AnnotatedStatementLike {
	if uti.IsUndefined(xyz) {
		panic("The \"xyz\" attribute is required by this class.")
	}
	if uti.IsUndefined(statement) {
		panic("The \"statement\" attribute is required by this class.")
	}
	var instance = &annotatedStatement_{
		// Initialize the instance attributes.
		xyz_:          xyz,
		statement_:    statement,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *annotatedStatement_) GetClass() AnnotatedStatementClassLike {
	return annotatedStatementClass()
}

// Attribute Methods

func (v *annotatedStatement_) GetXyz() string {
	return v.xyz_
}

func (v *annotatedStatement_) GetStatement() StatementLike {
	return v.statement_
}

func (v *annotatedStatement_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type annotatedStatement_ struct {
	// Declare the instance attributes.
	xyz_          string
	statement_    StatementLike
	optionalNote_ string
}

// Class Structure

type annotatedStatementClass_ struct {
	// Declare the class constants.
}

// Class Reference

func annotatedStatementClass() *annotatedStatementClass_ {
	return annotatedStatementClassReference_
}

var annotatedStatementClassReference_ = &annotatedStatementClass_{
	// Initialize the class constants.
}
