package main

import (
	"github.com/ewaldhorn/dommie/dom"
	"github.com/ewaldhorn/tinycanvas/colour"
	"github.com/ewaldhorn/tinycanvas/tinycanvas"
)

const (
	canvasWidth  = 900
	canvasHeight = 500
)

var mainCanvas *tinycanvas.TinyCanvas
var canvasBackgroundColour *colour.Colour

// ----------------------------------------------------------------------------
// bootstrap is a JavaScript-side defined function, called by Wasm in the main
// Go function
//
//export bootstrap
func bootstrap()

//export loadimage
func loadimage()

//export startAnimation
func startAnimation()

// ----------------------------------------------------------------------------
func main() {
	dom.Log("Starting...")

	setCallbacks()
	dom.Hide("loading")
	bootstrap()
	canvasBackgroundColour = colour.NewColourBlack()
	mainCanvas = tinycanvas.NewTinyCanvas(canvasWidth, canvasHeight)
	mainCanvas.ClearScreen(*canvasBackgroundColour)
	mainCanvas.Render()

	loadimage()

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch

	dom.Log("All done.")
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
	setLoadImageCallback()
	setRefreshEffectCallback()
}
