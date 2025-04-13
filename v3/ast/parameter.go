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

func ParameterClass() ParameterClassLike {
	return parameterClass()
}

// Constructor Methods

func (c *parameterClass_) Parameter(
	label LabelLike,
	colon string,
	component ComponentLike,
) ParameterLike {
	if uti.IsUndefined(label) {
		panic("The \"label\" attribute is required by this class.")
	}
	if uti.IsUndefined(colon) {
		panic("The \"colon\" attribute is required by this class.")
	}
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &parameter_{
		// Initialize the instance attributes.
		label_:     label,
		colon_:     colon,
		component_: component,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parameter_) GetClass() ParameterClassLike {
	return parameterClass()
}

// Attribute Methods

func (v *parameter_) GetLabel() LabelLike {
	return v.label_
}

func (v *parameter_) GetColon() string {
	return v.colon_
}

func (v *parameter_) GetComponent() ComponentLike {
	return v.component_
}

// PROTECTED INTERFACE

// Instance Structure

type parameter_ struct {
	// Declare the instance attributes.
	label_     LabelLike
	colon_     string
	component_ ComponentLike
}

// Class Structure

type parameterClass_ struct {
	// Declare the class constants.
}

// Class Reference

func parameterClass() *parameterClass_ {
	return parameterClassReference_
}

var parameterClassReference_ = &parameterClass_{
	// Initialize the class constants.
}
