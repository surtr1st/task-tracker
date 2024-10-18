package internal

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

var parser = UseTaskParser()

func TestCompose(t *testing.T) {
	parser.SetFile(fmt.Sprintf("%s\\%s\\%s", UseEnv(WINDOWS), BASE_FOLDER, FILEDATA))
	task := Task{
		Id:          1,
		Description: "Test",
		Status:      TODO,
		CreatedAt:   FormatDate(time.Now()),
		UpdatedAt:   "-",
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

func TestAppend(t *testing.T) {
	parser.SetFile(fmt.Sprintf("%s\\%s\\%s", UseEnv(WINDOWS), BASE_FOLDER, FILEDATA))
	prevTasks := parser.Get().Tasks
	task := Task{
		Id:          2,
		Description: "Test",
		Status:      TODO,
		CreatedAt:   FormatDate(time.Now()),
		UpdatedAt:   "-",
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
