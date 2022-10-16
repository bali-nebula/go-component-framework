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

// INTRINSIC EXPRESSION IMPLEMENTATION

// This constructor creates a new intrinsic expression.
func Intrinsic(function string, arguments abstractions.ArrayLike[any]) abstractions.IntrinsicLike {
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
	arguments abstractions.ArrayLike[any]
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

// This method returns the argument at the specified index from this intrinsic
// expression.
func (v *intrinsicExpression) GetArgument(index int) any {
	return v.arguments.GetItem(index)
}

// This method sets the argument at the specified index for this intrinsic
// expression.
func (v *intrinsicExpression) SetArgument(index int, argument any) {
	if argument == nil {
		panic("Each argument for an intrinsic expression requires a value.")
	}
	v.arguments.SetItem(index, argument)
}

// This method returns the list of arguments for this intrinsic expression.
func (v *intrinsicExpression) GetArguments() abstractions.ArrayLike[any] {
	return v.arguments
}

// This method sets the list of arguments for this intrinsic expression.
func (v *intrinsicExpression) SetArguments(arguments abstractions.ArrayLike[any]) {
	if arguments == nil {
		panic("An intrinsic expression requires an array (possibly empty) of arguments.")
	}
	v.arguments = arguments
}
