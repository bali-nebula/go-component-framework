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
)

// CLASS DEFINITIONS

// This private type implements the BooleanLike interface.  It extends the
// native Go `bool` type.
type boolean_ bool

// This private type defines the structure associated with the class constants
// and class functions for the boolean elements.
type booleans_ struct {
	false_ abs.BooleanLike
	true_  abs.BooleanLike
}

// CLASS CONSTANTS

// This class constant represents a false boolean element.
func (c *booleans_) False() abs.BooleanLike {
	return c.false_
}

// This class constant represents a true boolean element.
func (c *booleans_) True() abs.BooleanLike {
	return c.true_
}

// CLASS CONSTRUCTORS

// This constructor creates a new boolean from the specified bool value.
func (c *booleans_) FromBoolean(boolean bool) abs.BooleanLike {
	return boolean_(boolean)
}

// This constructor creates a new boolean from the specified string value.
func (c *booleans_) FromString(string_ string) abs.BooleanLike {
	var matches = uti.BooleanMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a boolean from an invalid string: %v", string_)
		panic(message)
	}
	var boolean = c.FromBoolean(matches[0] == "true")
	return boolean
}

// CLASS METHODS

// Discrete Interface

// This instance method returns a boolean value for this discrete element.
func (v boolean_) AsBoolean() bool {
	return bool(v)
}

// This instance method returns an integer value for this discrete element.
func (v boolean_) AsInteger() int {
	if v {
		return 1
	}
	return 0
}

// Lexical Interface

// This instance method returns a string value for this lexical element.
func (v boolean_) AsString() string {
	var string_ string
	switch v {
	case true:
		string_ = "true"
	case false:
		string_ = "false"
	}
	return string_
}

// CLASS FUNCTIONS

// This class method returns the logical inverse of the specified boolean.
func (c *booleans_) Not(boolean abs.BooleanLike) abs.BooleanLike {
	return c.FromBoolean(!boolean.AsBoolean())
}

// This class method returns the logical conjunction of the specified
// boolean elements.
func (c *booleans_) And(first, second abs.BooleanLike) abs.BooleanLike {
	return c.FromBoolean(first.AsBoolean() && second.AsBoolean())
}

// This class method returns the logical material non-implication of the
// specified boolean elements.
func (c *booleans_) Sans(first, second abs.BooleanLike) abs.BooleanLike {
	return c.FromBoolean(first.AsBoolean() && !second.AsBoolean())
}

// This class method returns the logical disjunction of the specified
// boolean elements.
func (c *booleans_) Or(first, second abs.BooleanLike) abs.BooleanLike {
	return c.FromBoolean(first.AsBoolean() || second.AsBoolean())
}

// This class method returns the logical exclusive disjunction of the
// specified boolean elements.
func (c *booleans_) Xor(first, second abs.BooleanLike) abs.BooleanLike {
	return c.Or(c.Sans(first, second), c.Sans(second, first))
}
