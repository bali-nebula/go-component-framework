/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

import (
	"bytes"
	"regexp"
)

// LEXICAL INTERFACES

// This interface defines the methods supported by all lexical elements.
type Lexical interface {
	AsString() string
}

// ANGLE ELEMENT SYNTAX

// These constants are used to form a regular expression for valid angle
// entities. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Angle
const (
	e        = `e`
	pi       = `pi|π`
	phi      = `phi|φ`
	tau      = `tau|τ`
	sign     = `[+-]`
	zero     = `0`
	ordinal  = `[1-9][0-9]*`
	fraction = `\.[0-9]+`
	exponent = `E` + sign + `?` + ordinal
	scalar   = `(?:` + ordinal + `(?:` + fraction + `)?|` + zero + fraction + `)(?:` + exponent + `)?`
	real     = sign + `?(?:` + e + `|` + pi + `|` + phi + `|` + tau + `|` + scalar + `)`
	angle    = `~(` + real + `|` + zero + `)`
)

// This scanner is used for matching angle entities.
var angleScanner = regexp.MustCompile(`^(?:` + angle + `)`)

// This function returns for the specified string an array of the matching
// subgroups for an angle element. The first item in the array is the entire
// matched string.
func ScanAngle(v []byte) []string {
	return bytesToStrings(angleScanner.FindSubmatch(v))
}

// BINARY STRING SYNTAX

// These constants are used to form a regular expression for valid binary
// strings. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Binary
const (
	whitespace = `\pS` // All unicode whitespace characters.
	base64     = `[A-Za-z0-9+/]`
	binary     = `'((?:` + base64 + `|` + whitespace + `)*)'`
)

// This scanner is used for matching binary strings.
var binaryScanner = regexp.MustCompile(`^(?:` + binary + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a binary string. The first item in the array is the entire
// matched string.
func ScanBinary(v []byte) []string {
	return bytesToStrings(binaryScanner.FindSubmatch(v))
}

// BOOLEAN ELEMENT SYNTAX

// These constants are used to form a regular expression for valid boolean
// entities. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Boolean
const (
	boolean = `false|true`
)

// This scanner is used for matching boolean entities.
var booleanScanner = regexp.MustCompile(`^(?:` + boolean + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a boolean element. The first item in the array is the entire
// matched string.
func ScanBoolean(v []byte) []string {
	return bytesToStrings(booleanScanner.FindSubmatch(v))
}

// COMMENT SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a commment. The first item in the array is the entire matched
// string. Since a comment can be recursive a regular expression is not used in
// this implementation.
func ScanComment(v []byte) []string {
	var result []string
	if !bytes.HasPrefix(v, bangAngle) {
		return result
	}
	var angleBangAllowed = false
	var current = 3 // Skip the leading '!>\n' characters.
	var level = 1
	for level > 0 {
		var s = v[current:]
		switch {
		case len(s) == 0:
			return result
		case bytes.HasPrefix(s, eol):
			angleBangAllowed = true
			current++
		case bytes.HasPrefix(s, bangAngle):
			current += 3 // Skip the '!>\n' characters.
			level++      // Start a nested narrative.
		case bytes.HasPrefix(s, angleBang):
			if !angleBangAllowed {
				return result
			}
			current += 2 // Skip the '<!' characters.
			level--      // Terminate the current narrative.
		default:
			if angleBangAllowed && !bytes.HasPrefix(s, space) {
				angleBangAllowed = false
			}
			current++ // Accept the next character.
		}
	}
	result = append(result, string(v[:current])) // Includes bang delimeters.
	return result
}

// DELIMITER SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a delimiter. The first item in the array is the entire
// matched string.
func ScanDelimiter(v []byte) []string {
	var result []string
	for _, delimiter := range delimiters {
		if bytes.HasPrefix(v, delimiter) {
			result = append(result, string(delimiter))
		}
	}
	return result
}

// DURATION ELEMENT SYNTAX

// These constants are used to form a regular expression for valid duration
// entities. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Duration
const (
	tspan    = `(?:` + zero + `|` + ordinal + `(?:` + fraction + `)?)`
	weeks    = `(` + tspan + `W)`
	dates    = `(` + tspan + `Y)?(` + tspan + `M)?(` + tspan + `D)?`
	times    = `(T)(` + tspan + `H)?(` + tspan + `M)?(` + tspan + `S)?`
	duration = `~(` + sign + `?)P(?:` + weeks + `|` + dates + `(?:` + times + `)?)`
)

// This scanner is used for matching duration entities.
var durationScanner = regexp.MustCompile(`^(?:` + duration + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a duration element. The first item in the array is the entire
// matched string.
func ScanDuration(v []byte) []string {
	return bytesToStrings(durationScanner.FindSubmatch(v))
}

// IDENTIFIER SYNTAX

// These constants are used to form a regular expression for valid identifiers.
// See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Identifier
const (
	letter     = `\pL` // All unicode letters and connectors like underscores.
	digit      = `\pN` // All unicode digits.
	identifier = letter + `(?:` + letter + `|` + digit + `)*`
)

// This scanner is used for matching identifiers.
var identifierScanner = regexp.MustCompile(`^(?:` + identifier + `)`)

// This function returns for the specified string an array of the matching
// subgroups for an identifier. The first item in the array is the entire
// matched string.
func ScanIdentifier(v []byte) []string {
	return bytesToStrings(identifierScanner.FindSubmatch(v))
}

// KEYWORD SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a keyword. The first item in the array is the entire
// matched string.
func ScanKeyword(v []byte) []string {
	var result []string
	for _, keyword := range keywords {
		if bytes.HasPrefix(v, keyword) {
			result = append(result, string(keyword))
		}
	}
	return result
}

// MOMENT ELEMENT SYNTAX

// These constants are used to form a regular expression for valid moment
// entities. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Moment
const (
	year   = sign + `?` + ordinal
	month  = `(?:[0][1-9])|(?:[1][012])`
	day    = `(?:[012][1-9])|(?:[3][01])`
	hour   = `(?:[01][0-9])|(?:[2][0-3])`
	minute = `[0-5][0-9]`
	second = `(?:[0-5][0-9])|(?:[6][01])`
	moment = `<(` + year + `)(?:-(` + month + `)(?:-(` + day + `)` +
		`(?:T(` + hour + `)(?::(` + minute + `)(?::((?:` + second + `)(?:` + fraction + `)?))?)?)?)?)?>`
)

// This scanner is used for matching moment entities.
var momentScanner = regexp.MustCompile(`^(?:` + moment + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a moment element. The first item in the array is the entire
// matched string.
func ScanMoment(v []byte) []string {
	return bytesToStrings(momentScanner.FindSubmatch(v))
}

// MONIKER STRING SYNTAX

// These constants are used to form a regular expression for valid moniker
// strings. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Moniker
const (
	separator = `[-+.]`
	name      = letter + `(?:` + separator + `?(?:` + letter + `|` + digit + `))*`
	moniker   = `(?:/` + name + `)+` // Cannot capture each name...
)

// This scanner is used for matching moniker strings.
var monikerScanner = regexp.MustCompile(`^(?:` + moniker + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a moniker string. The first item in the array is the entire
// matched string.
func ScanMoniker(v []byte) []string {
	return bytesToStrings(monikerScanner.FindSubmatch(v))
}

// NARRATIVE STRING SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a narrative string. The first item in the array is the entire
// matched string. Since a narrative can be recursive a regular expression is
// not used in this implementation.
func ScanNarrative(v []byte) []string {
	var result []string
	if !bytes.HasPrefix(v, quoteAngle) {
		return result
	}
	var angleQuoteAllowed = false
	var current = 3 // Skip the leading '">\n' characters.
	var level = 1
	for level > 0 {
		var s = v[current:]
		switch {
		case len(s) == 0:
			return result
		case bytes.HasPrefix(s, eol):
			angleQuoteAllowed = true
			current++
		case bytes.HasPrefix(s, quoteAngle):
			current += 3 // Skip the '">\n' characters.
			level++      // Start a nested narrative.
		case bytes.HasPrefix(s, angleQuote):
			if !angleQuoteAllowed {
				return result
			}
			current += 2 // Skip the '<"' characters.
			level--      // Terminate the current narrative.
		default:
			if angleQuoteAllowed && !bytes.HasPrefix(s, space) {
				angleQuoteAllowed = false
			}
			current++ // Accept the next character.
		}
	}
	result = append(result, string(v[:current])) // Includes quote delimeters.
	return result
}

// NOTE SYNTAX

// These constants are used to form a regular expression for valid notes. See
// the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Note
const (
	note = `! [^\n]*`
)

// This scanner is used for matching notes.
var noteScanner = regexp.MustCompile(`^(?:` + note + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a note. The first item in the array is the entire matched
// string.
func ScanNote(v []byte) []string {
	return bytesToStrings(noteScanner.FindSubmatch(v))
}

// NUMBER ELEMENT SYNTAX

// These constants are used to form a regular expression for valid number
// entities. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Number
const (
	infinity    = `infinity|∞`
	undefined   = `undefined`
	imaginary   = `(?:` + sign + `|` + real + `)?i`
	rectangular = `(` + real + `)(, )(` + imaginary + `)`
	polar       = `(` + real + `)(e\^)(` + angle + `i)`
	number      = imaginary + `|` + real + `|` + zero + `|` + infinity + `|` + undefined + `|` +
		`\((?:` + rectangular + `|` + polar + `)\)`
)

// This scanner is used for matching number entities.
var numberScanner = regexp.MustCompile(`^(?:` + number + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a number element. The first item in the array is the entire
// matched string.
func ScanNumber(v []byte) []string {
	return bytesToStrings(numberScanner.FindSubmatch(v))
}

// PATTERN ELEMENT SYNTAX

// These constants are used to form a regular expression for valid pattern
// entities. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Pattern
const (
	base16  = `[0-9a-f]`
	unicode = `u` + base16 + `{4}`
	escape  = `\\(?:` + unicode + `|[bfrnt\\])`
	roon    = `(?:` + escape + `|\\"|[^"\r\n]` + `)`
	regex   = `"(` + roon + `+)"\?`
	pattern = `none` + `|` + regex + `|` + `any`
)

// This scanner is used for matching pattern entities.
var patternScanner = regexp.MustCompile(`^(?:` + pattern + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a pattern element. The first item in the array is the entire
// matched string.
func ScanPattern(v []byte) []string {
	return bytesToStrings(patternScanner.FindSubmatch(v))
}

// PERCENTAGE ELEMENT SYNTAX

// These constants are used to form a regular expression for valid percentage
// entities. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Percentage
const (
	percentage = `(` + real + `|` + zero + `)%`
)

// This scanner is used for matching percentage entities.
var percentageScanner = regexp.MustCompile(`^(?:` + percentage + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a percentage element. The first item in the array is the entire
// matched string.
func ScanPercentage(v []byte) []string {
	return bytesToStrings(percentageScanner.FindSubmatch(v))
}

// PROBABILITY ELEMENT SYNTAX

// These constants are used to form a regular expression for valid probability
// entities. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Probability
const (
	probability = fraction + `|1\.`
)

// This scanner is used for matching probability entities.
var probabilityScanner = regexp.MustCompile(`^(?:` + probability + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a probability element. The first item in the array is the
// entire matched string.
func ScanProbability(v []byte) []string {
	return bytesToStrings(probabilityScanner.FindSubmatch(v))
}

// QUOTE STRING SYNTAX

// These constants are used to form a regular expression for valid quoted
// strings. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Quote
const (
	quote = `"(` + roon + `*)"`
)

// This scanner is used for matching quoted strings.
var quoteScanner = regexp.MustCompile(`^(?:` + quote + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a quoted string. The first item in the array is the entire
// matched string.
func ScanQuote(v []byte) []string {
	return bytesToStrings(quoteScanner.FindSubmatch(v))
}

// RESOURCE ELEMENT SYNTAX

// These constants are used to form a regular expression for valid resource
// entities. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Resource
const (
	scheme    = `[a-zA-Z][0-9a-zA-Z+-.]*`
	authority = `[^/]+`
	path      = `[^?#>]*`
	query     = `[^#>]*`
	fragment  = `[^>]+`
	resource  = `<(` +
		`(` + scheme + `):` +
		`(?:` + `//(` + authority + `)` + `)?` +
		`(` + path + `)` +
		`(?:` + `\?(` + query + `)` + `)?` +
		`(?:` + `#(` + fragment + `)` + `)?` +
		`)>`
)

// This scanner is used for matching resource entities.
var resourceScanner = regexp.MustCompile(`^(?:` + resource + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a resource element. The first item in the array is the entire
// matched string.
func ScanResource(v []byte) []string {
	return bytesToStrings(resourceScanner.FindSubmatch(v))
}

// SYMBOL STRING SYNTAX

// These constants are used to form a regular expression for valid symbol
// strings. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Symbol
const (
	symbol = `\$(` + identifier + `(?:-` + ordinal + `)?)`
)

// This scanner is used for matching symbol strings.
var symbolScanner = regexp.MustCompile(`^(?:` + symbol + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a symbol string. The first item in the array is the entire
// matched string.
func ScanSymbol(v []byte) []string {
	return bytesToStrings(symbolScanner.FindSubmatch(v))
}

// TAG ELEMENT SYNTAX

// These constants are used to form a regular expression for valid tag entities.
// See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Tag
const (
	base32 = `[0-9A-DF-HJ-NP-TV-Z]` // "E", "I", "O", and "U" have been removed.
	tag    = `#(` + base32 + `+)`
)

// This scanner is used for matching tag entities.
var tagScanner = regexp.MustCompile(`^(?:` + tag + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a tag element. The first item in the array is the entire
// matched string.
func ScanTag(v []byte) []string {
	return bytesToStrings(tagScanner.FindSubmatch(v))
}

// VERSION STRING SYNTAX

// These constants are used to form a regular expression for valid version
// strings. See the language specification for the exact grammar:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification#Version
const (
	version = `v(` + ordinal + `(?:\.` + ordinal + `)*)` // Cannot capture each ordinal...
)

// This scanner is used for matching version strings.
var versionScanner = regexp.MustCompile(`^(?:` + version + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a version string. The first item in the array is the entire
// matched string.
func ScanVersion(v []byte) []string {
	return bytesToStrings(versionScanner.FindSubmatch(v))
}

// PRIVATE CONSTANTS

// The following constants define some important byte sequences.
var (
	eol         = []byte("\n")
	space       = []byte(" ")
	doubleQuote = []byte(`"`)
	quoteAngle  = []byte(`">` + "\n")
	angleQuote  = []byte(`<"`)
	bangAngle   = []byte("!>" + "\n")
	angleBang   = []byte("<!")
)

// This array contains the full set of keywords used by the Bali Document
// Notation™ (BDN). They are in reverse order for proper matching.
var keywords = [][]byte{
	[]byte("with"),
	[]byte("while"),
	[]byte("to"),
	[]byte("throw"),
	[]byte("select"),
	[]byte("save"),
	[]byte("return"),
	[]byte("retrieve"),
	[]byte("reject"),
	[]byte("publish"),
	[]byte("post"),
	[]byte("notarize"),
	[]byte("matching"),
	[]byte("loop"),
	[]byte("level"),
	[]byte("in"),
	[]byte("if"),
	[]byte("handle"),
	[]byte("from"),
	[]byte("each"),
	[]byte("do"),
	[]byte("discard"),
	[]byte("continue"),
	[]byte("checkout"),
	[]byte("break"),
	[]byte("at"),
	[]byte("as"),
	[]byte("any"),
	[]byte("accept"),
	[]byte("XOR"),
	[]byte("SANS"),
	[]byte("OR"),
	[]byte("NOT"),
	[]byte("MATCHES"),
	[]byte("IS"),
	[]byte("AND")}

// This array contains the full set of delimiters used by the Bali Document
// Notation™ (BDN). They are in reverse order for proper matching.
var delimiters = [][]byte{
	[]byte("}"),
	[]byte("|"),
	[]byte("{"),
	[]byte("^"),
	[]byte("]"),
	[]byte("["),
	[]byte("@"),
	[]byte("?"),
	[]byte(">"),
	[]byte("="),
	[]byte("<..<"),
	[]byte("<.."),
	[]byte("<~"),
	[]byte("<"),
	[]byte(";"),
	[]byte(":="),
	[]byte(":"),
	[]byte("/="),
	[]byte("//"),
	[]byte("/"),
	[]byte("..<"),
	[]byte(".."),
	[]byte("."),
	[]byte("-="),
	[]byte("-"),
	[]byte(","),
	[]byte("+="),
	[]byte("+"),
	[]byte("*="),
	[]byte("*"),
	[]byte(")"),
	[]byte("("),
	[]byte("&")}

// PRIVATE FUNCTIONS

func bytesToStrings(bytes [][]byte) []string {
	var strings = make([]string, len(bytes))
	for index, array := range bytes {
		strings[index] = string(array)
	}
	return strings
}
