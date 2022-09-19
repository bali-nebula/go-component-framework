/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"regexp"
)

// PATTERN INTERFACE

// This constructor attempts to create a new regular expression pattern from the
// specified formatted string. It returns a pattern value and whether or not the
// string contained a valid regular expression.
// For valid string formats for this type see `../abstractions/language.go`.
func PatternFromString(v string) (Pattern, bool) {
	var pattern, ok = stringToPattern(v)
	return Pattern(pattern), ok
}

// This type defines the methods associated with a character pattern that
// extends the native Go string type and represents the regular expression
// corresponding to that pattern.
type Pattern string

// LEXICAL INTERFACE

// This method returns the canonical string for this element.
func (v Pattern) AsString() string {
	return string(v)
}

// This method implements the standard Go Stringer interface.
func (v Pattern) String() string {
	return v.AsString()
}

// MATCHABLE INTERFACE

// This method determines whether or not this pattern matches the specified text
// string.
func (v Pattern) MatchesText(text string) bool {
	var matched, _ = regexp.MatchString(patternToRegexp(v), text)
	return matched
}

// This method returns an array of strings containing all matching strings and
// substrings found in the specified text.
func (v Pattern) GetMatches(text string) []string {
	var scanner = regexp.MustCompile(patternToRegexp(v))
	return scanner.FindStringSubmatch(text)
}

// PRIVATE FUNCTIONS

// This function parses a pattern string and returns the corresponding regular
// expression. We take advantage of the Go regexp package to do the validation
// of the regular expression string.
func stringToPattern(v string) (string, bool) {
	var pattern string
	var ok = true
	var matches = abstractions.ScanPattern([]byte(v))
	if len(matches) == 0 {
		ok = false
	} else {
		pattern = matches[0]
	}
	return pattern, ok
}

// This function removes the delimiters from a pattern to expose the raw regexp
// string.
func patternToRegexp(v Pattern) string {
	var re string
	var pattern = string(v)
	switch pattern {
	case `none`:
		re = `^none$`
	case `any`:
		re = `.*`
	default:
		re = pattern[1 : len(pattern)-2]
	}
	return re
}
