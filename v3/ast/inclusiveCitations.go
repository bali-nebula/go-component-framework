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

func InclusiveCitationsClass() InclusiveCitationsClassLike {
	return inclusiveCitationsClass()
}

// Constructor Methods

func (c *inclusiveCitationsClass_) InclusiveCitations(
	citation1 string,
	citation2 string,
) InclusiveCitationsLike {
	if uti.IsUndefined(citation1) {
		panic("The \"citation1\" attribute is required by this class.")
	}
	if uti.IsUndefined(citation2) {
		panic("The \"citation2\" attribute is required by this class.")
	}
	var instance = &inclusiveCitations_{
		// Initialize the instance attributes.
		citation1_: citation1,
		citation2_: citation2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inclusiveCitations_) GetClass() InclusiveCitationsClassLike {
	return inclusiveCitationsClass()
}

// Attribute Methods

func (v *inclusiveCitations_) GetCitation1() string {
	return v.citation1_
}

func (v *inclusiveCitations_) GetCitation2() string {
	return v.citation2_
}

// PROTECTED INTERFACE

// Instance Structure

type inclusiveCitations_ struct {
	// Declare the instance attributes.
	citation1_ string
	citation2_ string
}

// Class Structure

type inclusiveCitationsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inclusiveCitationsClass() *inclusiveCitationsClass_ {
	return inclusiveCitationsClassReference_
}

var inclusiveCitationsClassReference_ = &inclusiveCitationsClass_{
	// Initialize the class constants.
}
