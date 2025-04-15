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

func AmplitudeClass() AmplitudeClassLike {
	return amplitudeClass()
}

// Constructor Methods

func (c *amplitudeClass_) Amplitude(
	bar1 string,
	numerical NumericalLike,
	bar2 string,
) AmplitudeLike {
	if uti.IsUndefined(bar1) {
		panic("The \"bar1\" attribute is required by this class.")
	}
	if uti.IsUndefined(numerical) {
		panic("The \"numerical\" attribute is required by this class.")
	}
	if uti.IsUndefined(bar2) {
		panic("The \"bar2\" attribute is required by this class.")
	}
	var instance = &amplitude_{
		// Initialize the instance attributes.
		bar1_:      bar1,
		numerical_: numerical,
		bar2_:      bar2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *amplitude_) GetClass() AmplitudeClassLike {
	return amplitudeClass()
}

// Attribute Methods

func (v *amplitude_) GetBar1() string {
	return v.bar1_
}

func (v *amplitude_) GetNumerical() NumericalLike {
	return v.numerical_
}

func (v *amplitude_) GetBar2() string {
	return v.bar2_
}

// PROTECTED INTERFACE

// Instance Structure

type amplitude_ struct {
	// Declare the instance attributes.
	bar1_      string
	numerical_ NumericalLike
	bar2_      string
}

// Class Structure

type amplitudeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func amplitudeClass() *amplitudeClass_ {
	return amplitudeClassReference_
}

var amplitudeClassReference_ = &amplitudeClass_{
	// Initialize the class constants.
}
