/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package procedures

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// ON CLAUSE IMPLEMENTATION

// This constructor creates a new on clause.
func OnClause(exception abs.Symbolic, handlers abs.ListLike[abs.BlockLike]) abs.OnClauseLike {
	var v = &onClause{}
	// Perform argument validation.
	v.SetException(exception)
	v.SetHandlers(handlers)
	return v
}

// This type defines the structure and methods associated with an on clause.
type onClause struct {
	exception abs.Symbolic
	handlers  abs.ListLike[abs.BlockLike]
}

// This method returns the exception symbol for this on clause.
func (v *onClause) GetException() abs.Symbolic {
	return v.exception
}

// This method sets the exception symbol for this on clause.
func (v *onClause) SetException(exception abs.Symbolic) {
	if exception == nil {
		panic("An on clause requires an exception symbol.")
	}
	v.exception = exception
}

// This method returns the handler at the specified index from this on clause.
func (v *onClause) GetHandler(index int) abs.BlockLike {
	return v.handlers.GetItem(index)
}

// This method sets the handler at the specified index for this on clause.
func (v *onClause) SetHandler(index int, handler abs.BlockLike) {
	if handler == nil {
		panic("Each index in an on clause requires a handler block.")
	}
	v.handlers.SetItem(index, handler)
}

// This method returns the list of handlers for this on clause.
func (v *onClause) GetHandlers() abs.ListLike[abs.BlockLike] {
	return v.handlers
}

// This method sets the list of handlers for this on clause.
func (v *onClause) SetHandlers(handlers abs.ListLike[abs.BlockLike]) {
	if handlers == nil || handlers.IsEmpty() {
		panic("An on clause requires at least one handler.")
	}
	v.handlers = handlers
}
