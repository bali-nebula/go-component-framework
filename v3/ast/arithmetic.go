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

func ArithmeticClass() ArithmeticClassLike {
	return arithmeticClass()
}

// Constructor Methods

func (c *arithmeticClass_) Arithmetic(
	numerical NumericalLike,
	arithmeticOperations col.Sequential[ArithmeticOperationLike],
) ArithmeticLike {
	if uti.IsUndefined(numerical) {
		panic("The \"numerical\" attribute is required by this class.")
	}
	if uti.IsUndefined(arithmeticOperations) {
		panic("The \"arithmeticOperations\" attribute is required by this class.")
	}
	var instance = &arithmetic_{
		// Initialize the instance attributes.
		numerical_:            numerical,
		arithmeticOperations_: arithmeticOperations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *arithmetic_) GetClass() ArithmeticClassLike {
	return arithmeticClass()
}

// Attribute Methods

func (v *arithmetic_) GetNumerical() NumericalLike {
	return v.numerical_
}

func (v *arithmetic_) GetArithmeticOperations() col.Sequential[ArithmeticOperationLike] {
	return v.arithmeticOperations_
}

// PROTECTED INTERFACE

// Instance Structure

type arithmetic_ struct {
	// Declare the instance attributes.
	numerical_            NumericalLike
	arithmeticOperations_ col.Sequential[ArithmeticOperationLike]
}

// Class Structure

type arithmeticClass_ struct {
	// Declare the class constants.
}

// Class Reference

func arithmeticClass() *arithmeticClass_ {
	return arithmeticClassReference_
}

var arithmeticClassReference_ = &arithmeticClass_{
	// Initialize the class constants.
}
