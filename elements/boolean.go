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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// CONSTANT DEFINITIONS

var False abs.BooleanLike = Boolean(false)

var True abs.BooleanLike = Boolean(true)

// BOOLEAN IMPLEMENTATION

// This type defines the methods associated with boolean elements. It extends
// the native Go bool type.
type Boolean bool

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
func (l *booleans) Not(boolean abs.BooleanLike) abs.BooleanLike {
	return Boolean(!boolean.AsBoolean())
}

// This library function returns the logical conjunction of the specified
// boolean elements.
func (l *booleans) And(first, second abs.BooleanLike) abs.BooleanLike {
	return Boolean(first.AsBoolean() && second.AsBoolean())
}

// This library function returns the logical material non-implication of the
// specified boolean elements.
func (l *booleans) Sans(first, second abs.BooleanLike) abs.BooleanLike {
	return Boolean(first.AsBoolean() && !second.AsBoolean())
}

// This library function returns the logical disjunction of the specified
// boolean elements.
func (l *booleans) Or(first, second abs.BooleanLike) abs.BooleanLike {
	return Boolean(first.AsBoolean() || second.AsBoolean())
}

// This library function returns the logical exclusive disjunction of the
// specified boolean elements.
func (l *booleans) Xor(first, second abs.BooleanLike) abs.BooleanLike {
	return Boolean((first.AsBoolean() && !second.AsBoolean()) || (!first.AsBoolean() && second.AsBoolean()))
}
