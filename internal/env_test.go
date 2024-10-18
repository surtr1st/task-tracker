package internal

import "testing"

func TestGetWindowsLocal(t *testing.T) {
	localFolder := UseEnv(WINDOWS)
	if localFolder == "" {
		t.Fatal("Env undefined")
	}
}
