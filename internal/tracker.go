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

func (tracker taskTracker) Add(description string) (string, error) {
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
		return "", err
	}
	if err := tracker.parser.Compose(value); err != nil {
		return "", err
	}

	return fmt.Sprintf(ADDED_TASK, id), nil
}

func (tracker taskTracker) Update(id int, value string, filter FilterUpdateProperty) (string, error) {
	tasks := tracker.parser.Get().Tasks

	if _, ok := tracker.parser.VerifyTable()[id]; !ok {
		return "", fmt.Errorf(INEXISTENCE_TASK, id)
	}

	for i, task := range tasks {
		if task.Id == id {
			updatedAt := FormatDate(time.Now())
			switch filter {
			case UPDATE_DESCRIPTION:
				tasks[i].Description = value
				tasks[i].UpdatedAt = updatedAt
			case UPDATE_STATUS:
				tasks[i].Status = TaskStatus(value)
				tasks[i].UpdatedAt = updatedAt
			}
		}
	}

	updatedContent, err := json.MarshalIndent(TaskList{
		Tasks: tasks,
	}, "", " ")
	if err != nil {
		return "", err
	}
	if err := tracker.parser.Compose(updatedContent); err != nil {
		return "", err
	}

	return fmt.Sprintf(UPDATED_TASK, id), nil
}

func (tracker taskTracker) List(filter FilterStatus) {
	switch filter {
	case FILTER_DONE:
		PrintTable(tracker.parser.Filter(DONE))

	case FILTER_IN_PROGRESS:
		PrintTable(tracker.parser.Filter(IN_PROGRESS))

	case FILTER_TODO:
		PrintTable(tracker.parser.Filter(TODO))

	case FILTER_NONE:
		PrintTable(tracker.parser.Get().Tasks)
	}
}

func (tracker taskTracker) Remove(id int) (string, error) {
	tasks := tracker.parser.Get().Tasks

	if _, ok := tracker.parser.VerifyTable()[id]; !ok {
		return "", fmt.Errorf(INEXISTENCE_TASK, id)
	}

	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}

	updatedContent, err := json.MarshalIndent(TaskList{
		Tasks: tasks,
	}, "", " ")
	if err != nil {
		return "", err
	}
	if err := tracker.parser.Compose(updatedContent); err != nil {
		return "", err
	}

	return fmt.Sprintf(REMOVED_TASK, id), nil
}
