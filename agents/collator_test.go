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

	ass.True(t, age.CompareItems(nil, nil))
	ass.True(t, age.CompareItems(nil, ShouldBeNil))
	ass.True(t, age.CompareItems(ShouldBeNil, ShouldBeNil))
	ass.True(t, age.CompareItems(ShouldBeNil, nil))

	// Boolean
	var False = false
	var True = true
	var ShouldBeFalse bool

	ass.True(t, age.CompareItems(ShouldBeFalse, False))
	ass.False(t, age.CompareItems(True, ShouldBeFalse))

	ass.False(t, age.CompareItems(False, True))
	ass.True(t, age.CompareItems(False, False))
	ass.False(t, age.CompareItems(True, False))
	ass.True(t, age.CompareItems(True, True))

	// Byte
	var Zero byte = 0x00
	var One byte = 0x01
	var ShouldBeZero byte

	ass.True(t, age.CompareItems(ShouldBeZero, Zero))
	ass.False(t, age.CompareItems(One, ShouldBeZero))

	ass.False(t, age.CompareItems(Zero, One))
	ass.True(t, age.CompareItems(Zero, Zero))
	ass.False(t, age.CompareItems(One, Zero))
	ass.True(t, age.CompareItems(One, One))

	// Integer
	var Zilch = 0
	var Two = 2
	var Three = 3
	var ShouldBeZilch int

	ass.True(t, age.CompareItems(ShouldBeZilch, Zilch))
	ass.False(t, age.CompareItems(Two, ShouldBeZilch))

	ass.False(t, age.CompareItems(Two, Three))
	ass.True(t, age.CompareItems(Two, Two))
	ass.False(t, age.CompareItems(Three, Two))
	ass.True(t, age.CompareItems(Three, Three))

	// Float
	var Negligible = 0.0
	var Fourth = 0.25
	var Half = 0.5
	var ShouldBeNegligible float64

	ass.True(t, age.CompareItems(ShouldBeNegligible, Negligible))
	ass.False(t, age.CompareItems(Half, ShouldBeNegligible))

	ass.False(t, age.CompareItems(Fourth, Half))
	ass.True(t, age.CompareItems(Fourth, Fourth))
	ass.False(t, age.CompareItems(Half, Fourth))
	ass.True(t, age.CompareItems(Half, Half))

	// Complex
	var Origin = 0 + 0i
	var PiOver4 = 1 + 1i
	var PiOver2 = 1 + 0i
	var ShouldBeOrigin complex128

	ass.True(t, age.CompareItems(ShouldBeOrigin, Origin))
	ass.False(t, age.CompareItems(PiOver4, ShouldBeOrigin))

	ass.False(t, age.CompareItems(PiOver4, PiOver2))
	ass.True(t, age.CompareItems(PiOver4, PiOver4))
	ass.False(t, age.CompareItems(PiOver2, PiOver4))
	ass.True(t, age.CompareItems(PiOver2, PiOver2))

	// Rune
	var Null = rune(0)
	var Sad = '☹'
	var Happy = '☺'
	var ShouldBeNull rune

	ass.True(t, age.CompareItems(ShouldBeNull, Null))
	ass.False(t, age.CompareItems(Sad, ShouldBeNull))

	ass.False(t, age.CompareItems(Happy, Sad))
	ass.True(t, age.CompareItems(Happy, Happy))
	ass.False(t, age.CompareItems(Sad, Happy))
	ass.True(t, age.CompareItems(Sad, Sad))

	// String
	var Empty = ""
	var Hello = "Hello"
	var World = "World"
	var ShouldBeEmpty string

	ass.True(t, age.CompareItems(ShouldBeEmpty, Empty))
	ass.False(t, age.CompareItems(Hello, ShouldBeEmpty))

	ass.False(t, age.CompareItems(World, Hello))
	ass.True(t, age.CompareItems(World, World))
	ass.False(t, age.CompareItems(Hello, World))
	ass.True(t, age.CompareItems(Hello, Hello))

	// Array
	var Universe = "Universe"
	var a0 = []any{}
	var a1 = []any{Hello, World}
	var a2 = []any{Hello, Universe}
	var ShouldBeA0 []any

	ass.True(t, age.CompareItems(ShouldBeA0, a0))
	ass.False(t, age.CompareItems(a1, ShouldBeA0))

	ass.False(t, age.CompareItems(a1, a2))
	ass.True(t, age.CompareItems(a1, a1))
	ass.False(t, age.CompareItems(a2, a1))
	ass.True(t, age.CompareItems(a2, a2))

	// Map
	var m0 = map[any]any{}
	var m1 = map[any]any{
		One: True,
		Two: World}
	var m2 = map[any]any{
		One: True,
		Two: Hello}
	var ShouldBeM0 map[any]any

	ass.True(t, age.CompareItems(ShouldBeM0, m0))
	ass.False(t, age.CompareItems(m1, ShouldBeM0))

	ass.False(t, age.CompareItems(m1, m2))
	ass.True(t, age.CompareItems(m1, m1))
	ass.False(t, age.CompareItems(m2, m1))
	ass.True(t, age.CompareItems(m2, m2))
}

func TestRanking(t *tes.T) {
	// Nil
	var ShouldBeNil any

	ass.Equal(t, 0, age.RankItems(nil, nil))
	ass.Equal(t, 0, age.RankItems(nil, ShouldBeNil))
	ass.Equal(t, 0, age.RankItems(ShouldBeNil, ShouldBeNil))
	ass.Equal(t, 0, age.RankItems(ShouldBeNil, nil))

	// Boolean
	var False = false
	var True = true
	var ShouldBeFalse bool

	ass.Equal(t, 0, age.RankItems(ShouldBeFalse, ShouldBeFalse))
	ass.Equal(t, -1, age.RankItems(ShouldBeFalse, True))
	ass.Equal(t, 0, age.RankItems(False, ShouldBeFalse))
	ass.Equal(t, 1, age.RankItems(True, ShouldBeFalse))
	ass.Equal(t, 0, age.RankItems(ShouldBeFalse, False))
	ass.Equal(t, -1, age.RankItems(False, True))
	ass.Equal(t, 0, age.RankItems(False, False))
	ass.Equal(t, 1, age.RankItems(True, False))
	ass.Equal(t, 0, age.RankItems(True, True))

	// Byte
	var Zero byte = 0x00
	var One byte = 0x01
	var ShouldBeZero byte

	ass.Equal(t, 0, age.RankItems(ShouldBeZero, ShouldBeZero))
	ass.Equal(t, -1, age.RankItems(ShouldBeZero, One))
	ass.Equal(t, 0, age.RankItems(Zero, ShouldBeZero))
	ass.Equal(t, 1, age.RankItems(One, ShouldBeZero))
	ass.Equal(t, 0, age.RankItems(ShouldBeZero, Zero))
	ass.Equal(t, -1, age.RankItems(Zero, One))
	ass.Equal(t, 0, age.RankItems(Zero, Zero))
	ass.Equal(t, 1, age.RankItems(One, Zero))
	ass.Equal(t, 0, age.RankItems(One, One))

	// Integer
	var Zilch = 0
	var Two = 2
	var Three = 3
	var ShouldBeZilch int

	ass.Equal(t, 0, age.RankItems(ShouldBeZilch, ShouldBeZilch))
	ass.Equal(t, -1, age.RankItems(ShouldBeZilch, Two))
	ass.Equal(t, 0, age.RankItems(Zilch, ShouldBeZilch))
	ass.Equal(t, 1, age.RankItems(Two, ShouldBeZilch))
	ass.Equal(t, 0, age.RankItems(ShouldBeZilch, Zilch))
	ass.Equal(t, -1, age.RankItems(Two, Three))
	ass.Equal(t, 0, age.RankItems(Two, Two))
	ass.Equal(t, 1, age.RankItems(Three, Two))
	ass.Equal(t, 0, age.RankItems(Three, Three))

	// Float
	var Negligible = 0.0
	var Fourth = 0.25
	var Half = 0.5
	var ShouldBeNegligible float64

	ass.Equal(t, 0, age.RankItems(ShouldBeNegligible, ShouldBeNegligible))
	ass.Equal(t, -1, age.RankItems(ShouldBeNegligible, Half))
	ass.Equal(t, 0, age.RankItems(Negligible, ShouldBeNegligible))
	ass.Equal(t, 1, age.RankItems(Half, ShouldBeNegligible))
	ass.Equal(t, 0, age.RankItems(ShouldBeNegligible, Negligible))
	ass.Equal(t, -1, age.RankItems(Fourth, Half))
	ass.Equal(t, 0, age.RankItems(Fourth, Fourth))
	ass.Equal(t, 1, age.RankItems(Half, Fourth))
	ass.Equal(t, 0, age.RankItems(Half, Half))

	// Complex
	var Origin = 0 + 0i
	var PiOver4 = 1 + 1i
	var PiOver2 = 1 + 0i
	var ShouldBeOrigin complex128

	ass.Equal(t, 0, age.RankItems(ShouldBeOrigin, ShouldBeOrigin))
	ass.Equal(t, -1, age.RankItems(ShouldBeOrigin, PiOver4))
	ass.Equal(t, 0, age.RankItems(Origin, ShouldBeOrigin))
	ass.Equal(t, 1, age.RankItems(PiOver4, ShouldBeOrigin))
	ass.Equal(t, 0, age.RankItems(ShouldBeOrigin, Origin))
	ass.Equal(t, -1, age.RankItems(PiOver2, PiOver4))
	ass.Equal(t, 0, age.RankItems(PiOver2, PiOver2))
	ass.Equal(t, 1, age.RankItems(PiOver4, PiOver2))
	ass.Equal(t, 0, age.RankItems(PiOver4, PiOver4))

	// Rune
	var Null = rune(0)
	var Sad = '☹'
	var Happy = '☺'
	var ShouldBeNull rune

	ass.Equal(t, 0, age.RankItems(ShouldBeNull, ShouldBeNull))
	ass.Equal(t, -1, age.RankItems(ShouldBeNull, Sad))
	ass.Equal(t, 0, age.RankItems(Null, ShouldBeNull))
	ass.Equal(t, 1, age.RankItems(Sad, ShouldBeNull))
	ass.Equal(t, 0, age.RankItems(ShouldBeNull, Null))
	ass.Equal(t, -1, age.RankItems(Sad, Happy))
	ass.Equal(t, 0, age.RankItems(Sad, Sad))
	ass.Equal(t, 1, age.RankItems(Happy, Sad))
	ass.Equal(t, 0, age.RankItems(Happy, Happy))

	// String
	var Empty = ""
	var Hello = "Hello"
	var World = "World"
	var ShouldBeEmpty string

	ass.Equal(t, 0, age.RankItems(ShouldBeEmpty, ShouldBeEmpty))
	ass.Equal(t, -1, age.RankItems(ShouldBeEmpty, Hello))
	ass.Equal(t, 0, age.RankItems(Empty, ShouldBeEmpty))
	ass.Equal(t, 1, age.RankItems(Hello, ShouldBeEmpty))
	ass.Equal(t, 0, age.RankItems(ShouldBeEmpty, Empty))
	ass.Equal(t, -1, age.RankItems(Hello, World))
	ass.Equal(t, 0, age.RankItems(Hello, Hello))
	ass.Equal(t, 1, age.RankItems(World, Hello))
	ass.Equal(t, 0, age.RankItems(World, World))

	// Array
	var Universe = "Universe"
	var a0 = []any{}
	var a1 = []any{Hello, World}
	var a2 = []any{Hello, Universe}
	var a3 = []any{Hello, World, Universe}
	var a4 = []any{Hello, Universe, World}
	var ShouldBeA0 []any

	ass.Equal(t, 0, age.RankItems(ShouldBeA0, ShouldBeA0))
	ass.Equal(t, -1, age.RankItems(ShouldBeA0, a1))
	ass.Equal(t, 0, age.RankItems(a0, ShouldBeA0))
	ass.Equal(t, 1, age.RankItems(a1, ShouldBeA0))
	ass.Equal(t, 0, age.RankItems(ShouldBeA0, a0))
	ass.Equal(t, -1, age.RankItems(a2, a1))
	ass.Equal(t, 0, age.RankItems(a2, a2))
	ass.Equal(t, 1, age.RankItems(a1, a2))
	ass.Equal(t, 0, age.RankItems(a1, a1))
	ass.Equal(t, -1, age.RankItems(a2, a3))
	ass.Equal(t, 0, age.RankItems(a2, a2))
	ass.Equal(t, 1, age.RankItems(a3, a2))
	ass.Equal(t, 0, age.RankItems(a3, a3))
	ass.Equal(t, -1, age.RankItems(a4, a1))
	ass.Equal(t, 0, age.RankItems(a4, a4))
	ass.Equal(t, 1, age.RankItems(a1, a4))
	ass.Equal(t, 0, age.RankItems(a1, a1))

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

	ass.Equal(t, 0, age.RankItems(ShouldBeM0, ShouldBeM0))
	ass.Equal(t, -1, age.RankItems(ShouldBeM0, m1))
	ass.Equal(t, 0, age.RankItems(m0, ShouldBeM0))
	ass.Equal(t, 1, age.RankItems(m1, ShouldBeM0))
	ass.Equal(t, 0, age.RankItems(ShouldBeM0, m0))
	ass.Equal(t, -1, age.RankItems(m2, m1))
	ass.Equal(t, 0, age.RankItems(m2, m2))
	ass.Equal(t, 1, age.RankItems(m1, m2))
	ass.Equal(t, 0, age.RankItems(m1, m1))
	ass.Equal(t, -1, age.RankItems(m2, m3))
	ass.Equal(t, 0, age.RankItems(m2, m2))
	ass.Equal(t, 1, age.RankItems(m3, m2))
	ass.Equal(t, 0, age.RankItems(m3, m3))
	ass.Equal(t, -1, age.RankItems(m4, m1))
	ass.Equal(t, 0, age.RankItems(m4, m4))
	ass.Equal(t, 1, age.RankItems(m1, m4))
	ass.Equal(t, 0, age.RankItems(m1, m1))
}
