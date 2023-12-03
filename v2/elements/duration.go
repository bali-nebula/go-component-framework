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
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	mat "math"
	stc "strconv"
	sts "strings"
)

// CLASS DEFINITIONS

// This private type implements the DurationLike interface.  It extends the
// native Go `int` type and its value represents the number of milliseconds
// for the entire duration of time. Durations can be negative.
type duration_ int

// This private type defines the structure associated with the class constants
// and class functions for the duration elements.
type durationClass_ struct {
	millisecondsPerSecond int
	millisecondsPerMinute int
	millisecondsPerHour   int
	millisecondsPerDay    int
	millisecondsPerWeek   int
	millisecondsPerMonth  int
	millisecondsPerYear   int
	daysPerMonth          float64
	daysPerYear           float64
	weeksPerMonth         float64
}

// CLASS CONSTANTS

// This class constant represents the minimum value for a duration of time.
func (c *durationClass_) MinimumValue() DurationLike {
	var duration = c.FromMilliseconds(0)
	return duration
}

// This class constant represents the maximum value for a duration of time.
func (c *durationClass_) MaximumValue() DurationLike {
	var duration = c.FromMilliseconds(mat.MaxInt)
	return duration
}

// This class constant represents the number of milliseconds in a second.
func (c *durationClass_) MillisecondsPerSecond() int {
	return c.millisecondsPerSecond
}

// This class constant represents the number of milliseconds in a minute.
func (c *durationClass_) MillisecondsPerMinute() int {
	return c.millisecondsPerMinute
}

// This class constant represents the number of milliseconds in a hour.
func (c *durationClass_) MillisecondsPerHour() int {
	return c.millisecondsPerHour
}

// This class constant represents the number of milliseconds in a day.
func (c *durationClass_) MillisecondsPerDay() int {
	return c.millisecondsPerDay
}

// This class constant represents the number of milliseconds in a week.
func (c *durationClass_) MillisecondsPerWeek() int {
	return c.millisecondsPerWeek
}

// This class constant represents the number of milliseconds in a month.
func (c *durationClass_) MillisecondsPerMonth() int {
	return c.millisecondsPerMonth
}

// This class constant represents the number of milliseconds in a year.
func (c *durationClass_) MillisecondsPerYear() int {
	return c.millisecondsPerYear
}

// This class constant represents the number of days in a month.
func (c *durationClass_) DaysPerMonth() float64 {
	return c.daysPerMonth
}

// This class constant represents the number of days in a year.
func (c *durationClass_) DaysPerYear() float64 {
	return c.daysPerYear
}

// This class constant represents the number of weeks in a month.
func (c *durationClass_) WeeksPerMonth() float64 {
	return c.weeksPerMonth
}

// CLASS CONSTRUCTORS

// This constructor creates a new duration of time element from the specified
// integer number of milliseconds.
func (c *durationClass_) FromMilliseconds(milliseconds int) DurationLike {
	var duration = duration_(milliseconds)
	return duration
}

// This constructor creates a new duration from the specified string value.
func (c *durationClass_) FromString(string_ string) DurationLike {
	var matches = uti.DurationMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a duration from an invalid string: %v", string_)
		panic(message)
	}
	var milliseconds = millisecondsFromMatches(matches)
	var duration = c.FromMilliseconds(milliseconds)
	return duration
}

// CLASS METHODS

// Discrete Interface

// This method returns a boolean value for this discrete element.
func (v duration_) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this discrete element.
func (v duration_) AsInteger() int {
	return int(v)
}

// Lexical Interface

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

// Polarized Interface

// This method determines whether or not this polarized component is negative.
func (v duration_) IsNegative() bool {
	return v < 0
}

// Temporal Interface

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
