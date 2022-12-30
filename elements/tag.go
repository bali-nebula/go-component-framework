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
	uti "github.com/bali-nebula/go-component-framework/utilities"
)

// TAG INTERFACE

// This constructor creates a new random tag identifier element with the
// specified number of bytes.
func TagOfSize(size int) Tag {
	if size < 1 {
		panic("A tag must be at least one byte long!")
	}
	var bytes = uti.RandomBytes(size)
	var tag = uti.Base32Encode(bytes)
	return Tag(tag)
}

// This type defines the methods associated with a tag identifier that extends
// the native Go string type.
type Tag string

// SEQUENTIAL INTERFACE

// This method determines whether or not this tag is empty.
func (v Tag) IsEmpty() bool {
	return false
}

// This method returns the number of bytes contained in this tag.
func (v Tag) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the bytes in this tag. The bytes retrieved
// are in the same order as they are in the tag.
func (v Tag) AsArray() []byte {
	var encoded = string(v)
	var bytes = uti.Base64Decode(encoded)
	return bytes
}
