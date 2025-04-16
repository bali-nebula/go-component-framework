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

func ExclusiveVersionsClass() ExclusiveVersionsClassLike {
	return exclusiveVersionsClass()
}

// Constructor Methods

func (c *exclusiveVersionsClass_) ExclusiveVersions(
	version1 string,
	version2 string,
) ExclusiveVersionsLike {
	if uti.IsUndefined(version1) {
		panic("The \"version1\" attribute is required by this class.")
	}
	if uti.IsUndefined(version2) {
		panic("The \"version2\" attribute is required by this class.")
	}
	var instance = &exclusiveVersions_{
		// Initialize the instance attributes.
		version1_: version1,
		version2_: version2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exclusiveVersions_) GetClass() ExclusiveVersionsClassLike {
	return exclusiveVersionsClass()
}

// Attribute Methods

func (v *exclusiveVersions_) GetVersion1() string {
	return v.version1_
}

func (v *exclusiveVersions_) GetVersion2() string {
	return v.version2_
}

// PROTECTED INTERFACE

// Instance Structure

type exclusiveVersions_ struct {
	// Declare the instance attributes.
	version1_ string
	version2_ string
}

// Class Structure

type exclusiveVersionsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exclusiveVersionsClass() *exclusiveVersionsClass_ {
	return exclusiveVersionsClassReference_
}

var exclusiveVersionsClassReference_ = &exclusiveVersionsClass_{
	// Initialize the class constants.
}
