/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents_test

import (
	age "github.com/craterdog-bali/go-bali-document-notation/agents"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestComparison(t *tes.T) {
	// Nil
	var ShouldBeNil any

	ass.True(t, age.CompareValues(nil, nil))
	ass.True(t, age.CompareValues(nil, ShouldBeNil))
	ass.True(t, age.CompareValues(ShouldBeNil, ShouldBeNil))
	ass.True(t, age.CompareValues(ShouldBeNil, nil))

	// Boolean
	var False = false
	var True = true
	var ShouldBeFalse bool

	ass.True(t, age.CompareValues(ShouldBeFalse, False))
	ass.False(t, age.CompareValues(True, ShouldBeFalse))

	ass.False(t, age.CompareValues(False, True))
	ass.True(t, age.CompareValues(False, False))
	ass.False(t, age.CompareValues(True, False))
	ass.True(t, age.CompareValues(True, True))

	// Byte
	var Zero byte = 0x00
	var One byte = 0x01
	var ShouldBeZero byte

	ass.True(t, age.CompareValues(ShouldBeZero, Zero))
	ass.False(t, age.CompareValues(One, ShouldBeZero))

	ass.False(t, age.CompareValues(Zero, One))
	ass.True(t, age.CompareValues(Zero, Zero))
	ass.False(t, age.CompareValues(One, Zero))
	ass.True(t, age.CompareValues(One, One))

	// Integer
	var Zilch = 0
	var Two = 2
	var Three = 3
	var ShouldBeZilch int

	ass.True(t, age.CompareValues(ShouldBeZilch, Zilch))
	ass.False(t, age.CompareValues(Two, ShouldBeZilch))

	ass.False(t, age.CompareValues(Two, Three))
	ass.True(t, age.CompareValues(Two, Two))
	ass.False(t, age.CompareValues(Three, Two))
	ass.True(t, age.CompareValues(Three, Three))

	// Float
	var Negligible = 0.0
	var Fourth = 0.25
	var Half = 0.5
	var ShouldBeNegligible float64

	ass.True(t, age.CompareValues(ShouldBeNegligible, Negligible))
	ass.False(t, age.CompareValues(Half, ShouldBeNegligible))

	ass.False(t, age.CompareValues(Fourth, Half))
	ass.True(t, age.CompareValues(Fourth, Fourth))
	ass.False(t, age.CompareValues(Half, Fourth))
	ass.True(t, age.CompareValues(Half, Half))

	// Complex
	var Origin = 0 + 0i
	var PiOver4 = 1 + 1i
	var PiOver2 = 1 + 0i
	var ShouldBeOrigin complex128

	ass.True(t, age.CompareValues(ShouldBeOrigin, Origin))
	ass.False(t, age.CompareValues(PiOver4, ShouldBeOrigin))

	ass.False(t, age.CompareValues(PiOver4, PiOver2))
	ass.True(t, age.CompareValues(PiOver4, PiOver4))
	ass.False(t, age.CompareValues(PiOver2, PiOver4))
	ass.True(t, age.CompareValues(PiOver2, PiOver2))

	// Rune
	var Null = rune(0)
	var Sad = '☹'
	var Happy = '☺'
	var ShouldBeNull rune

	ass.True(t, age.CompareValues(ShouldBeNull, Null))
	ass.False(t, age.CompareValues(Sad, ShouldBeNull))

	ass.False(t, age.CompareValues(Happy, Sad))
	ass.True(t, age.CompareValues(Happy, Happy))
	ass.False(t, age.CompareValues(Sad, Happy))
	ass.True(t, age.CompareValues(Sad, Sad))

	// String
	var Empty = ""
	var Hello = "Hello"
	var World = "World"
	var ShouldBeEmpty string

	ass.True(t, age.CompareValues(ShouldBeEmpty, Empty))
	ass.False(t, age.CompareValues(Hello, ShouldBeEmpty))

	ass.False(t, age.CompareValues(World, Hello))
	ass.True(t, age.CompareValues(World, World))
	ass.False(t, age.CompareValues(Hello, World))
	ass.True(t, age.CompareValues(Hello, Hello))

	// Array
	var Universe = "Universe"
	var a0 = []any{}
	var a1 = []any{Hello, World}
	var a2 = []any{Hello, Universe}
	var ShouldBeA0 []any

	ass.True(t, age.CompareValues(ShouldBeA0, a0))
	ass.False(t, age.CompareValues(a1, ShouldBeA0))

	ass.False(t, age.CompareValues(a1, a2))
	ass.True(t, age.CompareValues(a1, a1))
	ass.False(t, age.CompareValues(a2, a1))
	ass.True(t, age.CompareValues(a2, a2))

	// Map
	var m0 = map[any]any{}
	var m1 = map[any]any{
		One: True,
		Two: World}
	var m2 = map[any]any{
		One: True,
		Two: Hello}
	var ShouldBeM0 map[any]any

	ass.True(t, age.CompareValues(ShouldBeM0, m0))
	ass.False(t, age.CompareValues(m1, ShouldBeM0))

	ass.False(t, age.CompareValues(m1, m2))
	ass.True(t, age.CompareValues(m1, m1))
	ass.False(t, age.CompareValues(m2, m1))
	ass.True(t, age.CompareValues(m2, m2))
}

func TestRanking(t *tes.T) {
	// Nil
	var ShouldBeNil any

	ass.Equal(t, 0, age.RankValues(nil, nil))
	ass.Equal(t, 0, age.RankValues(nil, ShouldBeNil))
	ass.Equal(t, 0, age.RankValues(ShouldBeNil, ShouldBeNil))
	ass.Equal(t, 0, age.RankValues(ShouldBeNil, nil))

	// Boolean
	var False = false
	var True = true
	var ShouldBeFalse bool

	ass.Equal(t, 0, age.RankValues(ShouldBeFalse, ShouldBeFalse))
	ass.Equal(t, -1, age.RankValues(ShouldBeFalse, True))
	ass.Equal(t, 0, age.RankValues(False, ShouldBeFalse))
	ass.Equal(t, 1, age.RankValues(True, ShouldBeFalse))
	ass.Equal(t, 0, age.RankValues(ShouldBeFalse, False))
	ass.Equal(t, -1, age.RankValues(False, True))
	ass.Equal(t, 0, age.RankValues(False, False))
	ass.Equal(t, 1, age.RankValues(True, False))
	ass.Equal(t, 0, age.RankValues(True, True))

	// Byte
	var Zero byte = 0x00
	var One byte = 0x01
	var ShouldBeZero byte

	ass.Equal(t, 0, age.RankValues(ShouldBeZero, ShouldBeZero))
	ass.Equal(t, -1, age.RankValues(ShouldBeZero, One))
	ass.Equal(t, 0, age.RankValues(Zero, ShouldBeZero))
	ass.Equal(t, 1, age.RankValues(One, ShouldBeZero))
	ass.Equal(t, 0, age.RankValues(ShouldBeZero, Zero))
	ass.Equal(t, -1, age.RankValues(Zero, One))
	ass.Equal(t, 0, age.RankValues(Zero, Zero))
	ass.Equal(t, 1, age.RankValues(One, Zero))
	ass.Equal(t, 0, age.RankValues(One, One))

	// Integer
	var Zilch = 0
	var Two = 2
	var Three = 3
	var ShouldBeZilch int

	ass.Equal(t, 0, age.RankValues(ShouldBeZilch, ShouldBeZilch))
	ass.Equal(t, -1, age.RankValues(ShouldBeZilch, Two))
	ass.Equal(t, 0, age.RankValues(Zilch, ShouldBeZilch))
	ass.Equal(t, 1, age.RankValues(Two, ShouldBeZilch))
	ass.Equal(t, 0, age.RankValues(ShouldBeZilch, Zilch))
	ass.Equal(t, -1, age.RankValues(Two, Three))
	ass.Equal(t, 0, age.RankValues(Two, Two))
	ass.Equal(t, 1, age.RankValues(Three, Two))
	ass.Equal(t, 0, age.RankValues(Three, Three))

	// Float
	var Negligible = 0.0
	var Fourth = 0.25
	var Half = 0.5
	var ShouldBeNegligible float64

	ass.Equal(t, 0, age.RankValues(ShouldBeNegligible, ShouldBeNegligible))
	ass.Equal(t, -1, age.RankValues(ShouldBeNegligible, Half))
	ass.Equal(t, 0, age.RankValues(Negligible, ShouldBeNegligible))
	ass.Equal(t, 1, age.RankValues(Half, ShouldBeNegligible))
	ass.Equal(t, 0, age.RankValues(ShouldBeNegligible, Negligible))
	ass.Equal(t, -1, age.RankValues(Fourth, Half))
	ass.Equal(t, 0, age.RankValues(Fourth, Fourth))
	ass.Equal(t, 1, age.RankValues(Half, Fourth))
	ass.Equal(t, 0, age.RankValues(Half, Half))

	// Complex
	var Origin = 0 + 0i
	var PiOver4 = 1 + 1i
	var PiOver2 = 1 + 0i
	var ShouldBeOrigin complex128

	ass.Equal(t, 0, age.RankValues(ShouldBeOrigin, ShouldBeOrigin))
	ass.Equal(t, -1, age.RankValues(ShouldBeOrigin, PiOver4))
	ass.Equal(t, 0, age.RankValues(Origin, ShouldBeOrigin))
	ass.Equal(t, 1, age.RankValues(PiOver4, ShouldBeOrigin))
	ass.Equal(t, 0, age.RankValues(ShouldBeOrigin, Origin))
	ass.Equal(t, -1, age.RankValues(PiOver2, PiOver4))
	ass.Equal(t, 0, age.RankValues(PiOver2, PiOver2))
	ass.Equal(t, 1, age.RankValues(PiOver4, PiOver2))
	ass.Equal(t, 0, age.RankValues(PiOver4, PiOver4))

	// Rune
	var Null = rune(0)
	var Sad = '☹'
	var Happy = '☺'
	var ShouldBeNull rune

	ass.Equal(t, 0, age.RankValues(ShouldBeNull, ShouldBeNull))
	ass.Equal(t, -1, age.RankValues(ShouldBeNull, Sad))
	ass.Equal(t, 0, age.RankValues(Null, ShouldBeNull))
	ass.Equal(t, 1, age.RankValues(Sad, ShouldBeNull))
	ass.Equal(t, 0, age.RankValues(ShouldBeNull, Null))
	ass.Equal(t, -1, age.RankValues(Sad, Happy))
	ass.Equal(t, 0, age.RankValues(Sad, Sad))
	ass.Equal(t, 1, age.RankValues(Happy, Sad))
	ass.Equal(t, 0, age.RankValues(Happy, Happy))

	// String
	var Empty = ""
	var Hello = "Hello"
	var World = "World"
	var ShouldBeEmpty string

	ass.Equal(t, 0, age.RankValues(ShouldBeEmpty, ShouldBeEmpty))
	ass.Equal(t, -1, age.RankValues(ShouldBeEmpty, Hello))
	ass.Equal(t, 0, age.RankValues(Empty, ShouldBeEmpty))
	ass.Equal(t, 1, age.RankValues(Hello, ShouldBeEmpty))
	ass.Equal(t, 0, age.RankValues(ShouldBeEmpty, Empty))
	ass.Equal(t, -1, age.RankValues(Hello, World))
	ass.Equal(t, 0, age.RankValues(Hello, Hello))
	ass.Equal(t, 1, age.RankValues(World, Hello))
	ass.Equal(t, 0, age.RankValues(World, World))

	// Array
	var Universe = "Universe"
	var a0 = []any{}
	var a1 = []any{Hello, World}
	var a2 = []any{Hello, Universe}
	var a3 = []any{Hello, World, Universe}
	var a4 = []any{Hello, Universe, World}
	var ShouldBeA0 []any

	ass.Equal(t, 0, age.RankValues(ShouldBeA0, ShouldBeA0))
	ass.Equal(t, -1, age.RankValues(ShouldBeA0, a1))
	ass.Equal(t, 0, age.RankValues(a0, ShouldBeA0))
	ass.Equal(t, 1, age.RankValues(a1, ShouldBeA0))
	ass.Equal(t, 0, age.RankValues(ShouldBeA0, a0))
	ass.Equal(t, -1, age.RankValues(a2, a1))
	ass.Equal(t, 0, age.RankValues(a2, a2))
	ass.Equal(t, 1, age.RankValues(a1, a2))
	ass.Equal(t, 0, age.RankValues(a1, a1))
	ass.Equal(t, -1, age.RankValues(a2, a3))
	ass.Equal(t, 0, age.RankValues(a2, a2))
	ass.Equal(t, 1, age.RankValues(a3, a2))
	ass.Equal(t, 0, age.RankValues(a3, a3))
	ass.Equal(t, -1, age.RankValues(a4, a1))
	ass.Equal(t, 0, age.RankValues(a4, a4))
	ass.Equal(t, 1, age.RankValues(a1, a4))
	ass.Equal(t, 0, age.RankValues(a1, a1))

	// Map
	var m0 = map[any]any{}
	var m1 = map[any]any{
		One: True,
		Two: World}
	var m2 = map[any]any{
		One: True,
		Two: Hello}
	var m3 = map[any]any{
		One:   True,
		Two:   World,
		Three: Universe}
	var m4 = map[any]any{
		One:   True,
		Two:   Universe,
		Three: World}
	var ShouldBeM0 map[any]any

	ass.Equal(t, 0, age.RankValues(ShouldBeM0, ShouldBeM0))
	ass.Equal(t, -1, age.RankValues(ShouldBeM0, m1))
	ass.Equal(t, 0, age.RankValues(m0, ShouldBeM0))
	ass.Equal(t, 1, age.RankValues(m1, ShouldBeM0))
	ass.Equal(t, 0, age.RankValues(ShouldBeM0, m0))
	ass.Equal(t, -1, age.RankValues(m2, m1))
	ass.Equal(t, 0, age.RankValues(m2, m2))
	ass.Equal(t, 1, age.RankValues(m1, m2))
	ass.Equal(t, 0, age.RankValues(m1, m1))
	ass.Equal(t, -1, age.RankValues(m2, m3))
	ass.Equal(t, 0, age.RankValues(m2, m2))
	ass.Equal(t, 1, age.RankValues(m3, m2))
	ass.Equal(t, 0, age.RankValues(m3, m3))
	ass.Equal(t, -1, age.RankValues(m4, m1))
	ass.Equal(t, 0, age.RankValues(m4, m4))
	ass.Equal(t, 1, age.RankValues(m1, m4))
	ass.Equal(t, 0, age.RankValues(m1, m1))
}
