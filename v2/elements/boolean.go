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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
)

// CONSTANT DEFINITIONS

var False abs.BooleanLike = boolean_(false)

var True abs.BooleanLike = boolean_(true)

// BOOLEAN INTERFACE

// This constructor creates a new boolean from the specified bool value.
func BooleanFromBoolean(bool_ bool) abs.BooleanLike {
	return boolean_(bool_)
}

// This constructor creates a new boolean from the specified string value.
func BooleanFromString(string_ string) abs.BooleanLike {
	var matches = uti.BooleanMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a boolean from an invalid string: %v", string_)
		panic(message)
	}
	var bool_ = matches[0] == "true"
	var boolean = BooleanFromBoolean(bool_)
	return boolean
}

// BOOLEAN IMPLEMENTATION

// This type defines the methods associated with boolean elements. It extends
// the native Go bool type.
type boolean_ bool

// DISCRETE INTERFACE

// This method returns a boolean value for this discrete element.
func (v boolean_) AsBoolean() bool {
	return bool(v)
}

// This method returns an integer value for this discrete element.
func (v boolean_) AsInteger() int {
	if v {
		return 1
	}
	return 0
}

// LEXICAL INTERFACE

// This method returns a string value for this lexical element.
func (v boolean_) AsString() string {
	var s string
	switch v {
	case true:
		s = "true"
	case false:
		s = "false"
	}
	return s
}

// BOOLEAN LIBRARY

// This singleton creates a unique name space for the library functions for
// boolean elements.
var Boolean = &booleans_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for boolean elements.
type booleans_ struct{}

// LOGICAL INTERFACE

// This library function returns the logical inverse of the specified boolean.
func (l *booleans_) Not(boolean abs.BooleanLike) abs.BooleanLike {
	return BooleanFromBoolean(!boolean.AsBoolean())
}

// This library function returns the logical conjunction of the specified
// boolean elements.
func (l *booleans_) And(first, second abs.BooleanLike) abs.BooleanLike {
	return BooleanFromBoolean(first.AsBoolean() && second.AsBoolean())
}

// This library function returns the logical material non-implication of the
// specified boolean elements.
func (l *booleans_) Sans(first, second abs.BooleanLike) abs.BooleanLike {
	return BooleanFromBoolean(first.AsBoolean() && !second.AsBoolean())
}

// This library function returns the logical disjunction of the specified
// boolean elements.
func (l *booleans_) Or(first, second abs.BooleanLike) abs.BooleanLike {
	return BooleanFromBoolean(first.AsBoolean() || second.AsBoolean())
}

// This library function returns the logical exclusive disjunction of the
// specified boolean elements.
func (l *booleans_) Xor(first, second abs.BooleanLike) abs.BooleanLike {
	return Boolean.Or(Boolean.Sans(first, second), Boolean.Sans(second, first))
}
