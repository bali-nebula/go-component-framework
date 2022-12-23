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

// BOOLEAN INTERFACE

// This type defines the methods associated with boolean elements. It extends
// the native Go bool type.
type Boolean bool

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
