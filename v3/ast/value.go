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

func ValueClass() ValueClassLike {
	return valueClass()
}

// Constructor Methods

func (c *valueClass_) Value(
	component ComponentLike,
) ValueLike {
	if uti.IsUndefined(component) {
		panic("The \"component\" attribute is required by this class.")
	}
	var instance = &value_{
		// Initialize the instance attributes.
		component_: component,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *value_) GetClass() ValueClassLike {
	return valueClass()
}

// Attribute Methods

func (v *value_) GetComponent() ComponentLike {
	return v.component_
}

// PROTECTED INTERFACE

// Instance Structure

type value_ struct {
	// Declare the instance attributes.
	component_ ComponentLike
}

// Class Structure

type valueClass_ struct {
	// Declare the class constants.
}

// Class Reference

func valueClass() *valueClass_ {
	return valueClassReference_
}

var valueClassReference_ = &valueClass_{
	// Initialize the class constants.
}
