package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// NOTE: I have copied this code from internet
var (
	_iterations = uint32(1) // Number of iterations
)

const (
	_memory    = uint32(64 * 1024) // Memory size in KB (64 MB here)
	_threads   = uint8(100)        // Number of threads to use
	_keyLength = uint32(64)        // Length of the derived key
	_saltSize  = 24
)

func GenerateRandomSalt(size int) ([]byte, error) {
	salt := make([]byte, size)
	_, err := rand.Read(salt)
	if err != nil {
		return nil, err
	}
	return salt, nil
}

func HashPassword(password string) (string, error) {
	salt, err := GenerateRandomSalt(_saltSize)
	if err != nil {
		return "", err
	}

	// Hash the password
	hash := argon2.IDKey([]byte(password), salt, _iterations, _memory, _threads, _keyLength)

	// Encode the hash and salt as base64
	saltEncoded := base64.RawStdEncoding.EncodeToString(salt)
	hashEncoded := base64.RawStdEncoding.EncodeToString(hash)

	// Return combined hash with salt (you'll need to save both)
	return fmt.Sprintf("%s:%s", saltEncoded, hashEncoded), nil
}

func ComparePasswordMatching(password string, storedHash string) bool {
	// Split the stored hash into its components (salt and hashed password)
	parts := strings.Split(storedHash, ":")
	if len(parts) != 2 {
		return false
	}

	// Decode the stored salt and hashed password
	salt, err := base64.RawStdEncoding.DecodeString(parts[0])
	if err != nil {
		return false
	}

	storedPasswordHash, err := base64.RawStdEncoding.DecodeString(parts[1])
	if err != nil {
		return false
	}

	// Hash the provided password with the stored salt
	hash := argon2.IDKey([]byte(password), salt, _iterations, _memory, _threads, _keyLength)

	// Compare the hashed password with the stored hash
	return string(hash) == string(storedPasswordHash)
}
