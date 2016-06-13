package main

import (
	"testing"
)

func TestSet(t *testing.T) {
	bv := NewBitVector(10)
	bv.Set(8)
	if !bv.IsSet(8) {
		t.Error("Expected bit to be set")
	}
}

func TestNotSet(t *testing.T) {
	bv := NewBitVector(10)
	bv.Set(8)
	bv.Clear(8)
	if bv.IsSet(8) {
		t.Error("Expected bit to not be set")
	}
}
