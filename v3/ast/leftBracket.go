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

func LeftBracketClass() LeftBracketClassLike {
	return leftBracketClass()
}

// Constructor Methods

func (c *leftBracketClass_) LeftBracket(
	any_ any,
) LeftBracketLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &leftBracket_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *leftBracket_) GetClass() LeftBracketClassLike {
	return leftBracketClass()
}

// Attribute Methods

func (v *leftBracket_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type leftBracket_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type leftBracketClass_ struct {
	// Declare the class constants.
}

// Class Reference

func leftBracketClass() *leftBracketClass_ {
	return leftBracketClassReference_
}

var leftBracketClassReference_ = &leftBracketClass_{
	// Initialize the class constants.
}
