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
		var message = fmt.Sprintf("Expected a message expression following the 'accept' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$acceptClause")
		panic(message)
	}
	clause = &AcceptClause{message}
	return clause, token, true
}

// This type defines the node structure associated with an indexed attribute
// within a composite component.
type Attribute struct {
	Variable string
	Indices  abstractions.ListLike[any]
}

// This method attempts to parse an attribute. It returns the attribute and
// whether or not the attribute was successfully parsed.
func (v *parser) parseAttribute() (*Attribute, *Token, bool) {
	var ok bool
	var token *Token
	var variable string
	var indices abstractions.ListLike[any]
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
		var message = fmt.Sprintf("Expected indices following the '[' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("]")
	if !ok {
		var message = fmt.Sprintf("Expected a ']' character following the indices but received:\n%v\n\n", token)
		message += generateGrammar(
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	attribute = &Attribute{variable, indices}
	return attribute, token, true
}

// This type defines the node structure associated with a do block of statements
// that contains an expression and Bali Document Notation (BDN) procedural
// statements.
type DoBlock struct {
	Expression any
	Statements abstractions.ListLike[any]
}

// This method attempts to parse a do block. It returns the do block and whether
// or not the do block was successfully parsed.
func (v *parser) parseDoBlock() (*DoBlock, *Token, bool) {
	var ok bool
	var token *Token
	var expression any
	var statements abstractions.ListLike[any]
	var doBlock *DoBlock
	expression, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression before a statement block but received:\n%v\n\n", token)
		message += generateGrammar(
			"$ifClause",
			"$selectClause",
			"$withClause",
			"$whileClause",
			"$exceptionClause")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		var message = fmt.Sprintf("Expected a '{' character following the expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$ifClause",
			"$selectClause",
			"$withClause",
			"$whileClause",
			"$exceptionClause")
		panic(message)
	}
	statements, token, ok = v.parseStatements()
	if !ok {
		var message = fmt.Sprintf("Expected statements following the '{' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$ifClause",
			"$selectClause",
			"$withClause",
			"$whileClause",
			"$exceptionClause")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		var message = fmt.Sprintf("Expected a '}' character following the statements but received:\n%v\n\n", token)
		message += generateGrammar(
			"$ifClause",
			"$selectClause",
			"$withClause",
			"$whileClause",
			"$exceptionClause")
		panic(message)
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
		var message = fmt.Sprintf("Expected the 'loop' keyword following the 'break' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$breakClause")
		panic(message)
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
		var message = fmt.Sprintf("Expected a recipient following the 'checkout' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$checkoutClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	_, token, ok = v.parseKeyword("at")
	if ok {
		// There is an at level part to this clause.
		_, token, ok = v.parseKeyword("level")
		if !ok {
			var message = fmt.Sprintf("Expected a 'level' keyword following the 'at' keyword but received:\n%v\n\n", token)
			message += generateGrammar(
				"$checkoutClause",
				"$recipient",
				"$name",
				"$attribute",
				"$variable",
				"$indices")
			panic(message)
		}
		level, token, ok = v.parseExpression()
		if !ok {
			var message = fmt.Sprintf("Expected a level expression following the 'level' keyword but received:\n%v\n\n", token)
			message += generateGrammar(
				"$checkoutClause",
				"$recipient",
				"$name",
				"$attribute",
				"$variable",
				"$indices")
			panic(message)
		}
	}
	_, token, ok = v.parseKeyword("from")
	if !ok {
		var message = fmt.Sprintf("Expected the 'from' keyword after the recipient and level expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$checkoutClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	moniker, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a moniker expression following the 'from' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$checkoutClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
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
		var message = fmt.Sprintf("Expected the 'loop' keyword following the 'continue' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$continueClause")
		panic(message)
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
		var message = fmt.Sprintf("Expected a citation expression following the 'discard' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$discardClause")
		panic(message)
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
			var message = fmt.Sprintf("Expected an assignment operator but received:\n%v\n\n", token)
			message += generateGrammar(
				"$evaluateClause",
				"$recipient",
				"$name",
				"$attribute",
				"$variable",
				"$indices")
			panic(message)
		}
	}
	expression, token, ok = v.parseExpression()
	if !ok {
		if token != nil {
			var message = fmt.Sprintf("Expected an expression after the assignment operator '"+operator+"' but received:\n%v\n\n", token)
			message += generateGrammar(
				"$evaluateClause",
				"$recipient",
				"$name",
				"$attribute",
				"$variable",
				"$indices")
			panic(message)
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
	DoBlocks  abstractions.ListLike[*DoBlock]
}

// This method attempts to parse an exception clause. It returns the exception
// clause and whether or not the exception clause was successfully parsed.
func (v *parser) parseExceptionClause() (*ExceptionClause, *Token, bool) {
	var ok bool
	var token *Token
	var exception elements.Symbol
	var doBlock *DoBlock
	var doBlocks abstractions.ListLike[*DoBlock]
	var clause *ExceptionClause
	_, token, ok = v.parseKeyword("on")
	if !ok {
		// This is not an exception clause.
		return clause, token, false
	}
	exception, token, ok = v.parseSymbol()
	if !ok {
		var message = fmt.Sprintf("Expected an exception symbol following the 'on' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$exceptionClause",
			"$exception")
		panic(message)
	}
	for {
		_, token, ok = v.parseKeyword("matching")
		if !ok {
			break // No more possible matches.
		}
		doBlock, token, ok = v.parseDoBlock()
		if !ok {
			var message = fmt.Sprintf("Expected a pattern expression and statement block following the 'matching' keyword but received:\n%v\n\n", token)
			message += generateGrammar(
				"$exceptionClause",
				"$exception")
			panic(message)
		}
		doBlocks.AddItem(doBlock)
	}
	// There must be at least one matching block expression.
	if doBlocks.IsEmpty() {
		var message = fmt.Sprintf("Expected at least one pattern expression and statement block in the exception clause but received:\n%v\n\n", token)
		message += generateGrammar(
			"$exceptionClause",
			"$exception")
		panic(message)
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
		var message = fmt.Sprintf("Expected a condition expression and statement block following the 'if' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$ifClause")
		panic(message)
	}
	clause = &IfClause{doBlock}
	return clause, token, true
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
		var message = fmt.Sprintf("Expected a draft document expression following the 'notarize' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$notarizeClause")
		panic(message)
	}
	_, token, ok = v.parseKeyword("as")
	if !ok {
		var message = fmt.Sprintf("Expected the 'as' keyword after the draft document expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$notarizeClause")
		panic(message)
	}
	moniker, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a moniker expression following the 'as' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$notarizeClause")
		panic(message)
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
		var message = fmt.Sprintf("Expected a message expression following the 'post' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$postClause")
		panic(message)
	}
	_, token, ok = v.parseKeyword("to")
	if !ok {
		var message = fmt.Sprintf("Expected the 'to' keyword after the message expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$postClause")
		panic(message)
	}
	bag, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a bag expression following the 'to' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$postClause")
		panic(message)
	}
	clause = &PostClause{message, bag}
	return clause, token, true
}

// This method attempts to parse a procedure. It returns the procedure and
// whether or not the procedure was successfully parsed.
func (v *parser) parseProcedure() (abstractions.ListLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var statements abstractions.ListLike[any]
	_, token, ok = v.parseDelimiter("{")
	if !ok {
		// This is not a procedure.
		return statements, token, false
	}
	statements, token, ok = v.parseStatements()
	if !ok {
		var message = fmt.Sprintf("Expected statements following the '{' character but received:\n%v\n\n", token)
		message += generateGrammar(
			"$procedure",
			"$statements",
			"$statement")
		panic(message)
	}
	_, token, ok = v.parseDelimiter("}")
	if !ok {
		var message = fmt.Sprintf("Expected a '}' character following the statements but received:\n%v\n\n", token)
		message += generateGrammar(
			"$procedure",
			"$statements",
			"$statement")
		panic(message)
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
		var message = fmt.Sprintf("Expected an event expression following the 'publish' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$publishClause")
		panic(message)
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
		var message = fmt.Sprintf("Expected a message expression following the 'reject' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$rejectClause")
		panic(message)
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
		var message = fmt.Sprintf("Expected a message recipient following the 'retrieve' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$retrieveClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	_, token, ok = v.parseKeyword("from")
	if !ok {
		var message = fmt.Sprintf("Expected the 'from' keyword after the message recipient but received:\n%v\n\n", token)
		message += generateGrammar(
			"$retrieveClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	bag, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a bag expression following the 'from' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$retrieveClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
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
		var message = fmt.Sprintf("Expected a result expression following the 'return' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$returnClause")
		panic(message)
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
		var message = fmt.Sprintf("Expected a draft document expression following the 'save' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$saveClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	_, token, ok = v.parseKeyword("as")
	if !ok {
		var message = fmt.Sprintf("Expected the 'as' keyword after the draft document expression but received:\n%v\n\n", token)
		message += generateGrammar(
			"$saveClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	citation, token, ok = v.parseRecipient()
	if !ok {
		var message = fmt.Sprintf("Expected a citation recipient following the 'as' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$saveClause",
			"$recipient",
			"$name",
			"$attribute",
			"$variable",
			"$indices")
		panic(message)
	}
	clause = &SaveClause{draft, citation}
	return clause, token, true
}

// This type defines the node structure associated with a clause that selects a
// statement block to be executed based on the pattern of a control expression.
type SelectClause struct {
	Control  any
	DoBlocks abstractions.ListLike[*DoBlock]
}

// This method attempts to parse an select clause. It returns the select clause
// and whether or not the select clause was successfully parsed.
func (v *parser) parseSelectClause() (*SelectClause, *Token, bool) {
	var ok bool
	var token *Token
	var control any
	var doBlock *DoBlock
	var doBlocks abstractions.ListLike[*DoBlock]
	var clause *SelectClause
	_, token, ok = v.parseKeyword("select")
	if !ok {
		// This is not a select clause.
		return clause, token, false
	}
	control, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected a control expression following the 'select' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$selectClause")
		panic(message)
	}
	for {
		_, token, ok = v.parseKeyword("matching")
		if !ok {
			break // No more possible matches.
		}
		doBlock, token, ok = v.parseDoBlock()
		if !ok {
			var message = fmt.Sprintf("Expected a pattern expression and statement block following the 'matching' keyword but received:\n%v\n\n", token)
			message += generateGrammar(
				"$selectClause")
			panic(message)
		}
		doBlocks.AddItem(doBlock)
	}
	// There must be at least one matching block expression.
	if doBlocks.IsEmpty() {
		var message = fmt.Sprintf("Expected at least one pattern expression and statement block in the select clause but received:\n%v\n\n", token)
		message += generateGrammar(
			"$selectClause")
		panic(message)
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
func (v *parser) parseStatements() (abstractions.ListLike[any], *Token, bool) {
	var statement any
	var statements abstractions.ListLike[any]
	var _, token, ok = v.parseEOL()
	if !ok {
		// The statements are on a single line.
		statement, token, ok = v.parseStatement()
		// There must be at least one statement.
		if !ok {
			var message = fmt.Sprintf("Expected at least one statement in the component context but received:\n%v\n\n", token)
			message += generateGrammar(
				"$procedure",
				"$statements",
				"$statement")
			panic(message)
		}
		for {
			statements.AddItem(statement)
			// Every subsequent statement must be preceded by a ';'.
			_, token, ok = v.parseDelimiter(";")
			if !ok {
				// No more statements.
				break
			}
			statement, token, ok = v.parseStatement()
			if !ok {
				var message = fmt.Sprintf("Expected a statement after the ';' character but received:\n%v\n\n", token)
				message += generateGrammar(
					"$procedure",
					"$statements",
					"$statement")
				panic(message)
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
		statements.AddItem(statement)
		// Every statement must be followed by an EOL.
		_, token, ok = v.parseEOL()
		if !ok {
			var message = fmt.Sprintf("Expected an EOL character following the statement but received:\n%v\n\n", token)
			message += generateGrammar(
				"$procedure",
				"$statements",
				"$statement")
			panic(message)
		}
	}
	// There must be at least one statement.
	if statements.IsEmpty() {
		var message = fmt.Sprintf("Expected at least one statement in the component context but received:\n%v\n\n", token)
		message += generateGrammar(
			"$procedure",
			"$statements",
			"$statement")
		panic(message)
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
		var message = fmt.Sprintf("Expected an exception expression following the 'throw' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$throwClause")
		panic(message)
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
		var message = fmt.Sprintf("Expected a conditional expression and statement block following the 'while' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$whileClause")
		panic(message)
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
		var message = fmt.Sprintf("Expected an 'each' keyword following the 'with' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$withClause",
			"$item")
		panic(message)
	}
	item, token, ok = v.parseSymbol()
	if !ok {
		var message = fmt.Sprintf("Expected a symbol following the 'each' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$withClause",
			"$item")
		panic(message)
	}
	_, token, ok = v.parseKeyword("in")
	if !ok {
		var message = fmt.Sprintf("Expected an 'in' keyword following the symbol but received:\n%v\n\n", token)
		message += generateGrammar(
			"$withClause",
			"$item")
		panic(message)
	}
	doBlock, token, ok = v.parseDoBlock()
	if !ok {
		var message = fmt.Sprintf("Expected a sequential expression and statement block following the 'in' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$withClause",
			"$item")
		panic(message)
	}
	clause = &WithClause{item, doBlock}
	return clause, token, true
}
