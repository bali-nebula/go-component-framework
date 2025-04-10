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

func RightSquareClass() RightSquareClassLike {
	return rightSquareClass()
}

// Constructor Methods

func (c *rightSquareClass_) RightSquare() RightSquareLike {
	var instance = &rightSquare_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *rightSquare_) GetClass() RightSquareClassLike {
	return rightSquareClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type rightSquare_ struct {
	// Declare the instance attributes.
}

// Class Structure

type rightSquareClass_ struct {
	// Declare the class constants.
}

// Class Reference

func rightSquareClass() *rightSquareClass_ {
	return rightSquareClassReference_
}

var rightSquareClassReference_ = &rightSquareClass_{
	// Initialize the class constants.
}
