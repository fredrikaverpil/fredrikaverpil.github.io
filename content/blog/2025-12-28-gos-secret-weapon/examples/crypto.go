package examples

import (
	"crypto"
	"crypto/rand"
	"crypto/sha256"
	"hash"
)

// ============================================================================
// ACCEPTING INTERFACES
// ============================================================================

// GetDigest computes a digest with any hash algorithm using the hash.Hash interface.
func GetDigest(h hash.Hash, data []byte) []byte {
	h.Write(data)
	return h.Sum(nil)
}

// SignData signs data with any crypto.Signer using RSA, ECDSA, or custom implementations.
func SignData(signer crypto.Signer, data []byte) ([]byte, error) {
	hash := sha256.Sum256(data)
	return signer.Sign(rand.Reader, hash[:], crypto.SHA256)
}
