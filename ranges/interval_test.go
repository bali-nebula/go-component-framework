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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	ele "github.com/bali-nebula/go-component-framework/elements"
	ran "github.com/bali-nebula/go-component-framework/ranges"
	col "github.com/craterdog/go-collection-framework/v2"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestIntervalsWithImplicitDurations(t *tes.T) {
	var s = ran.Interval[ele.Duration](3, abs.INCLUSIVE, 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.False(t, s.ContainsValue(2))
	ass.True(t, s.ContainsValue(3))
	ass.True(t, s.ContainsValue(5))
	ass.True(t, s.ContainsValue(7))
	ass.False(t, s.ContainsValue(8))
	ass.Equal(t, ele.Duration(3), s.GetFirst())
	ass.Equal(t, abs.INCLUSIVE, s.GetExtent())
	ass.Equal(t, ele.Duration(7), s.GetLast())
	ass.Equal(t, ele.Duration(5), s.GetValue(3))
	ass.Equal(t, 0, s.GetIndex(2))
	ass.Equal(t, 1, s.GetIndex(3))
	ass.Equal(t, 3, s.GetIndex(5))
	ass.Equal(t, 5, s.GetIndex(7))
	ass.Equal(t, 0, s.GetIndex(8))
	ass.Equal(t, []ele.Duration{3, 4, 5, 6, 7}, s.AsArray())
	var iterator = col.Iterator[ele.Duration](s)
	ass.Equal(t, ele.Duration(3), iterator.GetNext())
	iterator.ToEnd()
	ass.Equal(t, ele.Duration(7), iterator.GetPrevious())
}
func TestIntervalsWithExplicitDurations(t *tes.T) {
	var two = ele.DurationFromInteger(2)
	var three = ele.DurationFromInteger(3)
	var four = ele.DurationFromInteger(4)
	var five = ele.DurationFromInteger(5)
	var six = ele.DurationFromInteger(6)
	var seven = ele.DurationFromInteger(7)
	var eight = ele.DurationFromInteger(8)
	var s = ran.Interval[abs.DurationLike](three, abs.INCLUSIVE, seven)
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
	ass.Equal(t, []abs.DurationLike{
		three,
		four,
		five,
		six,
		seven,
	}, s.AsArray())
	var iterator = col.Iterator[abs.DurationLike](s)
	ass.Equal(t, three, iterator.GetNext())
	iterator.ToEnd()
	ass.Equal(t, seven, iterator.GetPrevious())
}

func TestIntervalsWithMoments(t *tes.T) {
	var two = ele.Moment(2)
	var three = ele.Moment(3)
	var four = ele.Moment(4)
	var five = ele.Moment(5)
	var six = ele.Moment(6)
	var seven = ele.Moment(7)
	var eight = ele.Moment(8)
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
	ass.Equal(t, []ele.Moment{
		four,
		five,
		six,
		seven,
	}, s.AsArray())

	s = ran.Interval(three, abs.EXCLUSIVE, seven)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 3, s.GetSize())
	ass.False(t, s.ContainsValue(two))
	ass.False(t, s.ContainsValue(three))
	ass.True(t, s.ContainsValue(five))
	ass.False(t, s.ContainsValue(seven))
	ass.False(t, s.ContainsValue(eight))
	ass.Equal(t, three, s.GetFirst())
	ass.Equal(t, abs.EXCLUSIVE, s.GetExtent())
	ass.Equal(t, seven, s.GetLast())
	ass.Equal(t, four, s.GetValue(1))
	ass.Equal(t, 0, s.GetIndex(three))
	ass.Equal(t, 1, s.GetIndex(four))
	ass.Equal(t, 3, s.GetIndex(six))
	ass.Equal(t, 0, s.GetIndex(seven))
	ass.Equal(t, []ele.Moment{
		four,
		five,
		six,
	}, s.AsArray())
}

func TestIntervalsWithRunes(t *tes.T) {
	var s = ran.Interval[ran.Rune]('a', abs.LEFT, 'f')
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.False(t, s.ContainsValue('A'))
	ass.True(t, s.ContainsValue('a'))
	ass.True(t, s.ContainsValue('c'))
	ass.False(t, s.ContainsValue('f'))
	ass.False(t, s.ContainsValue('g'))
	ass.Equal(t, ran.Rune('a'), s.GetFirst())
	ass.Equal(t, abs.LEFT, s.GetExtent())
	ass.Equal(t, ran.Rune('f'), s.GetLast())
	ass.Equal(t, ran.Rune('a'), s.GetValue(1))
	ass.Equal(t, 1, s.GetIndex('a'))
	ass.Equal(t, 4, s.GetIndex('d'))
	ass.Equal(t, 0, s.GetIndex('f'))
	ass.Equal(t, []ran.Rune{'a', 'b', 'c', 'd', 'e'}, s.AsArray())
}
