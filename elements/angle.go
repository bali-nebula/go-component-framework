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
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"math"
	"strconv"
)

// ANGLE INTERFACE

// This constructor creates a new angle from the specified value and normalizes
// the value to be in the allowed range for angles [0..2π).
func AngleFromFloat(v float64) Angle {
	var twoPi = 2.0 * math.Pi
	if v < -twoPi || v >= twoPi {
		// Normalize the angle to the range [-2π..2π).
		v = math.Remainder(v, twoPi)
	}
	if v < 0.0 {
		// Normalize the angle to the range [0..2π).
		v = v + twoPi
	}
	return Angle(lockPhase(v))
}

// This constructor attempts to create a new angle from the specified formatted
// string. It returns an angle value and whether or not the string contained a
// valid angle.
// For valid string formats for this type see `../abstractions/language.go`.
func AngleFromString(v string) (Angle, bool) {
	var angle, ok = stringToAngle(v)
	return AngleFromFloat(angle), ok // This normalizes the angle.
}

// This type defines the methods associated with angle elements. It extends the
// native Go float64 type and represents a radian based angle in the range
// [0..2π).
type Angle float64

// LEXICAL INTERFACE

// This method returns the canonical string for this element.
func (v Angle) AsString() string {
	if float64(v) == math.Pi {
		return "~π"
	}
	return "~" + strconv.FormatFloat(float64(v), 'G', -1, 64)
}

// This method implements the standard Go Stringer interface.
func (v Angle) String() string {
	return v.AsString()
}

// CONTINUOUS INTERFACE

// This method returns a real value for this continuous component.
func (v Angle) AsReal() float64 {
	return float64(v)
}

// ANGLES LIBRARY

var Pi = Angle(complex(math.Pi, 0.0))

// This singleton creates a unique name space for the library functions for
// angle elements.
var Angles = &angles{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for angle elements.
type angles struct{}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified angle.
func (l *angles) Inverse(angle Angle) Angle {
	return AngleFromFloat(float64(angle) + math.Pi)
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
	return AngleFromFloat(float64(math.Pi/2.0 - angle))
}

// This library function returns the supplement of the specified angle. The
// supplementary angles add up to π.
func (l *angles) Supplement(angle Angle) Angle {
	return AngleFromFloat(float64(math.Pi - angle))
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
	case math.Pi * 0.25:
		return 0.5 * math.Sqrt2
	case math.Pi * 0.5:
		return 0.0
	case math.Pi * 0.75:
		return -0.5 * math.Sqrt2
	case math.Pi:
		return -1.0
	case math.Pi * 1.25:
		return -0.5 * math.Sqrt2
	case math.Pi * 1.5:
		return 0.0
	case math.Pi * 1.75:
		return 0.5 * math.Sqrt2
	default:
		return math.Cos(float64(angle))
	}
}

// This library function returns the angle whose trigonometric cosine is the
// specified distance along the x-axis.
func (l *angles) ArcCosine(x float64) Angle {
	return AngleFromFloat(math.Acos(x))
}

// This library function returns the trigonometric sine of the specified angle.
func (l *angles) Sine(angle Angle) float64 {
	switch angle {
	case 0.0:
		return 0.0
	case math.Pi * 0.25:
		return 0.5 * math.Sqrt2
	case math.Pi * 0.5:
		return 1.0
	case math.Pi * 0.75:
		return 0.5 * math.Sqrt2
	case math.Pi:
		return 0.0
	case math.Pi * 1.25:
		return -0.5 * math.Sqrt2
	case math.Pi * 1.5:
		return -1.0
	case math.Pi * 1.75:
		return -0.5 * math.Sqrt2
	default:
		return math.Sin(float64(angle))
	}
}

// This library function returns the angle whose trigonometric sine is the
// specified distance along the y-axis.
func (l *angles) ArcSine(y float64) Angle {
	return AngleFromFloat(math.Asin(y))
}

// This library function returns the trigonometric tangent of the specified
// angle.
func (l *angles) Tangent(angle Angle) float64 {
	switch angle {
	case 0.0:
		return 0.0
	case math.Pi * 0.25:
		return 1.0
	case math.Pi * 0.5:
		return math.Inf(1)
	case math.Pi * 0.75:
		return -1.0
	case math.Pi:
		return 0.0
	case math.Pi * 1.25:
		return 1.0
	case math.Pi * 1.5:
		return math.Inf(1)
	case math.Pi * 1.75:
		return -1.0
	default:
		return math.Tan(float64(angle))
	}
}

// This library function returns the angle whose trigonometric tangent is the
// specified ratio of the distances along the y-axis and x-axis.
func (l *angles) ArcTangent(x, y float64) Angle {
	return AngleFromFloat(math.Atan2(y, x))
}

// PRIVATE FUNCTIONS

// This function parses an angle string and returns the corresponding floating
// point number and whether or not the string contained a valid angle.
func stringToAngle(v string) (float64, bool) {
	var angle float64
	var ok = true
	var matches = abstractions.ScanAngle([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	case matches[1] == "pi" || matches[1] == "π":
		angle = math.Pi
	default:
		var err error
		angle, err = strconv.ParseFloat(matches[1], 64)
		if err != nil {
			ok = false
		}
	}
	return angle, ok
}

// This function uses the single precision floating point range to lock a double
// precision phase angle onto 0, π/2, π, or 3π/2 if the angle falls outside the
// single precision range for these values. Otherwise, the phase angle is
// returned unchanged.
func lockPhase(v float64) float64 {
	var v32 float32 = float32(v)
	switch {
	case math.Abs(v) <= 1.2246467991473515E-16:
		return 0
	case v32 == float32(0.5 * math.Pi):
		return 0.5 * math.Pi
	case v32 == float32(math.Pi):
		return math.Pi
	case v32 == float32(1.5 * math.Pi):
		return 1.5 * math.Pi
	default:
		return v
	}
}
