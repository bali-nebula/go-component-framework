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
	fmt "fmt"
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	sta "github.com/craterdog-bali/go-bali-document-notation/statements"
)

// This method attempts to parse an accept clause. It returns the accept
// clause and whether or not the accept clause was successfully parsed.
func (v *parser) parseAcceptClause() (abs.AcceptClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var message any
	var clause abs.AcceptClauseLike
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
	clause = sta.AcceptClause(message)
	return clause, token, true
}

// This method attempts to parse an attribute. It returns the attribute and
// whether or not the attribute was successfully parsed.
func (v *parser) parseAttribute() (abs.AttributeLike, *Token, bool) {
	var ok bool
	var token *Token
	var variable string
	var indices abs.ListLike[any]
	var attribute abs.AttributeLike
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
	attribute = sta.Attribute(variable, indices)
	return attribute, token, true
}

// This method attempts to parse a do block. It returns the do block and whether
// or not the do block was successfully parsed.
func (v *parser) parseBlock() (abs.BlockLike, *Token, bool) {
	var ok bool
	var token *Token
	var expression any
	var statements abs.ListLike[any]
	var block abs.BlockLike
	expression, token, ok = v.parseExpression()
	if !ok {
		var message = fmt.Sprintf("Expected an expression before a statement block but received:\n%v\n\n", token)
		message += generateGrammar(
			"$ifClause",
			"$selectClause",
			"$withClause",
			"$whileClause",
			"$onClause")
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
			"$onClause")
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
			"$onClause")
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
			"$onClause")
		panic(message)
	}
	block = sta.Block(expression, statements)
	return block, token, true
}

// This method attempts to parse a break clause. It returns the break
// clause and whether or not the break clause was successfully parsed.
func (v *parser) parseBreakClause() (abs.BreakClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var clause abs.BreakClauseLike
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

// This method attempts to parse a checkout clause. It returns the checkout
// clause and whether or not the checkout clause was successfully parsed.
func (v *parser) parseCheckoutClause() (abs.CheckoutClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var recipient any
	var level any
	var moniker any
	var clause abs.CheckoutClauseLike
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
	clause = sta.CheckoutClause(recipient, level, moniker)
	return clause, token, true
}

// This method attempts to parse a continue clause. It returns the continue
// clause and whether or not the continue clause was successfully parsed.
func (v *parser) parseContinueClause() (abs.ContinueClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var clause abs.ContinueClauseLike
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

// This method attempts to parse a discard clause. It returns the discard
// clause and whether or not the discard clause was successfully parsed.
func (v *parser) parseDiscardClause() (abs.DiscardClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var citation any
	var clause abs.DiscardClauseLike
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
	clause = sta.DiscardClause(citation)
	return clause, token, true
}

// This method attempts to parse an evaluate clause. It returns the evaluate
// clause and whether or not the evaluate clause was successfully parsed.
func (v *parser) parseEvaluateClause() (abs.EvaluateClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var recipient any
	var delimiter string
	var operator abs.Assignment
	var expression any
	var clause abs.EvaluateClauseLike
	recipient, token, ok = v.parseRecipient()
	if ok {
		delimiter, token, ok = v.parseDelimiter(":=")
		if !ok {
			delimiter, token, ok = v.parseDelimiter("+=")
		}
		if !ok {
			delimiter, token, ok = v.parseDelimiter("-=")
		}
		if !ok {
			delimiter, token, ok = v.parseDelimiter("*=")
		}
		if !ok {
			delimiter, token, ok = v.parseDelimiter("/=")
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
			var message = fmt.Sprintf("Expected an expression after the assignment operator '"+delimiter+"' but received:\n%v\n\n", token)
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
	switch delimiter {
	case ":=":
		operator = abs.REGULAR
	case "?=":
		operator = abs.DEFAULT
	case "*=":
		operator = abs.TIMES
	case "/=":
		operator = abs.DIVIDE
	case "+=":
		operator = abs.PLUS
	case "-=":
		operator = abs.MINUS
	}
	clause = sta.EvaluateClauseWithRecipient(recipient, operator, expression)
	return clause, token, true
}

// This method attempts to parse an if clause. It returns the if clause and
// whether or not the if clause was successfully parsed.
func (v *parser) parseIfClause() (abs.IfClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var block abs.BlockLike
	var clause abs.IfClauseLike
	_, token, ok = v.parseKeyword("if")
	if !ok {
		// This is not an if clause.
		return clause, token, false
	}
	block, token, ok = v.parseBlock()
	if !ok {
		var message = fmt.Sprintf("Expected a condition expression and statement block following the 'if' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$ifClause")
		panic(message)
	}
	clause = sta.IfClause(block.GetPattern(), block.GetStatements())
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

// This method attempts to parse a notarize clause. It returns the notarize
// clause and whether or not the notarize clause was successfully parsed.
func (v *parser) parseNotarizeClause() (abs.NotarizeClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var draft any
	var moniker any
	var clause abs.NotarizeClauseLike
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
	clause = sta.NotarizeClause(draft, moniker)
	return clause, token, true
}

// This method attempts to parse an exception clause. It returns the exception
// clause and whether or not the exception clause was successfully parsed.
func (v *parser) parseOnClause() (abs.OnClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var exception ele.Symbol
	var block abs.BlockLike
	var blocks abs.ListLike[abs.BlockLike]
	var clause abs.OnClauseLike
	_, token, ok = v.parseKeyword("on")
	if !ok {
		// This is not an exception clause.
		return clause, token, false
	}
	exception, token, ok = v.parseSymbol()
	if !ok {
		var message = fmt.Sprintf("Expected an exception symbol following the 'on' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$onClause",
			"$exception")
		panic(message)
	}
	for {
		_, token, ok = v.parseKeyword("matching")
		if !ok {
			break // No more possible matches.
		}
		block, token, ok = v.parseBlock()
		if !ok {
			var message = fmt.Sprintf("Expected a pattern expression and statement block following the 'matching' keyword but received:\n%v\n\n", token)
			message += generateGrammar(
				"$onClause",
				"$exception")
			panic(message)
		}
		blocks.AddItem(block)
	}
	// There must be at least one matching block expression.
	if blocks.IsEmpty() {
		var message = fmt.Sprintf("Expected at least one pattern expression and statement block in the exception clause but received:\n%v\n\n", token)
		message += generateGrammar(
			"$onClause",
			"$exception")
		panic(message)
	}
	clause = sta.OnClause(exception, blocks)
	return clause, token, true
}

// This method attempts to parse a post clause. It returns the post
// clause and whether or not the post clause was successfully parsed.
func (v *parser) parsePostClause() (abs.PostClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var message any
	var bag any
	var clause abs.PostClauseLike
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
	clause = sta.PostClause(message, bag)
	return clause, token, true
}

// This method attempts to parse a procedure. It returns the procedure and
// whether or not the procedure was successfully parsed.
func (v *parser) parseProcedure() (abs.ListLike[any], *Token, bool) {
	var ok bool
	var token *Token
	var statements abs.ListLike[any]
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

// This method attempts to parse a publish clause. It returns the publish
// clause and whether or not the publish clause was successfully parsed.
func (v *parser) parsePublishClause() (abs.PublishClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var event any
	var clause abs.PublishClauseLike
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
	clause = sta.PublishClause(event)
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

// This method attempts to parse a reject clause. It returns the reject
// clause and whether or not the reject clause was successfully parsed.
func (v *parser) parseRejectClause() (abs.RejectClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var message any
	var clause abs.RejectClauseLike
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
	clause = sta.RejectClause(message)
	return clause, token, true
}

// This method attempts to parse a retrieve clause. It returns the retrieve
// clause and whether or not the retrieve clause was successfully parsed.
func (v *parser) parseRetrieveClause() (abs.RetrieveClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var recipient any
	var bag any
	var clause abs.RetrieveClauseLike
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
	clause = sta.RetrieveClause(recipient, bag)
	return clause, token, true
}

// This method attempts to parse a return clause. It returns the return
// clause and whether or not the return clause was successfully parsed.
func (v *parser) parseReturnClause() (abs.ReturnClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var result any
	var clause abs.ReturnClauseLike
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
	clause = sta.ReturnClause(result)
	return clause, token, true
}

// This method attempts to parse a save clause. It returns the save
// clause and whether or not the save clause was successfully parsed.
func (v *parser) parseSaveClause() (abs.SaveClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var draft any
	var citation any
	var clause abs.SaveClauseLike
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
	clause = sta.SaveClause(draft, citation)
	return clause, token, true
}

// This method attempts to parse an select clause. It returns the select clause
// and whether or not the select clause was successfully parsed.
func (v *parser) parseSelectClause() (abs.SelectClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var control any
	var block abs.BlockLike
	var blocks abs.ListLike[abs.BlockLike]
	var clause abs.SelectClauseLike
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
		block, token, ok = v.parseBlock()
		if !ok {
			var message = fmt.Sprintf("Expected a pattern expression and statement block following the 'matching' keyword but received:\n%v\n\n", token)
			message += generateGrammar(
				"$selectClause")
			panic(message)
		}
		blocks.AddItem(block)
	}
	// There must be at least one matching block expression.
	if blocks.IsEmpty() {
		var message = fmt.Sprintf("Expected at least one pattern expression and statement block in the select clause but received:\n%v\n\n", token)
		message += generateGrammar(
			"$selectClause")
		panic(message)
	}
	clause = sta.SelectClause(control, blocks)
	return clause, token, true
}

// This method attempts to parse a statement. It returns the statement and
// whether or not the statement was successfully parsed.
func (v *parser) parseStatement() (abs.StatementLike, *Token, bool) {
	var ok bool
	var token *Token
	var statement abs.StatementLike
	var mainClause any
	var onClause abs.OnClauseLike
	mainClause, token, ok = v.parseMainClause()
	if ok {
		// The exception clause is optional.
		onClause, token, _ = v.parseOnClause()
	}
	statement = sta.StatementWithHandler(mainClause, onClause)
	return statement, token, ok
}

// This method attempts to parse the statements within a procedure. It returns
// an array of the statements and whether or not the statements were
// successfully parsed.
func (v *parser) parseStatements() (abs.ListLike[any], *Token, bool) {
	var statement any
	var statements abs.ListLike[any]
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

// This method attempts to parse a throw clause. It returns the throw
// clause and whether or not the throw clause was successfully parsed.
func (v *parser) parseThrowClause() (abs.ThrowClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var exception any
	var clause abs.ThrowClauseLike
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
	clause = sta.ThrowClause(exception)
	return clause, token, true
}

// This method attempts to parse a while clause. It returns the while clause and
// whether or not the while clause was successfully parsed.
func (v *parser) parseWhileClause() (abs.WhileClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var block abs.BlockLike
	var clause abs.WhileClauseLike
	_, token, ok = v.parseKeyword("while")
	if !ok {
		// This is not a while clause.
		return clause, token, false
	}
	block, token, ok = v.parseBlock()
	if !ok {
		var message = fmt.Sprintf("Expected a conditional expression and statement block following the 'while' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$whileClause")
		panic(message)
	}
	clause = sta.WhileClause(block.GetPattern(), block.GetStatements())
	return clause, token, true
}

// This method attempts to parse a with clause. It returns the with clause and
// whether or not the with clause was successfully parsed.
func (v *parser) parseWithClause() (abs.WithClauseLike, *Token, bool) {
	var ok bool
	var token *Token
	var item ele.Symbol
	var block abs.BlockLike
	var clause abs.WithClauseLike
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
	block, token, ok = v.parseBlock()
	if !ok {
		var message = fmt.Sprintf("Expected a sequential expression and statement block following the 'in' keyword but received:\n%v\n\n", token)
		message += generateGrammar(
			"$withClause",
			"$item")
		panic(message)
	}
	clause = sta.WithClause(item, block.GetPattern(), block.GetStatements())
	return clause, token, true
}
