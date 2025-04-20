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

func SlashEqualClass() SlashEqualClassLike {
	return slashEqualClass()
}

// Constructor Methods

func (c *slashEqualClass_) SlashEqual() SlashEqualLike {
	var instance = &slashEqual_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *slashEqual_) GetClass() SlashEqualClassLike {
	return slashEqualClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type slashEqual_ struct {
	// Declare the instance attributes.
}

// Class Structure

type slashEqualClass_ struct {
	// Declare the class constants.
}

// Class Reference

func slashEqualClass() *slashEqualClass_ {
	return slashEqualClassReference_
}

var slashEqualClassReference_ = &slashEqualClass_{
	// Initialize the class constants.
}
