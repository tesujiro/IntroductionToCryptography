package mySha1

import "hash"

// The size of a SHA-1 checksum in bytes.
const Size = 20

// The blocksize of SHA-1 in bytes.
const BlockSize = 64

const chunk = 512

const (
	H0 = 0x67452301
	H1 = 0xEFCDAB89
	H2 = 0x98BADCFE
	H3 = 0x10325476
	H4 = 0xC3D2E1F0
)

type digest struct {
	h   [5]int32
	x   [chunk / 8]byte
	nx  int
	len uint64
}

func New() hash.Hash {

}

func (h *digest) Write(p []byte) (n int, err error) {
	nn := len(p)
	h.len += uint64(nn)
}

func (h *digest) Sum(b []byte) []byte {

}

func (h *digest) checksum(b []byte) []byte {
}

func (h *digest) Reset() {
}

func (h *digest) Size() int {
	return Size
}

func (h *digest) BlockSize() int {
	return BlockSize
}
