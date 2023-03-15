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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	reg "regexp"
)

// CONSTANT DEFINITIONS

var Any abs.PatternLike = Pattern(`.*`)

var None abs.PatternLike = Pattern(`^none$`)

// PATTERN IMPLEMENTATION

// This constructor creates a new pattern from the specified string.
func PatternFromString(string_ string) abs.PatternLike {
	reg.MustCompile(string_)
	return Pattern(string_)
}

// This type defines the methods associated with a character pattern that
// extends the native Go string type and represents the regular expression
// corresponding to that pattern.
type Pattern string

// QUANTIZED INTERFACE

// This method returns a string value for this lexical element.
func (v Pattern) AsString() string {
	return string(v)
}

// MATCHABLE INTERFACE

// This method determines whether or not this pattern matches the specified text
// string.
func (v Pattern) MatchesText(text string) bool {
	var matched, _ = reg.MatchString(string(v), text)
	return matched
}

// This method returns an array of strings containing all matching strings and
// substrings found in the specified text.
func (v Pattern) GetMatches(text string) []string {
	var scanner = reg.MustCompile(string(v))
	return scanner.FindStringSubmatch(text)
}
