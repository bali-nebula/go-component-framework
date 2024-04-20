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
	fmt "fmt"
	mat "math"
	stc "strconv"
	sts "strings"
	tim "time"
)

// CLASS ACCESS

// Reference

var momentClass = &momentClass_{
	minimumValue_: moment_(mat.MinInt),
	maximumValue_: moment_(mat.MaxInt),
	epoch_:        moment_(0),
}

// Function

func Moment() MomentClassLike {
	return momentClass
}

// CLASS METHODS

// Target

type momentClass_ struct {
	minimumValue_ moment_
	maximumValue_ moment_
	epoch_        moment_
}

// Constants

func (c *momentClass_) MinimumValue() MomentLike {
	return c.minimumValue_
}

func (c *momentClass_) MaximumValue() MomentLike {
	return c.maximumValue_
}

func (c *momentClass_) Epoch() MomentLike {
	return c.epoch_
}

// Constructors

func (c *momentClass_) Make() MomentLike {
	var now = tim.Now().UTC().UnixMilli()
	return moment_(now)
}

func (c *momentClass_) MakeFromMilliseconds(milliseconds int64) MomentLike {
	return moment_(milliseconds)
}

func (c *momentClass_) MakeFromString(string_ string) MomentLike {
	var matches = matchMoment(string_)
	var milliseconds = momentFromMatches(matches)
	return moment_(milliseconds)
}

// Functions

func (c *momentClass_) Duration(
	first MomentLike,
	second MomentLike,
) DurationLike {
	return duration_(second.AsInteger() - first.AsInteger())
}

func (c *momentClass_) Earlier(
	moment MomentLike,
	duration DurationLike,
) MomentLike {
	return moment_(moment.AsInteger() - duration.AsInteger())
}

func (c *momentClass_) Later(
	moment MomentLike,
	duration DurationLike,
) MomentLike {
	return moment_(moment.AsInteger() + duration.AsInteger())
}

// INSTANCE METHODS

// Target

type moment_ int64

// Discrete

func (v moment_) AsBoolean() bool {
	return true
}

func (v moment_) AsInteger() int64 {
	return int64(v)
}

// Lexical

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
	builder.WriteString(stc.FormatInt(year, 10))
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

// Temporal

func (v moment_) AsMilliseconds() float64 {
	return float64(v)
}

func (v moment_) AsSeconds() float64 {
	return float64(v) / float64(durationClass.millisecondsPerSecond_)
}

func (v moment_) AsMinutes() float64 {
	return float64(v) / float64(durationClass.millisecondsPerMinute_)
}

func (v moment_) AsHours() float64 {
	return float64(v) / float64(durationClass.millisecondsPerHour_)
}

func (v moment_) AsDays() float64 {
	return float64(v) / float64(durationClass.millisecondsPerDay_)
}

func (v moment_) AsWeeks() float64 {
	return float64(v) / float64(durationClass.millisecondsPerWeek_)
}

func (v moment_) AsMonths() float64 {
	return float64(v) / float64(durationClass.millisecondsPerMonth_)
}

func (v moment_) AsYears() float64 {
	return float64(v) / float64(durationClass.millisecondsPerYear_)
}

// Factored

func (v moment_) GetMilliseconds() int64 {
	var time = v.asTime()
	var milliseconds = time.Nanosecond() / 1e6
	return int64(milliseconds)
}

func (v moment_) GetSeconds() int64 {
	var time = v.asTime()
	var seconds = time.Second()
	return int64(seconds)
}

func (v moment_) GetMinutes() int64 {
	var time = v.asTime()
	var minutes = time.Minute()
	return int64(minutes)
}

func (v moment_) GetHours() int64 {
	var time = v.asTime()
	var hours = time.Hour()
	return int64(hours)
}

func (v moment_) GetDays() int64 {
	var time = v.asTime()
	var days = time.Day()
	return int64(days)
}

func (v moment_) GetWeeks() int64 {
	var time = v.asTime()
	var _, weeks = time.ISOWeek()
	return int64(weeks)
}

func (v moment_) GetMonths() int64 {
	var time = v.asTime()
	var months = time.Month()
	return int64(months)
}

func (v moment_) GetYears() int64 {
	var time = v.asTime()
	var years = time.Year()
	return int64(years)
}

// Private

func (v moment_) asTime() tim.Time {
	return tim.UnixMilli(int64(v)).UTC()
}

// PACKAGE FUNCTIONS

// Private

func formatOrdinal(ordinal int64, digits int) string {
	return fmt.Sprintf("%0"+stc.Itoa(digits)+"d", ordinal)
}

func matchMoment(string_ string) []string {
	var matches = []string{
		string_,
	}
	// TBA - Add the pattern matching code...
	return matches
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
// So we must resort to some hacking with this private function...
func momentFromMatches(matches []string) int {
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
