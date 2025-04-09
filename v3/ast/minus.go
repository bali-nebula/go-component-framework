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

func MinusClass() MinusClassLike {
	return minusClass()
}

// Constructor Methods

func (c *minusClass_) Minus() MinusLike {
	var instance = &minus_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *minus_) GetClass() MinusClassLike {
	return minusClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type minus_ struct {
	// Declare the instance attributes.
}

// Class Structure

type minusClass_ struct {
	// Declare the class constants.
}

// Class Reference

func minusClass() *minusClass_ {
	return minusClassReference_
}

var minusClassReference_ = &minusClass_{
	// Initialize the class constants.
}
