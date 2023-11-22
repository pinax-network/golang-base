package keys

const (
	DefaultKeyLength = 16
)

type Generator struct {
	keyLength int
	rand      Rand
	encoder   Encoder
	decoder   Decoder
	signer    Signer
	verifier  Verifier
}

// NewGenerator returns a new key generator.
func NewGenerator(options ...Option) *Generator {

	encoder, decoder := NoopEncoding()
	res := &Generator{
		keyLength: DefaultKeyLength,
		rand:      Secure(),
		encoder:   encoder,
		decoder:   decoder,
		signer:    NoopSigner(),
	}

	for _, o := range options {
		o(res)
	}

	return res
}

// GenerateKey generates a secure key. The length of the key is set from DefaultKeyLength or WithKeyLength. The
// default encoding uses the NoopEncoding, but can be adapted using the WithEncoder option.
func (g *Generator) GenerateKey() ([]byte, error) {

	res, err := g.rand(g.keyLength)
	if err != nil {
		return res, err
	}

	signature, err := g.signer(res)
	if err != nil {
		return res, err
	}

	return g.encoder(append(res, signature...)), nil
}

func (g *Generator) VerifySignature(key []byte) (bool, error) {

	decodedKey, err := g.decoder(key)
	if err != nil {
		return false, err
	}

	return g.verifier(decodedKey[:g.keyLength], decodedKey[g.keyLength:])
}
