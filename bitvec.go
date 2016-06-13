// Variable sized bit vector that supports setting and unsetting operations.
// Implemented by treating a slice of bytes as a logically contiguous array of bits.

// Some notes about the bitwise operators used since their use may not be inuitive and
// commenting their use inline would make for a lot of repeated comments.
// 1) vec.len<<3 == vec.len*8, for getting total # of bits
// 2) idx>>3 == idx/8, converts bit position in bitstring to equivalent byte in bytestring
// 3) idx&7 == idx%8, converts our idx to the appropiate position to check wiithin the appropiate byte

package main

type BitVector struct {
	data []byte
	len  uint
}

func NewBitVector(len uint) *BitVector {
	return &BitVector{
		data: make([]byte, 4),
		len:  len,
	}
}

func (vec *BitVector) Length() uint {
	return vec.len
}

func (vec *BitVector) Set(idx uint) {
	if idx >= vec.len<<3 {
		panic("Invalid bit index")
	}

	byte := idx >> 8
	mask := uint8(1 << (idx & 7))
	vec.data[byte] |= mask
}

func (vec *BitVector) IsSet(idx uint) bool {
	if idx >= vec.len<<3 {
		panic("Invalid bit index")
	}

	byte := idx >> 3
	mask := uint8(1 << (idx & 7))
	return (vec.data[byte] & mask) > 0
}

func (vec *BitVector) Toggle(idx uint) {
	if idx >= vec.len<<3 {
		panic("Invalid bit index")
	}

	byte := idx >> 3
	mask := uint8(1 << (idx & 7))
	vec.data[byte] ^= mask

}

func (vec *BitVector) Clear(idx uint) {
	if idx >= vec.len<<3 {
		panic("Invalid bit index")
	}

	byte := idx >> 3
	mask := uint8(1 << (idx & 7))
	vec.data[byte] &^= mask
}
