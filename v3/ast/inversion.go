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

func InversionClass() InversionClassLike {
	return inversionClass()
}

// Constructor Methods

func (c *inversionClass_) Inversion(
	inverse InverseLike,
	numerical NumericalLike,
) InversionLike {
	if uti.IsUndefined(inverse) {
		panic("The \"inverse\" attribute is required by this class.")
	}
	if uti.IsUndefined(numerical) {
		panic("The \"numerical\" attribute is required by this class.")
	}
	var instance = &inversion_{
		// Initialize the instance attributes.
		inverse_:   inverse,
		numerical_: numerical,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inversion_) GetClass() InversionClassLike {
	return inversionClass()
}

// Attribute Methods

func (v *inversion_) GetInverse() InverseLike {
	return v.inverse_
}

func (v *inversion_) GetNumerical() NumericalLike {
	return v.numerical_
}

// PROTECTED INTERFACE

// Instance Structure

type inversion_ struct {
	// Declare the instance attributes.
	inverse_   InverseLike
	numerical_ NumericalLike
}

// Class Structure

type inversionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inversionClass() *inversionClass_ {
	return inversionClassReference_
}

var inversionClassReference_ = &inversionClass_{
	// Initialize the class constants.
}
