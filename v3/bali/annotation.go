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

var annotationClass = &annotationClass_{
	// This class has no private constants to initialize.
}

// Function

func Annotation() AnnotationClassLike {
	return annotationClass
}

// CLASS METHODS

// Target

type annotationClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *annotationClass_) MakeWithAttributes(
	note string,
	comment string,
) AnnotationLike {
	return &annotation_{
		note_:    note,
		comment_: comment,
	}
}

// Functions

// INSTANCE METHODS

// Target

type annotation_ struct {
	note_    string
	comment_ string
}

// Attributes

func (v *annotation_) GetNote() string {
	return v.note_
}

func (v *annotation_) GetComment() string {
	return v.comment_
}

// Public

// Private
