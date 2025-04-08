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

func EntityClass() EntityClassLike {
	return entityClass()
}

// Constructor Methods

func (c *entityClass_) Entity(
	any_ any,
) EntityLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &entity_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *entity_) GetClass() EntityClassLike {
	return entityClass()
}

// Attribute Methods

func (v *entity_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type entity_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type entityClass_ struct {
	// Declare the class constants.
}

// Class Reference

func entityClass() *entityClass_ {
	return entityClassReference_
}

var entityClassReference_ = &entityClass_{
	// Initialize the class constants.
}
