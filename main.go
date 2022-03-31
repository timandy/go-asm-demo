package main

import (
	"reflect"
	"routine-demo/ddd"
	_ "unsafe"
)

//go:nosplit
func getgt() reflect.Type {
	return reflect.TypeOf(ddd.Getgtype())
}

func main() {
	for i := 0; i < 100; i++ {
		of := getgt()
		println(of.NumField())
		field, _ := of.FieldByName("goid")
		println(field.Offset)
	}
}
