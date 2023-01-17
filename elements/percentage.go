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
	mat "math"
)

// PERCENTAGE INTERFACE

// This constructor returns the minimum value for a percentage.
func MinimumPercentage() Percentage {
	return Percentage(0)
}

// This constructor returns the maximum value for a percentage.
func MaximumPercentage() Percentage {
	return Percentage(100)
}

// This type defines the methods associated with percentage elements. It extends
// the native Go float64 type and represents a percentage. Percentages can be
// negative.
type Percentage float64

// CONTINUOUS INTERFACE

// This method determines whether or not this continuous element is zero.
func (v Percentage) IsZero() bool {
	return v == 0
}

// This method determines whether or not this continuous element is infinite.
func (v Percentage) IsInfinite() bool {
	return mat.IsInf(float64(v), 0)
}

// This method determines whether or not this continuous element is undefined.
func (v Percentage) IsUndefined() bool {
	return mat.IsNaN(float64(v))
}

// This method returns a real value for this continuous element.
func (v Percentage) AsReal() float64 {
	return float64(v / 100.0)
}

// POLARIZED INTERFACE

// This method determines whether or not this polarized component is negative.
func (v Percentage) IsNegative() bool {
	return v < 0.0
}

// PERCENTAGES LIBRARY

// This singleton creates a unique name space for the library functions for
// percentage elements.
var Percentages = &percentages{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for percentage elements.
type percentages struct{}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified percentage.
func (l *percentages) Inverse(percentage Percentage) Percentage {
	return Percentage(float64(-percentage))
}

// This library function returns the sum of the specified percentages.
func (l *percentages) Sum(first, second Percentage) Percentage {
	return Percentage(float64(first + second))
}

// This library function returns the difference of the specified percentages.
func (l *percentages) Difference(first, second Percentage) Percentage {
	return Percentage(float64(first - second))
}

// This library function returns the specified percentage scaled by the
// specified factor.
func (l *percentages) Scaled(percentage Percentage, factor float64) Percentage {
	return Percentage(float64(percentage) * factor)
}
