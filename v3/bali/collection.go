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

import ()

// CLASS ACCESS

// Reference

var collectionClass = &collectionClass_{
	// This class has no private constants to initialize.
}

// Function

func Collection() CollectionClassLike {
	return collectionClass
}

// CLASS METHODS

// Target

type collectionClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *collectionClass_) MakeWithAttributes(
	associations AssociationsLike,
	values ValuesLike,
) CollectionLike {
	return &collection_{
		associations_: associations,
		values_:       values,
	}
}

// Functions

// INSTANCE METHODS

// Target

type collection_ struct {
	associations_ AssociationsLike
	values_       ValuesLike
}

// Attributes

func (v *collection_) GetAssociations() AssociationsLike {
	return v.associations_
}

func (v *collection_) GetValues() ValuesLike {
	return v.values_
}

// Public

// Private
