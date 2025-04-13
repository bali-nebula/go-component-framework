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

func ComplementClass() ComplementClassLike {
	return complementClass()
}

// Constructor Methods

func (c *complementClass_) Complement(
	not string,
	expression ExpressionLike,
) ComplementLike {
	if uti.IsUndefined(not) {
		panic("The \"not\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &complement_{
		// Initialize the instance attributes.
		not_:        not,
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *complement_) GetClass() ComplementClassLike {
	return complementClass()
}

// Attribute Methods

func (v *complement_) GetNot() string {
	return v.not_
}

func (v *complement_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type complement_ struct {
	// Declare the instance attributes.
	not_        string
	expression_ ExpressionLike
}

// Class Structure

type complementClass_ struct {
	// Declare the class constants.
}

// Class Reference

func complementClass() *complementClass_ {
	return complementClassReference_
}

var complementClassReference_ = &complementClass_{
	// Initialize the class constants.
}
