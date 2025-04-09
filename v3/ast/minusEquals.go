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

func MinusEqualsClass() MinusEqualsClassLike {
	return minusEqualsClass()
}

// Constructor Methods

func (c *minusEqualsClass_) MinusEquals() MinusEqualsLike {
	var instance = &minusEquals_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *minusEquals_) GetClass() MinusEqualsClassLike {
	return minusEqualsClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type minusEquals_ struct {
	// Declare the instance attributes.
}

// Class Structure

type minusEqualsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func minusEqualsClass() *minusEqualsClass_ {
	return minusEqualsClassReference_
}

var minusEqualsClassReference_ = &minusEqualsClass_{
	// Initialize the class constants.
}
