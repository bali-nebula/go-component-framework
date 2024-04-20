/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package bali

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var indicesClass = &indicesClass_{
	// This class has no private constants to initialize.
}

// Function

func Indices() IndicesClassLike {
	return indicesClass
}

// CLASS METHODS

// Target

type indicesClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *indicesClass_) MakeWithIndexs(indexs col.ListLike[IndexLike]) IndicesLike {
	return &indices_{
		indexs_: indexs,
	}
}

// Functions

// INSTANCE METHODS

// Target

type indices_ struct {
	indexs_ col.ListLike[IndexLike]
}

// Attributes

func (v *indices_) GetIndexs() col.ListLike[IndexLike] {
	return v.indexs_
}

// Public

// Private
