/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologiesâ„˘.  All Rights Reserved.   *
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

// This interface defines the methods supported by all numeric elements.
type Discrete interface {
	AsBoolean() bool
	AsInteger() int
}

// This interface defines the methods supported by all numeric elements.
type Continuous interface {
	IsZero() bool
	IsInfinite() bool
	IsUndefined() bool
	AsReal() float64
}

// This interface defines the methods supported by all complex elements.
type Complex interface {
	AsRectangular() string
	AsPolar() string
	GetReal() float64
	GetImaginary() float64
	GetMagnitude() float64
	GetPhase() float64
}

// This interface defines the methods supported by all polarized elements.
type Polarized interface {
	IsNegative() bool
}

// This interface defines the methods supported by all matchable pattern
// elements.
type Matchable interface {
	MatchesText(text string) bool
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
