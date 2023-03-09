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
	Line        string
	Name        string
	Ordinal     uint
	Instruction uint16
	String      any
)

func InstructionFromBytes(firstByte, secondByte byte) Instruction {
	var v = Instruction((firstByte << 8) & secondByte)
	return v
}

func (v Instruction) GetFirstByte() byte {
	return byte(v >> 8)
}

func (v Instruction) GetSecondByte() byte {
	return byte(v)
}

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all quantized strings.
type Quantized interface {
	AsString() string
}

// CONSOLIDATED INTERFACES

type BinaryLike interface {
	Quantized
	Sequential[byte]
	Accessible[byte]
}

type BytecodeLike interface {
	Quantized
	Sequential[Instruction]
	Accessible[Instruction]
}

type MonikerLike interface {
	Quantized
	Sequential[Name]
	Accessible[Name]
}

type NarrativeLike interface {
	Quantized
	Sequential[Line]
	Accessible[Line]
}

type QuoteLike interface {
	Quantized
	Sequential[rune]
	Accessible[rune]
}

type VersionLike interface {
	Quantized
	Sequential[Ordinal]
	Accessible[Ordinal]
}
