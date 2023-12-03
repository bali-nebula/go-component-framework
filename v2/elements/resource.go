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
	uti "github.com/bali-nebula/go-component-framework/v2/utilities"
	uri "net/url"
)

// CLASS DEFINITIONS

// This private type implements the ResourceLike interface.  It extends the
// native Go `string` type and represents the URI corresponding to that web
// resource.
type resource_ string

// This private type defines the structure associated with the class constants
// and class functions for the resource elements.
type resourceClass_ struct {
	// This class has no class constants.
}

// CLASS CONSTRUCTORS

// This constructor creates a new resource element from the specified universal
// resource identifier (URI) string embedded in "<" and ">" delimiters.
func (c *resourceClass_) FromString(string_ string) ResourceLike {
	var matches = uti.ResourceMatcher.FindStringSubmatch(string_)
	if len(matches) == 0 {
		var message = fmt.Sprintf("Attempted to construct a resource from an invalid string: %v", string_)
		panic(message)
	}
	var resource = resource_(matches[1]) // Strip off the "<" and ">" delimiters.
	return resource
}

// This constructor creates a new resource element from the specified universal
// resource identifier (URI) pointer.
func (c *resourceClass_) FromURI(url *uri.URL) ResourceLike {
	var resource = resource_(url.String())
	return resource
}

// CLASS METHODS

// Lexical Interface

// This method returns a string value for this lexical element.
func (v resource_) AsString() string {
	return "<" + string(v) + ">"
}

// Segmented Interface

// This method returns the scheme part of this resource element.
func (v resource_) GetScheme() string {
	var u, _ = uri.Parse(string(v))
	return u.Scheme
}

// This method returns the authority part of this resource element.
func (v resource_) GetAuthority() string {
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
func (v resource_) GetPath() string {
	var u, _ = uri.Parse(string(v))
	return u.Path
}

// This method returns the query part of this resource element.
func (v resource_) GetQuery() string {
	var u, _ = uri.Parse(string(v))
	return u.RawQuery
}

// This method returns the fragment part of this resource element.
func (v resource_) GetFragment() string {
	var u, _ = uri.Parse(string(v))
	return u.Fragment
}
