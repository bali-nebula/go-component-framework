/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package language

import (
	"bytes"
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"strings"
	"unicode/utf8"
)

// TOKEN

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
	var value string
	switch {
	case v.Type == TokenEOF:
		value = "<EOF>"
	case v.Type == TokenEOL:
		value = "<EOL>"
	case len(v.Value) > 60:
		value = fmt.Sprintf("%.60q...", v.Value)
	default:
		value = fmt.Sprintf("%q", v.Value)
	}
	return fmt.Sprintf("Token [type: %2d, line: %2d, position: %2d]: %s", v.Type, v.Line, v.Position, value)
}

// SCANNER

// This constructor creates a new scanner initialized with the specified array
// of bytes. The scanner will scan in tokens matching Bali Document Notation™ as
// defined in the language specification:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification
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
	case v.foundPercentage(): // Must be before foundNumber().
	case v.foundNumber():
	case v.foundVersion():
	case v.foundBoolean():
	case v.foundKeyword():
	case v.foundIdentifier(): // Must be after foundVersion(), foundBoolean() and foundKeyword().
	case v.foundNote():
	case v.foundComment():
	case v.foundPattern():
	case v.foundQuote(): // Must be after foundPattern().
	case v.foundNarrative():
	case v.foundTag():
	case v.foundMoniker():
	case v.foundBinary():
	case v.foundResource():
	case v.foundMoment():
	case v.foundDuration():
	case v.foundAngle():
	case v.foundProbability():
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
	v.position += strings.Count(tValue, "") - 1 // Add the number of runes in the token.
	return tType
}

// This method adds an angle element token with the current scanner information to
// the token channel. It returns true if an angle token was found.
func (v *scanner) foundAngle() bool {
	var s = v.source[v.nextByte:]
	var matches = abstractions.ScanAngle(s)
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
	var matches = abstractions.ScanBoolean(s)
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
	var matches = abstractions.ScanBinary(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.line += strings.Count(matches[0], "\n")
		v.emitToken(TokenBinary)
		return true
	}
	return false
}

// This method adds a comment token with the current scanner information to the
// token channel. It returns true if a comment token was found.
func (v *scanner) foundComment() bool {
	var s = v.source[v.nextByte:]
	var matches = abstractions.ScanComment(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.line += strings.Count(matches[0], "\n")
		v.emitToken(TokenComment)
		return true
	}
	return false
}

// This method adds a delimiter token with the current scanner information to the
// token channel. It returns true if a delimiter token was found.
func (v *scanner) foundDelimiter() bool {
	var s = v.source[v.nextByte:]
	var matches = abstractions.ScanDelimiter(s)
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
	var matches = abstractions.ScanDuration(s)
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
	var _, width = utf8.DecodeRune(bytes)
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
	if bytes.HasPrefix(s, []byte("\n")) {
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
	var matches = abstractions.ScanIdentifier(s)
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
	var matches = abstractions.ScanKeyword(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenKeyword)
		return true
	}
	return false
}

// This method adds a moment element token with the current scanner information to
// the token channel. It returns true if a moment token was found.
func (v *scanner) foundMoment() bool {
	var s = v.source[v.nextByte:]
	var matches = abstractions.ScanMoment(s)
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
	var matches = abstractions.ScanMoniker(s)
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
	var matches = abstractions.ScanNarrative(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.line += strings.Count(matches[0], "\n")
		v.emitToken(TokenNarrative)
		return true
	}
	return false
}

// This method adds a note token with the current scanner information to the token
// channel. It returns true if a note token was found.
func (v *scanner) foundNote() bool {
	var s = v.source[v.nextByte:]
	var matches = abstractions.ScanNote(s)
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
	var matches = abstractions.ScanNumber(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenNumber)
		return true
	}
	return false
}

// This method adds a pattern element token with the current scanner information
// to the token channel. It returns true if a pattern token was found.
func (v *scanner) foundPattern() bool {
	var s = v.source[v.nextByte:]
	var matches = abstractions.ScanPattern(s)
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
	var matches = abstractions.ScanPercentage(s)
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
	var matches = abstractions.ScanProbability(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenProbability)
		return true
	}
	return false
}

// This method adds a quoted string token with the current scanner information to
// the token channel. It returns true if a quote token was found.
func (v *scanner) foundQuote() bool {
	var s = v.source[v.nextByte:]
	var matches = abstractions.ScanQuote(s)
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
	var matches = abstractions.ScanResource(s)
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
	var matches = abstractions.ScanSymbol(s)
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
	var matches = abstractions.ScanTag(s)
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
	var matches = abstractions.ScanVersion(s)
	if len(matches) > 0 {
		v.nextByte += len(matches[0])
		v.emitToken(TokenVersion)
		return true
	}
	return false
}
