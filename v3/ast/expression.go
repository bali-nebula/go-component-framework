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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func ExpressionClass() ExpressionClassLike {
	return expressionClass()
}

// Constructor Methods

func (c *expressionClass_) Expression(
	unary UnaryLike,
	operations col.Sequential[OperationLike],
) ExpressionLike {
	if uti.IsUndefined(unary) {
		panic("The \"unary\" attribute is required by this class.")
	}
	if uti.IsUndefined(operations) {
		panic("The \"operations\" attribute is required by this class.")
	}
	var instance = &expression_{
		// Initialize the instance attributes.
		unary_:      unary,
		operations_: operations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *expression_) GetClass() ExpressionClassLike {
	return expressionClass()
}

// Attribute Methods

func (v *expression_) GetUnary() UnaryLike {
	return v.unary_
}

func (v *expression_) GetOperations() col.Sequential[OperationLike] {
	return v.operations_
}

// PROTECTED INTERFACE

// Instance Structure

type expression_ struct {
	// Declare the instance attributes.
	unary_      UnaryLike
	operations_ col.Sequential[OperationLike]
}

// Class Structure

type expressionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func expressionClass() *expressionClass_ {
	return expressionClassReference_
}

var expressionClassReference_ = &expressionClass_{
	// Initialize the class constants.
}
