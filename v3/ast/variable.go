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

func VariableClass() VariableClassLike {
	return variableClass()
}

// Constructor Methods

func (c *variableClass_) Variable(
	identifier string,
	optionalSubcomponent SubcomponentLike,
) VariableLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	var instance = &variable_{
		// Initialize the instance attributes.
		identifier_:           identifier,
		optionalSubcomponent_: optionalSubcomponent,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *variable_) GetClass() VariableClassLike {
	return variableClass()
}

// Attribute Methods

func (v *variable_) GetIdentifier() string {
	return v.identifier_
}

func (v *variable_) GetOptionalSubcomponent() SubcomponentLike {
	return v.optionalSubcomponent_
}

// PROTECTED INTERFACE

// Instance Structure

type variable_ struct {
	// Declare the instance attributes.
	identifier_           string
	optionalSubcomponent_ SubcomponentLike
}

// Class Structure

type variableClass_ struct {
	// Declare the class constants.
}

// Class Reference

func variableClass() *variableClass_ {
	return variableClassReference_
}

var variableClassReference_ = &variableClass_{
	// Initialize the class constants.
}
