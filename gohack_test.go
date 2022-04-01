package gohack

import (
	"github.com/stretchr/testify/assert"
	"testing"
	_ "unsafe"
)

func TestGoid(t *testing.T) {
	runTest(t, func() {
		assert.Equal(t, curGoroutineID(), getg().goid)
	})
}

// curGoroutineID parse the current g's goid from caller stack.
//go:linkname curGoroutineID net/http.http2curGoroutineID
func curGoroutineID() int64
