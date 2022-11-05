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

// CONTINUE CLAUSE IMPLEMENTATION

// This constructor creates a new continue clause.
func ContinueClause() abs.ContinueClauseLike {
	var v = &continueClause{}
	return v
}

// This type defines the structure and methods associated with a continue clause.
type continueClause struct {
}

// This method is a dummy method that always returns true.
func (v *continueClause) IsContinueClause() bool {
	return true
}
