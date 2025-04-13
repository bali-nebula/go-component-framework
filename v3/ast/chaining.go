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

func ChainingClass() ChainingClassLike {
	return chainingClass()
}

// Constructor Methods

func (c *chainingClass_) Chaining(
	expression1 ExpressionLike,
	ampersand string,
	expression2 ExpressionLike,
) ChainingLike {
	if uti.IsUndefined(expression1) {
		panic("The \"expression1\" attribute is required by this class.")
	}
	if uti.IsUndefined(ampersand) {
		panic("The \"ampersand\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression2) {
		panic("The \"expression2\" attribute is required by this class.")
	}
	var instance = &chaining_{
		// Initialize the instance attributes.
		expression1_: expression1,
		ampersand_:   ampersand,
		expression2_: expression2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *chaining_) GetClass() ChainingClassLike {
	return chainingClass()
}

// Attribute Methods

func (v *chaining_) GetExpression1() ExpressionLike {
	return v.expression1_
}

func (v *chaining_) GetAmpersand() string {
	return v.ampersand_
}

func (v *chaining_) GetExpression2() ExpressionLike {
	return v.expression2_
}

// PROTECTED INTERFACE

// Instance Structure

type chaining_ struct {
	// Declare the instance attributes.
	expression1_ ExpressionLike
	ampersand_   string
	expression2_ ExpressionLike
}

// Class Structure

type chainingClass_ struct {
	// Declare the class constants.
}

// Class Reference

func chainingClass() *chainingClass_ {
	return chainingClassReference_
}

var chainingClassReference_ = &chainingClass_{
	// Initialize the class constants.
}
