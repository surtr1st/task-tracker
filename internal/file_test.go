package internal

import (
	"fmt"
	"testing"
)

func TestInitStorage(t *testing.T) {
	home := UseEnv(WINDOWS)
	filename := "data.json"
	file := fmt.Sprintf("%s\\%s\\%s", home, BASE_FOLDER, filename)
	if err := InitStorage(file); err != nil {
		t.Fatal(err)
	}
}
