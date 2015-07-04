package hoop

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sort"
)

var ErrEmptyHoop = errors.New("hoop is empty")

type Hoop struct {
	replicas    int
	nodesHashes []int
	nodesMap    map[int]string
}

func New(nodes []string, replicas int) *Hoop {
	h := &Hoop{
		replicas: replicas,
		nodesMap: make(map[int]string),
	}

	h.Add(nodes...)

	return h
}

func (h *Hoop) Get(key string) (string, error) {
	if h.isEmpty() {
		return "", ErrEmptyHoop
	}

	hash := hash(key)

	pos := sort.Search(len(h.nodesHashes), func(i int) bool {
		return h.nodesHashes[i] >= hash
	})

	if pos >= len(h.nodesHashes) {
		pos = 0
	}

	return h.nodesMap[h.nodesHashes[pos]], nil
}

func (h *Hoop) Add(nodes ...string) {
	for replica := 0; replica < h.replicas; replica++ {
		for _, node := range nodes {
			hash := hash(fmt.Sprintf("%v:%v", replica, node))
			h.nodesHashes = append(h.nodesHashes, hash)
			h.nodesMap[hash] = node
		}
	}

	sort.Ints(h.nodesHashes)
}

func (h *Hoop) Remove(nodes ...string) {
	toDelete := make(map[int]bool)
	aux := []int{}

	for replica := 0; replica < h.replicas; replica++ {
		for _, node := range nodes {
			hash := hash(fmt.Sprintf("%v:%v", replica, node))
			toDelete[hash] = true
			delete(h.nodesMap, hash)

		}
	}

	for _, nodeHash := range h.nodesHashes {
		if toDelete[nodeHash] {
			continue
		}

		aux = append(aux, nodeHash)
	}

	h.nodesHashes = aux
}

func (h *Hoop) isEmpty() bool {
	return len(h.nodesHashes) == 0
}

func hash(data string) int {
	return int(crc32.ChecksumIEEE([]byte(data)))
}
