/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// FORMATTER INTERFACE

// This function returns a canonical BDN string for the specified entity.
func FormatEntity(entity abs.EntityLike) string {
	var v = Formatter(0)
	return v.FormatEntity(entity)
}

// This function returns a canonical BDN string for the specified component.
func FormatComponent(component abs.ComponentLike) string {
	var v = Formatter(0)
	return v.FormatComponent(component)
}

// This function returns a canonical BDN string for the specified component
// using the specified indentation.
func FormatComponentWithIndentation(component abs.ComponentLike, indentation int) string {
	var v = Formatter(indentation)
	return v.FormatComponent(component)
}

// FORMATTER IMPLEMENTATION

// This constructor creates a new formatter using the specified indentation.
func Formatter(indentation int) *formatter {
	var v = &formatter{abs.FormatterStateWithIndentation(indentation)}
	return v
}

// This type defines the structure and methods for a canonical formatting agent.
type formatter struct {
	state abs.FormatterStateLike
}

// This method returns the number of levels that each line is indented in the
// resulting canonical string.
func (v *formatter) GetIndentation() int {
	return v.state.GetIndentation()
}

// This method returns the canonical string for the specified entity.
func (v *formatter) FormatEntity(entity abs.EntityLike) string {
	v.formatEntity(entity)
	return v.state.GetResult()
}

// This method returns the canonical string for the specified component.
func (v *formatter) FormatComponent(component abs.ComponentLike) string {
	v.formatComponent(component)
	return v.state.GetResult()
}
