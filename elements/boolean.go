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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	str "strconv"
)

// BOOLEAN INTERFACE

// This constructor attempts to create a new boolean from the specified
// formatted string. It returns a boolean value and whether or not the
// string contained a valid boolean.
// For valid string formats for this type see `../abstractions/language.go`.
func BooleanFromString(v string) (Boolean, bool) {
	var boolean, ok = stringToBoolean(v)
	return Boolean(boolean), ok
}

// This type defines the methods associated with boolean elements. It extends
// the native Go bool type.
type Boolean bool

// LEXICAL INTERFACE

// This method returns the canonical string for this element.
func (v Boolean) AsString() string {
	return str.FormatBool(bool(v))
}

// This method implements the standard Go Stringer interface.
func (v Boolean) String() string {
	return v.AsString()
}

// DISCRETE INTERFACE

// This method returns a boolean value for this discrete element.
func (v Boolean) AsBoolean() bool {
	return bool(v)
}

// This method returns an integer value for this discrete element.
func (v Boolean) AsInteger() int {
	if v {
		return 1
	}
	return 0
}

// BOOLEANS LIBRARY

// This singleton creates a unique name space for the library functions for
// boolean elements.
var Booleans = &booleans{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for boolean elements.
type booleans struct{}

// LOGICAL INTERFACE

// This library function returns the logical inverse of the specified boolean.
func (l *booleans) Not(boolean Boolean) Boolean {
	return !boolean
}

// This library function returns the logical conjunction of the specified
// boolean elements.
func (l *booleans) And(first, second Boolean) Boolean {
	return first && second
}

// This library function returns the logical material non-implication of the
// specified boolean elements.
func (l *booleans) Sans(first, second Boolean) Boolean {
	return first && !second
}

// This library function returns the logical disjunction of the specified
// boolean elements.
func (l *booleans) Or(first Boolean, second Boolean) Boolean {
	return first || second
}

// This library function returns the logical exclusive disjunction of the
// specified boolean elements.
func (l *booleans) Xor(first Boolean, second Boolean) Boolean {
	return (first && !second) || (!first && second)
}

// PRIVATE FUNCTIONS

// This function parses a boolean string and returns the corresponding boolean
// value and whether or not the string contained a valid boolean.
func stringToBoolean(v string) (bool, bool) {
	var boolean bool
	var ok = true
	var matches = abs.ScanBoolean([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		var err error
		boolean, err = str.ParseBool(matches[0])
		if err != nil {
			ok = false
		}
	}
	return boolean, ok
}
