package keys

import "crypto/rand"

type Rand func(int) ([]byte, error)

func Secure() Rand {
	return func(length int) ([]byte, error) {
		res := make([]byte, length)
		if _, err := rand.Read(res); err != nil {
			return res, err
		}
		return res, nil
	}
}
