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

func ExclusionClass() ExclusionClassLike {
	return exclusionClass()
}

// Constructor Methods

func (c *exclusionClass_) Exclusion() ExclusionLike {
	var instance = &exclusion_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exclusion_) GetClass() ExclusionClassLike {
	return exclusionClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type exclusion_ struct {
	// Declare the instance attributes.
}

// Class Structure

type exclusionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exclusionClass() *exclusionClass_ {
	return exclusionClassReference_
}

var exclusionClassReference_ = &exclusionClass_{
	// Initialize the class constants.
}
