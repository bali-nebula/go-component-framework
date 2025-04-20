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

func DoubleClass() DoubleClassLike {
	return doubleClass()
}

// Constructor Methods

func (c *doubleClass_) Double() DoubleLike {
	var instance = &double_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *double_) GetClass() DoubleClassLike {
	return doubleClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type double_ struct {
	// Declare the instance attributes.
}

// Class Structure

type doubleClass_ struct {
	// Declare the class constants.
}

// Class Reference

func doubleClass() *doubleClass_ {
	return doubleClassReference_
}

var doubleClassReference_ = &doubleClass_{
	// Initialize the class constants.
}
