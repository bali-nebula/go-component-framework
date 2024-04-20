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

var argumentsClass = &argumentsClass_{
	// This class has no private constants to initialize.
}

// Function

func Arguments() ArgumentsClassLike {
	return argumentsClass
}

// CLASS METHODS

// Target

type argumentsClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *argumentsClass_) MakeWithArguments(arguments col.ListLike[ArgumentLike]) ArgumentsLike {
	return &arguments_{
		arguments_: arguments,
	}
}

// Functions

// INSTANCE METHODS

// Target

type arguments_ struct {
	arguments_ col.ListLike[ArgumentLike]
}

// Attributes

func (v *arguments_) GetArguments() col.ListLike[ArgumentLike] {
	return v.arguments_
}

// Public

// Private
