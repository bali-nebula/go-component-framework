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
	leftBracket LeftBracketLike,
	primitive1 PrimitiveLike,
	primitive2 PrimitiveLike,
	rightBracket RightBracketLike,
) RangeLike {
	if uti.IsUndefined(leftBracket) {
		panic("The \"leftBracket\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitive1) {
		panic("The \"primitive1\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitive2) {
		panic("The \"primitive2\" attribute is required by this class.")
	}
	if uti.IsUndefined(rightBracket) {
		panic("The \"rightBracket\" attribute is required by this class.")
	}
	var instance = &range_{
		// Initialize the instance attributes.
		leftBracket_:  leftBracket,
		primitive1_:   primitive1,
		primitive2_:   primitive2,
		rightBracket_: rightBracket,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *range_) GetClass() RangeClassLike {
	return rangeClass()
}

// Attribute Methods

func (v *range_) GetLeftBracket() LeftBracketLike {
	return v.leftBracket_
}

func (v *range_) GetPrimitive1() PrimitiveLike {
	return v.primitive1_
}

func (v *range_) GetPrimitive2() PrimitiveLike {
	return v.primitive2_
}

func (v *range_) GetRightBracket() RightBracketLike {
	return v.rightBracket_
}

// PROTECTED INTERFACE

// Instance Structure

type range_ struct {
	// Declare the instance attributes.
	leftBracket_  LeftBracketLike
	primitive1_   PrimitiveLike
	primitive2_   PrimitiveLike
	rightBracket_ RightBracketLike
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
