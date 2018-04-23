package debug

import (
	"encoding/json"
	"io"
	"os"
	"os/user"

	"github.com/nicksnyder/hello-server/pkg/version"
)

// WriteData writes debug data to w.
func WriteData(w io.Writer) error {
	host, err := os.Hostname()
	if err != nil {
		return err
	}
	u, err := user.Current()
	if err != nil {
		return err
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	return enc.Encode(map[string]interface{}{
		"Version": version.Version(),
		"Host":    host,
		"User":    u,
	})
}
