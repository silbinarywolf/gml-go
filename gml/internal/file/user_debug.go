// +build debug

package file

import (
	"os/user"
	"path"
	"strings"
)

var (
	debugUsername string = "▲not-set▲"
)

func DebugUsernameFileSafe() string {
	return debugUsername
}

func init() {
	// Setup file-safe escaped username
	user, _ := user.Current()
	username := user.Username
	username = path.Clean(username)
	username = strings.Replace(username, "/", "-", -1)
	username = strings.Replace(username, "\\", "-", -1)
	username = strings.Replace(username, "_", "-", -1)
	debugUsername = username
}
