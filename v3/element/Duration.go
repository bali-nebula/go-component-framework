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

package element

import (
	mat "math"
	stc "strconv"
	sts "strings"
)

// CLASS ACCESS

// Reference

var durationClass = &durationClass_{
	// These are for non-negative durations.
	minimumValue_: duration_(0),
	maximumValue_: duration_(mat.MaxInt),

	// These are locked to the Earth's daily revolutions.
	millisecondsPerSecond_: 1000,
	millisecondsPerMinute_: 60000,
	millisecondsPerHour_:   3600000,
	millisecondsPerDay_:    86400000,
	millisecondsPerWeek_:   604800000,

	// These are locked to the Earth's yearly orbit around the sun.
	millisecondsPerMonth_: 2629746000, // An average but exact value.
	millisecondsPerYear_:  31556952000,

	// Tying the two together is where things get messy.
	daysPerMonth_:  30.436875,
	daysPerYear_:   365.2425,
	weeksPerMonth_: 4.348125,
}

// Function

func Duration() DurationClassLike {
	return durationClass
}

// CLASS METHODS

// Target

type durationClass_ struct {
	minimumValue_          duration_
	maximumValue_          duration_
	millisecondsPerSecond_ int64
	millisecondsPerMinute_ int64
	millisecondsPerHour_   int64
	millisecondsPerDay_    int64
	millisecondsPerWeek_   int64
	millisecondsPerMonth_  int64
	millisecondsPerYear_   int64
	daysPerMonth_          float64
	daysPerYear_           float64
	weeksPerMonth_         float64
}

// Constants

func (c *durationClass_) MinimumValue() DurationLike {
	return c.minimumValue_
}

func (c *durationClass_) MaximumValue() DurationLike {
	return c.maximumValue_
}

func (c *durationClass_) MillisecondsPerSecond() int64 {
	return c.millisecondsPerSecond_
}

func (c *durationClass_) MillisecondsPerMinute() int64 {
	return c.millisecondsPerMinute_
}

func (c *durationClass_) MillisecondsPerHour() int64 {
	return c.millisecondsPerHour_
}

func (c *durationClass_) MillisecondsPerDay() int64 {
	return c.millisecondsPerDay_
}

func (c *durationClass_) MillisecondsPerWeek() int64 {
	return c.millisecondsPerWeek_
}

func (c *durationClass_) MillisecondsPerMonth() int64 {
	return c.millisecondsPerMonth_
}

func (c *durationClass_) MillisecondsPerYear() int64 {
	return c.millisecondsPerYear_
}

func (c *durationClass_) DaysPerMonth() float64 {
	return c.daysPerMonth_
}

func (c *durationClass_) DaysPerYear() float64 {
	return c.daysPerYear_
}

func (c *durationClass_) WeeksPerMonth() float64 {
	return c.weeksPerMonth_
}

// Constructors

func (c *durationClass_) MakeFromMilliseconds(milliseconds int64) DurationLike {
	return duration_(milliseconds)
}

func (c *durationClass_) MakeFromString(string_ string) DurationLike {
	var matches = matchDuration(string_)
	var milliseconds = durationFromMatches(matches)
	return duration_(milliseconds)
}

// Functions

// INSTANCE METHODS

// Target

type duration_ int64

// Attributes

// Discrete

func (v duration_) AsBoolean() bool {
	return v != 0
}

func (v duration_) AsInteger() int64 {
	return int64(v)
}

// Lexical

func (v duration_) AsString() string {
	var builder sts.Builder
	builder.WriteString("~")
	if v.IsNegative() {
		builder.WriteString("-")
	}
	builder.WriteString("P")
	var float = mat.Abs(v.AsWeeks())
	var weeks = int64(float)
	if float64(weeks) == float {
		// It is an exact number of weeks.
		builder.WriteString(stc.FormatInt(weeks, 10))
		builder.WriteString("W")
		return builder.String()
	}
	var years = v.GetYears()
	if years > 0 {
		builder.WriteString(stc.FormatInt(years, 10))
		builder.WriteString("Y")
	}
	var months = v.GetMonths()
	if months > 0 {
		builder.WriteString(stc.FormatInt(months, 10))
		builder.WriteString("M")
	}
	var days = v.GetDays()
	if days > 0 {
		builder.WriteString(stc.FormatInt(days, 10))
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
		builder.WriteString(stc.FormatInt(hours, 10))
		builder.WriteString("H")
	}
	if minutes > 0 {
		builder.WriteString(stc.FormatInt(minutes, 10))
		builder.WriteString("M")
	}
	if seconds+milliseconds > 0 {
		builder.WriteString(stc.FormatInt(seconds, 10))
		if milliseconds > 0 {
			builder.WriteString(".")
			builder.WriteString(stc.FormatInt(milliseconds, 10))
		}
		builder.WriteString("S")
	}
	return builder.String()
}

// Polarized

func (v duration_) IsNegative() bool {
	return v < 0
}

// Temporal

func (v duration_) AsMilliseconds() float64 {
	return float64(v)
}

func (v duration_) AsSeconds() float64 {
	return float64(v) / float64(durationClass.millisecondsPerSecond_)
}

func (v duration_) AsMinutes() float64 {
	return float64(v) / float64(durationClass.millisecondsPerMinute_)
}

func (v duration_) AsHours() float64 {
	return float64(v) / float64(durationClass.millisecondsPerHour_)
}

func (v duration_) AsDays() float64 {
	return float64(v) / float64(durationClass.millisecondsPerDay_)
}

func (v duration_) AsWeeks() float64 {
	return float64(v) / float64(durationClass.millisecondsPerWeek_)
}

func (v duration_) AsMonths() float64 {
	return float64(v) / float64(durationClass.millisecondsPerMonth_)
}

func (v duration_) AsYears() float64 {
	return float64(v) / float64(durationClass.millisecondsPerYear_)
}

// Factored

func (v duration_) GetMilliseconds() int64 {
	// Retrieve the total number of milliseconds.
	var milliseconds = magnitude(int64(v))

	// Strip off everything but the milliseconds.
	milliseconds = milliseconds % durationClass.millisecondsPerSecond_
	return milliseconds
}

func (v duration_) GetSeconds() int64 {
	// Retrieve the total number of milliseconds.
	var milliseconds = magnitude(int64(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass.millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass.millisecondsPerMonth_)

	// Strip off the days.
	milliseconds = milliseconds - (v.GetDays() * durationClass.millisecondsPerDay_)

	// Strip off the hours.
	milliseconds = milliseconds - (v.GetHours() * durationClass.millisecondsPerHour_)

	// Strip off the minutes.
	milliseconds = milliseconds - (v.GetMinutes() * durationClass.millisecondsPerMinute_)

	// Convert to seconds.
	var seconds = milliseconds / durationClass.millisecondsPerSecond_
	return seconds
}

func (v duration_) GetMinutes() int64 {
	// Retrieve the total number of milliseconds.
	var milliseconds = magnitude(int64(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass.millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass.millisecondsPerMonth_)

	// Strip off the days.
	milliseconds = milliseconds - (v.GetDays() * durationClass.millisecondsPerDay_)

	// Strip off the hours.
	milliseconds = milliseconds - (v.GetHours() * durationClass.millisecondsPerHour_)

	// Convert to minutes.
	var minutes = milliseconds / durationClass.millisecondsPerMinute_
	return minutes
}

func (v duration_) GetHours() int64 {
	// Retrieve the total number of milliseconds.
	var milliseconds = magnitude(int64(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass.millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass.millisecondsPerMonth_)

	// Strip off the days.
	milliseconds = milliseconds - (v.GetDays() * durationClass.millisecondsPerDay_)

	// Convert to hours.
	var hours = milliseconds / durationClass.millisecondsPerHour_
	return hours
}

func (v duration_) GetDays() int64 {
	// Retrieve the total number of milliseconds.
	var milliseconds = magnitude(int64(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass.millisecondsPerYear_)

	// Strip off the months.
	milliseconds = milliseconds - (v.GetMonths() * durationClass.millisecondsPerMonth_)

	// Convert to days.
	var days = milliseconds / durationClass.millisecondsPerDay_
	return days
}

func (v duration_) GetWeeks() int64 {
	// Retrieve the total number of milliseconds.
	var milliseconds = magnitude(int64(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass.millisecondsPerYear_)

	// Convert to weeks.
	var weeks = milliseconds / durationClass.millisecondsPerWeek_
	return weeks
}

func (v duration_) GetMonths() int64 {
	// Retrieve the total number of milliseconds.
	var milliseconds = magnitude(int64(v))

	// Strip off the years.
	milliseconds = milliseconds - (v.GetYears() * durationClass.millisecondsPerYear_)

	// Convert to months.
	var months = milliseconds / durationClass.millisecondsPerMonth_
	return months
}

func (v duration_) GetYears() int64 {
	// Retrieve the total number of milliseconds.
	var milliseconds = magnitude(int64(v))

	// Convert to years.
	var years = milliseconds / durationClass.millisecondsPerYear_
	return years
}

// PACKAGE FUNCTIONS

// Private

func durationFromMatches(matches []string) int64 {
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
				sign = -1.0
			case "W":
				milliseconds += float * float64(durationClass.millisecondsPerWeek_)
			case "Y":
				milliseconds += float * float64(durationClass.millisecondsPerYear_)
			case "M":
				if isTime {
					milliseconds += float * float64(durationClass.millisecondsPerMinute_)
				} else {
					milliseconds += float * float64(durationClass.millisecondsPerMonth_)
				}
			case "D":
				milliseconds += float * float64(durationClass.millisecondsPerDay_)
			case "T":
				isTime = true
			case "H":
				milliseconds += float * float64(durationClass.millisecondsPerHour_)
			case "S":
				milliseconds += float * float64(durationClass.millisecondsPerSecond_)
			}
		}
	}
	return int64(sign * milliseconds)
}

func magnitude(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}

func matchDuration(string_ string) []string {
	var matches = []string{
		string_,
	}
	// TBA - Add the pattern matching code...
	return matches
}
