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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	mat "math"
	stc "strconv"
	str "strings"
)

// DURATION CONSTANTS

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

// This constructor attempts to create a new duration from the specified
// formatted string. It returns a duration value and whether or not the string
// contained a valid duration.
// For valid string formats for this type see `../abstractions/language.go`.
func DurationFromString(v string) (Duration, bool) {
	var duration, ok = stringToDuration(v)
	return Duration(duration), ok
}

// This type defines the methods associated with time duration elements. It
// extends the native Go `int` type and its value represents the number of
// milliseconds for the entire duration of time. Durations can be negative.
type Duration int

// POLARIZED INTERFACE

// This method determines whether or not this polarized component is negative.
func (v Duration) IsNegative() bool {
	return v < 0.0
}

	v.state.AppendString(s)
}

// TEMPORAL INTERFACE

// This method returns the total number of milliseconds in this duration.
func (v Duration) AsMilliseconds() float64 {
	return float64(v)
}

// This method returns the total number of seconds in this duration.
func (v Duration) AsSeconds() float64 {
	return float64(v.AsMilliseconds()) / 1000.0 // milliseconds per second
}

// This method returns the total number of minutes in this duration.
func (v Duration) AsMinutes() float64 {
	return v.AsSeconds() / 60.0 // seconds per minute
}

// This method returns the total number of hours in this duration.
func (v Duration) AsHours() float64 {
	return v.AsMinutes() / 60.0 // minutes per hour
}

// This method returns the total number of days in this duration.
func (v Duration) AsDays() float64 {
	return v.AsHours() / 24.0 // hours per day
}

// This method returns the total number of weeks in this duration.
func (v Duration) AsWeeks() float64 {
	return v.AsDays() / 7.0 // days per week
}

// This method returns the total number of months in this duration.
func (v Duration) AsMonths() float64 {
	return v.AsDays() / DaysPerMonth
}

// This method returns the total number of years in this duration.
func (v Duration) AsYears() float64 {
	return v.AsDays() / DaysPerYear
}

// This method returns the milliseconds part of this duration.
func (v Duration) GetMilliseconds() int {
	var milliseconds = magnitude(int(v)) % MillisecondsPerSecond
	return milliseconds
}

// This method returns the seconds part of this duration.
func (v Duration) GetSeconds() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * MillisecondsPerMonth)        // Strip off the months.
	milliseconds = milliseconds - (v.GetDays() * MillisecondsPerDay)            // Strip off the days.
	milliseconds = milliseconds - (v.GetHours() * MillisecondsPerHour)          // Strip off the hours.
	milliseconds = milliseconds - (v.GetMinutes() * MillisecondsPerMinute)      // Strip off the minutes.
	var seconds = milliseconds / MillisecondsPerSecond                          // Strip off the milliseconds.
	return seconds
}

// This method returns the minutes part of this duration.
func (v Duration) GetMinutes() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * MillisecondsPerMonth)        // Strip off the months.
	milliseconds = milliseconds - (v.GetDays() * MillisecondsPerDay)            // Strip off the days.
	milliseconds = milliseconds - (v.GetHours() * MillisecondsPerHour)          // Strip off the hours.
	var minutes = milliseconds / MillisecondsPerMinute                          // Strip off the seconds and below.
	return minutes
}

// This method returns the hours part of this duration.
func (v Duration) GetHours() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * MillisecondsPerMonth)        // Strip off the months.
	milliseconds = milliseconds - (v.GetDays() * MillisecondsPerDay)            // Strip off the days.
	var hours = milliseconds / MillisecondsPerHour                              // Strip off the minutes and below.
	return hours
}

// This method returns the days part of this duration.
func (v Duration) GetDays() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * MillisecondsPerMonth)        // Strip off the months.
	var days = milliseconds / MillisecondsPerDay                                // Strip off the hours and below.
	return days
}

// This method returns the weeks part of this duration.
func (v Duration) GetWeeks() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	var weeks = milliseconds / MillisecondsPerWeek                              // Strip off the days and below.
	return weeks
}

// This method returns the months part of this duration.
func (v Duration) GetMonths() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * MillisecondsPerYear) // Strip off the years.
	var months = milliseconds / MillisecondsPerMonth                            // Strip off the days and below.
	return months
}

// This method returns the years part of this duration.
func (v Duration) GetYears() int {
	var milliseconds = magnitude(int(v))
	var years = milliseconds / MillisecondsPerYear // Slice off the months and below.
	return years
}

// DURATIONS LIBRARY

// This singleton creates a unique name space for the library functions for
// duration elements.
var Durations = &durations{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for duration elements.
type durations struct{}

// SCALABLE INTERFACE

// This library function returns the inverse of the specified duration.
func (l *durations) Inverse(duration Duration) Duration {
	return Duration(int(-duration))
}

// This library function returns the sum of the specified durations.
func (l *durations) Sum(first, second Duration) Duration {
	return Duration(int(first + second))
}

// This library function returns the difference of the specified durations.
func (l *durations) Difference(first, second Duration) Duration {
	return Duration(int(first - second))
}

// This library function returns the specified duration scaled by the specified
// factor.
func (l *durations) Scaled(duration Duration, factor float64) Duration {
	return Duration(int(mat.Round(float64(duration) * factor)))
}

// PRIVATE FUNCTIONS

// This function returns the magnitude (absolute value) of the specified value.
func magnitude(value int) int {
	if value < 0 {
		return -value
	}
	return value
}

// This function parses a duration string and returns the corresponding number
// of milliseconds for that duration.
func stringToDuration(v string) (int, bool) {
	var matches = abs.ScanDuration([]byte(v))
	if len(matches) == 0 {
		return 0, false
	}
	var milliseconds = 0.0
	var sign = 1.0
	var isTime = false
	for _, match := range matches[1:] {
		if match != "" {
			var stype = match[len(match)-1:] // Strip off the time span.
			var tspan = match[:len(match)-1] // Strip off the span type.
			var value, err = stc.ParseFloat(tspan, 64)
			if len(tspan) > 0 && err != nil {
				return 0, false
			}
			switch stype {
			case "-":
				sign = -1
			case "W":
				milliseconds += value * float64(MillisecondsPerWeek)
			case "Y":
				milliseconds += value * float64(MillisecondsPerYear)
			case "M":
				if isTime {
					milliseconds += value * float64(MillisecondsPerMinute)
				} else {
					milliseconds += value * float64(MillisecondsPerMonth)
				}
			case "D":
				milliseconds += value * float64(MillisecondsPerDay)
			case "T":
				isTime = true
			case "H":
				milliseconds += value * float64(MillisecondsPerHour)
			case "S":
				milliseconds += value * float64(MillisecondsPerSecond)
			}
		}
	}
	return int(sign * milliseconds), true
}
