package goutils

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"
)

func MakeUniqueId() (string, error) {
	numBytes := 16
	bytes := make([]byte, numBytes)
	numBytes, err := io.ReadFull(rand.Reader, bytes)
	if numBytes != len(bytes) || err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func PrepHtmlForJson(target string, slashes bool) string {
	// only escape backslashes if slashes bool is true
	if slashes {
		target = strings.Replace(target, "\"", "\\\"", -1)
	}

	// remove breaks
	target = strings.Replace(target, "\n", "", -1)

	// remove tabs and double spaces
	target = strings.Replace(target, "	", "", -1)
	target = strings.Replace(target, "  ", "", -1)

	return target
}
