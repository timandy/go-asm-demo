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
		assert.Greater(t, gt.NumField(), 20)
		switch runtime.GOARCH {
		case "amd64":
			assert.Greater(t, getFieldOffset(gt, "goid"), 0)
			assert.Greater(t, getFieldOffset(gt, "paniconfault"), 0)
			assert.Greater(t, getFieldOffset(gt, "labels"), 0)
		case "386":
			assert.Greater(t, getFieldOffset(gt, "goid"), 0)
			assert.Greater(t, getFieldOffset(gt, "paniconfault"), 0)
			assert.Greater(t, getFieldOffset(gt, "labels"), 0)
		case "arm64":
			assert.Greater(t, getFieldOffset(gt, "goid"), 0)
			assert.Greater(t, getFieldOffset(gt, "paniconfault"), 0)
			assert.Greater(t, getFieldOffset(gt, "labels"), 0)
		case "arm":
			assert.Greater(t, getFieldOffset(gt, "goid"), 0)
			assert.Greater(t, getFieldOffset(gt, "paniconfault"), 0)
			assert.Greater(t, getFieldOffset(gt, "labels"), 0)

		default:
			panic("Not support GOARCH: " + runtime.GOARCH)
		}

	}
}

func getFieldOffset(typ reflect.Type, fieldName string) int {
	field, _ := typ.FieldByName(fieldName)
	return int(field.Offset)
}
