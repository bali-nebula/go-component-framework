/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	//fmt "fmt"
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestASCIICharacters(t *tes.T) {
	var v = ele.CharacterFromString(`"a"`)
	ass.Equal(t, `"a"`, v.AsString())

	v = ele.CharacterFromString(`"'"`)
	ass.Equal(t, `"'"`, v.AsString())
}

func TestUnicodeCharacters(t *tes.T) {
	var v = ele.CharacterFromString(`"😊"`)
	ass.Equal(t, `"😊"`, v.AsString())

	v = ele.CharacterFromString(`"界"`)
	ass.Equal(t, `"界"`, v.AsString())
}

func TestEscapedCharacters(t *tes.T) {
	var v = ele.CharacterFromString(`"\""`)
	ass.Equal(t, `"\""`, v.AsString())

	v = ele.CharacterFromString(`"\\"`)
	ass.Equal(t, `"\\"`, v.AsString())

	v = ele.CharacterFromString(`"\n"`)
	ass.Equal(t, `"\n"`, v.AsString())

	v = ele.CharacterFromString(`"\t"`)
	ass.Equal(t, `"\t"`, v.AsString())
}
