/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	ele "github.com/bali-nebula/go-component-framework/v2/elements"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestCitation(t *tes.T) {
	var v1 = ele.Citation.FromString("/bali/types/abstractions/String/v1.2.3")
	ass.Equal(t, "/bali/types/abstractions/String/v1.2.3", v1.AsString())
	ass.Equal(t, "/bali/types/abstractions/String", v1.GetName())
	ass.Equal(t, "v1.2.3", v1.GetVersion())
}
