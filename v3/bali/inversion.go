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

var inversionClass = &inversionClass_{
	// This class has no private constants to initialize.
}

// Function

func Inversion() InversionClassLike {
	return inversionClass
}

// CLASS METHODS

// Target

type inversionClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *inversionClass_) MakeWithExpression(expression ExpressionLike) InversionLike {
	return &inversion_{
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type inversion_ struct {
	expression_ ExpressionLike
}

// Attributes

func (v *inversion_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
