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

func PrecedenceClass() PrecedenceClassLike {
	return precedenceClass()
}

// Constructor Methods

func (c *precedenceClass_) Precedence(
	expression ExpressionLike,
) PrecedenceLike {
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &precedence_{
		// Initialize the instance attributes.
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *precedence_) GetClass() PrecedenceClassLike {
	return precedenceClass()
}

// Attribute Methods

func (v *precedence_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type precedence_ struct {
	// Declare the instance attributes.
	expression_ ExpressionLike
}

// Class Structure

type precedenceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func precedenceClass() *precedenceClass_ {
	return precedenceClassReference_
}

var precedenceClassReference_ = &precedenceClass_{
	// Initialize the class constants.
}
