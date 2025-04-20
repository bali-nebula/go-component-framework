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

func EqualClass() EqualClassLike {
	return equalClass()
}

// Constructor Methods

func (c *equalClass_) Equal() EqualLike {
	var instance = &equal_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *equal_) GetClass() EqualClassLike {
	return equalClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type equal_ struct {
	// Declare the instance attributes.
}

// Class Structure

type equalClass_ struct {
	// Declare the class constants.
}

// Class Reference

func equalClass() *equalClass_ {
	return equalClassReference_
}

var equalClassReference_ = &equalClass_{
	// Initialize the class constants.
}
