package internal

import "os"

func UseEnv(value string) string {
	return os.Getenv(value)
}
