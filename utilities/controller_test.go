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
	abs "github.com/bali-nebula/go-component-framework/abstractions"
	uti "github.com/bali-nebula/go-component-framework/utilities"
	ass "github.com/stretchr/testify/assert"
	tes "testing"
)

func TestController(t *tes.T) {
	var events = []abs.Event{"initialized", "processed", "finalized"}
	var states = []abs.State{"state1", "state2", "state3"}
	var table = [][]abs.State{
		{"state2", "invalid", "invalid"},
		{"invalid", "state2", "state3"},
		{"invalid", "invalid", "invalid"},
	}
	var controller = uti.Controller(events, states, table, "state1")
	ass.Equal(t, abs.State("state1"), controller.GetState())
	ass.Equal(t, abs.State("state2"), controller.TransitionState("initialized"))
	ass.Equal(t, abs.State("state2"), controller.TransitionState("processed"))
	ass.Equal(t, abs.State("state3"), controller.TransitionState("finalized"))
	controller.SetState("state1")
	ass.Equal(t, abs.State("state1"), controller.GetState())
}
