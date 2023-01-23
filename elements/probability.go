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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	uti "github.com/bali-nebula/go-component-framework/utilities"
)

// PROBABILITY IMPLEMENTATION

// This constructor creates a new probability and constrains the value to be in
// the allowed range for probabilities [0..1].
func ProbabilityFromFloat(v float64) abs.ProbabilityLike {
	if v < 0.0 {
		v = 0.0
	} else if v > 1.0 {
		v = 1.0
	}
	return Probability(v)
}

// This constructor creates a new probability from the specified boolean value.
func ProbabilityFromBoolean(v bool) abs.ProbabilityLike {
	if v {
		return Probability(1)
	}
	return Probability(0)
}

// This constructor returns the minimum value for a probability.
func MinimumProbability() abs.ProbabilityLike {
	return Probability(0)
}

// This constructor returns the maximum value for a probability.
func MaximumProbability() abs.ProbabilityLike {
	return Probability(1)
}

// This type defines the methods associated with probability elements. It
// extends the native Go float64 type and represents a probability in the range
// [0..1].
type Probability float64

// CONTINUOUS INTERFACE

// This method returns a real value for this continuous element.
func (v Probability) AsReal() float64 {
	return float64(v)
}

// This method determines whether or not this continuous element is zero.
func (v Probability) IsZero() bool {
	return v == 0
}

// This method determines whether or not this continuous element is infinite.
func (v Probability) IsInfinite() bool {
	return false
}

// This method determines whether or not this continuous element is undefined.
func (v Probability) IsUndefined() bool {
	return false
}

// PROBABILITIES LIBRARY

// This singleton creates a unique name space for the library functions for
// probability elements.
var Probabilities = &probabilities{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for probability elements.
type probabilities struct{}

// This library function returns a random probability.
func (l *probabilities) Random() abs.ProbabilityLike {
	return Probability(uti.RandomProbability())
}

// LOGICAL INTERFACE

// This library function returns the logical inverse of the specified
// probability.
func (l *probabilities) Not(probability abs.ProbabilityLike) abs.ProbabilityLike {
	return Probability(1.0 - probability.AsReal())
}

// This library function returns the logical conjunction of the specified
// probability elements.
func (l *probabilities) And(first, second abs.ProbabilityLike) abs.ProbabilityLike {
	return Probability(first.AsReal() * second.AsReal())
}

// This library function returns the logical material non-implication of the
// specified probability elements.
func (l *probabilities) Sans(first, second abs.ProbabilityLike) abs.ProbabilityLike {
	return Probability(first.AsReal() * (1.0 - second.AsReal()))
}

// This library function returns the logical disjunction of the specified
// probability elements.
func (l *probabilities) Or(first, second abs.ProbabilityLike) abs.ProbabilityLike {
	return Probability(first.AsReal() + second.AsReal() - (first.AsReal() * second.AsReal()))
}

// This library function returns the logical exclusive disjunction of the
// specified probability elements.
func (l *probabilities) Xor(first, second abs.ProbabilityLike) abs.ProbabilityLike {
	return Probability(first.AsReal() + second.AsReal() - (2.0 * first.AsReal() * second.AsReal()))
}
