package main

import (
	"testing"
	"time"
)

func TestOr(t *testing.T) {
	t.Run("AllChannelsClosed", func(t *testing.T) {
		c1 := make(chan interface{})
		c2 := make(chan interface{})
		defer close(c1)
		defer close(c2)

		done := or(c1, c2)

		c1 <- struct{}{}
		select {
		case _, ok := <-done:
			if !ok {
				t.Error("Expected channel not to be closed")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Expected value on channel")
		}

		c2 <- struct{}{}
		select {
		case _, ok := <-done:
			if !ok {
				t.Error("Expected channel not to be closed")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Expected value on channel")
		}
	})

	t.Run("SomeChannelsClosed", func(t *testing.T) {
		c1 := make(chan interface{})
		c2 := make(chan interface{})

		done := or(c1, c2)

		close(c1)

		select {
		case _, ok := <-done:
			if !ok {
				t.Error("Expected channel not to be closed")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Expected value on channel")
		}

		close(c2)

		select {
		case _, ok := <-done:
			if !ok {
				t.Error("Expected channel not to be closed")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Expected value on channel")
		}
	})

	t.Run("NoChannels", func(t *testing.T) {
		done := or()
		select {
		case _, ok := <-done:
			if ok {
				t.Error("Expected channel to be closed")
			}
		case <-time.After(100 * time.Millisecond):
			t.Error("Expected channel to be closed")
		}
	})
}
