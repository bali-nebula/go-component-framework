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
	threading ThreadingLike,
	identifier string,
	optionalArguments ArgumentsLike,
) InvocationLike {
	if uti.IsUndefined(threading) {
		panic("The \"threading\" attribute is required by this class.")
	}
	if uti.IsUndefined(identifier) {
		panic("The \"identifier\" attribute is required by this class.")
	}
	var instance = &invocation_{
		// Initialize the instance attributes.
		threading_:         threading,
		identifier_:        identifier,
		optionalArguments_: optionalArguments,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *invocation_) GetClass() InvocationClassLike {
	return invocationClass()
}

// Attribute Methods

func (v *invocation_) GetThreading() ThreadingLike {
	return v.threading_
}

func (v *invocation_) GetIdentifier() string {
	return v.identifier_
}

func (v *invocation_) GetOptionalArguments() ArgumentsLike {
	return v.optionalArguments_
}

// PROTECTED INTERFACE

// Instance Structure

type invocation_ struct {
	// Declare the instance attributes.
	threading_         ThreadingLike
	identifier_        string
	optionalArguments_ ArgumentsLike
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
