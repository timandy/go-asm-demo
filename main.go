package main

import (
	_ "unsafe"
)

func main() {
	for i := 0; i < 100; i++ {
		of := getgt()
		println(of.NumField())
		field, _ := of.FieldByName("goid")
		println(field.Offset)
	}
}
