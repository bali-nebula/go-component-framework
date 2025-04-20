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

func ArrowClass() ArrowClassLike {
	return arrowClass()
}

// Constructor Methods

func (c *arrowClass_) Arrow() ArrowLike {
	var instance = &arrow_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *arrow_) GetClass() ArrowClassLike {
	return arrowClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type arrow_ struct {
	// Declare the instance attributes.
}

// Class Structure

type arrowClass_ struct {
	// Declare the class constants.
}

// Class Reference

func arrowClass() *arrowClass_ {
	return arrowClassReference_
}

var arrowClassReference_ = &arrowClass_{
	// Initialize the class constants.
}
