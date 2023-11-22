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
	mat "math"
)

// CLASS DEFINITIONS

// This private type implements the FloatLike interface.  It extends the native
// Go `float64` type.
type float_ float64

// This private type defines the structure associated with the class constants
// and class functions for the float elements.
type floats_ struct {
	// This class has no class constants.
}

// CLASS CONSTANTS

// This class constant represents the minimum value for a float endpoint.
func (c *floats_) MinimumValue() abs.FloatLike {
	return float_(mat.Inf(-1))
}

// This class constant represents the maximum value for a float endpoint.
func (c *floats_) MaximumValue() abs.FloatLike {
	return float_(mat.Inf(1))
}

// CLASS CONSTRUCTORS

// This constructor creates a new float element from the specified float.
func (c *floats_) FromFloat(float float64) abs.FloatLike {
	return float_(float)
}

// This constructor creates a new float element from the specified string.
func (c *floats_) FromString(string_ string) abs.FloatLike {
	var matches = uti.FloatMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a float from an invalid string: %v", string_)
		panic(message)
	}
	var float = floatFromString(matches[0])
	return float_(float)
}

// CLASS METHODS

// Continuous Interface

// This method returns a float value for this floating point number.
func (v float_) AsFloat() float64 {
	return float64(v)
}

// This method determines whether or not this floating point number is zero.
func (v float_) IsZero() bool {
	return v == 0
}

// This method determines whether or not this floating point number is infinite.
func (v float_) IsInfinite() bool {
	return mat.IsInf(float64(v), 0)
}

// This method determines whether or not this floating point number is undefined.
func (v float_) IsUndefined() bool {
	return mat.IsNaN(float64(v))
}

// Lexical Interface

// This method returns a string value for this lexical element.
func (v float_) AsString() string {
	var string_ = stringFromFloat(v.AsFloat())
	return string_
}

// Polarized Interface

// This method determines whether or not this floating point number is negative.
func (v float_) IsNegative() bool {
	return v < 0
}
