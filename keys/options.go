package keys

type Option func(*Generator)

// WithKeyLength overrides the DefaultKeyLength used to define the byte length of a generated key.
// Note this method will panic if it receives a keyLength of zero or smaller.
func WithKeyLength(keyLength int) Option {

	if keyLength <= 0 {
		panic("default key length must be greater than zero")
	}
	return func(g *Generator) {
		g.keyLength = keyLength
	}
}

func WithRand(rand Rand) Option {
	return func(g *Generator) {
		g.rand = rand
	}
}

func WithEncoding(encoder Encoder, decoder Decoder) Option {
	return func(g *Generator) {
		g.encoder = encoder
		g.decoder = decoder
	}
}

func WithSigner(signer Signer) Option {
	return func(g *Generator) {
		g.signer = signer
	}
}

func WithVerifier(verifier Verifier) Option {
	return func(g *Generator) {
		g.verifier = verifier
	}
}
