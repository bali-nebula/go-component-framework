/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package language_test

import (
	str "github.com/craterdog-bali/go-bali-document-notation/strings"
	lan "github.com/craterdog-bali/go-bali-document-notation/language"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestRoundtripWithBinaries(t *tes.T) {
	var binaryStrings = []string{
		`''`,
		`'NyiXwRMG7OzG8P8y21A6lTFKWU6sFel5vCaYw1kslyq0gLkJTMHu5iNaFRwZ'`,
		`'
    EEzlwvUOVeuvUfs/IbPxLTZQqSFuy50r/XLCHdtnkB/fdwTzcwd3ManlLh+G
    WnQ4dN6fednRMep1pKYUVsEXaTz63jIwM62ID0hCukrP3CRiHZwYIgKnQ00w
    Rrk30KChDM1NILvA675+T6yR7P3xg/d8v/Rd35WPe4fZ5ACLzf6J7+ZTGDws
    SZX92PsBv+4fiMG4dLUMxjnNVvN9LElT9nVy7BAB9i5PiBOr2hFuV8+FTTBy
    oE89y7SjSM7BLw7m6DEfYJf5sapuz05pythmmqcqbbU9eDWkEfLbaS5MxYdN
    JocIaQNamhzDhmriK0WEWNdUdzRC992R2ZkTYAQ8vg
'`,
	}

	var binaryValues = []str.Binary{
		"",
		"NyiXwRMG7OzG8P8y21A6lTFKWU6sFel5vCaYw1kslyq0gLkJTMHu5iNaFRwZ",
		"EEzlwvUOVeuvUfs/IbPxLTZQqSFuy50r/XLCHdtnkB/fdwTzcwd3ManlLh+G" +
		"WnQ4dN6fednRMep1pKYUVsEXaTz63jIwM62ID0hCukrP3CRiHZwYIgKnQ00w" +
		"Rrk30KChDM1NILvA675+T6yR7P3xg/d8v/Rd35WPe4fZ5ACLzf6J7+ZTGDws" +
		"SZX92PsBv+4fiMG4dLUMxjnNVvN9LElT9nVy7BAB9i5PiBOr2hFuV8+FTTBy" +
		"oE89y7SjSM7BLw7m6DEfYJf5sapuz05pythmmqcqbbU9eDWkEfLbaS5MxYdN" +
		"JocIaQNamhzDhmriK0WEWNdUdzRC992R2ZkTYAQ8vg",
	}

	for index, s := range binaryStrings {
		var binary = lan.ParseSource(s).GetEntity().(str.Binary)
		ass.Equal(t, binaryValues[index], binary)
	}

	for index, binary := range binaryValues {
		var s = lan.FormatValue(binary)
		ass.Equal(t, binaryStrings[index], s)
	}
}

func TestRoundtripWithMonikers(t *tes.T) {
	var monikerStrings = []string{
		`/bali`,
		`/bali/abstractions`,
		`/bali/abstractions/Sequential/v1.2.3`,
	}

	var monikerValues = []str.Moniker{
		"/bali",
		"/bali/abstractions",
		"/bali/abstractions/Sequential/v1.2.3",
	}

	for index, s := range monikerStrings {
		var moniker = lan.ParseSource(s).GetEntity().(str.Moniker)
		ass.Equal(t, monikerValues[index], moniker)
	}

	for index, moniker := range monikerValues {
		var s = lan.FormatValue(moniker)
		ass.Equal(t, monikerStrings[index], s)
	}
}

func TestRoundtripWithNarratives(t *tes.T) {
	var narrativeStrings = []string{
		`">
    
<"`,
		`">
    This is a multiline
    narrative with a '"'
    character in it.
<"`,
	}

	var narrativeValues = []str.Narrative{
	`    `,
	"    This is a multiline\n" +
	"    narrative with a '\"'\n" + 
	"    character in it.",
	}

	for index, s := range narrativeStrings {
		var narrative = lan.ParseSource(s).GetEntity().(str.Narrative)
		ass.Equal(t, narrativeValues[index], narrative)
	}

	for index, narrative := range narrativeValues {
		var s = lan.FormatValue(narrative)
		ass.Equal(t, narrativeStrings[index], s)
	}
}

func TestRoundtripWithQuotes(t *tes.T) {
	var quoteStrings = []string{
		`""`,
		`"To be, or not to be, that is the question?"`,
		`"This quote contains '\"'s."`,
	}

	var quoteValues = []str.Quote{
		"",
		"To be, or not to be, that is the question?",
		`This quote contains '"'s.`,
	}

	for index, s := range quoteStrings {
		var quote = lan.ParseSource(s).GetEntity().(str.Quote)
		ass.Equal(t, quoteValues[index], quote)
	}

	for index, quote := range quoteValues {
		var s = lan.FormatValue(quote)
		ass.Equal(t, quoteStrings[index], s)
	}
}

func TestRoundtripWithVersions(t *tes.T) {
	var versionStrings = []string{
		`v1`,
		`v1.2`,
		`v1.2.3`,
	}

	var versionValues = []str.Version{
		"1",
		"1.2",
		"1.2.3",
	}

	for index, s := range versionStrings {
		var version = lan.ParseSource(s).GetEntity().(str.Version)
		ass.Equal(t, versionValues[index], version)
	}

	for index, version := range versionValues {
		var s = lan.FormatValue(version)
		ass.Equal(t, versionStrings[index], s)
	}
}
