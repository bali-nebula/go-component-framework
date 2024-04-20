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

var parameterClass = &parameterClass_{
	// This class has no private constants to initialize.
}

// Function

func Parameter() ParameterClassLike {
	return parameterClass
}

// CLASS METHODS

// Target

type parameterClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *parameterClass_) MakeWithAttributes(
	symbol string,
	component ComponentLike,
) ParameterLike {
	return &parameter_{
		symbol_:    symbol,
		component_: component,
	}
}

// Functions

// INSTANCE METHODS

// Target

type parameter_ struct {
	symbol_    string
	component_ ComponentLike
}

// Attributes

func (v *parameter_) GetSymbol() string {
	return v.symbol_
}

func (v *parameter_) GetComponent() ComponentLike {
	return v.component_
}

// Public

// Private
