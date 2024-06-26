// Package utils provides utility functions for various tasks.
package utils

import (
	"crypto/rand"
	"fmt"
	"strings"
)

// GenerateGUID generates a random GUID (UUID version 4) and returns it as a string.
func GenerateGUID() string {
	b := make([]byte, 16)
	_, _ = rand.Read(b)
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x",
		b[0:4], b[4:6], b[6:8], b[8:10], b[10:])

	return uuid
}

// MapToString converts a map[string]string to a string representation of key-value pairs.
// It returns a comma-separated string of key:value pairs.
func MapToString(m map[string]string) string {
	var strPairs []string
	for k, v := range m {
		strPairs = append(strPairs, fmt.Sprintf("%s:%s", k, v))
	}
	return strings.Join(strPairs, ",")
}
