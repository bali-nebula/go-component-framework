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
	abs "github.com/bali-nebula/go-component-framework/v1/abstractions"
	mat "math"
)

// INTEGER IMPLEMENTATION

// This constructor returns the minimum value for an integer endpoint.
func MinimumInteger() abs.IntegerLike {
	return Integer(0)
}

// This constructor returns the maximum value for an integer endpoint.
func MaximumInteger() abs.IntegerLike {
	return Integer(mat.MaxInt)
}

// This type defines the methods associated with integer endpoints. It extends the
// native Go int type.
type Integer int

// DISCRETE INTERFACE

// This method returns a boolean value for this rune.
func (v Integer) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this rune.
func (v Integer) AsInteger() int {
	return int(v)
}

// POLARIZED INTERFACE

// This method determines whether or not this rune is negative.
func (v Integer) IsNegative() bool {
	return v < 0
}
