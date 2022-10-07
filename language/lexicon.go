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

// This map captures the syntax rules for the Bali Document Notation™ (BDN)
// language grammar. The lowercase identifiers define rules for the grammar and
// the UPPERCASE identifiers represent tokens returned by the scanner. The
// official definition of the language grammar is here:
//
//	https://github.com/craterdog-bali/bali-nebula/wiki/Language-Specification
//
// This map is useful when creating scanner and parser error messages.
var Lexicon = map[string]string{
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
	"$VERSION": `"v" ORDINAL {"." ORDINAL}`,
	"$WEEKS":   `TSPAN "W"`,
	"$YEAR":    `[SIGN] ORDINAL`,
	"$ZERO":    `"0"`,
}
