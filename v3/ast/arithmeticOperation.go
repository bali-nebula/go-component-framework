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

func ArithmeticOperationClass() ArithmeticOperationClassLike {
	return arithmeticOperationClass()
}

// Constructor Methods

func (c *arithmeticOperationClass_) ArithmeticOperation(
	arithmeticOperator ArithmeticOperatorLike,
	numerical NumericalLike,
) ArithmeticOperationLike {
	if uti.IsUndefined(arithmeticOperator) {
		panic("The \"arithmeticOperator\" attribute is required by this class.")
	}
	if uti.IsUndefined(numerical) {
		panic("The \"numerical\" attribute is required by this class.")
	}
	var instance = &arithmeticOperation_{
		// Initialize the instance attributes.
		arithmeticOperator_: arithmeticOperator,
		numerical_:          numerical,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *arithmeticOperation_) GetClass() ArithmeticOperationClassLike {
	return arithmeticOperationClass()
}

// Attribute Methods

func (v *arithmeticOperation_) GetArithmeticOperator() ArithmeticOperatorLike {
	return v.arithmeticOperator_
}

func (v *arithmeticOperation_) GetNumerical() NumericalLike {
	return v.numerical_
}

// PROTECTED INTERFACE

// Instance Structure

type arithmeticOperation_ struct {
	// Declare the instance attributes.
	arithmeticOperator_ ArithmeticOperatorLike
	numerical_          NumericalLike
}

// Class Structure

type arithmeticOperationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func arithmeticOperationClass() *arithmeticOperationClass_ {
	return arithmeticOperationClassReference_
}

var arithmeticOperationClassReference_ = &arithmeticOperationClass_{
	// Initialize the class constants.
}
