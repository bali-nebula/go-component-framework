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

// REAL IMPLEMENTATION

// This constructor returns the minimum value for a real endpoint.
func MinimumReal() abs.RealLike {
	return Real(mat.Inf(-1))
}

// This constructor returns the maximum value for a real endpoint.
func MaximumReal() abs.RealLike {
	return Real(mat.Inf(1))
}

// This type defines the methods associated with real endpoints. It extends the
// native Go float64 type.
type Real float64

// CONTINUOUS INTERFACE

// This method returns a real value for this real number.
func (v Real) AsReal() float64 {
	return float64(v)
}

// This method determines whether or not this real number is zero.
func (v Real) IsZero() bool {
	return v == 0
}

// This method determines whether or not this real number is infinite.
func (v Real) IsInfinite() bool {
	return mat.IsInf(float64(v), 0)
}

// This method determines whether or not this real number is undefined.
func (v Real) IsUndefined() bool {
	return mat.IsNaN(float64(v))
}

// POLARIZED INTERFACE

// This method determines whether or not this real number is negative.
func (v Real) IsNegative() bool {
	return v < 0
}
