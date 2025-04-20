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

func PlusEqualClass() PlusEqualClassLike {
	return plusEqualClass()
}

// Constructor Methods

func (c *plusEqualClass_) PlusEqual() PlusEqualLike {
	var instance = &plusEqual_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *plusEqual_) GetClass() PlusEqualClassLike {
	return plusEqualClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type plusEqual_ struct {
	// Declare the instance attributes.
}

// Class Structure

type plusEqualClass_ struct {
	// Declare the class constants.
}

// Class Reference

func plusEqualClass() *plusEqualClass_ {
	return plusEqualClassReference_
}

var plusEqualClassReference_ = &plusEqualClass_{
	// Initialize the class constants.
}
