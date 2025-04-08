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

func AnnotatedAssociationClass() AnnotatedAssociationClassLike {
	return annotatedAssociationClass()
}

// Constructor Methods

func (c *annotatedAssociationClass_) AnnotatedAssociation(
	association AssociationLike,
	optionalNote string,
) AnnotatedAssociationLike {
	if uti.IsUndefined(association) {
		panic("The \"association\" attribute is required by this class.")
	}
	var instance = &annotatedAssociation_{
		// Initialize the instance attributes.
		association_:  association,
		optionalNote_: optionalNote,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *annotatedAssociation_) GetClass() AnnotatedAssociationClassLike {
	return annotatedAssociationClass()
}

// Attribute Methods

func (v *annotatedAssociation_) GetAssociation() AssociationLike {
	return v.association_
}

func (v *annotatedAssociation_) GetOptionalNote() string {
	return v.optionalNote_
}

// PROTECTED INTERFACE

// Instance Structure

type annotatedAssociation_ struct {
	// Declare the instance attributes.
	association_  AssociationLike
	optionalNote_ string
}

// Class Structure

type annotatedAssociationClass_ struct {
	// Declare the class constants.
}

// Class Reference

func annotatedAssociationClass() *annotatedAssociationClass_ {
	return annotatedAssociationClassReference_
}

var annotatedAssociationClassReference_ = &annotatedAssociationClass_{
	// Initialize the class constants.
}
