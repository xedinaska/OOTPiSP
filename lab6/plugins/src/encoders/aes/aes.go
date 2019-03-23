package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"io"

	"github.com/sirupsen/logrus"
)

var Exported AES

type AES struct{}

func (s *AES) Encode(value interface{}) string {
	block, _ := aes.NewCipher([]byte(createHash("")))

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logrus.Error(err.Error())
		return value.(string)
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		logrus.Error(err.Error())
		return value.(string)
	}

	return string(gcm.Seal(nonce, nonce, []byte(value.(string)), nil))
}

func (s *AES) Decode(hash string) interface{} {
	key := []byte(createHash(""))

	block, err := aes.NewCipher(key)
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}

	data := []byte(hash)
	nonceSize := gcm.NonceSize()
	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}

	return string(plaintext)
}

func (s *AES) Name() string {
	return "aes"
}

func createHash(key string) string {
	hasher := md5.New()
	hasher.Write([]byte(key))
	return hex.EncodeToString(hasher.Sum(nil))
}
