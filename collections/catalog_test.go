/*******************************************************************************
 *   Copyright (c) 2009-2022 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package collections_test

import (
	"github.com/craterdog-bali/go-bali-document-notation/abstractions"
	"github.com/craterdog-bali/go-bali-document-notation/agents"
	"github.com/craterdog-bali/go-bali-document-notation/collections"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCatalogsWithStringsAndIntegers(t *testing.T) {
	var association1 = collections.Association[string, int]("foo", 1)
	var association2 = collections.Association[string, int]("bar", 2)
	var association3 = collections.Association[string, int]("baz", 3)
	var associations = []abstractions.AssociationLike[string, int]{association2, association3}
	var catalog = collections.Catalog[string, int]()
	assert.True(t, catalog.IsEmpty())
	assert.Equal(t, 0, catalog.GetSize())
	assert.Equal(t, []string{}, catalog.GetKeys())
	assert.Equal(t, []abstractions.AssociationLike[string, int]{}, catalog.AsArray())
	var iterator = agents.Iterator[abstractions.AssociationLike[string, int]](catalog)
	assert.False(t, iterator.HasNext())
	assert.False(t, iterator.HasPrevious())
	iterator.ToStart()
	iterator.ToEnd()
	catalog.SortAssociations()
	catalog.RemoveAll()
	catalog.AddAssociation(association1)
	assert.False(t, catalog.IsEmpty())
	assert.Equal(t, 1, catalog.GetSize())
	assert.Equal(t, association1, catalog.GetItem(1))
	catalog.AddAssociations(associations)
	assert.Equal(t, 3, catalog.GetSize())
	assert.Equal(t, association1, catalog.GetItem(1))
	assert.Equal(t, associations, catalog.GetItems(2, 3))
	iterator = agents.Iterator[abstractions.AssociationLike[string, int]](catalog)
	assert.True(t, iterator.HasNext())
	assert.False(t, iterator.HasPrevious())
	assert.Equal(t, association1, iterator.GetNext())
	assert.True(t, iterator.HasPrevious())
	iterator.ToEnd()
	assert.False(t, iterator.HasNext())
	assert.True(t, iterator.HasPrevious())
	assert.Equal(t, association3, iterator.GetPrevious())
	assert.True(t, iterator.HasNext())
	assert.Equal(t, []string{"foo", "bar", "baz"}, catalog.GetKeys())
	assert.Equal(t, 3, int(catalog.GetValue("baz")))
	assert.Equal(t, 2, int(catalog.SetValue("bar", 5)))
	assert.Equal(t, []int{1, 5}, catalog.GetValues([]string{"foo", "bar"}))
	catalog.SortAssociations()
	assert.Equal(t, []string{"bar", "baz", "foo"}, catalog.GetKeys())
	catalog.ReverseAssociations()
	assert.Equal(t, []string{"foo", "baz", "bar"}, catalog.GetKeys())
	catalog.ReverseAssociations()
	assert.Equal(t, []int{1, 5}, catalog.RemoveValues([]string{"foo", "bar"}))
	assert.Equal(t, 1, catalog.GetSize())
	assert.Equal(t, 3, int(catalog.RemoveValue("baz")))
	assert.True(t, catalog.IsEmpty())
	assert.Equal(t, 0, catalog.GetSize())
	catalog.RemoveAll()
	assert.True(t, catalog.IsEmpty())
	assert.Equal(t, 0, catalog.GetSize())
}

func TestCatalogsWithMerge(t *testing.T) {
	var catalogs = collections.Catalogs[string, int]()
	var association1 = collections.Association[string, int]("foo", 1)
	var association2 = collections.Association[string, int]("bar", 2)
	var association3 = collections.Association[string, int]("baz", 3)
	var catalog1 = collections.Catalog[string, int]()
	catalog1.AddAssociation(association1)
	catalog1.AddAssociation(association2)
	var catalog2 = collections.Catalog[string, int]()
	catalog2.AddAssociation(association2)
	catalog2.AddAssociation(association3)
	var catalog3 = catalogs.Merge(catalog1, catalog2)
	var catalog4 = collections.Catalog[string, int]()
	catalog4.AddAssociation(association1)
	catalog4.AddAssociation(association2)
	catalog4.AddAssociation(association3)
	assert.True(t, agents.CompareValues(catalog3, catalog4))
}

func TestCatalogsWithExtract(t *testing.T) {
	var catalogs = collections.Catalogs[string, int]()
	var association1 = collections.Association[string, int]("foo", 1)
	var association2 = collections.Association[string, int]("bar", 2)
	var association3 = collections.Association[string, int]("baz", 3)
	var catalog1 = collections.Catalog[string, int]()
	catalog1.AddAssociation(association1)
	catalog1.AddAssociation(association2)
	catalog1.AddAssociation(association3)
	var catalog2 = catalogs.Extract(catalog1, []string{"foo", "baz"})
	var catalog3 = collections.Catalog[string, int]()
	catalog3.AddAssociation(association1)
	catalog3.AddAssociation(association3)
	assert.True(t, agents.CompareValues(catalog2, catalog3))
}

func TestCatalogsWithEmptyCatalogs(t *testing.T) {
	var keys = []int{}
	var catalogs = collections.Catalogs[int, string]()
	var catalog1 = collections.Catalog[int, string]()
	var catalog2 = collections.Catalog[int, string]()
	var catalog3 = catalogs.Merge(catalog1, catalog2)
	var catalog4 = catalogs.Extract(catalog1, keys)
	assert.True(t, agents.CompareValues(catalog1, catalog2))
	assert.True(t, agents.CompareValues(catalog2, catalog3))
	assert.True(t, agents.CompareValues(catalog3, catalog4))
	assert.True(t, agents.CompareValues(catalog4, catalog1))
}
