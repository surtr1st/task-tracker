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

func (parser taskParser) Content() TaskList {
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

func (parser taskParser) Update() {}
