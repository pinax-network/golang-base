package keys

import (
	"crypto"
	"crypto/ed25519"
	"crypto/hmac"
	"crypto/rand"
	"hash"
)

type Signer func(src []byte) ([]byte, error)

func NoopSigner() Signer {
	return func(_ []byte) ([]byte, error) {
		return []byte{}, nil
	}
}

func HmacSigner(h func() hash.Hash, signingKey []byte) Signer {
	return func(src []byte) ([]byte, error) {
		hmacSigner := hmac.New(h, signingKey)
		hmacSigner.Write(src)
		return hmacSigner.Sum(nil), nil
	}
}

func Ed25519Signer(privateKey ed25519.PrivateKey) Signer {
	return func(src []byte) ([]byte, error) {
		return privateKey.Sign(rand.Reader, src, crypto.Hash(0))
	}
}
