package internal

import (
	"encoding/json"
	"fmt"
	"time"
)

type taskTracker struct {
	parser *taskParser
}

func UseTaskTracker() *taskTracker {
	parser := UseTaskParser()
	parser.SetFile(Data())
	return &taskTracker{
		parser,
	}
}

func (tracker taskTracker) Add(description string) string {
	id := 1
	tasks := tracker.parser.Get().Tasks
	totalTask := len(tasks)
	if totalTask >= 1 {
		id = totalTask + 1
	}

	task := Task{
		Id:          id,
		Description: description,
		Status:      TODO,
		CreatedAt:   FormatDate(time.Now()),
		UpdatedAt:   "-",
	}
	tasks = append(tasks, task)
	value, err := json.MarshalIndent(TaskList{Tasks: tasks}, "", " ")
	if err != nil {
		LogError(err.Error())
	}
	if err := tracker.parser.Compose(value); err != nil {
		LogError(err.Error())
	}

	return fmt.Sprintf("Added new task! ID: %d", id)
}

func (tracker taskTracker) Update(id int, value string) string {
	tasks := tracker.parser.Get().Tasks

	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Description = value
			tasks[i].UpdatedAt = FormatDate(time.Now())
		}
	}

	updatedContent, err := json.MarshalIndent(TaskList{
		Tasks: tasks,
	}, "", " ")
	if err != nil {
		LogError(err.Error())
	}
	if err := tracker.parser.Compose(updatedContent); err != nil {
		LogError(err.Error())
	}

	return fmt.Sprintf("Updated task with ID: %d", id)
}

func (tracker taskTracker) List() {
	PrintTable(tracker.parser.Get().Tasks)
}

func (tracker taskTracker) Remove(id int) string {
	tasks := tracker.parser.Get().Tasks

	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	updatedContent, err := json.MarshalIndent(TaskList{
		Tasks: tasks,
	}, "", " ")
	if err != nil {
		LogError(err.Error())
	}
	if err := tracker.parser.Compose(updatedContent); err != nil {
		LogError(err.Error())
	}

	return fmt.Sprintf("Removed task with ID: %d", id)
}
