package internal

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestCompose(t *testing.T) {
	parser := UseTaskParser()
	parser.SetFile(fmt.Sprintf("%s\\%s\\%s", UseEnv(WINDOWS), BASE_FOLDER, FILEDATA))
	task := Task{
		Id:          1,
		Description: "Test",
		Status:      "To Do",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}
	tasks := TaskList{
		Tasks: []Task{task},
	}
	value, err := json.Marshal(tasks)
	if err != nil {
		t.Fatal(err)
	}
	if err := parser.Compose(value); err != nil {
		t.Fatal(err)
	}
}

func TestAppending(t *testing.T) {
	parser := UseTaskParser()
	parser.SetFile(fmt.Sprintf("%s\\%s\\%s", UseEnv(WINDOWS), BASE_FOLDER, FILEDATA))
	prevTasks := parser.Get().Tasks
	task := Task{
		Id:          2,
		Description: "Test",
		Status:      "To Do",
		CreatedAt:   time.Now().String(),
		UpdatedAt:   time.Now().String(),
	}
	prevTasks = append(prevTasks, task)
	tasks := TaskList{
		Tasks: prevTasks,
	}
	value, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		t.Fatal(err)
	}
	if err := parser.Compose(value); err != nil {
		t.Fatal(err)
	}
}
