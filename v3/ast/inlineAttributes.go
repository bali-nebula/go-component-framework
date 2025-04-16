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

func InlineAttributesClass() InlineAttributesClassLike {
	return inlineAttributesClass()
}

// Constructor Methods

func (c *inlineAttributesClass_) InlineAttributes(
	association AssociationLike,
	additionalAssociations col.Sequential[AdditionalAssociationLike],
) InlineAttributesLike {
	if uti.IsUndefined(association) {
		panic("The \"association\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalAssociations) {
		panic("The \"additionalAssociations\" attribute is required by this class.")
	}
	var instance = &inlineAttributes_{
		// Initialize the instance attributes.
		association_:            association,
		additionalAssociations_: additionalAssociations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inlineAttributes_) GetClass() InlineAttributesClassLike {
	return inlineAttributesClass()
}

// Attribute Methods

func (v *inlineAttributes_) GetAssociation() AssociationLike {
	return v.association_
}

func (v *inlineAttributes_) GetAdditionalAssociations() col.Sequential[AdditionalAssociationLike] {
	return v.additionalAssociations_
}

// PROTECTED INTERFACE

// Instance Structure

type inlineAttributes_ struct {
	// Declare the instance attributes.
	association_            AssociationLike
	additionalAssociations_ col.Sequential[AdditionalAssociationLike]
}

// Class Structure

type inlineAttributesClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineAttributesClass() *inlineAttributesClass_ {
	return inlineAttributesClassReference_
}

var inlineAttributesClassReference_ = &inlineAttributesClass_{
	// Initialize the class constants.
}
