package internal

import (
	"fmt"
	"os"
)

func InitStorage(name string) error {
	if _, err := os.Open(name); err != nil {
		var home string
		if IsWindows() {
			home = UseEnv(WINDOWS)
		} else {
			home = UseEnv(LINUX)
		}

		dir := fmt.Sprintf("%s\\%s", home, BASE_FOLDER)
		if mkdirErr := os.Mkdir(dir, os.ModePerm); mkdirErr != nil {
			return mkdirErr
		}

		f, createErr := os.Create(name)
		if createErr != nil {
			return createErr
		}
		defer f.Close()

		baseTemplate := `
		{
			"tasks": []
		}
		`
		if writeErr := os.WriteFile(name, []byte(baseTemplate), 0644); writeErr != nil {
			return writeErr
		}
	}
	return nil
}
