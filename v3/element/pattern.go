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

import ()

// CLASS ACCESS

// Reference

var patternClass = &patternClass_{
	// This class has no private constants to initialize.
}

// Function

func Pattern() PatternClassLike {
	return patternClass
}

// CLASS METHODS

// Target

type patternClass_ struct {
	none_ PatternLike
	any_  PatternLike
}

// Constants

func (c *patternClass_) None() PatternLike {
	return c.none_
}

func (c *patternClass_) Any() PatternLike {
	return c.any_
}

// Constructors

func (c *patternClass_) MakeFromString(string_ string) PatternLike {
	return &pattern_{}
}

// Functions

// INSTANCE METHODS

// Target

type pattern_ struct {
	// TBA - Add private instance attributes.
}

// Attributes

// Lexical

func (v *pattern_) AsString() string {
	var result_ string
	// TBA - Implement the method.
	return result_
}

// Matchable

func (v *pattern_) MatchesText(text string) bool {
	var result_ bool
	// TBA - Implement the method.
	return result_
}

func (v *pattern_) GetMatches(text string) []string {
	var result_ []string
	// TBA - Implement the method.
	return result_
}

// Public

// Private
