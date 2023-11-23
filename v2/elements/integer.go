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
	stc "strconv"
)

// CLASS DEFINITIONS

// This private type implements the IntegerLike interface.  It extends the
// native Go `int` type.
type integer_ int

// This private type defines the structure associated with the class constants
// and class functions for the integer elements.
type integers_ struct {
	minimum abs.IntegerLike
	maximum abs.IntegerLike
}

// CLASS CONSTANTS

// This class constant represents the minimum value for an integer endpoint.
func (c *integers_) MinimumValue() abs.IntegerLike {
	return c.minimum
}

// This class constant represents the maximum value for an integer endpoint.
func (c *integers_) MaximumValue() abs.IntegerLike {
	return c.maximum
}

// CLASS CONSTRUCTORS

// This constructor creates a new integer element from the specified integer.
func (c *integers_) FromInteger(integer int) abs.IntegerLike {
	return integer_(integer)
}

// This constructor creates a new integer element from the specified string.
func (c *integers_) FromString(string_ string) abs.IntegerLike {
	var matches = uti.IntegerMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct an integer from an invalid string: %v", string_)
		panic(message)
	}
	var integer, _ = stc.ParseInt(matches[0], 10, 0)
	return integer_(integer)
}

// CLASS METHODS

// Discrete Interface

// This method returns a boolean value for this rune.
func (v integer_) AsBoolean() bool {
	return v != 0
}

// This method returns an integer value for this rune.
func (v integer_) AsInteger() int {
	return int(v)
}

// Lexical Interface

// This method returns a string value for this lexical element.
func (v integer_) AsString() string {
	var string_ = stc.FormatInt(int64(v), 10)
	return string_
}

// Polarized Interface

// This method determines whether or not this rune is negative.
func (v integer_) IsNegative() bool {
	return v < 0
}
