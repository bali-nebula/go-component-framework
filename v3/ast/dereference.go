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

func DereferenceClass() DereferenceClassLike {
	return dereferenceClass()
}

// Constructor Methods

func (c *dereferenceClass_) Dereference(
	snail string,
	expression ExpressionLike,
) DereferenceLike {
	if uti.IsUndefined(snail) {
		panic("The \"snail\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &dereference_{
		// Initialize the instance attributes.
		snail_:      snail,
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *dereference_) GetClass() DereferenceClassLike {
	return dereferenceClass()
}

// Attribute Methods

func (v *dereference_) GetSnail() string {
	return v.snail_
}

func (v *dereference_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type dereference_ struct {
	// Declare the instance attributes.
	snail_      string
	expression_ ExpressionLike
}

// Class Structure

type dereferenceClass_ struct {
	// Declare the class constants.
}

// Class Reference

func dereferenceClass() *dereferenceClass_ {
	return dereferenceClassReference_
}

var dereferenceClassReference_ = &dereferenceClass_{
	// Initialize the class constants.
}
