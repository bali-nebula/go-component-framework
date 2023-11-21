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

// DURATION CLASS DEFINITION

// These private constants represents the ratios of various units of time.
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
	weeksPerMonth float64 = float64(millisecondsPerMonth) / float64(millisecondsPerWeek) // ~4.348125 weeks/month
	daysPerYear   float64 = float64(millisecondsPerYear) / float64(millisecondsPerDay)   // ~365.2425 days/year
)

// This public singleton creates a unique name space for the duration class.
var Duration = &durations_{
	millisecondsPerSecond,
	millisecondsPerMinute,
	millisecondsPerHour,
	millisecondsPerDay,
	millisecondsPerWeek,
	millisecondsPerMonth,
	millisecondsPerYear,
	daysPerMonth,
	weeksPerMonth,
	daysPerYear,
}

// This private type defines the structure associated with the class constants
// and class methods for the duration elements.
type durations_ struct {
	millisecondsPerSecond int
	millisecondsPerMinute int
	millisecondsPerHour   int
	millisecondsPerDay    int
	millisecondsPerWeek   int
	millisecondsPerMonth  int
	millisecondsPerYear   int
	daysPerMonth          float64
	weeksPerMonth         float64
	daysPerYear           float64
}

// DURATION CLASS CONSTANTS

// This class constant represents the number of milliseconds in a second.
func (t *durations_) MillisecondsPerSecond() int {
	return t.millisecondsPerSecond
}

// This class constant represents the number of milliseconds in a minute.
func (t *durations_) MillisecondsPerMinute() int {
	return t.millisecondsPerMinute
}

// This class constant represents the number of milliseconds in a hour.
func (t *durations_) MillisecondsPerHour() int {
	return t.millisecondsPerHour
}

// This class constant represents the number of milliseconds in a day.
func (t *durations_) MillisecondsPerDay() int {
	return t.millisecondsPerDay
}

// This class constant represents the number of milliseconds in a week.
func (t *durations_) MillisecondsPerWeek() int {
	return t.millisecondsPerWeek
}

// This class constant represents the number of milliseconds in a month.
func (t *durations_) MillisecondsPerMonth() int {
	return t.millisecondsPerMonth
}

// This class constant represents the number of milliseconds in a year.
func (t *durations_) MillisecondsPerYear() int {
	return t.millisecondsPerYear
}

// This class constant represents the number of days in a month.
func (t *durations_) DaysPerMonth() float64 {
	return t.daysPerMonth
}

// This class constant represents the number of weeks in a month.
func (t *durations_) WeeksPerMonth() float64 {
	return t.weeksPerMonth
}

// This class constant represents the number of days in a year.
func (t *durations_) DaysPerYear() float64 {
	return t.daysPerYear
}

// DURATION CLASS METHODS

// This constructor creates a new duration of time element from the specified
// integer number of milliseconds.
func (t *durations_) FromMilliseconds(milliseconds int) abs.DurationLike {
	var duration = duration_(milliseconds)
	return duration
}

// This constructor creates a new duration from the specified string value.
func (t *durations_) FromString(string_ string) abs.DurationLike {
	var matches = uti.DurationMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a duration from an invalid string: %v", string_)
		panic(message)
	}
	var milliseconds = millisecondsFromMatches(matches)
	var duration = t.FromMilliseconds(milliseconds)
	return duration
}

// This constructor returns the minimum value for a duration of time.
func (t *durations_) MinimumValue() abs.DurationLike {
	var duration = t.FromMilliseconds(0)
	return duration
}

// This constructor returns the maximum value for a duration of time.
func (t *durations_) MaximumValue() abs.DurationLike {
	var duration = t.FromMilliseconds(mat.MaxInt)
	return duration
}

// DURATION INSTANCE METHODS

// This private type implements the DurationLike interface.  It extends the
// native Go `int` type and its value represents the number of milliseconds
// for the entire duration of time. Durations can be negative.
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
	return v.AsDays() / daysPerMonth
}

// This method returns the total number of years in this duration.
func (v duration_) AsYears() float64 {
	return v.AsDays() / daysPerYear
}

// This method returns the milliseconds part of this duration.
func (v duration_) GetMilliseconds() int {
	var milliseconds = magnitude(int(v)) % millisecondsPerSecond
	return milliseconds
}

// This method returns the seconds part of this duration.
func (v duration_) GetSeconds() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * millisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * millisecondsPerMonth)        // Strip off the months.
	milliseconds = milliseconds - (v.GetDays() * millisecondsPerDay)            // Strip off the days.
	milliseconds = milliseconds - (v.GetHours() * millisecondsPerHour)          // Strip off the hours.
	milliseconds = milliseconds - (v.GetMinutes() * millisecondsPerMinute)      // Strip off the minutes.
	var seconds = milliseconds / millisecondsPerSecond                          // Strip off the milliseconds.
	return seconds
}

// This method returns the minutes part of this duration.
func (v duration_) GetMinutes() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * millisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * millisecondsPerMonth)        // Strip off the months.
	milliseconds = milliseconds - (v.GetDays() * millisecondsPerDay)            // Strip off the days.
	milliseconds = milliseconds - (v.GetHours() * millisecondsPerHour)          // Strip off the hours.
	var minutes = milliseconds / millisecondsPerMinute                          // Strip off the seconds and below.
	return minutes
}

// This method returns the hours part of this duration.
func (v duration_) GetHours() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * millisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * millisecondsPerMonth)        // Strip off the months.
	milliseconds = milliseconds - (v.GetDays() * millisecondsPerDay)            // Strip off the days.
	var hours = milliseconds / millisecondsPerHour                              // Strip off the minutes and below.
	return hours
}

// This method returns the days part of this duration.
func (v duration_) GetDays() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * millisecondsPerYear) // Strip off the years.
	milliseconds = milliseconds - (v.GetMonths() * millisecondsPerMonth)        // Strip off the months.
	var days = milliseconds / millisecondsPerDay                                // Strip off the hours and below.
	return days
}

// This method returns the weeks part of this duration.
func (v duration_) GetWeeks() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * millisecondsPerYear) // Strip off the years.
	var weeks = milliseconds / millisecondsPerWeek                              // Strip off the days and below.
	return weeks
}

// This method returns the months part of this duration.
func (v duration_) GetMonths() int {
	var milliseconds = magnitude(int(v)) - (v.GetYears() * millisecondsPerYear) // Strip off the years.
	var months = milliseconds / millisecondsPerMonth                            // Strip off the days and below.
	return months
}

// This method returns the years part of this duration.
func (v duration_) GetYears() int {
	var milliseconds = magnitude(int(v))
	var years = milliseconds / millisecondsPerYear // Strip off the months and below.
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
