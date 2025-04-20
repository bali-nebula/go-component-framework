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

func ColonEqualClass() ColonEqualClassLike {
	return colonEqualClass()
}

// Constructor Methods

func (c *colonEqualClass_) ColonEqual() ColonEqualLike {
	var instance = &colonEqual_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *colonEqual_) GetClass() ColonEqualClassLike {
	return colonEqualClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type colonEqual_ struct {
	// Declare the instance attributes.
}

// Class Structure

type colonEqualClass_ struct {
	// Declare the class constants.
}

// Class Reference

func colonEqualClass() *colonEqualClass_ {
	return colonEqualClassReference_
}

var colonEqualClassReference_ = &colonEqualClass_{
	// Initialize the class constants.
}
