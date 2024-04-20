/*
................................................................................
.                   Copyright (c) 2024.  All Rights Reserved.                  .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package bali

import (
	col "github.com/craterdog/go-collection-framework/v3/collection"
)

// CLASS ACCESS

// Reference

var valuesClass = &valuesClass_{
	// This class has no private constants to initialize.
}

// Function

func Values() ValuesClassLike {
	return valuesClass
}

// CLASS METHODS

// Target

type valuesClass_ struct {
	// This class has no private constants.
}

// Constants

// Constructors

func (c *valuesClass_) MakeWithValues(values col.ListLike[ValueLike]) ValuesLike {
	return &values_{
		values_: values,
	}
}

func (c *valuesClass_) MakeWithAttributes(
	value ValueLike,
	note string,
) ValuesLike {
	return &values_{
		value_: value,
		note_:  note,
	}
}

// Functions

// INSTANCE METHODS

// Target

type values_ struct {
	values_ col.ListLike[ValueLike]
	note_   string
	value_  ValueLike
}

// Attributes

func (v *values_) GetValues() col.ListLike[ValueLike] {
	return v.values_
}

func (v *values_) GetNote() string {
	return v.note_
}

// Public

// Private
