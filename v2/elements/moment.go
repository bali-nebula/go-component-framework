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
	tim "time"
)

// CLASS DEFINITIONS

// This private type implements the MomentLike interface.  It extends the native
// Go `int` type and represents the number of milliseconds after the UNIX Epoch
// (Midnight, January 1, 1970 UTC) for a moment of time.  All moments are based
// on UTC.  For moments before the UNIX Epoch the number of milliseconds is
// negative.
type moment_ int

// This private type defines the structure associated with the class constants
// and class functions for the moment elements.
type moments_ struct {
	epoch abs.MomentLike
}

// CLASS CONSTANTS

// This class constant represents the earliest value for a moment in time element.
func (c *moments_) MinimumValue() abs.MomentLike {
	var moment = c.FromMilliseconds(mat.MinInt)
	return moment
}

// This class constant represents the latest value for a moment in time element.
func (c *moments_) MaximumValue() abs.MomentLike {
	var moment = c.FromMilliseconds(mat.MaxInt)
	return moment
}

// This class constant represents the moment of the UNIX Epoch.
func (c *moments_) Epoch() abs.MomentLike {
	return c.epoch
}

// CLASS CONSTRUCTORS

// This constructor creates a new moment in time element for the current time
// in the UTC timezone.
func (c *moments_) Now() abs.MomentLike {
	var now = int(tim.Now().UTC().UnixMilli())
	var moment = c.FromMilliseconds(now)
	return moment
}

// This constructor creates a new moment in time element from the specified
// integer number of milliseconds since the UNIX Epoch in the UTC timezone.
func (c *moments_) FromMilliseconds(milliseconds int) abs.MomentLike {
	var moment = moment_(milliseconds)
	return moment
}

// This constructor creates a new moment in time element from the specified
// string value.
func (c *moments_) FromString(string_ string) abs.MomentLike {
	var matches = uti.MomentMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a moment from an invalid string: %v", string_)
		panic(message)
	}
	var milliseconds = hackedParseDateAsMilliseconds(matches)
	var moment = c.FromMilliseconds(milliseconds)
	return moment
}

// CLASS METHODS

// Discrete Interface

// This method returns a boolean value for this discrete element.
func (v moment_) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this discrete element.
func (v moment_) AsInteger() int {
	return int(v)
}

// Lexical Interface

// This method returns a string value for this lexical element.
func (v moment_) AsString() string {
	var builder sts.Builder
	var year = v.GetYears()
	var month = v.GetMonths()
	var day = v.GetDays()
	var hour = v.GetHours()
	var minute = v.GetMinutes()
	var second = v.GetSeconds()
	var millisecond = v.GetMilliseconds()
	builder.WriteString("<")
	builder.WriteString(stc.FormatInt(int64(year), 10))
	if month > 1 || day > 1 || hour > 0 || minute > 0 || second > 0 || millisecond > 0 {
		builder.WriteString("-")
		builder.WriteString(formatOrdinal(month, 2))
		if day > 1 || hour > 0 || minute > 0 || second > 0 || millisecond > 0 {
			builder.WriteString("-")
			builder.WriteString(formatOrdinal(day, 2))
			if hour > 0 || minute > 0 || second > 0 || millisecond > 0 {
				builder.WriteString("T")
				builder.WriteString(formatOrdinal(hour, 2))
				if minute > 0 || second > 0 || millisecond > 0 {
					builder.WriteString(":")
					builder.WriteString(formatOrdinal(minute, 2))
					if second > 0 || millisecond > 0 {
						builder.WriteString(":")
						builder.WriteString(formatOrdinal(second, 2))
						if millisecond > 0 {
							builder.WriteString(".")
							builder.WriteString(formatOrdinal(millisecond, 3))
						}
					}
				}
			}
		}
	}
	builder.WriteString(">")
	return builder.String()
}

// Temporal Interface

// This method returns the total number of milliseconds since the UNIX Epoch
// in this moment.
func (v moment_) AsMilliseconds() float64 {
	return float64(v)
}

// This method returns the total number of seconds since the UNIX Epoch
// in this moment.
func (v moment_) AsSeconds() float64 {
	return float64(v.AsMilliseconds()) / 1000.0 // milliseconds per second
}

// This method returns the total number of minutes since the UNIX Epoch
// in this moment.
func (v moment_) AsMinutes() float64 {
	return v.AsSeconds() / 60.0 // seconds per minute
}

// This method returns the total number of hours since the UNIX Epoch
// in this moment.
func (v moment_) AsHours() float64 {
	return v.AsMinutes() / 60.0 // minutes per hour
}

// This method returns the total number of days since the UNIX Epoch
// in this moment.
func (v moment_) AsDays() float64 {
	return v.AsHours() / 24.0 // hours per day
}

// This method returns the total number of weeks since the UNIX Epoch
// in this moment.
func (v moment_) AsWeeks() float64 {
	return v.AsDays() / 7.0 // days per week
}

// This method returns the total number of months since the UNIX Epoch
// in this moment.
func (v moment_) AsMonths() float64 {
	return v.AsDays() / Duration.DaysPerMonth()
}

// This method returns the total number of years since the UNIX Epoch
// in this moment.
func (v moment_) AsYears() float64 {
	return v.AsDays() / Duration.DaysPerYear()
}

// This method returns the millisecond part of this moment.
func (v moment_) GetMilliseconds() int {
	var time = v.asTime()
	return time.Nanosecond() / 1e6
}

// This method returns the second part of this moment.
func (v moment_) GetSeconds() int {
	var time = v.asTime()
	return time.Second()
}

// This method returns the minute part of this moment.
func (v moment_) GetMinutes() int {
	var time = v.asTime()
	return time.Minute()
}

// This method returns the hour part of this moment.
func (v moment_) GetHours() int {
	var time = v.asTime()
	return time.Hour()
}

// This method returns the day part of this moment.
func (v moment_) GetDays() int {
	var time = v.asTime()
	return time.Day()
}

// This method returns the week part of this moment.
func (v moment_) GetWeeks() int {
	var time = v.asTime()
	var _, week = time.ISOWeek()
	return week
}

// This method returns the month part of this moment.
func (v moment_) GetMonths() int {
	var time = v.asTime()
	return int(time.Month())
}

// This method returns the year part of this moment.
func (v moment_) GetYears() int {
	var time = v.asTime()
	return time.Year()
}

// Private Interface

// This private function returns the go Time value for the specified
// UNIX-based milliseconds.
func (v moment_) asTime() tim.Time {
	var milliseconds int64 = int64(v)
	return tim.UnixMilli(milliseconds).UTC()
}

// CLASS FUNCTIONS

// This library function returns the duration of time between the two specified
// moments in tim.
func (c *moments_) Duration(first, second abs.MomentLike) abs.DurationLike {
	return Duration.FromMilliseconds(second.AsInteger() - first.AsInteger())
}

// This library function returns the moment in time that is earlier than the
// specified moment in time by the specified duration of tim.
func (c *moments_) Earlier(moment abs.MomentLike, duration abs.DurationLike) abs.MomentLike {
	return c.FromMilliseconds(moment.AsInteger() - duration.AsInteger())
}

// This library function returns the moment in time that is later than the
// specified moment in time by the specified duration of tim.
func (c *moments_) Later(moment abs.MomentLike, duration abs.DurationLike) abs.MomentLike {
	return c.FromMilliseconds(moment.AsInteger() + duration.AsInteger())
}
