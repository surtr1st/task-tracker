package internal

import (
	"flag"
	"fmt"
	"strconv"
)

func ParseSubcommands() {
	task := UseTaskTracker()
	flag.Parse()
	args := flag.Args()

	switch args[0] {
	case INIT:
		if err := InitData(Data()); err != nil {
			LogError(err.Error())
		}

	case ADD:
		if len(args) > 1 {
			description := args[1]
			if description == "" {
				LogError(MISSING_CONTENT)
			}
			fmt.Println(task.Add(description))
		} else {
			LogError(MISSING_CONTENT)
		}

	case UPDATE:
		if len(args) > 1 {
			id, err := strconv.Atoi(args[1])
			if err != nil {
				LogError(INVALID_INPUT)
			}
			fmt.Println(task.Update(id, args[2], UPDATE_DESCRIPTION))
		} else {
			LogError(MISSING_TASK_ID)
		}

	case REMOVE:
		if len(args) > 1 {
			id, err := strconv.Atoi(args[1])
			if err != nil {
				LogError(INVALID_INPUT)
			}
			fmt.Println(task.Remove(id))
		} else {
			LogError(MISSING_TASK_ID)
		}

	case _DONE:
		if len(args) > 1 {
			id, err := strconv.Atoi(args[1])
			if err != nil {
				LogError(INVALID_INPUT)
			}
			fmt.Println(task.Update(id, string(DONE), UPDATE_STATUS))
		} else {
			LogError(MISSING_TASK_ID)
		}

	case LIST:
		if len(args) > 1 {
			task.List(FilterStatus(args[1]))
		} else {
			task.List(FILTER_NONE)
		}

	case _IN_PROGRESS:
		if len(args) > 1 {
			id, err := strconv.Atoi(args[1])
			if err != nil {
				LogError(INVALID_INPUT)
			}
			fmt.Println(task.Update(id, string(IN_PROGRESS), UPDATE_STATUS))
		} else {
			LogError(MISSING_TASK_ID)
		}
	}
}
