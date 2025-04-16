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

func MultilineAttributesClass() MultilineAttributesClassLike {
	return multilineAttributesClass()
}

// Constructor Methods

func (c *multilineAttributesClass_) MultilineAttributes(
	newline string,
	annotatedAssociations col.Sequential[AnnotatedAssociationLike],
) MultilineAttributesLike {
	if uti.IsUndefined(newline) {
		panic("The \"newline\" attribute is required by this class.")
	}
	if uti.IsUndefined(annotatedAssociations) {
		panic("The \"annotatedAssociations\" attribute is required by this class.")
	}
	var instance = &multilineAttributes_{
		// Initialize the instance attributes.
		newline_:               newline,
		annotatedAssociations_: annotatedAssociations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *multilineAttributes_) GetClass() MultilineAttributesClassLike {
	return multilineAttributesClass()
}

// Attribute Methods

func (v *multilineAttributes_) GetNewline() string {
	return v.newline_
}

func (v *multilineAttributes_) GetAnnotatedAssociations() col.Sequential[AnnotatedAssociationLike] {
	return v.annotatedAssociations_
}

// PROTECTED INTERFACE

// Instance Structure

type multilineAttributes_ struct {
	// Declare the instance attributes.
	newline_               string
	annotatedAssociations_ col.Sequential[AnnotatedAssociationLike]
}

// Class Structure

type multilineAttributesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func multilineAttributesClass() *multilineAttributesClass_ {
	return multilineAttributesClassReference_
}

var multilineAttributesClassReference_ = &multilineAttributesClass_{
	// Initialize the class constants.
}
