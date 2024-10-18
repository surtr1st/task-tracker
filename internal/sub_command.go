package internal

import (
	"flag"
	"fmt"
	"strconv"
)

const (
	ADD             = "add"
	UPDATE          = "update"
	REMOVE          = "rm"
	LIST            = "list"
	SET_DONE        = "done"
	SET_IN_PROGRESS = "ip"
	INIT            = "init"
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
	case SET_DONE:
		id := args[1]
		fmt.Println(id)
	case LIST:
		if len(args) > 1 {
			if args[1] == "--" {
				switch args[2] {
				case SET_DONE:
					break
				case SET_IN_PROGRESS:
					break
				}
			}
		} else {
			task.List()
		}
	case SET_IN_PROGRESS:
		id := args[1]
		fmt.Println(id)
	}
}
