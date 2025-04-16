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

func NoAttributesClass() NoAttributesClassLike {
	return noAttributesClass()
}

// Constructor Methods

func (c *noAttributesClass_) NoAttributes(
	colon string,
) NoAttributesLike {
	if uti.IsUndefined(colon) {
		panic("The \"colon\" attribute is required by this class.")
	}
	var instance = &noAttributes_{
		// Initialize the instance attributes.
		colon_: colon,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *noAttributes_) GetClass() NoAttributesClassLike {
	return noAttributesClass()
}

// Attribute Methods

func (v *noAttributes_) GetColon() string {
	return v.colon_
}

// PROTECTED INTERFACE

// Instance Structure

type noAttributes_ struct {
	// Declare the instance attributes.
	colon_ string
}

// Class Structure

type noAttributesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noAttributesClass() *noAttributesClass_ {
	return noAttributesClassReference_
}

var noAttributesClassReference_ = &noAttributesClass_{
	// Initialize the class constants.
}
