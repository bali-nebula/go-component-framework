/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element

import (
	mat "math"
	stc "strconv"
)

// CLASS ACCESS

// Reference

var floatClass = &floatClass_{
	minimumValue_: float_(mat.Inf(-1)),
	maximumValue_: float_(mat.Inf(1)),
}

// Function

func Float() FloatClassLike {
	return floatClass
}

// CLASS METHODS

// Target

type floatClass_ struct {
	minimumValue_ float_
	maximumValue_ float_
}

// Constants

func (c *floatClass_) MinimumValue() FloatLike {
	return c.minimumValue_
}

func (c *floatClass_) MaximumValue() FloatLike {
	return c.maximumValue_
}

// Constructors

func (c *floatClass_) MakeFromFloat(float float64) FloatLike {
	return float_(float)
}

func (c *floatClass_) MakeFromString(string_ string) FloatLike {
	var matches = matchFloat(string_)
	var float = floatFromString(matches[0])
	return float_(float)
}

// INSTANCE METHODS

// Target

type float_ float64

// Continuous

func (v float_) AsFloat() float64 {
	return float64(v)
}

func (v float_) IsZero() bool {
	return v == 0.0
}

func (v float_) IsInfinite() bool {
	return mat.IsInf(float64(v), 0)
}

func (v float_) IsUndefined() bool {
	return mat.IsNaN(float64(v))
}

// Lexical

func (v float_) AsString() string {
	return stringFromFloat(v.AsFloat())
}

// Polarized

func (v float_) IsNegative() bool {
	return v < 0
}

// PACKAGE FUNCTIONS

// Private

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
		var err error
		float, err = stc.ParseFloat(string_, 64)
		if err != nil {
			panic(err)
		}
	}
	return float
}

func matchFloat(string_ string) []string {
	var matches = []string{
		string_,
	}
	// TBA - Add the pattern matching code...
	return matches
}

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
