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

func ExInclusiveRangeClass() ExInclusiveRangeClassLike {
	return exInclusiveRangeClass()
}

// Constructor Methods

func (c *exInclusiveRangeClass_) ExInclusiveRange(
	bar1 string,
	primitive1 PrimitiveLike,
	primitive2 PrimitiveLike,
	bar2 string,
) ExInclusiveRangeLike {
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
	var instance = &exInclusiveRange_{
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

func (v *exInclusiveRange_) GetClass() ExInclusiveRangeClassLike {
	return exInclusiveRangeClass()
}

// Attribute Methods

func (v *exInclusiveRange_) GetBar1() string {
	return v.bar1_
}

func (v *exInclusiveRange_) GetPrimitive1() PrimitiveLike {
	return v.primitive1_
}

func (v *exInclusiveRange_) GetPrimitive2() PrimitiveLike {
	return v.primitive2_
}

func (v *exInclusiveRange_) GetBar2() string {
	return v.bar2_
}

// PROTECTED INTERFACE

// Instance Structure

type exInclusiveRange_ struct {
	// Declare the instance attributes.
	bar1_       string
	primitive1_ PrimitiveLike
	primitive2_ PrimitiveLike
	bar2_       string
}

// Class Structure

type exInclusiveRangeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exInclusiveRangeClass() *exInclusiveRangeClass_ {
	return exInclusiveRangeClassReference_
}

var exInclusiveRangeClassReference_ = &exInclusiveRangeClass_{
	// Initialize the class constants.
}
