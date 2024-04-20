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

var variableClass = &variableClass_{
	// This class has no private constants to initialize.
}

// Function

func Variable() VariableClassLike {
	return variableClass
}

// CLASS METHODS

// Target

type variableClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *variableClass_) MakeWithIdentifier(identifier string) VariableLike {
	return &variable_{
		identifier_: identifier,
	}
}

// Functions

// INSTANCE METHODS

// Target

type variable_ struct {
	identifier_ string
}

// Attributes

func (v *variable_) GetIdentifier() string {
	return v.identifier_
}

// Public

// Private
