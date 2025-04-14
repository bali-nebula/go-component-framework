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

func ThreadingClass() ThreadingClassLike {
	return threadingClass()
}

// Constructor Methods

func (c *threadingClass_) Threading(
	any_ any,
) ThreadingLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &threading_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *threading_) GetClass() ThreadingClassLike {
	return threadingClass()
}

// Attribute Methods

func (v *threading_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type threading_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type threadingClass_ struct {
	// Declare the class constants.
}

// Class Reference

func threadingClass() *threadingClass_ {
	return threadingClassReference_
}

var threadingClassReference_ = &threadingClass_{
	// Initialize the class constants.
}
