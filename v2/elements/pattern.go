/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	fmt "fmt"
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	reg "regexp"
)

// CLASS DEFINITIONS

// This private type implements the PatternLike interface.  It extends the
// native Go `string` type and represents the regular expression corresponding
// to that pattern.
type pattern_ string

// This private type defines the structure associated with the class constants
// and class functions for the pattern elements.
type patternClass_ struct {
	any_ PatternLike
	none PatternLike
}

// CLASS CONSTANTS

// This class constant represents the pattern that matches any string.
func (c *patternClass_) Any() PatternLike {
	return c.any_
}

// This class constant represents the pattern that only matches the "none"
// string.
func (c *patternClass_) None() PatternLike {
	return c.none
}

// CLASS CONSTRUCTORS

// This constructor creates a new pattern element from the specified string.
func (c *patternClass_) FromString(string_ string) PatternLike {
	var matches = uti.PatternMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a pattern from an invalid string: %v", string_)
		panic(message)
	}
	var regex string
	switch matches[0] {
	case `none`:
		regex = `^none$`
	case `any`:
		regex = `.*`
	default:
		regex = matches[1] // Strip off the '"' and '"?' delimiters.
		reg.MustCompile(regex)
	}
	var pattern = pattern_(regex)
	return pattern
}

// CLASS METHODS

// Lexical Interface

// This method returns a string value for this lexical element.
func (v pattern_) AsString() string {
	var string_ string
	switch string(v) {
	case `^none$`:
		string_ = `none`
	case `.*`:
		string_ = `any`
	default:
		string_ = `"` + string(v) + `"?`
	}
	return string_
}

// Matchable Interface

// This method determines whether or not this pattern matches the specified text
// string.
func (v pattern_) MatchesText(text string) bool {
	var matched, _ = reg.MatchString(string(v), text)
	return matched
}

// This method returns an array of strings containing all matching strings and
// substrings found in the specified text.
func (v pattern_) GetMatches(text string) []string {
	var scanner = reg.MustCompile(string(v))
	return scanner.FindStringSubmatch(text)
}
