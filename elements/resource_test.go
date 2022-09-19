/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/elements"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestResourceWithAuthorityAndPath(t *testing.T) {
	var v, ok = elements.ResourceFromString("<https://craterdog.com/About.html>")
	assert.True(t, ok)
	assert.Equal(t, "<https://craterdog.com/About.html>", string(v))
	assert.Equal(t, "<https://craterdog.com/About.html>", v.AsString())
	assert.Equal(t, "https", v.GetScheme())
	assert.Equal(t, "craterdog.com", v.GetAuthority())
	assert.Equal(t, "/About.html", v.GetPath())
	assert.Equal(t, "", v.GetQuery())
	assert.Equal(t, "", v.GetFragment())
}

func TestResourceWithPath(t *testing.T) {
	var v, ok = elements.ResourceFromString("<mailto:craterdog@google.com>")
	assert.True(t, ok)
	assert.Equal(t, "<mailto:craterdog@google.com>", string(v))
	assert.Equal(t, "<mailto:craterdog@google.com>", v.AsString())
	assert.Equal(t, "mailto", v.GetScheme())
	assert.Equal(t, "", v.GetAuthority())
	assert.Equal(t, "craterdog@google.com", v.GetPath())
	assert.Equal(t, "", v.GetQuery())
	assert.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQuery(t *testing.T) {
	var v, ok = elements.ResourceFromString("<https://craterdog.com/?foo=bar;bar=baz>")
	assert.True(t, ok)
	assert.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz>", string(v))
	assert.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz>", v.AsString())
	assert.Equal(t, "https", v.GetScheme())
	assert.Equal(t, "craterdog.com", v.GetAuthority())
	assert.Equal(t, "/", v.GetPath())
	assert.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	assert.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndFragment(t *testing.T) {
	var v, ok = elements.ResourceFromString("<https://craterdog.com/#Home>")
	assert.True(t, ok)
	assert.Equal(t, "<https://craterdog.com/#Home>", string(v))
	assert.Equal(t, "<https://craterdog.com/#Home>", v.AsString())
	assert.Equal(t, "https", v.GetScheme())
	assert.Equal(t, "craterdog.com", v.GetAuthority())
	assert.Equal(t, "/", v.GetPath())
	assert.Equal(t, "", v.GetQuery())
	assert.Equal(t, "Home", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQueryAndFragment(t *testing.T) {
	var v, ok = elements.ResourceFromString("<https://craterdog.com/?foo=bar;bar=baz#Home>")
	assert.True(t, ok)
	assert.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz#Home>", string(v))
	assert.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz#Home>", v.AsString())
	assert.Equal(t, "https", v.GetScheme())
	assert.Equal(t, "craterdog.com", v.GetAuthority())
	assert.Equal(t, "/", v.GetPath())
	assert.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	assert.Equal(t, "Home", v.GetFragment())
}

func TestBadResource(t *testing.T) {
	var v, ok = elements.ResourceFromString("<>")
	assert.False(t, ok)
	assert.Equal(t, "", string(v))
}
