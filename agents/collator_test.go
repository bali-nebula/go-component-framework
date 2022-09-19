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
	"github.com/craterdog-bali/go-bali-document-notation/agents"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComparison(t *testing.T) {
	// Nil
	var ShouldBeNil any

	assert.True(t, agents.CompareValues(nil, nil))
	assert.True(t, agents.CompareValues(nil, ShouldBeNil))
	assert.True(t, agents.CompareValues(ShouldBeNil, ShouldBeNil))
	assert.True(t, agents.CompareValues(ShouldBeNil, nil))

	// Boolean
	False := false
	True := true
	var ShouldBeFalse bool

	assert.True(t, agents.CompareValues(ShouldBeFalse, False))
	assert.False(t, agents.CompareValues(True, ShouldBeFalse))

	assert.False(t, agents.CompareValues(False, True))
	assert.True(t, agents.CompareValues(False, False))
	assert.False(t, agents.CompareValues(True, False))
	assert.True(t, agents.CompareValues(True, True))

	// Byte
	var Zero byte = 0x00
	var One byte = 0x01
	var ShouldBeZero byte

	assert.True(t, agents.CompareValues(ShouldBeZero, Zero))
	assert.False(t, agents.CompareValues(One, ShouldBeZero))

	assert.False(t, agents.CompareValues(Zero, One))
	assert.True(t, agents.CompareValues(Zero, Zero))
	assert.False(t, agents.CompareValues(One, Zero))
	assert.True(t, agents.CompareValues(One, One))

	// Integer
	Zilch := 0
	Two := 2
	Three := 3
	var ShouldBeZilch int

	assert.True(t, agents.CompareValues(ShouldBeZilch, Zilch))
	assert.False(t, agents.CompareValues(Two, ShouldBeZilch))

	assert.False(t, agents.CompareValues(Two, Three))
	assert.True(t, agents.CompareValues(Two, Two))
	assert.False(t, agents.CompareValues(Three, Two))
	assert.True(t, agents.CompareValues(Three, Three))

	// Float
	Negligible := 0.0
	Fourth := 0.25
	Half := 0.5
	var ShouldBeNegligible float64

	assert.True(t, agents.CompareValues(ShouldBeNegligible, Negligible))
	assert.False(t, agents.CompareValues(Half, ShouldBeNegligible))

	assert.False(t, agents.CompareValues(Fourth, Half))
	assert.True(t, agents.CompareValues(Fourth, Fourth))
	assert.False(t, agents.CompareValues(Half, Fourth))
	assert.True(t, agents.CompareValues(Half, Half))

	// Complex
	Origin := 0 + 0i
	PiOver4 := 1 + 1i
	PiOver2 := 1 + 0i
	var ShouldBeOrigin complex128

	assert.True(t, agents.CompareValues(ShouldBeOrigin, Origin))
	assert.False(t, agents.CompareValues(PiOver4, ShouldBeOrigin))

	assert.False(t, agents.CompareValues(PiOver4, PiOver2))
	assert.True(t, agents.CompareValues(PiOver4, PiOver4))
	assert.False(t, agents.CompareValues(PiOver2, PiOver4))
	assert.True(t, agents.CompareValues(PiOver2, PiOver2))

	// Rune
	Null := rune(0)
	Sad := '☹'
	Happy := '☺'
	var ShouldBeNull rune

	assert.True(t, agents.CompareValues(ShouldBeNull, Null))
	assert.False(t, agents.CompareValues(Sad, ShouldBeNull))

	assert.False(t, agents.CompareValues(Happy, Sad))
	assert.True(t, agents.CompareValues(Happy, Happy))
	assert.False(t, agents.CompareValues(Sad, Happy))
	assert.True(t, agents.CompareValues(Sad, Sad))

	// String
	Empty := ""
	Hello := "Hello"
	World := "World"
	var ShouldBeEmpty string

	assert.True(t, agents.CompareValues(ShouldBeEmpty, Empty))
	assert.False(t, agents.CompareValues(Hello, ShouldBeEmpty))

	assert.False(t, agents.CompareValues(World, Hello))
	assert.True(t, agents.CompareValues(World, World))
	assert.False(t, agents.CompareValues(Hello, World))
	assert.True(t, agents.CompareValues(Hello, Hello))

	// Array
	Universe := "Universe"
	a0 := []any{}
	a1 := []any{Hello, World}
	a2 := []any{Hello, Universe}
	var ShouldBeA0 []any

	assert.True(t, agents.CompareValues(ShouldBeA0, a0))
	assert.False(t, agents.CompareValues(a1, ShouldBeA0))

	assert.False(t, agents.CompareValues(a1, a2))
	assert.True(t, agents.CompareValues(a1, a1))
	assert.False(t, agents.CompareValues(a2, a1))
	assert.True(t, agents.CompareValues(a2, a2))

	// Map
	m0 := map[any]any{}
	m1 := map[any]any{
		One: True,
		Two: World}
	m2 := map[any]any{
		One: True,
		Two: Hello}
	var ShouldBeM0 map[any]any

	assert.True(t, agents.CompareValues(ShouldBeM0, m0))
	assert.False(t, agents.CompareValues(m1, ShouldBeM0))

	assert.False(t, agents.CompareValues(m1, m2))
	assert.True(t, agents.CompareValues(m1, m1))
	assert.False(t, agents.CompareValues(m2, m1))
	assert.True(t, agents.CompareValues(m2, m2))
}

func TestRanking(t *testing.T) {
	// Nil
	var ShouldBeNil any

	assert.Equal(t, 0, agents.RankValues(nil, nil))
	assert.Equal(t, 0, agents.RankValues(nil, ShouldBeNil))
	assert.Equal(t, 0, agents.RankValues(ShouldBeNil, ShouldBeNil))
	assert.Equal(t, 0, agents.RankValues(ShouldBeNil, nil))

	// Boolean
	False := false
	True := true
	var ShouldBeFalse bool

	assert.Equal(t, 0, agents.RankValues(ShouldBeFalse, ShouldBeFalse))
	assert.Equal(t, -1, agents.RankValues(ShouldBeFalse, True))
	assert.Equal(t, 0, agents.RankValues(False, ShouldBeFalse))
	assert.Equal(t, 1, agents.RankValues(True, ShouldBeFalse))
	assert.Equal(t, 0, agents.RankValues(ShouldBeFalse, False))
	assert.Equal(t, -1, agents.RankValues(False, True))
	assert.Equal(t, 0, agents.RankValues(False, False))
	assert.Equal(t, 1, agents.RankValues(True, False))
	assert.Equal(t, 0, agents.RankValues(True, True))

	// Byte
	var Zero byte = 0x00
	var One byte = 0x01
	var ShouldBeZero byte

	assert.Equal(t, 0, agents.RankValues(ShouldBeZero, ShouldBeZero))
	assert.Equal(t, -1, agents.RankValues(ShouldBeZero, One))
	assert.Equal(t, 0, agents.RankValues(Zero, ShouldBeZero))
	assert.Equal(t, 1, agents.RankValues(One, ShouldBeZero))
	assert.Equal(t, 0, agents.RankValues(ShouldBeZero, Zero))
	assert.Equal(t, -1, agents.RankValues(Zero, One))
	assert.Equal(t, 0, agents.RankValues(Zero, Zero))
	assert.Equal(t, 1, agents.RankValues(One, Zero))
	assert.Equal(t, 0, agents.RankValues(One, One))

	// Integer
	Zilch := 0
	Two := 2
	Three := 3
	var ShouldBeZilch int

	assert.Equal(t, 0, agents.RankValues(ShouldBeZilch, ShouldBeZilch))
	assert.Equal(t, -1, agents.RankValues(ShouldBeZilch, Two))
	assert.Equal(t, 0, agents.RankValues(Zilch, ShouldBeZilch))
	assert.Equal(t, 1, agents.RankValues(Two, ShouldBeZilch))
	assert.Equal(t, 0, agents.RankValues(ShouldBeZilch, Zilch))
	assert.Equal(t, -1, agents.RankValues(Two, Three))
	assert.Equal(t, 0, agents.RankValues(Two, Two))
	assert.Equal(t, 1, agents.RankValues(Three, Two))
	assert.Equal(t, 0, agents.RankValues(Three, Three))

	// Float
	Negligible := 0.0
	Fourth := 0.25
	Half := 0.5
	var ShouldBeNegligible float64

	assert.Equal(t, 0, agents.RankValues(ShouldBeNegligible, ShouldBeNegligible))
	assert.Equal(t, -1, agents.RankValues(ShouldBeNegligible, Half))
	assert.Equal(t, 0, agents.RankValues(Negligible, ShouldBeNegligible))
	assert.Equal(t, 1, agents.RankValues(Half, ShouldBeNegligible))
	assert.Equal(t, 0, agents.RankValues(ShouldBeNegligible, Negligible))
	assert.Equal(t, -1, agents.RankValues(Fourth, Half))
	assert.Equal(t, 0, agents.RankValues(Fourth, Fourth))
	assert.Equal(t, 1, agents.RankValues(Half, Fourth))
	assert.Equal(t, 0, agents.RankValues(Half, Half))

	// Complex
	Origin := 0 + 0i
	PiOver4 := 1 + 1i
	PiOver2 := 1 + 0i
	var ShouldBeOrigin complex128

	assert.Equal(t, 0, agents.RankValues(ShouldBeOrigin, ShouldBeOrigin))
	assert.Equal(t, -1, agents.RankValues(ShouldBeOrigin, PiOver4))
	assert.Equal(t, 0, agents.RankValues(Origin, ShouldBeOrigin))
	assert.Equal(t, 1, agents.RankValues(PiOver4, ShouldBeOrigin))
	assert.Equal(t, 0, agents.RankValues(ShouldBeOrigin, Origin))
	assert.Equal(t, -1, agents.RankValues(PiOver2, PiOver4))
	assert.Equal(t, 0, agents.RankValues(PiOver2, PiOver2))
	assert.Equal(t, 1, agents.RankValues(PiOver4, PiOver2))
	assert.Equal(t, 0, agents.RankValues(PiOver4, PiOver4))

	// Rune
	Null := rune(0)
	Sad := '☹'
	Happy := '☺'
	var ShouldBeNull rune

	assert.Equal(t, 0, agents.RankValues(ShouldBeNull, ShouldBeNull))
	assert.Equal(t, -1, agents.RankValues(ShouldBeNull, Sad))
	assert.Equal(t, 0, agents.RankValues(Null, ShouldBeNull))
	assert.Equal(t, 1, agents.RankValues(Sad, ShouldBeNull))
	assert.Equal(t, 0, agents.RankValues(ShouldBeNull, Null))
	assert.Equal(t, -1, agents.RankValues(Sad, Happy))
	assert.Equal(t, 0, agents.RankValues(Sad, Sad))
	assert.Equal(t, 1, agents.RankValues(Happy, Sad))
	assert.Equal(t, 0, agents.RankValues(Happy, Happy))

	// String
	Empty := ""
	Hello := "Hello"
	World := "World"
	var ShouldBeEmpty string

	assert.Equal(t, 0, agents.RankValues(ShouldBeEmpty, ShouldBeEmpty))
	assert.Equal(t, -1, agents.RankValues(ShouldBeEmpty, Hello))
	assert.Equal(t, 0, agents.RankValues(Empty, ShouldBeEmpty))
	assert.Equal(t, 1, agents.RankValues(Hello, ShouldBeEmpty))
	assert.Equal(t, 0, agents.RankValues(ShouldBeEmpty, Empty))
	assert.Equal(t, -1, agents.RankValues(Hello, World))
	assert.Equal(t, 0, agents.RankValues(Hello, Hello))
	assert.Equal(t, 1, agents.RankValues(World, Hello))
	assert.Equal(t, 0, agents.RankValues(World, World))

	// Array
	Universe := "Universe"
	a0 := []any{}
	a1 := []any{Hello, World}
	a2 := []any{Hello, Universe}
	a3 := []any{Hello, World, Universe}
	a4 := []any{Hello, Universe, World}
	var ShouldBeA0 []any

	assert.Equal(t, 0, agents.RankValues(ShouldBeA0, ShouldBeA0))
	assert.Equal(t, -1, agents.RankValues(ShouldBeA0, a1))
	assert.Equal(t, 0, agents.RankValues(a0, ShouldBeA0))
	assert.Equal(t, 1, agents.RankValues(a1, ShouldBeA0))
	assert.Equal(t, 0, agents.RankValues(ShouldBeA0, a0))
	assert.Equal(t, -1, agents.RankValues(a2, a1))
	assert.Equal(t, 0, agents.RankValues(a2, a2))
	assert.Equal(t, 1, agents.RankValues(a1, a2))
	assert.Equal(t, 0, agents.RankValues(a1, a1))
	assert.Equal(t, -1, agents.RankValues(a2, a3))
	assert.Equal(t, 0, agents.RankValues(a2, a2))
	assert.Equal(t, 1, agents.RankValues(a3, a2))
	assert.Equal(t, 0, agents.RankValues(a3, a3))
	assert.Equal(t, -1, agents.RankValues(a4, a1))
	assert.Equal(t, 0, agents.RankValues(a4, a4))
	assert.Equal(t, 1, agents.RankValues(a1, a4))
	assert.Equal(t, 0, agents.RankValues(a1, a1))

	// Map
	m0 := map[any]any{}
	m1 := map[any]any{
		One: True,
		Two: World}
	m2 := map[any]any{
		One: True,
		Two: Hello}
	m3 := map[any]any{
		One:   True,
		Two:   World,
		Three: Universe}
	m4 := map[any]any{
		One:   True,
		Two:   Universe,
		Three: World}
	var ShouldBeM0 map[any]any

	assert.Equal(t, 0, agents.RankValues(ShouldBeM0, ShouldBeM0))
	assert.Equal(t, -1, agents.RankValues(ShouldBeM0, m1))
	assert.Equal(t, 0, agents.RankValues(m0, ShouldBeM0))
	assert.Equal(t, 1, agents.RankValues(m1, ShouldBeM0))
	assert.Equal(t, 0, agents.RankValues(ShouldBeM0, m0))
	assert.Equal(t, -1, agents.RankValues(m2, m1))
	assert.Equal(t, 0, agents.RankValues(m2, m2))
	assert.Equal(t, 1, agents.RankValues(m1, m2))
	assert.Equal(t, 0, agents.RankValues(m1, m1))
	assert.Equal(t, -1, agents.RankValues(m2, m3))
	assert.Equal(t, 0, agents.RankValues(m2, m2))
	assert.Equal(t, 1, agents.RankValues(m3, m2))
	assert.Equal(t, 0, agents.RankValues(m3, m3))
	assert.Equal(t, -1, agents.RankValues(m4, m1))
	assert.Equal(t, 0, agents.RankValues(m4, m4))
	assert.Equal(t, 1, agents.RankValues(m1, m4))
	assert.Equal(t, 0, agents.RankValues(m1, m1))
}
