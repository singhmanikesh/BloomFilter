package main

import "fmt"

type BloomFilter struct {
	filter []byte
}

func NewBloomFilter(size int) *BloomFilter {
	return &BloomFilter{
		make([]byte, size),
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
