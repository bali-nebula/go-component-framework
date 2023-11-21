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
	stc "strconv"
)

// PROBABILITY ELEMENT CONSTRUCTORS

// This constructor creates a new probability element and constrains the value
// to be in the allowed range for probabilities [0..1].
func ProbabilityFromFloat(float float64) abs.ProbabilityLike {
	var probability abs.ProbabilityLike
	switch {
	case float < 0.0:
		probability = probability_(0.0)
	case float > 1.0:
		probability = probability_(1.0)
	default:
		probability = probability_(float)
	}
	return probability
}

// This constructor creates a new probability element from the specified boolean
// value.
func ProbabilityFromBool(boolean bool) abs.ProbabilityLike {
	var probability abs.ProbabilityLike
	switch boolean {
	case true:
		probability = probability_(1)
	case false:
		probability = probability_(0)
	}
	return probability
}

// This constructor creates a new probability element from the specified string.
func ProbabilityFromString(string_ string) abs.ProbabilityLike {
	var matches = uti.ProbabilityMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a probability from an invalid string: %v", string_)
		panic(message)
	}
	var float, _ = stc.ParseFloat(matches[1], 64) // Strip off the '%' suffix.
	var probability = probability_(float)
	return probability
}

// This constructor creates a new random probability.
func RandomProbability() abs.ProbabilityLike {
	var probability = probability_(uti.RandomProbability())
	return probability
}

// This constructor returns the minimum value for a probability.
func MinimumProbability() abs.ProbabilityLike {
	var probability = probability_(0)
	return probability
}

// This constructor returns the maximum value for a probability.
func MaximumProbability() abs.ProbabilityLike {
	var probability = probability_(1)
	return probability
}

// PROBABILITY ELEMENT METHODS

// This private type implements the ProbabilityLike interface.  It extends the
// native Go `float64` type and represents a probability in the range [0..1].
type probability_ float64

// CONTINUOUS INTERFACE

// This method returns a real value for this continuous element.
func (v probability_) AsFloat() float64 {
	return float64(v)
}

// This method determines whether or not this continuous element is zero.
func (v probability_) IsZero() bool {
	return v == 0
}

// This method determines whether or not this continuous element is infinite.
func (v probability_) IsInfinite() bool {
	return false
}

// This method determines whether or not this continuous element is undefined.
func (v probability_) IsUndefined() bool {
	return false
}

// DISCRETE INTERFACE

// This method returns a probability value for this discrete element.
func (v probability_) AsBoolean() bool {
	if v == 0.5 {
		return uti.RandomBoolean()
	}
	return v > 0.5
}

// This method returns an integer value for this discrete element.
func (v probability_) AsInteger() int {
	if v.AsBoolean() {
		return 1
	}
	return 0
}

// LEXICAL INTERFACE

// This method returns a string value for this lexical element.
func (v probability_) AsString() string {
	var float = v.AsFloat()
	var string_ = stc.FormatFloat(float, 'f', -1, 64)
	switch float {
	case 0:
		// Zero is formatted as ".0".
		string_ = "." + string_
	case 1:
		// One is formatted as "1.".
		string_ += "."
	default:
		// Otherwise, remove the leading "0".
		string_ = string_[1:]
	}
	return string_
}

// PROBABILITY ELEMENT LIBRARY

// This singleton creates a unique name space for the library functions for
// probability elements.
var Probability = &probabilities_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for probability elements.
type probabilities_ struct{}

// LOGICAL INTERFACE

// This library function returns the logical inverse of the specified
// probability.
func (l *probabilities_) Not(probability abs.ProbabilityLike) abs.ProbabilityLike {
	return probability_(1.0 - probability.AsFloat())
}

// This library function returns the logical conjunction of the specified
// probability elements.
func (l *probabilities_) And(first, second abs.ProbabilityLike) abs.ProbabilityLike {
	return probability_(first.AsFloat() * second.AsFloat())
}

// This library function returns the logical material non-implication of the
// specified probability elements.
func (l *probabilities_) Sans(first, second abs.ProbabilityLike) abs.ProbabilityLike {
	return probability_(first.AsFloat() * (1.0 - second.AsFloat()))
}

// This library function returns the logical disjunction of the specified
// probability elements.
func (l *probabilities_) Or(first, second abs.ProbabilityLike) abs.ProbabilityLike {
	return probability_(first.AsFloat() + second.AsFloat() - (first.AsFloat() * second.AsFloat()))
}

// This library function returns the logical exclusive disjunction of the
// specified probability elements.
func (l *probabilities_) Xor(first, second abs.ProbabilityLike) abs.ProbabilityLike {
	return probability_(first.AsFloat() + second.AsFloat() - (2.0 * first.AsFloat() * second.AsFloat()))
}
