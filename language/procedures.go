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

// This method attempts to parse an accept clause. It returns the accept
// clause and whether or not the accept clause was successfully parsed.
func (v *parser) parseAcceptClause() (*AcceptClause, *Token, bool) {
	var ok bool
	var token *Token
	var message any
	var clause *AcceptClause
	_, token, ok = v.parseKeyword("accept")
	if !ok {
		// This is not a accept clause.
		return clause, token, false
	}
	message, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a message expression following the 'accept' keyword.")
	}
	clause = &AcceptClause{message}
	return clause, token, true
}

// This type defines the node structure associated with an indexed attribute
// within a composite component.
type Attribute struct {
	Variable string
	Indices  []any
}

// This method attempts to parse an attribute. It returns the attribute and
// whether or not the attribute was successfully parsed.
func (v *parser) parseAttribute() (*Attribute, *Token, bool) {
	var ok bool
	var token *Token
	var variable string
	var indices []any
	var attribute *Attribute
	variable, token, ok = v.parseIdentifier()
	if !ok {
		// This is not an attribute.
		return attribute, token, false
	}
	_, token, ok = v.parseDelimiter("[")
	if !ok {
		// This is not an attribute.
		v.backupOne() // Put back the variable token.
		return attribute, token, false
	}
	indices, token, ok = v.parseIndices()
	if !ok {
		panic("Expected indices following the '[' character.")
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		panic(fmt.Sprintf("Expected a ']' character following the indices and received: %v", *token))
	}
	attribute = &Attribute{variable, indices}
	return attribute, token, true
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
func (v *parser) parseDoBlock() (*DoBlock, *Token, bool) {
	var ok bool
	var token *Token
	var expression any
	var statements []any
	var doBlock *DoBlock
	expression, token, ok = v.parseExpression()
	if !ok {
		// We know a do block is expected so time to panic...
		panic("Expected an expression before a statement block.")
	}
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		panic(fmt.Sprintf("Expected a '{' character following the expression and received: %v", *token))
	}
	statements, token, ok = v.parseStatements()
	if !ok {
		panic("Expected statements following the '{' character.")
	}
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		panic(fmt.Sprintf("Expected a '}' character following the statements and received: %v", *token))
	}
	doBlock = &DoBlock{expression, statements}
	return doBlock, token, true
}

// This type defines the node structure associated with a clause that causes the
// execution of a loop to end.
type BreakClause struct {
}

// This method attempts to parse a break clause. It returns the break
// clause and whether or not the break clause was successfully parsed.
func (v *parser) parseBreakClause() (*BreakClause, *Token, bool) {
	var ok bool
	var token *Token
	var clause *BreakClause
	_, token, ok = v.parseKeyword("break")
	if !ok {
		// This is not a break clause.
		return clause, token, false
	}
	_, token, ok = v.parseKeyword("loop")
	if !ok {
		panic("Expected the 'loop' keyword following the 'break' keyword.")
	}
	return clause, token, true
}

// This type defines the node structure associated with a clause that checks out
// a draft version of a released document at an optional release level from the
// document repository and assigns it to a recipient.
type CheckoutClause struct {
	Recipient any // The recipient is a symbol or attribute.
	Level     any
	Moniker   any
}

// This method attempts to parse a checkout clause. It returns the checkout
// clause and whether or not the checkout clause was successfully parsed.
func (v *parser) parseCheckoutClause() (*CheckoutClause, *Token, bool) {
	var ok bool
	var token *Token
	var recipient any
	var level any
	var moniker any
	var clause *CheckoutClause
	_, token, ok = v.parseKeyword("checkout")
	if !ok {
		// This is not a checkout clause.
		return clause, token, false
	}
	recipient, token, ok = v.parseRecipient()
	if !ok {
		panic("Expected a recipient following the 'checkout' keyword.")
	}
	_, token, ok = v.parseKeyword("at")
	if ok {
		// There is an at level part to this clause.
		_, token, ok = v.parseKeyword("level")
		if !ok {
			panic("Expected a 'level' keyword following the 'at' keyword.")
		}
		level, token, ok = v.parseExpression()
		if !ok {
			panic("Expected a level expression following the 'level' keyword.")
		}
	}
	_, token, ok = v.parseKeyword("from")
	if !ok {
		panic("Expected the 'from' keyword after the recipient and level expression.")
	}
	moniker, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a moniker expression following the 'from' keyword.")
	}
	clause = &CheckoutClause{recipient, level, moniker}
	return clause, token, true
}

// This type defines the node structure associated with a clause that causes the
// execution of a loop to continue at the beginning.
type ContinueClause struct {
}

// This method attempts to parse a continue clause. It returns the continue
// clause and whether or not the continue clause was successfully parsed.
func (v *parser) parseContinueClause() (*ContinueClause, *Token, bool) {
	var ok bool
	var token *Token
	var clause *ContinueClause
	_, token, ok = v.parseKeyword("continue")
	if !ok {
		// This is not a continue clause.
		return clause, token, false
	}
	_, token, ok = v.parseKeyword("loop")
	if !ok {
		panic("Expected the 'loop' keyword following the 'continue' keyword.")
	}
	return clause, token, true
}

// This type defines the node structure associated with a clause that discards
// a draft document referred to by an expression from the document repository.
type DiscardClause struct {
	Citation any
}

// This method attempts to parse a discard clause. It returns the discard
// clause and whether or not the discard clause was successfully parsed.
func (v *parser) parseDiscardClause() (*DiscardClause, *Token, bool) {
	var ok bool
	var token *Token
	var citation any
	var clause *DiscardClause
	_, token, ok = v.parseKeyword("discard")
	if !ok {
		// This is not a discard clause.
		return clause, token, false
	}
	citation, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a citation expression following the 'discard' keyword.")
	}
	clause = &DiscardClause{citation}
	return clause, token, true
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
func (v *parser) parseEvaluateClause() (*EvaluateClause, *Token, bool) {
	var ok bool
	var token *Token
	var recipient any
	var operator string
	var expression any
	var clause *EvaluateClause
	recipient, token, ok = v.parseRecipient()
	if ok {
		operator, token, ok = v.parseDelimiter(":=")
		if !ok {
			operator, token, ok = v.parseDelimiter("+=")
		}
		if !ok {
			operator, token, ok = v.parseDelimiter("-=")
		}
		if !ok {
			operator, token, ok = v.parseDelimiter("*=")
		}
		if !ok {
			operator, token, ok = v.parseDelimiter("/=")
		}
		if !ok {
			panic(fmt.Sprintf("Expected an assignment operator and received: %v", *token))
		}
	}
	expression, token, ok = v.parseExpression()
	if !ok {
		if token != nil {
			panic(fmt.Sprintf("Expected an expression after the assignment operator: %q.", operator))
		}
		// This is not an evaluate clause.
		return clause, token, false
	}
	clause = &EvaluateClause{recipient, operator, expression}
	return clause, token, true
}

// This type defines the node structure associated with a clause that handles
// an exception.
type ExceptionClause struct {
	Exception elements.Symbol
	DoBlocks  []*DoBlock
}

// This method attempts to parse an exception clause. It returns the exception
// clause and whether or not the exception clause was successfully parsed.
func (v *parser) parseExceptionClause() (*ExceptionClause, *Token, bool) {
	var ok bool
	var token *Token
	var exception elements.Symbol
	var doBlock *DoBlock
	var doBlocks []*DoBlock
	var clause *ExceptionClause
	_, token, ok = v.parseKeyword("on")
	if !ok {
		// This is not an exception clause.
		return clause, token, false
	}
	exception, token, ok = v.parseSymbol()
	if !ok {
		panic("Expected an exception symbol following the 'on' keyword.")
	}
	for {
		_, token, ok = v.parseKeyword("matching")
		if !ok {
			break // No more possible matches.
		}
		doBlock, token, ok = v.parseDoBlock()
		if !ok {
			panic("Expected a pattern expression and statement block following the 'matching' keyword.")
		}
		doBlocks = append(doBlocks, doBlock)
	}
	// There must be at least one matching block expression.
	if len(doBlocks) == 0 {
		panic("Expected at least one pattern expression and statement block in the exception clause.")
	}
	clause = &ExceptionClause{exception, doBlocks}
	return clause, token, true
}

// This type defines the node structure associated with a clause that selects a
// statement block to be executed based on the value of a conditional expression.
type IfClause struct {
	DoBlock *DoBlock
}

// This method attempts to parse an if clause. It returns the if clause and
// whether or not the if clause was successfully parsed.
func (v *parser) parseIfClause() (*IfClause, *Token, bool) {
	var ok bool
	var token *Token
	var doBlock *DoBlock
	var clause *IfClause
	_, token, ok = v.parseKeyword("if")
	if !ok {
		// This is not an if clause.
		return clause, token, false
	}
	doBlock, token, ok = v.parseDoBlock()
	if !ok {
		panic("Expected a condition expression and statement block following the 'if' keyword.")
	}
	clause = &IfClause{doBlock}
	return clause, token, true
}

// This method attempts to parse a sequence of indices. It returns an array of
// the indices and whether or not the indices were successfully parsed.
func (v *parser) parseIndices() ([]any, *Token, bool) {
	var ok bool
	var token *Token
	var index any
	var indices []any
	index, token, ok = v.parseExpression()
	// There must be at least one index.
	if !ok {
		panic("Expected at least one index in the sequence of indices.")
	}
	for {
		indices = append(indices, index)
		// Every subsequent index must be preceded by a ','.
		_, token, ok = v.parseDelimiter(",")
		if !ok {
			// No more indices.
			break
		}
		index, token, ok = v.parseExpression()
		if !ok {
			panic("Expected an index after the ',' character.")
		}
	}
	return indices, token, true
}

// This method attempts to parse a main clause. It returns the main clause and
// whether or not the main clause was successfully parsed.
func (v *parser) parseMainClause() (any, *Token, bool) {
	// TODO: Reorder these based on how often each type occurs.
	var ok bool
	var token *Token
	var mainClause any
	mainClause, token, ok = v.parseIfClause()
	if !ok {
		mainClause, token, ok = v.parseSelectClause()
	}
	if !ok {
		mainClause, token, ok = v.parseWithClause()
	}
	if !ok {
		mainClause, token, ok = v.parseWhileClause()
	}
	if !ok {
		mainClause, token, ok = v.parseContinueClause()
	}
	if !ok {
		mainClause, token, ok = v.parseBreakClause()
	}
	if !ok {
		mainClause, token, ok = v.parseReturnClause()
	}
	if !ok {
		mainClause, token, ok = v.parseThrowClause()
	}
	if !ok {
		mainClause, token, ok = v.parseSaveClause()
	}
	if !ok {
		mainClause, token, ok = v.parseDiscardClause()
	}
	if !ok {
		mainClause, token, ok = v.parseNotarizeClause()
	}
	if !ok {
		mainClause, token, ok = v.parseCheckoutClause()
	}
	if !ok {
		mainClause, token, ok = v.parsePublishClause()
	}
	if !ok {
		mainClause, token, ok = v.parsePostClause()
	}
	if !ok {
		mainClause, token, ok = v.parseRetrieveClause()
	}
	if !ok {
		mainClause, token, ok = v.parseAcceptClause()
	}
	if !ok {
		mainClause, token, ok = v.parseRejectClause()
	}
	if !ok {
		// This clause should be checked last since it is slower to fail.
		mainClause, token, ok = v.parseEvaluateClause()
	}
	return mainClause, token, ok
}

// This type defines the node structure associated with a clause that notarizes
// a draft document as a named released document in the document repository.
type NotarizeClause struct {
	Draft   any
	Moniker any
}

// This method attempts to parse a notarize clause. It returns the notarize
// clause and whether or not the notarize clause was successfully parsed.
func (v *parser) parseNotarizeClause() (*NotarizeClause, *Token, bool) {
	var ok bool
	var token *Token
	var draft any
	var moniker any
	var clause *NotarizeClause
	_, token, ok = v.parseKeyword("notarize")
	if !ok {
		// This is not a notarize clause.
		return clause, token, false
	}
	draft, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a draft document expression following the 'notarize' keyword.")
	}
	_, token, ok = v.parseKeyword("as")
	if !ok {
		panic("Expected the 'as' keyword after the draft document expression.")
	}
	moniker, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a moniker expression following the 'as' keyword.")
	}
	clause = &NotarizeClause{draft, moniker}
	return clause, token, true
}

// This type defines the node structure associated with a clause that posts a
// message to a named message bag.
type PostClause struct {
	Message any
	Bag     any
}

// This method attempts to parse a post clause. It returns the post
// clause and whether or not the post clause was successfully parsed.
func (v *parser) parsePostClause() (*PostClause, *Token, bool) {
	var ok bool
	var token *Token
	var message any
	var bag any
	var clause *PostClause
	_, token, ok = v.parseKeyword("post")
	if !ok {
		// This is not a post clause.
		return clause, token, false
	}
	message, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a message expression following the 'post' keyword.")
	}
	_, token, ok = v.parseKeyword("to")
	if !ok {
		panic("Expected the 'to' keyword after the message expression.")
	}
	bag, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a bag expression following the 'to' keyword.")
	}
	clause = &PostClause{message, bag}
	return clause, token, true
}

// This type defines the node structure associated with a procedure that
// contains Bali Document Notation™ (BDN) procedural statements.
type Procedure struct {
	Statements []any // This includes statements and annotations.
}

// This method attempts to parse a procedure. It returns the procedure and
// whether or not the procedure was successfully parsed.
func (v *parser) parseProcedure() ([]any, *Token, bool) {
	var ok bool
	var token *Token
	var statements []any
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		// This is not a procedure.
		return statements, token, false
	}
	statements, token, ok = v.parseStatements()
	if !ok {
		panic("Expected statements following the '{' character.")
	}
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		panic(fmt.Sprintf("Expected a '}' character following the statements and received: %v", *token))
	}
	return statements, token, true
}

// This type defines the node structure associated with a clause that publishes
// an event to be delivered to all interested parties.
type PublishClause struct {
	Event any
}

// This method attempts to parse a publish clause. It returns the publish
// clause and whether or not the publish clause was successfully parsed.
func (v *parser) parsePublishClause() (*PublishClause, *Token, bool) {
	var ok bool
	var token *Token
	var event any
	var clause *PublishClause
	_, token, ok = v.parseKeyword("publish")
	if !ok {
		// This is not a publish clause.
		return clause, token, false
	}
	event, token, ok = v.parseExpression()
	if !ok {
		panic("Expected an event expression following the 'publish' keyword.")
	}
	clause = &PublishClause{event}
	return clause, token, true
}

// This method attempts to parse a recipient. It returns the recipient and
// whether or not the recipient was successfully parsed.
func (v *parser) parseRecipient() (any, *Token, bool) {
	var ok bool
	var token *Token
	var recipient any
	recipient, token, ok = v.parseSymbol()
	if !ok {
		recipient, token, ok = v.parseAttribute()
	}
	return recipient, token, ok
}

// This type defines the node structure associated with a clause that rejects a
// message that was previously retrieved from a named message bag so that it
// can be retrieved by another party.
type RejectClause struct {
	Message any
}

// This method attempts to parse a reject clause. It returns the reject
// clause and whether or not the reject clause was successfully parsed.
func (v *parser) parseRejectClause() (*RejectClause, *Token, bool) {
	var ok bool
	var token *Token
	var message any
	var clause *RejectClause
	_, token, ok = v.parseKeyword("reject")
	if !ok {
		// This is not a reject clause.
		return clause, token, false
	}
	message, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a message expression following the 'reject' keyword.")
	}
	clause = &RejectClause{message}
	return clause, token, true
}

// This type defines the node structure associated with a clause that retrieves
// a random message from a named message bag and assigns it to a recipient.
type RetrieveClause struct {
	Recipient any // The recipient is a symbol or attribute.
	Bag       any
}

// This method attempts to parse a retrieve clause. It returns the retrieve
// clause and whether or not the retrieve clause was successfully parsed.
func (v *parser) parseRetrieveClause() (*RetrieveClause, *Token, bool) {
	var ok bool
	var token *Token
	var recipient any
	var bag any
	var clause *RetrieveClause
	_, token, ok = v.parseKeyword("retrieve")
	if !ok {
		// This is not a retrieve clause.
		return clause, token, false
	}
	recipient, token, ok = v.parseRecipient()
	if !ok {
		panic("Expected a message recipient following the 'retrieve' keyword.")
	}
	_, token, ok = v.parseKeyword("from")
	if !ok {
		panic("Expected the 'from' keyword after the message recipient.")
	}
	bag, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a bag expression following the 'from' keyword.")
	}
	clause = &RetrieveClause{recipient, bag}
	return clause, token, true
}

// This type defines the node structure associated with a clause that causes an
// executing procedure to return with a result.
type ReturnClause struct {
	Result any
}

// This method attempts to parse a return clause. It returns the return
// clause and whether or not the return clause was successfully parsed.
func (v *parser) parseReturnClause() (*ReturnClause, *Token, bool) {
	var ok bool
	var token *Token
	var result any
	var clause *ReturnClause
	_, token, ok = v.parseKeyword("return")
	if !ok {
		// This is not a return clause.
		return clause, token, false
	}
	result, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a result expression following the 'return' keyword.")
	}
	clause = &ReturnClause{result}
	return clause, token, true
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
func (v *parser) parseSaveClause() (*SaveClause, *Token, bool) {
	var ok bool
	var token *Token
	var draft any
	var citation any
	var clause *SaveClause
	_, token, ok = v.parseKeyword("save")
	if !ok {
		// This is not a save clause.
		return clause, token, false
	}
	draft, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a draft document expression following the 'save' keyword.")
	}
	_, token, ok = v.parseKeyword("as")
	if !ok {
		panic("Expected the 'as' keyword after the draft document expression.")
	}
	citation, token, ok = v.parseRecipient()
	if !ok {
		panic("Expected a citation recipient following the 'as' keyword.")
	}
	clause = &SaveClause{draft, citation}
	return clause, token, true
}

// This type defines the node structure associated with a clause that selects a
// statement block to be executed based on the pattern of a control expression.
type SelectClause struct {
	Control  any
	DoBlocks []*DoBlock
}

// This method attempts to parse an select clause. It returns the select clause
// and whether or not the select clause was successfully parsed.
func (v *parser) parseSelectClause() (*SelectClause, *Token, bool) {
	var ok bool
	var token *Token
	var control any
	var doBlock *DoBlock
	var doBlocks []*DoBlock
	var clause *SelectClause
	_, token, ok = v.parseKeyword("select")
	if !ok {
		// This is not a select clause.
		return clause, token, false
	}
	control, token, ok = v.parseExpression()
	if !ok {
		panic("Expected a control expression following the 'select' keyword.")
	}
	for {
		_, token, ok = v.parseKeyword("matching")
		if !ok {
			break // No more possible matches.
		}
		doBlock, token, ok = v.parseDoBlock()
		if !ok {
			panic("Expected a pattern expression and statement block following the 'matching' keyword.")
		}
		doBlocks = append(doBlocks, doBlock)
	}
	// There must be at least one matching block expression.
	if len(doBlocks) == 0 {
		panic("Expected at least one pattern expression and statement block in the select clause.")
	}
	clause = &SelectClause{control, doBlocks}
	return clause, token, true
}

// This type defines the node structure associated with a Bali Document
// Notation (BDN) statement containing a main clause and an optional
// exception handling clause.
type Statement struct {
	MainClause      any
	ExceptionClause *ExceptionClause
}

// This method attempts to parse a statement. It returns the statement and
// whether or not the statement was successfully parsed.
func (v *parser) parseStatement() (*Statement, *Token, bool) {
	var ok bool
	var token *Token
	var statement *Statement
	var mainClause any
	var exceptionClause *ExceptionClause
	mainClause, token, ok = v.parseMainClause()
	if ok {
		// The exception clause is optional.
		exceptionClause, token, _ = v.parseExceptionClause()
	}
	statement = &Statement{mainClause, exceptionClause}
	return statement, token, ok
}

// This method attempts to parse the statements within a procedure. It returns
// an array of the statements and whether or not the statements were
// successfully parsed.
func (v *parser) parseStatements() ([]any, *Token, bool) {
	var statement any
	var statements []any
	var _, token, ok = v.parseEOL()
	if !ok {
		// The statements are on a single line.
		statement, token, ok = v.parseStatement()
		// There must be at least one statement.
		if !ok {
			panic("Expected at least one statement in the component context.")
		}
		for {
			statements = append(statements, statement)
			// Every subsequent statement must be preceded by a ';'.
			_, token, ok = v.parseDelimiter(";")
			if !ok {
				// No more statements.
				break
			}
			statement, token, ok = v.parseStatement()
			if !ok {
				panic("Expected a statement after the ';' character.")
			}
		}
		return statements, token, true
	}
	// The statements are on separate lines.
	for {
		statement, token, ok = v.parseDocumentation()
		if !ok {
			statement, token, ok = v.parseStatement()
		}
		if !ok {
			// No more statements.
			break
		}
		statements = append(statements, statement)
		// Every statement must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			panic("Expected an EOL character following the statement.")
		}
	}
	// There must be at least one statement.
	if len(statements) == 0 {
		panic("Expected at least one statement in the component context.")
	}
	return statements, token, true
}

// This type defines the node structure associated with a clause that causes an
// exception to be thrown in the executing procedure.
type ThrowClause struct {
	Exception any
}

// This method attempts to parse a throw clause. It returns the throw
// clause and whether or not the throw clause was successfully parsed.
func (v *parser) parseThrowClause() (*ThrowClause, *Token, bool) {
	var ok bool
	var token *Token
	var exception any
	var clause *ThrowClause
	_, token, ok = v.parseKeyword("throw")
	if !ok {
		// This is not a throw clause.
		return clause, token, false
	}
	exception, token, ok = v.parseExpression()
	if !ok {
		panic("Expected an exception expression following the 'throw' keyword.")
	}
	clause = &ThrowClause{exception}
	return clause, token, true
}

// This type defines the node structure associated with a clause that executes
// a statement block while a condition is true.
type WhileClause struct {
	DoBlock *DoBlock
}

// This method attempts to parse a while clause. It returns the while clause and
// whether or not the while clause was successfully parsed.
func (v *parser) parseWhileClause() (*WhileClause, *Token, bool) {
	var ok bool
	var token *Token
	var doBlock *DoBlock
	var clause *WhileClause
	_, token, ok = v.parseKeyword("while")
	if !ok {
		// This is not a while clause.
		return clause, token, false
	}
	doBlock, token, ok = v.parseDoBlock()
	if !ok {
		panic("Expected a conditional expression and statement block following the 'while' keyword.")
	}
	clause = &WhileClause{doBlock}
	return clause, token, true
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
func (v *parser) parseWithClause() (*WithClause, *Token, bool) {
	var ok bool
	var token *Token
	var item elements.Symbol
	var doBlock *DoBlock
	var clause *WithClause
	_, token, ok = v.parseKeyword("with")
	if !ok {
		// This is not a with clause.
		return clause, token, false
	}
	_, token, ok = v.parseKeyword("each")
	if !ok {
		panic("Expected an 'each' keyword following the 'with' keyword.")
	}
	item, token, ok = v.parseSymbol()
	if !ok {
		panic("Expected a symbol following the 'each' keyword.")
	}
	_, token, ok = v.parseKeyword("in")
	if !ok {
		panic("Expected an 'in' keyword following the symbol.")
	}
	doBlock, token, ok = v.parseDoBlock()
	if !ok {
		panic("Expected a sequential expression and statement block following the 'in' keyword.")
	}
	clause = &WithClause{item, doBlock}
	return clause, token, true
}
