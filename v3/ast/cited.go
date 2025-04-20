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

func CitedClass() CitedClassLike {
	return citedClass()
}

// Constructor Methods

func (c *citedClass_) Cited(
	expression ExpressionLike,
) CitedLike {
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &cited_{
		// Initialize the instance attributes.
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *cited_) GetClass() CitedClassLike {
	return citedClass()
}

// Attribute Methods

func (v *cited_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type cited_ struct {
	// Declare the instance attributes.
	expression_ ExpressionLike
}

// Class Structure

type citedClass_ struct {
	// Declare the class constants.
}

// Class Reference

func citedClass() *citedClass_ {
	return citedClassReference_
}

var citedClassReference_ = &citedClass_{
	// Initialize the class constants.
}
