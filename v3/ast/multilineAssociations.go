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

func MultilineAssociationsClass() MultilineAssociationsClassLike {
	return multilineAssociationsClass()
}

// Constructor Methods

func (c *multilineAssociationsClass_) MultilineAssociations(
	annotatedAssociations col.Sequential[AnnotatedAssociationLike],
) MultilineAssociationsLike {
	if uti.IsUndefined(annotatedAssociations) {
		panic("The \"annotatedAssociations\" attribute is required by this class.")
	}
	var instance = &multilineAssociations_{
		// Initialize the instance attributes.
		annotatedAssociations_: annotatedAssociations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multilineAssociations_) GetClass() MultilineAssociationsClassLike {
	return multilineAssociationsClass()
}

// Attribute Methods

func (v *multilineAssociations_) GetAnnotatedAssociations() col.Sequential[AnnotatedAssociationLike] {
	return v.annotatedAssociations_
}

// PROTECTED INTERFACE

// Instance Structure

type multilineAssociations_ struct {
	// Declare the instance attributes.
	annotatedAssociations_ col.Sequential[AnnotatedAssociationLike]
}

// Class Structure

type multilineAssociationsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multilineAssociationsClass() *multilineAssociationsClass_ {
	return multilineAssociationsClassReference_
}

var multilineAssociationsClassReference_ = &multilineAssociationsClass_{
	// Initialize the class constants.
}
