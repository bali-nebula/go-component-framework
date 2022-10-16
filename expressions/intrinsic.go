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
	if len(function) == 0 {
		panic("An intrinsic expression requires a function name.")
	}
	if arguments == nil {
		panic("An intrinsic expression requires an array (possibly empty) of arguments.")
	}
	return &intrinsicExpression{function, arguments}
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

// This method returns the argument at the specified index from this intrinsic
// expression.
func (v *intrinsicExpression) GetArgument(index int) any {
	return v.arguments.GetItem(index)
}

// This method returns the list of arguments for this intrinsic expression.
func (v *intrinsicExpression) GetArguments() abstractions.ArrayLike[any] {
	return v.arguments
}
