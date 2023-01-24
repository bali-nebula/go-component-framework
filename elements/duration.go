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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	mat "math"
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

// DURATION IMPLEMENTATION

// This constructor creates a new duration of time element from the specified
// integer number of milliseconds.
func DurationFromInteger(milliseconds int) abs.DurationLike {
	return Duration(milliseconds)
}

// This constructor returns the minimum value for a duration of time.
func MinimumDuration() abs.DurationLike {
	return Duration(0)
}

// This constructor returns the maximum value for a duration of time.
func MaximumDuration() abs.DurationLike {
	return Duration(mat.MaxInt)
}

// This type defines the methods associated with time duration elements. It
// extends the native Go `int` type and its value represents the number of
// milliseconds for the entire duration of time. Durations can be negative.
type Duration int

// DISCRETE INTERFACE

// This method returns a boolean value for this discrete element.
func (v Duration) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this discrete element.
func (v Duration) AsInteger() int {
	return int(v)
}

// POLARIZED INTERFACE

// This method determines whether or not this polarized component is negative.
func (v Duration) IsNegative() bool {
	return v < 0
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
