/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
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
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
	sts "strings"
)

// PARSER INTERFACE

// This function parses the specified Bali Document Notation™ (BDN) source
// bytes retrieved from a POSIX compliant file and returns the corresponding
// abstract syntax tree as defined in the language specification:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification
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
			var message = parser.formatError(token)
			message += generateGrammar("EOL",
				"$source",
				"$component")
			panic(message)
		}
		_, token, ok = parser.parseEOF()
		if !ok {
			var message = parser.formatError(token)
			message += generateGrammar("EOF",
				"$source",
				"$component")
			panic(message)
		}
	} else {
		var message = parser.formatError(token)
		message += generateGrammar("component",
			"$source",
			"$component",
			"$entity",
			"$context")
		panic(message)
	}
	return component
}

// This function parses a source string rather than the bytes from a BDN
// document file. It is useful when parsing strings within source code.
func ParseComponent(source string) abs.ComponentLike {
	var document = []byte(source + EOL) // Append the POSIX compliant EOL character.
	return ParseDocument(document)
}

// This function parses an entity from a source string.
func ParseEntity(source string) abs.Entity {
	var ok bool
	var token *Token
	var entity abs.Entity
	var parser = Parser([]byte(source + EOL))
	entity, token, ok = parser.parseEntity()
	if !ok {
		var message = parser.formatError(token)
		message += generateGrammar("entity",
			"element",
			"string",
			"range",
			"collection",
			"procedure",
		)
		panic(message)
	}
	return entity
}

// This function parses an entity from a source string.
func ParseContext(source string) abs.ContextLike {
	var ok bool
	var token *Token
	var context abs.ContextLike
	var parser = Parser([]byte(source))
	context, token, ok = parser.parseContext()
	if !ok {
		var message = parser.formatError(token)
		message += generateGrammar("context",
			"$component",
			"$entity",
			"$context",
			"$parameters")
		panic(message)
	}
	return context
}

// PARSER IMPLEMENTATION

// This constructor creates a new parser using the specified byte array.
func Parser(source []byte) *parser {
	var tokens = make(chan Token, 256)
	ScanTokens(source, tokens) // Starts scanning in a separate go routine.
	var p = &parser{
		source: source,
		next:   col.StackWithCapacity[*Token](4),
		tokens: tokens,
	}
	return p
}

// This type defines the structure and methods for the parser agent.
type parser struct {
	source []byte
	next   col.StackLike[*Token] // The stack of the retrieved tokens that have been put back.
	tokens chan Token            // The queue of unread tokens coming from the scanner.
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
		next = &token
		if next.Type == TokenERROR {
			var message = v.formatError(next)
			panic(message)
		}
	} else {
		next = v.next.RemoveTop()
	}
	return next
}

// This method puts back the current token onto the token stream so that it can
// be retrieved by another parsing method.
func (v *parser) backupOne(token *Token) {
	v.next.AddValue(token)
}

// This method returns an error message containing the context for a parsing
// error.
func (v *parser) formatError(token *Token) string {
	var message = fmt.Sprintf("An unexpected token was received by the parser: %v\n", token)
	var line = token.Line
	var lines = sts.Split(string(v.source), EOL)

	message += "\033[36m"
	if line > 1 {
		message += fmt.Sprintf("%04d: ", line-1) + string(lines[line-2]) + EOL
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + EOL

	message += " \033[32m>>>─"
	var count = 0
	for count < token.Position {
		message += "─"
		count++
	}
	message += "⌃\033[36m\n"

	if line < len(lines) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + EOL
	}
	message += "\033[0m\n"

	return message
}
