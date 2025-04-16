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

func ExInclusiveMomentsClass() ExInclusiveMomentsClassLike {
	return exInclusiveMomentsClass()
}

// Constructor Methods

func (c *exInclusiveMomentsClass_) ExInclusiveMoments(
	moment1 string,
	moment2 string,
) ExInclusiveMomentsLike {
	if uti.IsUndefined(moment1) {
		panic("The \"moment1\" attribute is required by this class.")
	}
	if uti.IsUndefined(moment2) {
		panic("The \"moment2\" attribute is required by this class.")
	}
	var instance = &exInclusiveMoments_{
		// Initialize the instance attributes.
		moment1_: moment1,
		moment2_: moment2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exInclusiveMoments_) GetClass() ExInclusiveMomentsClassLike {
	return exInclusiveMomentsClass()
}

// Attribute Methods

func (v *exInclusiveMoments_) GetMoment1() string {
	return v.moment1_
}

func (v *exInclusiveMoments_) GetMoment2() string {
	return v.moment2_
}

// PROTECTED INTERFACE

// Instance Structure

type exInclusiveMoments_ struct {
	// Declare the instance attributes.
	moment1_ string
	moment2_ string
}

// Class Structure

type exInclusiveMomentsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exInclusiveMomentsClass() *exInclusiveMomentsClass_ {
	return exInclusiveMomentsClassReference_
}

var exInclusiveMomentsClassReference_ = &exInclusiveMomentsClass_{
	// Initialize the class constants.
}
