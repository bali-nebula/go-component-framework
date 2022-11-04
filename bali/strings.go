/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	str "github.com/craterdog-bali/go-bali-document-notation/strings"
	stc "strconv"
	sts "strings"
	uni "unicode"
)

// This method attempts to parse a binary string. It returns the binary
// string and whether or not the binary string was successfully parsed.
func (v *parser) parseBinary() (str.Binary, *Token, bool) {
	var token *Token
	var binary str.Binary
	token = v.nextToken()
	if token.Type != TokenBinary {
		v.backupOne()
		return binary, token, false
	}
	var matches = scanBinary([]byte(token.Value))
	// Remove all whitespace and the "'" delimiters.
	binary = str.Binary(sts.Map(func(r rune) rune {
		if uni.IsSpace(r) {
			return -1
		}
		return r
	}, matches[1]))
	return binary, token, true
}

const lineLength = 60 // 60 base 64 characters encode 45 bytes per line.

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatBinary(binary str.Binary) {
	v.state.AppendString("'")
	var s = string(binary)
	var length = len(s)
	if length > lineLength {
		// Spans multiple lines.
		v.state.IncrementDepth()
		for index := 0; index < length; {
			v.state.AppendNewline()
			var next = index + lineLength
			if next > length {
				next = length
			}
			v.state.AppendString(s[index:next])
			index = next
		}
		v.state.DecrementDepth()
		v.state.AppendNewline()
	} else {
		v.state.AppendString(s)
	}
	v.state.AppendString("'")
}

// This method attempts to parse a moniker string. It returns the moniker string
// and whether or not the moniker string was successfully parsed.
func (v *parser) parseMoniker() (str.Moniker, *Token, bool) {
	var token *Token
	var moniker str.Moniker
	token = v.nextToken()
	if token.Type != TokenMoniker {
		v.backupOne()
		return moniker, token, false
	}
	var matches = scanMoniker([]byte(token.Value))
	moniker = str.Moniker(matches[0])
	return moniker, token, true
}

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatMoniker(moniker str.Moniker) {
	var s = string(moniker)
	v.state.AppendString(s)
}

// This method attempts to parse a narrative string. It returns the narrative
// string and whether or not the narrative string was successfully parsed.
func (v *parser) parseNarrative() (str.Narrative, *Token, bool) {
	var token *Token
	var narrative str.Narrative
	token = v.nextToken()
	if token.Type != TokenNarrative {
		v.backupOne()
		return narrative, token, false
	}
	narrative = str.Narrative(trimIndentation(token.Value))
	return narrative, token, true
}

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatNarrative(narrative str.Narrative) {
	var s = string(narrative)
	var lines = sts.Split(s, "\n")
	v.state.AppendString(`">`)
	for _, line := range lines {
		v.state.AppendNewline()
		v.state.AppendString(line)
	}
	v.state.AppendNewline()
	v.state.AppendString(`<"`)
}

// This method attempts to parse a quote string. It returns the quote string
// and whether or not the quote string was successfully parsed.
func (v *parser) parseQuote() (str.Quote, *Token, bool) {
	var token *Token
	var quote str.Quote
	token = v.nextToken()
	if token.Type != TokenQuote {
		v.backupOne()
		return quote, token, false
	}
	var matches = scanQuote([]byte(token.Value))
	// We must unquote the full token string properly.
	var unquoted, _ = stc.Unquote(matches[0])
	quote = str.Quote(unquoted)
	return quote, token, true
}

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatQuote(quote str.Quote) {
	// We must requote the string string properly.
	var s = stc.Quote(string(quote))
	v.state.AppendString(s)
}

// This method attempts to parse a string sequence. It returns the
// string sequence and whether or not the string sequence was
// successfully parsed.
func (v *parser) parseString() (abs.StringLike, *Token, bool) {
	var ok bool
	var token *Token
	var s abs.StringLike
	s, token, ok = v.parseQuote()
	if !ok {
		s, token, ok = v.parseMoniker()
	}
	if !ok {
		s, token, ok = v.parseVersion()
	}
	if !ok {
		s, token, ok = v.parseBinary()
	}
	if !ok {
		s, token, ok = v.parseNarrative()
	}
	if !ok {
		// Override any empty strings returned from failed parsing attempts.
		s = nil
	}
	return s, token, ok
}

// This method adds the canonical string format for the specified string to the
// state of the formatter.
func (v *formatter) formatString(s abs.StringLike) {
	pattern, ok := s.(ele.Pattern)
	if ok {
		v.formatPattern(pattern)
		return
	}
	resource, ok := s.(ele.Resource)
	if ok {
		v.formatResource(resource)
		return
	}
	symbol, ok := s.(ele.Symbol)
	if ok {
		v.formatSymbol(symbol)
		return
	}
	tag, ok := s.(ele.Tag)
	if ok {
		v.formatTag(tag)
		return
	}
	binary, ok := s.(str.Binary)
	if ok {
		v.formatBinary(binary)
		return
	}
	moniker, ok := s.(str.Moniker)
	if ok {
		v.formatMoniker(moniker)
		return
	}
	narrative, ok := s.(str.Narrative)
	if ok {
		v.formatNarrative(narrative)
		return
	}
	quote, ok := s.(str.Quote)
	if ok {
		v.formatQuote(quote)
		return
	}
	version, ok := s.(str.Version)
	if ok {
		v.formatVersion(version)
		return
	}
	panic("An invalid string was passed to the formatter.")
}

// This method attempts to parse a version string. It returns the version
// string and whether or not the version string was successfully parsed.
func (v *parser) parseVersion() (str.Version, *Token, bool) {
	var token *Token
	var version str.Version
	token = v.nextToken()
	if token.Type != TokenVersion {
		v.backupOne()
		return version, token, false
	}
	var matches = scanVersion([]byte(token.Value))
	version = str.Version(matches[1]) // Remove the leading "v".
	return version, token, true
}

// This method adds the canonical format for the specified string to the state
// of the formatter.
func (v *formatter) formatVersion(version str.Version) {
	var s = "v" + string(version)
	v.state.AppendString(s)
}
