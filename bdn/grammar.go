/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package bdn

import (
	fmt "fmt"
)

// This map captures the syntax rules for the Bali Document Notation™ (BDN)
// language grammar. The lowercase identifiers define rules for the grammar and
// the UPPERCASE identifiers represent tokens returned by the scanner. The
// official definition of the language grammar is here:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification
//
// This map is useful when creating scanner and parser error messages.
var grammar = map[string]string{
	"$acceptClause": `"accept" message`,
	"$annotation":   `NOTE | COMMENT`,
	"$arguments":    `"(" [expression {"," expression}] ")"`,
	"$arithmetic":   `expression ("*" | "/" | "//" | "+" | "-") expression`,
	"$association":  `primitive ":" component`,
	"$associations": `
    association {"," association} |
    EOL <association EOL> |
    ":"  ! No associations.`,
	"$attribute":      `variable "[" indices "]"`,
	"$bag":            `expression`,
	"$breakClause":    `"break" "loop"`,
	"$catalog":        `"[" associations "]"`,
	"$chaining":       `expression "&" expression`,
	"$checkoutClause": `"checkout" recipient ["at" "level" ordinal] "from" moniker`,
	"$collection":     `catalog | list | range`,
	"$comparison":     `expression ("<" | "=" | ">" | "≠" | "IS" | "MATCHES") expression`,
	"$complement":     `"NOT" expression`,
	"$component":      `entity [context] [NOTE]`,
	"$condition":      `expression`,
	"$context":        `"(" parameters ")"`,
	"$continueClause": `"continue" "loop"`,
	"$dereference":    `"@" expression`,
	"$discardClause":  `"discard" document`,
	"$document":       `expression`,
	"$element": `
    ANGLE | BOOLEAN | DURATION | MOMENT | NUMBER | PATTERN |
    PERCENTAGE | PROBABILITY | RESOURCE | SYMBOL | TAG`,
	"$entity":         `element | string | collection | procedure`,
	"$evaluateClause": `[recipient (":=" | "?=" | "+=" | "-=" | "*=" | "/=")] expression`,
	"$event":          `expression`,
	"$exception":      `expression`,
	"$exponential":    `expression "^" expression`,
	"$expression": `
    component   |  ! entity [context] [NOTE]
    intrinsic   |  ! function arguments
    variable    |  ! IDENTIFIER
    precedence  |  ! "(" expression ")"
    dereference |  ! "@" expression
    invocation  |  ! expression ("." | "<~") method arguments
    value       |  ! expression "[" indices "]"
    chaining    |  ! expression "&" expression
    exponential |  ! expression "^" expression
    inversion   |  ! ("-" | "/" | "*") expression
    arithmetic  |  ! expression ("*" | "/" | "//" | "+" | "-") expression
    magnitude   |  ! "|" expression "|"
    comparison  |  ! expression ("<" | "=" | ">" | "≠" | "IS" | "MATCHES") expression
    complement  |  ! "NOT" expression
    logical        ! expression ("AND" | "SANS" | "XOR" | "OR") expression`,
	"$function":   `IDENTIFIER`,
	"$ifClause":   `"if" condition "do" procedure`,
	"$indices":    `expression {"," expression}`,
	"$intrinsic":  `function arguments`,
	"$inversion":  `("-" | "/" | "*") expression`,
	"$invocation": `target ("." | "<~") method arguments`,
	"$item":       `SYMBOL`,
	"$items": `
    component {"," component} |
    EOL <component EOL>       |
    ! No components.`,
	"$list":      `"[" items "]"`,
	"$logical":   `expression ("AND" | "SANS" | "XOR" | "OR") expression`,
	"$magnitude": `"|" expression "|"`,
	"$mainClause": `
    ifClause |
    selectClause |
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
    evaluateClause`,
	"$message":        `expression`,
	"$method":         `IDENTIFIER`,
	"$moniker":        `expression`,
	"$name":           `SYMBOL`,
	"$notarizeClause": `"notarize" document "as" moniker`,
	"$onClause":       `"on" "$exception" <"matching" pattern "do" procedure>`,
	"$ordinal":        `expression`,
	"$parameter":      `name ":" component`,
	"$parameters": `
    parameter {"," parameter} |
    EOL <parameter EOL>  ! At least one parameter is required.`,
	"$pattern":        `expression`,
	"$postClause":     `"post" message "to" bag`,
	"$precedence":     `"(" expression ")"`,
	"$primitive":      `element | string`,
	"$procedure":      `"{" statements "}"`,
	"$publishClause":  `"publish" event`,
	"$range":          `( "[" | "(" ) [primitive] ".." [primitive] ( ")" | "]" )`,
	"$recipient":      `name | attribute`,
	"$rejectClause":   `"reject" message`,
	"$result":         `expression`,
	"$retrieveClause": `"retrieve" recipient "from" bag`,
	"$returnClause":   `"return" result`,
	"$saveClause":     `"save" document "as" recipient`,
	"$selectClause":   `"select" target <"matching" pattern "do" procedure>`,
	"$sequence":       `expression`,
	"$source":         `component EOF  ! EOF is the end-of-file marker.`,
	"$statement":      `[annotation EOL] [mainClause] [onClause] [NOTE]`,
	"$statements": `
    statement {";" statement} |
    EOL {(annotation | statement) EOL} |
    ! An empty procedure.`,
	"$string":      `BINARY | MONIKER | NARRATIVE | QUOTE | VERSION`,
	"$target":      `expression`,
	"$throwClause": `"throw" exception`,
	"$value":       `expression "[" indices "]"`,
	"$variable":    `IDENTIFIER`,
	"$whileClause": `"while" condition "do" procedure`,
	"$withClause":  `"with" "each" item "in" sequence "do" procedure`,
	"$ANGLE":       `"~" (REAL | ZERO)`,
	"$ANY":         `"any"`,
	"$AUTHORITY":   `<~"/">`,
	"$BASE16":      `"0".."9" | "a".."f"`,
	"$BASE32":      `"0".."9" | "A".."D" | "F".."H" | "J".."N" | "P".."T" | "V".."Z"`,
	"$BASE64":      `"A".."Z" | "a".."z" | "0".."9" | "+" | "/"`,
	"$BINARY":      `"'" {BASE64 | SPACE | EOL} "'"`,
	"$BOOLEAN":     `"false" | "true"`,
	"$COMMENT":     `"!>" EOL  {COMMENT | ~"<!"} EOL {SPACE} "<!"`,
	"$DATES":       `[TIMESPAN "Y"] [TIMESPAN "M"] [TIMESPAN "D"]`,
	"$DAY":         `"0".."2" "1".."9" | "3" "0".."1"`,
	"$DELIMITER": `
    "}" | "|" | "{" | "^" | "]" | "[" | "@" | "?=" | ">" | "=" | "<~" | "<" |
	";" | ":=" | ":" | "/=" | "//" | "/" | ".." | "." | "-=" | "-" | "," |
	"+=" | "+" | "*=" | "*" | ")" | "(" | "&"`,
	"$DURATION":   `"~" [SIGN] "P" (WEEKS | DATES [TIMES])`,
	"$E":          `"e"`,
	"$EOL":        `"\n"`,
	"$ESCAPE":     `'\' ('\' | 'a' | 'b' | 'f' | 'n' | 'r' | 't' | 'v' | '"' | "'" | UNICODE)`,
	"$EXPONENT":   `"E" [SIGN] ORDINAL`,
	"$FRACTION":   `"." <"0".."9">`,
	"$FRAGMENT":   `{~">"}`,
	"$HOUR":       `"0".."1" "0".."9" | "2" "0".."3"`,
	"$IDENTIFIER": `LETTER {LETTER | DIGIT}`,
	"$IMAGINARY":  ` [SIGN | REAL] "i"`,
	"$INFINITY":   `"infinity" | "∞"`,
	"$KEYWORD": `
    "with" | "while" | "to" | "throw" | "select" | "save" |
    "return" | "retrieve" | "reject" | "publish" | "post" |
    "on" | "notarize" | "matching" | "loop" | "level" | "in" |
    "if" | "from" | "each" | "do" | "discard" | "continue" |
    "checkout" | "break" | "at" | "as" | "accept" |
    "XOR" | "SANS" | "OR" | "NOT" | "MATCHES" | "IS" | "AND"`,
	"$MINUTE":      `"0".."5" "0".."9"`,
	"$MOMENT":      `"<" [SIGN] YEAR ["-" MONTH ["-" DAY ["T" HOUR [":" MINUTE [":" SECOND [FRACTION]]]]]] ">"`,
	"$MONIKER":     `<"/" NAME>`,
	"$MONTH":       `"0" "1".."9" | "1" "0".."2"`,
	"$NAME":        `LETTER {[SEPARATOR] (LETTER | DIGIT)}`,
	"$NARRATIVE":   `'">' EOL {NARRATIVE | ~'<"'} EOL {SPACE} '<"'`,
	"$NONE":        `"none"`,
	"$NOTE":        `"! " <~EOL>`,
	"$NUMBER":      `INFINITY | IMAGINARY | REAL | ZERO | UNDEFINED | "(" (RECTANGULAR | POLAR) ")"`,
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
	"$RUNE":        `ESCAPE | ~EOL`,
	"$SCALAR":      `(ORDINAL [FRACTION] | ZERO FRACTION) [EXPONENT]`,
	"$SCHEME":      `("a".."z" | "A".."Z") {"a".."z" | "A".."Z" | "0".."9" | "+" | "-" | "."}`,
	"$SECOND":      `"0".."5" "0".."9" | "6" "0".."1"`,
	"$SEPARATOR":   `"-" | "+" | "."`,
	"$SIGN":        `"+" | "-"`,
	"$SPACE":       `" "`,
	"$SYMBOL":      `"$" IDENTIFIER`,
	"$TAG":         `"#" <BASE32>`,
	"$TAU":         `"tau" | "τ"`,
	"$TIMES":       `"T" [TIMESPAN "H"] [TIMESPAN "M"] [TIMESPAN "S"]`,
	"$TIMESPAN":    `ZERO | ORDINAL [FRACTION]`,
	"$UNDEFINED":   `"undefined"`,
	"$UNICODE": `
    "u" BASE16 BASE16 BASE16 BASE16 |
    "U" BASE16 BASE16 BASE16 BASE16 BASE16 BASE16 BASE16 BASE16`,
	"$VERSION": `"v" ORDINAL {"." ORDINAL}`,
	"$WEEKS":   `TIMESPAN "W"`,
	"$YEAR":    `ORDINAL | ZERO`,
	"$ZERO":    `"0"`,
}

// PRIVATE FUNCTIONS

func generateGrammar(expected string, symbols ...string) string {
	var message = "Was expecting '" + expected + "' from:\n"
	for _, symbol := range symbols {
		message += fmt.Sprintf("  %v: %v\n\n", symbol, grammar[symbol])
	}
	return message
}
