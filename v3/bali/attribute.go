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

var attributeClass = &attributeClass_{
	// This class has no private constants to initialize.
}

// Function

func Attribute() AttributeClassLike {
	return attributeClass
}

// CLASS METHODS

// Target

type attributeClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *attributeClass_) MakeWithAttributes(
	variable VariableLike,
	indices IndicesLike,
) AttributeLike {
	return &attribute_{
		variable_: variable,
		indices_:  indices,
	}
}

// Functions

// INSTANCE METHODS

// Target

type attribute_ struct {
	variable_ VariableLike
	indices_  IndicesLike
}

// Attributes

func (v *attribute_) GetVariable() VariableLike {
	return v.variable_
}

func (v *attribute_) GetIndices() IndicesLike {
	return v.indices_
}

// Public

// Private
