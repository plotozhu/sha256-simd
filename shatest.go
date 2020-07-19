//+build !noasm,!appengine

package sha256

import (
	"crypto/sha256"
	"encoding/binary"
	"shalib"
	"testing"
)


func runTestSha(hashfunc func([]byte) [32]byte) bool {
	var m = []byte("This is a message. This is a message. This is a message. This is a message.")

	ar := hashfunc(m)
	br := sha256.Sum256(m)

	return ar == br
}

func TestSha0(t *testing.T) {
	if !runTestSha(Sum256) {
		t.Errorf("FAILED")
	}
}

func TestSha1(t *testing.T) {
	if sha && ssse3 && sse41 && !runTestSha(shalib.Sha256hash) {
		t.Errorf("FAILED")
	}
}
