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
)

// CLASS DEFINITIONS

// This private type implements the BooleanLike interface.  It extends the
// native Go `bool` type.
type boolean_ bool

// This private type defines the structure associated with the class constants
// and class functions for the boolean elements.
type booleanClass_ struct {
	false_ BooleanLike
	true_  BooleanLike
}

// CLASS CONSTANTS

// This class constant represents a false boolean element.
func (c *booleanClass_) False() BooleanLike {
	return c.false_
}

// This class constant represents a true boolean element.
func (c *booleanClass_) True() BooleanLike {
	return c.true_
}

// CLASS CONSTRUCTORS

// This constructor creates a new boolean from the specified bool value.
func (c *booleanClass_) FromBoolean(boolean bool) BooleanLike {
	return boolean_(boolean)
}

// This constructor creates a new boolean from the specified string value.
func (c *booleanClass_) FromString(string_ string) BooleanLike {
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
func (c *booleanClass_) Not(boolean BooleanLike) BooleanLike {
	return c.FromBoolean(!boolean.AsBoolean())
}

// This class method returns the logical conjunction of the specified
// boolean elements.
func (c *booleanClass_) And(first, second BooleanLike) BooleanLike {
	return c.FromBoolean(first.AsBoolean() && second.AsBoolean())
}

// This class method returns the logical material non-implication of the
// specified boolean elements.
func (c *booleanClass_) Sans(first, second BooleanLike) BooleanLike {
	return c.FromBoolean(first.AsBoolean() && !second.AsBoolean())
}

// This class method returns the logical disjunction of the specified
// boolean elements.
func (c *booleanClass_) Or(first, second BooleanLike) BooleanLike {
	return c.FromBoolean(first.AsBoolean() || second.AsBoolean())
}

// This class method returns the logical exclusive disjunction of the
// specified boolean elements.
func (c *booleanClass_) Xor(first, second BooleanLike) BooleanLike {
	return c.Or(c.Sans(first, second), c.Sans(second, first))
}
