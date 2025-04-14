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

func ComparisonClass() ComparisonClassLike {
	return comparisonClass()
}

// Constructor Methods

func (c *comparisonClass_) Comparison(
	numerical1 NumericalLike,
	compareOperator CompareOperatorLike,
	numerical2 NumericalLike,
) ComparisonLike {
	if uti.IsUndefined(numerical1) {
		panic("The \"numerical1\" attribute is required by this class.")
	}
	if uti.IsUndefined(compareOperator) {
		panic("The \"compareOperator\" attribute is required by this class.")
	}
	if uti.IsUndefined(numerical2) {
		panic("The \"numerical2\" attribute is required by this class.")
	}
	var instance = &comparison_{
		// Initialize the instance attributes.
		numerical1_:      numerical1,
		compareOperator_: compareOperator,
		numerical2_:      numerical2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *comparison_) GetClass() ComparisonClassLike {
	return comparisonClass()
}

// Attribute Methods

func (v *comparison_) GetNumerical1() NumericalLike {
	return v.numerical1_
}

func (v *comparison_) GetCompareOperator() CompareOperatorLike {
	return v.compareOperator_
}

func (v *comparison_) GetNumerical2() NumericalLike {
	return v.numerical2_
}

// PROTECTED INTERFACE

// Instance Structure

type comparison_ struct {
	// Declare the instance attributes.
	numerical1_      NumericalLike
	compareOperator_ CompareOperatorLike
	numerical2_      NumericalLike
}

// Class Structure

type comparisonClass_ struct {
	// Declare the class constants.
}

// Class Reference

func comparisonClass() *comparisonClass_ {
	return comparisonClassReference_
}

var comparisonClassReference_ = &comparisonClass_{
	// Initialize the class constants.
}
