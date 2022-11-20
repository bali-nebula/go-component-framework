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
	sts "strings"
)

// COMMENT IMPLEMENTATION

// This type defines the methods associated with a comment annotation that
// extends the native Go string type and represents the string of runes that
// make up the comment.
type Comment string

// This method determines whether or not this string is empty.
func (v Comment) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of lines contained in this comment.
func (v Comment) GetSize() int {
	return len(v.AsArray())
}

// This method returns all the lines in this comment. The lines retrieved
// are in the same order as they are in the comment.
func (v Comment) AsArray() []string {
	var array = sts.Split(string(v), "\n")
	return array
}
