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

func ExclusiveMomentsClass() ExclusiveMomentsClassLike {
	return exclusiveMomentsClass()
}

// Constructor Methods

func (c *exclusiveMomentsClass_) ExclusiveMoments(
	moment1 string,
	moment2 string,
) ExclusiveMomentsLike {
	if uti.IsUndefined(moment1) {
		panic("The \"moment1\" attribute is required by this class.")
	}
	if uti.IsUndefined(moment2) {
		panic("The \"moment2\" attribute is required by this class.")
	}
	var instance = &exclusiveMoments_{
		// Initialize the instance attributes.
		moment1_: moment1,
		moment2_: moment2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exclusiveMoments_) GetClass() ExclusiveMomentsClassLike {
	return exclusiveMomentsClass()
}

// Attribute Methods

func (v *exclusiveMoments_) GetMoment1() string {
	return v.moment1_
}

func (v *exclusiveMoments_) GetMoment2() string {
	return v.moment2_
}

// PROTECTED INTERFACE

// Instance Structure

type exclusiveMoments_ struct {
	// Declare the instance attributes.
	moment1_ string
	moment2_ string
}

// Class Structure

type exclusiveMomentsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exclusiveMomentsClass() *exclusiveMomentsClass_ {
	return exclusiveMomentsClassReference_
}

var exclusiveMomentsClassReference_ = &exclusiveMomentsClass_{
	// Initialize the class constants.
}
