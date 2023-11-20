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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// BYTE ITERATOR IMPLEMENTATION

// This constructor creates a new instance of a byte iterator that can be used
// to traverse the bytes in the specified sequence.
func ByteIterator(sequence abs.Sequential[byte]) abs.ByteIteratorLike {
	var iterator = col.Iterator[byte](sequence)
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

// IDENTIFIER ITERATOR IMPLEMENTATION

// This constructor creates a new instance of a identifier iterator that can be
// used to traverse the identifiers in the specified sequence.
func IdentifierIterator(sequence abs.Sequential[abs.Identifier]) abs.IdentifierIteratorLike {
	var iterator = col.Iterator[abs.Identifier](sequence)
	return &identifierIterator{iterator}
}

// This type defines the structure and methods for a identifier iterator. The
// iterator operates on a sequence of identifiers.
type identifierIterator struct {
	abs.IdentifierIteratorLike
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
func RuneIterator(sequence abs.Sequential[rune]) abs.RuneIteratorLike {
	var iterator = col.Iterator[rune](sequence)
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
