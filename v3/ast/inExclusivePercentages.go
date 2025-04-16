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

func InExclusivePercentagesClass() InExclusivePercentagesClassLike {
	return inExclusivePercentagesClass()
}

// Constructor Methods

func (c *inExclusivePercentagesClass_) InExclusivePercentages(
	percentage1 string,
	percentage2 string,
) InExclusivePercentagesLike {
	if uti.IsUndefined(percentage1) {
		panic("The \"percentage1\" attribute is required by this class.")
	}
	if uti.IsUndefined(percentage2) {
		panic("The \"percentage2\" attribute is required by this class.")
	}
	var instance = &inExclusivePercentages_{
		// Initialize the instance attributes.
		percentage1_: percentage1,
		percentage2_: percentage2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inExclusivePercentages_) GetClass() InExclusivePercentagesClassLike {
	return inExclusivePercentagesClass()
}

// Attribute Methods

func (v *inExclusivePercentages_) GetPercentage1() string {
	return v.percentage1_
}

func (v *inExclusivePercentages_) GetPercentage2() string {
	return v.percentage2_
}

// PROTECTED INTERFACE

// Instance Structure

type inExclusivePercentages_ struct {
	// Declare the instance attributes.
	percentage1_ string
	percentage2_ string
}

// Class Structure

type inExclusivePercentagesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inExclusivePercentagesClass() *inExclusivePercentagesClass_ {
	return inExclusivePercentagesClassReference_
}

var inExclusivePercentagesClassReference_ = &inExclusivePercentagesClass_{
	// Initialize the class constants.
}
