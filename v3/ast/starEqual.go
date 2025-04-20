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

func StarEqualClass() StarEqualClassLike {
	return starEqualClass()
}

// Constructor Methods

func (c *starEqualClass_) StarEqual() StarEqualLike {
	var instance = &starEqual_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *starEqual_) GetClass() StarEqualClassLike {
	return starEqualClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type starEqual_ struct {
	// Declare the instance attributes.
}

// Class Structure

type starEqualClass_ struct {
	// Declare the class constants.
}

// Class Reference

func starEqualClass() *starEqualClass_ {
	return starEqualClassReference_
}

var starEqualClassReference_ = &starEqualClass_{
	// Initialize the class constants.
}
