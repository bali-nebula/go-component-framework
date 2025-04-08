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

func ExceptionClass() ExceptionClassLike {
	return exceptionClass()
}

// Constructor Methods

func (c *exceptionClass_) Exception(
	expression ExpressionLike,
) ExceptionLike {
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &exception_{
		// Initialize the instance attributes.
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exception_) GetClass() ExceptionClassLike {
	return exceptionClass()
}

// Attribute Methods

func (v *exception_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type exception_ struct {
	// Declare the instance attributes.
	expression_ ExpressionLike
}

// Class Structure

type exceptionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exceptionClass() *exceptionClass_ {
	return exceptionClassReference_
}

var exceptionClassReference_ = &exceptionClass_{
	// Initialize the class constants.
}
