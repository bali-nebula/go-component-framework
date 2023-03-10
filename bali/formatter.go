/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	sts "strings"
)

// FORMATTER INTERFACE

// This function returns a canonical BDN string for the specified entity.
func FormatEntity(entity abs.Entity) string {
	var v = Formatter(0)
	return v.FormatEntity(entity)
}

// This function returns a canonical BDN string for the specified entity
// using the specified indentation.
func FormatEntityWithIndentation(entity abs.Entity, indentation int) string {
	var v = Formatter(indentation)
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

// This function returns a canonical BDN bytes for the specified component
// including the POSIX standard trailing EOL.
func FormatDocument(component abs.ComponentLike) []byte {
	var v = Formatter(0)
	var s = v.FormatComponent(component) + EOL
	return []byte(s)
}

// FORMATTER IMPLEMENTATION

// This constructor creates a new formatter using the specified indentation.
func Formatter(indentation int) *formatter {
	var v = &formatter{indentation: indentation, depth: 0}
	return v
}

// This type defines the structure and methods for a canonical formatting agent.
type formatter struct {
	indentation int
	depth       int
	result      sts.Builder
}

// This method returns the number of levels that each line is indented in the
// resulting canonical string.
func (v *formatter) GetIndentation() int {
	return v.indentation
}

// This method appends the specified string to the result.
func (v *formatter) AppendString(string_ string) {
	v.result.WriteString(string_)
}

// This method appends a properly indented newline to the result.
func (v *formatter) AppendNewline() {
	var separator = EOL
	var levels = v.depth + v.indentation
	for level := 0; level < levels; level++ {
		separator += "    "
	}
	v.result.WriteString(separator)
}

// This method returns the canonically formatted string result.
func (v *formatter) GetResult() string {
	var result = v.result.String()
	v.result.Reset()
	return result
}

// This method returns the canonical string for the specified entity.
func (v *formatter) FormatEntity(entity abs.Entity) string {
	v.formatEntity(entity)
	return v.GetResult()
}

// This method returns the canonical string for the specified component.
func (v *formatter) FormatComponent(component abs.ComponentLike) string {
	v.formatComponent(component)
	return v.GetResult()
}
