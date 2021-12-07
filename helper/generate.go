package helper

import (
	"crypto/rand"
	"fmt"
	"math"
	"math/big"
)

type Pool string

const (
	UppercaseLetterPool Pool = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowercaseLetterPool Pool = "abcdefghijklmnopqrstuvwxyz"
	NumberPool          Pool = "0123456789"
	AlphanumericPool    Pool = UppercaseLetterPool + LowercaseLetterPool + NumberPool
)

func GenerateRandomString(length int, pools ...Pool) (string, error) {

	if length <= 0 {
		return "", fmt.Errorf("invalid length given (%q), must be > 0", length)
	}

	if len(pools) < 1 {
		return "", fmt.Errorf("no pools given to generate random strings from")
	}

	source := getSourceFromPools(pools)

	res := make([]byte, length)
	for i := 0; i < length; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(source))))
		if err != nil {
			return "", err
		}
		res[i] = source[num.Int64()]
	}

	return string(res), nil
}

func MustGenerateNewUid() uint64 {
	res, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt64))
	if err != nil {
		panic(fmt.Sprintf("failed to generate random uid: %e", err))
	}

	return res.Uint64()
}

func getSourceFromPools(pools []Pool) string {

	res := ""
	for _, pool := range pools {
		res += string(pool)
	}

	return res
}
