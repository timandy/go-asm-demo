package gohack

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
	"unsafe"
)

func TestGoid(t *testing.T) {
	runTest(t, func() {
		gp := getg()
		runtime.GC()
		assert.Equal(t, curGoroutineID(), gp.goid)
	})
}

func TestPaniconfault(t *testing.T) {
	runTest(t, func() {
		gp := getg()
		runtime.GC()
		//read-1
		assert.False(t, setPanicOnFault(false))
		assert.False(t, gp.getPanicOnFault())
		//read-2
		setPanicOnFault(true)
		assert.True(t, gp.getPanicOnFault())
		//write-1
		gp.setPanicOnFault(false)
		assert.False(t, setPanicOnFault(false))
		//write-2
		gp.setPanicOnFault(true)
		assert.True(t, setPanicOnFault(true))
		//write-read-1
		gp.setPanicOnFault(false)
		assert.False(t, gp.getPanicOnFault())
		//write-read-2
		gp.setPanicOnFault(true)
		assert.True(t, gp.getPanicOnFault())
	})
}

func TestProfLabel(t *testing.T) {
	runTest(t, func() {
		ptr := unsafe.Pointer(&struct{}{})
		null := unsafe.Pointer(uintptr(0))
		assert.NotEqual(t, ptr, null)
		//
		gp := getg()
		runtime.GC()
		//read-1
		assert.Equal(t, null, getProfLabel())
		assert.Equal(t, null, gp.getLabels())
		//read-2
		setProfLabel(ptr)
		assert.Equal(t, ptr, gp.getLabels())
		//write-1
		gp.setLabels(nil)
		assert.Equal(t, null, getProfLabel())
		//write-2
		gp.setLabels(ptr)
		assert.Equal(t, ptr, getProfLabel())
		//write-read-1
		gp.setLabels(nil)
		assert.Equal(t, null, gp.getLabels())
		//write-read-2
		gp.setLabels(ptr)
		assert.Equal(t, ptr, gp.getLabels())
	})
}

// curGoroutineID parse the current g's goid from caller stack.
//go:linkname curGoroutineID net/http.http2curGoroutineID
func curGoroutineID() int64

//go:linkname setPanicOnFault runtime/debug.setPanicOnFault
func setPanicOnFault(new bool) (old bool)

// getProfLabel get current g's labels which will be inherited by new goroutine.
//go:linkname getProfLabel runtime/pprof.runtime_getProfLabel
func getProfLabel() unsafe.Pointer

// setProfLabel set current g's labels which will be inherited by new goroutine.
//go:linkname setProfLabel runtime/pprof.runtime_setProfLabel
func setProfLabel(labels unsafe.Pointer)
