package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println("====GOOS:", runtime.GOOS)
	fmt.Println("====GOARCH:", runtime.GOARCH)
	//
	gt := getgt()
	numField := gt.NumField()
	offsetGoid := getFieldOffset(gt, "goid")
	offsetPaniconfault := getFieldOffset(gt, "paniconfault")
	offsetLabels := getFieldOffset(gt, "labels")
	//
	fmt.Println("====numField:", numField)
	fmt.Println("====offsetGoid:", offsetGoid)
	fmt.Println("====offsetPaniconfault:", offsetPaniconfault)
	fmt.Println("====offsetLabels:", offsetLabels)
	//
	assert.Greater(t, numField, 20)
	assert.Greater(t, offsetGoid, 0)
	assert.Greater(t, offsetPaniconfault, 0)
	assert.Greater(t, offsetLabels, 0)

	for i := 0; i < 20; i++ {
		gt := getgt()
		switch runtime.GOARCH {
		case "amd64":
		case "386":
		case "arm64":
		case "arm":
			assert.Equal(t, numField, gt.NumField())
			assert.Equal(t, offsetGoid, getFieldOffset(gt, "goid"))
			assert.Equal(t, offsetPaniconfault, getFieldOffset(gt, "paniconfault"))
			assert.Equal(t, offsetLabels, getFieldOffset(gt, "labels"))

		default:
			panic("Not support GOARCH: " + runtime.GOARCH)
		}

	}
}

func getFieldOffset(typ reflect.Type, fieldName string) int {
	field, _ := typ.FieldByName(fieldName)
	return int(field.Offset)
}
