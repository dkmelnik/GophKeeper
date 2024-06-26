package utils

import (
	"log"
	"strings"
)

func StringToMap(str string) map[string]string {
	out := make(map[string]string)
	if str != "" {
		pairs := strings.Split(str, ",")
		for _, pair := range pairs {
			kv := strings.Split(pair, ":")
			if len(kv) == 2 {
				out[kv[0]] = kv[1]
			} else {
				log.Fatalf("Invalid format: %s", pair)
			}
		}
	}
	return out
}
