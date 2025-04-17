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

func CommentLineClass() CommentLineClassLike {
	return commentLineClass()
}

// Constructor Methods

func (c *commentLineClass_) CommentLine(
	comment string,
) CommentLineLike {
	if uti.IsUndefined(comment) {
		panic("The \"comment\" attribute is required by this class.")
	}
	var instance = &commentLine_{
		// Initialize the instance attributes.
		comment_: comment,
	}
	return instance
}

// INSTANCE INTERFACE

// Principal Methods

func (v *commentLine_) GetClass() CommentLineClassLike {
	return commentLineClass()
}

// Attribute Methods

func (v *commentLine_) GetComment() string {
	return v.comment_
}

// PROTECTED INTERFACE

// Instance Structure

type commentLine_ struct {
	// Declare the instance attributes.
	comment_ string
}

// Class Structure

type commentLineClass_ struct {
	// Declare the class constants.
}

// Class Reference

func commentLineClass() *commentLineClass_ {
	return commentLineClassReference_
}

var commentLineClassReference_ = &commentLineClass_{
	// Initialize the class constants.
}
