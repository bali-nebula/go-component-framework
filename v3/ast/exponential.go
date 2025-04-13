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

func ExponentialClass() ExponentialClassLike {
	return exponentialClass()
}

// Constructor Methods

func (c *exponentialClass_) Exponential(
	expression1 ExpressionLike,
	caret string,
	expression2 ExpressionLike,
) ExponentialLike {
	if uti.IsUndefined(expression1) {
		panic("The \"expression1\" attribute is required by this class.")
	}
	if uti.IsUndefined(caret) {
		panic("The \"caret\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression2) {
		panic("The \"expression2\" attribute is required by this class.")
	}
	var instance = &exponential_{
		// Initialize the instance attributes.
		expression1_: expression1,
		caret_:       caret,
		expression2_: expression2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exponential_) GetClass() ExponentialClassLike {
	return exponentialClass()
}

// Attribute Methods

func (v *exponential_) GetExpression1() ExpressionLike {
	return v.expression1_
}

func (v *exponential_) GetCaret() string {
	return v.caret_
}

func (v *exponential_) GetExpression2() ExpressionLike {
	return v.expression2_
}

// PROTECTED INTERFACE

// Instance Structure

type exponential_ struct {
	// Declare the instance attributes.
	expression1_ ExpressionLike
	caret_       string
	expression2_ ExpressionLike
}

// Class Structure

type exponentialClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exponentialClass() *exponentialClass_ {
	return exponentialClassReference_
}

var exponentialClassReference_ = &exponentialClass_{
	// Initialize the class constants.
}
