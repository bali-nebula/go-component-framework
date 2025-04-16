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

func InlineParametersClass() InlineParametersClassLike {
	return inlineParametersClass()
}

// Constructor Methods

func (c *inlineParametersClass_) InlineParameters(
	association AssociationLike,
	additionalAssociations col.Sequential[AdditionalAssociationLike],
) InlineParametersLike {
	if uti.IsUndefined(association) {
		panic("The \"association\" attribute is required by this class.")
	}
	if uti.IsUndefined(additionalAssociations) {
		panic("The \"additionalAssociations\" attribute is required by this class.")
	}
	var instance = &inlineParameters_{
		// Initialize the instance attributes.
		association_:            association,
		additionalAssociations_: additionalAssociations,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *inlineParameters_) GetClass() InlineParametersClassLike {
	return inlineParametersClass()
}

// Attribute Methods

func (v *inlineParameters_) GetAssociation() AssociationLike {
	return v.association_
}

func (v *inlineParameters_) GetAdditionalAssociations() col.Sequential[AdditionalAssociationLike] {
	return v.additionalAssociations_
}

// PROTECTED INTERFACE

// Instance Structure

type inlineParameters_ struct {
	// Declare the instance attributes.
	association_            AssociationLike
	additionalAssociations_ col.Sequential[AdditionalAssociationLike]
}

// Class Structure

type inlineParametersClass_ struct {
	// Declare the class constants.
}

// Class Reference

func inlineParametersClass() *inlineParametersClass_ {
	return inlineParametersClassReference_
}

var inlineParametersClassReference_ = &inlineParametersClass_{
	// Initialize the class constants.
}
