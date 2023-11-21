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

// MOMENT ELEMENT CONSTANTS

const Epoch = moment_(0)

// MOMENT ELEMENT CONSTRUCTORS

// This constructor creates a new moment in time element from the specified
// integer number of milliseconds since the UNIX Epoch in the UTC timezone.
func MomentFromMilliseconds(milliseconds int) abs.MomentLike {
	var moment = moment_(milliseconds)
	return moment
}

// This constructor creates a new moment in time element from the specified
// string value.
func MomentFromString(string_ string) abs.MomentLike {
	var matches = uti.MomentMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a moment from an invalid string: %v", string_)
		panic(message)
	}
	var milliseconds = hackedParseDateAsMilliseconds(matches)
	var moment = MomentFromMilliseconds(milliseconds)
	return moment
}

// This constructor creates a new moment in time element for the current time
// in the UTC timezone.
func Now() abs.MomentLike {
	var now = int(tim.Now().UTC().UnixMilli())
	var moment = MomentFromMilliseconds(now)
	return moment
}

// This constructor returns the earliest value for a moment in time element.
func MinimumMoment() abs.MomentLike {
	var moment = MomentFromMilliseconds(mat.MinInt)
	return moment
}

// This constructor returns the latest value for a moment in time element.
func MaximumMoment() abs.MomentLike {
	var moment = MomentFromMilliseconds(mat.MaxInt)
	return moment
}

// MOMENT ELEMENT METHODS

// This private type implements the MomentLike interface.  It extends the native
// Go `int` type and represents the number of milliseconds after the UNIX Epoch
// (Midnight, January 1, 1970 UTC) for a moment of time.  All moments are based
// on UTC.  For moments before the UNIX Epoch the number of milliseconds is
// negative.
type moment_ int

// DISCRETE INTERFACE

// This method returns a boolean value for this discrete element.
func (v moment_) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this discrete element.
func (v moment_) AsInteger() int {
	return int(v)
}

// LEXICAL INTERFACE

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

// TEMPORAL INTERFACE

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
	return v.AsDays() / DaysPerMonth
}

// This method returns the total number of years since the UNIX Epoch
// in this moment.
func (v moment_) AsYears() float64 {
	return v.AsDays() / DaysPerYear
}

// This method returns the millisecond part of this moment.
func (v moment_) GetMilliseconds() int {
	var t = v.asTime()
	return t.Nanosecond() / 1e6
}

// This method returns the second part of this moment.
func (v moment_) GetSeconds() int {
	var t = v.asTime()
	return t.Second()
}

// This method returns the minute part of this moment.
func (v moment_) GetMinutes() int {
	var t = v.asTime()
	return t.Minute()
}

// This method returns the hour part of this moment.
func (v moment_) GetHours() int {
	var t = v.asTime()
	return t.Hour()
}

// This method returns the day part of this moment.
func (v moment_) GetDays() int {
	var t = v.asTime()
	return t.Day()
}

// This method returns the week part of this moment.
func (v moment_) GetWeeks() int {
	var t = v.asTime()
	var _, week = t.ISOWeek()
	return week
}

// This method returns the month part of this moment.
func (v moment_) GetMonths() int {
	var t = v.asTime()
	return int(t.Month())
}

// This method returns the year part of this moment.
func (v moment_) GetYears() int {
	var t = v.asTime()
	return t.Year()
}

// PRIVATE METHODS

// This function returns the go Time value for the specified UNIX milliseconds.
func (v moment_) asTime() tim.Time {
	var milliseconds int64 = int64(v)
	return tim.UnixMilli(milliseconds).UTC()
}

// MOMENT ELEMENT LIBRARY

// This singleton creates a unique name space for the library functions for
// moment elements.
var Moment = &moments_{}

// This type defines an empty structure and the group of methods bound to it
// that define the library functions for moment elements.
type moments_ struct{}

// This library function returns the duration of time between the two specified
// moments in tim.
func (l *moments_) Duration(first, second abs.MomentLike) abs.DurationLike {
	return DurationFromMilliseconds(second.AsInteger() - first.AsInteger())
}

// This library function returns the moment in time that is earlier than the
// specified moment in time by the specified duration of tim.
func (l *moments_) Earlier(moment abs.MomentLike, duration abs.DurationLike) abs.MomentLike {
	return MomentFromMilliseconds(moment.AsInteger() - duration.AsInteger())
}

// This library function returns the moment in time that is later than the
// specified moment in time by the specified duration of tim.
func (l *moments_) Later(moment abs.MomentLike, duration abs.DurationLike) abs.MomentLike {
	return MomentFromMilliseconds(moment.AsInteger() + duration.AsInteger())
}

// PRIVATE FUNCTIONS

// This function formats the specified ordinal value to the specified number of
// digits.
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
var isoFormats = [...]string{
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
// we must resort to some hacking...
func hackedParseDateAsMilliseconds(matches []string) int {

	// First, we replace the year with year zero.
	var yearString = matches[2]
	var patched = sts.Replace(matches[0], yearString, "0000", 1)

	// Next, we attempt to parse the patched moment using our Go based formats.
	for _, format := range isoFormats {
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
