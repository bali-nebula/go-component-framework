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
	"github.com/craterdog-bali/go-bali-document-notation/utilities"
	"math"
	"strconv"
)

// PROBABILITY INTERFACE

// This constructor creates a new probability and constrains the value to be in
// the allowed range for probabilities [0..1].
func ProbabilityFromFloat(v float64) Probability {
	if v < 0.0 {
		v = 0.0
	} else if v > 1.0 {
		v = 1.0
	}
	return Probability(v)
}

// This constructor creates a new probability from the specified boolean value.
func ProbabilityFromBoolean(v bool) Probability {
	if v {
		return Probability(1)
	}
	return Probability(0)
}

// This constructor attempts to create a new probability from the specified
// formatted string. It returns a probability value and whether or not the
// string contained a valid probability.
// For valid string formats for this type see `../abstractions/language.go`.
func ProbabilityFromString(v string) (Probability, bool) {
	var probability, ok = stringToProbability(v)
	return ProbabilityFromFloat(probability), ok // Will be normalized.
}

// This type defines the methods associated with probability elements. It
// extends the native Go float64 type and represents a probability in the range
// [0..1].
type Probability float64

// LEXICAL INTERFACE

// This method returns the canonical string for this element.
func (v Probability) AsString() string {
	var s = strconv.FormatFloat(float64(v), 'f', -1, 64)
	switch {
	case v == 0:
		return "." + s
	case v == 1:
		return s + "."
	default:
		return s[1:]
	}
}

// This method implements the standard Go Stringer interface.
func (v Probability) String() string {
	return v.AsString()
}

// DISCRETE INTERFACE

// This method returns a boolean value for this discrete element.
func (v Probability) AsBoolean() bool {
	return v.AsInteger() > 0
}

// This method returns an integer value for this discrete element.
func (v Probability) AsInteger() int {
	return int(math.Round(float64(v)))
}

// CONTINUOUS INTERFACE

// This method returns a real value for this continuous component.
func (v Probability) AsReal() float64 {
	return float64(v)
}

// PROBABILITIES LIBRARY

// This singleton creates a unique name space for the library functions for
// probability elements.
var Probabilities = &probabilities{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for probability elements.
type probabilities struct{}

// This library function returns a random probability.
func (l *probabilities) Random() Probability {
	return Probability(utilities.RandomProbability())
}

// LOGICAL INTERFACE

// This library function returns the logical inverse of the specified
// probability.
func (l *probabilities) Not(probability Probability) Probability {
	return Probability(1.0 - probability)
}

// This library function returns the logical conjunction of the specified
// probability elements.
func (l *probabilities) And(first, second Probability) Probability {
	return Probability(first * second)
}

// This library function returns the logical material non-implication of the
// specified probability elements.
func (l *probabilities) Sans(first, second Probability) Probability {
	return l.And(first, l.Not(second))
}

// This library function returns the logical disjunction of the specified
// probability elements.
func (l *probabilities) Or(first Probability, second Probability) Probability {
	return l.Not(l.And(l.Not(first), l.Not(second)))
}

// This library function returns the logical exclusive disjunction of the
// specified probability elements.
func (l *probabilities) Xor(first Probability, second Probability) Probability {
	return l.Or(l.Sans(first, second), l.Sans(second, first))
}

// PRIVATE FUNCTIONS

// This function parses a probability string and returns the corresponding
// floating point number and whether or not the string contained a valid
// probability.
func stringToProbability(v string) (float64, bool) {
	var probability float64
	var ok = true
	var matches = abstractions.ScanProbability([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		var err error
		probability, err = strconv.ParseFloat(matches[1], 64)
		if err != nil {
			ok = false
		}
	}
	// Check for mistaken floating point value.
	matches = abstractions.ScanNumber([]byte(v))
	if len(matches) > 0 && len(matches[1]) > 1 {
		ok = false
		probability = 0
	}
	return probability, ok
}
