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
	expression1 ExpressionLike,
	comparisonOperator ComparisonOperatorLike,
	expression2 ExpressionLike,
) ComparisonLike {
	if uti.IsUndefined(expression1) {
		panic("The \"expression1\" attribute is required by this class.")
	}
	if uti.IsUndefined(comparisonOperator) {
		panic("The \"comparisonOperator\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression2) {
		panic("The \"expression2\" attribute is required by this class.")
	}
	var instance = &comparison_{
		// Initialize the instance attributes.
		expression1_:        expression1,
		comparisonOperator_: comparisonOperator,
		expression2_:        expression2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *comparison_) GetClass() ComparisonClassLike {
	return comparisonClass()
}

// Attribute Methods

func (v *comparison_) GetExpression1() ExpressionLike {
	return v.expression1_
}

func (v *comparison_) GetComparisonOperator() ComparisonOperatorLike {
	return v.comparisonOperator_
}

func (v *comparison_) GetExpression2() ExpressionLike {
	return v.expression2_
}

// PROTECTED INTERFACE

// Instance Structure

type comparison_ struct {
	// Declare the instance attributes.
	expression1_        ExpressionLike
	comparisonOperator_ ComparisonOperatorLike
	expression2_        ExpressionLike
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
