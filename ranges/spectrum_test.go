/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
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
	ran "github.com/bali-nebula/go-component-framework/ranges"
	str "github.com/bali-nebula/go-component-framework/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestSpectrWithBinaries(t *tes.T) {
	var first = str.BinaryFromArray([]byte{0x0, 0x1, 0x3})
	var last = str.BinaryFromArray([]byte{0xa, 0xb, 0xc})
	var s = ran.Spectrum[str.Binary](first, abs.INCLUSIVE, last)
	ass.Equal(t, first, s.GetFirst())
	ass.Equal(t, abs.INCLUSIVE, s.GetExtent())
	ass.Equal(t, last, s.GetLast())
}

func TestSpectrWithMonikers(t *tes.T) {
	var first = str.Moniker("/bali/abstractions")
	var last = str.Moniker("/bali/strings")
	var s = ran.Spectrum[str.Moniker](first, abs.RIGHT, last)
	ass.Equal(t, first, s.GetFirst())
	ass.Equal(t, abs.RIGHT, s.GetExtent())
	ass.Equal(t, last, s.GetLast())
}

func TestSpectrWithQuotes(t *tes.T) {
	var first = str.Quote("A")
	var last = str.Quote("Fe")
	var s = ran.Spectrum[str.Quote](first, abs.EXCLUSIVE, last)
	ass.Equal(t, first, s.GetFirst())
	ass.Equal(t, abs.EXCLUSIVE, s.GetExtent())
	ass.Equal(t, last, s.GetLast())
}

func TestSpectrWithVersions(t *tes.T) {
	var first = str.Version("v1.2")
	var last = str.Version("v1.6.1")
	var s = ran.Spectrum[str.Version](first, abs.LEFT, last)
	ass.Equal(t, first, s.GetFirst())
	ass.Equal(t, abs.LEFT, s.GetExtent())
	ass.Equal(t, last, s.GetLast())
}
