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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func InlineParametersClass() InlineParametersClassLike {
	return inlineParametersClass()
}

// Constructor Methods

func (c *inlineParametersClass_) InlineParameters(
	parameter ParameterLike,
	additionalParameters col.Sequential[AdditionalParameterLike],
) InlineParametersLike {
	if uti.IsUndefined(parameter) {
		panic("The \"parameter\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalParameters) {
		panic("The \"additionalParameters\" attribute is required by this class.")
	}
	var instance = &inlineParameters_{
		// Initialize the instance attributes.
		parameter_:            parameter,
		additionalParameters_: additionalParameters,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inlineParameters_) GetClass() InlineParametersClassLike {
	return inlineParametersClass()
}

// Attribute Methods

func (v *inlineParameters_) GetParameter() ParameterLike {
	return v.parameter_
}

func (v *inlineParameters_) GetAdditionalParameters() col.Sequential[AdditionalParameterLike] {
	return v.additionalParameters_
}

// PROTECTED INTERFACE

// Instance Structure

type inlineParameters_ struct {
	// Declare the instance attributes.
	parameter_            ParameterLike
	additionalParameters_ col.Sequential[AdditionalParameterLike]
}

// Class Structure

type inlineParametersClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineParametersClass() *inlineParametersClass_ {
	return inlineParametersClassReference_
}

var inlineParametersClassReference_ = &inlineParametersClass_{
	// Initialize the class constants.
}
