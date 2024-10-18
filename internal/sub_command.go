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
		description := args[2]
		fmt.Println(task.Update(id, description))
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
