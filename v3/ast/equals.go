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

func EqualsClass() EqualsClassLike {
	return equalsClass()
}

// Constructor Methods

func (c *equalsClass_) Equals() EqualsLike {
	var instance = &equals_{
		// Initialize the instance attributes.
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *equals_) GetClass() EqualsClassLike {
	return equalsClass()
}

// Attribute Methods

// PROTECTED INTERFACE

// Instance Structure

type equals_ struct {
	// Declare the instance attributes.
}

// Class Structure

type equalsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func equalsClass() *equalsClass_ {
	return equalsClassReference_
}

var equalsClassReference_ = &equalsClass_{
	// Initialize the class constants.
}
