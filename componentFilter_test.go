package ecs

import (
	"testing"

	bitset "github.com/bits-and-blooms/bitset"
)

func TestBitset(t *testing.T) {

	b := bitset.New(32)

	b.Set(0)
	bytes := b.Bytes()
	t.Log("bytes =", bytes)
	if bytes[0] != 1 {

		t.Fatal("Check Fail1")
	}
}
