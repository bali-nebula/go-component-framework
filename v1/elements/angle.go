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

// CONSTANT DEFINITIONS

// See "The Tau Manifesto" at https://tauday.com/tau-manifesto
const tau = 2.0 * mat.Pi

var Pi abs.AngleLike = Angle(mat.Pi)

var Tau abs.AngleLike = Angle(tau)

// ANGLE IMPLEMENTATION

// This constructor returns the minimum value for an angle.
func MinimumAngle() abs.AngleLike {
	return Angle(0)
}

// This constructor returns the maximum value for an angle.
func MaximumAngle() abs.AngleLike {
	return Tau
}

// This constructor creates a new angle from the specified value and normalizes
// the value to be in the allowed range for angles [0..τ).
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
	return Angle(lockPhase(float))
}

// This type defines the methods associated with angle elements. It extends the
// native Go float64 type and represents a radian based angle in the range
// [0..2π).
type Angle float64

// CONTINUOUS INTERFACE

// This method returns a real value for this continuous element.
func (v Angle) AsReal() float64 {
	return float64(v)
}

// This method determines whether or not this continuous element is zero.
func (v Angle) IsZero() bool {
	return v == 0 || v == Tau
}

// This method determines whether or not this continuous element is infinite.
func (v Angle) IsInfinite() bool {
	return false
}

// This method determines whether or not this continuous element is undefined.
func (v Angle) IsUndefined() bool {
	return false
}

// ANGLES LIBRARY

// This singleton creates a unique name space for the library functions for
// angle elements.
var Angles = &angles{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for angle elements.
type angles struct{}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified angle.
func (l *angles) Inverse(angle abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(angle.AsReal() - mat.Pi)
}

// This library function returns the sum of the specified angles.
func (l *angles) Sum(first, second abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(first.AsReal() + second.AsReal())
}

// This library function returns the difference of the specified angles.
func (l *angles) Difference(first, second abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(first.AsReal() - second.AsReal())
}

// This library function returns the specified angle scaled by the specified
// factor.
func (l *angles) Scaled(angle abs.AngleLike, factor float64) abs.AngleLike {
	return AngleFromFloat(angle.AsReal() * factor)
}

// ANGULAR INTERFACE

// This library function returns the complement of the specified angle. The
// complementary angles add up to π/2.
func (l *angles) Complement(angle abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(mat.Pi/2.0 - angle.AsReal())
}

// This library function returns the supplement of the specified angle. The
// supplementary angles add up to π.
func (l *angles) Supplement(angle abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(mat.Pi - angle.AsReal())
}

// This library function returns the conjugate of the specified angle. The
// conjugate angles add up to 2π (zero).
func (l *angles) Conjugate(angle abs.AngleLike) abs.AngleLike {
	return AngleFromFloat(-angle.AsReal())
}

// This library function returns the trigonometric cosine of the specified
// angle.
func (l *angles) Cosine(angle abs.AngleLike) float64 {
	switch angle.AsReal() {
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
		return mat.Cos(angle.AsReal())
	}
}

// This library function returns the angle whose trigonometric cosine is the
// specified distance along the x-axis.
func (l *angles) ArcCosine(x float64) abs.AngleLike {
	return AngleFromFloat(mat.Acos(x))
}

// This library function returns the trigonometric sine of the specified angle.
func (l *angles) Sine(angle abs.AngleLike) float64 {
	switch angle.AsReal() {
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
		return mat.Sin(angle.AsReal())
	}
}

// This library function returns the angle whose trigonometric sine is the
// specified distance along the y-axis.
func (l *angles) ArcSine(y float64) abs.AngleLike {
	return AngleFromFloat(mat.Asin(y))
}

// This library function returns the trigonometric tangent of the specified
// angle.
func (l *angles) Tangent(angle abs.AngleLike) float64 {
	switch angle.AsReal() {
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
		return mat.Tan(angle.AsReal())
	}
}

// This library function returns the angle whose trigonometric tangent is the
// specified ratio of the distances along the y-axis and x-axis.
func (l *angles) ArcTangent(x, y float64) abs.AngleLike {
	return AngleFromFloat(mat.Atan2(y, x))
}

// PRIVATE FUNCTIONS

// This function uses the single precision floating point range to lock a double
// precision phase angle onto 0, π/2, π, or 3π/2 if the angle falls outside the
// single precision range for these values. Otherwise, the phase angle is
// returned unchanged.
func lockPhase(v float64) float64 {
	var v32 float32 = float32(v)
	switch {
	case mat.Abs(v) <= 1.2246467991473515e-16:
		return 0
	case v32 == float32(0.5*mat.Pi):
		return 0.5 * mat.Pi
	case v32 == float32(mat.Pi):
		return mat.Pi
	case v32 == float32(1.5*mat.Pi):
		return 1.5 * mat.Pi
	default:
		return v
	}
}
