/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bdn_test

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	bdn "github.com/craterdog-bali/go-bali-document-notation/bdn"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRoundtripWithComponents(t *tes.T) {
	var componentStrings = []string{
		`none`,
		`false  ! This is not true.`,
		`[ ]($type: /bali/collections/Set/v1)  ! This is a note`,
	}

	for index, s := range componentStrings {
		var component = bdn.ParseSource(s).(abs.ComponentLike)
		var s = bdn.FormatComponent(component)
		ass.Equal(t, componentStrings[index], s)
	}
}
