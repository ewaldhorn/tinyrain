package main

import "syscall/js"

// ----------------------------------------------------------------------------
const VERSION = "0.0.1h"
const NAME = "TinyRain"

// ----------------------------------------------------------------------------
func getVersionString() string {
	return NAME + " v" + VERSION
}

// ----------------------------------------------------------------------------
func setVersionCallback() {
	js.Global().Set("getVersion", js.FuncOf(func(this js.Value, args []js.Value) any {
		return getVersionString()
	}))
}
