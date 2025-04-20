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

func AmpersandClass() AmpersandClassLike {
	return ampersandClass()
}

// Constructor Methods

func (c *ampersandClass_) Ampersand() AmpersandLike {
	var instance = &ampersand_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *ampersand_) GetClass() AmpersandClassLike {
	return ampersandClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type ampersand_ struct {
	// Declare the instance attributes.
}

// Class Structure

type ampersandClass_ struct {
	// Declare the class constants.
}

// Class Reference

func ampersandClass() *ampersandClass_ {
	return ampersandClassReference_
}

var ampersandClassReference_ = &ampersandClass_{
	// Initialize the class constants.
}
