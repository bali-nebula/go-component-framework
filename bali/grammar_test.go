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
	fmt "fmt"
	bal "github.com/bali-nebula/go-component-framework/bali"
	osx "os"
	tes "testing"
)

const specifications = "../../specifications/bali.bwsn"

func TestGenerateGrammar(t *tes.T) {
	var err = osx.WriteFile(specifications, []byte(bal.FormatGrammar()), 0644)
	if err != nil {
		var message = fmt.Sprintf("Could not create the specifications file: %v.", err)
		panic(message)
	}
}
