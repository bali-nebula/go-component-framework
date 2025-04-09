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

func InvocationClass() InvocationClassLike {
	return invocationClass()
}

// Constructor Methods

func (c *invocationClass_) Invocation(
	target TargetLike,
	invocationOperator InvocationOperatorLike,
	method MethodLike,
	optionalArguments ArgumentsLike,
) InvocationLike {
	if uti.IsUndefined(target) {
		panic("The \"target\" attribute is required by this class.")
	}
	if uti.IsUndefined(invocationOperator) {
		panic("The \"invocationOperator\" attribute is required by this class.")
	}
	if uti.IsUndefined(method) {
		panic("The \"method\" attribute is required by this class.")
	}
	var instance = &invocation_{
		// Initialize the instance attributes.
		target_:             target,
		invocationOperator_: invocationOperator,
		method_:             method,
		optionalArguments_:  optionalArguments,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *invocation_) GetClass() InvocationClassLike {
	return invocationClass()
}

// Attribute Methods

func (v *invocation_) GetTarget() TargetLike {
	return v.target_
}

func (v *invocation_) GetInvocationOperator() InvocationOperatorLike {
	return v.invocationOperator_
}

func (v *invocation_) GetMethod() MethodLike {
	return v.method_
}

func (v *invocation_) GetOptionalArguments() ArgumentsLike {
	return v.optionalArguments_
}

// PROTECTED INTERFACE

// Instance Structure

type invocation_ struct {
	// Declare the instance attributes.
	target_             TargetLike
	invocationOperator_ InvocationOperatorLike
	method_             MethodLike
	optionalArguments_  ArgumentsLike
}

// Class Structure

type invocationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func invocationClass() *invocationClass_ {
	return invocationClassReference_
}

var invocationClassReference_ = &invocationClass_{
	// Initialize the class constants.
}
