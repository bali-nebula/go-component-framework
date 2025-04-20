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

func DashEqualClass() DashEqualClassLike {
	return dashEqualClass()
}

// Constructor Methods

func (c *dashEqualClass_) DashEqual() DashEqualLike {
	var instance = &dashEqual_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *dashEqual_) GetClass() DashEqualClassLike {
	return dashEqualClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type dashEqual_ struct {
	// Declare the instance attributes.
}

// Class Structure

type dashEqualClass_ struct {
	// Declare the class constants.
}

// Class Reference

func dashEqualClass() *dashEqualClass_ {
	return dashEqualClassReference_
}

var dashEqualClassReference_ = &dashEqualClass_{
	// Initialize the class constants.
}
