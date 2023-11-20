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

// REAL INTERFACE

// This constructor creates a new real element from the specified string.
func RealFromString(string_ string) abs.RealLike {
	var matches = uti.RealMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a real from an invalid string: %v", string_)
		panic(message)
	}
	var float float64
	var err error
	switch matches[0] {
	case "undefined":
		float = mat.NaN()
	case "+infinity", "+∞":
		float = mat.Inf(1)
	case "-infinity", "-∞":
		float = mat.Inf(-1)
	case "+pi", "pi", "-pi", "+phi", "phi", "-phi":
		float = floatFromString(matches[0])
	default:
		float, err = stc.ParseFloat(matches[0], 64)
		if err != nil {
			float = floatFromString(matches[0])
		}
	}
	return real_(float)
}

// This constructor returns the minimum value for a real endpoint.
func MinimumReal() abs.RealLike {
	return real_(mat.Inf(-1))
}

// This constructor returns the maximum value for a real endpoint.
func MaximumReal() abs.RealLike {
	return real_(mat.Inf(1))
}

// REAL IMPLEMENTATION

// This type defines the methods associated with real endpoints. It extends the
// native Go float64 type.
type real_ float64

// CONTINUOUS INTERFACE

// This method returns a real value for this real number.
func (v real_) AsReal() float64 {
	return float64(v)
}

// This method determines whether or not this real number is zero.
func (v real_) IsZero() bool {
	return v == 0
}

// This method determines whether or not this real number is infinite.
func (v real_) IsInfinite() bool {
	return mat.IsInf(float64(v), 0)
}

// This method determines whether or not this real number is undefined.
func (v real_) IsUndefined() bool {
	return mat.IsNaN(float64(v))
}

// LEXICAL INTERFACE

// This method returns a string value for this lexical element.
func (v real_) AsString() string {
	var string_ string
	switch {
	case v.IsZero():
		string_ = "0"
	case v.IsInfinite():
		if v.IsNegative() {
			string_ += "-"
		} else {
			string_ += "+"
		}
		string_ += "∞"
	case v.IsUndefined():
		string_ = "undefined"
	default:
		string_ = stringFromReal(v.AsReal())
	}
	return string_
}

// POLARIZED INTERFACE

// This method determines whether or not this real number is negative.
func (v real_) IsNegative() bool {
	return v < 0
}
