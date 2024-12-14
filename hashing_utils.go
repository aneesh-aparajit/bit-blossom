package main

import (
	"hash"
	"time"

	"github.com/spaolacci/murmur3"
)

type HashingFunctions struct {
	fns []hash.Hash32
}

func NewHashingFunctions(numFunctions int) *HashingFunctions {
	hasher := &HashingFunctions{
		fns: make([]hash.Hash32, 0),
	}
	for i := 0; i < numFunctions; i++ {
		hasher.fns = append(hasher.fns, murmur3.New32WithSeed(uint32(time.Now().Unix())))
	}
	return hasher
}

func (h *HashingFunctions) GetHashValuesForKey(key []byte) []int {
	hashes := make([]int, 0)
	for _, fn := range h.fns {
		fn.Write(key)
		hashes = append(hashes, int(fn.Sum32()))
		fn.Reset()
	}
	return hashes
}

func (h *HashingFunctions) Size() int {
	return len(h.fns)
}

func (h *HashingFunctions) Reset() {
	for i := 0; i < h.Size(); i++ {
		h.fns[i].Reset()
	}
}
