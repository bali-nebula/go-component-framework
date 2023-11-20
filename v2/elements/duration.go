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
	mat "math"
	stc "strconv"
	sts "strings"
)

// CONSTANT DEFINITIONS

// These are locked to the Earth's daily revolutions.

const MillisecondsPerSecond int = 1000
const MillisecondsPerMinute int = MillisecondsPerSecond * 60
const MillisecondsPerHour int = MillisecondsPerMinute * 60
const MillisecondsPerDay int = MillisecondsPerHour * 24
const MillisecondsPerWeek int = MillisecondsPerDay * 7

// These are locked to the Earth's yearly orbit around the sun.

const MillisecondsPerYear int = 31556952000
const MillisecondsPerMonth int = MillisecondsPerYear / 12 // An average but exact value.

// Tying the two together is where things get messy.

const DaysPerMonth float64 = float64(MillisecondsPerMonth) / float64(MillisecondsPerDay)   // ~30.436875 days/month
const WeeksPerMonth float64 = float64(MillisecondsPerMonth) / float64(MillisecondsPerWeek) // ~4.348125 weeks/month
const DaysPerYear float64 = float64(MillisecondsPerYear) / float64(MillisecondsPerDay)     // ~365.2425 days/year

// DURATION INTERFACE

// This constructor creates a new duration of time element from the specified
// integer number of milliseconds.
func DurationFromMilliseconds(milliseconds int) abs.DurationLike {
	var duration = duration_(milliseconds)
	return duration
}

// This constructor creates a new duration from the specified string value.
func DurationFromString(string_ string) abs.DurationLike {
	var matches = uti.DurationMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a duration from an invalid string: %v", string_)
		panic(message)
	}
	var milliseconds = millisecondsFromMatches(matches)
	var duration = DurationFromMilliseconds(milliseconds)
	return duration
}

// This constructor returns the minimum value for a duration of time.
func MinimumDuration() abs.DurationLike {
	var duration = DurationFromMilliseconds(0)
	return duration
}

// This constructor returns the maximum value for a duration of time.
func MaximumDuration() abs.DurationLike {
	var duration = DurationFromMilliseconds(mat.MaxInt)
	return duration
}

// DURATION IMPLEMENTATION

// This type defines the methods associated with time duration elements. It
// extends the native Go `int` type and its value represents the number of
// milliseconds for the entire duration of time. Durations can be negative.
type duration_ int

// DISCRETE INTERFACE

// This method returns a boolean value for this discrete element.
func (v duration_) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this discrete element.
func (v duration_) AsInteger() int {
	return int(v)
}

// LEXICAL INTERFACE

// This method returns a string value for this lexical element.
func (v duration_) AsString() string {
	var builder sts.Builder
	builder.WriteString("~")
	if v.IsNegative() {
		builder.WriteString("-")
	}
	builder.WriteString("P")
	var weeks = mat.Abs(v.AsWeeks())
	if float64(int(weeks)) == weeks {
		// It is an exact number of weeks.
		builder.WriteString(stc.FormatInt(int64(weeks), 10))
		builder.WriteString("W")
		return builder.String()
	}
	var years = v.GetYears()
	if years > 0 {
		builder.WriteString(stc.FormatInt(int64(years), 10))
		builder.WriteString("Y")
	}
	var months = v.GetMonths()
	if months > 0 {
		builder.WriteString(stc.FormatInt(int64(months), 10))
		builder.WriteString("M")
	}
	var days = v.GetDays()
	if days > 0 {
		builder.WriteString(stc.FormatInt(int64(days), 10))
		builder.WriteString("D")
	}
	var hours = v.GetHours()
	var minutes = v.GetMinutes()
	var seconds = v.GetSeconds()
	var milliseconds = v.GetMilliseconds()
	if hours+minutes+seconds+milliseconds == 0 {
		// There is no time part of the duration.
		return builder.String()
	}
	builder.WriteString("T")
	if hours > 0 {
		builder.WriteString(stc.FormatInt(int64(hours), 10))
		builder.WriteString("H")
	}
	if minutes > 0 {
		builder.WriteString(stc.FormatInt(int64(minutes), 10))
		builder.WriteString("M")
	}
	if seconds+milliseconds > 0 {
		builder.WriteString(stc.FormatInt(int64(seconds), 10))
		if milliseconds > 0 {
			builder.WriteString(".")
			builder.WriteString(stc.FormatInt(int64(milliseconds), 10))
		}
		builder.WriteString("S")
	}
	return builder.String()
}

// POLARIZED INTERFACE

// This method determines whether or not this polarized component is negative.
func (v duration_) IsNegative() bool {
	return v < 0
}

// TEMPORAL INTERFACE

// This method returns the total number of milliseconds in this duration.
func (v duration_) AsMilliseconds() float64 {
	return float64(v)
}

// This method returns the total number of seconds in this duration.
func (v duration_) AsSeconds() float64 {
	return float64(v.AsMilliseconds()) / 1000.0 // milliseconds per second
}

// This method returns the total number of minutes in this duration.
func (v duration_) AsMinutes() float64 {
	return v.AsSeconds() / 60.0 // seconds per minute
}

// This method returns the total number of hours in this duration.
func (v duration_) AsHours() float64 {
	return v.AsMinutes() / 60.0 // minutes per hour
}

// This method returns the total number of days in this duration.
func (v duration_) AsDays() float64 {
	return v.AsHours() / 24.0 // hours per day
}

// This method returns the total number of weeks in this duration.
func (v duration_) AsWeeks() float64 {
	return v.AsDays() / 7.0 // days per week
}

// This method returns the total number of months in this duration.
func (v duration_) AsMonths() float64 {
	return v.AsDays() / DaysPerMonth
}

// This method returns the total number of years in this duration.
func (v duration_) AsYears() float64 {
	return v.AsDays() / DaysPerYear
}

// This method returns the milliseconds part of this duration.
func (v duration_) GetMilliseconds() int {
	var milliseconds = magnitude(int(v)) % MillisecondsPerSecond
	return milliseconds
}

// This method returns the seconds part of this duration.
func (v duration_) GetSeconds() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * MillisecondsPerMonth)        // Strip off the months.
	milliseconds = milliseconds - (v.GetDays() * MillisecondsPerDay)            // Strip off the days.
	milliseconds = milliseconds - (v.GetHours() * MillisecondsPerHour)          // Strip off the hours.
	milliseconds = milliseconds - (v.GetMinutes() * MillisecondsPerMinute)      // Strip off the minutes.
	var seconds = milliseconds / MillisecondsPerSecond                          // Strip off the milliseconds.
	return seconds
}

// This method returns the minutes part of this duration.
func (v duration_) GetMinutes() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * MillisecondsPerMonth)        // Strip off the months.
	milliseconds = milliseconds - (v.GetDays() * MillisecondsPerDay)            // Strip off the days.
	milliseconds = milliseconds - (v.GetHours() * MillisecondsPerHour)          // Strip off the hours.
	var minutes = milliseconds / MillisecondsPerMinute                          // Strip off the seconds and below.
	return minutes
}

// This method returns the hours part of this duration.
func (v duration_) GetHours() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * MillisecondsPerMonth)        // Strip off the months.
	milliseconds = milliseconds - (v.GetDays() * MillisecondsPerDay)            // Strip off the days.
	var hours = milliseconds / MillisecondsPerHour                              // Strip off the minutes and below.
	return hours
}

// This method returns the days part of this duration.
func (v duration_) GetDays() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * MillisecondsPerMonth)        // Strip off the months.
	var days = milliseconds / MillisecondsPerDay                                // Strip off the hours and below.
	return days
}

// This method returns the weeks part of this duration.
func (v duration_) GetWeeks() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	var weeks = milliseconds / MillisecondsPerWeek                              // Strip off the days and below.
	return weeks
}

// This method returns the months part of this duration.
func (v duration_) GetMonths() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	var months = milliseconds / MillisecondsPerMonth                            // Strip off the days and below.
	return months
}

// This method returns the years part of this duration.
func (v duration_) GetYears() int {
	var milliseconds = magnitude(int(v))
	var years = milliseconds / MillisecondsPerYear // Strip off the months and below.
	return years
}

// PRIVATE FUNCTIONS

// This function returns the magnitude (absolute value) of the specified value.
func magnitude(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

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
				milliseconds += float * float64(MillisecondsPerWeek)
			case "Y":
				milliseconds += float * float64(MillisecondsPerYear)
			case "M":
				if isTime {
					milliseconds += float * float64(MillisecondsPerMinute)
				} else {
					milliseconds += float * float64(MillisecondsPerMonth)
				}
			case "D":
				milliseconds += float * float64(MillisecondsPerDay)
			case "T":
				isTime = true
			case "H":
				milliseconds += float * float64(MillisecondsPerHour)
			case "S":
				milliseconds += float * float64(MillisecondsPerSecond)
			}
		}
	}
	return int(sign * milliseconds)
}
