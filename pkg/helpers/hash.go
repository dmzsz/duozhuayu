package helpers

import (
	"encoding/hex"

	"golang.org/x/crypto/blake2b"
)

// CalcHash generates a fixed-sized BLAKE2b-256 hash of the given text
func CalcHash(plaintext, keyOptional string) ([]byte, error) {
	blake2b256Hash, err := blake2b.New256([]byte(keyOptional))
	if err != nil {
		return nil, err
	}

	_, err = blake2b256Hash.Write([]byte(plaintext))
	if err != nil {
		return nil, err
	}

	blake2b256Sum := blake2b256Hash.Sum(nil)

	return blake2b256Sum, nil
}

// DecryptEmail returns the plaintext email from the given cipher and nonce
func DecryptEmail(emailNonce, emailCipher string, key string) (email string, err error) {
	nonce, err := hex.DecodeString(emailNonce)
	if err != nil {
		return
	}
	cipherEmail, err := hex.DecodeString(emailCipher)
	if err != nil {
		return
	}

	email, err = DecryptChacha20poly1305(
		[]byte(key),
		nonce,
		cipherEmail,
	)
	return
}
