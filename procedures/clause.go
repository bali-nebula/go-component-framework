/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package procedures

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
)

// PUBLIC FUNCTIONS

// This function returns a string describing the type of the specified
// clause.
func GetType(clause abs.Clause) string {
	switch value := clause.(type) {
	case *acceptClause:
		return "AcceptClause"
	case *breakClause:
		return "BreakClause"
	case *checkoutClause:
		return "CheckoutClause"
	case *continueClause:
		return "ContinueClause"
	case *discardClause:
		return "DiscardClause"
	case *ifClause:
		return "IfClause"
	case *letClause:
		return "LetClause"
	case *notarizeClause:
		return "NotarizeClause"
	case *postClause:
		return "PostClause"
	case *publishClause:
		return "PublishClause"
	case *rejectClause:
		return "RejectClause"
	case *retrieveClause:
		return "RetrieveClause"
	case *returnClause:
		return "ReturnClause"
	case *saveClause:
		return "SaveClause"
	case *selectClause:
		return "SelectClause"
	case *throwClause:
		return "ThrowClause"
	case *whileClause:
		return "WhileClause"
	case *withClause:
		return "WithClause"
	default:
		var message = fmt.Sprintf("An invalid expression type was found: %T", value)
		panic(message)
	}
}
