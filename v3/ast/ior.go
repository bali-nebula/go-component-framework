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

func IorClass() IorClassLike {
	return iorClass()
}

// Constructor Methods

func (c *iorClass_) Ior() IorLike {
	var instance = &ior_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *ior_) GetClass() IorClassLike {
	return iorClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type ior_ struct {
	// Declare the instance attributes.
}

// Class Structure

type iorClass_ struct {
	// Declare the class constants.
}

// Class Reference

func iorClass() *iorClass_ {
	return iorClassReference_
}

var iorClassReference_ = &iorClass_{
	// Initialize the class constants.
}
