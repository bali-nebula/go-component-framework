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

func DotClass() DotClassLike {
	return dotClass()
}

// Constructor Methods

func (c *dotClass_) Dot() DotLike {
	var instance = &dot_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *dot_) GetClass() DotClassLike {
	return dotClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type dot_ struct {
	// Declare the instance attributes.
}

// Class Structure

type dotClass_ struct {
	// Declare the class constants.
}

// Class Reference

func dotClass() *dotClass_ {
	return dotClassReference_
}

var dotClassReference_ = &dotClass_{
	// Initialize the class constants.
}
