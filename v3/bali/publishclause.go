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

var publishClauseClass = &publishClauseClass_{
	// This class has no private constants to initialize.
}

// Function

func PublishClause() PublishClauseClassLike {
	return publishClauseClass
}

// CLASS METHODS

// Target

type publishClauseClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *publishClauseClass_) MakeWithEvent(event EventLike) PublishClauseLike {
	return &publishClause_{
		event_: event,
	}
}

// Functions

// INSTANCE METHODS

// Target

type publishClause_ struct {
	event_ EventLike
}

// Attributes

func (v *publishClause_) GetEvent() EventLike {
	return v.event_
}

// Public

// Private
