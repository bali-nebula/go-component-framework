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
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// INTRINSIC EXPRESSION IMPLEMENTATION

// This constructor creates a new intrinsic expression.
func Intrinsic(function string, arguments abs.ListLike[abs.ExpressionLike]) abs.IntrinsicLike {
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
	arguments abs.ListLike[abs.ExpressionLike]
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
func (v *intrinsicExpression) GetArgument(index int) abs.ExpressionLike {
	return v.arguments.GetItem(index)
}

// This method sets the argument at the specified index for this intrinsic
// expression.
func (v *intrinsicExpression) SetArgument(index int, argument abs.ExpressionLike) {
	if argument == nil {
		panic("Each argument for an intrinsic expression requires a value.")
	}
	v.arguments.SetItem(index, argument)
}

// This method returns the list of arguments for this intrinsic expression.
func (v *intrinsicExpression) GetArguments() abs.ListLike[abs.ExpressionLike] {
	return v.arguments
}

// This method sets the list of arguments for this intrinsic expression.
func (v *intrinsicExpression) SetArguments(arguments abs.ListLike[abs.ExpressionLike]) {
	if arguments == nil {
		panic("An intrinsic expression requires an array (possibly empty) of arguments.")
	}
	v.arguments = arguments
}

// This method returns the type of this expression.
func (v *intrinsicExpression) GetType() abs.Type {
	return abs.INTRINSIC
}
