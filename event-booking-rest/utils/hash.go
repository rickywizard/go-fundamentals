package utils

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type HashSalt struct {
	Hash []byte
	Salt []byte
}

type Argon2id struct {
	time    uint32
	memory  uint32
	threads uint32
	keyLen  uint32
	saltLen uint32
}

func NewArgon2id() *Argon2id {
	return &Argon2id{
		time:    1,
		memory:  64 * 1024,
		threads: 2,
		keyLen:  256,
		saltLen: 32,
	}
}

func randomSecret(length uint32) ([]byte, error) {
	secret := make([]byte, length)

	_, err := rand.Read(secret)
	if err != nil {
		return nil, err
	}

	return secret, nil
}

func (a *Argon2id) Hash(password, salt []byte) (string, error) {
	var err error

	if len(salt) == 0 {
		salt, err = randomSecret(a.saltLen)
	}

	if err != nil {
		return "", err
	}

	hash := argon2.IDKey(password, salt, a.time, a.memory, uint8(a.threads), a.keyLen)

	encodedSalt := base64.StdEncoding.EncodeToString(salt)
	encodedHash := base64.StdEncoding.EncodeToString(hash)

	return fmt.Sprintf("%s:%s", encodedSalt, encodedHash), nil
}

func (a *Argon2id) Compare(encoded string, password []byte) error {
	parts := strings.Split(encoded, ":")
	if len(parts) != 2 {
		return errors.New("invalid stored password format")
	}

	saltBytes, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return errors.New("invalid encoding")
	}

	hashBytes, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return errors.New("invalidencoding")
	}

	computed := argon2.IDKey(password, saltBytes, a.time, a.memory, uint8(a.threads), a.keyLen)

	if !bytes.Equal(hashBytes, computed) {
		return errors.New("invalid credentials")
	}

	return nil
}
