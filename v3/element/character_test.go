/*
................................................................................
.    Copyright (c) 2009-2024 Crater Dog Technologies™.  All Rights Reserved.   .
................................................................................
.  DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               .
.                                                                              .
.  This code is free software; you can redistribute it and/or modify it under  .
.  the terms of The MIT License (MIT), as published by the Open Source         .
.  Initiative. (See https://opensource.org/license/MIT)                        .
................................................................................
*/

package element_test

import (
	ele "github.com/bali-nebula/go-component-framework/v3/element"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestASCIICharacters(t *tes.T) {
	var Character = ele.Character()
	var v = Character.MakeFromString(`"a"`)
	ass.Equal(t, `"a"`, v.AsString())
	v = Character.ToUppercase(v)
	ass.Equal(t, `"A"`, v.AsString())
	v = Character.ToLowercase(v)
	ass.Equal(t, `"a"`, v.AsString())

	v = Character.MakeFromString(`"'"`)
	ass.Equal(t, `"'"`, v.AsString())
}

func TestUnicodeCharacters(t *tes.T) {
	var Character = ele.Character()
	var v = Character.MakeFromString(`"😊"`)
	ass.Equal(t, `"😊"`, v.AsString())

	v = Character.MakeFromString(`"界"`)
	ass.Equal(t, `"界"`, v.AsString())
}

func TestEscapedCharacters(t *tes.T) {
	var Character = ele.Character()
	var v = Character.MakeFromString(`"\""`)
	ass.Equal(t, `"\""`, v.AsString())

	v = Character.MakeFromString(`"\\"`)
	ass.Equal(t, `"\\"`, v.AsString())

	v = Character.MakeFromString(`"\n"`)
	ass.Equal(t, `"\n"`, v.AsString())

	v = Character.MakeFromString(`"\t"`)
	ass.Equal(t, `"\t"`, v.AsString())
}
