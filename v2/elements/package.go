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

// PACKAGE CONSTANTS

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

// PACKAGE TYPES

// This public singleton exports a unique name space for the angle class.
var Angle = &angles_{
	angle_(0.0),    // ~0
	angle_(mat.Pi), // ~pi
	angle_(tau),    // ~tau
}

// This public singleton exports a unique name space for the boolean class.
var Boolean = &booleans_{
	boolean_(false),
	boolean_(true),
}

// This public singleton exports a unique name space for the character class.
var Character = &characters_{
	character_(0),            // minimum
	character_(mat.MaxInt32), // maximum
}

// This public singleton exports a unique name space for the citation class.
var Citation = &citations_{}

// This public singleton exports a unique name space for the duration class.
var Duration = &durations_{
	millisecondsPerSecond,
	millisecondsPerMinute,
	millisecondsPerHour,
	millisecondsPerDay,
	millisecondsPerWeek,
	millisecondsPerMonth,
	millisecondsPerYear,
	daysPerMonth,
	daysPerYear,
	weeksPerMonth,
}

// This public singleton exports a unique name space for the float class.
var Float = &floats_{
	float_(mat.Inf(-1)), // minimum
	float_(mat.Inf(1)),  // maximum
}

// This public singleton exports a unique name space for the integer class.
var Integer = &integers_{
	integer_(mat.MinInt), // minimum
	integer_(mat.MaxInt), // maximum
}

// This public singleton exports a unique name space for the moment class.
var Moment = &moments_{
	moment_(0), // UNIX Epoch
}

// This public singleton exports a unique name space for the number class.
var Number = &numbers_{
	number_(complex(0, 0)),          // 0
	number_(complex(1, 0)),          // 1
	number_(complex(0, 1)),          // i
	number_(complex(mat.E, 0)),      // e
	number_(complex(mat.Pi, 0)),     // pi
	number_(complex(mat.Phi, 0)),    // phi
	number_(complex(2.0*mat.Pi, 0)), // tau
	number_(cmp.Inf()),              // infinity
	number_(cmp.NaN()),              // undefined
}

// This public singleton exports a unique name space for the pattern class.
var Pattern = &patterns_{
	pattern_(`.*`),     // any
	pattern_(`^none$`), // none
}

// This public singleton exports a unique name space for the percentage class.
var Percentage = &percentages_{}

// This public singleton exports a unique name space for the probability class.
var Probability = &probabilities_{
	probability_(0), // minimum
	probability_(1), // maximum
}

// This public singleton exports a unique name space for the resource class.
var Resource = &resources_{}

// PACKAGE FUNCTIONS

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
