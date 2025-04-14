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

func NumericalClass() NumericalClassLike {
	return numericalClass()
}

// Constructor Methods

func (c *numericalClass_) Numerical(
	any_ any,
) NumericalLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &numerical_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *numerical_) GetClass() NumericalClassLike {
	return numericalClass()
}

// Attribute Methods

func (v *numerical_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type numerical_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type numericalClass_ struct {
	// Declare the class constants.
}

// Class Reference

func numericalClass() *numericalClass_ {
	return numericalClassReference_
}

var numericalClassReference_ = &numericalClass_{
	// Initialize the class constants.
}
