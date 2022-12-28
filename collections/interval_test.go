/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	abs "github.com/craterdog-bali/go-component-framework/abstractions"
	col "github.com/craterdog-bali/go-component-framework/collections"
	ele "github.com/craterdog-bali/go-component-framework/elements"
	co2 "github.com/craterdog/go-collection-framework"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestIntervalsWithDurations(t *tes.T) {
	var s = col.Interval[ele.Duration](3, abs.INCLUSIVE, 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.Equal(t, 5, int(s.GetValue(3)))
	ass.Equal(t, 0, s.GetIndex(2))
	ass.Equal(t, 1, s.GetIndex(3))
	ass.Equal(t, 3, s.GetIndex(5))
	ass.Equal(t, 5, s.GetIndex(7))
	ass.Equal(t, 0, s.GetIndex(8))
	ass.Equal(t, []ele.Duration{3, 4, 5, 6, 7}, s.AsArray())
	var iterator = co2.Iterator[ele.Duration](s)
	ass.Equal(t, 3, int(iterator.GetNext()))
	iterator.ToEnd()
	ass.Equal(t, 7, int(iterator.GetPrevious()))
}

func TestIntervalsWithMoments(t *tes.T) {
	var s = col.Interval[ele.Moment](3, abs.RIGHT, 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 4, s.GetSize())
	ass.Equal(t, 4, int(s.GetValue(1)))
	ass.Equal(t, 0, s.GetIndex(3))
	ass.Equal(t, 1, s.GetIndex(4))
	ass.Equal(t, 2, s.GetIndex(5))
	ass.Equal(t, 4, s.GetIndex(7))
	ass.Equal(t, 0, s.GetIndex(8))
	ass.Equal(t, []ele.Moment{4, 5, 6, 7}, s.AsArray())

	s = col.Interval[ele.Moment](3, abs.EXCLUSIVE, 7)
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 3, s.GetSize())
	ass.Equal(t, 4, int(s.GetValue(1)))
	ass.Equal(t, 0, s.GetIndex(3))
	ass.Equal(t, 1, s.GetIndex(4))
	ass.Equal(t, 3, s.GetIndex(6))
	ass.Equal(t, 0, s.GetIndex(7))
	ass.Equal(t, []ele.Moment{4, 5, 6}, s.AsArray())
}

func TestIntervalsWithRunes(t *tes.T) {
	var s = col.Interval[col.Rune]('a', abs.LEFT, 'f')
	ass.False(t, s.IsEmpty())
	ass.Equal(t, 5, s.GetSize())
	ass.Equal(t, 'a', rune(s.GetValue(1)))
	ass.Equal(t, 1, s.GetIndex('a'))
	ass.Equal(t, 4, s.GetIndex('d'))
	ass.Equal(t, 0, s.GetIndex('f'))
	ass.Equal(t, []col.Rune{'a', 'b', 'c', 'd', 'e'}, s.AsArray())
}
