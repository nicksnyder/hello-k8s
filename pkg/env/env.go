package env

import "os"

// Get returns the value of the environment variable
// or the default value if it is empty.
func Get(key, defaultValue string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return defaultValue
}
