package hoop_test

import (
	"fmt"

	"../hoop"
)

func ExampleHoop_Get() {
	h := hoop.New([]string{
		"node0",
		"node1",
		"node2",
	}, 3)

	node, err := h.Get("alpha")

	if err != nil {
		panic(err)
	}

	fmt.Println(node)

	//Output: node1
}
