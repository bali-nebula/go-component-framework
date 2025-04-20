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

func MethodClass() MethodClassLike {
	return methodClass()
}

// Constructor Methods

func (c *methodClass_) Method(
	identifier1 string,
	blocking BlockingLike,
	identifier2 string,
	optionalArguments ArgumentsLike,
) MethodLike {
	if uti.IsUndefined(identifier1) {
		panic("The \"identifier1\" attribute is required by this class.")
	}
	if uti.IsUndefined(blocking) {
		panic("The \"blocking\" attribute is required by this class.")
	}
	if uti.IsUndefined(identifier2) {
		panic("The \"identifier2\" attribute is required by this class.")
	}
	var instance = &method_{
		// Initialize the instance attributes.
		identifier1_:       identifier1,
		blocking_:          blocking,
		identifier2_:       identifier2,
		optionalArguments_: optionalArguments,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *method_) GetClass() MethodClassLike {
	return methodClass()
}

// Attribute Methods

func (v *method_) GetIdentifier1() string {
	return v.identifier1_
}

func (v *method_) GetBlocking() BlockingLike {
	return v.blocking_
}

func (v *method_) GetIdentifier2() string {
	return v.identifier2_
}

func (v *method_) GetOptionalArguments() ArgumentsLike {
	return v.optionalArguments_
}

// PROTECTED INTERFACE

// Instance Structure

type method_ struct {
	// Declare the instance attributes.
	identifier1_       string
	blocking_          BlockingLike
	identifier2_       string
	optionalArguments_ ArgumentsLike
}

// Class Structure

type methodClass_ struct {
	// Declare the class constants.
}

// Class Reference

func methodClass() *methodClass_ {
	return methodClassReference_
}

var methodClassReference_ = &methodClass_{
	// Initialize the class constants.
}
