package main

// ----------------------------------------------------------------------------
type Pixel struct {
	r, g, b, a uint8
}

// ----------------------------------------------------------------------------
func NewBlankPixel() Pixel {
	return Pixel{a: 255}
}

// ----------------------------------------------------------------------------
func NewWhitePixel() Pixel {
	return Pixel{r: 255, g: 255, b: 255, a: 255}
}

// ----------------------------------------------------------------------------
// Built using information from https://en.wikipedia.org/wiki/Grayscale
// and https://stackoverflow.com/questions/42516203/converting-rgba-image-to-grayscale-golang
//
// Turns the pixel value in a grayscale pixel
func (p *Pixel) toGrayscale() Pixel {
	asGray := uint8((0.299*(float64(p.r)) + 0.587*(float64(p.g)) + 0.144*(float64(p.b))) / 256)
	return Pixel{r: asGray, g: asGray, b: asGray, a: p.a}
}
