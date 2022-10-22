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
	fmt "fmt"
	str "strings"
	tim "time"
)

// MOMENT INTERFACE

// This constructor creates a new moment in time element for the current time
// in the UTC timezone.
func Now() Moment {
	var now = tim.Now().UTC().UnixMilli()
	return Moment(now)
}

// This constructor attempts to create a new moment in time from the specified
// formatted string. It returns a moment value and whether or not the string
// contained a valid moment in tim.
// For valid string formats for this type see `../abstractions/language.go`.
func MomentFromString(v string) (Moment, bool) {
	var moment, ok = stringToMoment(v)
	return Moment(moment), ok
}

// This type defines the methods associated with moment in time elements. It
// extends the native Go int type and represents the number of milliseconds
// after the UNIX epoc (Midnight, January 1, 1970 UTC) for a moment of tim.
// All moments are based on UTC.
type Moment int

// TEMPORAL INTERFACE

// This method returns the total number of milliseconds since epoc in this moment.
func (v Moment) AsMilliseconds() float64 {
	return float64(v)
}

// This method returns the total number of seconds since epoc in this moment.
func (v Moment) AsSeconds() float64 {
	return float64(v.AsMilliseconds()) / 1000.0 // milliseconds per second
}

// This method returns the total number of minutes since epoc in this moment.
func (v Moment) AsMinutes() float64 {
	return v.AsSeconds() / 60.0 // seconds per minute
}

// This method returns the total number of hours since epoc in this moment.
func (v Moment) AsHours() float64 {
	return v.AsMinutes() / 60.0 // minutes per hour
}

// This method returns the total number of days since epoc in this moment.
func (v Moment) AsDays() float64 {
	return v.AsHours() / 24.0 // hours per day
}

// This method returns the total number of weeks since epoc in this moment.
func (v Moment) AsWeeks() float64 {
	return v.AsDays() / 7.0 // days per week
}

// This method returns the total number of months since epoc in this moment.
func (v Moment) AsMonths() float64 {
	return v.AsDays() / DaysPerMonth
}

// This method returns the total number of years since epoc in this moment.
func (v Moment) AsYears() float64 {
	return v.AsDays() / DaysPerYear
}

// This method returns the millisecond part of this moment.
func (v Moment) GetMilliseconds() int {
	var t = momentToTime(v)
	return t.Nanosecond() / 1e6
}

// This method returns the second part of this moment.
func (v Moment) GetSeconds() int {
	var t = momentToTime(v)
	return t.Second()
}

// This method returns the minute part of this moment.
func (v Moment) GetMinutes() int {
	var t = momentToTime(v)
	return t.Minute()
}

// This method returns the hour part of this moment.
func (v Moment) GetHours() int {
	var t = momentToTime(v)
	return t.Hour()
}

// This method returns the day part of this moment.
func (v Moment) GetDays() int {
	var t = momentToTime(v)
	return t.Day()
}

// This method returns the week part of this moment.
func (v Moment) GetWeeks() int {
	var t = momentToTime(v)
	var _, week = t.ISOWeek()
	return week
}

// This method returns the month part of this moment.
func (v Moment) GetMonths() int {
	var t = momentToTime(v)
	return int(t.Month())
}

// This method returns the year part of this moment.
func (v Moment) GetYears() int {
	var t = momentToTime(v)
	return t.Year()
}

// MOMENTS LIBRARY

// This singleton creates a unique name space for the library functions for
// moment elements.
var Moments = &moments{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for moment elements.
type moments struct{}

// This library function returns the duration of time between the two specified
// moments in tim.
func (l *moments) Duration(first, second Moment) Duration {
	return Duration(int(second - first))
}

// This library function returns the moment in time that is earlier than the
// specified moment in time by the specified duration of tim.
func (l *moments) Earlier(moment Moment, duration Duration) Moment {
	return Moment(int(moment) - int(duration))
}

// This library function returns the moment in time that is later than the
// specified moment in time by the specified duration of tim.
func (l *moments) Later(moment Moment, duration Duration) Moment {
	return Moment(int(moment) + int(duration))
}

// PRIVATE FUNCTIONS

// This list contains the supported ISO 8601 date-time formats. Note: the Go
// templates in this list must contain their exact values. If you are curious
// why check out this posting:
//   - https://medium.com/@simplyianm/how-go-solves-date-and-time-formatting-8a932117c41c
//
// A good mnemonic to use to remember the pattern is:
//
//	  1    2    3     4      5     6      7
//	month day hour minute second year timezone
//
//	"January 2nd, 2006 at 3:04:05pm in the MST timezone"
//
// The "Z" in the templates corresponds to the UTC timezone.
//
// Anyway, it is what it is, but this hides it from the rest of the code.
var isoFormats = [...]string{ // The "..." makes it a fixed sized array.
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
	"<-2006-01-02T15:04:05.000>"}

// This function parses a moment string and returns the corresponding number
// of milliseconds since UNIX epoc for that moment in tim. We take advantage
// of the Go time package to do the parsing and decoding of the string.
func stringToMoment(v string) (int, bool) {
	var result int
	for _, format := range isoFormats {
		var t, err = hackedParse(format, string(v)) // Parsed in UTC.
		if err == nil {
			result = int(t.UnixMilli())
			return result, true
		}
	}
	return result, false
}

// This function returns the go Time value for the specified UNIX milliseconds.
func momentToTime(v Moment) tim.Time {
	var milliseconds int64 = int64(v)
	return tim.UnixMilli(milliseconds).UTC()
}

// This function formats the specified ordinal value to exactly two digits.
func formatOrdinal(ordinal int) string {
	return fmt.Sprintf("%02d", ordinal)
}

// The Go tim.Parse() function cannot handle negative years even though the
// tim.Time.Format() method will correctly print negative years. The Go team
// has labeled this issue as "unfortunate" and will not fix it. Alas...
func hackedParse(layout string, value string) (tim.Time, error) {
	var date, err = tim.Parse(layout, value)
	if err != nil {
		return date, err
	}
	if str.HasPrefix(layout, "<-") {
		date = date.AddDate(-2*date.Year(), 0, 0)
	}
	return date, err
}
