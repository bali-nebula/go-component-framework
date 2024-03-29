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

type Custodial interface {
	Exists() bool
	Load() []byte
	Store(configuration []byte)
	Delete()
}

type Mechanized interface {
	GetState() int
	SetState(state int)
	TransitionState(event int) int
}

// CONSOLIDATED INTERFACES

type ConfiguratorLike interface {
	Custodial
}

type ControllerLike interface {
	Mechanized
}
