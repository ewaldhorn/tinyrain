package main

import "math/rand"

// ----------------------------------------------------------------------------
type Particle struct {
	x, y, velocity, speed float64
	size                  float32
	maxHeight             uint32
}

// ----------------------------------------------------------------------------
func NewParticle(width, height uint32) *Particle {
	p := Particle{
		x:         float64(rand.Intn(int(width))),
		y:         0.0,
		speed:     0.0,
		velocity:  0.5 + rand.Float64()*2.5,
		size:      0.5 + rand.Float32()*2,
		maxHeight: height,
	}

	return &p
}

// ----------------------------------------------------------------------------
func (p *Particle) update() {

}
