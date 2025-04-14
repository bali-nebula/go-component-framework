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

func ParametersClass() ParametersClassLike {
	return parametersClass()
}

// Constructor Methods

func (c *parametersClass_) Parameters(
	associations AssociationsLike,
) ParametersLike {
	if uti.IsUndefined(associations) {
		panic("The \"associations\" attribute is required by this class.")
	}
	var instance = &parameters_{
		// Initialize the instance attributes.
		associations_: associations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *parameters_) GetClass() ParametersClassLike {
	return parametersClass()
}

// Attribute Methods

func (v *parameters_) GetAssociations() AssociationsLike {
	return v.associations_
}

// PROTECTED INTERFACE

// Instance Structure

type parameters_ struct {
	// Declare the instance attributes.
	associations_ AssociationsLike
}

// Class Structure

type parametersClass_ struct {
	// Declare the class constants.
}

// Class Reference

func parametersClass() *parametersClass_ {
	return parametersClassReference_
}

var parametersClassReference_ = &parametersClass_{
	// Initialize the class constants.
}
