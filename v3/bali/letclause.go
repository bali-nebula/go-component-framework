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

var letClauseClass = &letClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func LetClause() LetClauseClassLike {
	return letClauseClass
}

// CLASS METHODS

// Target

type letClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *letClauseClass_) MakeWithAttributes(
	recipient RecipientLike,
	expression ExpressionLike,
) LetClauseLike {
	return &letClause_{
		recipient_:  recipient,
		expression_: expression,
	}
}

// Functions

// INSTANCE METHODS

// Target

type letClause_ struct {
	recipient_  RecipientLike
	expression_ ExpressionLike
}

// Attributes

func (v *letClause_) GetRecipient() RecipientLike {
	return v.recipient_
}

func (v *letClause_) GetExpression() ExpressionLike {
	return v.expression_
}

// Public

// Private
