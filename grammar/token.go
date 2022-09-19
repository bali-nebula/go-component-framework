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
	"fmt"
)

// This enumeration defines all possible token types including the error token.
const (
	// The first two token types must be first.
	tokenError tokenType = iota
	tokenEOF
	tokenEOL
	tokenAngle
	tokenBinary
	tokenBoolean
	tokenComment
	tokenDelimiter
	tokenDuration
	tokenIdentifier
	tokenKeyword
	tokenMoment
	tokenName
	tokenNarrative
	tokenNote
	tokenNumber
	tokenPattern
	tokenPercentage
	tokenProbability
	tokenQuote
	tokenRegex
	tokenResource
	tokenSymbol
	tokenTag
	tokenVersion
)

// This integer type is used as a type identifier for each token.
type tokenType int

// This type defines the structure and methods for each token returned by the
// scanner.
type token struct {
	typ tokenType
	val string
	pos int // The byte index of the position of the token in the input.
}

// This method returns the a canonical string version of this token.
func (v *token) String() string {
	switch {
	case v.typ == tokenEOF:
		return "<EOF>"
	case v.typ == tokenEOL:
		return "<EOL>"
	case v.typ == tokenError:
		return v.val
	case len(v.val) > 10:
		return fmt.Sprintf("%.10q...", v.val)
	default:
		return fmt.Sprintf("%q", v.val)
	}
}
