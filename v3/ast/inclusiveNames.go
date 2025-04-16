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

func InclusiveNamesClass() InclusiveNamesClassLike {
	return inclusiveNamesClass()
}

// Constructor Methods

func (c *inclusiveNamesClass_) InclusiveNames(
	name1 string,
	name2 string,
) InclusiveNamesLike {
	if uti.IsUndefined(name1) {
		panic("The \"name1\" attribute is required by this class.")
	}
	if uti.IsUndefined(name2) {
		panic("The \"name2\" attribute is required by this class.")
	}
	var instance = &inclusiveNames_{
		// Initialize the instance attributes.
		name1_: name1,
		name2_: name2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inclusiveNames_) GetClass() InclusiveNamesClassLike {
	return inclusiveNamesClass()
}

// Attribute Methods

func (v *inclusiveNames_) GetName1() string {
	return v.name1_
}

func (v *inclusiveNames_) GetName2() string {
	return v.name2_
}

// PROTECTED INTERFACE

// Instance Structure

type inclusiveNames_ struct {
	// Declare the instance attributes.
	name1_ string
	name2_ string
}

// Class Structure

type inclusiveNamesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inclusiveNamesClass() *inclusiveNamesClass_ {
	return inclusiveNamesClassReference_
}

var inclusiveNamesClassReference_ = &inclusiveNamesClass_{
	// Initialize the class constants.
}
