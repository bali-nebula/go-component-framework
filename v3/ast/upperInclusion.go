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

func UpperInclusionClass() UpperInclusionClassLike {
	return upperInclusionClass()
}

// Constructor Methods

func (c *upperInclusionClass_) UpperInclusion() UpperInclusionLike {
	var instance = &upperInclusion_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *upperInclusion_) GetClass() UpperInclusionClassLike {
	return upperInclusionClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type upperInclusion_ struct {
	// Declare the instance attributes.
}

// Class Structure

type upperInclusionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func upperInclusionClass() *upperInclusionClass_ {
	return upperInclusionClassReference_
}

var upperInclusionClassReference_ = &upperInclusionClass_{
	// Initialize the class constants.
}
