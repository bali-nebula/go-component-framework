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

var complementClass = &complementClass_{
	// This class has no private constants to initialize.
}

// Function

func Complement() ComplementClassLike {
	return complementClass
}

// CLASS METHODS

// Target

type complementClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *complementClass_) MakeWithExpression(expression ExpressionLike) ComplementLike {
	return &complement_{
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type complement_ struct {
	expression_ ExpressionLike
}

// Attributes

func (v *complement_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
