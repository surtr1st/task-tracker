package internal

import (
	"testing"
)

var task = UseTaskTracker()

func TestAddTask(t *testing.T) {
	description := "Test adding #1"
	if _, err := task.Add(description); err != nil {
		t.Fatal(err)
	}
}

func TestUpdateTaskDesc(t *testing.T) {
	id := 1
	description := "Test updated #1"
	if _, err := task.Update(id, description, UPDATE_DESCRIPTION); err != nil {
		t.Fatal(err)
	}
}

func TestSetTaskDone(t *testing.T) {
	id := 1
	if _, err := task.Update(id, string(DONE), UPDATE_STATUS); err != nil {
		t.Fatal(err)
	}
}

func TestRemoveTask(t *testing.T) {
	id := 1
	if _, err := task.Remove(id); err != nil {
		t.Fatal(err)
	}
}
