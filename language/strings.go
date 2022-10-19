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
	str "github.com/craterdog-bali/go-bali-document-notation/strings"
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
	binary, _ = str.BinaryFromString(token.Value)
	return binary, token, true
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
	moniker, _ = str.MonikerFromString(token.Value)
	return moniker, token, true
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
	narrative, _ = str.NarrativeFromString(token.Value)
	return narrative, token, true
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
	quote, _ = str.QuoteFromString(token.Value)
	return quote, token, true
}

// This method attempts to parse a string primitive. It returns the
// string primitive and whether or not the string primitive was
// successfully parsed.
func (v *parser) parseString() (any, *Token, bool) {
	var ok bool
	var token *Token
	var s any
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
	version, _ = str.VersionFromString(token.Value)
	return version, token, true
}
