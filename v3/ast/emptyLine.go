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

func EmptyLineClass() EmptyLineClassLike {
	return emptyLineClass()
}

// Constructor Methods

func (c *emptyLineClass_) EmptyLine(
	newline string,
) EmptyLineLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	var instance = &emptyLine_{
		// Initialize the instance attributes.
		newline_: newline,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *emptyLine_) GetClass() EmptyLineClassLike {
	return emptyLineClass()
}

// Attribute Methods

func (v *emptyLine_) GetNewline() string {
	return v.newline_
}

// PROTECTED INTERFACE

// Instance Structure

type emptyLine_ struct {
	// Declare the instance attributes.
	newline_ string
}

// Class Structure

type emptyLineClass_ struct {
	// Declare the class constants.
}

// Class Reference

func emptyLineClass() *emptyLineClass_ {
	return emptyLineClassReference_
}

var emptyLineClassReference_ = &emptyLineClass_{
	// Initialize the class constants.
}
