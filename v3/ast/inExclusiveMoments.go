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

func InExclusiveMomentsClass() InExclusiveMomentsClassLike {
	return inExclusiveMomentsClass()
}

// Constructor Methods

func (c *inExclusiveMomentsClass_) InExclusiveMoments(
	moment1 string,
	moment2 string,
) InExclusiveMomentsLike {
	if uti.IsUndefined(moment1) {
		panic("The \"moment1\" attribute is required by this class.")
	}
	if uti.IsUndefined(moment2) {
		panic("The \"moment2\" attribute is required by this class.")
	}
	var instance = &inExclusiveMoments_{
		// Initialize the instance attributes.
		moment1_: moment1,
		moment2_: moment2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inExclusiveMoments_) GetClass() InExclusiveMomentsClassLike {
	return inExclusiveMomentsClass()
}

// Attribute Methods

func (v *inExclusiveMoments_) GetMoment1() string {
	return v.moment1_
}

func (v *inExclusiveMoments_) GetMoment2() string {
	return v.moment2_
}

// PROTECTED INTERFACE

// Instance Structure

type inExclusiveMoments_ struct {
	// Declare the instance attributes.
	moment1_ string
	moment2_ string
}

// Class Structure

type inExclusiveMomentsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inExclusiveMomentsClass() *inExclusiveMomentsClass_ {
	return inExclusiveMomentsClassReference_
}

var inExclusiveMomentsClassReference_ = &inExclusiveMomentsClass_{
	// Initialize the class constants.
}
