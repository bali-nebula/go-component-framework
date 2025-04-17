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

func LowerInclusionClass() LowerInclusionClassLike {
	return lowerInclusionClass()
}

// Constructor Methods

func (c *lowerInclusionClass_) LowerInclusion() LowerInclusionLike {
	var instance = &lowerInclusion_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *lowerInclusion_) GetClass() LowerInclusionClassLike {
	return lowerInclusionClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type lowerInclusion_ struct {
	// Declare the instance attributes.
}

// Class Structure

type lowerInclusionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func lowerInclusionClass() *lowerInclusionClass_ {
	return lowerInclusionClassReference_
}

var lowerInclusionClassReference_ = &lowerInclusionClass_{
	// Initialize the class constants.
}
