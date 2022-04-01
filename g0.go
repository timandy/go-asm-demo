package gohack

import (
	"reflect"
	_ "unsafe"
)

// runtime.g0
func getg0() interface{}

//go:nosplit
func getgt() reflect.Type {
	return reflect.TypeOf(getg0())
}
