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

var associationsClass = &associationsClass_{
	// This class has no private constants to initialize.
}

// Function

func Associations() AssociationsClassLike {
	return associationsClass
}

// CLASS METHODS

// Target

type associationsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *associationsClass_) MakeWithAssociations(associations col.ListLike[AssociationLike]) AssociationsLike {
	return &associations_{
		associations_: associations,
	}
}

func (c *associationsClass_) MakeWithAttributes(
	association AssociationLike,
	note string,
) AssociationsLike {
	return &associations_{
		association_: association,
		note_:        note,
	}
}

// Functions

// INSTANCE METHODS

// Target

type associations_ struct {
	associations_ col.ListLike[AssociationLike]
	note_         string
	association_  AssociationLike
}

// Attributes

func (v *associations_) GetAssociations() col.ListLike[AssociationLike] {
	return v.associations_
}

func (v *associations_) GetNote() string {
	return v.note_
}

// Public

// Private
