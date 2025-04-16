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

func InclusiveProbabilitiesClass() InclusiveProbabilitiesClassLike {
	return inclusiveProbabilitiesClass()
}

// Constructor Methods

func (c *inclusiveProbabilitiesClass_) InclusiveProbabilities(
	probability1 string,
	probability2 string,
) InclusiveProbabilitiesLike {
	if uti.IsUndefined(probability1) {
		panic("The \"probability1\" attribute is required by this class.")
	}
	if uti.IsUndefined(probability2) {
		panic("The \"probability2\" attribute is required by this class.")
	}
	var instance = &inclusiveProbabilities_{
		// Initialize the instance attributes.
		probability1_: probability1,
		probability2_: probability2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inclusiveProbabilities_) GetClass() InclusiveProbabilitiesClassLike {
	return inclusiveProbabilitiesClass()
}

// Attribute Methods

func (v *inclusiveProbabilities_) GetProbability1() string {
	return v.probability1_
}

func (v *inclusiveProbabilities_) GetProbability2() string {
	return v.probability2_
}

// PROTECTED INTERFACE

// Instance Structure

type inclusiveProbabilities_ struct {
	// Declare the instance attributes.
	probability1_ string
	probability2_ string
}

// Class Structure

type inclusiveProbabilitiesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inclusiveProbabilitiesClass() *inclusiveProbabilitiesClass_ {
	return inclusiveProbabilitiesClassReference_
}

var inclusiveProbabilitiesClassReference_ = &inclusiveProbabilitiesClass_{
	// Initialize the class constants.
}
