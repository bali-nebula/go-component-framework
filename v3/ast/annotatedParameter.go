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

func AnnotatedParameterClass() AnnotatedParameterClassLike {
	return annotatedParameterClass()
}

// Constructor Methods

func (c *annotatedParameterClass_) AnnotatedParameter(
	parameter ParameterLike,
	optionalNote string,
) AnnotatedParameterLike {
	if uti.IsUndefined(parameter) {
		panic("The \"parameter\" attribute is required by this class.")
	}
	var instance = &annotatedParameter_{
		// Initialize the instance attributes.
		parameter_:    parameter,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *annotatedParameter_) GetClass() AnnotatedParameterClassLike {
	return annotatedParameterClass()
}

// Attribute Methods

func (v *annotatedParameter_) GetParameter() ParameterLike {
	return v.parameter_
}

func (v *annotatedParameter_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type annotatedParameter_ struct {
	// Declare the instance attributes.
	parameter_    ParameterLike
	optionalNote_ string
}

// Class Structure

type annotatedParameterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func annotatedParameterClass() *annotatedParameterClass_ {
	return annotatedParameterClassReference_
}

var annotatedParameterClassReference_ = &annotatedParameterClass_{
	// Initialize the class constants.
}
