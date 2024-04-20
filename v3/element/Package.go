/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

/*
Package "element" provides the elemental types that are found in Bali Document
Notation™ (Bali). All element types are immutable and—for better
performance—are implemented as extensions to existing Go primitive types.

This package follows the Crater Dog Technologies™ Go Coding Conventions located
here:
  - https://github.com/craterdog/go-model-framework/wiki

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-component-framework/wiki/Elements

Additional implementations of the concrete classes provided by this package can
be developed and used seamlessly since the interface definitions only depend on
other interfaces and primitive types—and the class implementations only depend
on interfaces, not on each other.
*/
package element

// Aspects

/*
Complex is an aspect interface that defines a set of method signatures
that must be supported by each instance of a complex elemental class.
*/
type Complex interface {
	// Methods
	AsComplex() complex128
	GetReal() float64
	GetImaginary() float64
	GetMagnitude() float64
	GetPhase() float64
}

/*
Continuous is an aspect interface that defines a set of method signatures
that must be supported by each instance of a continuous elemental class.
*/
type Continuous interface {
	// Methods
	AsFloat() float64
	IsZero() bool
	IsInfinite() bool
	IsUndefined() bool
}

/*
Discrete is an aspect interface that defines a set of method signatures
that must be supported by each instance of a discrete elemental class.
*/
type Discrete interface {
	// Methods
	AsBoolean() bool
	AsInteger() int
}

/*
Factored is an aspect interface that defines a set of method signatures
that must be supported by each instance of a factored elemental class.
*/
type Factored interface {
	// Methods
	GetMilliseconds() int
	GetSeconds() int
	GetMinutes() int
	GetHours() int
	GetDays() int
	GetWeeks() int
	GetMonths() int
	GetYears() int
}

/*
Lexical is an aspect interface that defines a set of method signatures
that must be supported by each instance of a lexical elemental class.
*/
type Lexical interface {
	// Methods
	AsString() string
}

/*
Matchable is an aspect interface that defines a set of method signatures
that must be supported by each instance of a matchable elemental class.
*/
type Matchable interface {
	// Methods
	MatchesText(text string) bool
	GetMatches(text string) []string
}

/*
Named is an aspect interface that defines a set of method signatures
that must be supported by each instance of a named elemental class.
*/
type Named interface {
	// Methods
	GetName() string
}

/*
Polarized is an aspect interface that defines a set of method signatures
that must be supported by each instance of a polarized elemental class.
*/
type Polarized interface {
	// Methods
	IsNegative() bool
}

/*
Segmented is an aspect interface that defines a set of method signatures
that must be supported by each instance of a segmented elemental class.
*/
type Segmented interface {
	// Methods
	GetScheme() string
	GetAuthority() string
	GetPath() string
	GetQuery() string
	GetFragment() string
}

/*
Temporal is an aspect interface that defines a set of method signatures
that must be supported by each instance of a temporal elemental class.
*/
type Temporal interface {
	// Methods
	AsMilliseconds() float64
	AsSeconds() float64
	AsMinutes() float64
	AsHours() float64
	AsDays() float64
	AsWeeks() float64
	AsMonths() float64
	AsYears() float64
}

/*
Versioned is an aspect interface that defines a set of method signatures
that must be supported by each instance of a versioned elemental class.
*/
type Versioned interface {
	// Methods
	GetVersion() string
}

// Classes

/*
AngleClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
angle-like concrete class.
*/
type AngleClassLike interface {
	// Constants
	MinimumValue() AngleLike
	MaximumValue() AngleLike
	Zero() AngleLike
	Pi() AngleLike
	Tau() AngleLike

	// Constructors
	MakeFromFloat(float float64) AngleLike
	MakeFromString(string_ string) AngleLike

	// Functions
	Inverse(angle AngleLike) AngleLike
	Sum(
		first AngleLike,
		second AngleLike,
	) AngleLike
	Difference(
		first AngleLike,
		second AngleLike,
	) AngleLike
	Scaled(
		angle AngleLike,
		factor float64,
	) AngleLike
	Complement(angle AngleLike) AngleLike
	Supplement(angle AngleLike) AngleLike
	Conjugate(angle AngleLike) AngleLike
	Cosine(angle AngleLike) float64
	ArcCosine(x float64) AngleLike
	Sine(angle AngleLike) float64
	ArcSine(y float64) AngleLike
	Tangent(angle AngleLike) float64
	ArcTangent(
		x float64,
		y float64,
	) AngleLike
}

/*
BooleanClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
boolean-like concrete class.
*/
type BooleanClassLike interface {
	// Constants
	False() BooleanLike
	True() BooleanLike

	// Constructors
	MakeFromBoolean(boolean bool) BooleanLike
	MakeFromString(string_ string) BooleanLike

	// Functions
	Not(boolean BooleanLike) BooleanLike
	And(
		first BooleanLike,
		second BooleanLike,
	) BooleanLike
	Sans(
		first BooleanLike,
		second BooleanLike,
	) BooleanLike
	Or(
		first BooleanLike,
		second BooleanLike,
	) BooleanLike
	Xor(
		first BooleanLike,
		second BooleanLike,
	) BooleanLike
}

/*
CharacterClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
character-like concrete class.
*/
type CharacterClassLike interface {
	// Constants
	MinimumValue() CharacterLike
	MaximumValue() CharacterLike

	// Constructors
	MakeFromRune(rune_ rune) CharacterLike
	MakeFromInteger(integer int) CharacterLike
	MakeFromString(string_ string) CharacterLike

	// Functions
	ToLowercase(character CharacterLike) CharacterLike
	ToUppercase(character CharacterLike) CharacterLike
}

/*
CitationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
citation-like concrete class.
*/
type CitationClassLike interface {
	// Constructors
	MakeFromString(string_ string) CitationLike
}

/*
DurationClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
duration-like concrete class.
*/
type DurationClassLike interface {
	// Constants
	MinimumValue() DurationLike
	MaximumValue() DurationLike
	MillisecondsPerSecond() int
	MillisecondsPerMinute() int
	MillisecondsPerHour() int
	MillisecondsPerDay() int
	MillisecondsPerWeek() int
	MillisecondsPerMonth() int
	MillisecondsPerYear() int
	DaysPerMonth() float64
	DaysPerYear() float64
	WeeksPerMonth() float64

	// Constructors
	MakeFromMilliseconds(milliseconds int) DurationLike
	MakeFromString(string_ string) DurationLike
}

/*
FloatClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
float-like concrete class.
*/
type FloatClassLike interface {
	// Constants
	MinimumValue() FloatLike
	MaximumValue() FloatLike

	// Constructors
	MakeFromFloat(float float64) FloatLike
	MakeFromString(string_ string) FloatLike
}

/*
IntegerClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
integer-like concrete class.
*/
type IntegerClassLike interface {
	// Constants
	MinimumValue() IntegerLike
	MaximumValue() IntegerLike

	// Constructors
	MakeFromInteger(integer int) IntegerLike
	MakeFromString(string_ string) IntegerLike
}

/*
MomentClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
moment-like concrete class.
*/
type MomentClassLike interface {
	// Constants
	MinimumValue() MomentLike
	MaximumValue() MomentLike
	Epoch() MomentLike

	// Constructors
	Make() MomentLike
	MakeFromMilliseconds(milliseconds int) MomentLike
	MakeFromString(string_ string) MomentLike

	// Functions
	Duration(
		first MomentLike,
		second MomentLike,
	) DurationLike
	Earlier(
		moment MomentLike,
		duration DurationLike,
	) MomentLike
	Later(
		moment MomentLike,
		duration DurationLike,
	) MomentLike
}

/*
NumberClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
number-like concrete class.
*/
type NumberClassLike interface {
	// Constants
	MinimumValue() NumberLike
	MaximumValue() NumberLike
	Zero() NumberLike
	One() NumberLike
	I() NumberLike
	E() NumberLike
	Pi() NumberLike
	Phi() NumberLike
	Tau() NumberLike
	Infinity() NumberLike
	Undefined() NumberLike

	// Constructors
	MakeFromComplex(complex_ complex128) NumberLike
	MakeFromPolar(
		magnitude float64,
		phase float64,
	) NumberLike
	MakeFromString(string_ string) NumberLike

	// Functions
	Inverse(number NumberLike) NumberLike
	Sum(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Difference(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Scaled(
		number NumberLike,
		factor float64,
	) NumberLike
	Reciprocal(number NumberLike) NumberLike
	Conjugate(number NumberLike) NumberLike
	Product(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Quotient(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Remainder(
		first NumberLike,
		second NumberLike,
	) NumberLike
	Power(
		base NumberLike,
		exponent NumberLike,
	) NumberLike
	Logarithm(
		base NumberLike,
		number NumberLike,
	) NumberLike
}

/*
PatternClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
pattern-like concrete class.
*/
type PatternClassLike interface {
	// Constants
	None() PatternLike
	Any() PatternLike

	// Constructors
	MakeFromString(string_ string) PatternLike
}

/*
PercentageClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
percentage-like concrete class.
*/
type PercentageClassLike interface {
	// Constructors
	MakeFromInteger(integer int) PercentageLike
	MakeFromFloat(float float64) PercentageLike
	MakeFromString(string_ string) PercentageLike
}

/*
ProbabilityClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
probability-like concrete class.
*/
type ProbabilityClassLike interface {
	// Constants
	MinimumValue() ProbabilityLike
	MaximumValue() ProbabilityLike

	// Constructors
	MakeFromFloat(float float64) ProbabilityLike
	MakeFromBoolean(boolean bool) ProbabilityLike
	MakeFromString(string_ string) ProbabilityLike

	// Functions
	Random() ProbabilityLike
}

/*
ResourceClassLike is a class interface that defines the complete set of
class constants, constructors and functions that must be supported by each
resource-like concrete class.
*/
type ResourceClassLike interface {
	// Constructors
	MakeFromString(string_ string) ResourceLike
}

// Instances

/*
AngleLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a angle-like elemental class.
*/
type AngleLike interface {
	// Abstractions
	Continuous
	Lexical
}

/*
BooleanLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a boolean-like elemental class.
*/
type BooleanLike interface {
	// Abstractions
	Discrete
	Lexical
}

/*
CharacterLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a character-like elemental class.
*/
type CharacterLike interface {
	// Abstractions
	Discrete
	Lexical
}

/*
CitationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a citation-like elemental class.
*/
type CitationLike interface {
	// Abstractions
	Lexical
	Named
	Versioned
}

/*
DurationLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a duration-like elemental class.
*/
type DurationLike interface {
	// Abstractions
	Discrete
	Lexical
	Polarized
	Temporal
	Factored
}

/*
FloatLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a float-like elemental class.
*/
type FloatLike interface {
	// Abstractions
	Continuous
	Lexical
	Polarized
}

/*
IntegerLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a integer-like elemental class.
*/
type IntegerLike interface {
	// Abstractions
	Discrete
	Lexical
	Polarized
}

/*
MomentLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a moment-like elemental class.
*/
type MomentLike interface {
	// Abstractions
	Discrete
	Lexical
	Temporal
}

/*
NumberLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a number-like elemental class.
*/
type NumberLike interface {
	// Abstractions
	Complex
	Continuous
	Lexical
	Polarized
}

/*
PatternLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a pattern-like elemental class.
*/
type PatternLike interface {
	// Abstractions
	Lexical
	Matchable
}

/*
PercentageLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a percentage-like elemental class.
*/
type PercentageLike interface {
	// Abstractions
	Continuous
	Discrete
	Lexical
	Polarized
}

/*
ProbabilityLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a probability-like elemental class.
*/
type ProbabilityLike interface {
	// Abstractions
	Continuous
	Discrete
	Lexical
}

/*
ResourceLike is an instance interface that defines the complete set of
instance attributes, abstractions and methods that must be supported by each
instance of a resource-like elemental class.
*/
type ResourceLike interface {
	// Abstractions
	Lexical
	Segmented
}
