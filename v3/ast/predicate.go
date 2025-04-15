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

func PredicateClass() PredicateClassLike {
	return predicateClass()
}

// Constructor Methods

func (c *predicateClass_) Predicate(
	operator OperatorLike,
	expression ExpressionLike,
) PredicateLike {
	if uti.IsUndefined(operator) {
		panic("The \"operator\" attribute is required by this class.")
	}
	if uti.IsUndefined(expression) {
		panic("The \"expression\" attribute is required by this class.")
	}
	var instance = &predicate_{
		// Initialize the instance attributes.
		operator_:   operator,
		expression_: expression,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *predicate_) GetClass() PredicateClassLike {
	return predicateClass()
}

// Attribute Methods

func (v *predicate_) GetOperator() OperatorLike {
	return v.operator_
}

func (v *predicate_) GetExpression() ExpressionLike {
	return v.expression_
}

// PROTECTED INTERFACE

// Instance Structure

type predicate_ struct {
	// Declare the instance attributes.
	operator_   OperatorLike
	expression_ ExpressionLike
}

// Class Structure

type predicateClass_ struct {
	// Declare the class constants.
}

// Class Reference

func predicateClass() *predicateClass_ {
	return predicateClassReference_
}

var predicateClassReference_ = &predicateClass_{
	// Initialize the class constants.
}
