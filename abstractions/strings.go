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
	Byte    byte
	Line    string
	Name    string
	Ordinal uint
	Rune    rune
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
	Sequential[Byte]
	Accessible[Byte]
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
	Sequential[Rune]
	Accessible[Rune]
}

type VersionLike interface {
	Spectral
	Sequential[Ordinal]
	Accessible[Ordinal]
}
