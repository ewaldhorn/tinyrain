package main

import (
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/colour"
)

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
	width := args[1].Int()
	height := args[2].Int()
	length := uint8Array.Length()
	data := make([]byte, length)
	js.CopyBytesToGo(data, uint8Array)

	for x := range 800 {
		for y := range height {
			offset := (x * 4) + (y * 4 * width)

			// don't bother if we are outside our area
			if offset < 0 || offset >= length {
				continue
			}

			col := colour.NewColour(data[offset], data[offset+1], data[offset+2], data[offset+3])
			mainCanvas.ColourPutPixel(x, y, *col)
		}
	}
	mainCanvas.Render()

	return nil
}
