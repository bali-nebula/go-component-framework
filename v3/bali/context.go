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

var contextClass = &contextClass_{
	// This class has no private constants to initialize.
}

// Function

func Context() ContextClassLike {
	return contextClass
}

// CLASS METHODS

// Target

type contextClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *contextClass_) MakeWithParameters(parameters ParametersLike) ContextLike {
	return &context_{
		parameters_: parameters,
	}
}

// Functions

// INSTANCE METHODS

// Target

type context_ struct {
	parameters_ ParametersLike
}

// Attributes

func (v *context_) GetParameters() ParametersLike {
	return v.parameters_
}

// Public

// Private
