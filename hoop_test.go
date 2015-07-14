package hoop

import (
	"fmt"
	"testing"
)

const replicas = 13

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

func TestAdd(t *testing.T) {
	nodes := []string{
		"node0",
	}

	h := New(nodes, replicas)

	node0, _ := h.Get("beta")

	if node0 != nodes[0] {
		t.Errorf("Expected \"%s\" but got \"%s\"", nodes[0], node0)
	}

	h.Add("node1")
	h.Add("node2")

	node0, _ = h.Get("beta")

	if node0 != nodes[0] {
		t.Errorf("Expected \"%s\" but got \"%s\"", nodes[0], node0)
	}

	node1, _ := h.Get("alpha")

	if node1 != "node1" {
		t.Errorf("Expected \"%s\" but got \"%s\"", "node1", node1)
	}

	node2, _ := h.Get("lambda")

	if node2 != "node2" {
		t.Errorf("Expected \"%s\" but got \"%s\"", "node2", node2)
	}
}

func TestRemove(t *testing.T) {
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

	h.Remove("node0")

	node1, _ := h.Get("beta")

	if node1 != nodes[1] {
		t.Errorf("Expected \"%s\" but got \"%s\"", nodes[1], node1)
	}

	node1, _ = h.Get("alpha")

	if node1 != nodes[1] {
		t.Errorf("Expected \"%s\" but got \"%s\"", nodes[1], node1)
	}

	node2, _ := h.Get("lambda")

	if node2 != nodes[2] {
		t.Errorf("Expected \"%s\" but got \"%s\"", nodes[2], node2)
	}
}

func BenchmarkGet(b *testing.B) {
	nodes := []string{
		"node0",
		"node1",
		"node2",
	}

	results := make(map[string]int, len(nodes))

	h := New(nodes, replicas)
	for i := 0; i < b.N; i++ {
		node, _ := h.Get(fmt.Sprintf("key:%d", i))
		results[node] += 1
	}

	fmt.Println(results)
}
