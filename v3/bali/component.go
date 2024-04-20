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

var componentClass = &componentClass_{
	// This class has no private constants to initialize.
}

// Function

func Component() ComponentClassLike {
	return componentClass
}

// CLASS METHODS

// Target

type componentClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *componentClass_) MakeWithAttributes(
	entity EntityLike,
	context ContextLike,
) ComponentLike {
	return &component_{
		entity_:  entity,
		context_: context,
	}
}

// Functions

// INSTANCE METHODS

// Target

type component_ struct {
	entity_  EntityLike
	context_ ContextLike
}

// Attributes

func (v *component_) GetEntity() EntityLike {
	return v.entity_
}

func (v *component_) GetContext() ContextLike {
	return v.context_
}

// Public

// Private
