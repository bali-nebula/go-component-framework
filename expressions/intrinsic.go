/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package expressions

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// INTRINSIC EXPRESSION IMPLEMENTATION

// This constructor creates a new intrinsic expression.
func Intrinsic(function string, arguments abs.Arguments) abs.IntrinsicLike {
	var v = &intrinsicExpression{}
	// Perform argument validation.
	v.SetFunction(function)
	v.SetArguments(arguments)
	return v
}

// This type defines the structure and methods associated with an intrinsic
// expression.
type intrinsicExpression struct {
	function  string
	arguments abs.Arguments
}

// This method returns the function name for this intrinsic expression.
func (v *intrinsicExpression) GetFunction() string {
	return v.function
}

// This method sets the function name for this intrinsic expression.
func (v *intrinsicExpression) SetFunction(function string) {
	if len(function) == 0 {
		panic("An intrinsic expression requires a function name.")
	}
	v.function = function
}

// This method returns the list of arguments for this intrinsic expression.
func (v *intrinsicExpression) GetArguments() abs.Arguments {
	return v.arguments
}

// This method sets the list of arguments for this intrinsic expression.
func (v *intrinsicExpression) SetArguments(arguments abs.Arguments) {
	if arguments == nil {
		panic("An intrinsic expression requires an array (possibly empty) of arguments.")
	}
	v.arguments = arguments
}
