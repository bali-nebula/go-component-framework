/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	mat "math"
)

// ANGLE INTERFACE

// This constructor creates a new angle from the specified value and normalizes
// the value to be in the allowed range for angles [0..2π).
func AngleFromFloat(v float64) Angle {
	var twoPi = 2.0 * mat.Pi
	if v < -twoPi || v >= twoPi {
		// Normalize the angle to the range [-2π..2π).
		v = mat.Remainder(v, twoPi)
	}
	if v < 0.0 {
		// Normalize the angle to the range [0..2π).
		v = v + twoPi
	}
	return Angle(lockPhase(v))
}

// This type defines the methods associated with angle elements. It extends the
// native Go float64 type and represents a radian based angle in the range
// [0..2π).
type Angle float64

// CONTINUOUS INTERFACE

// This method determines whether or not this angle is zero.
func (v Angle) IsZero() bool {
	return v == 0
}

// This method returns a real value for this continuous component.
func (v Angle) AsReal() float64 {
	return float64(v)
}

// ANGLES LIBRARY

var Pi = Angle(complex(mat.Pi, 0.0))

// This singleton creates a unique name space for the library functions for
// angle elements.
var Angles = &angles{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for angle elements.
type angles struct{}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified angle.
func (l *angles) Inverse(angle Angle) Angle {
	return AngleFromFloat(float64(angle) + mat.Pi)
}

// This library function returns the sum of the specified angles.
func (l *angles) Sum(first, second Angle) Angle {
	return AngleFromFloat(float64(first + second))
}

// This library function returns the difference of the specified angles.
func (l *angles) Difference(first, second Angle) Angle {
	return AngleFromFloat(float64(first - second))
}

// This library function returns the specified angle scaled by the specified
// factor.
func (l *angles) Scaled(angle Angle, factor float64) Angle {
	return AngleFromFloat(float64(angle) * factor)
}

// ANGULAR INTERFACE

// This library function returns the complement of the specified angle. The
// complementary angles add up to π/2.
func (l *angles) Complement(angle Angle) Angle {
	return AngleFromFloat(float64(mat.Pi/2.0 - angle))
}

// This library function returns the supplement of the specified angle. The
// supplementary angles add up to π.
func (l *angles) Supplement(angle Angle) Angle {
	return AngleFromFloat(float64(mat.Pi - angle))
}

// This library function returns the conjugate of the specified angle. The
// conjugate angles add up to 2π (zero).
func (l *angles) Conjugate(angle Angle) Angle {
	return AngleFromFloat(float64(-angle))
}

// This library function returns the trigonometric cosine of the specified
// angle.
func (l *angles) Cosine(angle Angle) float64 {
	switch angle {
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
	default:
		return mat.Cos(float64(angle))
	}
}

// This library function returns the angle whose trigonometric cosine is the
// specified distance along the x-axis.
func (l *angles) ArcCosine(x float64) Angle {
	return AngleFromFloat(mat.Acos(x))
}

// This library function returns the trigonometric sine of the specified angle.
func (l *angles) Sine(angle Angle) float64 {
	switch angle {
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
	default:
		return mat.Sin(float64(angle))
	}
}

// This library function returns the angle whose trigonometric sine is the
// specified distance along the y-axis.
func (l *angles) ArcSine(y float64) Angle {
	return AngleFromFloat(mat.Asin(y))
}

// This library function returns the trigonometric tangent of the specified
// angle.
func (l *angles) Tangent(angle Angle) float64 {
	switch angle {
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
	default:
		return mat.Tan(float64(angle))
	}
}

// This library function returns the angle whose trigonometric tangent is the
// specified ratio of the distances along the y-axis and x-axis.
func (l *angles) ArcTangent(x, y float64) Angle {
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
