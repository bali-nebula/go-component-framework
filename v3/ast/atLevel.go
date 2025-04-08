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

func AtLevelClass() AtLevelClassLike {
	return atLevelClass()
}

// Constructor Methods

func (c *atLevelClass_) AtLevel(
	expression ExpressionLike,
) AtLevelLike {
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &atLevel_{
		// Initialize the instance attributes.
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *atLevel_) GetClass() AtLevelClassLike {
	return atLevelClass()
}

// Attribute Methods

func (v *atLevel_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type atLevel_ struct {
	// Declare the instance attributes.
	expression_ ExpressionLike
}

// Class Structure

type atLevelClass_ struct {
	// Declare the class constants.
}

// Class Reference

func atLevelClass() *atLevelClass_ {
	return atLevelClassReference_
}

var atLevelClassReference_ = &atLevelClass_{
	// Initialize the class constants.
}
