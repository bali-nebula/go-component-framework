/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package grammar

import (
	"github.com/craterdog-bali/go-bali-document-notation/elements"
)

// This map captures the syntax rules for the Bali Document Notation™ (BDN)
// language grammar. The lowercase identifiers define rules for the grammar and
// the UPPERCASE identifiers represent tokens returned by the scanner. The token
// definitions and their associated regular expressions can be found here:
//
//	../abstractions/grammar.go
//
// This map is useful when creating parser error messages.
var Grammar = map[string]string{
	// Components
	"$component": `entity context?`,

	"$entity": `primitive | collection | procedure`,

	"$primitive": `element | string`,

	"$collection": `"[" sequence "]"`,

	"$sequence": `catalog | slice | list`,

	"$procedure": `"{" statements "}"`,

	"$statements": `
		statement (";" statement)* |
		EOL ((documentation | statement) EOL)* |
		! No statements.
`,

	"$documentation": `NOTE | COMMENT`,

	"$context": `"(" parameters ")"`,

	"$parameters": `
		parameter ("," parameter)* |
		EOL (parameter EOL)+
`,

	"$parameter": `name ":" component`,

	"$name": `SYMBOL`,

	// Primitives
	"$element": `
		ANGLE | BOOLEAN | DURATION | MOMENT | NUMBER | PATTERN |
		PERCENTAGE | PROBABILITY | RESOURCE | SYMBOL | TAG
`,

	"$string": `
		BINARY | MONIKER | NARRATIVE | QUOTE | VERSION
`,

	// Collections
	"$list": `
		component ("," component)* |
		EOL (component NOTE? EOL)* |
		! No items.`,

	"$catalog": `
		association ("," association)* |
		EOL (association NOTE? EOL)* |
		":" ! No associations.
`,

	"$association": `primitive ":" component`,

	"$slice": `value? (".." | "..<" | "<..<" | "<..") value?`,

	"$value": `primitive | variable`,

	"$variable": `IDENTIFIER`,

	// Statements
	"$statement": `mainClause handleClause?`,

	"$mainClause": `
		evaluateClause |
		onClause |
		ifClause |
		withClause |
		whileClause |
		continueClause |
		breakClause |
		returnClause |
		throwClause |
		saveClause |
		discardClause |
		notarizeClause |
		checkoutClause |
		publishClause |
		postClause |
		retrieveClause |
		acceptClause |
		rejectClause
`,

	"$evaluateClause": `(recipient (":=" | "+=" | "-=" | "*=" | "/="))? expression`,

	"$recipient": `name | attribute`,

	"$attribute": `variable "[" indices "]"`,

	"$indices": `expression ("," expression)*`,

	"$onClause": `"on" expression ("matching" expression "do" block)+`,

	"$block": `"{" statements "}"`,

	"$ifClause": `"if" (expression "do" block)+`,

	"$withClause": `"with" ("each" item "in")? expression "do" block`,

	"$item": `SYMBOL`,

	"$whileClause": `"while" expression "do" block`,

	"$continueClause": `"continue" "loop"`,

	"$breakClause": `"break" "loop"`,

	"$returnClause": `"return" expression?`,

	"$throwClause": `"throw" expression`,

	"$saveClause": `"save" expression ("as" recipient)?`,

	"$discardClause": `"discard" expression`,

	"$notarizeClause": `"notarize" expression "as" expression`,

	"$checkoutClause": `"checkout" recipient ("at" "level" expression)? "from" expression`,

	"$publishClause": `"publish" expression`,

	"$postClause": `"post" expression "to" expression`,

	"$retrieveClause": `"retrieve" recipient "from" expression`,

	"$acceptClause": `"accept" expression`,

	"$rejectClause": `"reject" expression`,

	"$handleClause": `"handle" exception ("matching" expression "with" block)+`,

	"$exception": `SYMBOL`,

	// Expressions
	"$expression": `
		component |
		variable |
		functionExpression |
		precedenceExpression |
		dereferenceExpression |
		messageExpression |
		attributeExpression |
		chainExpression |
		powerExpression |
		inversionExpression |
		arithmeticExpression |
		magnitudeExpression |
		comparisonExpression |
		complementExpression |
		logicalExpression |
		defaultExpression
`,

	"$functionExpression": `function "(" arguments? ")"`,

	"$function": `IDENTIFIER`,

	"$arguments": `expression ("," expression)*`,

	"$precedenceExpression": `"(" expression ")"`,

	"$dereferenceExpression": `"@" expression`,

	"$messageExpression": `expression ("." | "<-") message "(" arguments? ")"`,

	"$message": `IDENTIFIER`,

	"$attributeExpression": `expression "[" indices "]"`,

	"$chainExpression": `expression "&" expression`,

	"$powerExpression": `expression "^" expression {right associative}`,

	"$inversionExpression": `("-" | "/" | "*") expression`,

	"$arithmeticExpression": `expression ("*" | "/" | "//" | "+" | "-") expression`,

	"$magnitudeExpression": `"|" expression "|"`,

	"$comparisonExpression": `expression ("<" | "=" | ">" | "IS" | "MATCHES") expression`,

	"$complementExpression": `"NOT" expression`,

	"$logicalExpression": `expression ("AND" | "SANS" | "XOR" | "OR") expression`,

	"$defaultExpression": `expression "?" expression`,
}

// COMPONENT NODES

// This type defines the node structure associated with a component.
type Component struct {
	Entity     any // A entity is a primitive, collection or procedure.
	Parameters []*Parameter
	Note       string
}

// This type defines the node structure associated with a name-value pair. It is
// used by a component to maintain its parameters.
type Parameter struct {
	Name  elements.Symbol // The name of the parameter.
	Value any             // The value of the parameter.
}

// This type defines the node structure associated with a procedure that
// contains Bali Document Notation™ (BDN) procedural statements.
type Procedure struct {
	Statements []any
}

// STATEMENT NODES

// This type defines the node structure associated with a Bali Document
// Notation (BDN) statement containing a main clause and an optional
// exception handling clause.
type Statement struct {
	MainClause   any
	HandleClause *HandleClause
}

// This type defines the node structure associated with a clause that evaluates
// an expression and optionally assigns the result to a recipient. The recipient
// must support the `Scalable` interface.
type EvaluateClause struct {
	Recipient  any // The recipient is a symbol or attribute.
	Operator   string
	Expression any
}

// This type defines the node structure associated with a clause that selects a
// statement block to be executed based on the pattern of a control expression.
type OnClause struct {
	Control  any
	Patterns []any
	Blocks   []*Block
}

// This type defines the node structure associated with a clause that selects a
// statement block to be executed based on the value of a conditional expression.
type IfClause struct {
	Conditions []any
	Blocks     []*Block
}

// This type defines the node structure associated with a clause that executes
// a statement block for each item in a collection. Each item may be optionally
// assigned to a symbol that is referenced in the statement block.
type WithClause struct {
	Item       elements.Symbol
	Collection any
	Block      *Block
}

// This type defines the node structure associated with a clause that executes
// a statement block while a condition is true.
type WhileClause struct {
	Condition any
	Block     *Block
}

// This type defines the node structure associated with a clause that causes the
// execution of a loop to continue at the beginning.
type ContinueClause struct {
}

// This type defines the node structure associated with a clause that causes the
// execution of a loop to end.
type BreakClause struct {
}

// This type defines the node structure associated with a clause that causes an
// executing procedure to return with an optional result.
type ReturnClause struct {
	Result any
}

// This type defines the node structure associated with a clause that causes an
// exception to be thrown in the executing procedure.
type ThrowClause struct {
	Exception any
}

// This type defines the node structure associated with a clause that saves
// a draft document referred to by an expression to the document repository
// and returns a citation to the document which is optionally assigned to a
// recipient.
type SaveClause struct {
	Draft     any
	Recipient any // The recipient is a symbol or attribute.
}

// This type defines the node structure associated with a clause that discards
// a draft document referred to by an expression from the document repository.
type DiscardClause struct {
	Citation any
}

// This type defines the node structure associated with a clause that notarizes
// a draft document as a named released document in the document repository.
type NotarizeClause struct {
	Draft   any
	Moniker any
}

// This type defines the node structure associated with a clause that checks out
// a draft version of a released document at an optional release level from the
// document repository and assigns it to a recipient.
type CheckoutClause struct {
	Recipient any // The recipient is a symbol or attribute.
	Level     any
	Moniker   any
}

// This type defines the node structure associated with a clause that publishes
// an event to be delivered to all interested parties.
type PublishClause struct {
	Event any
}

// This type defines the node structure associated with a clause that posts a
// message to a named message bag.
type PostClause struct {
	Message any
	Bag     any
}

// This type defines the node structure associated with a clause that retrieves
// a random message from a named message bag and assigns it to a recipient.
type RetrieveClause struct {
	Recipient any // The recipient is a symbol or attribute.
	Bag       any
}

// This type defines the node structure associated with a clause that accepts a
// message that was previously retrieved from a named message bag so that it
// cannot be retrieved by another party.
type AcceptClause struct {
	Message any
}

// This type defines the node structure associated with a clause that rejects a
// message that was previously retrieved from a named message bag so that it
// can be retrieved by another party.
type RejectClause struct {
	Message any
}

// This type defines the node structure associated with a clause that handles
// an exception.
type HandleClause struct {
	Exception elements.Symbol
	Patterns  []any
	Blocks    []*Block
}

// This type defines the node structure associated with a block of statements
// that contains Bali Document Notation (BDN) procedural statements.
type Block struct {
	Statements []any
}

// This type defines the node structure associated with an indexed attribute
// within a composite component.
type Attribute struct {
	Variable string
	Indices  []any
}

// EXPRESSION NODES (highest precedence to lowest precedence)

// This type defines the node structure associated with an expression made up of
// a single component.
type ComponentExpression struct {
	Component any
}

// This type defines the node structure associated with an expression that
// returns the value of a variable.
type VariableExpression struct {
	Variable string
}

// This type defines the node structure associated with an expression that
// returns the result of a function call.
type FunctionExpression struct {
	Function  string
	Arguments []any
}

// This type defines the node structure associated with an expression that
// takes precdence over its surrounding expressions.
type PrecedenceExpression struct {
	Expression any
}

// This type defines the node structure associated with an expression that
// returns the value that the expression references.
type DereferenceExpression struct {
	Expression any
}

// This type defines the node structure associated with an expression that
// sends a message to a target component. The message can be sent synchronously
// or asynchronously.
type MessageExpression struct {
	Target    any
	Operator  string
	Message   string
	Arguments []any
}

// This type defines the node structure associated with an expression that
// returns an indexed attribute from within a composite component.
type AttributeExpression struct {
	Composite any
	Indices   []any
}

// This type defines the node structure associated with an expression that
// returns the concatenation of two values.
type ChainExpression struct {
	Left  any
	Right any
}

// This type defines the node structure associated with an expression that
// returns the power of a base value raised to an exponential value.
type PowerExpression struct {
	Base     any
	Exponent any
}

// This type defines the node structure associated with an expression that
// returns the inverse of a value.
type InversionExpression struct {
	Operator string
	Numeric  any
}

// This type defines the node structure associated with an expression that
// returns the result of an arithmetic operation on two values.
type ArithmeticExpression struct {
	Left     any
	Operator string
	Right    any
}

// This type defines the node structure associated with an expression that
// returns the magnitude of a value.
type MagnitudeExpression struct {
	Numeric any
}

// This type defines the node structure associated with an expression that
// returns a comparison of two values.
type ComparisonExpression struct {
	Left     any
	Operator string
	Right    any
}

// This type defines the node structure associated with an expression that
// returns the logical complement of a value.
type ComplementExpression struct {
	Logical any
}

// This type defines the node structure associated with an expression that
// returns the result of a logical operation on two values.
type LogicalExpression struct {
	Left     any
	Operator string
	Right    any
}

// This type defines the node structure associated with an expression that
// returns a value or its default value if not set.
type DefaultExpression struct {
	Left  any
	Right any
}
