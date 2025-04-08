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

func RepositoryClass() RepositoryClassLike {
	return repositoryClass()
}

// Constructor Methods

func (c *repositoryClass_) Repository(
	any_ any,
) RepositoryLike {
	if uti.IsUndefined(any_) {
		panic("The \"any\" attribute is required by this class.")
	}
	var instance = &repository_{
		// Initialize the instance attributes.
		any_: any_,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *repository_) GetClass() RepositoryClassLike {
	return repositoryClass()
}

// Attribute Methods

func (v *repository_) GetAny() any {
	return v.any_
}

// PROTECTED INTERFACE

// Instance Structure

type repository_ struct {
	// Declare the instance attributes.
	any_ any
}

// Class Structure

type repositoryClass_ struct {
	// Declare the class constants.
}

// Class Reference

func repositoryClass() *repositoryClass_ {
	return repositoryClassReference_
}

var repositoryClassReference_ = &repositoryClass_{
	// Initialize the class constants.
}
