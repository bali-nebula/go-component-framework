/*
................................................................................
.    Copyright (c) 2009-2025 Crater Dog Technologies.  All Rights Reserved.    .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
┌────────────────────────────────── WARNING ───────────────────────────────────┐
│                 This class file was automatically generated.                 │
│                     Any updates to it may be overwritten.                    │
└──────────────────────────────────────────────────────────────────────────────┘
*/

package ast

import (
	uti "github.com/craterdog/go-missing-utilities/v2"
)

// CLASS INTERFACE

// Access Function

func NoteLineClass() NoteLineClassLike {
	return noteLineClass()
}

// Constructor Methods

func (c *noteLineClass_) NoteLine(
	note string,
) NoteLineLike {
	if uti.IsUndefined(note) {
		panic("The \"note\" attribute is required by this class.")
	}
	var instance = &noteLine_{
		// Initialize the instance attributes.
		note_: note,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *noteLine_) GetClass() NoteLineClassLike {
	return noteLineClass()
}

// Attribute Methods

func (v *noteLine_) GetNote() string {
	return v.note_
}

// PROTECTED INTERFACE

// Instance Structure

type noteLine_ struct {
	// Declare the instance attributes.
	note_ string
}

// Class Structure

type noteLineClass_ struct {
	// Declare the class constants.
}

// Class Reference

func noteLineClass() *noteLineClass_ {
	return noteLineClassReference_
}

var noteLineClassReference_ = &noteLineClass_{
	// Initialize the class constants.
}
