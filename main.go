package main

import (
	"fmt"
	"hash"
	"time"

	"github.com/spaolacci/murmur3"
)

var mHasher hash.Hash32

func init() {
	mHasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

func murmurhash(key string, size int32) int32 {
	mHasher.Write([]byte(key))
	result := mHasher.Sum32() % uint32(size)
	mHasher.Reset()
	return int32(result)

}

type BloomFilter struct {
	filter []bool
	size   int32
}

func NewBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		filter: make([]bool, size),
		size:   size,
	}
}

func (b *BloomFilter) Add(key string) {
	idx := murmurhash(key, b.size)
	b.filter[idx] = true

}

func (b *BloomFilter) Exists(key string) bool {
	idx := murmurhash(key, b.size)
	return b.filter[idx]

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
