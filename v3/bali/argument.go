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

var argumentClass = &argumentClass_{
	// This class has no private constants to initialize.
}

// Function

func Argument() ArgumentClassLike {
	return argumentClass
}

// CLASS METHODS

// Target

type argumentClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *argumentClass_) MakeWithExpression(expression ExpressionLike) ArgumentLike {
	return &argument_{
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type argument_ struct {
	expression_ ExpressionLike
}

// Attributes

func (v *argument_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
