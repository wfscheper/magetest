// +build mage

package main

import (
	// mage:import
	. "github.com/wfscheper/magetest"

	// mage:import foo
	_ "github.com/wfscheper/magetest/foo"
)

var bar = Bar
