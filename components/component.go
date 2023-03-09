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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
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

// COMPONENT IMPLEMENTATION

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

// This method returns the entity for this component as a continuum.
func (v *component) ExtractContinuum() abs.ContinuumLike[abs.Continuous] {
	return v.entity.(abs.ContinuumLike[abs.Continuous])
}

// This method returns the entity for this component as a duration.
func (v *component) ExtractDuration() abs.DurationLike {
	return v.entity.(abs.DurationLike)
}

// This method returns the entity for this component as an interval.
func (v *component) ExtractInterval() abs.IntervalLike[abs.Discrete] {
	return v.entity.(abs.IntervalLike[abs.Discrete])
}

// This method returns the entity for this component as a list.
func (v *component) ExtractList() abs.ListLike {
	return v.entity.(abs.ListLike)
}

// This method returns the entity for this component as a moment.
func (v *component) ExtractMoment() abs.MomentLike {
	return v.entity.(abs.MomentLike)
}

// This method returns the entity for this component as a moniker.
func (v *component) ExtractMoniker() abs.MonikerLike {
	return v.entity.(abs.MonikerLike)
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
func (v *component) ExtractSpectrum() abs.SpectrumLike[abs.Quantized] {
	return v.entity.(abs.SpectrumLike[abs.Quantized])
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
