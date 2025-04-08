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

func InlineValuesClass() InlineValuesClassLike {
	return inlineValuesClass()
}

// Constructor Methods

func (c *inlineValuesClass_) InlineValues(
	value ValueLike,
	additionalValues col.Sequential[AdditionalValueLike],
) InlineValuesLike {
	if uti.IsUndefined(value) {
		panic("The \"value\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalValues) {
		panic("The \"additionalValues\" attribute is required by this class.")
	}
	var instance = &inlineValues_{
		// Initialize the instance attributes.
		value_:            value,
		additionalValues_: additionalValues,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inlineValues_) GetClass() InlineValuesClassLike {
	return inlineValuesClass()
}

// Attribute Methods

func (v *inlineValues_) GetValue() ValueLike {
	return v.value_
}

func (v *inlineValues_) GetAdditionalValues() col.Sequential[AdditionalValueLike] {
	return v.additionalValues_
}

// PROTECTED INTERFACE

// Instance Structure

type inlineValues_ struct {
	// Declare the instance attributes.
	value_            ValueLike
	additionalValues_ col.Sequential[AdditionalValueLike]
}

// Class Structure

type inlineValuesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineValuesClass() *inlineValuesClass_ {
	return inlineValuesClassReference_
}

var inlineValuesClassReference_ = &inlineValuesClass_{
	// Initialize the class constants.
}
