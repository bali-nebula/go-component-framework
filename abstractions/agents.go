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

// INDIVIDUAL INTERFACES

// This interface defines the methods supported by all custodial agents.
type Custodial interface {
	Exists() bool
	Load() []byte
	Store(configuration []byte)
	Delete()
}

// This interface defines the methods supported by all mechanized agents.
type Mechanized interface {
	GetState() int
	SetState(state int)
	TransitionState(event int) int
}

// This interface defines the methods supported by all trusted security
// modules.
type Trusted interface {
	GetProtocol() string
	DigestBytes(bytes []byte) []byte
	IsValid(key []byte, signature []byte, bytes []byte) bool
}

// This interface defines the methods supported by all hardened security
// modules.
type Hardened interface {
	GetTag() string
	GenerateKeys() []byte
	SignBytes(bytes []byte) []byte
	RotateKeys() []byte
	EraseKeys()
}

// CONSOLIDATED INTERFACES

type ConfiguratorLike interface {
	Custodial
}

type ControllerLike interface {
	Mechanized
}

type SecurityModuleLike interface {
	Trusted
	Hardened
}
