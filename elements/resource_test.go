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
	ele "github.com/craterdog-bali/go-bali-document-notation/elements"
	lan "github.com/craterdog-bali/go-bali-document-notation/language"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestResourceWithAuthorityAndPath(t *tes.T) {
	var v, ok = ele.ResourceFromString("<https://craterdog.com/About.html>")
	ass.True(t, ok)
	ass.Equal(t, "<https://craterdog.com/About.html>", string(v))
	ass.Equal(t, "<https://craterdog.com/About.html>", lan.FormatValue(v))
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithPath(t *tes.T) {
	var v, ok = ele.ResourceFromString("<mailto:craterdog@google.com>")
	ass.True(t, ok)
	ass.Equal(t, "<mailto:craterdog@google.com>", string(v))
	ass.Equal(t, "<mailto:craterdog@google.com>", lan.FormatValue(v))
	ass.Equal(t, "mailto", v.GetScheme())
	ass.Equal(t, "", v.GetAuthority())
	ass.Equal(t, "craterdog@google.com", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQuery(t *tes.T) {
	var v, ok = ele.ResourceFromString("<https://craterdog.com/?foo=bar;bar=baz>")
	ass.True(t, ok)
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz>", string(v))
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz>", lan.FormatValue(v))
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndFragment(t *tes.T) {
	var v, ok = ele.ResourceFromString("<https://craterdog.com/#Home>")
	ass.True(t, ok)
	ass.Equal(t, "<https://craterdog.com/#Home>", string(v))
	ass.Equal(t, "<https://craterdog.com/#Home>", lan.FormatValue(v))
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQueryAndFragment(t *tes.T) {
	var v, ok = ele.ResourceFromString("<https://craterdog.com/?foo=bar;bar=baz#Home>")
	ass.True(t, ok)
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz#Home>", string(v))
	ass.Equal(t, "<https://craterdog.com/?foo=bar;bar=baz#Home>", lan.FormatValue(v))
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestBadResource(t *tes.T) {
	var v, ok = ele.ResourceFromString("<>")
	ass.False(t, ok)
	ass.Equal(t, "", string(v))
}
