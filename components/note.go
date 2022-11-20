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

// NOTE IMPLEMENTATION

// This type defines the methods associated with a note annotation that
// extends the native Go string type and represents the string of runes that
// make up the note.
type Note string

// This method determines whether or not this note is empty.
func (v Note) IsEmpty() bool {
	return len(v) == 0
}

// This method returns the number of runes contained in this note.
func (v Note) GetSize() int {
	return len(v)
}

// This method returns all the runes in this note. The runes retrieved
// are in the same order as they are in the note.
func (v Note) AsArray() []rune {
	return []rune(v)
}
