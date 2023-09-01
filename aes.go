package gconv

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type AES struct {
	Key []byte
}

// Encrypt function to encrypt plaintext using AES.
func (s *AES) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(s.Key)
	if err != nil {
		return "", err
	}

	// Create a new GCM instance.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Generate a random nonce.
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "", err
	}

	// Encrypt the plaintext using GCM.
	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)
	return string(append(nonce, ciphertext...)), nil
}

// Decrypt function to decrypt ciphertext using AES.
func (s *AES) Decrypt(ciphertext string) (string, error) {
	block, err := aes.NewCipher(s.Key)
	if err != nil {
		return "", err
	}

	// Create a new GCM instance.
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Extract the nonce from the ciphertext.
	nonce := ciphertext[:gcm.NonceSize()]
	ciphertext = ciphertext[gcm.NonceSize():]

	// Decrypt the ciphertext using GCM.
	plaintext, err := gcm.Open(nil, []byte(nonce), []byte(ciphertext), nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
