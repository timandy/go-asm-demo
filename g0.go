package main

import (
	"reflect"
	_ "unsafe"
)

func getg0() interface{}

//go:nosplit
func getgt() reflect.Type {
	return reflect.TypeOf(getg0())
}
