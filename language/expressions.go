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
)

// This method attempts to parse an expression. It returns the expression and
// whether or not the expression was successfully parsed.
func (v *parser) parseExpression() (any, bool) {
	// TODO: Reorder these based on how often each type of expression occurs.
	var ok bool
	var expression any
	expression, ok = v.parseComponent()
	/*
		if !ok {
			expression, ok = v.parseVariable()
		}
		if !ok {
			expression, ok = v.parseFunctionExpression()
		}
		if !ok {
			expression, ok = v.parsePrecedenceExpression()
		}
		if !ok {
			expression, ok = v.parseDereferenceExpression()
		}
		if !ok {
			expression, ok = v.parseMessageExpression()
		}
		if !ok {
			expression, ok = v.parseAttributeExpression()
		}
		if !ok {
			expression, ok = v.parseChainExpression()
		}
		if !ok {
			expression, ok = v.parsePowerExpression()
		}
		if !ok {
			expression, ok = v.parseInversionExpression()
		}
		if !ok {
			expression, ok = v.parseArithmeticExpression()
		}
		if !ok {
			expression, ok = v.parseMagnitudeExpression()
		}
		if !ok {
			expression, ok = v.parseComparisonExpression()
		}
		if !ok {
			expression, ok = v.parseComplementExpression()
		}
		if !ok {
			expression, ok = v.parseLogicalExpression()
		}
		if !ok {
			expression, ok = v.parseDefaultExpression()
		}
	*/
	return expression, ok
}
