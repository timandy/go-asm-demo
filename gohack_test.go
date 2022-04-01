package gohack

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"runtime"
	"testing"
	_ "unsafe"
)

func TestG0(t *testing.T) {
	for i := 0; i < 20; i++ {
		g0 := getg0()
		m := reflect.ValueOf(g0).FieldByName("m").Elem()
		curg := m.FieldByName("curg").Elem()
		goid := curg.FieldByName("goid").Int()
		assert.Equal(t, int64(curGoroutineID()), goid)
	}
}

func TestType(t *testing.T) {
	fmt.Println("====GOOS:", runtime.GOOS)
	fmt.Println("====GOARCH:", runtime.GOARCH)
	//
	gt := getgt()
	assert.Equal(t, "g", gt.Name())
	//
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
		tt := getgt()
		switch runtime.GOARCH {
		case "amd64":
			fallthrough
		case "386":
			fallthrough
		case "arm64":
			fallthrough
		case "arm":
			assert.Equal(t, numField, tt.NumField())
			assert.Equal(t, offsetGoid, getFieldOffset(tt, "goid"))
			assert.Equal(t, offsetPaniconfault, getFieldOffset(tt, "paniconfault"))
			assert.Equal(t, offsetLabels, getFieldOffset(tt, "labels"))

		default:
			panic("Not support GOARCH: " + runtime.GOARCH)
		}
	}
}

func getFieldOffset(typ reflect.Type, fieldName string) int {
	field, found := typ.FieldByName(fieldName)
	if !found {
		panic("No field [" + fieldName + "] found in type " + typ.Name())
	}
	return int(field.Offset)
}

// curGoroutineID parse the current g's goid from caller stack.
//go:linkname curGoroutineID net/http.http2curGoroutineID
func curGoroutineID() uint64
