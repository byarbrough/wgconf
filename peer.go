package wgconf

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/curve25519"
)

// GenKey returns a base64 encoded private key
func GenKey() (string, error) {
	scalar := make([]byte, 32)
	_, err := rand.Read(scalar)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return "", err
	}
	privateKey, err := curve25519.X25519(scalar, curve25519.Basepoint)
	if err != nil {
		return "", err
	}
	encodedKey := base64.StdEncoding.EncodeToString(privateKey)
	return encodedKey, nil
}

// PubKey returns a base64 public key for the provided
// private key
func PubKey(privateKey string) (string, error) {
	decodedPrivate, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", err
	}
	publicKey, err := curve25519.X25519(decodedPrivate, curve25519.Basepoint)
	if err != nil {
		return "", err
	}
	encodedPublic := base64.StdEncoding.EncodeToString(publicKey)
	return encodedPublic, nil
}
