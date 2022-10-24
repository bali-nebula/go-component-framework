/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

// SYMBOL INTERFACE

// This type defines the methods associated with a symbol element that
// extends the native Go string type and represents the runes that
// make up the symbol.
type Symbol string

// SYMBOLIC INTERFACE

// This method returns the identifier for this symbol.
func (v Symbol) GetIdentifier() string {
	return string(v)
}
