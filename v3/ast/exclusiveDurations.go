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

func ExclusiveDurationsClass() ExclusiveDurationsClassLike {
	return exclusiveDurationsClass()
}

// Constructor Methods

func (c *exclusiveDurationsClass_) ExclusiveDurations(
	duration1 string,
	duration2 string,
) ExclusiveDurationsLike {
	if uti.IsUndefined(duration1) {
		panic("The \"duration1\" attribute is required by this class.")
	}
	if uti.IsUndefined(duration2) {
		panic("The \"duration2\" attribute is required by this class.")
	}
	var instance = &exclusiveDurations_{
		// Initialize the instance attributes.
		duration1_: duration1,
		duration2_: duration2,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *exclusiveDurations_) GetClass() ExclusiveDurationsClassLike {
	return exclusiveDurationsClass()
}

// Attribute Methods

func (v *exclusiveDurations_) GetDuration1() string {
	return v.duration1_
}

func (v *exclusiveDurations_) GetDuration2() string {
	return v.duration2_
}

// PROTECTED INTERFACE

// Instance Structure

type exclusiveDurations_ struct {
	// Declare the instance attributes.
	duration1_ string
	duration2_ string
}

// Class Structure

type exclusiveDurationsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func exclusiveDurationsClass() *exclusiveDurationsClass_ {
	return exclusiveDurationsClassReference_
}

var exclusiveDurationsClassReference_ = &exclusiveDurationsClass_{
	// Initialize the class constants.
}
