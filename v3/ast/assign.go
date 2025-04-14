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

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func AssignClass() AssignClassLike {
	return assignClass()
}

// Constructor Methods

func (c *assignClass_) Assign(
	any_ any,
) AssignLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &assign_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *assign_) GetClass() AssignClassLike {
	return assignClass()
}

// Attribute Methods

func (v *assign_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type assign_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type assignClass_ struct {
	// Declare the class constants.
}

// Class Reference

func assignClass() *assignClass_ {
	return assignClassReference_
}

var assignClassReference_ = &assignClass_{
	// Initialize the class constants.
}
