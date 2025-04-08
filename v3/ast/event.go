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

func EventClass() EventClassLike {
	return eventClass()
}

// Constructor Methods

func (c *eventClass_) Event(
	expression ExpressionLike,
) EventLike {
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &event_{
		// Initialize the instance attributes.
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *event_) GetClass() EventClassLike {
	return eventClass()
}

// Attribute Methods

func (v *event_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type event_ struct {
	// Declare the instance attributes.
	expression_ ExpressionLike
}

// Class Structure

type eventClass_ struct {
	// Declare the class constants.
}

// Class Reference

func eventClass() *eventClass_ {
	return eventClassReference_
}

var eventClassReference_ = &eventClass_{
	// Initialize the class constants.
}
