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

func XorClass() XorClassLike {
	return xorClass()
}

// Constructor Methods

func (c *xorClass_) Xor() XorLike {
	var instance = &xor_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *xor_) GetClass() XorClassLike {
	return xorClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type xor_ struct {
	// Declare the instance attributes.
}

// Class Structure

type xorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func xorClass() *xorClass_ {
	return xorClassReference_
}

var xorClassReference_ = &xorClass_{
	// Initialize the class constants.
}
