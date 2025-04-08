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

func ContextClass() ContextClassLike {
	return contextClass()
}

// Constructor Methods

func (c *contextClass_) Context(
	parameters ParametersLike,
) ContextLike {
	if uti.IsUndefined(parameters) {
		panic("The \"parameters\" attribute is required by this class.")
	}
	var instance = &context_{
		// Initialize the instance attributes.
		parameters_: parameters,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *context_) GetClass() ContextClassLike {
	return contextClass()
}

// Attribute Methods

func (v *context_) GetParameters() ParametersLike {
	return v.parameters_
}

// PROTECTED INTERFACE

// Instance Structure

type context_ struct {
	// Declare the instance attributes.
	parameters_ ParametersLike
}

// Class Structure

type contextClass_ struct {
	// Declare the class constants.
}

// Class Reference

func contextClass() *contextClass_ {
	return contextClassReference_
}

var contextClassReference_ = &contextClass_{
	// Initialize the class constants.
}
