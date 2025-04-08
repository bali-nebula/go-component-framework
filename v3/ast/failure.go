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

func FailureClass() FailureClassLike {
	return failureClass()
}

// Constructor Methods

func (c *failureClass_) Failure(
	symbol string,
) FailureLike {
	if uti.IsUndefined(symbol) {
		panic("The \"symbol\" attribute is required by this class.")
	}
	var instance = &failure_{
		// Initialize the instance attributes.
		symbol_: symbol,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *failure_) GetClass() FailureClassLike {
	return failureClass()
}

// Attribute Methods

func (v *failure_) GetSymbol() string {
	return v.symbol_
}

// PROTECTED INTERFACE

// Instance Structure

type failure_ struct {
	// Declare the instance attributes.
	symbol_ string
}

// Class Structure

type failureClass_ struct {
	// Declare the class constants.
}

// Class Reference

func failureClass() *failureClass_ {
	return failureClassReference_
}

var failureClassReference_ = &failureClass_{
	// Initialize the class constants.
}
