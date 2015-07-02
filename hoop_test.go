package hoop

import (
	"testing"
)

func TestEmpty(t *testing.T) {
	h := New([]string{})

	_, err := h.Get("key")

	if err != ErrEmptyHoop {
		t.Errorf("Expected \"%s\" but got \"%s\"", ErrEmptyHoop, err)
	}
}
