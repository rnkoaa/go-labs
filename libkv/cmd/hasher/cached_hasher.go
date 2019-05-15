package main

import (
	"crypto/sha256"
	"hash"
)

// Cached -
type Cached struct {
	hasher hash.Hash
	result [sha256.Size]byte
}

// NewCached -
func NewCached() *Cached {
	return &Cached{
		hasher: sha256.New(),
	}
}

// Sum -
func (c *Cached) Sum(data []byte) []byte {
	c.hasher.Reset()
	c.hasher.Write(data)
	return c.hasher.Sum(c.result[:0])
}
