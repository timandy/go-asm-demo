package g

import (
	"reflect"
	"unsafe"
)

// getgp returns the pointer to the current runtime.g.
func getgp() unsafe.Pointer

// getg0 returns the value of runtime.g0.
func getg0() interface{}

// getgt returns the type of runtime.g.
//go:nosplit
func getgt() reflect.Type {
	return reflect.TypeOf(getg0())
}
