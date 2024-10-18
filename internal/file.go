package internal

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitData(targetFile string) error {
	if _, err := os.OpenFile(targetFile, os.O_RDWR, 0666); err != nil {
		if os.IsNotExist(err) {
			var dir string
			if IsWindows() {
				dir = filepath.Join(UseEnv(WINDOWS), BASE_FOLDER)
			} else {
				dir = filepath.Join(UseEnv(LINUX), LOCALLINUX, BASE_FOLDER)
			}

			if mkdirErr := os.MkdirAll(dir, os.ModePerm); mkdirErr != nil {
				return fmt.Errorf("failed to create directory: %w", mkdirErr)
			}

			f, createErr := os.Create(targetFile)
			if createErr != nil {
				return fmt.Errorf("failed to create target file: %w", createErr)
			}
			defer f.Close()

			baseTemplate := `{
				"tasks": []
			}`

			if writeErr := os.WriteFile(targetFile, []byte(baseTemplate), os.ModePerm); writeErr != nil {
				return fmt.Errorf("failed to write to target file: %w", writeErr)
			}
		} else {
			return fmt.Errorf("failed to open target file: %w", err)
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
