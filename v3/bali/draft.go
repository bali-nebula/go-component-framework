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

var draftClass = &draftClass_{
	// This class has no private constants to initialize.
}

// Function

func Draft() DraftClassLike {
	return draftClass
}

// CLASS METHODS

// Target

type draftClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *draftClass_) MakeWithExpression(expression ExpressionLike) DraftLike {
	return &draft_{
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type draft_ struct {
	expression_ ExpressionLike
}

// Attributes

func (v *draft_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
