package env

import (
	"fmt"
	"os"
)

// Get returns the value of the environment variable
// or the default value if it is empty.
func Get(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}

// MustGet returns the value of the environment variable
// if it is not empty, or panics otherwise.
func MustGet(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("env not set: %s", key))
	}
	return v
}
