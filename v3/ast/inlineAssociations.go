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

func InlineAssociationsClass() InlineAssociationsClassLike {
	return inlineAssociationsClass()
}

// Constructor Methods

func (c *inlineAssociationsClass_) InlineAssociations(
	association AssociationLike,
	additionalAssociations col.Sequential[AdditionalAssociationLike],
) InlineAssociationsLike {
	if uti.IsUndefined(association) {
		panic("The \"association\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalAssociations) {
		panic("The \"additionalAssociations\" attribute is required by this class.")
	}
	var instance = &inlineAssociations_{
		// Initialize the instance attributes.
		association_:            association,
		additionalAssociations_: additionalAssociations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inlineAssociations_) GetClass() InlineAssociationsClassLike {
	return inlineAssociationsClass()
}

// Attribute Methods

func (v *inlineAssociations_) GetAssociation() AssociationLike {
	return v.association_
}

func (v *inlineAssociations_) GetAdditionalAssociations() col.Sequential[AdditionalAssociationLike] {
	return v.additionalAssociations_
}

// PROTECTED INTERFACE

// Instance Structure

type inlineAssociations_ struct {
	// Declare the instance attributes.
	association_            AssociationLike
	additionalAssociations_ col.Sequential[AdditionalAssociationLike]
}

// Class Structure

type inlineAssociationsClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineAssociationsClass() *inlineAssociationsClass_ {
	return inlineAssociationsClassReference_
}

var inlineAssociationsClassReference_ = &inlineAssociationsClass_{
	// Initialize the class constants.
}
