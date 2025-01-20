package main

import (
	"syscall/js"
)

// ----------------------------------------------------------------------------
var imageData []byte
var imageWidth, imageHeight int

// ----------------------------------------------------------------------------
func setLoadImageCallback() {
	js.Global().Set("passImageDataToWasm", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		passImageDataToWasm(args)
		return nil
	}))
}

// ----------------------------------------------------------------------------
func passImageDataToWasm(args []js.Value) interface{} {
	uint8Array := args[0]
	imageWidth = args[1].Int()
	imageHeight = args[2].Int()
	length := uint8Array.Length()
	imageData = make([]byte, length)
	js.CopyBytesToGo(imageData, uint8Array)
	calculateBrightnessValues()
	renderBrightness()
	return nil
}
