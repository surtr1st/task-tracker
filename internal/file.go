package internal

import (
	"fmt"
	"os"
)

func InitData(name string) error {
	if _, err := os.Open(name); err != nil {
		var dir string
		if IsWindows() {
			dir = fmt.Sprintf("%s\\%s", UseEnv(WINDOWS), BASE_FOLDER)
		} else {
			dir = fmt.Sprintf("%s/.local/%s", UseEnv(LINUX), BASE_FOLDER)
		}

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

func Data() string {
	if IsWindows() {
		return fmt.Sprintf("%s\\%s\\%s", UseEnv(WINDOWS), BASE_FOLDER, FILEDATA)
	}
	return fmt.Sprintf("%s/.local/%s/%s", UseEnv(LINUX), BASE_FOLDER, FILEDATA)
}
