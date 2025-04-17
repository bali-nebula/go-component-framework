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

func InExclusiveRangeClass() InExclusiveRangeClassLike {
	return inExclusiveRangeClass()
}

// Constructor Methods

func (c *inExclusiveRangeClass_) InExclusiveRange(
	bar1 string,
	primitive1 PrimitiveLike,
	primitive2 PrimitiveLike,
	bar2 string,
) InExclusiveRangeLike {
	if uti.IsUndefined(bar1) {
		panic("The \"bar1\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitive1) {
		panic("The \"primitive1\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitive2) {
		panic("The \"primitive2\" attribute is required by this class.")
	}
	if uti.IsUndefined(bar2) {
		panic("The \"bar2\" attribute is required by this class.")
	}
	var instance = &inExclusiveRange_{
		// Initialize the instance attributes.
		bar1_:       bar1,
		primitive1_: primitive1,
		primitive2_: primitive2,
		bar2_:       bar2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inExclusiveRange_) GetClass() InExclusiveRangeClassLike {
	return inExclusiveRangeClass()
}

// Attribute Methods

func (v *inExclusiveRange_) GetBar1() string {
	return v.bar1_
}

func (v *inExclusiveRange_) GetPrimitive1() PrimitiveLike {
	return v.primitive1_
}

func (v *inExclusiveRange_) GetPrimitive2() PrimitiveLike {
	return v.primitive2_
}

func (v *inExclusiveRange_) GetBar2() string {
	return v.bar2_
}

// PROTECTED INTERFACE

// Instance Structure

type inExclusiveRange_ struct {
	// Declare the instance attributes.
	bar1_       string
	primitive1_ PrimitiveLike
	primitive2_ PrimitiveLike
	bar2_       string
}

// Class Structure

type inExclusiveRangeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inExclusiveRangeClass() *inExclusiveRangeClass_ {
	return inExclusiveRangeClassReference_
}

var inExclusiveRangeClassReference_ = &inExclusiveRangeClass_{
	// Initialize the class constants.
}
