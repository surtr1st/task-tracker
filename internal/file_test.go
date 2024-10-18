package internal

import (
	"testing"
)

func TestInitData(t *testing.T) {
	if err := InitData(Data()); err != nil {
		t.Fatal(err)
	}
}
