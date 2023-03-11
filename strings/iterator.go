/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// BYTE ITERATOR IMPLEMENTATION

// This constructor creates a new instance of a byte iterator that can be used
// to traverse the bytes in the specified sequence.
func ByteIterator(sequence abs.Sequential[abs.Byte]) abs.ByteIteratorLike {
	var iterator = col.Iterator[abs.Byte](sequence)
	return &byteIterator{iterator}
}

// This type defines the structure and methods for a byte iterator. The
// iterator operates on a sequence of bytes.
type byteIterator struct {
	abs.ByteIteratorLike
}

// INSTRUCTION ITERATOR IMPLEMENTATION

// This constructor creates a new instance of an instruction iterator that can be used
// to traverse the instructions in the specified sequence.
func InstructionIterator(sequence abs.Sequential[abs.Instruction]) abs.InstructionIteratorLike {
	var iterator = col.Iterator[abs.Instruction](sequence)
	return &instructionIterator{iterator}
}

// This type defines the structure and methods for an instruction iterator. The
// iterator operates on a sequence of instructions.
type instructionIterator struct {
	abs.InstructionIteratorLike
}

// NAME ITERATOR IMPLEMENTATION

// This constructor creates a new instance of a name iterator that can be used
// to traverse the names in the specified sequence.
func NameIterator(sequence abs.Sequential[abs.Name]) abs.NameIteratorLike {
	var iterator = col.Iterator[abs.Name](sequence)
	return &nameIterator{iterator}
}

// This type defines the structure and methods for a name iterator. The
// iterator operates on a sequence of names.
type nameIterator struct {
	abs.NameIteratorLike
}

// LINE ITERATOR IMPLEMENTATION

// This constructor creates a new instance of a line iterator that can be used
// to traverse the lines in the specified sequence.
func LineIterator(sequence abs.Sequential[abs.Line]) abs.LineIteratorLike {
	var iterator = col.Iterator[abs.Line](sequence)
	return &lineIterator{iterator}
}

// This type defines the structure and methods for a line iterator. The
// iterator operates on a sequence of lines.
type lineIterator struct {
	abs.LineIteratorLike
}

// RUNE ITERATOR IMPLEMENTATION

// This constructor creates a new instance of a rune iterator that can be used
// to traverse the runes in the specified sequence.
func RuneIterator(sequence abs.Sequential[abs.Rune]) abs.RuneIteratorLike {
	var iterator = col.Iterator[abs.Rune](sequence)
	return &runeIterator{iterator}
}

// This type defines the structure and methods for a rune iterator. The
// iterator operates on a sequence of runes.
type runeIterator struct {
	abs.RuneIteratorLike
}

// ORDINAL ITERATOR IMPLEMENTATION

// This constructor creates a new instance of an ordinal iterator that can be used
// to traverse the ordinals in the specified sequence.
func OrdinalIterator(sequence abs.Sequential[abs.Ordinal]) abs.OrdinalIteratorLike {
	var iterator = col.Iterator[abs.Ordinal](sequence)
	return &ordinalIterator{iterator}
}

// This type defines the structure and methods for an ordinal iterator. The
// iterator operates on a sequence of ordinals.
type ordinalIterator struct {
	abs.OrdinalIteratorLike
}
