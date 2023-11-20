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
	//fmt "fmt"
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	reg "regexp"
	sts "strings"
	uni "unicode"
	utf "unicode/utf8"
)

// SCANNER INTERFACE

// The POSIX standard end-of-line character.
const EOL = "\n"

// This enumeration defines all possible token types including the error token.
const (
	TokenANGLE       TokenType = "ANGLE"
	TokenBINARY      TokenType = "BINARY"
	TokenBOOLEAN     TokenType = "BOOLEAN"
	TokenBYTECODE    TokenType = "BYTECODE"
	TokenCOMMENT     TokenType = "COMMENT"
	TokenDELIMITER   TokenType = "DELIMITER"
	TokenDURATION    TokenType = "DURATION"
	TokenEOL         TokenType = "EOL"
	TokenEOF         TokenType = "EOF"
	TokenERROR       TokenType = "ERROR"
	TokenIDENTIFIER  TokenType = "IDENTIFIER"
	TokenINTRINSIC   TokenType = "INTRINSIC"
	TokenKEYWORD     TokenType = "KEYWORD"
	TokenMOMENT      TokenType = "MOMENT"
	TokenNAME        TokenType = "NAME"
	TokenNARRATIVE   TokenType = "NARRATIVE"
	TokenNOTE        TokenType = "NOTE"
	TokenNUMBER      TokenType = "NUMBER"
	TokenPATTERN     TokenType = "PATTERN"
	TokenPERCENTAGE  TokenType = "PERCENTAGE"
	TokenPROBABILITY TokenType = "PROBABILITY"
	TokenQUOTE       TokenType = "QUOTE"
	TokenRESOURCE    TokenType = "RESOURCE"
	TokenSYMBOL      TokenType = "SYMBOL"
	TokenTAG         TokenType = "TAG"
	TokenVERSION     TokenType = "VERSION"
	TokenWHITESPACE  TokenType = "WHITESPACE"
)

// This function creates a new scanner initialized with the specified array
// of bytes. The scanner will automatically generating tokens that match the
// corresponding regular expressions.
func ScanTokens(source []byte, tokens chan Token) *scanner {
	var v = &scanner{source: source, line: 1, position: 1, tokens: tokens}
	go v.generateTokens() // Start scanning in the background.
	return v
}

// SCANNER IMPLEMENTATION

// This private function converts an array of byte arrays into an array of
// strings.
func bytesToStrings(bytes [][]byte) []string {
	var strings = make([]string, len(bytes))
	for index, array := range bytes {
		strings[index] = string(array)
	}
	return strings
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
	line      int // The line number in the source bytes of the next rune.
	position  int // The position in the current line of the first rune in the next token.
	tokens    chan Token
}

// This method determines whether or not the scanner is at the end of the source
// bytes and adds an EOF token with the current scanner information to the token
// channel if it is at the end.
func (v *scanner) atEOF() bool {
	if v.nextByte == len(v.source) {
		// The last byte in a POSIX standard file must be an EOL character.
		if byt.HasPrefix(v.source[v.nextByte-1:], []byte(EOL)) {
			v.emitToken("", TokenEOF)
			return true
		}
	}
	return false
}

// This method adds a new error token with the current scanner information
// to the token channel.
func (v *scanner) atError() {
	var bytes = v.source[v.nextByte:]
	var character, _ = utf.DecodeRune(bytes)
	v.emitToken(string(character), TokenERROR)
}

// This method adds a token of the specified type with the current scanner
// information to the token channel. It then resets the first byte index to the
// next byte index position. It returns the token type of the type added to the
// channel.
func (v *scanner) emitToken(tValue string, tType TokenType) {
	var byteCount = len(tValue)
	var runeCount = sts.Count(tValue, "") - 1 // Empty string adds one to count.
	var eolCount = sts.Count(tValue, EOL)
	var lastEOL = sts.LastIndex(tValue, EOL) + 1 // Convert to ordinal indexing.
	if tType != TokenWHITESPACE {
		if tType == TokenEOF {
			tValue = "<EOFL>"
		}
		if tType == TokenERROR {
			switch tValue {
			case "\a":
				tValue = "<BELL>"
			case "\b":
				tValue = "<BKSP>"
			case "\t":
				tValue = "<HTAB>"
			case "\f":
				tValue = "<FMFD>"
			case "\n":
				tValue = "<EOLN>"
			case "\r":
				tValue = "<CRTN>"
			case "\v":
				tValue = "<VTAB>"
			}
		}
		var token = Token{tType, tValue, v.line, v.position}
		//fmt.Println(token)
		v.tokens <- token
	}
	v.nextByte += byteCount
	v.firstByte = v.nextByte
	v.line += eolCount
	if eolCount > 0 {
		v.position = runeCount - lastEOL + 1
	} else {
		v.position += runeCount
	}
}

// This method continues scanning tokens from the source bytes until an error
// occurs or the end of file is reached. It then closes the token channel.
func (v *scanner) generateTokens() {
	for v.processToken() {
	}
	close(v.tokens)
}

// This method attempts to scan any token starting with the next rune in the
// source bytes. It checks for each type of token as the cases for the switch
// statement. If that token type is found, this method returns true and skips
// the rest of the cases.  If no valid token is found, or if the scanner is at
// the end of the source bytes, this method returns false.
func (v *scanner) processToken() bool {
	switch {
	// Check for end-of-file marker.
	case v.atEOF():
		return false
	// Must be scanned first.
	case v.scanWHITESPACE():
	case v.scanINTRINSIC():
	case v.scanKEYWORD():
	// Annotation token types.
	case v.scanNOTE():
	case v.scanCOMMENT():
	// Element token types.
	case v.scanANGLE():
	case v.scanBOOLEAN():
	case v.scanDURATION():
	case v.scanMOMENT():
	case v.scanPATTERN():
	case v.scanPERCENTAGE():
	case v.scanPROBABILITY():
	case v.scanRESOURCE():
	// String token types.
	case v.scanBINARY():
	case v.scanBYTECODE():
	case v.scanNAME():
	case v.scanNARRATIVE():
	case v.scanQUOTE():
	case v.scanSYMBOL():
	case v.scanTAG():
	case v.scanVERSION():
	// Must be scanned last.
	case v.scanNUMBER():
	case v.scanIDENTIFIER():
	case v.scanDELIMITER():
	case v.scanEOL():
	// No valid token was found.
	default:
		v.atError()
		return false
	}
	// Successfully processed a token.
	return true
}

// This method adds a new angle token with the current scanner information
// to the token channel. It returns true if a new angle token was found.
func (v *scanner) scanANGLE() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.AngleMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenANGLE)
		return true
	}
	return false
}

// This method adds a new binary token with the current scanner information
// to the token channel. It returns true if a new binary token was found.
func (v *scanner) scanBINARY() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.BinaryMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenBINARY)
		return true
	}
	return false
}

// This method adds a new boolean token with the current scanner information
// to the token channel. It returns true if a new boolean token was found.
func (v *scanner) scanBOOLEAN() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.BooleanMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenBOOLEAN)
		return true
	}
	return false
}

// This method adds a new bytecode token with the current scanner information
// to the token channel. It returns true if a new bytecode token was found.
func (v *scanner) scanBYTECODE() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.BytecodeMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenBYTECODE)
		return true
	}
	return false
}

// This method adds a new comment token with the current scanner information
// to the token channel. It returns true if a new comment token was found.
func (v *scanner) scanCOMMENT() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.CommentMatcher.FindSubmatch(s))
	//matches = scanComment(s)
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenCOMMENT)
		return true
	}
	return false
}

// This method adds a new delimiter token with the current scanner information
// to the token channel. It returns true if a new delimiter token was found.
func (v *scanner) scanDELIMITER() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.DelimiterMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenDELIMITER)
		return true
	}
	return false
}

// This method adds a new duration token with the current scanner information
// to the token channel. It returns true if a new duration token was found.
func (v *scanner) scanDURATION() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.DurationMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenDURATION)
		return true
	}
	return false
}

// This method adds a new EOL token with the current scanner information
// to the token channel. It returns true if a new EOL token was found.
func (v *scanner) scanEOL() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.EolMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenEOL)
		return true
	}
	return false
}

// This method adds a new identifier token with the current scanner information
// to the token channel. It returns true if a new identifier token was found.
func (v *scanner) scanIDENTIFIER() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.IdentifierMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenIDENTIFIER)
		return true
	}
	return false
}

// This method adds a new intrinsic token with the current scanner information
// to the token channel. It returns true if a new intrinsic token was found.
func (v *scanner) scanINTRINSIC() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.IntrinsicMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		// Check to see if the match is part of an identifier.
		var r, _ = utf.DecodeRune(v.source[v.nextByte+len(matches[0]):])
		if !uni.IsLetter(r) {
			v.emitToken(matches[0], TokenINTRINSIC)
			return true
		}
	}
	return false
}

// This method adds a new keyword token with the current scanner information
// to the token channel. It returns true if a new keyword token was found.
func (v *scanner) scanKEYWORD() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.KeywordMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		// Check to see if the match is part of an identifier.
		var r, _ = utf.DecodeRune(v.source[v.nextByte+len(matches[0]):])
		if !uni.IsLetter(r) {
			v.emitToken(matches[0], TokenKEYWORD)
			return true
		}
	}
	return false
}

// This method adds a new moment token with the current scanner information
// to the token channel. It returns true if a new moment token was found.
func (v *scanner) scanMOMENT() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.MomentMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenMOMENT)
		return true
	}
	return false
}

// This method adds a new name token with the current scanner information
// to the token channel. It returns true if a new name token was found.
func (v *scanner) scanNAME() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.NameMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenNAME)
		return true
	}
	return false
}

// This method adds a new narrative token with the current scanner information
// to the token channel. It returns true if a new narrative token was found.
func (v *scanner) scanNARRATIVE() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.NarrativeMatcher.FindSubmatch(s))
	//matches = scanNarrative(s)
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenNARRATIVE)
		return true
	}
	return false
}

// This method adds a new note token with the current scanner information
// to the token channel. It returns true if a new note token was found.
func (v *scanner) scanNOTE() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.NoteMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenNOTE)
		return true
	}
	return false
}

// This method adds a new number token with the current scanner information
// to the token channel. It returns true if a new number token was found.
func (v *scanner) scanNUMBER() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.NumberMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		// Check to see if the match is part of an identifier or keyword.
		var r, _ = utf.DecodeRune(v.source[v.nextByte+len(matches[0]):])
		if !uni.IsLetter(r) {
			v.emitToken(matches[0], TokenNUMBER)
			return true
		}
	}
	return false
}

// This method adds a new pattern token with the current scanner information
// to the token channel. It returns true if a new pattern token was found.
func (v *scanner) scanPATTERN() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.PatternMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenPATTERN)
		return true
	}
	return false
}

// This method adds a new percentage token with the current scanner information
// to the token channel. It returns true if a new percentage token was found.
func (v *scanner) scanPERCENTAGE() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.PercentageMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenPERCENTAGE)
		return true
	}
	return false
}

// This method adds a new probability token with the current scanner information
// to the token channel. It returns true if a new probability token was found.
func (v *scanner) scanPROBABILITY() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.ProbabilityMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		// Check to see if the match is part of a range.
		if matches[0] != "1." ||
			byt.HasPrefix(v.source[v.nextByte+len(matches[0]):], []byte("..")) || // "1...x" is ok.
			!byt.HasPrefix(v.source[v.nextByte+len(matches[0]):], []byte(".")) { // "1.x" is ok.
			v.emitToken(matches[0], TokenPROBABILITY)
			return true
		} // "1..x" is NOT ok.
	}
	return false
}

// This method adds a new quote token with the current scanner information
// to the token channel. It returns true if a new quote token was found.
func (v *scanner) scanQUOTE() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.QuoteMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenQUOTE)
		return true
	}
	return false
}

// This method adds a new resource token with the current scanner information
// to the token channel. It returns true if a new resource token was found.
func (v *scanner) scanRESOURCE() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.ResourceMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenRESOURCE)
		return true
	}
	return false
}

// This method adds a new symbol token with the current scanner information
// to the token channel. It returns true if a new symbol token was found.
func (v *scanner) scanSYMBOL() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.SymbolMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenSYMBOL)
		return true
	}
	return false
}

// This method adds a new tag token with the current scanner information
// to the token channel. It returns true if a new tag token was found.
func (v *scanner) scanTAG() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.TagMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenTAG)
		return true
	}
	return false
}

// This method adds a new version token with the current scanner information
// to the token channel. It returns true if a new version token was found.
func (v *scanner) scanVERSION() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.VersionMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenVERSION)
		return true
	}
	return false
}

// This method adds a new whitespace token with the current scanner information
// to the token channel. It returns true if a new whitespace token was found.
func (v *scanner) scanWHITESPACE() bool {
	var s = v.source[v.nextByte:]
	var matches = bytesToStrings(uti.WhitespaceMatcher.FindSubmatch(s))
	if len(matches) > 0 {
		v.emitToken(matches[0], TokenWHITESPACE)
		return true
	}
	return false
}

/*
// COMMENT SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a commment. The first string in the array is the entire matched
// string. Since a comment can be recursive a regular expression is not used in
// this implementation.
func scanComment(v []byte) []string {
	var result []string
	var space = []byte(" ")
	var eol = []byte(EOL)
	var bangAngle = []byte("!>" + EOL)
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
			level++      // Start a nested comment.
		case byt.HasPrefix(s, angleBang):
			if !angleBangAllowed {
				return result
			}
			current += 2 // Skip the '<!' characters.
			level--      // Terminate the current comment.
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

// NARRATIVE STRING SYNTAX

// This function returns for the specified string an array of the matching
// subgroups for a narrative string. The first string in the array is the entire
// matched string. Since a narrative can be recursive a regular expression is
// not used in this implementation.
func scanNarrative(v []byte) []string {
	var result []string
	var quoteAngle = []byte(`">` + EOL)
	var angleQuote = []byte(`<"`)
	var space = []byte(" ")
	var eol = []byte(EOL)
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
*/
