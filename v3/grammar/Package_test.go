/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar_test

import (
	fmt "fmt"
	gra "github.com/bali-nebula/go-component-framework/v3/grammar"
	uti "github.com/craterdog/go-missing-utilities/v2"
	ass "github.com/stretchr/testify/assert"
	osx "os"
	sts "strings"
	tes "testing"
)

const testDirectory = "./test/"

func TestParsingRoundtrips(t *tes.T) {
	var parser = gra.ParserClass().Parser()
	var formatter = gra.FormatterClass().Formatter()

	var files, err = osx.ReadDir(testDirectory)
	if err != nil {
		panic("Could not find the ./test directory.")
	}

	for _, file := range files {
		var filename = testDirectory + file.Name()
		if sts.HasSuffix(filename, ".bali") {
			fmt.Println(filename)
			var source = uti.ReadFile(filename)
			var document = parser.ParseSource(source)
			var formatted = formatter.FormatDocument(document)
			ass.Equal(t, source, formatted)
		}
	}
}
