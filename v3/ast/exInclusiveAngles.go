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

func ExInclusiveAnglesClass() ExInclusiveAnglesClassLike {
	return exInclusiveAnglesClass()
}

// Constructor Methods

func (c *exInclusiveAnglesClass_) ExInclusiveAngles(
	angle1 string,
	angle2 string,
) ExInclusiveAnglesLike {
	if uti.IsUndefined(angle1) {
		panic("The \"angle1\" attribute is required by this class.")
	}
	if uti.IsUndefined(angle2) {
		panic("The \"angle2\" attribute is required by this class.")
	}
	var instance = &exInclusiveAngles_{
		// Initialize the instance attributes.
		angle1_: angle1,
		angle2_: angle2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exInclusiveAngles_) GetClass() ExInclusiveAnglesClassLike {
	return exInclusiveAnglesClass()
}

// Attribute Methods

func (v *exInclusiveAngles_) GetAngle1() string {
	return v.angle1_
}

func (v *exInclusiveAngles_) GetAngle2() string {
	return v.angle2_
}

// PROTECTED INTERFACE

// Instance Structure

type exInclusiveAngles_ struct {
	// Declare the instance attributes.
	angle1_ string
	angle2_ string
}

// Class Structure

type exInclusiveAnglesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exInclusiveAnglesClass() *exInclusiveAnglesClass_ {
	return exInclusiveAnglesClassReference_
}

var exInclusiveAnglesClassReference_ = &exInclusiveAnglesClass_{
	// Initialize the class constants.
}
