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

import (
	abs "github.com/craterdog-bali/go-bali-document-notation/abstractions"
)

// RESOURCE INTERFACE

// This constructor attempts to create a new web resource from the specified
// formatted string. It returns a resource value and whether or not the string
// contained a valid resource.
// For valid string formats for this type see `../abstractions/language.go`.
func ResourceFromString(v string) (Resource, bool) {
	var resource, ok = stringToResource(v)
	return Resource(resource), ok
}

// This type defines the methods associated with a web resource element that
// extends the native Go string type and represents the URI corresponding to
// that web resource.
type Resource string

// LEXICAL INTERFACE

// This method returns the canonical string for this element.
func (v Resource) AsString() string {
	return string(v)
}

// This method implements the standard Go Stringer interface.
func (v Resource) String() string {
	return v.AsString()
}

// SEGMENTED INTERFACE

// This method returns the scheme part of this resource element.
func (v Resource) GetScheme() string {
	var matches = abs.ScanResource([]byte(v))
	var scheme = matches[2]
	return scheme
}

// This method returns the authority part of this resource element.
func (v Resource) GetAuthority() string {
	var matches = abs.ScanResource([]byte(v))
	var authority = matches[3]
	return authority
}

// This method returns the path part of this resource element.
func (v Resource) GetPath() string {
	var matches = abs.ScanResource([]byte(v))
	return matches[4]
}

// This method returns the query part of this resource element.
func (v Resource) GetQuery() string {
	var matches = abs.ScanResource([]byte(v))
	var query = matches[5]
	return query
}

// This method returns the fragment part of this resource element.
func (v Resource) GetFragment() string {
	var matches = abs.ScanResource([]byte(v))
	var fragment = matches[6]
	return fragment
}

// PRIVATE FUNCTIONS

// This function parses a resource string and returns the corresponding web URI.
func stringToResource(v string) (string, bool) {
	var resource string
	var ok = true
	var matches = abs.ScanResource([]byte(v))
	switch {
	case len(matches) == 0:
		ok = false
	default:
		resource = matches[0]
	}
	return resource, ok
}
