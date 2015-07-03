package hoop

import (
	"testing"
)

const replicas = 3

func TestEmpty(t *testing.T) {
	h := New([]string{}, replicas)

	_, err := h.Get("key")

	if err != ErrEmptyHoop {
		t.Errorf("Expected \"%s\" but got \"%s\"", ErrEmptyHoop, err)
	}
}

func TestGet(t *testing.T) {
	nodes := []string{
		"node0",
		"node1",
		"node2",
	}

	h := New(nodes, replicas)

	node0, _ := h.Get("beta")

	if node0 != nodes[0] {
		t.Errorf("Expected \"%s\" but got \"%s\"", nodes[0], node0)
	}

	node1, _ := h.Get("alpha")

	if node1 != nodes[1] {
		t.Errorf("Expected \"%s\" but got \"%s\"", nodes[1], node1)
	}

	node2, _ := h.Get("lambda")

	if node2 != nodes[2] {
		t.Errorf("Expected \"%s\" but got \"%s\"", nodes[2], node2)
	}
}
