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

func ConcatenationOperationClass() ConcatenationOperationClassLike {
	return concatenationOperationClass()
}

// Constructor Methods

func (c *concatenationOperationClass_) ConcatenationOperation(
	ampersand string,
	textual TextualLike,
) ConcatenationOperationLike {
	if uti.IsUndefined(ampersand) {
		panic("The \"ampersand\" attribute is required by this class.")
	}
	if uti.IsUndefined(textual) {
		panic("The \"textual\" attribute is required by this class.")
	}
	var instance = &concatenationOperation_{
		// Initialize the instance attributes.
		ampersand_: ampersand,
		textual_:   textual,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *concatenationOperation_) GetClass() ConcatenationOperationClassLike {
	return concatenationOperationClass()
}

// Attribute Methods

func (v *concatenationOperation_) GetAmpersand() string {
	return v.ampersand_
}

func (v *concatenationOperation_) GetTextual() TextualLike {
	return v.textual_
}

// PROTECTED INTERFACE

// Instance Structure

type concatenationOperation_ struct {
	// Declare the instance attributes.
	ampersand_ string
	textual_   TextualLike
}

// Class Structure

type concatenationOperationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func concatenationOperationClass() *concatenationOperationClass_ {
	return concatenationOperationClassReference_
}

var concatenationOperationClassReference_ = &concatenationOperationClass_{
	// Initialize the class constants.
}
