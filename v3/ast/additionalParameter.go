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

func AdditionalParameterClass() AdditionalParameterClassLike {
	return additionalParameterClass()
}

// Constructor Methods

func (c *additionalParameterClass_) AdditionalParameter(
	parameter ParameterLike,
) AdditionalParameterLike {
	if uti.IsUndefined(parameter) {
		panic("The \"parameter\" attribute is required by this class.")
	}
	var instance = &additionalParameter_{
		// Initialize the instance attributes.
		parameter_: parameter,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalParameter_) GetClass() AdditionalParameterClassLike {
	return additionalParameterClass()
}

// Attribute Methods

func (v *additionalParameter_) GetParameter() ParameterLike {
	return v.parameter_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalParameter_ struct {
	// Declare the instance attributes.
	parameter_ ParameterLike
}

// Class Structure

type additionalParameterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalParameterClass() *additionalParameterClass_ {
	return additionalParameterClassReference_
}

var additionalParameterClassReference_ = &additionalParameterClass_{
	// Initialize the class constants.
}
