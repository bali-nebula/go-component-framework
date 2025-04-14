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

func TextualClass() TextualClassLike {
	return textualClass()
}

// Constructor Methods

func (c *textualClass_) Textual(
	any_ any,
) TextualLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &textual_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *textual_) GetClass() TextualClassLike {
	return textualClass()
}

// Attribute Methods

func (v *textual_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type textual_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type textualClass_ struct {
	// Declare the class constants.
}

// Class Reference

func textualClass() *textualClass_ {
	return textualClassReference_
}

var textualClassReference_ = &textualClass_{
	// Initialize the class constants.
}
