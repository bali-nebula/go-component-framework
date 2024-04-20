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

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var linesClass = &linesClass_{
	// This class has no private constants to initialize.
}

// Function

func Lines() LinesClassLike {
	return linesClass
}

// CLASS METHODS

// Target

type linesClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *linesClass_) MakeWithLines(lines col.ListLike[LineLike]) LinesLike {
	return &lines_{
		lines_: lines,
	}
}

func (c *linesClass_) MakeWithAttributes(
	line LineLike,
	note string,
) LinesLike {
	return &lines_{
		line_: line,
		note_: note,
	}
}

// Functions

// INSTANCE METHODS

// Target

type lines_ struct {
	lines_ col.ListLike[LineLike]
	note_  string
	line_  LineLike
}

// Attributes

func (v *lines_) GetLines() col.ListLike[LineLike] {
	return v.lines_
}

func (v *lines_) GetNote() string {
	return v.note_
}

// Public

// Private
