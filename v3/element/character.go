/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element

import (
	mat "math"
	stc "strconv"
	uni "unicode"
	utf "unicode/utf8"
)

// CLASS ACCESS

// Reference

var characterClass = &characterClass_{
	minimumValue_: character_(0),
	maximumValue_: character_(mat.MaxInt32),
}

// Function

func Character() CharacterClassLike {
	return characterClass
}

// CLASS METHODS

// Target

type characterClass_ struct {
	minimumValue_ character_
	maximumValue_ character_
}

// Constants

func (c *characterClass_) MinimumValue() CharacterLike {
	return c.minimumValue_
}

func (c *characterClass_) MaximumValue() CharacterLike {
	return c.maximumValue_
}

// Constructors

func (c *characterClass_) MakeFromRune(rune_ rune) CharacterLike {
	return character_(rune_)
}

func (c *characterClass_) MakeFromInteger(integer int64) CharacterLike {
	return character_(integer)
}

func (c *characterClass_) MakeFromString(string_ string) CharacterLike {
	var matches = matchCharacter(string_)
	var rune_ = runeFromString(matches[0])
	return character_(rune_)
}

// Functions

func (c *characterClass_) ToLowercase(character CharacterLike) CharacterLike {
	var rune_ = rune(character.AsInteger())
	rune_ = uni.ToLower(rune_)
	return character_(rune_)
}

func (c *characterClass_) ToUppercase(character CharacterLike) CharacterLike {
	var rune_ = rune(character.AsInteger())
	rune_ = uni.ToUpper(rune_)
	return character_(rune_)
}

// INSTANCE METHODS

// Target

type character_ rune

// Discrete

func (v character_) AsBoolean() bool {
	return v > -1
}

func (v character_) AsInteger() int64 {
	return int64(v)
}

// Lexical

func (v character_) AsString() string {
	return stc.Quote(string([]rune{rune(v)}))
}

// PACKAGE FUNCTIONS

// Private

func matchCharacter(string_ string) []string {
	var matches = []string{
		string_,
	}
	// TBA - Add the pattern matching code...
	return matches
}

func runeFromString(string_ string) rune {
	var err error
	string_, err = stc.Unquote(string_)
	if err != nil {
		panic(err)
	}
	var rune_, _ = utf.DecodeRuneInString(string_)
	if err != nil {
		panic(err)
	}
	return rune_
}
