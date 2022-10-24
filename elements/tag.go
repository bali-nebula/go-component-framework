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
	var tag = uti.Base32Encode(bytes)
	return Tag(tag)
}

// This type defines the methods associated with a tag identifier that extends
// the native Go byte array type.
type Tag string
