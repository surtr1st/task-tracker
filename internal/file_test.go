package internal

import (
	"fmt"
	"testing"
)

func TestInitStorage(t *testing.T) {
	file := fmt.Sprintf("%s\\%s\\%s", UseEnv(WINDOWS), BASE_FOLDER, FILEDATA)
	if err := InitStorage(file); err != nil {
		t.Fatal(err)
	}
}
