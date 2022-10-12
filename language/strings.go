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
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/strings"
)

// This method attempts to parse a binary string. It returns the binary
// string and whether or not the binary string was successfully parsed.
func (v *parser) parseBinary() (strings.Binary, *Token, bool) {
	var ok bool
	var token *Token
	var binary strings.Binary
	token = v.nextToken()
	if token.Type != TokenBinary {
		v.backupOne()
		return binary, token, false
	}
	binary, ok = strings.BinaryFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid binary token was found: %v", token))
	}
	return binary, token, true
}

// This method attempts to parse a moniker string. It returns the moniker string
// and whether or not the moniker string was successfully parsed.
func (v *parser) parseMoniker() (strings.Moniker, *Token, bool) {
	var ok bool
	var token *Token
	var moniker strings.Moniker
	token = v.nextToken()
	if token.Type != TokenMoniker {
		v.backupOne()
		return moniker, token, false
	}
	moniker, ok = strings.MonikerFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid moniker token was found: %v", token))
	}
	return moniker, token, true
}

// This method attempts to parse a narrative string. It returns the narrative
// string and whether or not the narrative string was successfully parsed.
func (v *parser) parseNarrative() (strings.Narrative, *Token, bool) {
	var ok bool
	var token *Token
	var narrative strings.Narrative
	token = v.nextToken()
	if token.Type != TokenNarrative {
		v.backupOne()
		return narrative, token, false
	}
	narrative, ok = strings.NarrativeFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid narrative token was found: %v", token))
	}
	return narrative, token, true
}

// This method attempts to parse a quote string. It returns the quote string
// and whether or not the quote string was successfully parsed.
func (v *parser) parseQuote() (strings.Quote, *Token, bool) {
	var ok bool
	var token *Token
	var quote strings.Quote
	token = v.nextToken()
	if token.Type != TokenQuote {
		v.backupOne()
		return quote, token, false
	}
	quote, ok = strings.QuoteFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid quote token was found: %v", token))
	}
	return quote, token, true
}

// This method attempts to parse a string primitive. It returns the
// string primitive and whether or not the string primitive was
// successfully parsed.
func (v *parser) parseString() (any, *Token, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var token *Token
	var str any
	str, token, ok = v.parseBinary()
	if !ok {
		str, token, ok = v.parseMoniker()
	}
	if !ok {
		str, token, ok = v.parseNarrative()
	}
	if !ok {
		str, token, ok = v.parseQuote()
	}
	if !ok {
		str, token, ok = v.parseVersion()
	}
	if !ok {
		str = nil
	}
	return str, token, ok
}

// This method attempts to parse a version string. It returns the version
// string and whether or not the version string was successfully parsed.
func (v *parser) parseVersion() (strings.Version, *Token, bool) {
	var ok bool
	var token *Token
	var version strings.Version
	token = v.nextToken()
	if token.Type != TokenVersion {
		v.backupOne()
		return version, token, false
	}
	version, ok = strings.VersionFromString(token.Value)
	if !ok {
		panic(fmt.Sprintf("An invalid version token was found: %v", token))
	}
	return version, token, true
}
