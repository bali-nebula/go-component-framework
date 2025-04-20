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

func MagnitudeClass() MagnitudeClassLike {
	return magnitudeClass()
}

// Constructor Methods

func (c *magnitudeClass_) Magnitude(
	numerical NumericalLike,
) MagnitudeLike {
	if uti.IsUndefined(numerical) {
		panic("The \"numerical\" attribute is required by this class.")
	}
	var instance = &magnitude_{
		// Initialize the instance attributes.
		numerical_: numerical,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *magnitude_) GetClass() MagnitudeClassLike {
	return magnitudeClass()
}

// Attribute Methods

func (v *magnitude_) GetNumerical() NumericalLike {
	return v.numerical_
}

// PROTECTED INTERFACE

// Instance Structure

type magnitude_ struct {
	// Declare the instance attributes.
	numerical_ NumericalLike
}

// Class Structure

type magnitudeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func magnitudeClass() *magnitudeClass_ {
	return magnitudeClassReference_
}

var magnitudeClassReference_ = &magnitudeClass_{
	// Initialize the class constants.
}
