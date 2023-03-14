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
	bal "github.com/bali-nebula/go-component-framework/v1/bali"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRoundtripWithStatements(t *tes.T) {
	var statementStrings = []string{
		`{ }`,
		`{
    return none
}`,
		`{
    let $foo := "bar"
    throw [
        $type: $bad
        $kind: $worse
    ]
}`,
		`{
    if NOT list.isEmpty() do {
        let $first := list[1]
    }
}`,
	}

	for index, s := range statementStrings {
		var component = bal.ParseComponent(s)
		var s = bal.FormatComponent(component)
		ass.Equal(t, statementStrings[index], s)
	}
}
