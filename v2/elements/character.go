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
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	stc "strconv"
	utf "unicode/utf8"
)

// CLASS DEFINITIONS

// This private type implements the CharacterLike interface.  It extends the
// native Go `int32` type.
type character_ int32

// This private type defines the structure associated with the class constants
// and class functions for the character elements.
type characterClass_ struct {
	minimum CharacterLike
	maximum CharacterLike
}

// CLASS CONSTANTS

// This class constant represents the minimum value for a character endpoint.
func (c *characterClass_) MinimumValue() CharacterLike {
	return c.minimum
}

// This class constant represents the maximum value for a character endpoint.
func (c *characterClass_) MaximumValue() CharacterLike {
	return c.maximum
}

// CLASS CONSTRUCTORS

// This constructor creates a new character element from the specified rune.
func (c *characterClass_) FromRune(rune_ rune) CharacterLike {
	return character_(rune_)
}

// This constructor creates a new character element from the specified integer.
func (c *characterClass_) FromInteger(integer int) CharacterLike {
	return character_(int32(integer))
}

// This constructor creates a new character element from the specified string.
func (c *characterClass_) FromString(string_ string) CharacterLike {
	var matches = uti.CharacterMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a character from an invalid string: %v", string_)
		panic(message)
	}
	var character, _ = stc.Unquote(matches[0])
	var rune_, _ = utf.DecodeRuneInString(character)
	return character_(rune_)
}

// CLASS METHODS

// Discrete Interface

// This method returns a boolean value for this character.
func (v character_) AsBoolean() bool {
	return v > -1
}

// This method returns an integer value for this character.
func (v character_) AsInteger() int {
	return int(v)
}

// Lexical Interface

// This method returns a string value for this lexical element.
func (v character_) AsString() string {
	//var string_ = `"` + string([]rune{rune(v)}) + `"`
	var string_ = stc.Quote(string([]rune{rune(v)}))
	return string_
}
