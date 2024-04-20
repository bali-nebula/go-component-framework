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

var eventClass = &eventClass_{
	// This class has no private constants to initialize.
}

// Function

func Event() EventClassLike {
	return eventClass
}

// CLASS METHODS

// Target

type eventClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *eventClass_) MakeWithExpression(expression ExpressionLike) EventLike {
	return &event_{
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type event_ struct {
	expression_ ExpressionLike
}

// Attributes

func (v *event_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
