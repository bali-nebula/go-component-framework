/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package ranges_test

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	ran "github.com/bali-nebula/go-component-framework/v2/ranges"
	col "github.com/craterdog/go-collection-framework/v2"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestIntervalsWithDurations(t *tes.T) {
	var two = ele.DurationFromInt(2)
	var three = ele.DurationFromInt(3)
	var four = ele.DurationFromInt(4)
	var five = ele.DurationFromInt(5)
	var six = ele.DurationFromInt(6)
	var seven = ele.DurationFromInt(7)
	var eight = ele.DurationFromInt(8)
	var s = ran.Interval(three, abs.INCLUSIVE, seven)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.False(t, s.ContainsValue(two))
	ass.True(t, s.ContainsValue(three))
	ass.True(t, s.ContainsValue(five))
	ass.True(t, s.ContainsValue(seven))
	ass.False(t, s.ContainsValue(eight))
	ass.Equal(t, three, s.GetFirst())
	ass.Equal(t, abs.INCLUSIVE, s.GetExtent())
	ass.Equal(t, seven, s.GetLast())
	ass.Equal(t, five, s.GetValue(3))
	ass.Equal(t, 0, s.GetIndex(two))
	ass.Equal(t, 1, s.GetIndex(three))
	ass.Equal(t, 3, s.GetIndex(five))
	ass.Equal(t, 5, s.GetIndex(seven))
	ass.Equal(t, 0, s.GetIndex(eight))
	ass.Equal(t, []abs.Discrete{
		three,
		four,
		five,
		six,
		seven,
	}, s.AsArray())
	var iterator = col.Iterator(s)
	ass.Equal(t, three, iterator.GetNext())
	iterator.ToEnd()
	ass.Equal(t, seven, iterator.GetPrevious())
}

func TestIntervalsWithMoments(t *tes.T) {
	var two = ele.MomentFromInt(2)
	var three = ele.MomentFromInt(3)
	var four = ele.MomentFromInt(4)
	var five = ele.MomentFromInt(5)
	var six = ele.MomentFromInt(6)
	var seven = ele.MomentFromInt(7)
	var eight = ele.MomentFromInt(8)
	var s = ran.Interval(three, abs.RIGHT, seven)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 4, s.GetSize())
	ass.False(t, s.ContainsValue(two))
	ass.False(t, s.ContainsValue(three))
	ass.True(t, s.ContainsValue(five))
	ass.True(t, s.ContainsValue(seven))
	ass.False(t, s.ContainsValue(eight))
	ass.Equal(t, three, s.GetFirst())
	ass.Equal(t, abs.RIGHT, s.GetExtent())
	ass.Equal(t, seven, s.GetLast())
	ass.Equal(t, four, s.GetValue(1))
	ass.Equal(t, 0, s.GetIndex(three))
	ass.Equal(t, 1, s.GetIndex(four))
	ass.Equal(t, 2, s.GetIndex(five))
	ass.Equal(t, 4, s.GetIndex(seven))
	ass.Equal(t, 0, s.GetIndex(eight))
	ass.Equal(t, []abs.Discrete{
		four,
		five,
		six,
		seven,
	}, s.AsArray())
}

func TestIntervalsWithEmojis(t *tes.T) {
	var r1 = ele.Character('😀')
	var r2 = ele.Character('😆')
	var r3 = ele.Character('🤣')
	var s = ran.Interval(r1, abs.INCLUSIVE, r3)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 804, s.GetSize())
	ass.True(t, s.ContainsValue(r2))
	ass.Equal(t, r1, s.GetFirst())
	ass.Equal(t, abs.INCLUSIVE, s.GetExtent())
	ass.Equal(t, r3, s.GetLast())
	ass.Equal(t, r2, s.GetValue(7))
	ass.True(t, s.ContainsValue(r2))
}

func TestIntervalsWithCharacters(t *tes.T) {
	var rA = ele.Character('A')
	var ra = ele.Character('a')
	var rb = ele.Character('b')
	var rc = ele.Character('c')
	var rd = ele.Character('d')
	var re = ele.Character('e')
	var rf = ele.Character('f')
	var rg = ele.Character('g')
	var s = ran.Interval(ra, abs.LEFT, rf)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.False(t, s.ContainsValue(rA))
	ass.True(t, s.ContainsValue(ra))
	ass.True(t, s.ContainsValue(rc))
	ass.False(t, s.ContainsValue(rf))
	ass.False(t, s.ContainsValue(rg))
	ass.Equal(t, ra, s.GetFirst())
	ass.Equal(t, abs.LEFT, s.GetExtent())
	ass.Equal(t, rf, s.GetLast())
	ass.Equal(t, ra, s.GetValue(1))
	ass.Equal(t, 1, s.GetIndex(ra))
	ass.Equal(t, 4, s.GetIndex(rd))
	ass.Equal(t, 0, s.GetIndex(rf))
	ass.Equal(t, []abs.Discrete{ra, rb, rc, rd, re}, s.AsArray())
}

func TestIntervalsWithIntegers(t *tes.T) {
	var i1 = ele.Integer(1)
	var i2 = ele.Integer(2)
	var i3 = ele.Integer(3)
	var i4 = ele.Integer(4)
	var i5 = ele.Integer(5)
	var i6 = ele.Integer(6)
	var s = ran.Interval(i1, abs.RIGHT, i5)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 4, s.GetSize())
	ass.False(t, s.ContainsValue(i1))
	ass.True(t, s.ContainsValue(i2))
	ass.True(t, s.ContainsValue(i5))
	ass.False(t, s.ContainsValue(i6))
	ass.Equal(t, i1, s.GetFirst())
	ass.Equal(t, abs.RIGHT, s.GetExtent())
	ass.Equal(t, i5, s.GetLast())
	ass.Equal(t, i2, s.GetValue(1))
	ass.Equal(t, 0, s.GetIndex(i1))
	ass.Equal(t, 1, s.GetIndex(i2))
	ass.Equal(t, 4, s.GetIndex(i5))
	ass.Equal(t, 0, s.GetIndex(i6))
	ass.Equal(t, []abs.Discrete{i2, i3, i4, i5}, s.AsArray())
}
