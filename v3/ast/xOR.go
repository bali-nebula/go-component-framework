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

func XORClass() XORClassLike {
	return xORClass()
}

// Constructor Methods

func (c *xORClass_) XOR() XORLike {
	var instance = &xOR_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *xOR_) GetClass() XORClassLike {
	return xORClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type xOR_ struct {
	// Declare the instance attributes.
}

// Class Structure

type xORClass_ struct {
	// Declare the class constants.
}

// Class Reference

func xORClass() *xORClass_ {
	return xORClassReference_
}

var xORClassReference_ = &xORClass_{
	// Initialize the class constants.
}
