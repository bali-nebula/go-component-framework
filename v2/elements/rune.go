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
	mat "math"
)

// RUNE IMPLEMENTATION

// This constructor returns the minimum value for a rune endpoint.
func MinimumRune() abs.RuneLike {
	return Rune(0)
}

// This constructor returns the maximum value for a rune endpoint.
func MaximumRune() abs.RuneLike {
	return Rune(mat.MaxInt32)
}

// This type defines the methods associated with rune endpoints. It extends the
// native Go int32 type.
type Rune int32

// DISCRETE INTERFACE

// This method returns a boolean value for this rune.
func (v Rune) AsBoolean() bool {
	return v > -1
}

// This method returns an integer value for this rune.
func (v Rune) AsInteger() int {
	return int(v)
}
