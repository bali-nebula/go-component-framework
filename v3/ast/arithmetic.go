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

func ArithmeticClass() ArithmeticClassLike {
	return arithmeticClass()
}

// Constructor Methods

func (c *arithmeticClass_) Arithmetic(
	expression1 ExpressionLike,
	arithmeticOperator ArithmeticOperatorLike,
	expression2 ExpressionLike,
) ArithmeticLike {
	if uti.IsUndefined(expression1) {
		panic("The \"expression1\" attribute is required by this class.")
	}
	if uti.IsUndefined(arithmeticOperator) {
		panic("The \"arithmeticOperator\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression2) {
		panic("The \"expression2\" attribute is required by this class.")
	}
	var instance = &arithmetic_{
		// Initialize the instance attributes.
		expression1_:        expression1,
		arithmeticOperator_: arithmeticOperator,
		expression2_:        expression2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *arithmetic_) GetClass() ArithmeticClassLike {
	return arithmeticClass()
}

// Attribute Methods

func (v *arithmetic_) GetExpression1() ExpressionLike {
	return v.expression1_
}

func (v *arithmetic_) GetArithmeticOperator() ArithmeticOperatorLike {
	return v.arithmeticOperator_
}

func (v *arithmetic_) GetExpression2() ExpressionLike {
	return v.expression2_
}

// PROTECTED INTERFACE

// Instance Structure

type arithmetic_ struct {
	// Declare the instance attributes.
	expression1_        ExpressionLike
	arithmeticOperator_ ArithmeticOperatorLike
	expression2_        ExpressionLike
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
