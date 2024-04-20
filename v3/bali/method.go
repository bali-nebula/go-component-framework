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

var methodClass = &methodClass_{
	// This class has no private constants to initialize.
}

// Function

func Method() MethodClassLike {
	return methodClass
}

// CLASS METHODS

// Target

type methodClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *methodClass_) MakeWithIdentifier(identifier string) MethodLike {
	return &method_{
		identifier_: identifier,
	}
}

// Functions

// INSTANCE METHODS

// Target

type method_ struct {
	identifier_ string
}

// Attributes

func (v *method_) GetIdentifier() string {
	return v.identifier_
}

// Public

// Private
