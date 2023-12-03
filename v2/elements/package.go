/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

/*
This package defines the elemental types that are found in Bali Document
Notation™ (Bali). All element types are immutable and—for better
performance—are implemented as extensions to existing Go primitive types.

For detailed documentation on this package refer to the wiki:
  - https://github.com/bali-nebula/go-component-framework/wiki/Elements

All UpperCase type, constant, abstraction, function and class names are visible
to other projects that import this package. All lowercase names are private to
this package.
*/
package elements

import (
	fmt "fmt"
	mat "math"
	cmp "math/cmplx"
	stc "strconv"
	sts "strings"
	tim "time"
)

// PACKAGE TYPES

// Specialized Types

// This specialized type definition represents a specialization of the primitive
// Go any data type.  It is used as a generic type for any of the element types
// defined in this package.
type (
	Element any
)

// PACKAGE CONSTANTS

// Private Constants

// This private constant represents the value tau (τ).
// See "The Tau Manifesto" at https://tauday.com/tau-manifesto
const tau = 2.0 * mat.Pi

// These private constants represent the ratios of various units of time.
const (
	// These are locked to the Earth's daily revolutions.
	millisecondsPerSecond int = 1000
	millisecondsPerMinute int = millisecondsPerSecond * 60
	millisecondsPerHour   int = millisecondsPerMinute * 60
	millisecondsPerDay    int = millisecondsPerHour * 24
	millisecondsPerWeek   int = millisecondsPerDay * 7

	// These are locked to the Earth's yearly orbit around the sun.
	millisecondsPerYear  int = 31556952000
	millisecondsPerMonth int = millisecondsPerYear / 12 // An average but exact value.

	// Tying the two together is where things get messy.
	daysPerMonth  float64 = float64(millisecondsPerMonth) / float64(millisecondsPerDay)  // ~30.436875 days/month
	daysPerYear   float64 = float64(millisecondsPerYear) / float64(millisecondsPerDay)   // ~365.2425 days/year
	weeksPerMonth float64 = float64(millisecondsPerMonth) / float64(millisecondsPerWeek) // ~4.348125 weeks/month
)

// PACKAGE ABSTRACTIONS

// Abstract Interfaces

// This abstract interface defines the set of method signatures that must be
// supported by all complex numeric types.
type Complex interface {
	AsComplex() complex128
	GetReal() float64
	GetImaginary() float64
	GetMagnitude() float64
	GetPhase() float64
}

// This abstract interface defines the set of method signatures that must be
// supported by all continuous numeric types.
type Continuous interface {
	AsFloat() float64
	IsZero() bool
	IsInfinite() bool
	IsUndefined() bool
}

// This abstract interface defines the set of method signatures that must be
// supported by all discrete numeric types.
type Discrete interface {
	AsBoolean() bool
	AsInteger() int
}

// This abstract interface defines the set of method signatures that must be
// supported by all lexical string types.
type Lexical interface {
	AsString() string
}

// This abstract interface defines the set of method signatures that must be
// supported by all matchable pattern types.
type Matchable interface {
	MatchesText(text string) bool
	GetMatches(text string) []string
}

// This abstract interface defines the set of method signatures that must be
// supported by all named identifier types.
type Named interface {
	GetName() string
}

// This abstract interface defines the set of method signatures that must be
// supported by all polarized numeric types.
type Polarized interface {
	IsNegative() bool
}

// This abstract interface defines the set of method signatures that must be
// supported by all segmented string types.
type Segmented interface {
	GetScheme() string
	GetAuthority() string
	GetPath() string
	GetQuery() string
	GetFragment() string
}

// This abstract interface defines the set of method signatures that must be
// supported by all temporal duration types.
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

// This abstract interface defines the set of method signatures that must be
// supported by all versioned identifier types.
type Versioned interface {
	GetVersion() string
}

// Abstract Types

// This abstract type defines the set of abstract interfaces that must be
// supported by all angle-like types.
type AngleLike interface {
	Continuous
	Lexical
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all boolean-like types.
type BooleanLike interface {
	Discrete
	Lexical
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all character-like types.
type CharacterLike interface {
	Discrete
	Lexical
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all citation-like types.
type CitationLike interface {
	Lexical
	Named
	Versioned
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all duration-like types.
type DurationLike interface {
	Discrete
	Lexical
	Polarized
	Temporal
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all float-like types.
type FloatLike interface {
	Continuous
	Lexical
	Polarized
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all integer-like types.
type IntegerLike interface {
	Discrete
	Lexical
	Polarized
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all moment-like types.
type MomentLike interface {
	Discrete
	Lexical
	Temporal
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all number-like types.
type NumberLike interface {
	Complex
	Continuous
	Lexical
	Polarized
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all pattern-like types.
type PatternLike interface {
	Lexical
	Matchable
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all percentage-like types.
type PercentageLike interface {
	Continuous
	Discrete
	Lexical
	Polarized
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all probability-like types.
type ProbabilityLike interface {
	Continuous
	Discrete
	Lexical
}

// This abstract type defines the set of abstract interfaces that must be
// supported by all resource-like types.
type ResourceLike interface {
	Lexical
	Segmented
}

// PACKAGE FUNCTIONS

// Private Functions

// This private function returns the complex number associated with the
// specified regular expression matches.
func complexFromMatches(matches []string) complex128 {
	var complex_ complex128
	switch {
	case len(matches[2]) > 0:
		// This is a complex number in rectangular form.
		var realPart = floatFromString(matches[1])
		var imaginaryPart = floatFromString(matches[3][:len(matches[3])-1])
		complex_ = complex(realPart, imaginaryPart)
	case len(matches[5]) > 0:
		// This is a complex number in polar form.
		var magnitude = floatFromString(matches[4])
		var phase = floatFromString(matches[7])
		complex_ = cmp.Rect(magnitude, phase)
	default:
		// This is a pure (non-complex) number.
		switch matches[0] {
		case "undefined":
			complex_ = cmp.NaN()
		case "infinity", "∞":
			complex_ = cmp.Inf()
		case "0":
			complex_ = complex(0, 0)
		case "+i", "i":
			complex_ = complex(0, 1)
		case "-i":
			complex_ = complex(0, -1)
		case "+pi", "pi", "-pi", "+phi", "phi", "-phi":
			// We must handle the constants that end in "i" separately.
			complex_ = complex(floatFromString(matches[1]), 0)
		default:
			if sts.HasSuffix(matches[0], "i") {
				// This is a pure imaginary number.
				complex_ = complex(0, floatFromString(matches[0][:len(matches[0])-1]))
			} else {
				// This is a pure real number.
				complex_ = complex(floatFromString(matches[0]), 0)
			}
		}
	}
	return complex_
}

// This private function returns the floating point value for the specified
// string.
func floatFromString(string_ string) float64 {
	var float float64
	switch string_ {
	case "+e", "e":
		float = mat.E
	case "-e":
		float = -mat.E
	case "+pi", "+π", "pi", "π":
		float = mat.Pi
	case "-pi", "-π":
		float = -mat.Pi
	case "+phi", "+φ", "phi", "φ":
		float = mat.Phi
	case "-phi", "-φ":
		float = -mat.Phi
	case "+tau", "+τ", "tau", "τ":
		float = mat.Pi * 2.0
	case "-tau", "-τ":
		float = -mat.Pi * 2.0
	case "+infinity", "+∞", "infinity", "∞":
		float = mat.Inf(1)
	case "-infinity", "-∞":
		float = mat.Inf(-1)
	default:
		float, _ = stc.ParseFloat(string_, 64)
	}
	return float
}

// This private function formats the specified ordinal value to the specified
// number of digits.
func formatOrdinal(ordinal, digits int) string {
	return fmt.Sprintf("%0"+stc.Itoa(digits)+"d", ordinal)
}

// This list contains the supported ISO 8601 date-time formats delimited by
// angle brackets. Note: the Go templates in this list must contain their exact
// numeric values. If you are curious why this is, check out this posting:
//   - https://medium.com/@simplyianm/how-go-solves-date-and-time-formatting-8a932117c41c
//
// A good mnemonic to use to remember the pattern is:
//
//	  1    2    3     4      5     6      7
//	month day hour minute second year timezone
//
//	"January 2nd, 2006 at 3:04:05pm in the MST timezone"
//
// Anyway, it is what it is, but this hides it from the rest of the code.
var hackedIsoFormats = [...]string{
	"<2006-01-02T15:04:05.000>",
	"<2006-01-02T15:04:05>",
	"<2006-01-02T15:04>",
	"<2006-01-02T15>",
	"<2006-01-02>",
	"<2006-01>",
	"<2006>",
	"<-2006>",
	"<-2006-01>",
	"<-2006-01-02>",
	"<-2006-01-02T15>",
	"<-2006-01-02T15:04>",
	"<-2006-01-02T15:04:05>",
	"<-2006-01-02T15:04:05.000>",
}

// The Go time.Parse() function can only handle years in the range [0000..9999]
// even though the time.Time.Format() method will correctly print any size year
// positive or negative. The Go team has labeled this issue as "unfortunate" and
// will not fix it.
//
// To support the entire ISO 8601 calendar (and beyond) as shown here:
//
//	https://en.wikipedia.org/wiki/Holocene_calendar#Conversion
//
// we must resort to some hacking with this private function...
func hackedParseDateAsMilliseconds(matches []string) int {

	// First, we replace the year with year zero.
	var yearString = matches[2]
	var patched = sts.Replace(matches[0], yearString, "0000", 1)

	// Next, we attempt to parse the patched moment using our Go based formats.
	for _, format := range hackedIsoFormats {
		var date, err = tim.Parse(format, patched) // Parsed in UTC.
		if err == nil {

			// We found a match, now we add back in the correct year.
			var year, _ = stc.ParseInt(yearString, 10, 64)
			date = date.AddDate(int(year), 0, 0)
			if sts.HasPrefix(format, "<-") {

				// We change the positive date to a negative one.
				date = date.AddDate(-2*date.Year(), 0, 0)
			}

			// And return the correct date as milliseconds.
			var milliseconds = int(date.UnixMilli())
			return milliseconds
		}
	}

	// This panic will only happen if the scanner regular expressions are out
	// of sync with the ISO 8601 standard formats. The moment has already been
	// succussfully scanned by the the scanner.
	panic(fmt.Sprintf("The moment does not match a known format: %v", matches[0]))
}

// This private function uses the single precision floating point range to lock
// a double precision magnitude onto 0, 1, -1, or ∞ if the magnitude falls
// outside the single precision range for these values. Otherwise, the magnitude
// is returned unchanged.
func lockMagnitude(magnitude float64) float64 {
	var magnitude32 = float32(magnitude)
	switch {
	case mat.Abs(magnitude) <= 1.2246467991473515e-16:
		magnitude = 0
	case magnitude32 == -1:
		magnitude = -1
	case magnitude32 == 1:
		magnitude = 1
	case mat.IsInf(magnitude, 0):
		magnitude = mat.Inf(1)
	}
	return magnitude
}

// This private function uses the single precision floating point range to lock
// a double precision phase onto 0, π/2, π, or 3π/2 if the phase falls outside
// the single precision range for these values. Otherwise, the phase is returned
// unchanged.
func lockPhase(phase float64) float64 {
	var phase32 = float32(phase)
	switch {
	case mat.Abs(phase) <= 1.2246467991473515e-16:
		phase = 0
	case phase32 == float32(0.5*mat.Pi):
		phase = 0.5 * mat.Pi
	case phase32 == float32(mat.Pi):
		phase = mat.Pi
	case phase32 == float32(1.5*mat.Pi):
		phase = 1.5 * mat.Pi
	}
	return phase
}

// This private function returns the magnitude (absolute value) of the specified value.
func magnitude(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// This private function returns the milliseconds value from the specified matches.
func millisecondsFromMatches(matches []string) int {
	var milliseconds = 0.0
	var sign = 1.0
	var isTime = false
	for _, match := range matches[1:] {
		if match != "" {
			var stype = match[len(match)-1:] // Strip off the time span.
			var tspan = match[:len(match)-1] // Strip off the span type.
			var float, _ = stc.ParseFloat(tspan, 64)
			switch stype {
			case "-":
				sign = -1
			case "W":
				milliseconds += float * float64(millisecondsPerWeek)
			case "Y":
				milliseconds += float * float64(millisecondsPerYear)
			case "M":
				if isTime {
					milliseconds += float * float64(millisecondsPerMinute)
				} else {
					milliseconds += float * float64(millisecondsPerMonth)
				}
			case "D":
				milliseconds += float * float64(millisecondsPerDay)
			case "T":
				isTime = true
			case "H":
				milliseconds += float * float64(millisecondsPerHour)
			case "S":
				milliseconds += float * float64(millisecondsPerSecond)
			}
		}
	}
	return int(sign * milliseconds)
}

// This private function returns the string for the specified floating point
// number.
func stringFromFloat(float float64) string {
	var string_ string
	switch float {
	case mat.E:
		string_ = "e"
	case -mat.E:
		string_ = "-e"
	case mat.Pi:
		string_ = "π"
	case -mat.Pi:
		string_ = "-π"
	case mat.Phi:
		string_ = "φ"
	case -mat.Phi:
		string_ = "-φ"
	case mat.Pi * 2.0:
		string_ = "τ"
	case -mat.Pi * 2.0:
		string_ = "-τ"
	case mat.Inf(1):
		string_ = "∞"
	case mat.Inf(-1):
		string_ = "-∞"
	default:
		string_ = stc.FormatFloat(float, 'G', -1, 64)
	}
	return string_
}

// This private function returns the string for the specified imaginary floating
// point number.
func stringFromImaginary(imaginary float64) string {
	var string_ string
	switch imaginary {
	case 1:
		string_ = "i"
	case -1:
		string_ = "-i"
	default:
		string_ = stringFromFloat(imaginary) + "i"
	}
	return string_
}

// PACKAGE CLASSES

// This function returns a reference to the angle class type and
// initializes any class constants.
func Angle() *angleClass_ {
	var class = &angleClass_{
		angle_(0.0),    // Angle.Zero()
		angle_(mat.Pi), // Angle.Pi()
		angle_(tau),    // Angle.Tau()
	}
	return class
}

// This function returns a reference to the boolean class type and
// initializes any class constants.
func Boolean() *booleanClass_ {
	var class = &booleanClass_{
		boolean_(false), // Boolean.False()
		boolean_(true),  // Boolean.True()
	}
	return class
}

// This function returns a reference to the character class type and
// initializes any class constants.
func Character() *characterClass_ {
	var class = &characterClass_{
		character_(0),            // Character.Minimum()
		character_(mat.MaxInt32), // Character.Maximum()
	}
	return class
}

// This function returns a reference to the citation class type and
// initializes any class constants.
func Citation() *citationClass_ {
	var class = &citationClass_{
		// No class constants to initialize.
	}
	return class
}

// This function returns a reference to the duration class type and
// initializes any class constants.
func Duration() *durationClass_ {
	var class = &durationClass_{
		millisecondsPerSecond, // Duration.MillisecondsPerSecond()
		millisecondsPerMinute, // Duration.MillisecondsPerMinute()
		millisecondsPerHour,   // Duration.MillisecondsPerHour()
		millisecondsPerDay,    // Duration.MillisecondsPerDay()
		millisecondsPerWeek,   // Duration.MillisecondsPerWeek()
		millisecondsPerMonth,  // Duration.MillisecondsPerMonth()
		millisecondsPerYear,   // Duration.MillisecondsPerYear()
		daysPerMonth,          // Duration.DaysPerMonth()
		daysPerYear,           // Duration.DaysPerYear()
		weeksPerMonth,         // Duration.WeeksPerMonth()
	}
	return class
}

// This function returns a reference to the float class type and
// initializes any class constants.
func Float() *floatClass_ {
	var class = &floatClass_{
		float_(mat.Inf(-1)), // Float.Minimum()
		float_(mat.Inf(1)),  // Float.Maximum()
	}
	return class
}

// This function returns a reference to the integer class type and
// initializes any class constants.
func Integer() *integerClass_ {
	var class = &integerClass_{
		integer_(mat.MinInt), // Integer.Minimum()
		integer_(mat.MaxInt), // Integer.Maximum()
	}
	return class
}

// This function returns a reference to the moment class type and
// initializes any class constants.
func Moment() *momentClass_ {
	var class = &momentClass_{
		moment_(0), // Moment.Epoch()
	}
	return class
}

// This function returns a reference to the number class type and
// initializes any class constants.
func Number() *numberClass_ {
	var class = &numberClass_{
		number_(complex(0, 0)),          // Number.Zero()
		number_(complex(1, 0)),          // Number.One()
		number_(complex(0, 1)),          // Number.I()
		number_(complex(mat.E, 0)),      // Number.E()
		number_(complex(mat.Pi, 0)),     // Number.Pi()
		number_(complex(mat.Phi, 0)),    // Number.Phi()
		number_(complex(2.0*mat.Pi, 0)), // Number.Tau()
		number_(cmp.Inf()),              // Number.Infinity()
		number_(cmp.NaN()),              // Number.Undefined()
	}
	return class
}

// This function returns a reference to the pattern class type and
// initializes any class constants.
func Pattern() *patternClass_ {
	var class = &patternClass_{
		pattern_(`.*`),     // Pattern.Any()
		pattern_(`^none$`), // Pattern.None()
	}
	return class
}

// This function returns a reference to the percentage class type and
// initializes any class constants.
func Percentage() *percentageClass_ {
	var class = &percentageClass_{
		// No class constants to initialize.
	}
	return class
}

// This function returns a reference to the probability class type and
// initializes any class constants.
func Probability() *probabilityClass_ {
	var class = &probabilityClass_{
		probability_(0), // Probability.Minimum()
		probability_(1), // Probability.Maximum()
	}
	return class
}

// This function returns a reference to the resource class type and
// initializes any class constants.
func Resource() *resourceClass_ {
	var class = &resourceClass_{
		// No class constants to initialize.
	}
	return class
}
