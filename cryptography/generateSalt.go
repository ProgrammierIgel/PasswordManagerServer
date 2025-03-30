package cryptography

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"github.com/programmierigel/pwmanager/logger"
)

// GenerateSalt creates a cryptographically secure random string
// suitable for use as a salt in password hashing or other security applications.
// The length parameter specifies the number of random bytes to generate
// before base64 encoding (final output will be longer).
func GenerateSalt(length int) (string, error) {
	// Allocate a byte slice to store the random bytes
	randomBytes := make([]byte, length)

	// Read random bytes from crypto/rand
	// This is a secure random number generator suitable for cryptographic use
	_, err := rand.Read(randomBytes)
	if err != nil {
		logger.Critiacal(fmt.Sprintf("failed to generate random bytes: %s", err.Error()))
		return "", fmt.Errorf("failed to generate random bytes: %w", err)
	}

	// Encode the random bytes to base64 to get a string
	// Using RawURLEncoding to avoid characters that might need escaping in URLs
	return base64.RawURLEncoding.EncodeToString(randomBytes), nil
}
