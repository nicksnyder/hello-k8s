package debug

import (
	"encoding/json"
	"net/http"
	"os"
	"os/user"

	"github.com/nicksnyder/service/pkg/version"
)

// Serve serves a debug http endpoint.
func Serve(w http.ResponseWriter, r *http.Request) {
	host, err := os.Hostname()
	if err != nil {
		panic(err)
	}
	u, err := user.Current()
	if err != nil {
		panic(err)
	}
	enc := json.NewEncoder(w)
	enc.SetIndent("", "  ")
	if err := enc.Encode(map[string]interface{}{
		"version": version.Version(),
		"host":    host,
		"user":    u,
		// "user": map[string]string{
		// 	"gid":      u.Gid,
		// 	"homedir":  u.HomeDir,
		// 	"name":     u.Name,
		// 	"uid":      u.Uid,
		// 	"username": u.Username,
		// },
	}); err != nil {
		panic(err)
	}
}
