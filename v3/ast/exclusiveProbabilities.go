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

func ExclusiveProbabilitiesClass() ExclusiveProbabilitiesClassLike {
	return exclusiveProbabilitiesClass()
}

// Constructor Methods

func (c *exclusiveProbabilitiesClass_) ExclusiveProbabilities(
	probability1 string,
	probability2 string,
) ExclusiveProbabilitiesLike {
	if uti.IsUndefined(probability1) {
		panic("The \"probability1\" attribute is required by this class.")
	}
	if uti.IsUndefined(probability2) {
		panic("The \"probability2\" attribute is required by this class.")
	}
	var instance = &exclusiveProbabilities_{
		// Initialize the instance attributes.
		probability1_: probability1,
		probability2_: probability2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exclusiveProbabilities_) GetClass() ExclusiveProbabilitiesClassLike {
	return exclusiveProbabilitiesClass()
}

// Attribute Methods

func (v *exclusiveProbabilities_) GetProbability1() string {
	return v.probability1_
}

func (v *exclusiveProbabilities_) GetProbability2() string {
	return v.probability2_
}

// PROTECTED INTERFACE

// Instance Structure

type exclusiveProbabilities_ struct {
	// Declare the instance attributes.
	probability1_ string
	probability2_ string
}

// Class Structure

type exclusiveProbabilitiesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exclusiveProbabilitiesClass() *exclusiveProbabilitiesClass_ {
	return exclusiveProbabilitiesClassReference_
}

var exclusiveProbabilitiesClassReference_ = &exclusiveProbabilitiesClass_{
	// Initialize the class constants.
}
