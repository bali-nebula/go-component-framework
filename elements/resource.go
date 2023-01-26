/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package elements

import (
	fmt "fmt"
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	uri "net/url"
)

// RESOURCE IMPLEMENTATION

// This constructor creates a new resource from the specified universal resource
// identifier (URI) string.
func ResourceFromString(v string) abs.ResourceLike {
	var _, err = uri.Parse(v)
	if err != nil {
		var message = fmt.Sprintf("Attempted to construct a resource from an invalid URI: %v", v)
		panic(message)
	}
	return Resource(v)
}

// This constructor creates a new resource from the specified universal resource
// identifier (URI) pointer.
func ResourceFromURI(v *uri.URL) abs.ResourceLike {
	return Resource(v.String())
}

// This type defines the methods associated with a web resource element that
// extends the native Go string type and represents the URI corresponding to
// that web resource.
type Resource string

// SPECTRAL INTERFACE

// This method returns a string value for this spectral element.
func (v Resource) AsString() string {
	return string(v)
}

// SEGMENTED INTERFACE

// This method returns the scheme part of this resource element.
func (v Resource) GetScheme() string {
	var u, _ = uri.Parse(string(v))
	return u.Scheme
}

// This method returns the authority part of this resource element.
func (v Resource) GetAuthority() string {
	var authority string
	var u, _ = uri.Parse(string(v))
	var user = u.User.String()
	var host = u.Host
	if len(user) > 0 {
		authority = user + "@"
	}
	authority += host
	return authority
}

// This method returns the path part of this resource element.
func (v Resource) GetPath() string {
	var u, _ = uri.Parse(string(v))
	return u.Path
}

// This method returns the query part of this resource element.
func (v Resource) GetQuery() string {
	var u, _ = uri.Parse(string(v))
	return u.RawQuery
}

// This method returns the fragment part of this resource element.
func (v Resource) GetFragment() string {
	var u, _ = uri.Parse(string(v))
	return u.Fragment
}
