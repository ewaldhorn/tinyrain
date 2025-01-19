package main

import (
	"github.com/ewaldhorn/dommie/dom"
	"github.com/ewaldhorn/tinycanvas/colour"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

var mainCanvas *tinycanvas.TinyCanvas

// ----------------------------------------------------------------------------
// bootstrap is a JavaScript-side defined function, called by Wasm in the main
// Go function
//
//export bootstrap
func bootstrap()

// ----------------------------------------------------------------------------
func main() {
	dom.Log("Starting...")

	setCallbacks()
	dom.Hide("loading")
	bootstrap()

	mainCanvas = tinycanvas.NewTinyCanvas(800, 600)
	mainCanvas.ClearScreen(*colour.NewRandomColour())
	mainCanvas.ColourCircle(300, 300, 100, *colour.NewRandomColour())
	mainCanvas.Render()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch

	dom.Log("All done.")
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
}
