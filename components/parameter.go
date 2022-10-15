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

// PARAMETER IMPLEMENTATION

// This constructor creates a new parameter.
func Parameter(name string, value any) abstractions.ParameterLike {
	if len(name) == 0 || name[0] != '$' {
		panic("A parameter name must start with '$'.")
	}
	return &parameter{name, value}
}

// This type defines the structure and methods associated with a context
// parameter.
type parameter struct {
	name  string
	value any
}

// This method returns the name of this parameter.
func (v *parameter) GetName() string {
	return v.name
}

// This method returns the value of this parameter.
func (v *parameter) GetValue() any {
	return v.value
}
