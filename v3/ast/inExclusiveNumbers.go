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

func InExclusiveNumbersClass() InExclusiveNumbersClassLike {
	return inExclusiveNumbersClass()
}

// Constructor Methods

func (c *inExclusiveNumbersClass_) InExclusiveNumbers(
	number1 string,
	number2 string,
) InExclusiveNumbersLike {
	if uti.IsUndefined(number1) {
		panic("The \"number1\" attribute is required by this class.")
	}
	if uti.IsUndefined(number2) {
		panic("The \"number2\" attribute is required by this class.")
	}
	var instance = &inExclusiveNumbers_{
		// Initialize the instance attributes.
		number1_: number1,
		number2_: number2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inExclusiveNumbers_) GetClass() InExclusiveNumbersClassLike {
	return inExclusiveNumbersClass()
}

// Attribute Methods

func (v *inExclusiveNumbers_) GetNumber1() string {
	return v.number1_
}

func (v *inExclusiveNumbers_) GetNumber2() string {
	return v.number2_
}

// PROTECTED INTERFACE

// Instance Structure

type inExclusiveNumbers_ struct {
	// Declare the instance attributes.
	number1_ string
	number2_ string
}

// Class Structure

type inExclusiveNumbersClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inExclusiveNumbersClass() *inExclusiveNumbersClass_ {
	return inExclusiveNumbersClassReference_
}

var inExclusiveNumbersClassReference_ = &inExclusiveNumbersClass_{
	// Initialize the class constants.
}
