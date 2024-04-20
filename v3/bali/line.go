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

var lineClass = &lineClass_{
	// This class has no private constants to initialize.
}

// Function

func Line() LineClassLike {
	return lineClass
}

// CLASS METHODS

// Target

type lineClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *lineClass_) MakeWithAttributes(
	annotation AnnotationLike,
	statement StatementLike,
) LineLike {
	return &line_{
		annotation_: annotation,
		statement_:  statement,
	}
}

// Functions

// INSTANCE METHODS

// Target

type line_ struct {
	annotation_ AnnotationLike
	statement_  StatementLike
}

// Attributes

func (v *line_) GetAnnotation() AnnotationLike {
	return v.annotation_
}

func (v *line_) GetStatement() StatementLike {
	return v.statement_
}

// Public

// Private
