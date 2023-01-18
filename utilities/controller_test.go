/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package utilities_test

import (
	uti "github.com/bali-nebula/go-component-framework/utilities"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestController(t *tes.T) {
	var events = []uti.Event{"initialized", "processed", "finalized"}
	var states = []uti.State{"state1", "state2", "state3"}
	var table = [][]uti.State{
		{"state2", "invalid", "invalid"},
		{"invalid", "state2", "state3"},
		{"invalid", "invalid", "invalid"},
	}
	var controller = uti.Controller(events, states, table)
	ass.Equal(t, uti.State("state1"), controller.GetState())
	ass.True(t, controller.IsValid("initialized"))
	ass.False(t, controller.IsValid("processed"))
	ass.False(t, controller.IsValid("finalized"))
	ass.Equal(t, uti.State("state2"), controller.TransitionState("initialized"))
	ass.False(t, controller.IsValid("initialized"))
	ass.True(t, controller.IsValid("processed"))
	ass.True(t, controller.IsValid("finalized"))
	ass.Equal(t, uti.State("state2"), controller.TransitionState("processed"))
	ass.False(t, controller.IsValid("initialized"))
	ass.True(t, controller.IsValid("processed"))
	ass.True(t, controller.IsValid("finalized"))
	ass.Equal(t, uti.State("state3"), controller.TransitionState("finalized"))
	ass.False(t, controller.IsValid("initialized"))
	ass.False(t, controller.IsValid("processed"))
	ass.False(t, controller.IsValid("finalized"))
	controller.ResetState()
	ass.Equal(t, uti.State("state1"), controller.GetState())
	ass.True(t, controller.IsValid("initialized"))
	ass.False(t, controller.IsValid("processed"))
	ass.False(t, controller.IsValid("finalized"))
}
