/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element

import (
	sts "strings"
)

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

// Constructors

func (c *citationClass_) MakeFromString(string_ string) CitationLike {
	var matches = matchCitation(string_)
	var citation = citation_(matches[0])
	return citation
}

// INSTANCE METHODS

// Target

type citation_ string

// Lexical

func (v citation_) AsString() string {
	return string(v)
}

// Named

func (v citation_) GetName() string {
	var index = sts.LastIndex(string(v), "/")
	return string(v[:index])
}

// Versioned

func (v citation_) GetVersion() string {
	var index = sts.LastIndex(string(v), "/") + 1 // Skip the '/' character.
	return string(v[index:])
}

// PACKAGE FUNCTIONS

// Private

func matchCitation(string_ string) []string {
	var matches = []string{
		string_,
	}
	// TBA - Add the pattern matching code...
	return matches
}
