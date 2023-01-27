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
	byt "bytes"
	fmt "fmt"
	reg "regexp"
	sts "strings"
	uni "unicode"
	utf "unicode/utf8"
)

// TOKENS

// This integer type is used as a type identifier for each token.
type TokenType int

// This enumeration defines all possible token types including the error token.
const (
	// The first two token types must be first.
	TokenError TokenType = iota
	TokenEOF
	TokenEOL
	TokenAngle
	TokenBinary
	TokenBoolean
	TokenComment
	TokenDelimiter
	TokenDuration
	TokenIdentifier
	TokenKeyword
	TokenMoment
	TokenMoniker
	TokenNarrative
	TokenNote
	TokenNumber
	TokenPattern
	TokenPercentage
	TokenProbability
	TokenQuote
	TokenResource
	TokenSymbol
	TokenTag
	TokenVersion
)

// This method returns the string representation for each token type.
func (v TokenType) String() string {
	return [...]string{
		"Error",
		"EOF",
		"EOL",
		"Angle",
		"Binary",
		"Boolean",
		"Comment",
		"Delimiter",
		"Duration",
		"Identifier",
		"Keyword",
		"Moment",
		"Moniker",
		"Narrative",
		"Note",
		"Number",
		"Pattern",
		"Percentage",
		"Probability",
		"Quote",
		"Resource",
		"Symbol",
		"Tag",
		"Version",
	}[v]
}

// This type defines the structure and methods for each token returned by the
// scanner.
type Token struct {
	Type     TokenType
	Value    string
	Line     int // The line number of the token in the input string.
	Position int // The position in the line of the first rune of the token.
}

// This method returns the a canonical string version of this token.
func (v Token) String() string {
	var s string
	switch {
	case v.Type == TokenEOF:
		s = "<EOF>"
	case v.Type == TokenEOL:
		s = "<EOL>"
	case len(v.Value) > 60:
		s = fmt.Sprintf("%.60q...", v.Value)
	default:
		s = fmt.Sprintf("%q", v.Value)
	}
	return fmt.Sprintf("Token [type: %s, line: %d, position: %d]: %s", v.Type, v.Line, v.Position, s)
}

// SCANNER

// This constructor creates a new scanner initialized with the specified array
// of bytes. The scanner will scan in tokens matching Bali Document Notation™ as
// defined in the language specification:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification
//
// All token types in the specification are shown in UPPERCASE.
func Scanner(source []byte, tokens chan Token) *scanner {
	var v = &scanner{source: source, line: 1, position: 1, tokens: tokens}
	go v.scanTokens() // Start scanning in the background.
	return v
}

// This type defines the structure and methods for the scanner agent. The source
// bytes can be viewed like this:
//
//   | byte 0 | byte 1 | byte 2 | byte 3 | byte 4 | byte 5 | ... | byte N-1 |
//   | rune 0 |      rune 1     |      rune 2     | rune 3 | ... | rune R-1 |
//
// Runes can be one to eight bytes long.

type scanner struct {
	source    []byte
	firstByte int // The zero based index of the first possible byte in the next token.
	nextByte  int // The zero based index of the next possible byte in the next token.
	line      int // The line number in the source string of the next rune.
	position  int // The position in the current line of the first rune in the next token.
	tokens    chan Token
}

// This method continues scanning tokens from the source array until an error
// occurs or the end of file is reached. It then closes the token channel.
func (v *scanner) scanTokens() {
	for v.scanToken() {
	}
	close(v.tokens)
}

// This method attempts to scan any token starting with the next rune in
// the source array. It returns true if it found a valid token or false
// if no valid token or a TokenEOF was found.
func (v *scanner) scanToken() bool {
	v.skipSpaces()
	switch {
	// The order of the cases below has been set to help optimize the
	// performance by placing the most common token types first.
	//
	// Must be first in case the input buffer is empty.
	case v.foundEOL():
	case v.foundSymbol():
	case v.foundBoolean():
	case v.foundPattern():
	case v.foundPercentage():
	case v.foundProbability():
	case v.foundNumber(): // Must be after foundPercentage() and foundProbability().
	case v.foundVersion():
	case v.foundKeyword():
	case v.foundIdentifier(): // Must be after foundBoolean(), foundKeyword(), foundPattern() and foundVersion().
	case v.foundNote():
	case v.foundComment():
	case v.foundNarrative():
	case v.foundQuote(): // Must be after foundPattern() and foundNarrative().
	case v.foundTag():
	case v.foundMoniker():
	case v.foundBinary():
	case v.foundResource():
	case v.foundMoment():
	case v.foundDuration():
	case v.foundAngle():
	case v.foundDelimiter(): // Must be after all element and string types.
	case v.foundEOF():
		// We are done scanning.
		return false
	default:
		v.foundError()
		return false
	}
	return true
}

// This method scans through any spaces in the source array and
// sets the next byte index to the next non-white-space rune.
func (v *scanner) skipSpaces() {
	if v.nextByte < len(v.source) {
		for {
			if v.source[v.nextByte] != ' ' {
				break
			}
			v.nextByte++
			v.position++
		}
		v.firstByte = v.nextByte
	}
}

// This method adds a token of the specified type with the current scanner
// information to the token channel. It then resets the first byte index to the
// next byte index position. It returns the token type of the type added to the
// channel.
func (v *scanner) emitToken(tType TokenType) TokenType {
	var tValue = string(v.source[v.firstByte:v.nextByte])
	if tType == TokenEOF {
		tValue = "<EOF>"
	}
	if tType == TokenError {
		switch tValue {
		case "\a":
			tValue = "<BELL>"
		case "\b":
			tValue = "<BKSP>"
		case "\t":
			tValue = "<TAB>"
		case "\f":
			tValue = "<FF>"
		case "\r":
			tValue = "<CR>"
		case "\v":
			tValue = "<VTAB>"
		}
	}
	var token = Token{tType, tValue, v.line, v.position}
	//fmt.Println(token)
	v.tokens <- token
	v.firstByte = v.nextByte
	v.position += sts.Count(tValue, "") - 1 // Add the number of runes in the token.
	return tType
}

// This method adds an angle element token with the current scanner information to
// the token channel. It returns true if an angle token was found.
func (v *scanner) foundAngle() bool {
	var s = v.source[v.nextByte:]
	var matches = scanAngle(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenAngle)
		return true
	}
	return false
}

// This method adds a boolean element token with the current scanner information
// to the token channel. It returns true if a boolean token was found.
func (v *scanner) foundBoolean() bool {
	var s = v.source[v.nextByte:]
	var matches = scanBoolean(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenBoolean)
		return true
	}
	return false
}

// This method adds a binary string token with the current scanner information to
// the token channel. It returns true if a binary token was found.
func (v *scanner) foundBinary() bool {
	var s = v.source[v.nextByte:]
	var matches = scanBinary(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.line += sts.Count(matches[0], "\n")
		v.emitToken(TokenBinary)
		return true
	}
	return false
}

// This method adds a comment token with the current scanner information to the
// token channel. It returns true if a comment token was found.
func (v *scanner) foundComment() bool {
	var s = v.source[v.nextByte:]
	var matches = scanComment(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.line += sts.Count(matches[0], "\n")
		v.emitToken(TokenComment)
		return true
	}
	return false
}

// This method adds a delimiter token with the current scanner information to the
// token channel. It returns true if a delimiter token was found.
func (v *scanner) foundDelimiter() bool {
	var s = v.source[v.nextByte:]
	var matches = scanDelimiter(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenDelimiter)
		return true
	}
	return false
}

// This method adds a duration element token with the current scanner information
// to the token channel. It returns true if a duration token was found.
func (v *scanner) foundDuration() bool {
	var s = v.source[v.nextByte:]
	var matches = scanDuration(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenDuration)
		return true
	}
	return false
}

// This method adds an error token with the current scanner information to the token
// channel.
func (v *scanner) foundError() {
	var bytes = v.source[v.nextByte:]
	var _, width = utf.DecodeRune(bytes)
	v.nextByte += width
	v.emitToken(TokenError)
}

// This method adds an EOF token with the current scanner information to the token
// channel. It returns true if an EOF token was found.
func (v *scanner) foundEOF() bool {
	if v.nextByte == len(v.source) {
		v.emitToken(TokenEOF)
		return true
	}
	return false
}

// This method adds an EOL token with the current scanner information to the token
// channel. It returns true if an EOL token was found.
func (v *scanner) foundEOL() bool {
	var s = v.source[v.nextByte:]
	if byt.HasPrefix(s, []byte("\n")) {
		v.nextByte++
		v.emitToken(TokenEOL)
		v.line++
		v.position = 1
		return true
	}
	return false
}

// This method adds an identifier token with the current scanner information to
// the token channel. It returns true if an identifier token was found.
func (v *scanner) foundIdentifier() bool {
	var s = v.source[v.nextByte:]
	var matches = scanIdentifier(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenIdentifier)
		return true
	}
	return false
}

// This method adds a keyword token with the current scanner information to the
// token channel. It returns true if a keyword token was found.
func (v *scanner) foundKeyword() bool {
	var s = v.source[v.nextByte:]
	var matches = scanKeyword(s)
	if len(matches) > 0 {
		// Check to see if the match is part of an identifier.
		var r, _ = utf.DecodeRune(v.source[v.nextByte+len(matches[0]):])
		if !uni.IsLetter(r) {
			v.nextByte += len(matches[0])
			v.emitToken(TokenKeyword)
			return true
		}
	}
	return false
}

// This method adds a moment element token with the current scanner information to
// the token channel. It returns true if a moment token was found.
func (v *scanner) foundMoment() bool {
	var s = v.source[v.nextByte:]
	var matches = scanMoment(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenMoment)
		return true
	}
	return false
}

// This method adds a moniker string token with the current scanner information to
// the token channel. It returns true if a moniker token was found.
func (v *scanner) foundMoniker() bool {
	var s = v.source[v.nextByte:]
	var matches = scanMoniker(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenMoniker)
		return true
	}
	return false
}

// This method adds a narrative string token with the current scanner information
// to the token channel. It returns true if a narrative token was found.
func (v *scanner) foundNarrative() bool {
	var s = v.source[v.nextByte:]
	var matches = scanNarrative(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.line += sts.Count(matches[0], "\n")
		v.emitToken(TokenNarrative)
		return true
	}
	return false
}

// This method adds a note token with the current scanner information to the token
// channel. It returns true if a note token was found.
func (v *scanner) foundNote() bool {
	var s = v.source[v.nextByte:]
	var matches = scanNote(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenNote)
		return true
	}
	return false
}

// This method adds a number element token with the current scanner information to
// the token channel. It returns true if a number token was found.
func (v *scanner) foundNumber() bool {
	var s = v.source[v.nextByte:]
	var matches = scanNumber(s)
	if len(matches) > 0 {
		// Check to see if the match is part of an identifier or keyword.
		var r, _ = utf.DecodeRune(v.source[v.nextByte+len(matches[0]):])
		if !uni.IsLetter(r) {
			v.nextByte += len(matches[0])
			v.emitToken(TokenNumber)
			return true
		}
	}
	return false
}

// This method adds a pattern element token with the current scanner information
// to the token channel. It returns true if a pattern token was found.
func (v *scanner) foundPattern() bool {
	var s = v.source[v.nextByte:]
	var matches = scanPattern(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenPattern)
		return true
	}
	return false
}

// This method adds a percentage element token with the current scanner
// information to the token channel. It returns true if a percentage token was
// found.
func (v *scanner) foundPercentage() bool {
	var s = v.source[v.nextByte:]
	var matches = scanPercentage(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenPercentage)
		return true
	}
	return false
}

// This method adds a probability element token with the current scanner
// information to the token channel. It returns true if a probability token was
// found.
func (v *scanner) foundProbability() bool {
	var s = v.source[v.nextByte:]
	var matches = scanProbability(s)
	if len(matches) > 0 {
		// Check to see if the match is part of a range.
		if matches[0] != "1." ||
			byt.HasPrefix(v.source[v.nextByte+len(matches[0]):], []byte("..")) || // "1...x" is ok.
			!byt.HasPrefix(v.source[v.nextByte+len(matches[0]):], []byte(".")) { // "1.x" is ok.
			v.nextByte += len(matches[0])
			v.emitToken(TokenProbability)
			return true
		} // "1..x" is NOT ok.
	}
	return false
}

// This method adds a quoted string token with the current scanner information to
// the token channel. It returns true if a quote token was found.
func (v *scanner) foundQuote() bool {
	var s = v.source[v.nextByte:]
	var matches = scanQuote(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenQuote)
		return true
	}
	return false
}

// This method adds a resource element token with the current scanner information
// the token channel. It returns true if a resource token was found.
func (v *scanner) foundResource() bool {
	var s = v.source[v.nextByte:]
	var matches = scanResource(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenResource)
		return true
	}
	return false
}

// This method adds a symbol string token with the current scanner information to
// the token channel. It returns true if a symbol token was found.
func (v *scanner) foundSymbol() bool {
	var s = v.source[v.nextByte:]
	var matches = scanSymbol(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenSymbol)
		return true
	}
	return false
}

// This method adds a tag element token with the current scanner information to
// the token channel. It returns true if a tag token was found.
func (v *scanner) foundTag() bool {
	var s = v.source[v.nextByte:]
	var matches = scanTag(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenTag)
		return true
	}
	return false
}

// This method adds a version string token with the current scanner information to
// the token channel. It returns true if a version token was found.
func (v *scanner) foundVersion() bool {
	var s = v.source[v.nextByte:]
	var matches = scanVersion(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenVersion)
		return true
	}
	return false
}

// ANGLE ELEMENT SYNTAX

// These constants are used to form a regular expression for valid angle
// entities. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Angle
const (
	e         = `e`
	pi        = `pi|π`
	phi       = `phi|φ`
	tau       = `tau|τ`
	sign      = `[+-]`
	zero      = `0`
	ordinal   = `[1-9][0-9]*`
	fraction  = `\.[0-9]+`
	exponent  = `E` + sign + `?` + ordinal
	scalar    = `(?:` + ordinal + `(?:` + fraction + `)?|` + zero + fraction + `)(?:` + exponent + `)?`
	magnitude = `(?:` + e + `|` + pi + `|` + phi + `|` + tau + `|` + scalar + `)`
	angle     = `~(` + magnitude + `|` + zero + `)`
)

// This scanner is used for matching angle entities.
var angleScanner = reg.MustCompile(`^(?:` + angle + `)`)

// This function returns for the specified string an array of the matching
// subgroups for an angle element. The first string in the array is the entire
// matched string.
func scanAngle(v []byte) []string {
	return bytesToStrings(angleScanner.FindSubmatch(v))
}

// BINARY STRING SYNTAX

// These constants are used to form a regular expression for valid binary
// strings. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Binary
const (
	space  = ` `
	eol    = "\n"
	base64 = `[A-Za-z0-9+/]`
	binary = `'>((?:` + base64 + `|` + space + `|` + eol + `)*)<'`
)

// This scanner is used for matching binary strings.
var binaryScanner = reg.MustCompile(`^(?:` + binary + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a binary string. The first string in the array is the entire
// matched string.
func scanBinary(v []byte) []string {
	return bytesToStrings(binaryScanner.FindSubmatch(v))
}

// BOOLEAN ELEMENT SYNTAX

// These constants are used to form a regular expression for valid boolean
// entities. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Boolean
const (
	boolean = `false|true`
)

// This scanner is used for matching boolean entities.
var booleanScanner = reg.MustCompile(`^(?:` + boolean + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a boolean element. The first string in the array is the entire
// matched string.
func scanBoolean(v []byte) []string {
	return bytesToStrings(booleanScanner.FindSubmatch(v))
}

// COMMENT SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a commment. The first string in the array is the entire matched
// string. Since a comment can be recursive a regular expression is not used in
// this implementation.
func scanComment(v []byte) []string {
	var result []string
	var space = []byte(" ")
	var eol = []byte("\n")
	var bangAngle = []byte("!>" + "\n")
	var angleBang = []byte("<!")
	if !byt.HasPrefix(v, bangAngle) {
		return result
	}
	var angleBangAllowed = false
	var current = 3 // Skip the leading '!>\n' characters.
	var last = 0
	var level = 1
	for level > 0 {
		var s = v[current:]
		switch {
		case len(s) == 0:
			return result
		case byt.HasPrefix(s, eol):
			angleBangAllowed = true
			last = current
			current++
		case byt.HasPrefix(s, bangAngle):
			current += 3 // Skip the '!>\n' characters.
			level++      // Start a nested narrative.
		case byt.HasPrefix(s, angleBang):
			if !angleBangAllowed {
				return result
			}
			current += 2 // Skip the '<!' characters.
			level--      // Terminate the current narrative.
		default:
			if angleBangAllowed && !byt.HasPrefix(s, space) {
				angleBangAllowed = false
			}
			current++ // Accept the next character.
		}
	}
	result = append(result, string(v[:current])) // Includes bang delimeters.
	result = append(result, string(v[3:last]))   // Excludes bang delimeters.
	return result
}

// DELIMITER SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a delimiter. The first string in the array is the entire
// matched string.
func scanDelimiter(v []byte) []string {
	var result []string
	for _, delimiter := range delimiters {
		if byt.HasPrefix(v, delimiter) {
			result = append(result, string(delimiter))
		}
	}
	return result
}

// DURATION ELEMENT SYNTAX

// These constants are used to form a regular expression for valid duration
// entities. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Duration
const (
	tspan    = `(?:` + zero + `|` + ordinal + `(?:` + fraction + `)?)`
	weeks    = `(` + tspan + `W)`
	years    = `(` + tspan + `Y)`
	months   = `(` + tspan + `M)`
	days     = `(` + tspan + `D)`
	hours    = `(` + tspan + `H)`
	minutes  = `(` + tspan + `M)`
	seconds  = `(` + tspan + `S)`
	dates    = years + `?` + months + `?` + days + `?`
	times    = `(T)` + hours + `?` + minutes + `?` + seconds + `?`
	duration = `~(` + sign + `?)P(?:` + weeks + `|` + dates + `(?:` + times + `)?)`
)

// This scanner is used for matching duration entities.
var durationScanner = reg.MustCompile(`^(?:` + duration + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a duration element. The first string in the array is the entire
// matched string.
func scanDuration(v []byte) []string {
	return bytesToStrings(durationScanner.FindSubmatch(v))
}

// IDENTIFIER SYNTAX

// These constants are used to form a regular expression for valid identifiers.
// See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Identifier
const (
	letter     = `\pL` // All unicode letters and connectors like underscores.
	digit      = `\pN` // All unicode digits.
	identifier = letter + `(?:` + letter + `|` + digit + `)*`
)

// This scanner is used for matching identifiers.
var identifierScanner = reg.MustCompile(`^(?:` + identifier + `)`)

// This function returns for the specified string an array of the matching
// subgroups for an identifier. The first string in the array is the entire
// matched string.
func scanIdentifier(v []byte) []string {
	return bytesToStrings(identifierScanner.FindSubmatch(v))
}

const (
	integer = zero + `|` + sign + `?` + ordinal
)

// This scanner is used for matching integers.
var integerScanner = reg.MustCompile(`^(?:` + integer + `)`)

// This function returns for the specified string an array of the matching
// subgroups for an integer. The first string in the array is the entire
// matched string.
func scanInteger(v []byte) []string {
	return bytesToStrings(integerScanner.FindSubmatch(v))
}

// KEYWORD SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a keyword. The first string in the array is the entire
// matched string.
func scanKeyword(v []byte) []string {
	var result []string
	for _, keyword := range keywords {
		if byt.HasPrefix(v, keyword) {
			result = append(result, string(keyword))
		}
	}
	return result
}

// MOMENT ELEMENT SYNTAX

// These constants are used to form a regular expression for valid moment
// entities. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Moment
const (
	year   = `(?:` + ordinal + `|` + zero + `)`
	month  = `(?:[0][1-9])|(?:[1][012])`
	day    = `(?:[012][1-9])|(?:[3][01])`
	hour   = `(?:[01][0-9])|(?:[2][0-3])`
	minute = `[0-5][0-9]`
	second = `(?:[0-5][0-9])|(?:[6][01])`
	moment = `<(` + sign + `)?(` + year + `)(?:-(` + month + `)(?:-(` + day + `)` +
		`(?:T(` + hour + `)(?::(` + minute + `)(?::((?:` + second + `)(?:` + fraction + `)?))?)?)?)?)?>`
)

// This scanner is used for matching moment entities.
var momentScanner = reg.MustCompile(`^(?:` + moment + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a moment element. The first string in the array is the entire
// matched string.
func scanMoment(v []byte) []string {
	return bytesToStrings(momentScanner.FindSubmatch(v))
}

// MONIKER STRING SYNTAX

// These constants are used to form a regular expression for valid moniker
// strings. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Moniker
const (
	separator = `[-+.]`
	name      = letter + `(?:` + separator + `?(?:` + letter + `|` + digit + `))*`
	moniker   = `(?:/` + name + `)+` // Cannot capture each name...
)

// This scanner is used for matching moniker strings.
var monikerScanner = reg.MustCompile(`^(?:` + moniker + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a moniker string. The first string in the array is the entire
// matched string.
func scanMoniker(v []byte) []string {
	return bytesToStrings(monikerScanner.FindSubmatch(v))
}

// NARRATIVE STRING SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a narrative string. The first string in the array is the entire
// matched string. Since a narrative can be recursive a regular expression is
// not used in this implementation.
func scanNarrative(v []byte) []string {
	var result []string
	var quoteAngle = []byte(`">` + "\n")
	var angleQuote = []byte(`<"`)
	var space = []byte(" ")
	var eol = []byte("\n")
	if !byt.HasPrefix(v, quoteAngle) {
		return result
	}
	var angleQuoteAllowed = false
	var current = 3 // Skip the leading '">\n' characters.
	var last = 0
	var level = 1
	for level > 0 {
		var s = v[current:]
		switch {
		case len(s) == 0:
			return result
		case byt.HasPrefix(s, eol):
			angleQuoteAllowed = true
			last = current
			current++
		case byt.HasPrefix(s, quoteAngle):
			current += 3 // Skip the '">\n' characters.
			level++      // Start a nested narrative.
		case byt.HasPrefix(s, angleQuote):
			if !angleQuoteAllowed {
				return result
			}
			current += 2 // Skip the '<"' characters.
			level--      // Terminate the current narrative.
		default:
			if angleQuoteAllowed && !byt.HasPrefix(s, space) {
				angleQuoteAllowed = false
			}
			current++ // Accept the next character.
		}
	}
	result = append(result, string(v[:current])) // Includes quote delimeters.
	result = append(result, string(v[3:last]))   // Excludes quote delimeters.
	return result
}

// NOTE SYNTAX

// These constants are used to form a regular expression for valid notes. See
// the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Note
const (
	note = `! [^\n]*`
)

// This scanner is used for matching notes.
var noteScanner = reg.MustCompile(`^(?:` + note + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a note. The first string in the array is the entire matched
// string.
func scanNote(v []byte) []string {
	return bytesToStrings(noteScanner.FindSubmatch(v))
}

// NUMBER ELEMENT SYNTAX

// These constants are used to form a regular expression for valid number
// entities. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Number
const (
	infinity    = sign + `?(?:infinity|∞)`
	undefined   = `undefined`
	float       = sign + `?` + magnitude
	imaginary   = sign + `?` + magnitude + `?i`
	rectangular = `(` + float + `)(, )(` + imaginary + `)`
	polar       = `(` + magnitude + `)(e\^)(` + angle + `i)`
	real_       = undefined + `|` + infinity + `|` + float + `|` + zero
	complex_    = `\((?:` + rectangular + `|` + polar + `)\)`
	number      = imaginary + `|` + real_ + `|` + complex_
)

// This scanner is used for matching number entities.
var numberScanner = reg.MustCompile(`^(?:` + number + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a number element. The first string in the array is the entire
// matched string.
func scanNumber(v []byte) []string {
	return bytesToStrings(numberScanner.FindSubmatch(v))
}

// This scanner is used for matching real numbers.
var realScanner = reg.MustCompile(`^(?:` + real_ + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a real number. The first string in the array is the entire
// matched string.
func scanReal(v []byte) []string {
	return bytesToStrings(realScanner.FindSubmatch(v))
}

// PATTERN ELEMENT SYNTAX

// These constants are used to form a regular expression for valid pattern
// entities. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Pattern
const (
	base16  = `[0-9a-f]`
	unicode = `u` + base16 + `{4}`
	escape  = `\\(?:` + unicode + `|["frnt\\])`
	rune_   = `(?:` + escape + `|[^"\f\r\n\t]` + `)`
	regex   = `"(` + rune_ + `+)"\?`
	pattern = `none` + `|` + regex + `|` + `any`
)

// This scanner is used for matching pattern entities.
var patternScanner = reg.MustCompile(`^(?:` + pattern + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a pattern element. The first string in the array is the entire
// matched string.
func scanPattern(v []byte) []string {
	return bytesToStrings(patternScanner.FindSubmatch(v))
}

// PERCENTAGE ELEMENT SYNTAX

// These constants are used to form a regular expression for valid percentage
// entities. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Percentage
const (
	percentage = `(` + real_ + `)%`
)

// This scanner is used for matching percentage entities.
var percentageScanner = reg.MustCompile(`^(?:` + percentage + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a percentage element. The first string in the array is the entire
// matched string.
func scanPercentage(v []byte) []string {
	return bytesToStrings(percentageScanner.FindSubmatch(v))
}

// PROBABILITY ELEMENT SYNTAX

// These constants are used to form a regular expression for valid probability
// entities. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Probability
const (
	probability = fraction + `|1\.`
)

// This scanner is used for matching probability entities.
var probabilityScanner = reg.MustCompile(`^(?:` + probability + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a probability element. The first string in the array is the
// entire matched string.
func scanProbability(v []byte) []string {
	return bytesToStrings(probabilityScanner.FindSubmatch(v))
}

// QUOTE STRING SYNTAX

// These constants are used to form a regular expression for valid quoted
// strings. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Quote
const (
	quote = `"(` + rune_ + `*)"`
)

// This scanner is used for matching quoted strings.
var quoteScanner = reg.MustCompile(`^(?:` + quote + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a quoted string. The first string in the array is the entire
// matched string.
func scanQuote(v []byte) []string {
	return bytesToStrings(quoteScanner.FindSubmatch(v))
}

// RESOURCE ELEMENT SYNTAX

// These constants are used to form a regular expression for valid resource
// entities. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Resource
const (
	scheme    = `[a-zA-Z][0-9a-zA-Z+-.]*`
	authority = `[^/\n]+`
	path      = `[^?#>\n]*`
	query     = `[^#>\n]*`
	fragment  = `[^>\n]+`
	resource  = `<(` +
		`(` + scheme + `):` +
		`(?:` + `//(` + authority + `)` + `)?` +
		`(` + path + `)` +
		`(?:` + `\?(` + query + `)` + `)?` +
		`(?:` + `#(` + fragment + `)` + `)?` +
		`)>`
)

// This scanner is used for matching resource entities.
var resourceScanner = reg.MustCompile(`^(?:` + resource + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a resource element. The first string in the array is the entire
// matched string.
func scanResource(v []byte) []string {
	return bytesToStrings(resourceScanner.FindSubmatch(v))
}

// SYMBOL STRING SYNTAX

// These constants are used to form a regular expression for valid symbol
// strings. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Symbol
const (
	symbol = `\$(` + identifier + `(?:-` + ordinal + `)?)`
)

// This scanner is used for matching symbol strings.
var symbolScanner = reg.MustCompile(`^(?:` + symbol + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a symbol string. The first string in the array is the entire
// matched string.
func scanSymbol(v []byte) []string {
	return bytesToStrings(symbolScanner.FindSubmatch(v))
}

// TAG ELEMENT SYNTAX

// These constants are used to form a regular expression for valid tag entities.
// See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Tag
const (
	base32 = `[0-9A-DF-HJ-NP-TV-Z]` // "E", "I", "O", and "U" have been removed.
	tag    = `#(` + base32 + `+)`
)

// This scanner is used for matching tag entities.
var tagScanner = reg.MustCompile(`^(?:` + tag + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a tag element. The first string in the array is the entire
// matched string.
func scanTag(v []byte) []string {
	return bytesToStrings(tagScanner.FindSubmatch(v))
}

// VERSION STRING SYNTAX

// These constants are used to form a regular expression for valid version
// strings. See the language specification for the exact grammar:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification#Version
const (
	version = `v(` + ordinal + `(?:\.` + ordinal + `)*)` // Cannot capture each ordinal...
)

// This scanner is used for matching version strings.
var versionScanner = reg.MustCompile(`^(?:` + version + `)`)

// This function returns for the specified string an array of the matching
// subgroups for a version string. The first string in the array is the entire
// matched string.
func scanVersion(v []byte) []string {
	return bytesToStrings(versionScanner.FindSubmatch(v))
}

// PRIVATE CONSTANTS

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
	[]byte("on"),
	[]byte("notarize"),
	[]byte("matching"),
	[]byte("loop"),
	[]byte("level"),
	[]byte("let"),
	[]byte("in"),
	[]byte("if"),
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
	[]byte("AND"),
}

// This array contains the full set of delimiters used by the Bali Document
// Notation™ (BDN). They are in reverse order for proper matching.
var delimiters = [][]byte{
	[]byte("~"),
	[]byte("}"),
	[]byte("|"),
	[]byte("{"),
	[]byte("^"),
	[]byte("]"),
	[]byte("["),
	[]byte("@"),
	[]byte("?="),
	[]byte(">"),
	[]byte("="),
	[]byte("≠"),
	[]byte("<-"),
	[]byte("<"),
	[]byte(";"),
	[]byte(":="),
	[]byte(":"),
	[]byte("/="),
	[]byte("//"),
	[]byte("/"),
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
	[]byte("&"),
}

// PRIVATE FUNCTIONS

func bytesToStrings(bytes [][]byte) []string {
	var strings = make([]string, len(bytes))
	for index, array := range bytes {
		strings[index] = string(array)
	}
	return strings
}

func stringsToBytes(strings []string) [][]byte {
	var bytes = make([][]byte, len(strings))
	for index, s := range strings {
		bytes[index] = []byte(s)
	}
	return bytes
}
