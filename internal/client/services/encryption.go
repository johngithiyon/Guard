package services

import (
	"crypto/rand"

	"golang.org/x/crypto/chacha20"
)

func Encrypt(key, plaintext []byte) ([]byte, []byte, error) {

	nonce := make([]byte, chacha20.NonceSize)

	_, err := rand.Read(nonce)
	if err != nil {
		return nil, nil, err
	}

	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, nil, err
	}

	ciphertext := make([]byte, len(plaintext))

	cipher.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nonce, nil
}

func Decrypt(key, nonce, ciphertext []byte) ([]byte, error) {

	cipher, err := chacha20.NewUnauthenticatedCipher(key, nonce)
	if err != nil {
		return nil, err
	}

	plaintext := make([]byte, len(ciphertext))

	cipher.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}