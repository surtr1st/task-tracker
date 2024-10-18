package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type taskTracker struct {
	parser *taskParser
}

func UseTaskTracker() *taskTracker {
	parser := UseTaskParser()
	if IsWindows() {
		parser.SetFile(fmt.Sprintf("%s\\%s\\%s", UseEnv(WINDOWS), BASE_FOLDER, FILEDATA))
	} else {
		parser.SetFile(fmt.Sprintf("%s/.local/%s/%s", UseEnv(LINUX), BASE_FOLDER, FILEDATA))
	}
	return &taskTracker{
		parser,
	}
}

func (tracker taskTracker) Add(description string) string {
	id := 1
	totalTask := len(tracker.parser.Get().Tasks)
	if totalTask > 1 {
		id = totalTask + 1
	}

	task := Task{
		Id:          id,
		Description: description,
		Status:      TODO,
		CreatedAt:   time.Now().String(),
		UpdatedAt:   "",
	}
	tasks := TaskList{
		Tasks: []Task{task},
	}

	value, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	if err := tracker.parser.Compose(value); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Added new task! ID: %d", id)
}

func (tracker taskTracker) Update(id int, value string) string {
	tasks := tracker.parser.Get().Tasks
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Description = value
			tasks[i].UpdatedAt = time.Now().String()
		}
	}

	updatedContent, err := json.MarshalIndent(TaskList{
		Tasks: tasks,
	}, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	if err := tracker.parser.Compose(updatedContent); err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Updated task with ID: %d", id)
}

func (tracker taskTracker) List() {}
