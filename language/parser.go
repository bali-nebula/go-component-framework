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
	"github.com/craterdog-bali/go-bali-document-notation/strings"
)

// PARSER INTERFACE

// This function parses the specified Bali Document Notation™ (BDN) source
// string and returns the corresponding abstract syntax tree.
func ParseSource(source string) (*Component, bool) {
	var ok bool
	var component *Component
	var v = Parser([]byte(source))
	component, ok = v.parseComponent()
	if ok {
		_, ok = v.parseEOF()
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
		v.parseEOL() // Required by POSIX.
		_, ok = v.parseEOF()
	}
	return component, ok
}

// PARSER IMPLEMENTATION

// This constructor creates a new parser using the specified byte array.
func Parser(source []byte) *parser {
	var tokens = make(chan token)
	Scanner(source, tokens) // Starts scanning in a go routine.
	var p = &parser{
		previous: collections.Stack[*token](),
		next:     collections.Stack[*token](),
		tokens:   tokens}
	return p
}

// This type defines the structure and methods for the parser agent.
type parser struct {
	previous abstractions.StackLike[*token] // The stack of the previously retrieved tokens.
	next     abstractions.StackLike[*token] // The stack of the retrieved tokens that have been put back.
	tokens   chan token                     // The queue of unread tokens coming from the scanner.
}

// This method attempts to read the next token from the token stream and return
// it.
func (v *parser) nextToken() *token {
	var next *token
	if v.next.IsEmpty() {
		var t, ok = <-v.tokens
		if !ok {
			panic("The token channel terminated without an EOF token.")
		}
		if t.typ == tokenError {
			panic(fmt.Sprintf("%s", t.val))
		}
		next = &t
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

// This method attempts to parse a angle element. It returns the angle element
// and whether or not the angle element was successfully parsed.
func (v *parser) parseAngle() (elements.Angle, bool) {
	var ok bool
	var angle elements.Angle
	var token = v.nextToken()
	if token.typ != tokenAngle {
		v.backupOne()
		return angle, false
	}
	angle, ok = elements.AngleFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid angle token was found: %v", token))
	}
	return angle, true
}

// This method attempts to parse an association between a key and value. It
// returns the association and whether or not the association was successfully
// parsed.
func (v *parser) parseAssociation() (abstractions.AssociationLike[any, any], bool) {
	var ok bool
	var key any
	var value any
	key, ok = v.parsePrimitive()
	if !ok {
		// This is not an association.
		return nil, false
	}
	_, ok = v.parseDelimiter(":")
	if !ok {
		// This is not an association.
		v.backupOne() // Put back the primitive token.
		return nil, false
	}
	// This must be an association.
	value, ok = v.parseComponent()
	if !ok {
		panic("Expected a value after the ':' character.")
	}
	var association = collections.Association[any, any](key, value)
	return association, true
}

// This method attempts to parse a binary string. It returns the binary
// string and whether or not the binary string was successfully parsed.
func (v *parser) parseBinary() (strings.Binary, bool) {
	var ok bool
	var binary strings.Binary
	var token = v.nextToken()
	if token.typ != tokenBinary {
		v.backupOne()
		return binary, false
	}
	binary, ok = strings.BinaryFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid binary token was found: %v", token))
	}
	return binary, true
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

// This method attempts to parse a boolean element. It returns the boolean
// element and whether or not the boolean element was successfully parsed.
func (v *parser) parseBoolean() (elements.Boolean, bool) {
	var ok bool
	var boolean elements.Boolean
	var token = v.nextToken()
	if token.typ != tokenBoolean {
		v.backupOne()
		return boolean, false
	}
	boolean, ok = elements.BooleanFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid boolean token was found: %v", token))
	}
	return boolean, true
}

// This method attempts to parse a catalog collection. It returns the
// catalog collection and whether or not the catalog collection was
// successfully parsed.
func (v *parser) parseCatalog() (abstractions.CatalogLike[any, any], bool) {
	var ok bool
	var association abstractions.AssociationLike[any, any]
	var catalog = collections.Catalog[any, any]()
	_, ok = v.parseEOL()
	if !ok {
		// The associations are on a single line.
		_, ok = v.parseDelimiter(":")
		if ok {
			// This is an empty catalog.
			return catalog, true
		}
		association, ok = v.parseAssociation()
		if !ok {
			// A non-empty catalog must have at least one association.
			return nil, false
		}
		for {
			catalog.AddAssociation(association)
			// Every subsequent association must be preceded by a ','.
			_, ok = v.parseDelimiter(",")
			if !ok {
				// No more associations.
				break
			}
			association, ok = v.parseAssociation()
			if !ok {
				panic("Expected an association after the ',' character.")
			}
		}
		return catalog, true
	}
	// The associations are on separate lines.
	association, ok = v.parseAssociation()
	if !ok {
		// A non-empty catalog must have at least one association.
		v.backupOne() // Put back the EOL token.
		return nil, false
	}
	for {
		catalog.AddAssociation(association)
		// Every association must be followed by an EOL.
		_, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the association.")
		}
		association, ok = v.parseAssociation()
		if !ok {
			// No more associations.
			break
		}
	}
	return catalog, true
}

// This method attempts to parse a sequence of items. It returns the
// sequence and whether or not the sequence was successfully parsed.
func (v *parser) parseSequence() (any, bool) {
	var ok bool
	var sequence any
	sequence, ok = v.parseCatalog()
	if !ok {
		sequence, ok = v.parseSlice()
	}
	if !ok {
		sequence, ok = v.parseList() // The list must be last.
	}
	return sequence, ok
}

// This method attempts to parse a component. It returns the component and
// whether or not the component was successfully parsed.
func (v *parser) parseComponent() (*Component, bool) {
	var context []*Parameter
	var note string
	var component *Component
	var entity, ok = v.parseEntity()
	if ok {
		context, _ = v.parseContext() // The context is optional.
		note, _ = v.parseNote()       // The note is optional.
		component = &Component{entity, context, note}
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

// This method attempts to parse a duration element. It returns the duration
// element and whether or not the duration element was successfully parsed.
func (v *parser) parseDuration() (elements.Duration, bool) {
	var ok bool
	var duration elements.Duration
	var token = v.nextToken()
	if token.typ != tokenDuration {
		v.backupOne()
		return duration, false
	}
	duration, ok = elements.DurationFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid duration token was found: %v", token))
	}
	return duration, true
}

// This method attempts to parse an element primitive. It returns the
// element primitive and whether or not the element primitive was
// successfully parsed.
func (v *parser) parseElement() (any, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var element any
	element, ok = v.parseAngle()
	if !ok {
		element, ok = v.parseBoolean()
	}
	if !ok {
		element, ok = v.parseDuration()
	}
	if !ok {
		element, ok = v.parseMoment()
	}
	if !ok {
		element, ok = v.parseNumber()
	}
	if !ok {
		element, ok = v.parsePattern()
	}
	if !ok {
		element, ok = v.parsePercentage()
	}
	if !ok {
		element, ok = v.parseProbability()
	}
	if !ok {
		element, ok = v.parseResource()
	}
	if !ok {
		element, ok = v.parseTag()
	}
	if !ok {
		element = nil
	}
	return element, ok
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

// This method attempts to parse an expression. It returns the expression and
// whether or not the expression was successfully parsed.
func (v *parser) parseExpression() (any, bool) {
	// TODO: Reorder these based on how often each type of expression occurs.
	var ok bool
	var expression any
	expression, ok = v.parseComponent()
	/*
		if !ok {
			expression, ok = v.parseVariable()
		}
		if !ok {
			expression, ok = v.parseFunctionExpression()
		}
		if !ok {
			expression, ok = v.parsePrecedenceExpression()
		}
		if !ok {
			expression, ok = v.parseDereferenceExpression()
		}
		if !ok {
			expression, ok = v.parseMessageExpression()
		}
		if !ok {
			expression, ok = v.parseAttributeExpression()
		}
		if !ok {
			expression, ok = v.parseChainExpression()
		}
		if !ok {
			expression, ok = v.parsePowerExpression()
		}
		if !ok {
			expression, ok = v.parseInversionExpression()
		}
		if !ok {
			expression, ok = v.parseArithmeticExpression()
		}
		if !ok {
			expression, ok = v.parseMagnitudeExpression()
		}
		if !ok {
			expression, ok = v.parseComparisonExpression()
		}
		if !ok {
			expression, ok = v.parseComplementExpression()
		}
		if !ok {
			expression, ok = v.parseLogicalExpression()
		}
		if !ok {
			expression, ok = v.parseDefaultExpression()
		}
	*/
	return expression, ok
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

// This method attempts to parse a list of items. It returns the
// list of items and whether or not the list of items was
// successfully parsed.
func (v *parser) parseList() (abstractions.ListLike[any], bool) {
	var ok bool
	var item any
	var list = collections.List[any]()
	_, ok = v.parseEOL()
	if !ok {
		// The items are on a single line.
		_, ok = v.parseDelimiter("]")
		if ok {
			// This is an empty list.
			v.backupOne() // Put back the ']' token.
			return list, true
		}
		item, ok = v.parseComponent()
		if !ok {
			// A non-empty list must have at least one item.
			panic("Expected an item after the '[' character.")
		}
		for {
			list.AddItem(item)
			// Every subsequent item must be preceded by a ','.
			_, ok = v.parseDelimiter(",")
			if !ok {
				// No more items.
				break
			}
			item, ok = v.parseComponent()
			if !ok {
				panic("Expected an item after the ',' character.")
			}
		}
		return list, true
	}
	// The items are on separate lines.
	item, ok = v.parseComponent()
	if !ok {
		// A non-empty list must have at least one item.
		panic("Expected an item after the '[' character.")
	}
	for {
		list.AddItem(item)
		// Every item must be followed by an EOL.
		_, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the item.")
		}
		item, ok = v.parseComponent()
		if !ok {
			// No more items.
			break
		}
	}
	return list, true
}

// This method attempts to parse a moment element. It returns the moment
// element and whether or not the moment element was successfully parsed.
func (v *parser) parseMoment() (elements.Moment, bool) {
	var ok bool
	var moment elements.Moment
	var token = v.nextToken()
	if token.typ != tokenMoment {
		v.backupOne()
		return moment, false
	}
	moment, ok = elements.MomentFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid moment token was found: %v", token))
	}
	return moment, true
}

// This method attempts to parse a moniker string. It returns the moniker string
// and whether or not the moniker string was successfully parsed.
func (v *parser) parseMoniker() (strings.Moniker, bool) {
	var ok bool
	var moniker strings.Moniker
	var token = v.nextToken()
	if token.typ != tokenMoniker {
		v.backupOne()
		return moniker, false
	}
	moniker, ok = strings.MonikerFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid moniker token was found: %v", token))
	}
	return moniker, true
}

// This method attempts to parse a narrative string. It returns the narrative
// string and whether or not the narrative string was successfully parsed.
func (v *parser) parseNarrative() (strings.Narrative, bool) {
	var ok bool
	var narrative strings.Narrative
	var token = v.nextToken()
	if token.typ != tokenNarrative {
		v.backupOne()
		return narrative, false
	}
	narrative, ok = strings.NarrativeFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid narrative token was found: %v", token))
	}
	return narrative, true
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

// This method attempts to parse a number element. It returns the number
// element and whether or not the number element was successfully parsed.
func (v *parser) parseNumber() (elements.Number, bool) {
	var ok bool
	var number elements.Number
	var token = v.nextToken()
	if token.typ != tokenNumber {
		v.backupOne()
		return number, false
	}
	number, ok = elements.NumberFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid number token was found: %v", token))
	}
	return number, true
}

// This method attempts to parse a pattern element. It returns the pattern
// element and whether or not the pattern element was successfully parsed.
func (v *parser) parsePattern() (elements.Pattern, bool) {
	var ok bool
	var pattern elements.Pattern
	var token = v.nextToken()
	if token.typ != tokenPattern {
		v.backupOne()
		return pattern, false
	}
	pattern, ok = elements.PatternFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid pattern token was found: %v", token))
	}
	return pattern, true
}

// This method attempts to parse a parameter. It returns the parameter and
// whether or not the parameter was successfully parsed.
func (v *parser) parseParameter() (*Parameter, bool) {
	var ok bool
	var name elements.Symbol
	var value any
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

// This method attempts to parse a percentage element. It returns the percentage
// element and whether or not the percentage element was successfully parsed.
func (v *parser) parsePercentage() (elements.Percentage, bool) {
	var ok bool
	var percentage elements.Percentage
	var token = v.nextToken()
	if token.typ != tokenPercentage {
		v.backupOne()
		return percentage, false
	}
	percentage, ok = elements.PercentageFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid percentage token was found: %v", token))
	}
	return percentage, true
}

// This method attempts to parse a primitive. It returns the primitive and
// whether or not the primitive was successfully parsed.
func (v *parser) parsePrimitive() (any, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var primitive any
	primitive, ok = v.parseElement()
	if !ok {
		primitive, ok = v.parseString()
	}
	if !ok {
		primitive = nil
	}
	return primitive, ok
}

func (v *parser) parseProbability() (elements.Probability, bool) {
	var ok bool
	var probability elements.Probability
	var token = v.nextToken()
	if token.typ != tokenProbability {
		v.backupOne()
		return probability, false
	}
	probability, ok = elements.ProbabilityFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid probability token was found: %v", token))
	}
	return probability, true
}

// This method attempts to parse a procedure. It returns the procedure and
// whether or not the procedure was successfully parsed.
func (v *parser) parseProcedure() ([]*Statement, bool) {
	var ok bool
	var statements []*Statement
	_, ok = v.parseDelimiter("{")
	if !ok {
		return nil, false
	}
	statements, ok = v.parseStatements()
	if !ok {
		panic("Expected statements following the '{' character.")
	}
	_, ok = v.parseDelimiter("}")
	if !ok {
		panic("Expected a '}' character following the statements.")
	}
	return statements, true
}

// This method attempts to parse a quote string. It returns the quote string
// and whether or not the quote string was successfully parsed.
func (v *parser) parseQuote() (strings.Quote, bool) {
	var ok bool
	var quote strings.Quote
	var token = v.nextToken()
	if token.typ != tokenQuote {
		v.backupOne()
		return quote, false
	}
	quote, ok = strings.QuoteFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid quote token was found: %v", token))
	}
	return quote, true
}

// This method attempts to parse a slice collection. It returns the slice
// collection and whether or not the slice collection was successfully parsed.
func (v *parser) parseSlice() (abstractions.SliceLike[any], bool) {
	var ok bool
	var t *token
	var first any
	var last any
	first, _ = v.parseValue() // The first value is optional.
	t, ok = v.parseDelimiter("..")
	if !ok {
		t, ok = v.parseDelimiter("<..")
	}
	if !ok {
		t, ok = v.parseDelimiter("<..<")
	}
	if !ok {
		t, ok = v.parseDelimiter("..<")
	}
	if !ok {
		// This is not a slice collection.
		if first != nil {
			v.backupOne() // Put back the Value token.
		}
		return nil, false
	}
	var connector = t.val
	last, _ = v.parseValue() // The last value is optional.
	var slice = collections.Slice(first, connector, last)
	return slice, true
}

// This method attempts to parse a statement. It returns the statement and
// whether or not the statement was successfully parsed.
func (v *parser) parseStatement() (*Statement, bool) {
	var ok bool
	var statement *Statement
	return statement, ok
}

// This method attempts to parse the statements within a procedure. It returns
// an array of the statements and whether or not the statements were
// successfully parsed.
func (v *parser) parseStatements() ([]*Statement, bool) {
	var statement *Statement
	var statements []*Statement
	var _, ok = v.parseEOL()
	if !ok {
		// The statements are on a single line.
		statement, ok = v.parseStatement()
		// There must be at least one statement.
		if !ok {
			panic("Expected at least one statement in the component context.")
		}
		for {
			statements = append(statements, statement)
			// Every subsequent statement must be preceded by a ';'.
			_, ok = v.parseDelimiter(";")
			if !ok {
				// No more statements.
				break
			}
			statement, ok = v.parseStatement()
			if !ok {
				panic("Expected a statement after the ';' character.")
			}
		}
		return statements, true
	}
	// The statements are on separate lines.
	for {
		statement, ok = v.parseStatement()
		if !ok {
			// No more statements.
			break
		}
		statements = append(statements, statement)
		// Every statement must be followed by an EOL.
		_, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the statement.")
		}
	}
	// There must be at least one statement.
	if len(statements) == 0 {
		panic("Expected at least one statement in the component context.")
	}
	return statements, true
}

// This method attempts to parse a resource element. It returns the
// resource element and whether or not the resource element was
// successfully parsed.
func (v *parser) parseResource() (elements.Resource, bool) {
	var ok bool
	var resource elements.Resource
	var token = v.nextToken()
	if token.typ != tokenResource {
		v.backupOne()
		return resource, false
	}
	resource, ok = elements.ResourceFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid resource token was found: %v", token))
	}
	return resource, true
}

// This method attempts to parse a collection of items. It returns the
// collection and whether or not the collection was successfully parsed.
func (v *parser) parseCollection() (any, bool) {
	var ok bool
	var bad *token
	var sequence any
	_, ok = v.parseDelimiter("[")
	if !ok {
		return nil, false
	}
	sequence, ok = v.parseSequence()
	if !ok {
		panic("Expected a sequence following the '[' character.")
	}
	bad, ok = v.parseDelimiter("]")
	if !ok {
		panic(fmt.Sprintf("Expected a ']' character following the sequence and received: %v", *bad))
	}
	return sequence, true
}

// This method attempts to parse a string primitive. It returns the
// string primitive and whether or not the string primitive was
// successfully parsed.
func (v *parser) parseString() (any, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var str any
	str, ok = v.parseBinary()
	if !ok {
		str, ok = v.parseMoniker()
	}
	if !ok {
		str, ok = v.parseNarrative()
	}
	if !ok {
		str, ok = v.parseQuote()
	}
	if !ok {
		str, ok = v.parseSymbol()
	}
	if !ok {
		str, ok = v.parseVersion()
	}
	if !ok {
		str = nil
	}
	return str, ok
}

// This method attempts to parse a probability element. It returns the
// probability element and whether or not the probability element was
// successfully parsed.
// This method attempts to parse a symbol string. It returns the symbol
// string and whether or not the symbol string was successfully parsed.
func (v *parser) parseSymbol() (elements.Symbol, bool) {
	var ok bool
	var symbol elements.Symbol
	var token = v.nextToken()
	if token.typ != tokenSymbol {
		v.backupOne()
		return symbol, false
	}
	symbol, ok = elements.SymbolFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid symbol token was found: %v", token))
	}
	return symbol, true
}

// This method attempts to parse a tag element. It returns the
// tag element and whether or not the tag element was
// successfully parsed.
func (v *parser) parseTag() (elements.Tag, bool) {
	var ok bool
	var tag elements.Tag
	var token = v.nextToken()
	if token.typ != tokenTag {
		v.backupOne()
		return tag, false
	}
	tag, ok = elements.TagFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid tag token was found: %v", token))
	}
	return tag, true
}

// This method attempts to parse a primitive value. It returns the primitive
// value and whether or not the primitive value was successfully parsed.
func (v *parser) parseValue() (any, bool) {
	var ok bool
	var value any
	value, ok = v.parseElement()
	if !ok {
		value, ok = v.parseString()
	}
	if !ok {
		value, ok = v.parseIdentifier()
	}
	if !ok {
		value = nil
	}
	return value, ok
}

// This method attempts to parse a version string. It returns the version
// string and whether or not the version string was successfully parsed.
func (v *parser) parseVersion() (strings.Version, bool) {
	var ok bool
	var version strings.Version
	var token = v.nextToken()
	if token.typ != tokenVersion {
		v.backupOne()
		return version, false
	}
	version, ok = strings.VersionFromString(token.val)
	if !ok {
		panic(fmt.Sprintf("An invalid version token was found: %v", token))
	}
	return version, true
}
