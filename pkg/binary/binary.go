package binary

import "os"

var productVersion = "0.0.0-dev"
var dockerTag = "0000-00-00-dev"

func init() {
	if v := os.Getenv("PRODUCT_VERSION"); v != "" {
		productVersion = v
	}
}

// ProductVersion returns the version of the product.
func ProductVersion() string {
	return productVersion
}

// DockerTag returns the docker tag of the binary.
func DockerTag() string {
	return dockerTag
}
