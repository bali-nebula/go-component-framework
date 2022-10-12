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
func (v *parser) parseBinary() (strings.Binary, bool) {
	var ok bool
	var binary strings.Binary
	var token = v.nextToken()
	if token.tType != tokenBinary {
		v.backupOne()
		return binary, false
	}
	binary, ok = strings.BinaryFromString(token.value)
	if !ok {
		panic(fmt.Sprintf("An invalid binary token was found: %v", token))
	}
	return binary, true
}

// This method attempts to parse a moniker string. It returns the moniker string
// and whether or not the moniker string was successfully parsed.
func (v *parser) parseMoniker() (strings.Moniker, bool) {
	var ok bool
	var moniker strings.Moniker
	var token = v.nextToken()
	if token.tType != tokenMoniker {
		v.backupOne()
		return moniker, false
	}
	moniker, ok = strings.MonikerFromString(token.value)
	if !ok {
		panic(fmt.Sprintf("An invalid moniker token was found: %v", token))
	}
	return moniker, true
}

// This method attempts to parse a narrative string. It returns the narrative
// string and whether or not the narrative string was successfully parsed.
func (v *parser) parseNarrative() (strings.Narrative, bool) {
	var ok bool
	var narrative strings.Narrative
	var token = v.nextToken()
	if token.tType != tokenNarrative {
		v.backupOne()
		return narrative, false
	}
	narrative, ok = strings.NarrativeFromString(token.value)
	if !ok {
		panic(fmt.Sprintf("An invalid narrative token was found: %v", token))
	}
	return narrative, true
}

// This method attempts to parse a quote string. It returns the quote string
// and whether or not the quote string was successfully parsed.
func (v *parser) parseQuote() (strings.Quote, bool) {
	var ok bool
	var quote strings.Quote
	var token = v.nextToken()
	if token.tType != tokenQuote {
		v.backupOne()
		return quote, false
	}
	quote, ok = strings.QuoteFromString(token.value)
	if !ok {
		panic(fmt.Sprintf("An invalid quote token was found: %v", token))
	}
	return quote, true
}

// This method attempts to parse a string primitive. It returns the
// string primitive and whether or not the string primitive was
// successfully parsed.
func (v *parser) parseString() (any, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var str any
	str, ok = v.parseBinary()
	if !ok {
		str, ok = v.parseMoniker()
	}
	if !ok {
		str, ok = v.parseNarrative()
	}
	if !ok {
		str, ok = v.parseQuote()
	}
	if !ok {
		str, ok = v.parseVersion()
	}
	if !ok {
		str = nil
	}
	return str, ok
}

// This method attempts to parse a version string. It returns the version
// string and whether or not the version string was successfully parsed.
func (v *parser) parseVersion() (strings.Version, bool) {
	var ok bool
	var version strings.Version
	var token = v.nextToken()
	if token.tType != tokenVersion {
		v.backupOne()
		return version, false
	}
	version, ok = strings.VersionFromString(token.value)
	if !ok {
		panic(fmt.Sprintf("An invalid version token was found: %v", token))
	}
	return version, true
}
