/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	mat "math"
	stc "strconv"
	utf "unicode/utf8"
)

// CHARACTER INTERFACE

// This constructor creates a new character element from the specified rune.
func CharacterFromRune(rune_ rune) abs.CharacterLike {
	return character_(rune_)
}

// This constructor creates a new character element from the specified integer.
func CharacterFromInteger(integer int) abs.CharacterLike {
	return character_(int32(integer))
}

// This constructor creates a new character element from the specified string.
func CharacterFromString(string_ string) abs.CharacterLike {
	var matches = uti.CharacterMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a character from an invalid string: %v", string_)
		panic(message)
	}
	var character, _ = stc.Unquote(matches[0])
	var rune_, _ = utf.DecodeRuneInString(character)
	return character_(rune_)
}

// This constructor returns the minimum value for a character endpoint.
func MinimumCharacter() abs.CharacterLike {
	return character_(0)
}

// This constructor returns the maximum value for a character endpoint.
func MaximumCharacter() abs.CharacterLike {
	return character_(mat.MaxInt32)
}

// CHARACTER IMPLEMENTATION

// This type defines the methods associated with character endpoints. It extends the
// native Go int32 type.
type character_ int32

// DISCRETE INTERFACE

// This method returns a boolean value for this character.
func (v character_) AsBoolean() bool {
	return v > -1
}

// This method returns an integer value for this character.
func (v character_) AsInteger() int {
	return int(v)
}

// LEXICAL INTERFACE

// This method returns a string value for this lexical element.
func (v character_) AsString() string {
	//var string_ = `"` + string([]rune{rune(v)}) + `"`
	var string_ = stc.Quote(string([]rune{rune(v)}))
	return string_
}
