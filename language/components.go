/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package language

import (
	"fmt"
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/collections"
	"github.com/craterdog-bali/go-bali-document-notation/elements"
	"strings"
)

// PARSER INTERFACE

// This function parses the specified Bali Document Notation™ (BDN) source
// string and returns the corresponding abstract syntax tree as defined in
// the language specification:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification
//
// All parser rules in the specification are shown in lowerCamelCase.
func ParseSource(source string) *Component {
	var ok bool
	var token *Token
	var component *Component
	var parser = Parser([]byte(source))
	var message string
	component, token, ok = parser.parseComponent()
	if ok {
		_, token, ok = parser.parseEOF()
		if !ok {
			message = fmt.Sprintf("Expected an EOF following the component but received:\n%v\n\n", token)
			message += generateGrammar(
				"$component",
				"$entity",
				"$context",
				"$parameters",
				"$parameter",
				"$name")
			panic(message)
		}
	}
	return component
}

// This function parses the specified Bali Document Notation™ (BDN) source
// document and returns the corresponding abstract syntax tree. A POSIX
// compliant source file must end with a EOL character.
func ParseDocument(document []byte) *Component {
	var ok bool
	var token *Token
	var component *Component
	var parser = Parser(document)
	component, token, ok = parser.parseComponent()
	if ok {
		_, token, ok = parser.parseEOL() // Required by POSIX.
		if !ok {
			var message = fmt.Sprintf("Expected an EOL following the component but received:\n%v\n\n", token)
			message += generateGrammar(
				"$document",
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
			var message = fmt.Sprintf("Expected an EOF following the EOL but received:\n%v\n\n", token)
			message += generateGrammar(
				"$document",
				"$component",
				"$entity",
				"$context",
				"$parameters",
				"$parameter",
				"$name")
			panic(message)
		}
	}
	return component
}

// COMPONENT NODES

// This type defines the node structure associated with a component.
type Component struct {
	Entity  any // A entity is an element, string, collection or procedure.
	Context []*Parameter
}

// This type defines the node structure associated with a name-value pair. It is
// used by a component to maintain its parameters.
type Parameter struct {
	Name  elements.Symbol
	Value *Component
}

// PARSER IMPLEMENTATION

// This constructor creates a new parser using the specified byte array.
func Parser(source []byte) *parser {
	var tokens = make(chan Token, 100)
	Scanner(source, tokens) // Starts scanning in a go routine.
	var p = &parser{
		source:   source,
		previous: collections.Stack[*Token](),
		next:     collections.Stack[*Token](),
		tokens:   tokens}
	return p
}

// This type defines the structure and methods for the parser agent.
type parser struct {
	source   []byte
	previous abstractions.StackLike[*Token] // The stack of the previously retrieved tokens.
	next     abstractions.StackLike[*Token] // The stack of the retrieved tokens that have been put back.
	tokens   chan Token                     // The queue of unread tokens coming from the scanner.
}

// PARSING METHODS
//
// Each parsing method returns the following three results:
//  * The parsed component corresponding to the source string, or nil if the
//    parsing failed.
//  * The parsed token that caused the failure if the parsing did fail.
//  * Whether or not the parsing succeeded.

// This method attempts to parse a comment. It returns a string containing the
// comment and whether or not the comment was successfully parsed.
func (v *parser) parseComment() (string, *Token, bool) {
	var comment string
	var token = v.nextToken()
	if token.Type != TokenComment {
		v.backupOne()
		return comment, token, false
	}
	comment = token.Value
	return comment, token, true
}

// This method attempts to parse a component. It returns the component and
// whether or not the component was successfully parsed.
func (v *parser) parseComponent() (*Component, *Token, bool) {
	var component *Component
	var context []*Parameter
	var entity, token, ok = v.parseEntity()
	if ok {
		context, token, _ = v.parseContext() // The context is optional.
		component = &Component{entity, context}
	}
	return component, token, ok
}

// This method attempts to parse the context for a parameterized component. It
// returns an array of parameters and whether or not the context was
// successfully parsed.
func (v *parser) parseContext() ([]*Parameter, *Token, bool) {
	var ok bool
	var token *Token
	var parameters []*Parameter
	_, token, ok = v.parseDelimiter("(")
	if !ok {
		return nil, token, false
	}
	parameters, token, ok = v.parseParameters()
	if !ok {
		var message = fmt.Sprintf("Expected at least one parameter following the '(' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$context",
			"$parameters",
			"$parameter",
			"$name")
		panic(message)
	}
	_, token, ok = v.parseDelimiter(")")
	if !ok {
		var message = fmt.Sprintf("Expected a ')' character following the parameters but received:\n%v\n\n", token)
		message += generateGrammar(
			"$context",
			"$parameters",
			"$parameter",
			"$name")
		panic(message)
	}
	return parameters, token, true
}

// This method attempts to parse the specified delimiter. It returns
// the token and whether or not the delimiter was found.
func (v *parser) parseDelimiter(delimiter string) (string, *Token, bool) {
	var value string
	var token = v.nextToken()
	if token.Type == TokenEOF || token.Value != delimiter {
		v.backupOne()
		return value, token, false
	}
	value = token.Value
	return value, token, true
}

// This method attempts to parse documentation. It returns the documentation and
// whether or not the documentation was successfully parsed.
func (v *parser) parseDocumentation() (string, *Token, bool) {
	var ok bool
	var token *Token
	var documentation string
	documentation, token, ok = v.parseNote()
	if !ok {
		documentation, token, ok = v.parseComment()
	}
	return documentation, token, ok
}

// This method attempts to parse a component entity. It returns the component
// entity and whether or not the component entity was successfully parsed.
func (v *parser) parseEntity() (any, *Token, bool) {
	var ok bool
	var token *Token
	var entity any
	entity, token, ok = v.parseElement()
	if !ok {
		entity, token, ok = v.parseString()
	}
	if !ok {
		entity, token, ok = v.parseCollection()
	}
	if !ok {
		entity, token, ok = v.parseProcedure()
	}
	return entity, token, ok
}

// This method attempts to parse the end-of-file (EOF) marker. It returns
// the token and whether or not an EOL token was found.
func (v *parser) parseEOF() (*Token, *Token, bool) {
	var token = v.nextToken()
	if token.Type != TokenEOF {
		v.backupOne()
		return token, token, false
	}
	return token, token, true
}

// This method attempts to parse the end-of-line (EOL) marker. It returns
// the token and whether or not an EOF token was found.
func (v *parser) parseEOL() (*Token, *Token, bool) {
	var token = v.nextToken()
	if token.Type != TokenEOL {
		v.backupOne()
		return token, token, false
	}
	return token, token, true
}

// This method attempts to parse an identifier. It returns the identifier
// string and whether or not the identifier was successfully parsed.
func (v *parser) parseIdentifier() (string, *Token, bool) {
	var identifier string = "<UNKNOWN>"
	var token = v.nextToken()
	if token.Type != TokenIdentifier {
		v.backupOne()
		return identifier, token, false
	}
	identifier = token.Value
	return identifier, token, true
}

// This method attempts to parse the specified keyword. It returns
// the token and whether or not the keyword was found.
func (v *parser) parseKeyword(keyword string) (string, *Token, bool) {
	var value string
	var token = v.nextToken()
	if token.Type == TokenKeyword || token.Value != keyword {
		v.backupOne()
		return value, token, false
	}
	value = token.Value
	return value, token, true
}

// This method attempts to parse a note. It returns a string containing the
// note and whether or not the note was successfully parsed.
func (v *parser) parseNote() (string, *Token, bool) {
	var note string
	var token = v.nextToken()
	if token.Type != TokenNote {
		v.backupOne()
		return note, token, false
	}
	note = token.Value
	return note, token, true
}

// This method attempts to parse a parameter. It returns the parameter and
// whether or not the parameter was successfully parsed.
func (v *parser) parseParameter() (*Parameter, *Token, bool) {
	var ok bool
	var token *Token
	var name elements.Symbol
	var value *Component
	name, token, ok = v.parseSymbol()
	if !ok {
		return nil, token, false
	}
	_, token, ok = v.parseDelimiter(":")
	if !ok {
		var message = fmt.Sprintf("Expected a ':' character after the name but received:\n%v\n\n", token)
		message += generateGrammar(
			"$parameter",
			"$name")
		panic(message)
	}
	value, token, ok = v.parseComponent()
	if !ok {
		var message = fmt.Sprintf("Expected a component following the ':' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$parameter",
			"$name")
		panic(message)
	}
	var parameter = &Parameter{name, value}
	return parameter, token, true
}

// This method attempts to parse the parameters within a context. It returns an
// array of the parameters and whether or not the parameters were successfully
// parsed.
func (v *parser) parseParameters() ([]*Parameter, *Token, bool) {
	var parameter *Parameter
	var parameters []*Parameter
	var _, token, ok = v.parseEOL()
	if !ok {
		// The parameters are on a single line.
		parameter, token, ok = v.parseParameter()
		// There must be at least one parameter.
		if !ok {
			var message = fmt.Sprintf("Expected at least one parameter in the component context but received:\n%v\n\n", token)
			message += generateGrammar(
				"$context",
				"$parameters",
				"$parameter",
				"$name")
			panic(message)
		}
		for {
			parameters = append(parameters, parameter)
			// Every subsequent parameter must be preceded by a ','.
			_, token, ok = v.parseDelimiter(",")
			if !ok {
				// No more parameters.
				break
			}
			parameter, token, ok = v.parseParameter()
			if !ok {
				var message = fmt.Sprintf("Expected a parameter after the ',' character but received:\n%v\n\n", token)
				message += generateGrammar(
					"$context",
					"$parameters",
					"$parameter",
					"$name")
				panic(message)
			}
		}
		return parameters, token, true
	}
	// The parameters are on separate lines.
	for {
		parameter, token, ok = v.parseParameter()
		if !ok {
			// No more parameters.
			break
		}
		parameters = append(parameters, parameter)
		// Every parameter must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = fmt.Sprintf("Expected an EOL character following the parameter but received:\n%v\n\n", token)
			message += generateGrammar(
				"$context",
				"$parameters",
				"$parameter",
				"$name")
			panic(message)
		}
	}
	// There must be at least one parameter.
	if len(parameters) == 0 {
		var message = fmt.Sprintf("Expected at least one parameter in the component context but received:\n%v\n\n", token)
		message += generateGrammar(
			"$context",
			"$parameters",
			"$parameter",
			"$name")
		panic(message)
	}
	return parameters, token, true
}

// PRIVATE METHODS

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
	var lines = strings.Split(string(v.source), "\n")

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
