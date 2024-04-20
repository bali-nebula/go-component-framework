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
)

// CLASS ACCESS

// Reference

var integerClass = &integerClass_{
	minimumValue_: integer_(mat.MinInt),
	maximumValue_: integer_(mat.MaxInt),
}

// Function

func Integer() IntegerClassLike {
	return integerClass
}

// CLASS METHODS

// Target

type integerClass_ struct {
	minimumValue_ integer_
	maximumValue_ integer_
}

// Constants

func (c *integerClass_) MinimumValue() IntegerLike {
	return c.minimumValue_
}

func (c *integerClass_) MaximumValue() IntegerLike {
	return c.maximumValue_
}

// Constructors

func (c *integerClass_) MakeFromInteger(integer int64) IntegerLike {
	return integer_(integer)
}

func (c *integerClass_) MakeFromString(string_ string) IntegerLike {
	var matches = matchInteger(string_)
	var integer = integerFromString(matches[0])
	return integer_(integer)
}

// INSTANCE METHODS

// Target

type integer_ int64

// Attributes

// Discrete

func (v integer_) AsBoolean() bool {
	return v != 0
}

func (v integer_) AsInteger() int64 {
	return int64(v)
}

// Lexical

func (v integer_) AsString() string {
	return stc.FormatInt(int64(v), 10)
}

// Polarized

func (v integer_) IsNegative() bool {
	return v < 0
}

// PACKAGE FUNCTIONS

// Private

func integerFromString(string_ string) int64 {
	var integer, err = stc.ParseInt(string_, 10, 64)
	if err != nil {
		panic(err)
	}
	return integer
}

func matchInteger(string_ string) []string {
	var matches = []string{
		string_,
	}
	// TBA - Add the pattern matching code...
	return matches
}
