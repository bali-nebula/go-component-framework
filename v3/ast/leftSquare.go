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

func LeftSquareClass() LeftSquareClassLike {
	return leftSquareClass()
}

// Constructor Methods

func (c *leftSquareClass_) LeftSquare(
	bar string,
) LeftSquareLike {
	if uti.IsUndefined(bar) {
		panic("The \"bar\" attribute is required by this class.")
	}
	var instance = &leftSquare_{
		// Initialize the instance attributes.
		bar_: bar,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *leftSquare_) GetClass() LeftSquareClassLike {
	return leftSquareClass()
}

// Attribute Methods

func (v *leftSquare_) GetBar() string {
	return v.bar_
}

// PROTECTED INTERFACE

// Instance Structure

type leftSquare_ struct {
	// Declare the instance attributes.
	bar_ string
}

// Class Structure

type leftSquareClass_ struct {
	// Declare the class constants.
}

// Class Reference

func leftSquareClass() *leftSquareClass_ {
	return leftSquareClassReference_
}

var leftSquareClassReference_ = &leftSquareClass_{
	// Initialize the class constants.
}
