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
)

// This map captures the syntax rules for the Bali Document Notation™ (Bali)
// language grammar. The lowercase identifiers define rules for the grammar and
// the UPPERCASE identifiers represent tokens returned by the scanner. The
// official definition of the language grammar is here:
//
//	https://github.com/bali-nebula/bali-nebula/wiki/Language-Specification
//
// This map is useful when creating scanner and parser error messages.
var grammar = map[string]string{
	"$ANGLE":       `'~' (ZERO | MAGNITUDE)`,
	"$ANY":         `"any"`,
	"$AUTHORITY":   `~('/' | CONTROL)+`,
	"$BASE10":      `'0'..'9'`,
	"$BASE16":      `'0'..'9' | 'a'..'f'`,
	"$BASE32":      `'0'..'9' | 'A'..'D' | 'F'..'H' | 'J'..'N' | 'P'..'T' | 'V'..'Z'`,
	"$BASE64":      `'0'..'9' | 'a'..'z' | 'A'..'Z' | '+' | '/'`,
	"$BINARY":      `"'>" EOL SPACE* BASE64+ EOL* SPACE* "<'"`,
	"$BOOLEAN":     `"false" | "true"`,
	"$BYTECODE":    `"'" (INSTRUCTION (SPACE INSTRUCTION)*)* "'"`,
	"$CHARACTER":   `ESCAPE | ~('"' | CONTROL)`,
	"$COMMENT":     `"!>" ANY* "<!"`,
	"$COMPLEX":     `'(' (RECTANGULAR | POLAR) ')'`,
	"$DAY":         `'0'..'2' '1'..'9' | '3' '0'..'1'`,
	"$DAYS":        `TIMESPAN 'D'`,
	"$DURATION":    `'~' SIGN? 'P' (WEEKS | YEARS? MONTHS? DAYS? ('T' HOURS? MINUTES? SECONDS?)?)`,
	"$E":           `'e'`,
	"$EOL":         `"\n"`,
	"$EXPONENT":    `'E' SIGN? ORDINAL`,
	"$FLOAT":       `SIGN? MAGNITUDE`,
	"$FRACTION":    `'.' BASE10+`,
	"$FRAGMENT":    `~('>' | CONTROL)*`,
	"$HOUR":        `'0'..'1' '0'..'9' | '2' '0'..'3'`,
	"$HOURS":       `TIMESPAN 'H'`,
	"$IDENTIFIER":  `LETTER (LETTER | DIGIT)*`,
	"$IMAGINARY":   `SIGN? MAGNITUDE? 'i'`,
	"$INFINITY":    `SIGN? ("infinity" | '∞')`,
	"$INSTRUCTION": `BASE16{4}`,
	"$LETTER":      `LOWER_CASE | UPPER_CASE`,
	"$MAGNITUDE":   `E | PI | PHI | TAU | SCALAR`,
	"$MINUTE":      `'0'..'5' '0'..'9'`,
	"$MINUTES":     `TIMESPAN 'M'`,
	"$MOMENT":      `'<' SIGN? YEAR ('-' MONTH ('-' DAY ('T' HOUR (':' MINUTE (':' SECOND FRACTION?)?)?)?)?)? '>'`,
	"$MONIKER":     `('/' NAME)+`,
	"$MONTH":       `'0' '1'..'9' | '1' '0'..'2'`,
	"$MONTHS":      `TIMESPAN 'M'`,
	"$NAME":        `LETTER (SEPARATOR? (LETTER | DIGIT))*`,
	"$NARRATIVE":   `'"' '>' EOL (NARRATIVE | ~('<' '"'))* EOL SPACE* '<' '"'`,
	"$NONE":        `"none"`,
	"$NOTE":        `"! " (~CONTROL)*`,
	"$NUMBER":      `REAL | IMAGINARY | COMPLEX`,
	"$ONE":         `"1."`,
	"$ORDINAL":     `'1'..'9' '0'..'9'*`,
	"$PATH":        `~('?' | '#' | '>' | CONTROL)*`,
	"$PATTERN":     `NONE | REGEX | ANY`,
	"$PERCENTAGE":  `REAL '%'`,
	"$PHI":         `"phi" | 'φ'`,
	"$PI":          `"pi" | 'π'`,
	"$POLAR":       `MAGNITUDE "e^" ANGLE 'i'`,
	"$PROBABILITY": `FRACTION | ONE`,
	"$QUERY":       `~('#' | '>' | CONTROL)*`,
	"$QUOTE":       `'"' CHARACTER* '"'`,
	"$REAL":        `ZERO | FLOAT | INFINITY | UNDEFINED`,
	"$RECTANGULAR": `FLOAT ", " FLOAT 'i'`,
	"$REGEX":       `'"' CHARACTER+ '"' '?'`,
	"$RESOURCE":    `'<' SCHEME ':' ("//" AUTHORITY)? '/' PATH ('?' QUERY)? ('#' FRAGMENT)? '>'`,
	"$SCALAR":      `(ZERO FRACTION | ORDINAL FRACTION?) EXPONENT?`,
	"$SCHEME":      `('a'..'z' | 'A'..'Z') ('0'..'9' | 'a'..'z' | 'A'..'Z' | '+' | '-' | '.')*`,
	"$SECOND":      `'0'..'5' '0'..'9' | '6' '0'..'1'`,
	"$SECONDS":     `TIMESPAN 'S'`,
	"$SIGN":        `'+' | '-'`,
	"$SPACE":       `' '`,
	"$SYMBOL":      `'$' IDENTIFIER`,
	"$TAG":         `'#' BASE32+`,
	"$TAU":         `"tau" | 'τ'`,
	"$TIMESPAN":    `ZERO | ORDINAL FRACTION?`,
	"$UNDEFINED":   `"undefined"`,
	"$UNICODE":     `'u' BASE16{4} | 'U' BASE16{8}`,
	"$VERSION":     `'v' ORDINAL ('.' ORDINAL)*`,
	"$WEEKS":       `TIMESPAN 'W'`,
	"$YEAR":        `ZERO | ORDINAL`,
	"$YEARS":       `TIMESPAN 'Y'`,
	"$ZERO":        `'0'`,

	"$acceptClause": `"accept" message`,
	"$annotation":   `NOTE | COMMENT`,
	"$argument":     `expression`,
	"$arguments":    `argument ("," argument)*`,
	"$arithmetic":   `expression ("*" | "/" | "//" | "+" | "-") expression`,
	"$assignment":   `letClause  ! ("let" recipient (":=" | "?=" | "+=" | "-=" | "*=" | "/="))? expression`,
	"$association":  `key ":" value`,
	"$associations": `
      association ("," association)*
    | EOL (association EOL)+
    | ":"  ! No associations.`,
	"$attribute":      `variable "[" indices "]"`,
	"$bag":            `expression`,
	"$breakClause":    `"break" "loop"`,
	"$chaining":       `expression "&" expression`,
	"$checkoutClause": `"checkout" recipient ("at" "level" ordinal)? "from" moniker`,
	"$collection":     `"[" (associations | values) "]"`,
	"$comparison":     `expression ("<" | "=" | ">" | "≠" | "IS" | "MATCHES") expression`,
	"$complement":     `"NOT" expression`,
	"$component":      `entity context? NOTE?`,
	"$composite":      `expression`,
	"$condition":      `expression`,
	"$context":        `"(" parameters ")"`,
	"$continueClause": `"continue" "loop"`,
	"$control": `
      ifClause  ! "if" condition "do" procedure
    | selectClause  ! "select" target ("matching" pattern "do" procedure)+
    | whileClause  ! "while" condition "do" procedure
    | withClause  ! "with" "each" item "in" sequence "do" procedure
    | continueClause  ! "continue" "loop"
    | breakClause  ! "break" "loop"
    | returnClause  ! "return" result
    | throwClause  ! "throw" exception`,
	"$dereference":   `"@" expression`,
	"$discardClause": `"discard" document`,
	"$document":      `expression`,
	"$element":       `ANGLE | BOOLEAN | DURATION | MOMENT | NUMBER | PATTERN | PERCENTAGE | PROBABILITY | RESOURCE`,
	"$entity":        `element | string | range | collection | procedure`,
	"$event":         `expression`,
	"$exception":     `expression`,
	"$exponential":   `expression "^" expression`,
	"$expression": `
      value  ! component
    | intrinsic  ! function "(" arguments? ")"
    | variable  ! IDENTIFIER
    | precedence  ! "(" expression ")"
    | dereference  ! "@" expression
    | invocation  ! target ("." | "<-") method "(" arguments? ")"
    | subcomponent  ! composite "[" indices "]"
    | chaining  ! expression "&" expression
    | exponential  ! expression "^" expression
    | inversion  ! ("-" | "/" | "*") expression
    | arithmetic  ! expression ("*" | "/" | "//" | "+" | "-") expression
    | magnitude  ! "|" expression "|"
    | comparison  ! expression ("<" | "=" | ">" | "≠" | "IS" | "MATCHES") expression
    | complement  ! "NOT" expression
    | logical  ! expression ("AND" | "SANS" | "OR" | "XOR") expression`,
	"$failure":    `SYMBOL`,
	"$function":   `IDENTIFIER`,
	"$ifClause":   `"if" condition "do" procedure`,
	"$index":      `expression`,
	"$indices":    `index ("," index)*`,
	"$intrinsic":  `function "(" arguments? ")"`,
	"$inversion":  `("-" | "/" | "*") expression`,
	"$invocation": `target ("." | "<-") method "(" arguments? ")"`,
	"$item":       `SYMBOL`,
	"$key":        `primitive`,
	"$letClause":  `("let" recipient (":=" | "?=" | "+=" | "-=" | "*=" | "/="))? expression`,
	"$logical":    `expression ("AND" | "SANS" | "OR" | "XOR") expression`,
	"$magnitude":  `"|" expression "|"`,
	"$mainClause": `control | assignment | messaging | repository`,
	"$message":    `expression`,
	"$messaging": `
      postClause  ! "post" message "to" bag
    | retrieveClause  ! "retrieve" recipient "from" bag
    | acceptClause  ! "accept" message
    | rejectClause  ! "reject" message
    | publishClause  ! "publish" event`,
	"$method":         `IDENTIFIER`,
	"$moniker":        `expression`,
	"$name":           `SYMBOL`,
	"$notarizeClause": `"notarize" document "as" moniker`,
	"$onClause":       `"on" failure ("matching" pattern "do" procedure)+`,
	"$ordinal":        `expression`,
	"$parameter":      `name ":" value`,
	"$parameters": `
      parameter ("," parameter)*
    | EOL (parameter EOL)+  ! At least one parameter is required.`,
	"$pattern":       `expression`,
	"$postClause":    `"post" message "to" bag`,
	"$precedence":    `"(" expression ")"`,
	"$primitive":     `element | string`,
	"$procedure":     `"{" statements "}"`,
	"$publishClause": `"publish" event`,
	"$range":         `("[" | "(") primitive ".." primitive (")" | "]")`,
	"$recipient":     `name | attribute`,
	"$rejectClause":  `"reject" message`,
	"$repository": `
      checkoutClause  ! "checkout" recipient ("at" "level" ordinal)? "from" moniker
    | saveClause  ! "save" document "as" recipient
    | discardClause  ! "discard" document
    | notarizeClause  ! "notarize" document "as" moniker`,
	"$result":         `expression`,
	"$retrieveClause": `"retrieve" recipient "from" bag`,
	"$returnClause":   `"return" result`,
	"$saveClause":     `"save" document "as" recipient`,
	"$selectClause":   `"select" target ("matching" pattern "do" procedure)+`,
	"$sequence":       `expression`,
	"$source":         `component EOF  ! EOF is the end-of-file marker.`,
	"$statement":      `(annotation EOL)? mainClause onClause? NOTE?`,
	"$statements": `
      statement (";" statement)*
    | EOL (statement? EOL)+  ! Allows blank lines.
    | " "  ! No statements.`,
	"$string":       `BINARY | BYTECODE | MONIKER | NARRATIVE | QUOTE | SYMBOL | TAG | VERSION`,
	"$subcomponent": `composite "[" indices "]"`,
	"$target":       `expression`,
	"$throwClause":  `"throw" exception`,
	"$value":        `component`,
	"$values": `
      value ("," value)*
    | EOL (value EOL)+
    | " "  ! No values.`,
	"$variable":    `IDENTIFIER`,
	"$whileClause": `"while" condition "do" procedure`,
	"$withClause":  `"with" "each" item "in" sequence "do" procedure`,
}

// PRIVATE FUNCTIONS

func generateGrammar(expected string, symbols ...string) string {
	var message = "Was expecting '" + expected + "' from:\n"
	for _, symbol := range symbols {
		message += fmt.Sprintf("  \033[32m%v: \033[33m%v\033[0m\n\n", symbol, grammar[symbol])
	}
	return message
}
