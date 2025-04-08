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

func ValuesClass() ValuesClassLike {
	return valuesClass()
}

// Constructor Methods

func (c *valuesClass_) Values(
	any_ any,
) ValuesLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &values_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *values_) GetClass() ValuesClassLike {
	return valuesClass()
}

// Attribute Methods

func (v *values_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type values_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type valuesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func valuesClass() *valuesClass_ {
	return valuesClassReference_
}

var valuesClassReference_ = &valuesClass_{
	// Initialize the class constants.
}
