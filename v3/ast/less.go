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

func LessClass() LessClassLike {
	return lessClass()
}

// Constructor Methods

func (c *lessClass_) Less() LessLike {
	var instance = &less_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *less_) GetClass() LessClassLike {
	return lessClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type less_ struct {
	// Declare the instance attributes.
}

// Class Structure

type lessClass_ struct {
	// Declare the class constants.
}

// Class Reference

func lessClass() *lessClass_ {
	return lessClassReference_
}

var lessClassReference_ = &lessClass_{
	// Initialize the class constants.
}
