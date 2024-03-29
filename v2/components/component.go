/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package components

import (
	abs "github.com/bali-nebula/go-component-framework/v2/abstractions"
	col "github.com/craterdog/go-collection-framework/v2"
)

// COMPONENT IMPLEMENTATION

// This constructor creates a new component.
func Component(entity abs.Entity) abs.ComponentLike {
	var v = &component{}
	// Perform argument validation.
	v.SetEntity(entity)
	return v
}

// This constructor creates a new component with the specified context.
func ComponentWithContext(entity abs.Entity, context abs.ContextLike) abs.ComponentLike {
	var v = &component{}
	// Perform argument validation.
	v.SetEntity(entity)
	v.SetContext(context)
	return v
}

// This type defines the structure and methods associated with a component.
type component struct {
	entity  abs.Entity
	context abs.ContextLike
	note    abs.NoteLike
}

// ENCAPSULATED INTERFACE

// This method returns the entity for this component.
func (v *component) GetEntity() abs.Entity {
	return v.entity
}

// This method sets the entity for this component.
func (v *component) SetEntity(entity abs.Entity) {
	if entity == nil {
		panic("A component requires an entity.")
	}
	v.entity = entity
}

// This method returns the entity for this component as an angle.
func (v *component) ExtractAngle() abs.AngleLike {
	return v.entity.(abs.AngleLike)
}

// This method returns the entity for this component as a binary.
func (v *component) ExtractBinary() abs.BinaryLike {
	return v.entity.(abs.BinaryLike)
}

// This method returns the entity for this component as a boolean.
func (v *component) ExtractBoolean() abs.BooleanLike {
	return v.entity.(abs.BooleanLike)
}

// This method returns the entity for this component as a catalog.
func (v *component) ExtractCatalog() abs.CatalogLike {
	return v.entity.(abs.CatalogLike)
}

// This method returns the entity for this component as a citation.
func (v *component) ExtractCitation() abs.CitationLike {
	return v.entity.(abs.CitationLike)
}

// This method returns the entity for this component as a continuum.
func (v *component) ExtractContinuum() abs.ContinuumLike {
	return v.entity.(abs.ContinuumLike)
}

// This method returns the entity for this component as a duration.
func (v *component) ExtractDuration() abs.DurationLike {
	return v.entity.(abs.DurationLike)
}

// This method returns the entity for this component as an interval.
func (v *component) ExtractInterval() abs.IntervalLike {
	return v.entity.(abs.IntervalLike)
}

// This method returns the entity for this component as a list.
func (v *component) ExtractList() abs.ListLike {
	return v.entity.(abs.ListLike)
}

// This method returns the entity for this component as a moment.
func (v *component) ExtractMoment() abs.MomentLike {
	return v.entity.(abs.MomentLike)
}

// This method returns the entity for this component as a name.
func (v *component) ExtractName() abs.NameLike {
	return v.entity.(abs.NameLike)
}

// This method returns the entity for this component as a narrative.
func (v *component) ExtractNarrative() abs.NarrativeLike {
	return v.entity.(abs.NarrativeLike)
}

// This method returns the entity for this component as a number.
func (v *component) ExtractNumber() abs.NumberLike {
	return v.entity.(abs.NumberLike)
}

// This method returns the entity for this component as a pattern.
func (v *component) ExtractPattern() abs.PatternLike {
	return v.entity.(abs.PatternLike)
}

// This method returns the entity for this component as a percentage.
func (v *component) ExtractPercentage() abs.PercentageLike {
	return v.entity.(abs.PercentageLike)
}

// This method returns the entity for this component as a probability.
func (v *component) ExtractProbability() abs.ProbabilityLike {
	return v.entity.(abs.ProbabilityLike)
}

// This method returns the entity for this component as a procedure.
func (v *component) ExtractProcedure() abs.ProcedureLike {
	return v.entity.(abs.ProcedureLike)
}

// This method returns the entity for this component as a queue.
func (v *component) ExtractQueue() abs.QueueLike {
	return v.entity.(abs.QueueLike)
}

// This method returns the entity for this component as a quote.
func (v *component) ExtractQuote() abs.QuoteLike {
	return v.entity.(abs.QuoteLike)
}

// This method returns the entity for this component as a resource.
func (v *component) ExtractResource() abs.ResourceLike {
	return v.entity.(abs.ResourceLike)
}

// This method returns the entity for this component as a set.
func (v *component) ExtractSet() abs.SetLike {
	return v.entity.(abs.SetLike)
}

// This method returns the entity for this component as a spectrum.
func (v *component) ExtractSpectrum() abs.SpectrumLike {
	return v.entity.(abs.SpectrumLike)
}

// This method returns the entity for this component as a stack.
func (v *component) ExtractStack() abs.StackLike {
	return v.entity.(abs.StackLike)
}

// This method returns the entity for this component as a symbol.
func (v *component) ExtractSymbol() abs.SymbolLike {
	return v.entity.(abs.SymbolLike)
}

// This method returns the entity for this component as a tag.
func (v *component) ExtractTag() abs.TagLike {
	return v.entity.(abs.TagLike)
}

// This method returns the entity for this component as a version.
func (v *component) ExtractVersion() abs.VersionLike {
	return v.entity.(abs.VersionLike)
}

// This method determines whether or not this component is parameterized.
func (v *component) IsParameterized() bool {
	return v.context != nil
}

// This method returns the context for this component.
func (v *component) GetContext() abs.ContextLike {
	return v.context
}

// This method sets the context for this component.
func (v *component) SetContext(context abs.ContextLike) {
	v.context = context
}

// This method determines whether or not this component is annotated.
func (v *component) IsAnnotated() bool {
	return v.note != nil
}

// This method returns the note for this component.
func (v *component) GetNote() abs.NoteLike {
	return v.note
}

// This method sets the note for this component.
func (v *component) SetNote(note abs.NoteLike) {
	v.note = note
}

// COMPONENT ITERATOR IMPLEMENTATION

// This constructor creates a new instance of a components iterator that can be
// used to traverse the components in the specified sequence.
func ComponentIterator(sequence abs.Sequential[abs.ComponentLike]) abs.ComponentIteratorLike {
	var v = col.Iterator[abs.ComponentLike](sequence)
	return &components{v}
}

// This type defines the structure and methods for a components iterator. The
// iterator operates on a sequence of components.
type components struct {
	abs.ComponentIteratorLike
}
