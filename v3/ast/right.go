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

func RightClass() RightClassLike {
	return rightClass()
}

// Constructor Methods

func (c *rightClass_) Right() RightLike {
	var instance = &right_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *right_) GetClass() RightClassLike {
	return rightClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type right_ struct {
	// Declare the instance attributes.
}

// Class Structure

type rightClass_ struct {
	// Declare the class constants.
}

// Class Reference

func rightClass() *rightClass_ {
	return rightClassReference_
}

var rightClassReference_ = &rightClass_{
	// Initialize the class constants.
}
