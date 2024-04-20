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
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	mat "math"
	stc "strconv"
)

// CLASS DEFINITIONS

// This private type implements the AngleLike interface.  It extends the native
// Go `float64` type and represents a radian based angle in the range [0..2π).
type angle_ float64

// This private type defines the structure associated with the class constants
// and class functions for the angle elements.
type angleClass_ struct {
	zero AngleLike
	pi   AngleLike
	tau  AngleLike
}

// CLASS CONSTANTS

// This class constant represents the minimum value for an angle.
func (c *angleClass_) MinimumValue() AngleLike {
	return c.zero
}

// This class constant represents the maximum value for an angle.
func (c *angleClass_) MaximumValue() AngleLike {
	return c.tau
}

// This class constant represents an angle of zero.
func (c *angleClass_) Zero() AngleLike {
	return c.zero
}

// This class constant represents an angle of pi.
func (c *angleClass_) Pi() AngleLike {
	return c.pi
}

// This class constant represents an angle of tau.
func (c *angleClass_) Tau() AngleLike {
	return c.tau
}

// CLASS CONSTRUCTORS

// This constructor creates a new angle from the specified float value and
// normalizes the value to be in the allowed range for angleClass [0..τ).
func (c *angleClass_) FromFloat(float float64) AngleLike {
	if float < -Tau || float >= Tau {
		// Normalize the angle to the range [-τ..τ).
		float = mat.Remainder(float, Tau)
	}
	if float < 0.0 {
		// Normalize the angle to the range [0..τ).
		float = float + Tau
	}
	var angle = angle_(lockPhase(float))
	return angle
}

// This constructor creates a new angle from the specified string value and
// normalizes the value to be in the allowed range for angleClass [0..τ).
func (c *angleClass_) FromString(string_ string) AngleLike {
	var matches = uti.AngleMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct an angle from an invalid string: %v", string_)
		panic(message)
	}
	var float = floatFromString(matches[1]) // Strip off the leading '~' character.
	var angle = c.FromFloat(float)
	return angle
}

// CLASS METHODS

// Continuous Interface

// This instance method returns a real value for this continuous element.
func (v angle_) AsFloat() float64 {
	return float64(v)
}

// This instance method determines whether or not this continuous element is
// zero.
func (v angle_) IsZero() bool {
	return v == Angle().Zero() || v == Angle().Tau()
}

// This instance method determines whether or not this continuous element is
// infinite.
func (v angle_) IsInfinite() bool {
	return false
}

// This instance method determines whether or not this continuous element is
// undefined.
func (v angle_) IsUndefined() bool {
	return false
}

// Lexical Interface

// This instance method returns a string value for this lexical element.
func (v angle_) AsString() string {
	var string_ string
	switch v {
	case Angle().Pi():
		string_ = "~π"
	case Angle().Tau():
		string_ = "~τ"
	default:
		string_ = "~" + stc.FormatFloat(float64(v), 'G', -1, 64)
	}
	return string_
}

// CLASS FUNCTIONS

// This class method returns the inverse of the specified angle.
func (c *angleClass_) Inverse(angle AngleLike) AngleLike {
	return c.FromFloat(angle.AsFloat() - mat.Pi)
}

// This class method returns the sum of the specified angleClass.
func (c *angleClass_) Sum(first, second AngleLike) AngleLike {
	return c.FromFloat(first.AsFloat() + second.AsFloat())
}

// This class method returns the difference of the specified angleClass.
func (c *angleClass_) Difference(first, second AngleLike) AngleLike {
	return c.FromFloat(first.AsFloat() - second.AsFloat())
}

// This class method returns the specified angle scaled by the specified
// factor.
func (c *angleClass_) Scaled(angle AngleLike, factor float64) AngleLike {
	return c.FromFloat(angle.AsFloat() * factor)
}

// This class method returns the complement of the specified angle. The
// complementary angleClass add up to π/2.
func (c *angleClass_) Complement(angle AngleLike) AngleLike {
	return c.FromFloat(mat.Pi/2.0 - angle.AsFloat())
}

// This class method returns the supplement of the specified angle. The
// supplementary angleClass add up to π.
func (c *angleClass_) Supplement(angle AngleLike) AngleLike {
	return c.FromFloat(mat.Pi - angle.AsFloat())
}

// This class method returns the conjugate of the specified angle. The
// conjugate angleClass add up to 2π (zero).
func (c *angleClass_) Conjugate(angle AngleLike) AngleLike {
	return c.FromFloat(-angle.AsFloat())
}

// This class method returns the trigonometric cosine of the specified
// angle.
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
	case Tau:
		cosine = 1.0
	default:
		cosine = mat.Cos(angle.AsFloat())
	}
	return cosine
}

// This class method returns the angle whose trigonometric cosine is the
// specified distance along the x-axis.
func (c *angleClass_) ArcCosine(x float64) AngleLike {
	return c.FromFloat(mat.Acos(x))
}

// This class method returns the trigonometric sine of the specified angle.
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
	case Tau:
		sine = 0.0
	default:
		sine = mat.Sin(angle.AsFloat())
	}
	return sine
}

// This class method returns the angle whose trigonometric sine is the
// specified distance along the y-axis.
func (c *angleClass_) ArcSine(y float64) AngleLike {
	return c.FromFloat(mat.Asin(y))
}

// This class method returns the trigonometric tangent of the specified
// angle.
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
	case Tau:
		tangent = 0.0
	default:
		tangent = mat.Tan(angle.AsFloat())
	}
	return tangent
}

// This class method returns the angle whose trigonometric tangent is the
// specified ratio of the distances along the y-axis and x-axis.
func (c *angleClass_) ArcTangent(x, y float64) AngleLike {
	return c.FromFloat(mat.Atan2(y, x))
}
