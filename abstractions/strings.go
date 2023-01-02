/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

import (
	col "github.com/craterdog/go-collection-framework"
)

// TYPE DEFINITIONS

type (
	String  any
	Name    string
	Ordinal uint
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all numeric elements.
type Spectral interface {
	AsString() string
}

// TYPE ALIASES

// These type assignments hide the dependencies on the package used to implement
// the collection types. It preserves the collection interfaces in a way that
// will allow them to evolve separately as needed. Currently, the interfaces are
// synchronized.
type (
	Bytes    = col.Sequential[byte]
	Names    = col.Sequential[Name]
	Ordinals = col.Sequential[Ordinal]
)
