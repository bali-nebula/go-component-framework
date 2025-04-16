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

func InExclusiveProbabilitiesClass() InExclusiveProbabilitiesClassLike {
	return inExclusiveProbabilitiesClass()
}

// Constructor Methods

func (c *inExclusiveProbabilitiesClass_) InExclusiveProbabilities(
	probability1 string,
	probability2 string,
) InExclusiveProbabilitiesLike {
	if uti.IsUndefined(probability1) {
		panic("The \"probability1\" attribute is required by this class.")
	}
	if uti.IsUndefined(probability2) {
		panic("The \"probability2\" attribute is required by this class.")
	}
	var instance = &inExclusiveProbabilities_{
		// Initialize the instance attributes.
		probability1_: probability1,
		probability2_: probability2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inExclusiveProbabilities_) GetClass() InExclusiveProbabilitiesClassLike {
	return inExclusiveProbabilitiesClass()
}

// Attribute Methods

func (v *inExclusiveProbabilities_) GetProbability1() string {
	return v.probability1_
}

func (v *inExclusiveProbabilities_) GetProbability2() string {
	return v.probability2_
}

// PROTECTED INTERFACE

// Instance Structure

type inExclusiveProbabilities_ struct {
	// Declare the instance attributes.
	probability1_ string
	probability2_ string
}

// Class Structure

type inExclusiveProbabilitiesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inExclusiveProbabilitiesClass() *inExclusiveProbabilitiesClass_ {
	return inExclusiveProbabilitiesClassReference_
}

var inExclusiveProbabilitiesClassReference_ = &inExclusiveProbabilitiesClass_{
	// Initialize the class constants.
}
