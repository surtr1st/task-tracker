package internal

import (
	"encoding/json"
	"os"
)

type taskParser struct {
	path     string
	fileByte []byte
}

func UseTaskParser() *taskParser {
	return &taskParser{}
}

func (parser *taskParser) SetFile(filepath string) {
	if filepath == "" {
		LogError(MISSING_FILE_PATH)
	}

	dat, err := os.ReadFile(filepath)
	if err != nil {
		if err := InitData(filepath); err != nil {
			LogError(err.Error())
		}
	}
	parser.path = filepath
	parser.fileByte = dat

}

func (parser taskParser) Get() TaskList {
	content := parser.fileByte
	if content == nil {
		LogError(MISSING_CONTENT)
	}

	var tasks TaskList
	if err := json.Unmarshal(content, &tasks); err != nil {
		LogError(err.Error())
	}

	return tasks
}

func (parser taskParser) VerifyTable() map[int]interface{} {
	verifyTable := make(map[int]interface{})
	tasks := parser.Get().Tasks
	for _, task := range tasks {
		verifyTable[task.Id] = task
	}
	return verifyTable
}

func (parser taskParser) Compose(content []byte) error {
	if err := os.WriteFile(parser.path, content, 0644); err != nil {
		return err
	}
	return nil
}
