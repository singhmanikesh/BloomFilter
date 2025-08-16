// Bloom Filter implementation in Go
package main

import (
	"fmt"  // For printing output
	"hash" // For hash.Hash32 interface

	"github.com/google/uuid"
	"github.com/spaolacci/murmur3" // MurmurHash3 package
)

// mHasher is a global MurmurHash3 hasher used for hashing keys
var hashers [4]hash.Hash32

// init initializes the global hasher with a time-based seed
func init() {
	hashers[0] = murmur3.New32WithSeed(10)
	hashers[1] = murmur3.New32WithSeed(20)
	hashers[2] = murmur3.New32WithSeed(30)
	hashers[3] = murmur3.New32WithSeed(40)
}

// murmurhash hashes the key and maps it to a valid index in the filter
func murmurhashes(key string, size int32) [4]int32 {
	var idxs [4]int32
	for i, hasher := range hashers {
		hasher.Write([]byte(key))
		idxs[i] = int32(hasher.Sum32() % uint32(size))
		hasher.Reset()
	}
	return idxs
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
	idxs := murmurhashes(key, b.size)
	for _, idx := range idxs {
		b.filter[idx] = true
	}
}

func (b *BloomFilter) print() {
	fmt.Println(b.filter)
}

// Exists checks if a key might be in the Bloom filter
// Returns true if the bit is set, false otherwise
func (b *BloomFilter) Exists(key string) (string, [4]int32, bool) {
	idxs := murmurhashes(key, b.size)
	for _, idx := range idxs {
		if !b.filter[idx] {
			return key, idxs, false
		}
	}
	return key, idxs, true
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

	for j := 100; j < 10000; j += 100 {

		bloom := NewBloomFilter(int32(j)) // Create a Bloom filter of size j

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
