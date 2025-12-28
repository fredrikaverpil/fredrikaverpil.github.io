package examples

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
)

// ExampleGetDigest demonstrates accepting hash.Hash
func ExampleGetDigest() {
	data := []byte("hello world")
	digest := GetDigest(sha256.New(), data)
	fmt.Printf("SHA256 (%d bytes): %x\n", len(digest), digest)
	// Output: SHA256 (32 bytes): b94d27b9934d3e08a52e52d7da7dabfac484efe37a5380ee9088f7ace2efcde9
}

// ExampleSignData demonstrates accepting crypto.Signer
func ExampleSignData() {
	key, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	data := []byte("message to sign")
	sig, _ := SignData(key, data)
	// ECDSA signatures vary in length (68-72 bytes) due to DER encoding
	if len(sig) > 60 && len(sig) < 80 {
		fmt.Println("Signature generated successfully")
	}
	// Output: Signature generated successfully
}
