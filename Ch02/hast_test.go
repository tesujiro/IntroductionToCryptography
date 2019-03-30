package main

import (
	"crypto/sha1"
	"hash"
	"log"
	"testing"

	"github.com/tesujiro/IntroductionToCryptography/Ch02/mySha1"
)

func TestHash(t *testing.T) {
	s1 := "tesu"
	s2 := "jiro"
	for _, h := range []hash.Hash{
		sha1.New(),
		mySha1.New(),
		//sha256.New(),
		//sha512.New(),
	} {
		h.Write([]byte(s1))
		h.Write([]byte(s2))
		bs := h.Sum(nil)
		log.Printf("%T:\t%x\n", h, bs)
	}
}
