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

var resultClass = &resultClass_{
	// This class has no private constants to initialize.
}

// Function

func Result() ResultClassLike {
	return resultClass
}

// CLASS METHODS

// Target

type resultClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *resultClass_) MakeWithExpression(expression ExpressionLike) ResultLike {
	return &result_{
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type result_ struct {
	expression_ ExpressionLike
}

// Attributes

func (v *result_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
