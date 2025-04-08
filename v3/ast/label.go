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

func LabelClass() LabelClassLike {
	return labelClass()
}

// Constructor Methods

func (c *labelClass_) Label(
	symbol string,
) LabelLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &label_{
		// Initialize the instance attributes.
		symbol_: symbol,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *label_) GetClass() LabelClassLike {
	return labelClass()
}

// Attribute Methods

func (v *label_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type label_ struct {
	// Declare the instance attributes.
	symbol_ string
}

// Class Structure

type labelClass_ struct {
	// Declare the class constants.
}

// Class Reference

func labelClass() *labelClass_ {
	return labelClassReference_
}

var labelClassReference_ = &labelClass_{
	// Initialize the class constants.
}
