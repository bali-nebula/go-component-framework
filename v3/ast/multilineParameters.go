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

func MultilineParametersClass() MultilineParametersClassLike {
	return multilineParametersClass()
}

// Constructor Methods

func (c *multilineParametersClass_) MultilineParameters(
	annotatedParameters col.Sequential[AnnotatedParameterLike],
) MultilineParametersLike {
	if uti.IsUndefined(annotatedParameters) {
		panic("The \"annotatedParameters\" attribute is required by this class.")
	}
	var instance = &multilineParameters_{
		// Initialize the instance attributes.
		annotatedParameters_: annotatedParameters,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multilineParameters_) GetClass() MultilineParametersClassLike {
	return multilineParametersClass()
}

// Attribute Methods

func (v *multilineParameters_) GetAnnotatedParameters() col.Sequential[AnnotatedParameterLike] {
	return v.annotatedParameters_
}

// PROTECTED INTERFACE

// Instance Structure

type multilineParameters_ struct {
	// Declare the instance attributes.
	annotatedParameters_ col.Sequential[AnnotatedParameterLike]
}

// Class Structure

type multilineParametersClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multilineParametersClass() *multilineParametersClass_ {
	return multilineParametersClassReference_
}

var multilineParametersClassReference_ = &multilineParametersClass_{
	// Initialize the class constants.
}
