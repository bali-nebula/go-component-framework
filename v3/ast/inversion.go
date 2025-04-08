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

func InversionClass() InversionClassLike {
	return inversionClass()
}

// Constructor Methods

func (c *inversionClass_) Inversion(
	inversionOperator string,
	expression ExpressionLike,
) InversionLike {
	if uti.IsUndefined(inversionOperator) {
		panic("The \"inversionOperator\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &inversion_{
		// Initialize the instance attributes.
		inversionOperator_: inversionOperator,
		expression_:        expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inversion_) GetClass() InversionClassLike {
	return inversionClass()
}

// Attribute Methods

func (v *inversion_) GetInversionOperator() string {
	return v.inversionOperator_
}

func (v *inversion_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type inversion_ struct {
	// Declare the instance attributes.
	inversionOperator_ string
	expression_        ExpressionLike
}

// Class Structure

type inversionClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inversionClass() *inversionClass_ {
	return inversionClassReference_
}

var inversionClassReference_ = &inversionClass_{
	// Initialize the class constants.
}
