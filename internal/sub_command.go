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
		description := args[1]
		fmt.Println(task.Add(description))

	case UPDATE:
		id, err := strconv.Atoi(args[1])
		if err != nil {
			LogError(INVALID_INPUT)
		}
		switch args[2] {
		case string(FLAG_IN_PROGRESS), string(FLAG_DONE), string(FLAG_TODO):
			fmt.Println(task.Update(id, args[2], UPDATE_STATUS))
		default:
			fmt.Println(task.Update(id, args[2], UPDATE_DESCRIPTION))
		}

	case REMOVE:
		id, err := strconv.Atoi(args[1])
		if err != nil {
			LogError(INVALID_INPUT)
		}
		fmt.Println(task.Remove(id))

	case _DONE:
		id := args[1]
		fmt.Println(id)

	case LIST:
		if len(args) > 1 {
			task.List(FilterFlagList(args[1]))
		} else {
			task.List(FLAG_NONE)
		}

	case _IN_PROGRESS:
		id := args[1]
		fmt.Println(id)
	}
}
