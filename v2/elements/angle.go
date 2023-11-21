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

// ANGLE ELEMENT CONSTANTS

// See "The Tau Manifesto" at https://tauday.com/tau-manifesto
const tau = 2.0 * mat.Pi

var Pi abs.AngleLike = angle_(mat.Pi)

var Tau abs.AngleLike = angle_(tau)

// ANGLE ELEMENT CONSTRUCTORS

// This constructor creates a new angle from the specified float value and
// normalizes the value to be in the allowed range for angles [0..τ).
func AngleFromFloat(float float64) abs.AngleLike {
	var tau = 2.0 * mat.Pi
	if float < -tau || float >= tau {
		// Normalize the angle to the range [-τ..τ).
		float = mat.Remainder(float, tau)
	}
	if float < 0.0 {
		// Normalize the angle to the range [0..τ).
		float = float + tau
	}
	var angle = angle_(lockPhase(float))
	return angle
}

// This constructor creates a new angle from the specified string value and
// normalizes the value to be in the allowed range for angles [0..τ).
func AngleFromString(string_ string) abs.AngleLike {
	var matches = uti.AngleMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct an angle from an invalid string: %v", string_)
		panic(message)
	}
	var float = floatFromString(matches[1]) // Strip off the leading '~' character.
	var angle = AngleFromFloat(float)
	return angle
}

// This constructor returns the minimum value for an angle.
func MinimumAngle() abs.AngleLike {
	return AngleFromFloat(0)
}

// This constructor returns the maximum value for an angle.
func MaximumAngle() abs.AngleLike {
	return Tau
}

// ANGLE ELEMENT METHODS

// This private type implements the AngleLike interface.  It extends the native
// Go `float64` type and represents a radian based angle in the range [0..2π).
type angle_ float64

// CONTINUOUS INTERFACE

// This method returns a real value for this continuous element.
func (v angle_) AsFloat() float64 {
	return float64(v)
}

// This method determines whether or not this continuous element is zero.
func (v angle_) IsZero() bool {
	return v == 0 || v == Tau
}

// This method determines whether or not this continuous element is infinite.
func (v angle_) IsInfinite() bool {
	return false
}

// This method determines whether or not this continuous element is undefined.
func (v angle_) IsUndefined() bool {
	return false
}

// LEXICAL INTERFACE

// This method returns a string value for this lexical element.
func (v angle_) AsString() string {
	var s string
	switch v {
	case Pi:
		s = "~π"
	case Tau:
		s = "~τ"
	default:
		s = "~" + stc.FormatFloat(float64(v), 'G', -1, 64)
	}
	return s
}

// ANGLE ELEMENT LIBRARY

// This singleton creates a unique name space for the library functions for
// angle elements.
var Angle = &angles_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for angle elements.
type angles_ struct{}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified angle.
func (l *angles_) Inverse(angle abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(angle.AsFloat() - mat.Pi)
}

// This library function returns the sum of the specified angles.
func (l *angles_) Sum(first, second abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(first.AsFloat() + second.AsFloat())
}

// This library function returns the difference of the specified angles.
func (l *angles_) Difference(first, second abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(first.AsFloat() - second.AsFloat())
}

// This library function returns the specified angle scaled by the specified
// factor.
func (l *angles_) Scaled(angle abs.AngleLike, factor float64) abs.AngleLike {
	return AngleFromFloat(angle.AsFloat() * factor)
}

// ANGULAR INTERFACE

// This library function returns the complement of the specified angle. The
// complementary angles add up to π/2.
func (l *angles_) Complement(angle abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(mat.Pi/2.0 - angle.AsFloat())
}

// This library function returns the supplement of the specified angle. The
// supplementary angles add up to π.
func (l *angles_) Supplement(angle abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(mat.Pi - angle.AsFloat())
}

// This library function returns the conjugate of the specified angle. The
// conjugate angles add up to 2π (zero).
func (l *angles_) Conjugate(angle abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(-angle.AsFloat())
}

// This library function returns the trigonometric cosine of the specified
// angle.
func (l *angles_) Cosine(angle abs.AngleLike) float64 {
	switch angle.AsFloat() {
	case 0.0:
		return 1.0
	case mat.Pi * 0.25:
		return 0.5 * mat.Sqrt2
	case mat.Pi * 0.5:
		return 0.0
	case mat.Pi * 0.75:
		return -0.5 * mat.Sqrt2
	case mat.Pi:
		return -1.0
	case mat.Pi * 1.25:
		return -0.5 * mat.Sqrt2
	case mat.Pi * 1.5:
		return 0.0
	case mat.Pi * 1.75:
		return 0.5 * mat.Sqrt2
	case tau:
		return 1.0
	default:
		return mat.Cos(angle.AsFloat())
	}
}

// This library function returns the angle whose trigonometric cosine is the
// specified distance along the x-axis.
func (l *angles_) ArcCosine(x float64) abs.AngleLike {
	return AngleFromFloat(mat.Acos(x))
}

// This library function returns the trigonometric sine of the specified angle.
func (l *angles_) Sine(angle abs.AngleLike) float64 {
	switch angle.AsFloat() {
	case 0.0:
		return 0.0
	case mat.Pi * 0.25:
		return 0.5 * mat.Sqrt2
	case mat.Pi * 0.5:
		return 1.0
	case mat.Pi * 0.75:
		return 0.5 * mat.Sqrt2
	case mat.Pi:
		return 0.0
	case mat.Pi * 1.25:
		return -0.5 * mat.Sqrt2
	case mat.Pi * 1.5:
		return -1.0
	case mat.Pi * 1.75:
		return -0.5 * mat.Sqrt2
	case tau:
		return 0.0
	default:
		return mat.Sin(angle.AsFloat())
	}
}

// This library function returns the angle whose trigonometric sine is the
// specified distance along the y-axis.
func (l *angles_) ArcSine(y float64) abs.AngleLike {
	return AngleFromFloat(mat.Asin(y))
}

// This library function returns the trigonometric tangent of the specified
// angle.
func (l *angles_) Tangent(angle abs.AngleLike) float64 {
	switch angle.AsFloat() {
	case 0.0:
		return 0.0
	case mat.Pi * 0.25:
		return 1.0
	case mat.Pi * 0.5:
		return mat.Inf(1)
	case mat.Pi * 0.75:
		return -1.0
	case mat.Pi:
		return 0.0
	case mat.Pi * 1.25:
		return 1.0
	case mat.Pi * 1.5:
		return mat.Inf(1)
	case mat.Pi * 1.75:
		return -1.0
	case tau:
		return 0.0
	default:
		return mat.Tan(angle.AsFloat())
	}
}

// This library function returns the angle whose trigonometric tangent is the
// specified ratio of the distances along the y-axis and x-axis.
func (l *angles_) ArcTangent(x, y float64) abs.AngleLike {
	return AngleFromFloat(mat.Atan2(y, x))
}
