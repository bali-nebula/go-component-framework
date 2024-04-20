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

var angleClass = &angleClass_{
	zero_: angle_(0.0),
	pi_:   angle_(mat.Pi),
	tau_:  angle_(2.0 * mat.Pi),
}

// Function

func Angle() AngleClassLike {
	return angleClass
}

// CLASS METHODS

// Target

type angleClass_ struct {
	zero_ angle_
	pi_   angle_
	tau_  angle_
}

// Constants

func (c *angleClass_) MinimumValue() AngleLike {
	return c.zero_
}

func (c *angleClass_) MaximumValue() AngleLike {
	return c.tau_
}

func (c *angleClass_) Zero() AngleLike {
	return c.zero_
}

func (c *angleClass_) Pi() AngleLike {
	return c.pi_
}

func (c *angleClass_) Tau() AngleLike {
	return c.tau_
}

// Constructors

func (c *angleClass_) MakeFromFloat(float float64) AngleLike {
	var angle = angleFromFloat(float)
	return angle
}

func (c *angleClass_) MakeFromString(string_ string) AngleLike {
	var matches = matchAngle(string_)
	var float = floatFromString(matches[1]) // Strip off the leading '~' character.
	var angle = angleFromFloat(float)
	return angle
}

// Functions

func (c *angleClass_) Inverse(angle AngleLike) AngleLike {
	return angleFromFloat(angle.AsFloat() - Angle().Pi().AsFloat())
}

func (c *angleClass_) Sum(
	first AngleLike,
	second AngleLike,
) AngleLike {
	return angleFromFloat(first.AsFloat() + second.AsFloat())
}

func (c *angleClass_) Difference(
	first AngleLike,
	second AngleLike,
) AngleLike {
	return angleFromFloat(first.AsFloat() - second.AsFloat())
}

func (c *angleClass_) Scaled(
	angle AngleLike,
	factor float64,
) AngleLike {
	return angleFromFloat(angle.AsFloat() * factor)
}

func (c *angleClass_) Complement(angle AngleLike) AngleLike {
	return angleFromFloat(Angle().Pi().AsFloat()/2.0 - angle.AsFloat())
}

func (c *angleClass_) Supplement(angle AngleLike) AngleLike {
	return angleFromFloat(Angle().Pi().AsFloat() - angle.AsFloat())
}

func (c *angleClass_) Conjugate(angle AngleLike) AngleLike {
	return angleFromFloat(-angle.AsFloat())
}

func (c *angleClass_) Cosine(angle AngleLike) float64 {
	var cosine float64
	switch angle.AsFloat() {
	case 0.0:
		cosine = 1.0
	case mat.Pi * 0.25:
		cosine = 0.5 * mat.Sqrt2
	case mat.Pi * 0.5:
		cosine = 0.0
	case mat.Pi * 0.75:
		cosine = -0.5 * mat.Sqrt2
	case mat.Pi:
		cosine = -1.0
	case mat.Pi * 1.25:
		cosine = -0.5 * mat.Sqrt2
	case mat.Pi * 1.5:
		cosine = 0.0
	case mat.Pi * 1.75:
		cosine = 0.5 * mat.Sqrt2
	case mat.Pi * 2.0:
		cosine = 1.0
	default:
		cosine = mat.Cos(angle.AsFloat())
	}
	return cosine
}

func (c *angleClass_) ArcCosine(x float64) AngleLike {
	return angleFromFloat(mat.Acos(x))
}

func (c *angleClass_) Sine(angle AngleLike) float64 {
	var sine float64
	switch angle.AsFloat() {
	case 0.0:
		sine = 0.0
	case mat.Pi * 0.25:
		sine = 0.5 * mat.Sqrt2
	case mat.Pi * 0.5:
		sine = 1.0
	case mat.Pi * 0.75:
		sine = 0.5 * mat.Sqrt2
	case mat.Pi:
		sine = 0.0
	case mat.Pi * 1.25:
		sine = -0.5 * mat.Sqrt2
	case mat.Pi * 1.5:
		sine = -1.0
	case mat.Pi * 1.75:
		sine = -0.5 * mat.Sqrt2
	case mat.Pi * 2.0:
		sine = 0.0
	default:
		sine = mat.Sin(angle.AsFloat())
	}
	return sine
}

func (c *angleClass_) ArcSine(y float64) AngleLike {
	return angleFromFloat(mat.Asin(y))
}

func (c *angleClass_) Tangent(angle AngleLike) float64 {
	var tangent float64
	switch angle.AsFloat() {
	case 0.0:
		tangent = 0.0
	case mat.Pi * 0.25:
		tangent = 1.0
	case mat.Pi * 0.5:
		tangent = mat.Inf(1)
	case mat.Pi * 0.75:
		tangent = -1.0
	case mat.Pi:
		tangent = 0.0
	case mat.Pi * 1.25:
		tangent = 1.0
	case mat.Pi * 1.5:
		tangent = mat.Inf(1)
	case mat.Pi * 1.75:
		tangent = -1.0
	case mat.Pi * 2.0:
		tangent = 0.0
	default:
		tangent = mat.Tan(angle.AsFloat())
	}
	return tangent
}

func (c *angleClass_) ArcTangent(
	x float64,
	y float64,
) AngleLike {
	return angleFromFloat(mat.Atan2(y, x))
}

// INSTANCE METHODS

// Target

type angle_ float64

// Continuous

func (v angle_) AsFloat() float64 {
	return float64(v)
}

func (v angle_) IsZero() bool {
	return v == angleClass.zero_ || v == angleClass.tau_
}

func (v angle_) IsInfinite() bool {
	return false
}

func (v angle_) IsUndefined() bool {
	return false
}

// Lexical

func (v angle_) AsString() string {
	var string_ string
	switch v {
	case angleClass.pi_:
		string_ = "~π"
	case angleClass.tau_:
		string_ = "~τ"
	default:
		string_ = "~" + stc.FormatFloat(float64(v), 'G', -1, 64)
	}
	return string_
}

// PACKAGE FUNCTIONS

// Private

func angleFromFloat(float float64) angle_ {
	float = normalizeValue(float)
	float = lockPhase(float)
	return angle_(float)
}

func matchAngle(string_ string) []string {
	var matches = []string{
		string_,
		string_[1:],
	}
	// TBA - Add the pattern matching code...
	return matches
}

func lockPhase(value float64) float64 {
	var pi = float64(angleClass.pi_)
	var value32 = float32(value)
	switch {
	case mat.Abs(value) <= 1.2246467991473515e-16:
		value = 0
	case value32 == float32(0.5*pi):
		value = 0.5 * pi
	case value32 == float32(pi):
		value = pi
	case value32 == float32(1.5*pi):
		value = 1.5 * pi
	}
	return value
}

func normalizeValue(value float64) float64 {
	var tau = float64(angleClass.tau_)
	if value < -tau || value >= tau {
		// Normalize the value to the range [-τ..τ).
		value = mat.Remainder(value, tau)
	}
	if value < 0.0 {
		// Normalize the value to the range [0..τ).
		value = value + tau
	}
	return value
}
