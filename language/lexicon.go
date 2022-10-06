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
	"github.com/craterdog-bali/go-bali-document-notation/elements"
)

// This map captures the syntax rules for the Bali Document Notation™ (BDN)
// language grammar. The lowercase identifiers define rules for the grammar and
// the UPPERCASE identifiers represent tokens returned by the scanner. The token
// definitions and their associated regular expressions can be found here:
//
//	../abstractions/language.go
//
// This map is useful when creating parser error messages.
var Grammar = map[string]string{
	"$acceptClause":         `"accept" expression`,
	"$arguments":            `expression {"," expression}`,
	"$arithmeticExpression": `expression ("*" | "/" | "//" | "+" | "-") expression`,
	"$association":          `primitive ":" component`,
	"$attribute":            `variable "[" indices "]"`,
	"$attributeExpression":  `expression "[" indices "]"`,
	"$block":                `"{" statements "}"`,
	"$breakClause":          `"break" "loop"`,
	"$catalog": `
		association {"," association} |
		EOL <association [NOTE] EOL> |
		":"  ! An empty catalog.
	`,
	"$chainExpression":       `expression "&" expression`,
	"$checkoutClause":        `"checkout" recipient ["at" "level" expression] "from" expression`,
	"$collection":            `"[" sequence "]"`,
	"$commentary":            `NOTE | COMMENT`,
	"$comparisonExpression":  `expression ("<" | "=" | ">" | "≠" | "IS" | "MATCHES") expression`,
	"$complementExpression":  `"NOT" expression`,
	"$component":             `entity [context]`,
	"$context":               `"(" parameters ")"`,
	"$continueClause":        `"continue" "loop"`,
	"$defaultExpression":     `expression "?" expression`,
	"$dereferenceExpression": `"@" expression`,
	"$discardClause":         `"discard" expression`,
	"$document":              `component EOL EOF`,
	"$element": `
		ANGLE | BOOLEAN | DURATION | MOMENT | NUMBER | PATTERN |
		PERCENTAGE | PROBABILITY | RESOURCE | SYMBOL | TAG
	`,
	"$entity":         `element | string | collection | procedure`,
	"$evaluateClause": `[recipient (":=" | "+=" | "-=" | "*=" | "/=")] expression`,
	"$exception":      `SYMBOL`,
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
	"$function":            `IDENTIFIER`,
	"$functionExpression":  `function "(" [arguments] ")"`,
	"$handleClause":        `"handle" exception <"matching" expression "with" block>`,
	"$ifClause":            `"if" <expression "do" block>`,
	"$indices":             `expression {"," expression}`,
	"$inversionExpression": `("-" | "/" | "*") expression`,
	"$item":                `SYMBOL`,
	"$list": `
		component {"," component} |
		EOL <component [NOTE] EOL> |
		! An empty list.
	`,
	"$logicalExpression":   `expression ("AND" | "SANS" | "XOR" | "OR") expression`,
	"$magnitudeExpression": `"|" expression "|"`,
	"$mainClause": `
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
		rejectClause |
		evaluateClause
	`,
	"$message":           `IDENTIFIER`,
	"$messageExpression": `expression ("." | "<-") message "(" [arguments] ")"`,
	"$name":              `SYMBOL`,
	"$notarizeClause":    `"notarize" expression "as" expression`,
	"$onClause":          `"on" expression <"matching" expression "do" block>`,
	"$parameter":         `name ":" component`,
	"$parameters": `
		parameter {"," parameter} |
		EOL <parameter EOL>
	`,
	"$postClause":           `"post" expression "to" expression`,
	"$powerExpression":      `expression "^" expression`,
	"$precedenceExpression": `"(" expression ")"`,
	"$primitive":            `element | string`,
	"$procedure":            `"{" statements "}"`,
	"$publishClause":        `"publish" expression`,
	"$recipient":            `name | attribute`,
	"$rejectClause":         `"reject" expression`,
	"$retrieveClause":       `"retrieve" recipient "from" expression`,
	"$returnClause":         `"return" [expression]`,
	"$saveClause":           `"save" expression ["as" recipient]`,
	"$sequence":             `catalog | list | slice`,
	"$slice":                `[value] (".." | "..<" | "<..<" | "<..") [value]`,
	"$statement":            `mainClause [handleClause]`,
	"$statements": `
		statement {";" statement} |
		EOL {(commentary | statement) EOL} |
		! An empty procedure.
	`,
	"$string":      `BINARY | MONIKER | NARRATIVE | QUOTE | VERSION`,
	"$throwClause": `"throw" expression`,
	"$value":       `element | string | variable`,
	"$variable":    `IDENTIFIER`,
	"$whileClause": `"while" expression "do" block`,
	"$withClause":  `"with" ["each" item "in"] expression "do" block`,
	"$ANGLE":       `"~" (REAL | ZERO)`,
	"$ANY":         `"any"`,
	"$AUTHORITY":   `<~"/">`,
	"$BASE16":      `"0".."9" | "a".."f"`,
	"$BASE32":      `"0".."9" | "A".."D" | "F".."H" | "J".."N" | "P".."T" | "V".."Z"`,
	"$BASE64":      `"A".."Z" | "a".."z" | "0".."9" | "+" | "/"`,
	"$BINARY":      `"'" {BASE64 | WHITESPACE} "'"`,
	"$BOOLEAN":     `"false" | "true"`,
	"$COMMENT":     `"!>" EOL  {COMMENT | ~"<!"} EOL {TAB} "<!"`,
	"$DATES":       `[TSPAN "Y"] [TSPAN "M"] [TSPAN "D"]`,
	"$DAY":         `"0".."2" "1".."9" | "3" "0".."1"`,
	"$DELIMITER": `
		"}" | "|" | "{" | "^" | "]" | "[" | "@" | "?" | ">" | "=" |
		"<..<" | "<.." | "<-" | "<" | ";" | ":=" | ":" | "/=" | "//" | "/" |
		"..<" | ".." | "." | "-=" | "-" | "," | "+=" | "+" | "*=" | "*" |
		")" | "(" | "&"
	`,
	"$DURATION":   `"~" [SIGN] "P" (WEEKS | DATES [TIMES])`,
	"$E":          `"e"`,
	"$EOL":        `"\n"`,
	"$ESCAPE":     `'\' ('\' | 'a' | 'b' | 'f' | 'n' | 'r' | 't' | 'v' | '"' | "'" | UNICODE)`,
	"$EXPONENT":   `"E" [SIGN] ORDINAL`,
	"$FRACTION":   `"." <"0".."9">`,
	"$FRAGMENT":   `{~">"}`,
	"$HOUR":       `"0".."1" "0".."9" | "2" "0".."3"`,
	"$IDENTIFIER": `LETTER (LETTER | DIGIT)*`,
	"$IMAGINARY":  ` [SIGN | REAL] "i"`,
	"$INFINITY":   `"infinity" | "∞"`,
	"$KEYWORD": `
		"with" | "while" | "to" | "throw" | "select" | "save" |
		"return" | "retrieve" | "reject" | "publish" | "post" |
		"notarize" | "matching" | "loop" | "level" | "in" | "if" |
		"handle" | "from" | "each" | "do" | "discard" | "continue" |
		"checkout" | "break" | "at" | "as" | "any" | "accept" |
		"XOR" | "SANS" | "OR" | "NOT" | "MATCHES" | "IS" | "AND"
	`,
	"$MINUTE":      `"0".."5" "0".."9"`,
	"$MOMENT":      `"<" YEAR ["-" MONTH ["-" DAY ["T" HOUR [":" MINUTE [":" SECOND [FRACTION]]]]]] ">"`,
	"$MONIKER":     `<"/" NAME>`,
	"$MONTH":       `"0" "1".."9" | "1" "0".."2"`,
	"$NAME":        `LETTER {[SEPARATOR] (LETTER | DIGIT)}`,
	"$NARRATIVE":   `'">' EOL {NARRATIVE | ~'<"'} EOL {TAB} '<"'`,
	"$NONE":        `"none"`,
	"$NOTE":        `"! " {~EOL}`,
	"$NUMBER":      `IMAGINARY | REAL | ZERO | INFINITY | UNDEFINED | "(" (RECTANGULAR | POLAR) ")"`,
	"$ONE":         `"1."`,
	"$ORDINAL":     `"1".."9" {"0".."9"}`,
	"$PATH":        `{~("?" | "#" | ">")}`,
	"$PATTERN":     `NONE | REGEX | ANY`,
	"$PERCENTAGE":  `(REAL | ZERO) "%"`,
	"$PHI":         `"phi" | "φ"`,
	"$PI":          `"pi" | "π"`,
	"$POLAR":       `REAL "e^" ANGLE "i"`,
	"$PROBABILITY": `FRACTION | ONE`,
	"$QUERY":       `{~("#" | ">")}`,
	"$QUOTE":       `'"' {RUNE} '"'`,
	"$REAL":        `[SIGN] (E | PI | PHI | TAU | SCALAR)`,
	"$RECTANGULAR": ` REAL ", " IMAGINARY`,
	"$REGEX":       `'"' <RUNE> '"?'`,
	"$RESOURCE":    `"<" SCHEME ":" ["//" AUTHORITY] "/" PATH ["?" QUERY] ["#" FRAGMENT] ">"`,
	"$RUNE":        `ESCAPE | ~("\r" | "\n")`,
	"$SCALAR":      `(ORDINAL [FRACTION] | ZERO FRACTION) [EXPONENT]`,
	"$SCHEME":      `("a".."z" | "A".."Z") {"a".."z" | "A".."Z" | "0".."9" | "+" | "-" | "."}`,
	"$SECOND":      `"0".."5" "0".."9" | "6" "0".."1"`,
	"$SEPARATOR":   `"-" | "+" | "."`,
	"$SIGN":        `"+" | "-"`,
	"$SYMBOL":      `"$" <IDENTIFIER>`,
	"$TAB":         `"\t"`,
	"$TAG":         `"#" <BASE32>`,
	"$TAU":         `"tau" | "τ"`,
	"$TIMES":       `"T" [TSPAN "H"] [TSPAN "M"] [TSPAN "S"]`,
	"$TSPAN":       `ZERO | ORDINAL [FRACTION]`,
	"$UNDEFINED":   `"undefined"`,
	"$UNICODE": `
		"u" BASE16 BASE16 BASE16 BASE16 |
		"U" BASE16 BASE16 BASE16 BASE16 BASE16 BASE16 BASE16 BASE16 
	`,
	"$VERSION":    `"v" ORDINAL {"." ORDINAL}`,
	"$WEEKS":      `TSPAN "W"`,
	"$WHITESPACE": `" " | "\t" | "\n" | "\r"`,
	"$YEAR":       `[SIGN] ORDINAL`,
	"$ZERO":       `"0"`,
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
