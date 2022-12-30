/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package temporary_test

import (
	tem "github.com/bali-nebula/go-component-framework/temporary"
	ele "github.com/bali-nebula/go-component-framework/elements"
	co2 "github.com/craterdog/go-collection-framework"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestIntervalsWithDurations(t *tes.T) {
	var s = tem.Interval[ele.Duration](3, tem.INCLUSIVE, 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.Equal(t, ele.Duration(5), s.GetValue(3))
	ass.Equal(t, 0, s.GetIndex(2))
	ass.Equal(t, 1, s.GetIndex(3))
	ass.Equal(t, 3, s.GetIndex(5))
	ass.Equal(t, 5, s.GetIndex(7))
	ass.Equal(t, 0, s.GetIndex(8))
	ass.Equal(t, []ele.Duration{3, 4, 5, 6, 7}, s.AsArray())
	var iterator = co2.Iterator[ele.Duration](s)
	ass.Equal(t, ele.Duration(3), iterator.GetNext())
	iterator.ToEnd()
	ass.Equal(t, ele.Duration(7), iterator.GetPrevious())
}

func TestIntervalsWithMoments(t *tes.T) {
	var s = tem.Interval[ele.Moment](3, tem.RIGHT, 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 4, s.GetSize())
	ass.Equal(t, ele.Moment(4), s.GetValue(1))
	ass.Equal(t, 0, s.GetIndex(3))
	ass.Equal(t, 1, s.GetIndex(4))
	ass.Equal(t, 2, s.GetIndex(5))
	ass.Equal(t, 4, s.GetIndex(7))
	ass.Equal(t, 0, s.GetIndex(8))
	ass.Equal(t, []ele.Moment{4, 5, 6, 7}, s.AsArray())

	s = tem.Interval[ele.Moment](3, tem.EXCLUSIVE, 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 3, s.GetSize())
	ass.Equal(t, ele.Moment(4), s.GetValue(1))
	ass.Equal(t, 0, s.GetIndex(3))
	ass.Equal(t, 1, s.GetIndex(4))
	ass.Equal(t, 3, s.GetIndex(6))
	ass.Equal(t, 0, s.GetIndex(7))
	ass.Equal(t, []ele.Moment{4, 5, 6}, s.AsArray())
}

func TestIntervalsWithRunes(t *tes.T) {
	var s = tem.Interval[tem.Rune]('a', tem.LEFT, 'f')
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.Equal(t, tem.Rune('a'), s.GetValue(1))
	ass.Equal(t, 1, s.GetIndex('a'))
	ass.Equal(t, 4, s.GetIndex('d'))
	ass.Equal(t, 0, s.GetIndex('f'))
	ass.Equal(t, []tem.Rune{'a', 'b', 'c', 'd', 'e'}, s.AsArray())
}
