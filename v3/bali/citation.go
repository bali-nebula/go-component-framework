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

var citationClass = &citationClass_{
	// This class has no private constants to initialize.
}

// Function

func Citation() CitationClassLike {
	return citationClass
}

// CLASS METHODS

// Target

type citationClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *citationClass_) MakeWithExpression(expression ExpressionLike) CitationLike {
	return &citation_{
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type citation_ struct {
	expression_ ExpressionLike
}

// Attributes

func (v *citation_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
