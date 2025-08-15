// Bloom Filter implementation in Go
package main

import (
	"fmt"  // For printing output
	"hash" // For hash.Hash32 interface
	"time" // For seeding the hash function

	"github.com/spaolacci/murmur3" // MurmurHash3 package
)

// mHasher is a global MurmurHash3 hasher used for hashing keys
var mHasher hash.Hash32

// init initializes the global hasher with a time-based seed
func init() {
	mHasher = murmur3.New32WithSeed(uint32(time.Now().Unix()))
}

// murmurhash hashes the key and maps it to a valid index in the filter
func murmurhash(key string, size int32) int32 {
	mHasher.Write([]byte(key))               // Hash the key
	result := mHasher.Sum32() % uint32(size) // Map hash to filter size
	mHasher.Reset()                          // Reset hasher for next use
	return int32(result)
}

// BloomFilter represents the Bloom filter data structure
type BloomFilter struct {
	filter []bool // The bit array (using bool for simplicity)
	size   int32  // Size of the filter
}

// NewBloomFilter creates and returns a new BloomFilter of given size
func NewBloomFilter(size int32) *BloomFilter {
	return &BloomFilter{
		filter: make([]bool, size),
		size:   size,
	}
}

// Add inserts a key into the Bloom filter
func (b *BloomFilter) Add(key string) {
	idx := murmurhash(key, b.size) // Get index for the key
	b.filter[idx] = true           // Set the bit at that index
}

// Exists checks if a key might be in the Bloom filter
// Returns true if the bit is set, false otherwise
func (b *BloomFilter) Exists(key string) bool {
	idx := murmurhash(key, b.size) // Get index for the key
	return b.filter[idx]           // Return the bit value
}

// main demonstrates usage of the Bloom filter
func main() {
	bloom := NewBloomFilter(16)     // Create a Bloom filter of size 16
	keys := []string{"a", "b", "c"} // Keys to add

	// Add keys to the filter
	for _, key := range keys {
		bloom.Add(key)
	}

	// Check if keys exist in the filter
	for _, key := range keys {
		fmt.Println(key, bloom.Exists(key))
	}
}
