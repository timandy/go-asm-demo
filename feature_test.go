package routine

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestComplete(t *testing.T) {
	fea := NewFeature()
	go func() {
		fea.Complete(1)
	}()
	assert.Equal(t, 1, fea.Get())
}

func TestCompleteError(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			se := err.(StackError)
			assert.NotNil(t, se)
			assert.Equal(t, 1, se.Message())
			assert.NotNil(t, se.StackTrace())
		}
	}()

	fea := NewFeature()
	go func() {
		fea.CompleteError(1)
	}()
	fea.Get()
	assert.Fail(t, "should not be here")
}

func TestGet(t *testing.T) {
	run := false
	fea := NewFeature()
	go func() {
		time.Sleep(time.Millisecond * 500)
		run = true
		fea.Complete(nil)
	}()
	assert.Nil(t, fea.Get())
	assert.True(t, run)
}
