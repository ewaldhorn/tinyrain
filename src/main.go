package main

import "tinyrain/src/dom"

func main() {
	dom.Log("Starting...")
	dom.Hide("loading")

	// prevent the app for closing - it stays running for the life of the webpage
	ch := make(chan struct{})
	<-ch

	dom.Log("All done.")
}
