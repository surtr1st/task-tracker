package internal

import (
	"encoding/json"
	"log"
	"os"
)

type taskParser struct {
	path     string
	fileByte []byte
}

func UseTaskParser() taskParser {
	return taskParser{}
}

func (parser *taskParser) SetFile(filepath string) {
	if filepath == "" {
		log.Fatal("File path is empty!")
	}

	dat, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}

	parser.path = filepath
	parser.fileByte = dat
}

func (parser taskParser) Get() TaskList {
	content := parser.fileByte
	if content == nil {
		log.Fatal("Content is missing or invalid!")
	}

	var tasks TaskList
	if err := json.Unmarshal(content, &tasks); err != nil {
		log.Fatal(err)
	}

	return tasks
}

func (parser taskParser) ToMap() map[int]Task {
	tasks := parser.Get().Tasks
	table := make(map[int]Task)
	for _, task := range tasks {
		table[task.Id] = task
	}
	return table
}

func (parser taskParser) Compose(content []byte) error {
	if err := os.WriteFile(parser.path, content, 0644); err != nil {
		return err
	}
	return nil
}
