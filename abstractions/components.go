/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package abstractions

// TYPE DEFINITIONS

type (
	Annotation any
	Entity     any
)

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all component-like types.
type Encapsulated interface {
	GetEntity() Entity
	SetEntity(entity Entity)
	ExtractAngle() AngleLike
	ExtractBinary() BinaryLike
	ExtractBoolean() BooleanLike
	ExtractCatalog() CatalogLike
	ExtractContinuum() ContinuumLike[Continuous]
	ExtractDuration() DurationLike
	ExtractInterval() IntervalLike[Discrete]
	ExtractList() ListLike
	ExtractMoment() MomentLike
	ExtractMoniker() MonikerLike
	ExtractNarrative() NarrativeLike
	ExtractNumber() NumberLike
	ExtractPattern() PatternLike
	ExtractPercentage() PercentageLike
	ExtractProbability() ProbabilityLike
	ExtractProcedure() ProcedureLike
	ExtractQueue() QueueLike
	ExtractQuote() QuoteLike
	ExtractResource() ResourceLike
	ExtractSet() SetLike
	ExtractSpectrum() SpectrumLike[Lexical]
	ExtractStack() StackLike
	ExtractSymbol() SymbolLike
	ExtractTag() TagLike
	ExtractVersion() VersionLike
	IsParameterized() bool
	GetContext() ContextLike
	SetContext(context ContextLike)
	IsAnnotated() bool
	GetNote() NoteLike
	SetNote(note NoteLike)
}

// This interface defines the methods supported by all parameterized sequences
// whose values consist of name-value pairs.
type Parameterized interface {
	GetNames() Sequential[SymbolLike]
	GetValue(name SymbolLike) ComponentLike
	SetValue(name SymbolLike, value ComponentLike)
}

// CONSOLIDATED INTERFACES

type ComponentLike interface {
	Encapsulated
}

type NoteLike interface {
	Sequential[rune]
}

type CommentLike interface {
	Sequential[string]
}

type ParameterLike interface {
	Binding[SymbolLike, ComponentLike]
}

type ContextLike interface {
	Parameterized
}
