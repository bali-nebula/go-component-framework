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

func IntrinsicClass() IntrinsicClassLike {
	return intrinsicClass()
}

// Constructor Methods

func (c *intrinsicClass_) Intrinsic(
	function FunctionLike,
	optionalArguments ArgumentsLike,
) IntrinsicLike {
	if uti.IsUndefined(function) {
		panic("The \"function\" attribute is required by this class.")
	}
	var instance = &intrinsic_{
		// Initialize the instance attributes.
		function_:          function,
		optionalArguments_: optionalArguments,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *intrinsic_) GetClass() IntrinsicClassLike {
	return intrinsicClass()
}

// Attribute Methods

func (v *intrinsic_) GetFunction() FunctionLike {
	return v.function_
}

func (v *intrinsic_) GetOptionalArguments() ArgumentsLike {
	return v.optionalArguments_
}

// PROTECTED INTERFACE

// Instance Structure

type intrinsic_ struct {
	// Declare the instance attributes.
	function_          FunctionLike
	optionalArguments_ ArgumentsLike
}

// Class Structure

type intrinsicClass_ struct {
	// Declare the class constants.
}

// Class Reference

func intrinsicClass() *intrinsicClass_ {
	return intrinsicClassReference_
}

var intrinsicClassReference_ = &intrinsicClass_{
	// Initialize the class constants.
}
