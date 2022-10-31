/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package procedures

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// SELECT CLAUSE IMPLEMENTATION

// This constructor creates a new select clause.
func SelectClause(control abs.ExpressionLike, options abs.ListLike[abs.BlockLike]) abs.SelectClauseLike {
	var v = &selectClause{}
	// Perform argument validation.
	v.SetControl(control)
	v.SetOptions(options)
	return v
}

// This type defines the structure and methods associated with a select clause.
type selectClause struct {
	control abs.ExpressionLike
	options abs.ListLike[abs.BlockLike]
}

// This method returns the control expression for this select clause.
func (v *selectClause) GetControl() abs.ExpressionLike {
	return v.control
}

// This method sets the control expression for this select clause.
func (v *selectClause) SetControl(control abs.ExpressionLike) {
	if control == nil {
		panic("A select clause requires an control expression.")
	}
	v.control = control
}

// This method returns the option at the specified index from this select clause.
func (v *selectClause) GetOption(index int) abs.BlockLike {
	return v.options.GetItem(index)
}

// This method sets the option at the specified index for this select clause.
func (v *selectClause) SetOption(index int, option abs.BlockLike) {
	if option == nil {
		panic("Each option in a select clause requires a value.")
	}
	v.options.SetItem(index, option)
}

// This method returns the list of options for this select clause.
func (v *selectClause) GetOptions() abs.ListLike[abs.BlockLike] {
	return v.options
}

// This method sets the list of options for this select clause.
func (v *selectClause) SetOptions(options abs.ListLike[abs.BlockLike]) {
	if options == nil || options.IsEmpty() {
		panic("A select clause requires at least one option.")
	}
	v.options = options
}
