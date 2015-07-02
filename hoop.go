package hoop

import (
	"errors"
)

var ErrEmptyHoop = errors.New("hoop is empty")

type Hoop struct {
	nodes []string
}

func New(nodes []string) *Hoop {
	return &Hoop{nodes}
}

func (h *Hoop) Get(key string) (string, error) {
	if h.isEmpty() {
		return "", ErrEmptyHoop
	}

	return "", nil
}

func (h *Hoop) isEmpty() bool {
	return len(h.nodes) == 0
}
