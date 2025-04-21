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

func InclusionClass() InclusionClassLike {
	return inclusionClass()
}

// Constructor Methods

func (c *inclusionClass_) Inclusion() InclusionLike {
	var instance = &inclusion_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inclusion_) GetClass() InclusionClassLike {
	return inclusionClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type inclusion_ struct {
	// Declare the instance attributes.
}

// Class Structure

type inclusionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inclusionClass() *inclusionClass_ {
	return inclusionClassReference_
}

var inclusionClassReference_ = &inclusionClass_{
	// Initialize the class constants.
}
