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

func ConcatenationClass() ConcatenationClassLike {
	return concatenationClass()
}

// Constructor Methods

func (c *concatenationClass_) Concatenation(
	textual TextualLike,
	concatenationOperations col.Sequential[ConcatenationOperationLike],
) ConcatenationLike {
	if uti.IsUndefined(textual) {
		panic("The \"textual\" attribute is required by this class.")
	}
	if uti.IsUndefined(concatenationOperations) {
		panic("The \"concatenationOperations\" attribute is required by this class.")
	}
	var instance = &concatenation_{
		// Initialize the instance attributes.
		textual_:                 textual,
		concatenationOperations_: concatenationOperations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *concatenation_) GetClass() ConcatenationClassLike {
	return concatenationClass()
}

// Attribute Methods

func (v *concatenation_) GetTextual() TextualLike {
	return v.textual_
}

func (v *concatenation_) GetConcatenationOperations() col.Sequential[ConcatenationOperationLike] {
	return v.concatenationOperations_
}

// PROTECTED INTERFACE

// Instance Structure

type concatenation_ struct {
	// Declare the instance attributes.
	textual_                 TextualLike
	concatenationOperations_ col.Sequential[ConcatenationOperationLike]
}

// Class Structure

type concatenationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func concatenationClass() *concatenationClass_ {
	return concatenationClassReference_
}

var concatenationClassReference_ = &concatenationClass_{
	// Initialize the class constants.
}
