package internal

import (
	"flag"
	"fmt"
)

const (
	ADD             = "add"
	UPDATE          = "update"
	REMOVE          = "rm"
	LIST            = "list"
	SET_DONE        = "done"
	SET_IN_PROGRESS = "ip"
)

func ParseSubcommands() {
	flag.Parse()
	args := flag.Args()
	switch args[0] {
	case UPDATE:
		id := args[1]
		newDescription := args[2]
		fmt.Println(id, newDescription)
	case REMOVE:
		id := args[1]
		fmt.Println(id)
	case SET_DONE:
		id := args[1]
		fmt.Println(id)
	case LIST:
		if args[1] == "--" {
			switch args[2] {
			case SET_DONE:
				break
			default:
				break
			}
		}
	case SET_IN_PROGRESS:
		id := args[1]
		fmt.Println(id)
	default:
		description := args[1]
		fmt.Println(description)
	}
}
