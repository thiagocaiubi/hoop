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

	for _, node := range nodes {
		for i := 0; i < replicas; i++ {
			hash := hash(fmt.Sprintf("%v:%v", i, node))
			h.nodesHashes = append(h.nodesHashes, hash)
			h.nodesMap[hash] = node
		}
	}

	sort.Ints(h.nodesHashes)

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

func (h *Hoop) isEmpty() bool {
	return len(h.nodesHashes) == 0
}

func hash(data string) int {
	return int(crc32.ChecksumIEEE([]byte(data)))
}
