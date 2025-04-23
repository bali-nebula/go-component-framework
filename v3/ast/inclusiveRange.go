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

func InclusiveRangeClass() InclusiveRangeClassLike {
	return inclusiveRangeClass()
}

// Constructor Methods

func (c *inclusiveRangeClass_) InclusiveRange(
	primitive1 PrimitiveLike,
	primitive2 PrimitiveLike,
	bracket BracketLike,
) InclusiveRangeLike {
	if uti.IsUndefined(primitive1) {
		panic("The \"primitive1\" attribute is required by this class.")
	}
	if uti.IsUndefined(primitive2) {
		panic("The \"primitive2\" attribute is required by this class.")
	}
	if uti.IsUndefined(bracket) {
		panic("The \"bracket\" attribute is required by this class.")
	}
	var instance = &inclusiveRange_{
		// Initialize the instance attributes.
		primitive1_: primitive1,
		primitive2_: primitive2,
		bracket_:    bracket,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inclusiveRange_) GetClass() InclusiveRangeClassLike {
	return inclusiveRangeClass()
}

// Attribute Methods

func (v *inclusiveRange_) GetPrimitive1() PrimitiveLike {
	return v.primitive1_
}

func (v *inclusiveRange_) GetPrimitive2() PrimitiveLike {
	return v.primitive2_
}

func (v *inclusiveRange_) GetBracket() BracketLike {
	return v.bracket_
}

// PROTECTED INTERFACE

// Instance Structure

type inclusiveRange_ struct {
	// Declare the instance attributes.
	primitive1_ PrimitiveLike
	primitive2_ PrimitiveLike
	bracket_    BracketLike
}

// Class Structure

type inclusiveRangeClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inclusiveRangeClass() *inclusiveRangeClass_ {
	return inclusiveRangeClassReference_
}

var inclusiveRangeClassReference_ = &inclusiveRangeClass_{
	// Initialize the class constants.
}
