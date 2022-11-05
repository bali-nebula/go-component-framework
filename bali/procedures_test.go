/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali_test

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	bal "github.com/craterdog-bali/go-bali-document-notation/bali"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRoundtripWithStatements(t *tes.T) {
	var statementStrings = []string{
		`{ }`,
		`{
    return none
}`,
	}

	for index, s := range statementStrings {
		var component = bal.ParseSource(s).(abs.ComponentLike)
		var s = bal.FormatComponent(component)
		ass.Equal(t, statementStrings[index], s)
	}
}
