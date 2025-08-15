package main

import (
	"fmt"
	"hash"

	"github.com/spaolacci/murmur3"
)

var hasher hash.Hash32 = murmur3.New32WithSeed(32)

func murmurhash(key string, size int) int {
	hasher.Write(data)
	var result = hasher.Sum32()

}

type BloomFilter struct {
	filter []bool
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		make([]bool, size),
	}
}

func (b *BloomFilter) Add(key string) {

}

func (b *BloomFilter) Exists(key string) bool {
	return false

}

func main() {

	bloom := NewBloomFilter(16)
	keys := []string{"a", "b", "c"}
	for _, key := range keys {
		bloom.Add(key)
	}

	for _, key := range keys {
		fmt.Println(key, bloom.Exists(key))
	}

}
