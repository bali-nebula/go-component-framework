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

func ExInclusiveCitationsClass() ExInclusiveCitationsClassLike {
	return exInclusiveCitationsClass()
}

// Constructor Methods

func (c *exInclusiveCitationsClass_) ExInclusiveCitations(
	citation1 string,
	citation2 string,
) ExInclusiveCitationsLike {
	if uti.IsUndefined(citation1) {
		panic("The \"citation1\" attribute is required by this class.")
	}
	if uti.IsUndefined(citation2) {
		panic("The \"citation2\" attribute is required by this class.")
	}
	var instance = &exInclusiveCitations_{
		// Initialize the instance attributes.
		citation1_: citation1,
		citation2_: citation2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exInclusiveCitations_) GetClass() ExInclusiveCitationsClassLike {
	return exInclusiveCitationsClass()
}

// Attribute Methods

func (v *exInclusiveCitations_) GetCitation1() string {
	return v.citation1_
}

func (v *exInclusiveCitations_) GetCitation2() string {
	return v.citation2_
}

// PROTECTED INTERFACE

// Instance Structure

type exInclusiveCitations_ struct {
	// Declare the instance attributes.
	citation1_ string
	citation2_ string
}

// Class Structure

type exInclusiveCitationsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exInclusiveCitationsClass() *exInclusiveCitationsClass_ {
	return exInclusiveCitationsClassReference_
}

var exInclusiveCitationsClassReference_ = &exInclusiveCitationsClass_{
	// Initialize the class constants.
}
