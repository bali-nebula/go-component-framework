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

func RightBracketClass() RightBracketClassLike {
	return rightBracketClass()
}

// Constructor Methods

func (c *rightBracketClass_) RightBracket(
	any_ any,
) RightBracketLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &rightBracket_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *rightBracket_) GetClass() RightBracketClassLike {
	return rightBracketClass()
}

// Attribute Methods

func (v *rightBracket_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type rightBracket_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type rightBracketClass_ struct {
	// Declare the class constants.
}

// Class Reference

func rightBracketClass() *rightBracketClass_ {
	return rightBracketClassReference_
}

var rightBracketClassReference_ = &rightBracketClass_{
	// Initialize the class constants.
}
