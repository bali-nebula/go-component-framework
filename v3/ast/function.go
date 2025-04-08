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

func FunctionClass() FunctionClassLike {
	return functionClass()
}

// Constructor Methods

func (c *functionClass_) Function(
	identifier string,
) FunctionLike {
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	var instance = &function_{
		// Initialize the instance attributes.
		identifier_: identifier,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *function_) GetClass() FunctionClassLike {
	return functionClass()
}

// Attribute Methods

func (v *function_) GetIdentifier() string {
	return v.identifier_
}

// PROTECTED INTERFACE

// Instance Structure

type function_ struct {
	// Declare the instance attributes.
	identifier_ string
}

// Class Structure

type functionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func functionClass() *functionClass_ {
	return functionClassReference_
}

var functionClassReference_ = &functionClass_{
	// Initialize the class constants.
}
