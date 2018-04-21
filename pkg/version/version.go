package version

import "os"

var version = "dev"

func init() {
	if v := os.Getenv("VERSION"); v != "" {
		version = v
	}
}

// Version returns the version of the binary.
func Version() string {
	return version
}
