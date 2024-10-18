package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type taskTracker struct {
	parser taskParser
}

func UseTaskTracker() *taskTracker {
	return &taskTracker{
		parser: UseTaskParser(),
	}
}

func (tracker taskTracker) Add(description string) {}

func (tracker taskTracker) Update(id int, value string) (string, error) {
	tasks := tracker.parser.ToMap()
	if task, ok := tasks[id]; ok {
		task.Description = value
		task.UpdatedAt = time.Now()
		tasks[id] = task
	}

	updatedContent, err := json.Marshal(tasks)
	if err != nil {
		log.Fatal(err)
	}
	if err := tracker.parser.Compose(updatedContent); err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("Updated task with id: %d", id), nil
}

func (tracker taskTracker) List() {}
