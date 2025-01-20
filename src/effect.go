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

const (
	dropletCount = 100_000
)

var droplets []Droplet

// ----------------------------------------------------------------------------
func setupAnimation() {
	width := imageWidth
	if width > canvasWidth {
		width = canvasWidth
	}

	height := imageHeight
	if height > canvasHeight {
		height = canvasHeight
	}

	dropletsMade := 0
	droplets = make([]Droplet, dropletCount)
	for dropletsMade < dropletCount {
		for x := range width {
			for y := range height {
				if rand.Intn(100) > 55 && dropletsMade < dropletCount {
					droplet := Droplet{x: x, y: y, brightness: uint8(150 + rand.Intn(100)), speed: 1 + rand.Intn(15)}
					droplets[dropletsMade] = droplet
					dropletsMade++
				}
			}
		}
	}
	startAnimation()
}

// ----------------------------------------------------------------------------
func updateDroplets() {
	for pos := range droplets {
		droplets[pos].y += droplets[pos].speed

		if droplets[pos].y >= imageHeight {
			droplets[pos].y = rand.Intn(5)
			droplets[pos].x = rand.Intn(canvasWidth)
			droplets[pos].speed = 1 + rand.Intn(15)
			droplets[pos].brightness = uint8(150 + rand.Intn(100))
		}

	}
	renderDroplets()
}

// ----------------------------------------------------------------------------
func renderDroplets() {
	mainCanvas.ClearScreen(*canvasBackgroundColour)
	for _, droplet := range droplets {
		offset := (droplet.x * 4) + (droplet.y * 4 * imageWidth)
		col := colour.NewColour(imageData[offset], imageData[offset+1], imageData[offset+2], droplet.brightness)
		mainCanvas.ColourPutPixel(droplet.x, droplet.y, *col)
	}
	mainCanvas.Render()
}

// ----------------------------------------------------------------------------
func renderOriginal() {
	width := imageWidth
	if width > canvasWidth {
		width = canvasWidth
	}

	height := imageHeight
	if height > canvasHeight {
		height = canvasHeight
	}

	for x := range width {
		for y := range height {
			offset := (x * 4) + (y * 4 * imageWidth)

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
func setRefreshEffectCallback() {
	js.Global().Set("refreshEffect", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		updateDroplets()
		return nil
	}))
}
