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

func AdditionalIndexClass() AdditionalIndexClassLike {
	return additionalIndexClass()
}

// Constructor Methods

func (c *additionalIndexClass_) AdditionalIndex(
	index IndexLike,
) AdditionalIndexLike {
	if uti.IsUndefined(index) {
		panic("The \"index\" attribute is required by this class.")
	}
	var instance = &additionalIndex_{
		// Initialize the instance attributes.
		index_: index,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalIndex_) GetClass() AdditionalIndexClassLike {
	return additionalIndexClass()
}

// Attribute Methods

func (v *additionalIndex_) GetIndex() IndexLike {
	return v.index_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalIndex_ struct {
	// Declare the instance attributes.
	index_ IndexLike
}

// Class Structure

type additionalIndexClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalIndexClass() *additionalIndexClass_ {
	return additionalIndexClassReference_
}

var additionalIndexClassReference_ = &additionalIndexClass_{
	// Initialize the class constants.
}
