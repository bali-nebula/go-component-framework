/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package bali

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var rangeClass = &rangeClass_{
	// This class has no private constants to initialize.
}

// Function

func Range() RangeClassLike {
	return rangeClass
}

// CLASS METHODS

// Target

type rangeClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *rangeClass_) MakeWithPrimitives(primitives col.ListLike[PrimitiveLike]) RangeLike {
	return &range_{
		primitives_: primitives,
	}
}

// Functions

// INSTANCE METHODS

// Target

type range_ struct {
	primitives_ col.ListLike[PrimitiveLike]
}

// Attributes

func (v *range_) GetPrimitives() col.ListLike[PrimitiveLike] {
	return v.primitives_
}

// Public

// Private
