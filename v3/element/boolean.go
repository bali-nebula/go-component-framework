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
	stc "strconv"
)

// CLASS ACCESS

// Reference

var booleanClass = &booleanClass_{
	false_: boolean_(false),
	true_:  boolean_(true),
}

// Function

func Boolean() BooleanClassLike {
	return booleanClass
}

// CLASS METHODS

// Target

type booleanClass_ struct {
	false_ boolean_
	true_  boolean_
}

// Constants

func (c *booleanClass_) False() BooleanLike {
	return c.false_
}

func (c *booleanClass_) True() BooleanLike {
	return c.true_
}

// Constructors

func (c *booleanClass_) MakeFromBoolean(boolean bool) BooleanLike {
	return boolean_(boolean)
}

func (c *booleanClass_) MakeFromString(string_ string) BooleanLike {
	var matches = matchBoolean(string_)
	var boolean = booleanFromString(matches[0])
	return boolean_(boolean)
}

// Functions

func (c *booleanClass_) Not(boolean BooleanLike) BooleanLike {
	return boolean_(!boolean.AsBoolean())
}

func (c *booleanClass_) And(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	return boolean_(first.AsBoolean() && second.AsBoolean())
}

func (c *booleanClass_) Sans(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	return boolean_(first.AsBoolean() && !second.AsBoolean())
}

func (c *booleanClass_) Or(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	return boolean_(first.AsBoolean() || second.AsBoolean())
}

func (c *booleanClass_) Xor(
	first BooleanLike,
	second BooleanLike,
) BooleanLike {
	return c.Or(c.Sans(first, second), c.Sans(second, first))
}

// INSTANCE METHODS

// Target

type boolean_ bool

// Discrete

func (v boolean_) AsBoolean() bool {
	return bool(v)
}

func (v boolean_) AsInteger() int64 {
	if v {
		return 1
	}
	return 0
}

// Lexical

func (v boolean_) AsString() string {
	return stc.FormatBool(bool(v))
}

// PACKAGE FUNCTIONS

// Private

func booleanFromString(string_ string) bool {
	var boolean, err = stc.ParseBool(string_)
	if err != nil {
		panic(err)
	}
	return boolean
}

func matchBoolean(string_ string) []string {
	var matches = []string{
		string_,
	}
	// TBA - Add the pattern matching code...
	return matches
}
