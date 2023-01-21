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
	uri "net/url"
)

// RESOURCE IMPLEMENTATION

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
