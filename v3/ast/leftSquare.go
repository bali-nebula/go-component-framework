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

import ()

// CLASS INTERFACE

// Access Function

func LeftSquareClass() LeftSquareClassLike {
	return leftSquareClass()
}

// Constructor Methods

func (c *leftSquareClass_) LeftSquare() LeftSquareLike {
	var instance = &leftSquare_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *leftSquare_) GetClass() LeftSquareClassLike {
	return leftSquareClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type leftSquare_ struct {
	// Declare the instance attributes.
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
