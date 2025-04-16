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

func IsClass() IsClassLike {
	return isClass()
}

// Constructor Methods

func (c *isClass_) Is() IsLike {
	var instance = &is_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *is_) GetClass() IsClassLike {
	return isClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type is_ struct {
	// Declare the instance attributes.
}

// Class Structure

type isClass_ struct {
	// Declare the class constants.
}

// Class Reference

func isClass() *isClass_ {
	return isClassReference_
}

var isClassReference_ = &isClass_{
	// Initialize the class constants.
}
