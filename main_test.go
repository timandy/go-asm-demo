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
	for i := 0; i < 20; i++ {
		gt := getgt()
		assert.Equal(t, 51, gt.NumField())
		switch runtime.GOARCH {
		case "amd64":
			assert.Equal(t, 152, getFieldOffset(gt, "goid"))
			assert.Equal(t, 181, getFieldOffset(gt, "paniconfault"))
			assert.Equal(t, 360, getFieldOffset(gt, "labels"))
		case "386":
			assert.Equal(t, 80, getFieldOffset(gt, "goid"))
			assert.Equal(t, 105, getFieldOffset(gt, "paniconfault"))
			assert.Equal(t, 216, getFieldOffset(gt, "labels"))
		case "arm64":
			assert.Equal(t, 152, getFieldOffset(gt, "goid"))
			assert.Equal(t, 181, getFieldOffset(gt, "paniconfault"))
			assert.Equal(t, 360, getFieldOffset(gt, "labels"))
		case "arm":
			assert.Equal(t, 80, getFieldOffset(gt, "goid"))
			assert.Equal(t, 105, getFieldOffset(gt, "paniconfault"))
			assert.Equal(t, 216, getFieldOffset(gt, "labels"))

		default:
			panic("Not support GOARCH: " + runtime.GOARCH)
		}

	}
}

func getFieldOffset(typ reflect.Type, fieldName string) int {
	field, _ := typ.FieldByName(fieldName)
	return int(field.Offset)
}
