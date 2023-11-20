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

// PERCENTAGE INTERFACE

// This constructor creates a new percentage element from the specified integer.
func PercentageFromInt(integer int) abs.PercentageLike {
	return percentage_(float64(integer))
}

// This constructor creates a new percentage element from the specified float.
func PercentageFromFloat(float float64) abs.PercentageLike {
	return percentage_(float)
}

// This constructor creates a new percentage element from the specified string.
func PercentageFromString(string_ string) abs.PercentageLike {
	var matches = uti.PercentageMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a percentage from an invalid string: %v", string_)
		panic(message)
	}
	var float, _ = stc.ParseFloat(matches[1], 64) // Strip off the '%' suffix.
	var percentage = percentage_(float)
	return percentage
}

// This constructor returns the minimum value for a percentage.
func MinimumPercentage() abs.PercentageLike {
	return percentage_(0)
}

// This constructor returns the maximum value for a percentage.
func MaximumPercentage() abs.PercentageLike {
	return percentage_(100)
}

// PERCENTAGE IMPLEMENTATION

// This type defines the methods associated with percentage elements. It extends
// the native Go float64 type and represents a percentage. Percentages can be
// negative.
type percentage_ float64

// CONTINUOUS INTERFACE

// This method returns a real value for this continuous element.
func (v percentage_) AsReal() float64 {
	return float64(v / 100.0)
}

// This method determines whether or not this continuous element is zero.
func (v percentage_) IsZero() bool {
	return v == 0
}

// This method determines whether or not this continuous element is infinite.
func (v percentage_) IsInfinite() bool {
	return mat.IsInf(float64(v), 0)
}

// This method determines whether or not this continuous element is undefined.
func (v percentage_) IsUndefined() bool {
	return mat.IsNaN(float64(v))
}

// DISCRETE INTERFACE

// This method returns a boolean value for this discrete element.
func (v percentage_) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this discrete element.
func (v percentage_) AsInteger() int {
	return int(float64(v))
}

// LEXICAL INTERFACE

// This method returns a string value for this lexical element.
func (v percentage_) AsString() string {
	var string_ = stringFromReal(float64(v)) + "%"
	return string_
}

// POLARIZED INTERFACE

// This method determines whether or not this polarized component is negative.
func (v percentage_) IsNegative() bool {
	return v < 0.0
}
