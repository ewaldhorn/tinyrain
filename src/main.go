package main

import "github.com/ewaldhorn/dommie/dom"

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

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch

	dom.Log("All done.")
}

// ----------------------------------------------------------------------------
func setCallbacks() {
	setVersionCallback()
}
