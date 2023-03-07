/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali_test

import (
	bal "github.com/bali-nebula/go-component-framework/bali"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRoundtripWithComponents(t *tes.T) {
	var componentStrings = []string{
		`none`,
		`false  ! This is not true.`,
		`[ ]($type: /bali/types/collections/Set/v1)  ! This is a note`,
	}

	for index, s := range componentStrings {
		var component = bal.ParseComponent(s)
		var s = bal.FormatComponent(component)
		ass.Equal(t, componentStrings[index], s)
	}
}
