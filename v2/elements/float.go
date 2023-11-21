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
	stc "strconv"
)

// FLOAT ELEMENT CONSTRUCTORS

// This constructor creates a new float element from the specified float.
func FloatFromFloat(float float64) abs.FloatLike {
	return float_(float)
}

// This constructor creates a new float element from the specified string.
func FloatFromString(string_ string) abs.FloatLike {
	var matches = uti.FloatMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a float from an invalid string: %v", string_)
		panic(message)
	}
	var float = floatFromString(matches[0])
	return float_(float)
}

// This constructor returns the minimum value for a float endpoint.
func MinimumFloat() abs.FloatLike {
	return float_(mat.Inf(-1))
}

// This constructor returns the maximum value for a float endpoint.
func MaximumFloat() abs.FloatLike {
	return float_(mat.Inf(1))
}

// FLOAT ELEMENT METHODS

// This private type implements the FloatLike interface.  It extends the native
// Go `float64` type.
type float_ float64

// CONTINUOUS INTERFACE

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

// LEXICAL INTERFACE

// This method returns a string value for this lexical element.
func (v float_) AsString() string {
	var string_ = stringFromFloat(v.AsFloat())
	return string_
}

// POLARIZED INTERFACE

// This method determines whether or not this floating point number is negative.
func (v float_) IsNegative() bool {
	return v < 0
}

// PRIVATE FUNCTIONS

// This function converts a string into a real value.
func floatFromString(string_ string) float64 {
	var float float64
	switch string_ {
	case "+e", "e":
		float = mat.E
	case "-e":
		float = -mat.E
	case "+pi", "+π", "pi", "π":
		float = mat.Pi
	case "-pi", "-π":
		float = -mat.Pi
	case "+phi", "+φ", "phi", "φ":
		float = mat.Phi
	case "-phi", "-φ":
		float = -mat.Phi
	case "+tau", "+τ", "tau", "τ":
		float = mat.Pi * 2.0
	case "-tau", "-τ":
		float = -mat.Pi * 2.0
	case "+infinity", "+∞", "infinity", "∞":
		float = mat.Inf(1)
	case "-infinity", "-∞":
		float = mat.Inf(-1)
	default:
		float, _ = stc.ParseFloat(string_, 64)
	}
	return float
}

// This returns the string for the specified floating point number.
func stringFromFloat(float float64) string {
	var string_ string
	switch float {
	case mat.E:
		string_ = "e"
	case -mat.E:
		string_ = "-e"
	case mat.Pi:
		string_ = "π"
	case -mat.Pi:
		string_ = "-π"
	case mat.Phi:
		string_ = "φ"
	case -mat.Phi:
		string_ = "-φ"
	case mat.Pi * 2.0:
		string_ = "τ"
	case -mat.Pi * 2.0:
		string_ = "-τ"
	case mat.Inf(1):
		string_ = "∞"
	case mat.Inf(-1):
		string_ = "-∞"
	default:
		string_ = stc.FormatFloat(float, 'G', -1, 64)
	}
	return string_
}
