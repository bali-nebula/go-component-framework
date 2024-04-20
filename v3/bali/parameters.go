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

var parametersClass = &parametersClass_{
	// This class has no private constants to initialize.
}

// Function

func Parameters() ParametersClassLike {
	return parametersClass
}

// CLASS METHODS

// Target

type parametersClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *parametersClass_) MakeWithParameters(parameters col.ListLike[ParameterLike]) ParametersLike {
	return &parameters_{
		parameters_: parameters,
	}
}

func (c *parametersClass_) MakeWithAttributes(
	parameter ParameterLike,
	note string,
) ParametersLike {
	return &parameters_{
		parameter_: parameter,
		note_:      note,
	}
}

// Functions

// INSTANCE METHODS

// Target

type parameters_ struct {
	parameters_ col.ListLike[ParameterLike]
	note_       string
	parameter_  ParameterLike
}

// Attributes

func (v *parameters_) GetParameters() col.ListLike[ParameterLike] {
	return v.parameters_
}

func (v *parameters_) GetNote() string {
	return v.note_
}

// Public

// Private
