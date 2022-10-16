/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package components

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// EXPONENTIAL EXPRESSION IMPLEMENTATION

// This constructor creates a new exponential expression.
func Exponential(base, exponent any) abstractions.ExponentialLike {
	var v = &exponentialExpression{}
	// Perform argument validation.
	v.SetBase(base)
	v.SetExponent(exponent)
	return v
}

// This type defines the structure and methods associated with a exponential
// expression.
type exponentialExpression struct {
	base     any
	exponent any
}

// This method returns the base expression in this exponential expression.
func (v *exponentialExpression) GetBase() any {
	return v.base
}

// This method sets the base expression in this exponential expression to the
// specified value.
func (v *exponentialExpression) SetBase(base any) {
	if base == nil {
		panic("The base expression in a exponential expression cannot be nil.")
	}
	v.base = base
}

// This method returns the exponent expression in this exponential expression.
func (v *exponentialExpression) GetExponent() any {
	return v.exponent
}

// This method sets the exponent expression in this exponential expression to
// the specified value.
func (v *exponentialExpression) SetExponent(exponent any) {
	if exponent == nil {
		panic("The exponent expression in a exponential expression cannot be nil.")
	}
	v.exponent = exponent
}
