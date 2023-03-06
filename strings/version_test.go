/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings_test

import (
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	str "github.com/bali-nebula/go-component-framework/strings"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestVersion(t *tes.T) {
	var v1 = str.VersionFromString("1.2.3")
	ass.Equal(t, "1.2.3", v1.AsString())
	ass.False(t, v1.IsEmpty())
	ass.Equal(t, 3, v1.GetSize())
	ass.Equal(t, abs.Ordinal(1), v1.GetValue(1))
	ass.Equal(t, abs.Ordinal(3), v1.GetValue(-1))
	var v2 = str.VersionFromArray(v1.AsArray())
	ass.Equal(t, v1.AsString(), v2.AsString())
	var v3 = str.VersionFromSequence(v1.GetValues(1, 2))
	ass.Equal(t, "1.2", v3.AsString())
}

func TestVersionsLibrary(t *tes.T) {
	var v1 = str.VersionFromString("1")
	var v2 = str.VersionFromString("2.3")
	var v3 = str.Versions.Concatenate(v1, v2)
	ass.Equal(t, "1.2.3", v3.AsString())

	ass.False(t, str.Versions.IsValidNextVersion(v1, v1))
	ass.Equal(t, "2", str.Versions.GetNextVersion(v1, 0).AsString())
	ass.Equal(t, "2", str.Versions.GetNextVersion(v1, 1).AsString())
	ass.True(t, str.Versions.IsValidNextVersion(v1, str.Versions.GetNextVersion(v1, 1)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v1, 1), v1))
	ass.Equal(t, "1.1", str.Versions.GetNextVersion(v1, 2).AsString())
	ass.True(t, str.Versions.IsValidNextVersion(v1, str.Versions.GetNextVersion(v1, 2)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v1, 2), v1))
	ass.Equal(t, "1.1", str.Versions.GetNextVersion(v1, 3).AsString())

	ass.False(t, str.Versions.IsValidNextVersion(v2, v2))
	ass.Equal(t, "3", str.Versions.GetNextVersion(v2, 1).AsString())
	ass.True(t, str.Versions.IsValidNextVersion(v2, str.Versions.GetNextVersion(v2, 1)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v2, 1), v2))
	ass.Equal(t, "2.4", str.Versions.GetNextVersion(v2, 0).AsString())
	ass.Equal(t, "2.4", str.Versions.GetNextVersion(v2, 2).AsString())
	ass.True(t, str.Versions.IsValidNextVersion(v2, str.Versions.GetNextVersion(v2, 2)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v2, 2), v2))
	ass.Equal(t, "2.3.1", str.Versions.GetNextVersion(v2, 3).AsString())
	ass.True(t, str.Versions.IsValidNextVersion(v2, str.Versions.GetNextVersion(v2, 3)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v2, 3), v2))

	ass.False(t, str.Versions.IsValidNextVersion(v3, v3))
	ass.Equal(t, "2", str.Versions.GetNextVersion(v3, 1).AsString())
	ass.True(t, str.Versions.IsValidNextVersion(v3, str.Versions.GetNextVersion(v3, 1)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v3, 1), v3))
	ass.Equal(t, "1.3", str.Versions.GetNextVersion(v3, 2).AsString())
	ass.True(t, str.Versions.IsValidNextVersion(v3, str.Versions.GetNextVersion(v3, 2)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v3, 2), v3))
	ass.Equal(t, "1.2.4", str.Versions.GetNextVersion(v3, 0).AsString())
	ass.Equal(t, "1.2.4", str.Versions.GetNextVersion(v3, 3).AsString())
	ass.True(t, str.Versions.IsValidNextVersion(v3, str.Versions.GetNextVersion(v3, 3)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v3, 3), v3))
	ass.Equal(t, "1.2.3.1", str.Versions.GetNextVersion(v3, 4).AsString())
	ass.True(t, str.Versions.IsValidNextVersion(v3, str.Versions.GetNextVersion(v3, 4)))
	ass.False(t, str.Versions.IsValidNextVersion(str.Versions.GetNextVersion(v3, 4), v3))
}
