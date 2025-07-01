package main

import (
	"math/rand"
	"syscall/js"

	"github.com/ewaldhorn/tinycanvas/colour"
)

// ----------------------------------------------------------------------------
type Droplet struct {
	x, y, speed, iteration int
	brightness             uint8
}

// ----------------------------------------------------------------------------
const (
	dropletCount    = 120_000
	dropletMaxSpeed = 20
)

// ----------------------------------------------------------------------------
var droplets []Droplet
var effectWidth, effectHeight int

// ----------------------------------------------------------------------------
func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// ----------------------------------------------------------------------------
func setupAnimation() {
	// ensure width and height are within the canvas sizes
	effectWidth = minInt(imageWidth, canvasWidth)
	effectHeight = minInt(imageHeight, canvasHeight)

	// now create "random" droplets
	dropletsMade := 0
	droplets = make([]Droplet, dropletCount)

	for dropletsMade < dropletCount {
		x := rand.Intn(effectWidth)
		y := rand.Intn(effectHeight)

		droplet := Droplet{
			x:          x,
			y:          y,
			brightness: uint8(150 + rand.Intn(100)),
			speed:      1 + rand.Intn(dropletMaxSpeed),
		}
		droplets[dropletsMade] = droplet
		dropletsMade++
	}

	renderOriginal()
	startAnimation()
}

// ----------------------------------------------------------------------------
// Move droplets down by their speed, if they reach the bottom, reset them.
func updateDroplets() {
	for pos := range droplets {
		droplets[pos].y += droplets[pos].speed

		if droplets[pos].y >= imageHeight {
			droplets[pos].y = rand.Intn(10)
			droplets[pos].x = rand.Intn(canvasWidth)
			droplets[pos].speed = 1 + rand.Intn(dropletMaxSpeed)
			droplets[pos].brightness = uint8(150 + rand.Intn(100))
		}
	}

	renderDroplets()
}

// ----------------------------------------------------------------------------
func renderDroplets() {
	for _, droplet := range droplets {
		offset := (droplet.x * 4) + (droplet.y * 4 * imageWidth)
		col := colour.NewColour(imageData[offset], imageData[offset+1], imageData[offset+2], droplet.brightness)
		mainCanvas.ColourPutPixel(droplet.x, droplet.y, *col)
	}
	mainCanvas.Render()
}

// ----------------------------------------------------------------------------
// Renders the original image to the screen
func renderOriginal() {
	fourImageWidths := 4 * imageWidth

	for x := range effectWidth {
		for y := range effectHeight {
			offset := (x * 4) + (y * 4 * fourImageWidths)

			// don't bother if we are outside our area
			if offset < 0 || offset >= len(imageData) {
				continue
			}

			col := colour.NewColour(imageData[offset], imageData[offset+1], imageData[offset+2], imageData[offset+3])
			mainCanvas.ColourPutPixel(x, y, *col)
		}
	}
	mainCanvas.Render()
}

// ----------------------------------------------------------------------------
// Allows JS to call into Wasm to refresh the effect.
func setRefreshEffectCallback() {
	js.Global().Set("refreshEffect", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updateDroplets()
		return nil
	}))
}
