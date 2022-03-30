package main

import (
	"reflect"
	"routine-demo/ddd"
	"unsafe"
	_ "unsafe"
)

type gg struct {
	aa int64
	bb unsafe.Pointer
	cc unsafe.Pointer
	dd unsafe.Pointer
	ee unsafe.Pointer
	ff unsafe.Pointer
}

var g0 gg

//go:nosplit
func read666() interface{} {
	return g0
}

//go:nosplit
func getgt() reflect.Type {
	return reflect.TypeOf(ddd.Getgtype())
}

func main() {
	println(read666())
	for i := 0; i < 100; i++ {
		of := getgt()
		println(of.NumField())
		field, _ := of.FieldByName("goid")
		println(field.Offset)
	}
}
