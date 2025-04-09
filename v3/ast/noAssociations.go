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

func NoAssociationsClass() NoAssociationsClassLike {
	return noAssociationsClass()
}

// Constructor Methods

func (c *noAssociationsClass_) NoAssociations() NoAssociationsLike {
	var instance = &noAssociations_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *noAssociations_) GetClass() NoAssociationsClassLike {
	return noAssociationsClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type noAssociations_ struct {
	// Declare the instance attributes.
}

// Class Structure

type noAssociationsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noAssociationsClass() *noAssociationsClass_ {
	return noAssociationsClassReference_
}

var noAssociationsClassReference_ = &noAssociationsClass_{
	// Initialize the class constants.
}
