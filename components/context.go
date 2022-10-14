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

// CONTEXT IMPLEMENTATION

// This constructor creates a new context.
func Context(parameters []abstractions.ParameterLike) abstractions.ContextLike {
	return &context{parameters}
}

// This type defines the structure and methods associated with a component
// context.
type context struct {
	parameters []abstractions.ParameterLike
}

// This method returns the names of the parameters in this context.
func (v *context) GetNames() []string {
	var names = make([]string, len(v.parameters))
	var index = 0
	for _, parameter := range v.parameters {
		names[index] = parameter.GetName()
		index++
	}
	return names
}

// This method returns the value of the specified parameter from this context
// or nil if it does not exist.
func (v *context) GetValue(name string) any {
	for _, parameter := range v.parameters {
		if parameter.GetName() == name {
			return parameter.GetValue
		}
	}
	return nil
}
