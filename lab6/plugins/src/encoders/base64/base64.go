package main

import (
	"encoding/base64"

	"github.com/sirupsen/logrus"
)

var Exported Base64

type Base64 struct{}

func (s *Base64) Encode(value interface{}) string {
	data := []byte(value.(string))
	return base64.StdEncoding.EncodeToString(data)
}

func (s *Base64) Decode(hash string) interface{} {
	data, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		logrus.Error(err.Error())
		return ""
	}
	return string(data)
}

func (s *Base64) Name() string {
	return "base64"
}
