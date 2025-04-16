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

func InclusivePercentagesClass() InclusivePercentagesClassLike {
	return inclusivePercentagesClass()
}

// Constructor Methods

func (c *inclusivePercentagesClass_) InclusivePercentages(
	percentage1 string,
	percentage2 string,
) InclusivePercentagesLike {
	if uti.IsUndefined(percentage1) {
		panic("The \"percentage1\" attribute is required by this class.")
	}
	if uti.IsUndefined(percentage2) {
		panic("The \"percentage2\" attribute is required by this class.")
	}
	var instance = &inclusivePercentages_{
		// Initialize the instance attributes.
		percentage1_: percentage1,
		percentage2_: percentage2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inclusivePercentages_) GetClass() InclusivePercentagesClassLike {
	return inclusivePercentagesClass()
}

// Attribute Methods

func (v *inclusivePercentages_) GetPercentage1() string {
	return v.percentage1_
}

func (v *inclusivePercentages_) GetPercentage2() string {
	return v.percentage2_
}

// PROTECTED INTERFACE

// Instance Structure

type inclusivePercentages_ struct {
	// Declare the instance attributes.
	percentage1_ string
	percentage2_ string
}

// Class Structure

type inclusivePercentagesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inclusivePercentagesClass() *inclusivePercentagesClass_ {
	return inclusivePercentagesClassReference_
}

var inclusivePercentagesClassReference_ = &inclusivePercentagesClass_{
	// Initialize the class constants.
}
