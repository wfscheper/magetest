// +build mage

package main

import (
	mt "github.com/wfscheper/magetest"

	// mage:import foo
	_ "github.com/wfscheper/magetest/foo"
)

// Bar wraps magetest.Bar()
func Bar() {
	mt.Bar()
}
