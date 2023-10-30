/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// TYPE DEFINITIONS

type (
	String      any
	Value       any
	Ordinal     uint
	Instruction uint16
	Line        string
	Name        string
)

func InstructionFromBytes(leftByte, rightByte byte) Instruction {
	var instruction uint16 = uint16(leftByte)
	instruction = instruction<<8 | uint16(rightByte)
	var v = Instruction(instruction)
	return v
}

func (v Instruction) GetLeftByte() byte {
	return byte(v >> 8)
}

func (v Instruction) GetRightByte() byte {
	return byte(v)
}

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all sequences of values.
type Sequential[V Value] interface {
	IsEmpty() bool
	GetSize() int
	AsArray() []V
}

// This interface defines the methods supported by all sequences whose values can
// be accessed using indices. The indices of an accessible sequence are ORDINAL
// rather than ZERO based (which is "SO last century"). This allows for positive
// indices starting at the beginning of the sequence, and negative indices
// starting at the end of the sequence as follows:
//
//	    1           2           3             N
//	[value 1] . [value 2] . [value 3] ... [value N]
//	   -N        -(N-1)      -(N-2)          -1
//
// Notice that because the indices are ordinal based, the positive and negative
// indices are symmetrical.
type Accessible[V Value] interface {
	GetValue(index int) V
	GetValues(first int, last int) Sequential[V]
}

// This interface defines the methods supported by all ratcheted agents that
// are capable of moving forward and backward over the values in a sequence. It
// is used to implement the GoF Iterator Pattern:
//   - https://en.wikipedia.org/wiki/Iterator_pattern
//
// A ratcheted agent locks into the slots that reside between each value in the
// sequence:
//
//	    [value 1] . [value 2] . [value 3] ... [value N]
//	  ^           ^           ^                         ^
//	slot 0      slot 1      slot 2                    slot N
//
// It moves from slot to slot and has access to the values (if they exist) on
// each side of the slot.
type Ratcheted[V Value] interface {
	GetSlot() int
	ToSlot(slot int)
	ToStart()
	ToEnd()
	HasPrevious() bool
	GetPrevious() V
	HasNext() bool
	GetNext() V
}

// CONSOLIDATED INTERFACES

type BinaryLike interface {
	Lexical
	Sequential[byte]
	Accessible[byte]
}

type BytecodeLike interface {
	Lexical
	Sequential[Instruction]
	Accessible[Instruction]
}

type MonikerLike interface {
	Lexical
	Sequential[Name]
	Accessible[Name]
}

type NarrativeLike interface {
	Lexical
	Sequential[Line]
	Accessible[Line]
}

type QuoteLike interface {
	Lexical
	Sequential[rune]
	Accessible[rune]
}

type SymbolLike interface {
	Lexical
	Sequential[rune]
}

type TagLike interface {
	Lexical
	Sequential[byte]
}

type VersionLike interface {
	Lexical
	Sequential[Ordinal]
	Accessible[Ordinal]
}

type ByteIteratorLike interface {
	Ratcheted[byte]
}

type InstructionIteratorLike interface {
	Ratcheted[Instruction]
}

type NameIteratorLike interface {
	Ratcheted[Name]
}

type LineIteratorLike interface {
	Ratcheted[Line]
}

type RuneIteratorLike interface {
	Ratcheted[rune]
}

type OrdinalIteratorLike interface {
	Ratcheted[Ordinal]
}
