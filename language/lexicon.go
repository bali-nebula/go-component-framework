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
)

// This map captures the syntax rules for the Bali Document Notation™ (BDN)
// language grammar. The lowercase identifiers define rules for the grammar and
// the UPPERCASE identifiers represent tokens returned by the scanner. The
// official definition of the language grammar is here:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification
//
// This map is useful when creating scanner and parser error messages.
var lexicon = map[string]string{
	"$acceptClause": `"accept" expression`,
	"$annotation":   `NOTE | COMMENT`,
	"$arguments":    `expression {"," expression}`,
	"$arithmetic":   `expression ("*" | "/" | "//" | "+" | "-") expression`,
	"$association":  `primitive ":" component`,
	"$attribute":    `variable "[" indices "]"`,
	"$breakClause":  `"break" "loop"`,
	"$catalog": `
    association {"," association} |  ! Inline, no NOTEs allowed.
    EOL <association EOL> |
    ":"  ! An empty catalog.`,
	"$chaining":       `expression "&" expression`,
	"$checkoutClause": `"checkout" recipient ["at" "level" expression] "from" expression`,
	"$collection":     `"[" sequence "]"`,
	"$comparison":     `expression ("<" | "=" | ">" | "≠" | "IS" | "MATCHES") expression`,
	"$complement":     `"NOT" expression`,
	"$component":      `entity [context] [NOTE]`,
	"$context":        `"(" parameters ")"`,
	"$continueClause": `"continue" "loop"`,
	"$dereference":    `"@" expression`,
	"$discardClause":  `"discard" expression`,
	"$document":       `component EOL EOF`,
	"$element": `
    ANGLE | BOOLEAN | DURATION | MOMENT | NUMBER | PATTERN |
    PERCENTAGE | PROBABILITY | RESOURCE | SYMBOL | TAG`,
	"$entity":          `element | string | collection | procedure`,
	"$evaluateClause":  `[recipient (":=" | "+=" | "-=" | "*=" | "/=")] expression`,
	"$exception":       `SYMBOL`,
	"$exceptionClause": `"on" exception <"matching" expression "do" "{" statements "}">`,
	"$exponential":     `expression "^" expression`,
	"$expression": `
    component |
    variable |
    intrinsic |
    precedence |
    dereference |
    invocation |
    value |
    chaining |
    exponential |
    inversion |
    arithmetic |
    magnitude |
    comparison |
    complement |
    logical`,
	"$function":   `IDENTIFIER`,
	"$ifClause":   `"if" expression "do" "{" statements "}"`,
	"$indices":    `expression {"," expression}`,
	"$intrinsic":  `function "(" [arguments] ")"`,
	"$inversion":  `("-" | "/" | "*") expression`,
	"$invocation": `expression ("." | "<~") message "(" [arguments] ")"`,
	"$item":       `SYMBOL`,
	"$list": `
    component {"," component} |  ! Inline, no NOTEs allowed.
    EOL <component EOL> |
    ! An empty list.`,
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
	"$message":        `IDENTIFIER`,
	"$name":           `SYMBOL`,
	"$notarizeClause": `"notarize" expression "as" expression`,
	"$parameter":      `name ":" component`,
	"$parameters": `
    parameter {"," parameter} |  ! Inline, no NOTEs allowed.
    EOL <parameter EOL>`,
	"$postClause":     `"post" expression "to" expression`,
	"$precedence":     `"(" expression ")"`,
	"$primitive":      `element | string`,
	"$procedure":      `"{" statements "}"`,
	"$publishClause":  `"publish" expression`,
	"$range":          `[primitive] (".." | "..<" | "<..<" | "<..") [primitive]`,
	"$recipient":      `name | attribute`,
	"$rejectClause":   `"reject" expression`,
	"$retrieveClause": `"retrieve" recipient "from" expression`,
	"$returnClause":   `"return" expression`,
	"$saveClause":     `"save" expression "as" recipient`,
	"$selectClause":   `"select" expression <"matching" expression "do" "{" statements "}">`,
	"$sequence":       `catalog | list | range`,
	"$statement":      `mainClause [exceptionClause]`,
	"$statements": `
    statement {";" statement} |
    EOL {(annotation | statement) EOL} |
    ! An empty procedure.`,
	"$string":      `BINARY | MONIKER | NARRATIVE | QUOTE | VERSION`,
	"$throwClause": `"throw" expression`,
	"$value":       `expression "[" indices "]"`,
	"$variable":    `IDENTIFIER`,
	"$whileClause": `"while" expression "do" "{" statements "}"`,
	"$withClause":  `"with" "each" item "in" expression "do" "{" statements "}"`,
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
    "}" | "|" | "{" | "^" | "]" | "[" | "@" | "?=" | ">" | "=" |
    "<..<" | "<.." | "<~" | "<" | ";" | ":=" | ":" | "/=" | "//" | "/" |
    "..<" | ".." | "." | "-=" | "-" | "," | "+=" | "+" | "*=" | "*" |
    ")" | "(" | "&"`,
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
	"$MOMENT":      `"<" YEAR ["-" MONTH ["-" DAY ["T" HOUR [":" MINUTE [":" SECOND [FRACTION]]]]]] ">"`,
	"$MONIKER":     `<"/" NAME>`,
	"$MONTH":       `"0" "1".."9" | "1" "0".."2"`,
	"$NAME":        `LETTER {[SEPARATOR] (LETTER | DIGIT)}`,
	"$NARRATIVE":   `'">' EOL {NARRATIVE | ~'<"'} EOL {SPACE} '<"'`,
	"$NONE":        `"none"`,
	"$NOTE":        `"! " <~EOL>`,
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
	"$YEAR":    `[SIGN] ORDINAL`,
	"$ZERO":    `"0"`,
}

// PRIVATE FUNCTIONS

func generateGrammar(symbols ...string) string {
	var grammar = "The expected grammar was:\n"
	for _, s := range symbols {
		grammar += fmt.Sprintf("  %v: %v\n\n", s, lexicon[s])
	}
	return grammar
}
