// Bloom Filter implementation in Go
package main

import (
	"fmt"  // For printing output
	"hash" // For hash.Hash32 interface

	"github.com/google/uuid"
	"github.com/spaolacci/murmur3" // MurmurHash3 package
)

// mHasher is a global MurmurHash3 hasher used for hashing keys
var mHasher hash.Hash32

// init initializes the global hasher with a time-based seed
func init() {
	mHasher = murmur3.New32WithSeed(uint32(10))
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
	//fmt.Println("wrote", key, "at index", idx) // Print the index where the key was added
}

func (b *BloomFilter) print() {
	fmt.Println(b.filter)
}

// Exists checks if a key might be in the Bloom filter
// Returns true if the bit is set, false otherwise
func (b *BloomFilter) Exists(key string) (string, int32, bool) {
	idx := murmurhash(key, b.size) // Get index for the key
	return key, idx, b.filter[idx]
	// Return the bit value
}

// main demonstrates usage of the Bloom filter
func main() {
	dataset := make([]string, 0) // Initialize dataset to hold UUIDs
	dataset_exists := make(map[string]bool)
	dataset_notexists := make(map[string]bool)
	for i := 0; i < 500; i++ {
		u := uuid.New()
		dataset = append(dataset, u.String())
		dataset_exists[u.String()] = true
	}

	for i := 0; i < 500; i++ {
		u := uuid.New()
		dataset = append(dataset, u.String())
		dataset_notexists[u.String()] = false
	}

	bloom := NewBloomFilter(1000)

	for key, _ := range dataset_exists {
		bloom.Add(key)

	}

	falsePositive := 0
	for _, key := range dataset {
		_, _, exists := bloom.Exists(key)
		if exists {
			if _, ok := dataset_notexists[key]; ok {
				falsePositive++
			}
		}
	}
	fmt.Println(float64(falsePositive) / float64(len(dataset)))

}

// Create a Bloom filter of size 1000
//	keys := []string{"a", "b", "c", "d", "e", "f", "g"} // Keys to add
// Add keys to the filter
//for _, key := range keys {
//	bloom.Add(key)
//	}
// Check if keys exist in the filter
//	for _, key := range keys {
//		fmt.Println(bloom.Exists(key))
//	}
//	fmt.Println(bloom.Exists("h"))
//	bloom.print()
