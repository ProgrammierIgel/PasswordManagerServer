package cryptography

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/programmierigel/pwmanager/logger"
)

// Encrypt encrypts a plaintext string using a password
// Returns base64-encoded encrypted string
func Encrypt(plaintext, password string) (string, error) {
	// Create a hash of the password to use as the AES key
	key := sha256.Sum256([]byte(password))

	// Create AES cipher
	block, err := aes.NewCipher(key[:])
	if err != nil {
		logger.Critiacal(fmt.Sprintf("EncryptionAES: %s",err.Error()))
		return "", err
	}

	// Create a new GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logger.Critiacal(fmt.Sprintf("EncryptionAES: %s",err.Error()))
		return "", err
	}

	// Generate a nonce (Number used ONCE)
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		logger.Critiacal(fmt.Sprintf("EncryptionAES: %s",err.Error()))
		return "", err
	}

	// Encrypt the plaintext
	ciphertext := gcm.Seal(nonce, nonce, []byte(plaintext), nil)

	// Encode to base64 and return
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
