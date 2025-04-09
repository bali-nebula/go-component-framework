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

func LogicalClass() LogicalClassLike {
	return logicalClass()
}

// Constructor Methods

func (c *logicalClass_) Logical(
	expression1 ExpressionLike,
	logicalOperator LogicalOperatorLike,
	expression2 ExpressionLike,
) LogicalLike {
	if uti.IsUndefined(expression1) {
		panic("The \"expression1\" attribute is required by this class.")
	}
	if uti.IsUndefined(logicalOperator) {
		panic("The \"logicalOperator\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression2) {
		panic("The \"expression2\" attribute is required by this class.")
	}
	var instance = &logical_{
		// Initialize the instance attributes.
		expression1_:     expression1,
		logicalOperator_: logicalOperator,
		expression2_:     expression2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *logical_) GetClass() LogicalClassLike {
	return logicalClass()
}

// Attribute Methods

func (v *logical_) GetExpression1() ExpressionLike {
	return v.expression1_
}

func (v *logical_) GetLogicalOperator() LogicalOperatorLike {
	return v.logicalOperator_
}

func (v *logical_) GetExpression2() ExpressionLike {
	return v.expression2_
}

// PROTECTED INTERFACE

// Instance Structure

type logical_ struct {
	// Declare the instance attributes.
	expression1_     ExpressionLike
	logicalOperator_ LogicalOperatorLike
	expression2_     ExpressionLike
}

// Class Structure

type logicalClass_ struct {
	// Declare the class constants.
}

// Class Reference

func logicalClass() *logicalClass_ {
	return logicalClassReference_
}

var logicalClassReference_ = &logicalClass_{
	// Initialize the class constants.
}
