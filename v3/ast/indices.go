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
	col "github.com/craterdog/go-collection-framework/v5/collection"
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func IndicesClass() IndicesClassLike {
	return indicesClass()
}

// Constructor Methods

func (c *indicesClass_) Indices(
	index IndexLike,
	additionalIndexes col.Sequential[AdditionalIndexLike],
) IndicesLike {
	if uti.IsUndefined(index) {
		panic("The \"index\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalIndexes) {
		panic("The \"additionalIndexes\" attribute is required by this class.")
	}
	var instance = &indices_{
		// Initialize the instance attributes.
		index_:             index,
		additionalIndexes_: additionalIndexes,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *indices_) GetClass() IndicesClassLike {
	return indicesClass()
}

// Attribute Methods

func (v *indices_) GetIndex() IndexLike {
	return v.index_
}

func (v *indices_) GetAdditionalIndexes() col.Sequential[AdditionalIndexLike] {
	return v.additionalIndexes_
}

// PROTECTED INTERFACE

// Instance Structure

type indices_ struct {
	// Declare the instance attributes.
	index_             IndexLike
	additionalIndexes_ col.Sequential[AdditionalIndexLike]
}

// Class Structure

type indicesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func indicesClass() *indicesClass_ {
	return indicesClassReference_
}

var indicesClassReference_ = &indicesClass_{
	// Initialize the class constants.
}
