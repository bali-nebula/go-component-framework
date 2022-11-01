/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bali

import (
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	col "github.com/craterdog-bali/go-bali-document-notation/collections"
	sts "strings"
)

// PARSER INTERFACE

// This function parses the specified Bali Document Notation™ (BDN) source
// bytes retrieved from a POSIX compliant file and returns the corresponding
// abstract syntax tree as defined in the language specification:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification
//
// All parser rules in the specification are shown in lowerCamelCase and scanner
// tokens shown in UPPERCASE.
//
// A POSIX compliant file must end with a EOL character before the EOF marker.
func ParseDocument(document []byte) abs.ComponentLike {
	var ok bool
	var token *Token
	var component abs.ComponentLike
	var parser = Parser(document)
	component, token, ok = parser.parseComponent()
	if ok {
		_, token, ok = parser.parseEOL() // Required by POSIX.
		if !ok {
			var message = parser.formatError("An unexpected token was received by the parser:", token)
			message += generateGrammar("EOL",
				"$source",
				"$component",
				"$entity",
				"$context",
				"$parameters",
				"$parameter",
				"$name")
			panic(message)
		}
		_, token, ok = parser.parseEOF()
		if !ok {
			var message = parser.formatError("An unexpected token was received by the parser:", token)
			message += generateGrammar("EOF",
				"$source",
				"$component",
				"$entity",
				"$context",
				"$parameters",
				"$parameter",
				"$name")
			panic(message)
		}
	} else {
		var message = parser.formatError("An unexpected token was received by the parser:", token)
		message += generateGrammar("$component",
			"$source",
			"$component",
			"$entity",
			"$context",
			"$parameters",
			"$parameter",
			"$name")
		panic(message)
	}
	return component
}

// This function parses a source string rather than the bytes from a BDN
// document file. It is useful when parsing strings within source code.
func ParseSource(source string) abs.ComponentLike {
	var document = []byte(source + "\n") // Append the POSIX compliant EOL character.
	return ParseDocument(document)
}

// PARSER IMPLEMENTATION

// This constructor creates a new parser using the specified byte array.
func Parser(source []byte) *parser {
	var tokens = make(chan Token, 100)
	Scanner(source, tokens) // Starts scanning in a go routine.
	var p = &parser{
		source:   source,
		previous: col.Stack[*Token](),
		next:     col.Stack[*Token](),
		tokens:   tokens}
	return p
}

// This type defines the structure and methods for the parser agent.
type parser struct {
	source   []byte
	previous abs.StackLike[*Token] // The stack of the previously retrieved tokens.
	next     abs.StackLike[*Token] // The stack of the retrieved tokens that have been put back.
	tokens   chan Token            // The queue of unread tokens coming from the scanner.
}

// This method attempts to read the next token from the token stream and return
// it.
func (v *parser) nextToken() *Token {
	var next *Token
	if v.next.IsEmpty() {
		var token, ok = <-v.tokens
		if !ok {
			panic("The token channel terminated without an EOF or error token.")
		}
		if token.Type == TokenError {
			panic(v.formatError(
				"An unexpected character was encountered while scanning the input: '"+token.Value+"'",
				&token))
		}
		next = &token
	} else {
		next = v.next.RemoveTop()
	}
	v.previous.AddItem(next)
	return next
}

// This method puts back the current token onto the token stream so that it can
// be retrieved by another parsing method.
func (v *parser) backupOne() {
	if v.previous.IsEmpty() {
		panic("Attempted to put back a previous token that does not exist.")
	}
	var previous = v.previous.RemoveTop()
	v.next.AddItem(previous)
}

// This method returns an error message containing the context for a parsing
// error.
func (v *parser) formatError(message string, token *Token) string {
	message += "\n\n"
	var line = token.Line
	var lines = sts.Split(string(v.source), "\n")

	if line > 1 {
		message += fmt.Sprintf("%04d: ", line-1) + string(lines[line-2]) + "\n"
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	message += " >>>─"
	var count = 0
	for count < token.Position {
		message += "─"
		count++
	}
	message += "⌃\n"

	if line < len(lines) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}

	return message
}
