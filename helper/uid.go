package helper

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

func MustGenerateNewUid() uint64 {
	res, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(fmt.Sprintf("failed to generate random uid: %e", err))
	}

	return res.Uint64()
}
