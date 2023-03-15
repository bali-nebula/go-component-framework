/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	mat "math"
	cmp "math/cmplx"
	uri "net/url"
	stc "strconv"
	sts "strings"
	tim "time"
)

// UNIVERSAL CONSTRUCTORS

// This constructor returns a new angle element initialized with the specified
// value.
func Angle(value abs.Value) abs.AngleLike {
	var angle abs.AngleLike
	switch actual := value.(type) {
	case uint:
		angle = ele.AngleFromFloat(float64(actual))
	case uint8:
		angle = ele.AngleFromFloat(float64(actual))
	case uint16:
		angle = ele.AngleFromFloat(float64(actual))
	case uint32:
		angle = ele.AngleFromFloat(float64(actual))
	case uint64:
		angle = ele.AngleFromFloat(float64(actual))
	case int:
		angle = ele.AngleFromFloat(float64(actual))
	case int8:
		angle = ele.AngleFromFloat(float64(actual))
	case int16:
		angle = ele.AngleFromFloat(float64(actual))
	case int32:
		angle = ele.AngleFromFloat(float64(actual))
	case int64:
		angle = ele.AngleFromFloat(float64(actual))
	case float32:
		angle = ele.AngleFromFloat(float64(actual))
	case float64:
		angle = ele.AngleFromFloat(float64(actual))
	case string:
		angle = ParseEntity(actual).(abs.AngleLike)
	case abs.AngleLike:
		angle = actual
	case abs.ComponentLike:
		angle = actual.GetEntity().(abs.AngleLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to an angle: %v", actual, actual)
		panic(message)
	}
	return angle
}

// This constructor returns a new boolean element initialized with the specified
// value.
func Boolean(value abs.Value) abs.BooleanLike {
	var boolean abs.BooleanLike
	switch actual := value.(type) {
	case bool:
		boolean = ele.Boolean(actual)
	case string:
		boolean = ParseEntity(actual).(abs.BooleanLike)
	case abs.BooleanLike:
		boolean = actual
	case abs.ComponentLike:
		boolean = actual.GetEntity().(abs.BooleanLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a boolean: %v", actual, actual)
		panic(message)
	}
	return boolean
}

// This constructor returns a new duration element initialized with the specified
// value.
func Duration(value abs.Value) abs.DurationLike {
	var duration abs.DurationLike
	switch actual := value.(type) {
	case uint:
		duration = ele.DurationFromInt(int(actual))
	case uint8:
		duration = ele.DurationFromInt(int(actual))
	case uint16:
		duration = ele.DurationFromInt(int(actual))
	case uint32:
		duration = ele.DurationFromInt(int(actual))
	case uint64:
		duration = ele.DurationFromInt(int(actual))
	case int:
		duration = ele.DurationFromInt(int(actual))
	case int8:
		duration = ele.DurationFromInt(int(actual))
	case int16:
		duration = ele.DurationFromInt(int(actual))
	case int32:
		duration = ele.DurationFromInt(int(actual))
	case int64:
		duration = ele.DurationFromInt(int(actual))
	case string:
		duration = ParseEntity(actual).(abs.DurationLike)
	case abs.DurationLike:
		duration = actual
	case abs.ComponentLike:
		duration = actual.GetEntity().(abs.DurationLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a duration: %v", actual, actual)
		panic(message)
	}
	return duration
}

// This constructor returns a new moment element initialized with the specified
// value.
func Moment(value abs.Value) abs.MomentLike {
	var moment abs.MomentLike
	switch actual := value.(type) {
	case uint:
		moment = ele.MomentFromInt(int(actual))
	case uint8:
		moment = ele.MomentFromInt(int(actual))
	case uint16:
		moment = ele.MomentFromInt(int(actual))
	case uint32:
		moment = ele.MomentFromInt(int(actual))
	case uint64:
		moment = ele.MomentFromInt(int(actual))
	case int:
		moment = ele.MomentFromInt(int(actual))
	case int8:
		moment = ele.MomentFromInt(int(actual))
	case int16:
		moment = ele.MomentFromInt(int(actual))
	case int32:
		moment = ele.MomentFromInt(int(actual))
	case int64:
		moment = ele.MomentFromInt(int(actual))
	case string:
		moment = ParseEntity(actual).(abs.MomentLike)
	case abs.MomentLike:
		moment = actual
	case abs.ComponentLike:
		moment = actual.GetEntity().(abs.MomentLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a moment: %v", actual, actual)
		panic(message)
	}
	return moment
}

// This constructor returns a new number element initialized with the specified
// value.
func Number(value abs.Value) abs.NumberLike {
	var number abs.NumberLike
	switch actual := value.(type) {
	case uint:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case uint8:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case uint16:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case uint32:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case uint64:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int8:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int16:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int32:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case int64:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case float32:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case float64:
		number = ele.NumberFromComplex(complex(float64(actual), 0))
	case complex64:
		number = ele.NumberFromComplex(complex128(actual))
	case complex128:
		number = ele.NumberFromComplex(complex128(actual))
	case string:
		number = ParseEntity(actual).(abs.NumberLike)
	case abs.NumberLike:
		number = actual
	case abs.ComponentLike:
		number = actual.GetEntity().(abs.NumberLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a number: %v", actual, actual)
		panic(message)
	}
	return number
}

// This constructor returns a new pattern element initialized with the specified
// value.
func Pattern(value abs.Value) abs.PatternLike {
	var pattern abs.PatternLike
	switch actual := value.(type) {
	case string:
		pattern = ParseEntity(actual).(abs.PatternLike)
	case abs.PatternLike:
		pattern = actual
	case abs.ComponentLike:
		pattern = actual.GetEntity().(abs.PatternLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a pattern: %v", actual, actual)
		panic(message)
	}
	return pattern
}

// This constructor returns a new percentage element initialized with the specified
// value.
func Percentage(value abs.Value) abs.PercentageLike {
	var percentage abs.PercentageLike
	switch actual := value.(type) {
	case uint:
		percentage = ele.PercentageFromFloat(float64(actual))
	case uint8:
		percentage = ele.PercentageFromFloat(float64(actual))
	case uint16:
		percentage = ele.PercentageFromFloat(float64(actual))
	case uint32:
		percentage = ele.PercentageFromFloat(float64(actual))
	case uint64:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int8:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int16:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int32:
		percentage = ele.PercentageFromFloat(float64(actual))
	case int64:
		percentage = ele.PercentageFromFloat(float64(actual))
	case float32:
		percentage = ele.PercentageFromFloat(float64(actual))
	case float64:
		percentage = ele.PercentageFromFloat(float64(actual))
	case string:
		percentage = ParseEntity(actual).(abs.PercentageLike)
	case abs.PercentageLike:
		percentage = actual
	case abs.ComponentLike:
		percentage = actual.GetEntity().(abs.PercentageLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a percentage: %v", actual, actual)
		panic(message)
	}
	return percentage
}

// This constructor returns a new probability element initialized with the specified
// value.
func Probability(value abs.Value) abs.ProbabilityLike {
	var probability abs.ProbabilityLike
	switch actual := value.(type) {
	case bool:
		probability = ele.ProbabilityFromBool(actual)
	case uint:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case uint8:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case uint16:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case uint32:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case uint64:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int8:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int16:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int32:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case int64:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case float32:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case float64:
		probability = ele.ProbabilityFromFloat(float64(actual))
	case string:
		probability = ParseEntity(actual).(abs.ProbabilityLike)
	case abs.ProbabilityLike:
		probability = actual
	case abs.ComponentLike:
		probability = actual.GetEntity().(abs.ProbabilityLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a probability: %v", actual, actual)
		panic(message)
	}
	return probability
}

// This constructor returns a new resource element initialized with the specified
// value.
func Resource(value abs.Value) abs.ResourceLike {
	var resource abs.ResourceLike
	switch actual := value.(type) {
	case *uri.URL:
		resource = ele.ResourceFromString(actual.String())
	case string:
		resource = ParseEntity(actual).(abs.ResourceLike)
	case abs.ResourceLike:
		resource = actual
	case abs.ComponentLike:
		resource = actual.GetEntity().(abs.ResourceLike)
	default:
		var message = fmt.Sprintf("The value (of type %T) cannot be converted to a resource: %v", actual, actual)
		panic(message)
	}
	return resource
}

// PRIVATE METHODS

// This method attempts to parse an angle element. It returns the angle element
// and whether or not the angle element was successfully parsed.
func (v *parser) parseAngle() (abs.AngleLike, *Token, bool) {
	var token *Token
	var angle abs.AngleLike
	token = v.nextToken()
	if token.Type != TokenAngle {
		v.backupOne()
		return angle, token, false
	}
	var value = stringToFloat(token.Value[1:])
	if value == 2.0*mat.Pi {
		angle = ele.Tau
	} else {
		angle = ele.AngleFromFloat(value)
	}
	return angle, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatAngle(angle abs.AngleLike) {
	var s string
	switch angle {
	case ele.Pi:
		s = "~π"
	case ele.Tau:
		s = "~τ"
	default:
		var value = angle.AsReal()
		s = "~" + stc.FormatFloat(value, 'G', -1, 64)
	}
	v.AppendString(s)
}

// This method attempts to parse a boolean element. It returns the boolean
// element and whether or not the boolean element was successfully parsed.
func (v *parser) parseBoolean() (abs.BooleanLike, *Token, bool) {
	var token *Token
	var boolean abs.BooleanLike
	token = v.nextToken()
	if token.Type != TokenBoolean {
		v.backupOne()
		return boolean, token, false
	}
	var b, _ = stc.ParseBool(token.Value)
	boolean = ele.Boolean(b)
	return boolean, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatBoolean(boolean abs.BooleanLike) {
	var value = boolean.AsBoolean()
	var s = stc.FormatBool(value)
	v.AppendString(s)
}

// This method attempts to parse a duration element. It returns the duration
// element and whether or not the duration element was successfully parsed.
func (v *parser) parseDuration() (abs.DurationLike, *Token, bool) {
	var token *Token
	var duration abs.DurationLike
	token = v.nextToken()
	if token.Type != TokenDuration {
		v.backupOne()
		return duration, token, false
	}
	var matches = scanDuration([]byte(token.Value))
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
				milliseconds += float * float64(ele.MillisecondsPerWeek)
			case "Y":
				milliseconds += float * float64(ele.MillisecondsPerYear)
			case "M":
				if isTime {
					milliseconds += float * float64(ele.MillisecondsPerMinute)
				} else {
					milliseconds += float * float64(ele.MillisecondsPerMonth)
				}
			case "D":
				milliseconds += float * float64(ele.MillisecondsPerDay)
			case "T":
				isTime = true
			case "H":
				milliseconds += float * float64(ele.MillisecondsPerHour)
			case "S":
				milliseconds += float * float64(ele.MillisecondsPerSecond)
			}
		}
	}
	duration = ele.DurationFromInt(int(sign * milliseconds))
	return duration, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatDuration(duration abs.DurationLike) {
	var result sts.Builder
	result.WriteString("~")
	if duration.IsNegative() {
		result.WriteString("-")
	}
	result.WriteString("P")
	var weeks = mat.Abs(duration.AsWeeks())
	if float64(int(weeks)) == weeks {
		// It is an exact number of weeks.
		result.WriteString(stc.FormatInt(int64(weeks), 10))
		result.WriteString("W")
		v.AppendString(result.String())
		return
	}
	var years = duration.GetYears()
	if years > 0 {
		result.WriteString(stc.FormatInt(int64(years), 10))
		result.WriteString("Y")
	}
	var months = duration.GetMonths()
	if months > 0 {
		result.WriteString(stc.FormatInt(int64(months), 10))
		result.WriteString("M")
	}
	var days = duration.GetDays()
	if days > 0 {
		result.WriteString(stc.FormatInt(int64(days), 10))
		result.WriteString("D")
	}
	var hours = duration.GetHours()
	var minutes = duration.GetMinutes()
	var seconds = duration.GetSeconds()
	var milliseconds = duration.GetMilliseconds()
	if hours+minutes+seconds+milliseconds == 0 {
		// There is no time part of the duration.
		v.AppendString(result.String())
		return
	}
	result.WriteString("T")
	if hours > 0 {
		result.WriteString(stc.FormatInt(int64(hours), 10))
		result.WriteString("H")
	}
	if minutes > 0 {
		result.WriteString(stc.FormatInt(int64(minutes), 10))
		result.WriteString("M")
	}
	if seconds+milliseconds > 0 {
		result.WriteString(stc.FormatInt(int64(seconds), 10))
		if milliseconds > 0 {
			result.WriteString(".")
			result.WriteString(stc.FormatInt(int64(milliseconds), 10))
		}
		result.WriteString("S")
	}
	v.AppendString(result.String())
}

// This method attempts to parse an element primitive. It returns the
// element primitive and whether or not the element primitive was
// successfully parsed.
func (v *parser) parseElement() (abs.Element, *Token, bool) {
	var ok bool
	var token *Token
	var element abs.Element
	element, token, ok = v.parseAngle()
	if !ok {
		element, token, ok = v.parseBoolean()
	}
	if !ok {
		element, token, ok = v.parseDuration()
	}
	if !ok {
		element, token, ok = v.parseMoment()
	}
	if !ok {
		element, token, ok = v.parseNumber()
	}
	if !ok {
		element, token, ok = v.parsePattern()
	}
	if !ok {
		element, token, ok = v.parsePercentage()
	}
	if !ok {
		element, token, ok = v.parseProbability()
	}
	if !ok {
		element, token, ok = v.parseResource()
	}
	if !ok {
		// Override any zero values returned from failed parsing attempts.
		element = nil
	}
	return element, token, ok
}

// This method adds the canonical format for the specified imaginary number to
// the state of the formatter.
func (v *formatter) formatImaginary(i float64) {
	var s string
	switch i {
	case 1:
		s = "i"
	case -1:
		s = "-i"
	case mat.E:
		s = "ei"
	case -mat.E:
		s = "-ei"
	case mat.Pi:
		s = "πi"
	case -mat.Pi:
		s = "-πi"
	case mat.Phi:
		s = "φi"
	case -mat.Phi:
		s = "-φi"
	case mat.Pi * 2.0:
		s = "τi"
	case -mat.Pi * 2.0:
		s = "-τi"
	default:
		s = stc.FormatFloat(i, 'G', -1, 64) + "i"
	}
	v.AppendString(s)
}

// This method attempts to parse a moment element. It returns the moment
// element and whether or not the moment element was successfully parsed.
func (v *parser) parseMoment() (abs.MomentLike, *Token, bool) {
	var token *Token
	var moment abs.MomentLike
	token = v.nextToken()
	if token.Type != TokenMoment {
		v.backupOne()
		return moment, token, false
	}
	var date tim.Time = hackedParseDate(token.Value)
	var milliseconds = int(date.UnixMilli())
	moment = ele.MomentFromInt(milliseconds)
	return moment, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatMoment(moment abs.MomentLike) {
	var result sts.Builder
	var year = moment.GetYears()
	var month = moment.GetMonths()
	var day = moment.GetDays()
	var hour = moment.GetHours()
	var minute = moment.GetMinutes()
	var second = moment.GetSeconds()
	var millisecond = moment.GetMilliseconds()
	result.WriteString("<")
	result.WriteString(stc.FormatInt(int64(year), 10))
	if month > 1 || day > 1 || hour > 0 || minute > 0 || second > 0 || millisecond > 0 {
		result.WriteString("-")
		result.WriteString(formatOrdinal(month, 2))
		if day > 1 || hour > 0 || minute > 0 || second > 0 || millisecond > 0 {
			result.WriteString("-")
			result.WriteString(formatOrdinal(day, 2))
			if hour > 0 || minute > 0 || second > 0 || millisecond > 0 {
				result.WriteString("T")
				result.WriteString(formatOrdinal(hour, 2))
				if minute > 0 || second > 0 || millisecond > 0 {
					result.WriteString(":")
					result.WriteString(formatOrdinal(minute, 2))
					if second > 0 || millisecond > 0 {
						result.WriteString(":")
						result.WriteString(formatOrdinal(second, 2))
						if millisecond > 0 {
							result.WriteString(".")
							result.WriteString(formatOrdinal(millisecond, 3))
						}
					}
				}
			}
		}
	}
	result.WriteString(">")
	var s = result.String()
	v.AppendString(s)
}

// This method attempts to parse a number element. It returns the number
// element and whether or not the number element was successfully parsed.
func (v *parser) parseNumber() (abs.NumberLike, *Token, bool) {
	var number abs.NumberLike
	var token = v.nextToken()
	if token.Type != TokenNumber {
		v.backupOne()
		return number, token, false
	}
	var err any
	var c complex128
	var realPart float64
	var imaginaryPart float64
	var phasePart abs.AngleLike
	var matches = scanNumber([]byte(token.Value))
	switch {
	case matches[0] == "undefined":
		c = cmp.NaN()
	case matches[0] == "infinity" || matches[0] == "∞":
		c = cmp.Inf()
	case matches[0] == "i":
		c = complex(0, 1)
	case matches[0] == "-i":
		c = complex(0, -1)
	case matches[0] == "pi", matches[0] == "-pi", matches[0] == "phi", matches[0] == "-phi":
		// The value is a pure real constant ending in "i" so it must be handled first.
		realPart = stringToFloat(matches[0])
		c = complex(realPart, 0)
	case matches[0][len(matches[0])-1] == 'i':
		// The value is pure imaginary.
		var match = matches[0][:len(matches[0])-1] // Strip off the trailing "i".
		imaginaryPart, err = stc.ParseFloat(match, 64)
		if err != nil {
			// The value is a pure imaginary constant.
			imaginaryPart = stringToFloat(match)
		}
		c = complex(0, imaginaryPart)
	case matches[0][0] == '(':
		// The value is complex.
		switch {
		case matches[2] == ", ":
			// The complex number is in rectangular form.
			realPart, err = stc.ParseFloat(matches[1], 64)
			if err != nil {
				// The real part of the number is a constant.
				realPart = stringToFloat(matches[1])
			}
			if matches[3] == "i" {
				imaginaryPart = 1
			} else if matches[3] == "-i" {
				imaginaryPart = -1
			} else {
				var match = matches[3][:len(matches[3])-1] // Strip off the trailing "i".
				imaginaryPart, err = stc.ParseFloat(match, 64)
				if err != nil {
					// The imaginary part of the number is a constant.
					imaginaryPart = stringToFloat(match)
				}
			}
			c = complex(realPart, imaginaryPart)
		default:
			// The complex number is in polar form.
			realPart, err = stc.ParseFloat(matches[4], 64)
			if err != nil {
				// The real part of the number is a constant.
				realPart = stringToFloat(matches[4])
			}
			var match = matches[6][:len(matches[6])-1] // Strip off the trailing "i".
			var parser = Parser([]byte(match))
			phasePart, _, _ = parser.parseAngle()
			c = ele.NumberFromPolar(realPart, phasePart).AsComplex()
		}
	default:
		// The value is pure real.
		realPart, err = stc.ParseFloat(matches[0], 64)
		if err != nil {
			// The value is a pure real constant.
			realPart = stringToFloat(matches[0])
		}
		c = complex(realPart, 0)
	}
	number = ele.NumberFromComplex(c)
	return number, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatNumber(number abs.NumberLike) {
	switch {
	case number.IsZero():
		v.AppendString("0")
	case number.IsInfinite():
		v.AppendString("∞")
	case number.IsUndefined():
		v.AppendString("undefined")
	default:
		var realPart = number.GetReal()
		var imagPart = number.GetImaginary()
		switch {
		case imagPart == 0:
			v.formatFloat(realPart)
		case realPart == 0:
			v.formatImaginary(imagPart)
		default:
			v.AppendString("(")
			v.formatFloat(realPart)
			v.AppendString(", ")
			v.formatImaginary(imagPart)
			v.AppendString(")")
		}
	}
}

// This method adds the canonical format for the specified imaginary phase to
// the state of the formatter.
func (v *formatter) formatPhase(p abs.AngleLike) {
	var s string
	if p.AsReal() == mat.Pi {
		s = "~πi"
	} else {
		var value = p.AsReal()
		s = "~" + stc.FormatFloat(value, 'G', -1, 64) + "i"
	}
	v.AppendString(s)
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatPolar(number abs.NumberLike) {
	switch {
	case number.IsZero():
		v.AppendString("0")
	case number.IsInfinite():
		v.AppendString("∞")
	case number.IsUndefined():
		v.AppendString("undefined")
	default:
		var magnitude = number.GetMagnitude()
		var phase = number.GetPhase()
		switch {
		case phase.IsZero():
			v.formatFloat(magnitude)
		case magnitude == 0:
			v.AppendString("0")
		default:
			v.AppendString("(")
			v.formatFloat(magnitude)
			v.AppendString("e^")
			v.formatPhase(phase)
			v.AppendString(")")
		}
	}
}

// This method attempts to parse a pattern element. It returns the pattern
// element and whether or not the pattern element was successfully parsed.
func (v *parser) parsePattern() (abs.PatternLike, *Token, bool) {
	var token *Token
	var pattern abs.PatternLike
	token = v.nextToken()
	if token.Type != TokenPattern {
		v.backupOne()
		return pattern, token, false
	}
	var regex = token.Value
	switch regex {
	case `none`:
		regex = `^none$`
	case `any`:
		regex = `.*`
	default:
		regex = regex[1 : len(regex)-2] // Strip off the '"' and '"?' delimiters.
	}
	pattern = ele.PatternFromString(regex)
	return pattern, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatPattern(pattern abs.PatternLike) {
	var regexp = pattern.AsString()
	var s = `"` + regexp + `"?`
	switch regexp {
	case `^none$`:
		s = `none`
	case `.*`:
		s = `any`
	}
	v.AppendString(s)
}

// This method attempts to parse a percentage element. It returns the percentage
// element and whether or not the percentage element was successfully parsed.
func (v *parser) parsePercentage() (abs.PercentageLike, *Token, bool) {
	var token *Token
	var percentage abs.PercentageLike
	token = v.nextToken()
	if token.Type != TokenPercentage {
		v.backupOne()
		return percentage, token, false
	}
	var s = token.Value[:len(token.Value)-1] // Removed the trailing '%'.
	var float, _ = stc.ParseFloat(s, 64)
	percentage = ele.PercentageFromFloat(float)
	return percentage, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatPercentage(percentage abs.PercentageLike) {
	var float float64
	var element, ok = percentage.(ele.Percentage)
	if ok {
		float = float64(element)
	} else {
		// This approach introduces round-off errors.
		float = 100.0 * percentage.AsReal()
	}
	var s = stc.FormatFloat(float, 'G', -1, 64) + "%"
	v.AppendString(s)
}

// This method attempts to parse a probability element. It returns the probability
// element and whether or not the probability element was successfully parsed.
func (v *parser) parseProbability() (abs.ProbabilityLike, *Token, bool) {
	var token *Token
	var probability abs.ProbabilityLike
	token = v.nextToken()
	if token.Type != TokenProbability {
		v.backupOne()
		return probability, token, false
	}
	var matches = scanProbability([]byte(token.Value))
	var float, _ = stc.ParseFloat(matches[0], 64)
	probability = ele.ProbabilityFromFloat(float)
	return probability, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatProbability(probability abs.ProbabilityLike) {
	var value = probability.AsReal()
	var s = stc.FormatFloat(value, 'f', -1, 64)
	switch {
	// Zero is formatted as ".0".
	case value == 0:
		s = "." + s
	// One is formatted as "1.".
	case value == 1:
		s += "."
	// Otherwise, remove the leading "0".
	default:
		s = s[1:]
	}
	v.AppendString(s)
}

// This method adds the canonical format for the specified real number to the
// state of the formatter.
func (v *formatter) formatFloat(r float64) {
	var s string
	switch r {
	case mat.E:
		s = "e"
	case -mat.E:
		s = "-e"
	case mat.Pi:
		s = "π"
	case -mat.Pi:
		s = "-π"
	case mat.Phi:
		s = "φ"
	case -mat.Phi:
		s = "-φ"
	case mat.Pi * 2.0:
		s = "τ"
	case -mat.Pi * 2.0:
		s = "-τ"
	default:
		s = stc.FormatFloat(r, 'G', -1, 64)
	}
	v.AppendString(s)
}

// This method attempts to parse a resource element. It returns the
// resource element and whether or not the resource element was
// successfully parsed.
func (v *parser) parseResource() (abs.ResourceLike, *Token, bool) {
	var token *Token
	var resource abs.ResourceLike
	token = v.nextToken()
	if token.Type != TokenResource {
		v.backupOne()
		return resource, token, false
	}
	var matches = scanResource([]byte(token.Value))
	resource = ele.ResourceFromString(matches[1])
	return resource, token, true
}

// This method adds the canonical format for the specified element to the state
// of the formatter.
func (v *formatter) formatResource(resource abs.ResourceLike) {
	var uri = resource.AsString()
	var s = "<" + uri + ">"
	v.AppendString(s)
}

// PRIVATE FUNCTIONS

// This function formats the specified ordinal value to the specified number of
// digits.
func formatOrdinal(ordinal, digits int) string {
	return fmt.Sprintf("%0"+stc.Itoa(digits)+"d", ordinal)
}

// This function converts a string into a float value.
func stringToFloat(string_ string) float64 {
	switch string_ {
	case "":
		return 1
	case "-":
		return -1
	case "e":
		return mat.E
	case "-e":
		return -mat.E
	case "pi", "π":
		return mat.Pi
	case "-pi", "-π":
		return -mat.Pi
	case "phi", "φ":
		return mat.Phi
	case "-phi", "-φ":
		return -mat.Phi
	case "tau", "τ":
		return mat.Pi * 2.0
	case "-tau", "-τ":
		return -mat.Pi * 2.0
	}
	var float, _ = stc.ParseFloat(string_, 64)
	return float
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
func hackedParseDate(moment string) tim.Time {

	// First, we replace the year with year zero.
	var matches = scanMoment([]byte(moment))
	var yearString = matches[2]
	var patched = sts.Replace(moment, yearString, "0000", 1)

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

			// And return the correct date.
			return date
		}
	}

	// This panic will only happen if the scanner regular expressions are out
	// of sync with the ISO 8601 standard formats. The moment has already been
	// succussfully scanned by the the scanner.
	panic(fmt.Sprintf("The moment does not match a known format: %v", moment))
}
