/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar

import (
	"bytes"
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"unicode"
	"unicode/utf8"
)

// This constructor creates a new scanner initialized with the specified array
// of bytes.
func Scanner(source []byte, tokens chan token) *scanner {
	v := &scanner{source: source, tokens: tokens}
	go v.scanTokens() // Start scanning in the background.
	return v
}

// This type defines the structure and methods for the scanner agent.
type scanner struct {
	source  []byte
	start   int
	current int
	width   int
	last    int
	tokens  chan token
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
// if no valid token or a tokenEOF was found.
func (v *scanner) scanToken() bool {
	v.skipWhiteSpace()
	switch {
	// The order of the cases below has been set to help optimize the
	// performance by placing the most common token types first.
	//
	// Must be first in case the input buffer is empty.
	case v.foundEOF():
		return false
	case v.foundEOL():
	case v.foundSymbol():
	// Must be before foundNumber().
	case v.foundPercentage():
	case v.foundNumber():
	case v.foundPattern():
	case v.foundMoniker():
	case v.foundTag():
	case v.foundVersion():
	case v.foundBinary():
	case v.foundResource():
	// Must be before foundNote().
	case v.foundComment():
	case v.foundNote():
	// Must be before foundQuote().
	case v.foundNarrative():
	case v.foundQuote():
	case v.foundMoment():
	// Must be before foundAngle().
	case v.foundDuration():
	case v.foundAngle():
	case v.foundProbability():
	case v.foundBoolean():
	case v.foundDelimiter():
	// Must be after all element and string types.
	case v.foundKeyword():
	// Must be after foundKeyword().
	case v.foundIdentifier():
	default:
		return false
	}
	return true
}

// This method scans through any non-EOL white-space in the source array and
// sets the current index to the next non-white-space rune.
func (v *scanner) skipWhiteSpace() {
	if v.current < len(v.source) {
		r := v.nextRune()
		for unicode.IsSpace(r) && r != '\n' {
			r = v.nextRune() // Accept the whitespace rune.
		}
		v.backupOne() // The last rune wasn't whitespace.
		v.start = v.current
	}
}

// This method returns the next rune in the source array without advancing
// the current index.
func (v *scanner) peekRune() rune {
	var r rune = eof
	if v.current < len(v.source) {
		s := v.source[v.current:]
		r, _ = utf8.DecodeRune(s)
	}
	return r
}

// This method returns the next rune in the source array and advances the
// current index.
//
//	Before: | rune n | rune n+1 | rune n+2 |
//	                 ^
//	              current
//	        |-width--|
//
//	After:  | rune n | rune n+1 | rune n+2 |
//	                     next   ^
//	                         current
//	        |--last--|--width---|
//
//	EOF:    | rune n | EOF
//	                 ^
//	              current
//	        |-width--|
//
// The next rune, which may be EOF, is returned.
func (v *scanner) nextRune() rune {
	var next rune = eof
	if v.current < len(v.source) {
		s := v.source[v.current:]
		v.last = v.width
		next, v.width = utf8.DecodeRune(s)
		v.current += v.width
	}
	return next
}

// This method backs up the current index by one rune.
//
//	Before: | rune n | rune n+1 | rune n+2 |
//	                            ^
//	                         current
//	        |--last--|--width---|
//
//	After:  | rune n | rune n+1 | rune n+2 |
//	                 ^
//	              current
//	        |-width--|
//
//	EOF:    | rune n | rune n+1 | EOF
//	                 ^
//	              current
//	        |-width--|
func (v *scanner) backupOne() {
	if v.width == 0 {
		panic("The scanner can only backup by one rune.")
	}
	v.current -= v.width
	v.width = v.last
	v.last = 0
}

// This method accepts the specified string as a valid token and updates the
// state of the scanner accordingly.
func (v *scanner) acceptToken(s string) {
	v.current += len(s)
	v.width = 0
	v.last = 0
}

// This method adds a token of the specified type with the current token
// information to the token channel. It then resets the start index to the
// current index position. It returns the token type of the type added to the
// channel.
func (v *scanner) emitToken(tType tokenType) tokenType {
	tValue := string(v.source[v.start:v.current])
	if tType == tokenEOF {
		tValue = "<EOF>"
	}
	t := token{tType, tValue, v.start}
	v.tokens <- t
	v.start = v.current
	v.width = 0
	v.last = 0
	return tType
}

// This method handles an error by adding an error token to the token channel.
// It returns the token type of the type added to the channel.
func (v *scanner) emitError(message string, args ...any) tokenType {
	tType := tokenError
	tValue := fmt.Sprintf(message, args...)
	v.tokens <- token{tType, tValue, v.start}
	return tType
}

// This method adds an EOF token with the current token information to the token
// channel. It returns true if an EOF token was found.
func (v *scanner) foundEOF() bool {
	if v.current == len(v.source) {
		// Can't call v.acceptToken() here since EOF wasn't a rune.
		v.emitToken(tokenEOF)
		return true
	}
	return false
}

// This method adds an EOL token with the current token information to the token
// channel. It returns true if an EOL token was found.
func (v *scanner) foundEOL() bool {
	s := v.source[v.current:]
	if bytes.HasPrefix(s, eol) {
		v.acceptToken("\n")
		v.emitToken(tokenEOL)
		return true
	}
	return false
}

// This method adds an angle element token with the current token information to
// the token channel. It returns true if an angle token was found.
func (v *scanner) foundAngle() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanAngle(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenAngle)
		return true
	}
	return false
}

// This method adds a boolean element token with the current token information
// to the token channel. It returns true if a boolean token was found.
func (v *scanner) foundBoolean() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanBoolean(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenBoolean)
		return true
	}
	return false
}

// This method adds a binary string token with the current token information to
// the token channel. It returns true if a binary token was found.
func (v *scanner) foundBinary() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanBinary(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenBinary)
		return true
	}
	return false
}

// This method adds a comment token with the current token information to the
// token channel. It returns true if a comment token was found.
func (v *scanner) foundComment() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanComment(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenComment)
		return true
	}
	return false
}

// This method adds a duration element token with the current token information
// to the token channel. It returns true if a duration token was found.
func (v *scanner) foundDuration() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanDuration(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenDuration)
		return true
	}
	return false
}

// This method adds an identifier token with the current token information to
// the token channel. It returns true if an identifier token was found.
func (v *scanner) foundIdentifier() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanIdentifier(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenIdentifier)
		return true
	}
	return false
}

// This method adds a keyword token with the current token information to the
// token channel. It returns true if a keyword token was found.
func (v *scanner) foundKeyword() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanKeyword(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenKeyword)
		return true
	}
	return false
}

// This method adds a moment element token with the current token information to
// the token channel. It returns true if a moment token was found.
func (v *scanner) foundMoment() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanMoment(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenMoment)
		return true
	}
	return false
}

// This method adds a moniker string token with the current token information to
// the token channel. It returns true if a moniker token was found.
func (v *scanner) foundMoniker() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanMoniker(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenMoniker)
		return true
	}
	return false
}

// This method adds a narrative string token with the current token information
// to the token channel. It returns true if a narrative token was found.
func (v *scanner) foundNarrative() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanNarrative(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenNarrative)
		return true
	}
	return false
}

// This method adds a note token with the current token information to the token
// channel. It returns true if a note token was found.
func (v *scanner) foundNote() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanNote(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenNote)
		return true
	}
	return false
}

// This method adds a number element token with the current token information to
// the token channel. It returns true if a number token was found.
func (v *scanner) foundNumber() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanNumber(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenNumber)
		return true
	}
	return false
}

// This method adds a pattern element token with the current token information
// to the token channel. It returns true if a pattern token was found.
func (v *scanner) foundPattern() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanPattern(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenPattern)
		return true
	}
	return false
}

// This method adds a percentage element token with the current token
// information to the token channel. It returns true if a percentage token was
// found.
func (v *scanner) foundPercentage() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanPercentage(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenPercentage)
		return true
	}
	return false
}

// This method adds a probability element token with the current token
// information to the token channel. It returns true if a probability token was
// found.
func (v *scanner) foundProbability() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanProbability(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenProbability)
		return true
	}
	return false
}

// This method adds a quoted string token with the current token information to
// the token channel. It returns true if a quote token was found.
func (v *scanner) foundQuote() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanQuote(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenQuote)
		return true
	}
	return false
}

// This method adds a resource element token with the current token information
// the token channel. It returns true if a resource token was found.
func (v *scanner) foundResource() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanResource(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenResource)
		return true
	}
	return false
}

// This method adds a symbol string token with the current token information to
// the token channel. It returns true if a symbol token was found.
func (v *scanner) foundSymbol() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanSymbol(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenSymbol)
		return true
	}
	return false
}

// This method adds a tag element token with the current token information to
// the token channel. It returns true if a tag token was found.
func (v *scanner) foundTag() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanTag(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenTag)
		return true
	}
	return false
}

// This method adds a version string token with the current token information to
// the token channel. It returns true if a version token was found.
func (v *scanner) foundVersion() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanVersion(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenVersion)
		return true
	}
	return false
}

// This method adds a delimiter token with the current token information to the
// token channel. It returns true if a delimiter token was found.
func (v *scanner) foundDelimiter() bool {
	s := v.source[v.current:]
	matches := abstractions.ScanDelimiter(s)
	if len(matches) > 0 {
		v.acceptToken(matches[0])
		v.emitToken(tokenDelimiter)
		return true
	}
	return false
}

// PRIVATE CONSTANTS

// This constant defines the pseudo-rune for the end-of-file marker.
const eof = -1

// This constant defines the pseudo-rune for the end-of-line marker.
var eol = []byte("\n")
