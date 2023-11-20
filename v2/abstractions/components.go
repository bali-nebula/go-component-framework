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

type Encapsulated interface {
	IsParameterized() bool
	GetContext() ContextLike
	SetContext(context ContextLike)
	IsAnnotated() bool
	GetNote() NoteLike
	SetNote(note NoteLike)
	GetEntity() Entity
	SetEntity(entity Entity)
	/*
		IsElement() bool
		IsString() bool
		IsRange() bool
		IsCollection() bool
		IsProcedure() bool
	*/
	ExtractAngle() AngleLike
	ExtractBinary() BinaryLike
	ExtractBoolean() BooleanLike
	ExtractCatalog() CatalogLike
	ExtractCitation() CitationLike
	ExtractContinuum() ContinuumLike
	ExtractDuration() DurationLike
	ExtractInterval() IntervalLike
	ExtractList() ListLike
	ExtractMoment() MomentLike
	ExtractName() NameLike
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
	ExtractSpectrum() SpectrumLike
	ExtractStack() StackLike
	ExtractSymbol() SymbolLike
	ExtractTag() TagLike
	ExtractVersion() VersionLike
}

// CONSOLIDATED INTERFACES

type CommentLike interface {
	Sequential[string]
}

type ComponentLike interface {
	Encapsulated
}

type ContextLike interface {
	Sequential[ParameterLike]
	Associative[SymbolLike, ComponentLike]
}

type NoteLike interface {
	Sequential[rune]
}

type ParameterIteratorLike interface {
	Ratcheted[ParameterLike]
}

type ParameterLike interface {
	Binding[SymbolLike, ComponentLike]
}
