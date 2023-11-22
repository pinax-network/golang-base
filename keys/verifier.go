package keys

import (
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/subtle"
	"hash"
)

type Verifier func(key []byte, signature []byte) (bool, error)

func NoopVerifier() Verifier {
	return func(_ []byte, _ []byte) (bool, error) {
		return true, nil
	}
}

func HmacVerifier(h func() hash.Hash, signingKey []byte) Verifier {
	return func(key []byte, signature []byte) (bool, error) {
		hmacSigner := hmac.New(h, signingKey)
		hmacSigner.Write(key)
		resSig := hmacSigner.Sum(nil)

		return subtle.ConstantTimeCompare(resSig, signature) == 1, nil
	}
}

func Ed25519Verifier(publicKey ed25519.PublicKey) Verifier {
	return func(key []byte, signature []byte) (bool, error) {
		return ed25519.Verify(publicKey, key, signature), nil
	}
}
