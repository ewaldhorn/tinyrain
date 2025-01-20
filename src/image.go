package main

import (
	"fmt"
	"unsafe"

	"github.com/ewaldhorn/dommie/dom"
)

// ----------------------------------------------------------------------------
//
//export passImageDataToWasm
func passImageDataToWasm(length int, data *byte) {
	imageData := make([]byte, length)
	for i := 0; i < length; i++ {
		imageData[i] = *(*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(data)) + uintptr(i)))
	}

	for i := range 25 {
		dom.Log(fmt.Sprintf("%d: %d", i, imageData[i]))
	}
}
