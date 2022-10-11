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
func ParseSource(source string) (*Component, bool) {
	var ok bool
	var component *Component
	var v = Parser([]byte(source))
	component, ok = v.parseComponent()
	if ok {
		_, ok = v.parseEOF()
		if !ok {
			panic("Source contained more than just a component.")
		}
	}
	return component, ok
}

// This function parses the specified Bali Document Notation™ (BDN) source
// document and returns the corresponding abstract syntax tree. A POSIX
// compliant source file must end with a EOL character.
func ParseDocument(document []byte) (*Component, bool) {
	var ok bool
	var component *Component
	var v = Parser(document)
	component, ok = v.parseComponent()
	if ok {
		_, ok = v.parseEOL() // Required by POSIX.
		if !ok {
			panic("Document is missing a final EOL.")
		}
		_, ok = v.parseEOF()
		if !ok {
			panic("Source contained more than just a component.")
		}
	}
	return component, ok
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
	var tokens = make(chan token, 100)
	Scanner(source, tokens) // Starts scanning in a go routine.
	var p = &parser{
		source:   source,
		previous: collections.Stack[*token](),
		next:     collections.Stack[*token](),
		tokens:   tokens}
	return p
}

// This type defines the structure and methods for the parser agent.
type parser struct {
	source   []byte
	previous abstractions.StackLike[*token] // The stack of the previously retrieved tokens.
	next     abstractions.StackLike[*token] // The stack of the retrieved tokens that have been put back.
	tokens   chan token                     // The queue of unread tokens coming from the scanner.
}

// This method attempts to read the next token from the token stream and return
// it.
func (v *parser) nextToken() *token {
	var next *token
	if v.next.IsEmpty() {
		var token, ok = <-v.tokens
		if !ok {
			panic("The token channel terminated without an EOF or error token.")
		}
		if token.typ == tokenError {
			panic(v.formatError(
				"An unexpected character was encountered while scanning the input: '"+token.val+"'",
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
func (v *parser) formatError(message string, token *token) string {
	message += "\n\n"
	var line = token.lin
	var lines = strings.Split(string(v.source), "\n")

	if line > 1 {
		message += fmt.Sprintf("%04d: ", line-1) + string(lines[line-2]) + "\n"
	}
	message += fmt.Sprintf("%04d: ", line) + string(lines[line-1]) + "\n"

	message += " >>>─"
	var count = 0
	for count < token.pos {
		message += "─"
		count++
	}
	message += "⌃\n"

	if line < len(lines) {
		message += fmt.Sprintf("%04d: ", line+1) + string(lines[line]) + "\n"
	}

	return message
}

// This method attempts to parse a comment. It returns a string containing the
// comment and whether or not the comment was successfully parsed.
func (v *parser) parseComment() (string, bool) {
	var comment string
	var token = v.nextToken()
	if token.typ != tokenComment {
		v.backupOne()
		return comment, false
	}
	comment = token.val
	return comment, true
}

// This method attempts to parse a component. It returns the component and
// whether or not the component was successfully parsed.
func (v *parser) parseComponent() (*Component, bool) {
	var component *Component
	var context []*Parameter
	var entity, ok = v.parseEntity()
	if ok {
		context, _ = v.parseContext() // The context is optional.
		component = &Component{entity, context}
	}
	return component, ok
}

// This method attempts to parse the context for a parameterized component. It
// returns an array of parameters and whether or not the context was
// successfully parsed.
func (v *parser) parseContext() ([]*Parameter, bool) {
	var ok bool
	var parameters []*Parameter
	_, ok = v.parseDelimiter("(")
	if !ok {
		return nil, false
	}
	parameters, ok = v.parseParameters()
	if !ok {
		panic("Expected at least one parameter following the '(' character.")
	}
	_, ok = v.parseDelimiter(")")
	if !ok {
		panic("Expected a ')' character following the parameters.")
	}
	return parameters, true
}

// This method attempts to parse the specified delimiter. It returns
// the token and whether or not the delimiter was found.
func (v *parser) parseDelimiter(delimiter string) (*token, bool) {
	var token = v.nextToken()
	if token.typ == tokenEOF || token.val != delimiter {
		v.backupOne()
		return token, false
	}
	return token, true
}

// This method attempts to parse documentation. It returns the documentation and
// whether or not the documentation was successfully parsed.
func (v *parser) parseDocumentation() (string, bool) {
	var ok bool
	var documentation string
	documentation, ok = v.parseNote()
	if !ok {
		documentation, ok = v.parseComment()
	}
	return documentation, ok
}

// This method attempts to parse a component entity. It returns the component
// entity and whether or not the component entity was successfully parsed.
func (v *parser) parseEntity() (any, bool) {
	var ok bool
	var entity any
	entity, ok = v.parseElement()
	if !ok {
		entity, ok = v.parseString()
	}
	if !ok {
		entity, ok = v.parseCollection()
	}
	if !ok {
		entity, ok = v.parseProcedure()
	}
	return entity, ok
}

// This method attempts to parse the end-of-file (EOF) marker. It returns
// the token and whether or not an EOL token was found.
func (v *parser) parseEOF() (*token, bool) {
	var token = v.nextToken()
	if token.typ != tokenEOF {
		v.backupOne()
		return token, false
	}
	return token, true
}

// This method attempts to parse the end-of-line (EOL) marker. It returns
// the token and whether or not an EOF token was found.
func (v *parser) parseEOL() (*token, bool) {
	var token = v.nextToken()
	if token.typ != tokenEOL {
		v.backupOne()
		return token, false
	}
	return token, true
}

// This method attempts to parse an identifier. It returns the identifier
// string and whether or not the identifier was successfully parsed.
func (v *parser) parseIdentifier() (string, bool) {
	var identifier string = "<UNKNOWN>"
	var token = v.nextToken()
	if token.typ != tokenIdentifier {
		v.backupOne()
		return identifier, false
	}
	identifier = token.val
	return identifier, true
}

// This method attempts to parse the specified keyword. It returns
// the token and whether or not the keyword was found.
func (v *parser) parseKeyword(keyword string) (*token, bool) {
	var token = v.nextToken()
	if token.typ == tokenKeyword || token.val != keyword {
		v.backupOne()
		return token, false
	}
	return token, true
}

// This method attempts to parse a note. It returns a string containing the
// note and whether or not the note was successfully parsed.
func (v *parser) parseNote() (string, bool) {
	var note string
	var token = v.nextToken()
	if token.typ != tokenNote {
		v.backupOne()
		return note, false
	}
	note = token.val
	return note, true
}

// This method attempts to parse a parameter. It returns the parameter and
// whether or not the parameter was successfully parsed.
func (v *parser) parseParameter() (*Parameter, bool) {
	var ok bool
	var name elements.Symbol
	var value *Component
	name, ok = v.parseSymbol()
	if !ok {
		return nil, false
	}
	_, ok = v.parseDelimiter(":")
	if !ok {
		panic("Expected a ':' character.")
	}
	value, ok = v.parseComponent()
	if !ok {
		panic("Expected a component following the ':' character.")
	}
	var parameter = &Parameter{name, value}
	return parameter, true
}

// This method attempts to parse the parameters within a context. It returns an
// array of the parameters and whether or not the parameters were successfully
// parsed.
func (v *parser) parseParameters() ([]*Parameter, bool) {
	var parameter *Parameter
	var parameters []*Parameter
	var _, ok = v.parseEOL()
	if !ok {
		// The parameters are on a single line.
		parameter, ok = v.parseParameter()
		// There must be at least one parameter.
		if !ok {
			panic("Expected at least one parameter in the component context.")
		}
		for {
			parameters = append(parameters, parameter)
			// Every subsequent parameter must be preceded by a ','.
			_, ok = v.parseDelimiter(",")
			if !ok {
				// No more parameters.
				break
			}
			parameter, ok = v.parseParameter()
			if !ok {
				panic("Expected a parameter after the ',' character.")
			}
		}
		return parameters, true
	}
	// The parameters are on separate lines.
	for {
		parameter, ok = v.parseParameter()
		if !ok {
			// No more parameters.
			break
		}
		parameters = append(parameters, parameter)
		// Every parameter must be followed by an EOL.
		_, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the parameter.")
		}
	}
	// There must be at least one parameter.
	if len(parameters) == 0 {
		panic("Expected at least one parameter in the component context.")
	}
	return parameters, true
}
