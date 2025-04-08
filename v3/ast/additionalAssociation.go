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

func AdditionalAssociationClass() AdditionalAssociationClassLike {
	return additionalAssociationClass()
}

// Constructor Methods

func (c *additionalAssociationClass_) AdditionalAssociation(
	association AssociationLike,
) AdditionalAssociationLike {
	if uti.IsUndefined(association) {
		panic("The \"association\" attribute is required by this class.")
	}
	var instance = &additionalAssociation_{
		// Initialize the instance attributes.
		association_: association,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *additionalAssociation_) GetClass() AdditionalAssociationClassLike {
	return additionalAssociationClass()
}

// Attribute Methods

func (v *additionalAssociation_) GetAssociation() AssociationLike {
	return v.association_
}

// PROTECTED INTERFACE

// Instance Structure

type additionalAssociation_ struct {
	// Declare the instance attributes.
	association_ AssociationLike
}

// Class Structure

type additionalAssociationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func additionalAssociationClass() *additionalAssociationClass_ {
	return additionalAssociationClassReference_
}

var additionalAssociationClassReference_ = &additionalAssociationClass_{
	// Initialize the class constants.
}
