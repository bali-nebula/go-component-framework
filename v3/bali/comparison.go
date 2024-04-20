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

var comparisonClass = &comparisonClass_{
	// This class has no private constants to initialize.
}

// Function

func Comparison() ComparisonClassLike {
	return comparisonClass
}

// CLASS METHODS

// Target

type comparisonClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *comparisonClass_) MakeWithExpressions(expressions col.ListLike[ExpressionLike]) ComparisonLike {
	return &comparison_{
		expressions_: expressions,
	}
}

// Functions

// INSTANCE METHODS

// Target

type comparison_ struct {
	expressions_ col.ListLike[ExpressionLike]
}

// Attributes

func (v *comparison_) GetExpressions() col.ListLike[ExpressionLike] {
	return v.expressions_
}

// Public

// Private
