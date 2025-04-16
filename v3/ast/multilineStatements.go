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

func MultilineStatementsClass() MultilineStatementsClassLike {
	return multilineStatementsClass()
}

// Constructor Methods

func (c *multilineStatementsClass_) MultilineStatements(
	newline string,
	annotatedStatements col.Sequential[AnnotatedStatementLike],
) MultilineStatementsLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(annotatedStatements) {
		panic("The \"annotatedStatements\" attribute is required by this class.")
	}
	var instance = &multilineStatements_{
		// Initialize the instance attributes.
		newline_:             newline,
		annotatedStatements_: annotatedStatements,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multilineStatements_) GetClass() MultilineStatementsClassLike {
	return multilineStatementsClass()
}

// Attribute Methods

func (v *multilineStatements_) GetNewline() string {
	return v.newline_
}

func (v *multilineStatements_) GetAnnotatedStatements() col.Sequential[AnnotatedStatementLike] {
	return v.annotatedStatements_
}

// PROTECTED INTERFACE

// Instance Structure

type multilineStatements_ struct {
	// Declare the instance attributes.
	newline_             string
	annotatedStatements_ col.Sequential[AnnotatedStatementLike]
}

// Class Structure

type multilineStatementsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multilineStatementsClass() *multilineStatementsClass_ {
	return multilineStatementsClassReference_
}

var multilineStatementsClassReference_ = &multilineStatementsClass_{
	// Initialize the class constants.
}
