package utils

import (
	"crypto/sha256"
	"encoding/base64"
	"strings"
)

func SHAEncoding(target string) (output string) {
	buf := []byte(target)

	encrypted := sha256.New()
	encrypted.Write(buf)

	output = base64.StdEncoding.EncodeToString(encrypted.Sum(nil))

	return
}

type StringArr []string

func (arr StringArr) ContainsIgnoreCase(val string) bool {
	for _, v := range arr {
		if strings.EqualFold(v, val) {
			return true
		}
	}
	return false
}
