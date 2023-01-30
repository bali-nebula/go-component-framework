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
	Line    string
	Name    string
	Ordinal uint
	Instruction uint16
	String  any
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all spectral strings.
type Spectral interface {
	AsString() string
}

// CONSOLIDATED INTERFACES

type BinaryLike interface {
	Spectral
	Sequential[byte]
	Accessible[byte]
}

type BytecodeLike interface {
	Spectral
	Sequential[Instruction]
	Accessible[Instruction]
}

type MonikerLike interface {
	Spectral
	Sequential[Name]
	Accessible[Name]
}

type NarrativeLike interface {
	Spectral
	Sequential[Line]
	Accessible[Line]
}

type QuoteLike interface {
	Spectral
	Sequential[rune]
	Accessible[rune]
}

type VersionLike interface {
	Spectral
	Sequential[Ordinal]
	Accessible[Ordinal]
}
