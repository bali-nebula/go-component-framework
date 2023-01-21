/*******************************************************************************
 *   Copyright (c) 2009-2023 Crater Dog Technologies™.  All Rights Reserved.   *
 *******************************************************************************
 * DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.               *
 *                                                                             *
 * This code is free software; you can redistribute it and/or modify it under  *
 * the terms of The MIT License (MIT), as published by the Open Source         *
 * Initiative. (See http://opensource.org/licenses/MIT)                        *
 *******************************************************************************/

package agents_test

import (
	age "github.com/bali-nebula/go-component-framework/agents"
	ass "github.com/stretchr/testify/assert"
	sts "strings"
	tes "testing"
)

func TestConfiguratorWithNoDirectory(t *tes.T) {
	var directory = ""
	var filename = "test.bali"
	var configurator = age.Configurator(directory, filename)
	ass.False(t, configurator.Exists())
	var configuration = []byte("This is a test configuration.\n")

	configurator.Store(configuration)
	ass.True(t, configurator.Exists())
	ass.Equal(t, configuration, configurator.Load())
	configurator.Delete()
	defer func() {
		if e := recover(); e != nil {
			var message = e.(string)
			ass.True(t, sts.HasPrefix(message, "Could not read the configuration file"))
		} else {
			ass.Fail(t, "Test should result in recovered panic.")
		}
	}()
	configurator.Load() // This should panic.
}
