/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package strings_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/strings"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBadVersion(t *testing.T) {
	var _, ok = strings.VersionFromString(`v0`)
	assert.False(t, ok)
}

func TestEmptyVersion(t *testing.T) {
	var _, ok = strings.VersionFromString(`v`)
	assert.False(t, ok)
}

func TestVersion(t *testing.T) {
	var v1, ok = strings.VersionFromString("v1.2.3")
	assert.True(t, ok)
	assert.Equal(t, "v1.2.3", v1.AsString())
	assert.False(t, v1.IsEmpty())
	assert.Equal(t, 3, v1.GetSize())
	assert.Equal(t, 1, v1.GetItem(1))
	assert.Equal(t, 3, v1.GetItem(-1))
	var v2 = strings.VersionFromOrdinals(v1.AsArray())
	assert.Equal(t, v1.String(), v2.AsString())
	var v3 = strings.VersionFromOrdinals(v1.GetItems(1, 2))
	assert.Equal(t, "v1.2", string(v3))
	assert.Equal(t, 3, v1.GetIndex(3))
}

func TestVersionsLibrary(t *testing.T) {
	var v1, _ = strings.VersionFromString("v1")
	var v2, _ = strings.VersionFromString("v2.3")
	var v3 = strings.Versions.Concatenate(v1, v2)
	assert.Equal(t, "v1.2.3", v3.AsString())

	assert.False(t, strings.Versions.IsValidNextVersion(v1, v1))
	assert.Equal(t, "v2", strings.Versions.GetNextVersion(v1, 1).AsString())
	assert.True(t, strings.Versions.IsValidNextVersion(v1, strings.Versions.GetNextVersion(v1, 1)))
	assert.False(t, strings.Versions.IsValidNextVersion(strings.Versions.GetNextVersion(v1, 1), v1))
	assert.Equal(t, "v1.1", strings.Versions.GetNextVersion(v1, 2).AsString())
	assert.True(t, strings.Versions.IsValidNextVersion(v1, strings.Versions.GetNextVersion(v1, 2)))
	assert.False(t, strings.Versions.IsValidNextVersion(strings.Versions.GetNextVersion(v1, 2), v1))
	assert.Equal(t, "v1.1", strings.Versions.GetNextVersion(v1, 2).AsString())
	assert.Equal(t, "v1.1", strings.Versions.GetNextVersion(v1, 3).AsString())

	assert.False(t, strings.Versions.IsValidNextVersion(v2, v2))
	assert.Equal(t, "v3", strings.Versions.GetNextVersion(v2, 1).AsString())
	assert.True(t, strings.Versions.IsValidNextVersion(v2, strings.Versions.GetNextVersion(v2, 1)))
	assert.False(t, strings.Versions.IsValidNextVersion(strings.Versions.GetNextVersion(v2, 1), v2))
	assert.Equal(t, "v2.4", strings.Versions.GetNextVersion(v2, 2).AsString())
	assert.True(t, strings.Versions.IsValidNextVersion(v2, strings.Versions.GetNextVersion(v2, 2)))
	assert.False(t, strings.Versions.IsValidNextVersion(strings.Versions.GetNextVersion(v2, 2), v2))
	assert.Equal(t, "v2.3.1", strings.Versions.GetNextVersion(v2, 3).AsString())
	assert.True(t, strings.Versions.IsValidNextVersion(v2, strings.Versions.GetNextVersion(v2, 3)))
	assert.False(t, strings.Versions.IsValidNextVersion(strings.Versions.GetNextVersion(v2, 3), v2))

	assert.False(t, strings.Versions.IsValidNextVersion(v3, v3))
	assert.Equal(t, "v2", strings.Versions.GetNextVersion(v3, 1).AsString())
	assert.True(t, strings.Versions.IsValidNextVersion(v3, strings.Versions.GetNextVersion(v3, 1)))
	assert.False(t, strings.Versions.IsValidNextVersion(strings.Versions.GetNextVersion(v3, 1), v3))
	assert.Equal(t, "v1.3", strings.Versions.GetNextVersion(v3, 2).AsString())
	assert.True(t, strings.Versions.IsValidNextVersion(v3, strings.Versions.GetNextVersion(v3, 2)))
	assert.False(t, strings.Versions.IsValidNextVersion(strings.Versions.GetNextVersion(v3, 2), v3))
	assert.Equal(t, "v1.2.4", strings.Versions.GetNextVersion(v3, 3).AsString())
	assert.True(t, strings.Versions.IsValidNextVersion(v3, strings.Versions.GetNextVersion(v3, 3)))
	assert.False(t, strings.Versions.IsValidNextVersion(strings.Versions.GetNextVersion(v3, 3), v3))
	assert.Equal(t, "v1.2.3.1", strings.Versions.GetNextVersion(v3, 4).AsString())
	assert.True(t, strings.Versions.IsValidNextVersion(v3, strings.Versions.GetNextVersion(v3, 4)))
	assert.False(t, strings.Versions.IsValidNextVersion(strings.Versions.GetNextVersion(v3, 4), v3))
}
