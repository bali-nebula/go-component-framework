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

func LowerExclusionClass() LowerExclusionClassLike {
	return lowerExclusionClass()
}

// Constructor Methods

func (c *lowerExclusionClass_) LowerExclusion() LowerExclusionLike {
	var instance = &lowerExclusion_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *lowerExclusion_) GetClass() LowerExclusionClassLike {
	return lowerExclusionClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type lowerExclusion_ struct {
	// Declare the instance attributes.
}

// Class Structure

type lowerExclusionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func lowerExclusionClass() *lowerExclusionClass_ {
	return lowerExclusionClassReference_
}

var lowerExclusionClassReference_ = &lowerExclusionClass_{
	// Initialize the class constants.
}
