package internal

import (
	"fmt"
	"os"
)

func InitData(targetFile string) error {
	if _, err := os.OpenFile(targetFile, os.O_RDWR|os.O_EXCL, 0666); err != nil {
		if os.IsNotExist(err) {
			var dir string
			if IsWindows() {
				dir = fmt.Sprintf("%s\\%s", UseEnv(WINDOWS), BASE_FOLDER)
			} else {
				dir = fmt.Sprintf("%s/.local/%s", UseEnv(LINUX), BASE_FOLDER)
			}

			if mkdirErr := os.Mkdir(dir, os.ModePerm); mkdirErr != nil {
				if os.IsExist(err) {
					f, createErr := os.Create(targetFile)
					if createErr != nil {
						return createErr
					}
					defer f.Close()

					baseTemplate := `
					{
						"tasks": []
					}
					`
					if writeErr := os.WriteFile(targetFile, []byte(baseTemplate), os.ModePerm); writeErr != nil {
						return writeErr
					}
				}
				return mkdirErr
			}
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
