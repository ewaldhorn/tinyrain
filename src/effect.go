package main

import (
	"math"

	"github.com/ewaldhorn/tinycanvas/colour"
)

var brightnessMap []byte

// ----------------------------------------------------------------------------
func calculateBrightnessValues() {
	width := imageWidth
	if width > canvasWidth {
		width = canvasWidth
	}

	height := imageHeight
	if height > canvasHeight {
		height = canvasHeight
	}

	brightnessMap = make([]byte, width*height)

	for x := range width {
		for y := range height {
			offset := (x * 4) + (y * 4 * imageWidth)

			// don't bother if we are outside our area
			if offset < 0 || offset >= len(imageData) {
				continue
			}

			brigthness := calculateBrightness(imageData[offset], imageData[offset+1], imageData[offset+2])
			brightnessMap[x+(y*width)] = brigthness
		}
	}
}

// ----------------------------------------------------------------------------
func calculateBrightness(r, g, b byte) byte {
	tmpR := float64(r) * 0.299
	tmpG := float64(g) * 0.587
	tmpB := float64(b) * 0.144

	return byte(math.Sqrt(tmpR+tmpG+tmpB) / 100)
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
func renderBrightness() {
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
			offset := x + (y * width)

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
