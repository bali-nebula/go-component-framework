/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	uti "github.com/craterdog-bali/go-bali-document-notation/utilities"
)

// TAG INTERFACE

// This constructor creates a new random tag identifier element with the
// specified number of bytes.
func TagOfSize(size int) Tag {
	if size < 1 {
		panic("A tag must be at least one byte long!")
	}
	var bytes = uti.RandomBytes(size)
	var tag = "#" + uti.Base32Encode(bytes)
	return Tag(tag)
}

// This constructor attempts to create a new identifier tag from the specified
// formatted string. It returns a tag value and whether or not the string
// contained a valid base 32 encoded string.
// For valid string formats for this type see `../abstractions/language.go`.
func TagFromString(v string) (Tag, bool) {
	var tag, ok = stringToTag(v)
	return Tag(tag), ok
}

// This type defines the methods associated with a tag identifier that extends
// the native Go byte array type.
type Tag string

// PRIVATE FUNCTIONS

// This function parses a identifier tag string and returns the corresponding
// base 32 encoded string for the tag.
func stringToTag(v string) (string, bool) {
	var tag string
	var ok = true
	var matches = abs.ScanTag([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		tag = matches[0]
	}
	return tag, ok
}
