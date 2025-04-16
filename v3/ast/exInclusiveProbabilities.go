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

func ExInclusiveProbabilitiesClass() ExInclusiveProbabilitiesClassLike {
	return exInclusiveProbabilitiesClass()
}

// Constructor Methods

func (c *exInclusiveProbabilitiesClass_) ExInclusiveProbabilities(
	probability1 string,
	probability2 string,
) ExInclusiveProbabilitiesLike {
	if uti.IsUndefined(probability1) {
		panic("The \"probability1\" attribute is required by this class.")
	}
	if uti.IsUndefined(probability2) {
		panic("The \"probability2\" attribute is required by this class.")
	}
	var instance = &exInclusiveProbabilities_{
		// Initialize the instance attributes.
		probability1_: probability1,
		probability2_: probability2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exInclusiveProbabilities_) GetClass() ExInclusiveProbabilitiesClassLike {
	return exInclusiveProbabilitiesClass()
}

// Attribute Methods

func (v *exInclusiveProbabilities_) GetProbability1() string {
	return v.probability1_
}

func (v *exInclusiveProbabilities_) GetProbability2() string {
	return v.probability2_
}

// PROTECTED INTERFACE

// Instance Structure

type exInclusiveProbabilities_ struct {
	// Declare the instance attributes.
	probability1_ string
	probability2_ string
}

// Class Structure

type exInclusiveProbabilitiesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exInclusiveProbabilitiesClass() *exInclusiveProbabilitiesClass_ {
	return exInclusiveProbabilitiesClassReference_
}

var exInclusiveProbabilitiesClassReference_ = &exInclusiveProbabilitiesClass_{
	// Initialize the class constants.
}
