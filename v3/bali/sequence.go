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

var sequenceClass = &sequenceClass_{
	// This class has no private constants to initialize.
}

// Function

func Sequence() SequenceClassLike {
	return sequenceClass
}

// CLASS METHODS

// Target

type sequenceClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *sequenceClass_) MakeWithExpression(expression ExpressionLike) SequenceLike {
	return &sequence_{
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type sequence_ struct {
	expression_ ExpressionLike
}

// Attributes

func (v *sequence_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
