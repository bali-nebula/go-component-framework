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

func InExclusiveVersionsClass() InExclusiveVersionsClassLike {
	return inExclusiveVersionsClass()
}

// Constructor Methods

func (c *inExclusiveVersionsClass_) InExclusiveVersions(
	version1 string,
	version2 string,
) InExclusiveVersionsLike {
	if uti.IsUndefined(version1) {
		panic("The \"version1\" attribute is required by this class.")
	}
	if uti.IsUndefined(version2) {
		panic("The \"version2\" attribute is required by this class.")
	}
	var instance = &inExclusiveVersions_{
		// Initialize the instance attributes.
		version1_: version1,
		version2_: version2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inExclusiveVersions_) GetClass() InExclusiveVersionsClassLike {
	return inExclusiveVersionsClass()
}

// Attribute Methods

func (v *inExclusiveVersions_) GetVersion1() string {
	return v.version1_
}

func (v *inExclusiveVersions_) GetVersion2() string {
	return v.version2_
}

// PROTECTED INTERFACE

// Instance Structure

type inExclusiveVersions_ struct {
	// Declare the instance attributes.
	version1_ string
	version2_ string
}

// Class Structure

type inExclusiveVersionsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inExclusiveVersionsClass() *inExclusiveVersionsClass_ {
	return inExclusiveVersionsClassReference_
}

var inExclusiveVersionsClassReference_ = &inExclusiveVersionsClass_{
	// Initialize the class constants.
}
