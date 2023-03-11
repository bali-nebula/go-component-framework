/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// TYPE DEFINITIONS

type (
	Element any
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all spectrum endpoints.
type Lexical interface {
	AsString() string
}

// This interface defines the methods supported by all interval endpoints.
type Discrete interface {
	AsBoolean() bool
	AsInteger() int
}

// This interface defines the methods supported by all continuum endpoints.
type Continuous interface {
	AsReal() float64
	IsZero() bool
	IsInfinite() bool
	IsUndefined() bool
}

// This interface defines the methods supported by all polarized elements.
type Polarized interface {
	IsNegative() bool
}

// This interface defines the methods supported by all complex elements.
type Complex interface {
	AsComplex() complex128
	GetReal() float64
	GetImaginary() float64
	GetMagnitude() float64
	GetPhase() AngleLike
}

// This interface defines the methods supported by all matchable pattern
// elements.
type Matchable interface {
	MatchesText(text string) bool
	GetMatches(text string) []string
}

// This interface defines the methods supported by all segmented resource
// elements.
type Segmented interface {
	GetScheme() string
	GetAuthority() string
	GetPath() string
	GetQuery() string
	GetFragment() string
}

// This interface defines the methods supported by all temporal elements.
type Temporal interface {
	// Return the entire time in specific units.
	AsMilliseconds() float64
	AsSeconds() float64
	AsMinutes() float64
	AsHours() float64
	AsDays() float64
	AsWeeks() float64
	AsMonths() float64
	AsYears() float64

	// Return a specific part of the entire time.
	GetMilliseconds() int
	GetSeconds() int
	GetMinutes() int
	GetHours() int
	GetDays() int
	GetWeeks() int
	GetMonths() int
	GetYears() int
}

// CONSOLIDATED INTERFACES

type AngleLike interface {
	Continuous
}

type BooleanLike interface {
	Discrete
}

type DurationLike interface {
	Discrete
	Polarized
	Temporal
}

type MomentLike interface {
	Discrete
	Temporal
}

type NumberLike interface {
	Continuous
	Polarized
	Complex
}

type PatternLike interface {
	Lexical
	Matchable
}

type PercentageLike interface {
	Discrete
	Continuous
	Polarized
}

type ProbabilityLike interface {
	Discrete
	Continuous
}

type ResourceLike interface {
	Lexical
	Segmented
}

type SymbolLike interface {
	Lexical
}
