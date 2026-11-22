package crypto

import (
	"crypto/rand"
)

const DEKSize = 32
const NonceSize = 12

func fillRand(buf []byte) error {
	_, err := rand.Read(buf)
	if err != nil {
		return err
	}

	return nil
}

func createDek(buf *[DEKSize]byte) error {
	return fillRand(buf[:])
}

func createNonce(buf *[NonceSize]byte) error {
	return fillRand(buf[:])
}
