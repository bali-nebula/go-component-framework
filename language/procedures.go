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
	"github.com/craterdog-bali/go-bali-document-notation/elements"
)

// This type defines the node structure associated with a clause that accepts a
// message that was previously retrieved from a named message bag so that it
// cannot be retrieved by another party.
type AcceptClause struct {
	Message any
}

// This type defines the node structure associated with an indexed attribute
// within a composite component.
type Attribute struct {
	Variable string
	Indices  []any
}

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

// This type defines the node structure associated with a do block of statements
// that contains an expression and Bali Document Notation (BDN) procedural
// statements.
type DoBlock struct {
	Expression any
	Statements []any
}

// This method attempts to parse a do block. It returns the do block and whether
// or not the do block was successfully parsed.
func (v *parser) parseDoBlock() (*DoBlock, bool) {
	var ok bool
	var bad *token
	var expression any
	var statements []any
	var doBlock *DoBlock
	expression, ok = v.parseExpression()
	if !ok {
		// We know a do block is expected so time to panic...
		panic("Expected an expression before a statement block.")
	}
	bad, ok = v.parseDelimiter("{")
	if !ok {
		panic(fmt.Sprintf("Expected a '{' character following the expression and received: %v", *bad))
	}
	statements, ok = v.parseStatements()
	if !ok {
		panic("Expected statements following the '{' character.")
	}
	bad, ok = v.parseDelimiter("}")
	if !ok {
		panic(fmt.Sprintf("Expected a '}' character following the statements and received: %v", *bad))
	}
	doBlock = &DoBlock{expression, statements}
	return doBlock, true
}

// This type defines the node structure associated with a clause that causes the
// execution of a loop to end.
type BreakClause struct {
}

// This method attempts to parse a break clause. It returns the break
// clause and whether or not the break clause was successfully parsed.
func (v *parser) parseBreakClause() (*BreakClause, bool) {
	var ok bool
	var clause *BreakClause
	_, ok = v.parseKeyword("break")
	if !ok {
		// This is not a break clause.
		return clause, false
	}
	_, ok = v.parseKeyword("loop")
	if !ok {
		panic("Expected the 'loop' keyword following the 'break' keyword.")
	}
	return clause, true
}

// This type defines the node structure associated with a clause that checks out
// a draft version of a released document at an optional release level from the
// document repository and assigns it to a recipient.
type CheckoutClause struct {
	Recipient any // The recipient is a symbol or attribute.
	Level     any
	Moniker   any
}

// This type defines the node structure associated with a clause that causes the
// execution of a loop to continue at the beginning.
type ContinueClause struct {
}

// This method attempts to parse a continue clause. It returns the continue
// clause and whether or not the continue clause was successfully parsed.
func (v *parser) parseContinueClause() (*ContinueClause, bool) {
	var ok bool
	var clause *ContinueClause
	_, ok = v.parseKeyword("continue")
	if !ok {
		// This is not a continue clause.
		return clause, false
	}
	_, ok = v.parseKeyword("loop")
	if !ok {
		panic("Expected the 'loop' keyword following the 'continue' keyword.")
	}
	return clause, true
}

// This type defines the node structure associated with a clause that discards
// a draft document referred to by an expression from the document repository.
type DiscardClause struct {
	Citation any
}

// This method attempts to parse a discard clause. It returns the discard
// clause and whether or not the discard clause was successfully parsed.
func (v *parser) parseDiscardClause() (*DiscardClause, bool) {
	var ok bool
	var citation any
	var clause *DiscardClause
	_, ok = v.parseKeyword("discard")
	if !ok {
		// This is not a discard clause.
		return clause, false
	}
	citation, ok = v.parseExpression()
	if !ok {
		panic("Expected a citation expression following the 'discard' keyword.")
	}
	clause = &DiscardClause{citation}
	return clause, true
}

// This type defines the node structure associated with a clause that evaluates
// an expression and optionally assigns the result to a recipient. The recipient
// must support the `Scalable` interface.
type EvaluateClause struct {
	Recipient  any // The recipient is a symbol or attribute.
	Operator   string
	Expression any
}

// This method attempts to parse an evaluate clause. It returns the evaluate
// clause and whether or not the evaluate clause was successfully parsed.
func (v *parser) parseEvaluateClause() (*EvaluateClause, bool) {
	var ok bool
	var t *token
	var recipient any
	var operator string
	var expression any
	var clause *EvaluateClause
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
		// This is not an evaluate clause.
		return clause, false
	}
	clause = &EvaluateClause{recipient, operator, expression}
	return clause, true
}

// This type defines the node structure associated with a clause that handles
// an exception.
type ExceptionClause struct {
	Exception elements.Symbol
	DoBlocks  []*DoBlock
}

// This method attempts to parse an exception clause. It returns the exception
// clause and whether or not the exception clause was successfully parsed.
func (v *parser) parseExceptionClause() (*ExceptionClause, bool) {
	var ok bool
	var clause *ExceptionClause
	return clause, ok
}

// This type defines the node structure associated with a clause that selects a
// statement block to be executed based on the value of a conditional expression.
type IfClause struct {
	DoBlock *DoBlock
}

// This method attempts to parse an if clause. It returns the if clause and
// whether or not the if clause was successfully parsed.
func (v *parser) parseIfClause() (*IfClause, bool) {
	var ok bool
	var doBlock *DoBlock
	var clause *IfClause
	_, ok = v.parseKeyword("if")
	if !ok {
		// This is not an if clause.
		return clause, false
	}
	doBlock, ok = v.parseDoBlock()
	if !ok {
		panic("Expected a condition expression following the 'if' keyword.")
	}
	clause = &IfClause{doBlock}
	return clause, true
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
	mainClause, ok = v.parseIfClause()
	if !ok {
		mainClause, ok = v.parseSelectClause()
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
	/*
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

// This type defines the node structure associated with a clause that notarizes
// a draft document as a named released document in the document repository.
type NotarizeClause struct {
	Draft   any
	Moniker any
}

// This type defines the node structure associated with a clause that posts a
// message to a named message bag.
type PostClause struct {
	Message any
	Bag     any
}

// This type defines the node structure associated with a procedure that
// contains Bali Document Notation™ (BDN) procedural statements.
type Procedure struct {
	Statements []any // This includes statements and annotations.
}

// This method attempts to parse a procedure. It returns the procedure and
// whether or not the procedure was successfully parsed.
func (v *parser) parseProcedure() ([]any, bool) {
	var ok bool
	var bad *token
	var statements []any
	_, ok = v.parseDelimiter("{")
	if !ok {
		// This is not a procedure.
		return statements, false
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

// This type defines the node structure associated with a clause that publishes
// an event to be delivered to all interested parties.
type PublishClause struct {
	Event any
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

// This type defines the node structure associated with a clause that rejects a
// message that was previously retrieved from a named message bag so that it
// can be retrieved by another party.
type RejectClause struct {
	Message any
}

// This type defines the node structure associated with a clause that retrieves
// a random message from a named message bag and assigns it to a recipient.
type RetrieveClause struct {
	Recipient any // The recipient is a symbol or attribute.
	Bag       any
}

// This type defines the node structure associated with a clause that causes an
// executing procedure to return with a result.
type ReturnClause struct {
	Result any
}

// This method attempts to parse a return clause. It returns the return
// clause and whether or not the return clause was successfully parsed.
func (v *parser) parseReturnClause() (*ReturnClause, bool) {
	var ok bool
	var result any
	var clause *ReturnClause
	_, ok = v.parseKeyword("return")
	if !ok {
		// This is not a return clause.
		return clause, false
	}
	result, ok = v.parseExpression()
	if !ok {
		panic("Expected a result expression following the 'return' keyword.")
	}
	clause = &ReturnClause{result}
	return clause, true
}

// This type defines the node structure associated with a clause that saves
// a draft document referred to by an expression to the document repository
// and returns a citation to the document which is optionally assigned to a
// recipient.
type SaveClause struct {
	Draft    any
	Citation any
}

// This method attempts to parse a save clause. It returns the save
// clause and whether or not the save clause was successfully parsed.
func (v *parser) parseSaveClause() (*SaveClause, bool) {
	var ok bool
	var draft any
	var citation any
	var clause *SaveClause
	_, ok = v.parseKeyword("save")
	if !ok {
		// This is not a save clause.
		return clause, false
	}
	draft, ok = v.parseExpression()
	if !ok {
		panic("Expected a draft document expression following the 'save' keyword.")
	}
	_, ok = v.parseKeyword("as")
	if !ok {
		panic("Expected the 'as' keyword after the draft document expression.")
	}
	citation, ok = v.parseRecipient()
	if !ok {
		panic("Expected a citation recipient expression following the 'as' keyword.")
	}
	clause = &SaveClause{draft, citation}
	return clause, true
}

// This type defines the node structure associated with a clause that selects a
// statement block to be executed based on the pattern of a control expression.
type SelectClause struct {
	Control  any
	DoBlocks []*DoBlock
}

// This method attempts to parse an select clause. It returns the select clause
// and whether or not the select clause was successfully parsed.
func (v *parser) parseSelectClause() (*SelectClause, bool) {
	var ok bool
	var control any
	var doBlock *DoBlock
	var doBlocks []*DoBlock
	var clause *SelectClause
	_, ok = v.parseKeyword("select")
	if !ok {
		// This is not a select clause.
		return clause, false
	}
	control, ok = v.parseExpression()
	if !ok {
		panic("Expected a control expression following the 'select' keyword.")
	}
	for {
		_, ok = v.parseKeyword("matching")
		if !ok {
			break // No more possible matches.
		}
		doBlock, ok = v.parseDoBlock()
		if !ok {
			panic("Expected a block expression following the 'matching' keyword.")
		}
		doBlocks = append(doBlocks, doBlock)
	}
	// There must be at least one block expression.
	if len(doBlocks) == 0 {
		panic("Expected at least one block expression in the select clause.")
	}
	clause = &SelectClause{control, doBlocks}
	return clause, true
}

// This type defines the node structure associated with a Bali Document
// Notation (BDN) statement containing a main clause and an optional
// exception handling clause.
type Statement struct {
	MainClause   any
	ExceptionClause *ExceptionClause
}

// This method attempts to parse a statement. It returns the statement and
// whether or not the statement was successfully parsed.
func (v *parser) parseStatement() (*Statement, bool) {
	var ok bool
	var statement *Statement
	var mainClause any
	var exceptionClause *ExceptionClause
	mainClause, ok = v.parseMainClause()
	if ok {
		// The exception clause is optional.
		exceptionClause, _ = v.parseExceptionClause()
	}
	statement = &Statement{mainClause, exceptionClause}
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

// This type defines the node structure associated with a clause that causes an
// exception to be thrown in the executing procedure.
type ThrowClause struct {
	Exception any
}

// This method attempts to parse a throw clause. It returns the throw
// clause and whether or not the throw clause was successfully parsed.
func (v *parser) parseThrowClause() (*ThrowClause, bool) {
	var ok bool
	var exception any
	var clause *ThrowClause
	_, ok = v.parseKeyword("throw")
	if !ok {
		// This is not a throw clause.
		return clause, false
	}
	exception, ok = v.parseExpression()
	if !ok {
		panic("Expected an exception expression following the 'throw' keyword.")
	}
	clause = &ThrowClause{exception}
	return clause, true
}

// This type defines the node structure associated with a clause that executes
// a statement block while a condition is true.
type WhileClause struct {
	DoBlock *DoBlock
}

// This method attempts to parse a while clause. It returns the while clause and
// whether or not the while clause was successfully parsed.
func (v *parser) parseWhileClause() (*WhileClause, bool) {
	var ok bool
	var doBlock *DoBlock
	var clause *WhileClause
	_, ok = v.parseKeyword("while")
	if !ok {
		// This is not a while clause.
		return clause, false
	}
	doBlock, ok = v.parseDoBlock()
	if !ok {
		panic("Expected a conditional expression following the 'while' keyword.")
	}
	clause = &WhileClause{doBlock}
	return clause, true
}

// This type defines the node structure associated with a clause that executes
// a statement block for each item in a collection. Each item may be optionally
// assigned to a symbol that is referenced in the statement block.
type WithClause struct {
	Item    elements.Symbol
	DoBlock *DoBlock
}

// This method attempts to parse a with clause. It returns the with clause and
// whether or not the with clause was successfully parsed.
func (v *parser) parseWithClause() (*WithClause, bool) {
	var ok bool
	var item elements.Symbol
	var doBlock *DoBlock
	var clause *WithClause
	_, ok = v.parseKeyword("with")
	if !ok {
		// This is not a with clause.
		return clause, false
	}
	_, ok = v.parseKeyword("each")
	if ok {
		// There is an each item option.
		item, ok = v.parseSymbol()
		if !ok {
			panic("Expected a symbol following the 'each' keyword.")
		}
	}
	doBlock, ok = v.parseDoBlock()
	if !ok {
		panic("Expected a sequential expression following the item symbol.")
	}
	clause = &WithClause{item, doBlock}
	return clause, true
}
