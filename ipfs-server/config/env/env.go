package env

import (
	"fmt"
	"os"
)

func Get(key, def string) string {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	return val
}

func MustGet(key string) string {
	val := os.Getenv(key)
	if val == "" {
		panic("missing required env: " + key)
	}
	return val
}

func GetInt(key string, def int) int {
	val := os.Getenv(key)
	if val == "" {
		return def
	}
	var out int
	fmt.Sscanf(val, "%d", &out)
	return out
}
