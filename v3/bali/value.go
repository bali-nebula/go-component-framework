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

var valueClass = &valueClass_{
	// This class has no private constants to initialize.
}

// Function

func Value() ValueClassLike {
	return valueClass
}

// CLASS METHODS

// Target

type valueClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *valueClass_) MakeWithComponent(component ComponentLike) ValueLike {
	return &value_{
		component_: component,
	}
}

// Functions

// INSTANCE METHODS

// Target

type value_ struct {
	component_ ComponentLike
}

// Attributes

func (v *value_) GetComponent() ComponentLike {
	return v.component_
}

// Public

// Private
