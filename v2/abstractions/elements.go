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

type Complex interface {
	AsComplex() complex128
	GetReal() float64
	GetImaginary() float64
	GetMagnitude() float64
	GetPhase() float64
}

type Continuous interface {
	AsReal() float64
	IsZero() bool
	IsInfinite() bool
	IsUndefined() bool
}

type Discrete interface {
	AsBoolean() bool
	AsInteger() int
}

type Lexical interface {
	AsString() string
}

type Matchable interface {
	MatchesText(text string) bool
	GetMatches(text string) []string
}

type Named interface {
	GetName() string
}

type Polarized interface {
	IsNegative() bool
}

type Segmented interface {
	GetScheme() string
	GetAuthority() string
	GetPath() string
	GetQuery() string
	GetFragment() string
}

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

type Versioned interface {
	GetVersion() string
}

// CONSOLIDATED INTERFACES

type AngleLike interface {
	Continuous
	Lexical
}

type BooleanLike interface {
	Discrete
	Lexical
}

type CharacterLike interface {
	Discrete
	Lexical
}

type CitationLike interface {
	Lexical
	Named
	Versioned
}

type DurationLike interface {
	Discrete
	Lexical
	Polarized
	Temporal
}

type IntegerLike interface {
	Discrete
	Lexical
	Polarized
}

type MomentLike interface {
	Discrete
	Lexical
	Temporal
}

type NumberLike interface {
	Complex
	Continuous
	Lexical
	Polarized
}

type PatternLike interface {
	Lexical
	Matchable
}

type PercentageLike interface {
	Continuous
	Discrete
	Lexical
	Polarized
}

type ProbabilityLike interface {
	Continuous
	Discrete
	Lexical
}

type RealLike interface {
	Continuous
	Lexical
	Polarized
}

type ResourceLike interface {
	Lexical
	Segmented
}
