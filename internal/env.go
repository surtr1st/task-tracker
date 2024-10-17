package internal

import "os"

const (
	WINDOWS_HOME = "USERPROFILE"
	LINUX_HOME   = "HOME"
	MAC_HOME     = "HOME"
)

func UseEnv(value string) string {
	return os.Getenv(value)
}