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

// This method attempts to parse an attribute. It returns the attribute and
// whether or not the attribute was successfully parsed.
func (v *parser) parseAttribute() (*Attribute, bool) {
	var ok bool
	var bad *token
	var variable string
	var indices []any
	var attribute *Attribute
	variable, ok = v.parseIdentifier()
	if !ok {
		// This is not an attribute.
		return attribute, false
	}
	_, ok = v.parseDelimiter("[")
	if !ok {
		// This is not an attribute.
		v.backupOne() // Put back the variable token.
		return attribute, false
	}
	indices, ok = v.parseIndices()
	if !ok {
		panic("Expected indices following the '[' character.")
	}
	bad, ok = v.parseDelimiter("]")
	if !ok {
		panic(fmt.Sprintf("Expected a ']' character following the indices and received: %v", *bad))
	}
	attribute = &Attribute{variable, indices}
	return attribute, true
}

// This method attempts to parse an evaluate clause. It returns the evaluate
// clause and whether or not the evaluate clause was successfully parsed.
func (v *parser) parseEvaluateClause() (*EvaluateClause, bool) {
	var ok bool
	var t *token
	var recipient any
	var operator string
	var expression any
	var evaluateClause *EvaluateClause
	recipient, ok = v.parseRecipient()
	if ok {
		t, ok = v.parseDelimiter(":=")
		if !ok {
			t, ok = v.parseDelimiter("+=")
		}
		if !ok {
			t, ok = v.parseDelimiter("-=")
		}
		if !ok {
			t, ok = v.parseDelimiter("*=")
		}
		if !ok {
			t, ok = v.parseDelimiter("/=")
		}
		if !ok {
			panic(fmt.Sprintf("Expected an assignment operator and received: %v", *t))
		}
		operator = t.val
	}
	expression, ok = v.parseExpression()
	if !ok {
		return evaluateClause, false
	}
	evaluateClause = &EvaluateClause{recipient, operator, expression}
	return evaluateClause, true
}

// This method attempts to parse a handle clause. It returns the handle
// clause and whether or not the handle clause was successfully parsed.
func (v *parser) parseHandleClause() (*HandleClause, bool) {
	var ok bool
	var clause *HandleClause
	return clause, ok
}

// This method attempts to parse a sequence of indices. It returns an array of
// the indices and whether or not the indices were successfully parsed.
func (v *parser) parseIndices() ([]any, bool) {
	var ok bool
	var index any
	var indices []any
	index, ok = v.parseExpression()
	// There must be at least one index.
	if !ok {
		panic("Expected at least one index in the sequence of indices.")
	}
	for {
		indices = append(indices, index)
		// Every subsequent index must be preceded by a ','.
		_, ok = v.parseDelimiter(",")
		if !ok {
			// No more indices.
			break
		}
		index, ok = v.parseExpression()
		if !ok {
			panic("Expected an index after the ',' character.")
		}
	}
	return indices, true
}

// This method attempts to parse a main clause. It returns the main clause and
// whether or not the main clause was successfully parsed.
func (v *parser) parseMainClause() (any, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var mainClause any
	/*
	mainClause, ok = v.parseOnClause()
	if !ok {
		mainClause, ok = v.parseIfClause()
	}
	if !ok {
		mainClause, ok = v.parseWithClause()
	}
	if !ok {
		mainClause, ok = v.parseWhileClause()
	}
	if !ok {
		mainClause, ok = v.parseContinueClause()
	}
	if !ok {
		mainClause, ok = v.parseBreakClause()
	}
	if !ok {
		mainClause, ok = v.parseReturnClause()
	}
	if !ok {
		mainClause, ok = v.parseThrowClause()
	}
	if !ok {
		mainClause, ok = v.parseSaveClause()
	}
	if !ok {
		mainClause, ok = v.parseDiscardClause()
	}
	if !ok {
		mainClause, ok = v.parseNotarizeClause()
	}
	if !ok {
		mainClause, ok = v.parseCheckoutClause()
	}
	if !ok {
		mainClause, ok = v.parsePublishClause()
	}
	if !ok {
		mainClause, ok = v.parsePostClause()
	}
	if !ok {
		mainClause, ok = v.parseRetrieveClause()
	}
	if !ok {
		mainClause, ok = v.parseAcceptClause()
	}
	if !ok {
		mainClause, ok = v.parseRejectClause()
	}
	*/
	if !ok {
		// This clause should be check last as it is slower to fail.
		mainClause, ok = v.parseEvaluateClause()
	}
	return mainClause, ok
}

// This method attempts to parse a procedure. It returns the procedure and
// whether or not the procedure was successfully parsed.
func (v *parser) parseProcedure() ([]any, bool) {
	var ok bool
	var bad *token
	var statements []any
	_, ok = v.parseDelimiter("{")
	if !ok {
		return nil, false
	}
	statements, ok = v.parseStatements()
	if !ok {
		panic("Expected statements following the '{' character.")
	}
	bad, ok = v.parseDelimiter("}")
	if !ok {
		panic(fmt.Sprintf("Expected a '}' character following the statements and received: %v", *bad))
	}
	return statements, true
}

// This method attempts to parse a recipient. It returns the recipient and
// whether or not the recipient was successfully parsed.
func (v *parser) parseRecipient() (any, bool) {
	var ok bool
	var recipient any
	recipient, ok = v.parseSymbol()
	if !ok {
		recipient, ok = v.parseAttribute()
	}
	return recipient, ok
}

// This method attempts to parse a statement. It returns the statement and
// whether or not the statement was successfully parsed.
func (v *parser) parseStatement() (*Statement, bool) {
	var ok bool
	var statement *Statement
	var mainClause any
	var handleClause *HandleClause
	mainClause, ok = v.parseMainClause()
	if ok {
		handleClause, _ = v.parseHandleClause() // The handle clause is optional.
		statement = &Statement{mainClause, handleClause}
	}
	return statement, ok
}

// This method attempts to parse the statements within a procedure. It returns
// an array of the statements and whether or not the statements were
// successfully parsed.
func (v *parser) parseStatements() ([]any, bool) {
	var statement any
	var statements []any
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
		statement, ok = v.parseDocumentation()
		if !ok {
			statement, ok = v.parseStatement()
		}
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
