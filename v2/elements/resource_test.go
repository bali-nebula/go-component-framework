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

func TestResourceWithAuthorityAndPath(t *tes.T) {
	var v = ele.ResourceFromString("https://craterdog.com/About.html")
	ass.Equal(t, "https://craterdog.com/About.html", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/About.html", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithPath(t *tes.T) {
	var v = ele.ResourceFromString("mailto:craterdog@google.com")
	ass.Equal(t, "mailto:craterdog@google.com", v.AsString())
	ass.Equal(t, "mailto", v.GetScheme())
	ass.Equal(t, "", v.GetAuthority())
	ass.Equal(t, "", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQuery(t *tes.T) {
	var v = ele.ResourceFromString("https://craterdog.com/?foo=bar;bar=baz")
	ass.Equal(t, "https://craterdog.com/?foo=bar;bar=baz", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndFragment(t *tes.T) {
	var v = ele.ResourceFromString("https://craterdog.com/#Home")
	ass.Equal(t, "https://craterdog.com/#Home", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}

func TestResourceWithAuthorityAndPathAndQueryAndFragment(t *tes.T) {
	var v = ele.ResourceFromString("https://craterdog.com/?foo=bar;bar=baz#Home")
	ass.Equal(t, "https://craterdog.com/?foo=bar;bar=baz#Home", v.AsString())
	ass.Equal(t, "https", v.GetScheme())
	ass.Equal(t, "craterdog.com", v.GetAuthority())
	ass.Equal(t, "/", v.GetPath())
	ass.Equal(t, "foo=bar;bar=baz", v.GetQuery())
	ass.Equal(t, "Home", v.GetFragment())
}
