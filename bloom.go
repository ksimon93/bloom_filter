package main

import (
	"fmt"
	"hash/fnv"
)

type BloomFilter struct {
	bv *BitVector
}

// TODO: Add multiple hashes
func NewBloomFilter(len uint) *BloomFilter {
	return &BloomFilter{
		bv: NewBitVector(len),
	}
}

// TODO: Make the bloom filter generic
func (bf *BloomFilter) Add(s string) {
	idx := hash(s) % bf.bv.Length()
	bf.bv.Set(idx)
}

func (bf *BloomFilter) Contains(s string) bool {
	idx := hash(s) % bf.bv.Length()
	return bf.bv.IsSet(idx)
}

func hash(s string) uint {
	h := fnv.New32()
	h.Write([]byte(s))
	return uint(h.Sum32())
}

func main() {
	bf := NewBloomFilter(10)
	bf.Add("Hello")
	contains := bf.Contains("Hello")
	fmt.Println(contains)
}
