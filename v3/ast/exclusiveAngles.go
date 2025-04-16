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

func ExclusiveAnglesClass() ExclusiveAnglesClassLike {
	return exclusiveAnglesClass()
}

// Constructor Methods

func (c *exclusiveAnglesClass_) ExclusiveAngles(
	angle1 string,
	angle2 string,
) ExclusiveAnglesLike {
	if uti.IsUndefined(angle1) {
		panic("The \"angle1\" attribute is required by this class.")
	}
	if uti.IsUndefined(angle2) {
		panic("The \"angle2\" attribute is required by this class.")
	}
	var instance = &exclusiveAngles_{
		// Initialize the instance attributes.
		angle1_: angle1,
		angle2_: angle2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exclusiveAngles_) GetClass() ExclusiveAnglesClassLike {
	return exclusiveAnglesClass()
}

// Attribute Methods

func (v *exclusiveAngles_) GetAngle1() string {
	return v.angle1_
}

func (v *exclusiveAngles_) GetAngle2() string {
	return v.angle2_
}

// PROTECTED INTERFACE

// Instance Structure

type exclusiveAngles_ struct {
	// Declare the instance attributes.
	angle1_ string
	angle2_ string
}

// Class Structure

type exclusiveAnglesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exclusiveAnglesClass() *exclusiveAnglesClass_ {
	return exclusiveAnglesClassReference_
}

var exclusiveAnglesClassReference_ = &exclusiveAnglesClass_{
	// Initialize the class constants.
}
