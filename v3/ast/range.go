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

func RangeClass() RangeClassLike {
	return rangeClass()
}

// Constructor Methods

func (c *rangeClass_) Range(
	open string,
	primitive1 PrimitiveLike,
	primitive2 PrimitiveLike,
	close_ string,
) RangeLike {
	if uti.IsUndefined(open) {
		panic("The \"open\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitive1) {
		panic("The \"primitive1\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitive2) {
		panic("The \"primitive2\" attribute is required by this class.")
	}
	if uti.IsUndefined(close_) {
		panic("The \"close\" attribute is required by this class.")
	}
	var instance = &range_{
		// Initialize the instance attributes.
		open_:       open,
		primitive1_: primitive1,
		primitive2_: primitive2,
		close_:      close_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *range_) GetClass() RangeClassLike {
	return rangeClass()
}

// Attribute Methods

func (v *range_) GetOpen() string {
	return v.open_
}

func (v *range_) GetPrimitive1() PrimitiveLike {
	return v.primitive1_
}

func (v *range_) GetPrimitive2() PrimitiveLike {
	return v.primitive2_
}

func (v *range_) GetClose() string {
	return v.close_
}

// PROTECTED INTERFACE

// Instance Structure

type range_ struct {
	// Declare the instance attributes.
	open_       string
	primitive1_ PrimitiveLike
	primitive2_ PrimitiveLike
	close_      string
}

// Class Structure

type rangeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func rangeClass() *rangeClass_ {
	return rangeClassReference_
}

var rangeClassReference_ = &rangeClass_{
	// Initialize the class constants.
}
