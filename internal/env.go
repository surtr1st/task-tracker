package internal

import "os"

const (
	WINDOWS     = "LOCALAPPDATA"
	LINUX       = "HOME"
	MAC         = "HOME"
	BASE_FOLDER = "task-tracker"
	FILEDATA    = "data.json"
)

func UseEnv(value string) string {
	return os.Getenv(value)
}
