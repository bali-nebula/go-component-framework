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

func SanClass() SanClassLike {
	return sanClass()
}

// Constructor Methods

func (c *sanClass_) San() SanLike {
	var instance = &san_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *san_) GetClass() SanClassLike {
	return sanClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type san_ struct {
	// Declare the instance attributes.
}

// Class Structure

type sanClass_ struct {
	// Declare the class constants.
}

// Class Reference

func sanClass() *sanClass_ {
	return sanClassReference_
}

var sanClassReference_ = &sanClass_{
	// Initialize the class constants.
}
