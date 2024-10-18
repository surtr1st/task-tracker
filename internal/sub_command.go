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

	if len(args) == 0 {
		showHelpCommand()
		return
	}

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

			result, err := task.Add(description)
			if err != nil {
				LogError(err.Error())
			}
			fmt.Println(result)
		} else {
			LogError(MISSING_CONTENT)
		}

	case UPDATE:
		if len(args) > 1 {
			id, err := strconv.Atoi(args[1])
			if err != nil {
				LogError(INVALID_INPUT)
			}

			result, err := task.Update(id, args[2], UPDATE_DESCRIPTION)
			if err != nil {
				LogError(err.Error())
			}
			fmt.Println(result)
		} else {
			LogError(MISSING_TASK_ID)
		}

	case REMOVE:
		if len(args) > 1 {
			id, err := strconv.Atoi(args[1])
			if err != nil {
				LogError(INVALID_INPUT)
			}

			result, err := task.Remove(id)
			if err != nil {
				LogError(err.Error())
			}
			fmt.Println(result)
		} else {
			LogError(MISSING_TASK_ID)
		}

	case _DONE:
		if len(args) > 1 {
			id, err := strconv.Atoi(args[1])
			if err != nil {
				LogError(INVALID_INPUT)
			}

			result, err := task.Update(id, string(DONE), UPDATE_STATUS)
			if err != nil {
				LogError(err.Error())
			}
			fmt.Println(result)
		} else {
			LogError(MISSING_TASK_ID)
		}

	case LIST:
		if len(args) > 1 {
			if args[1] == HELP {
				showHelpCommandOfList()
			} else {
				task.List(FilterStatus(args[1]))
			}
		} else {
			task.List(FILTER_NONE)
		}

	case _IN_PROGRESS:
		if len(args) > 1 {
			id, err := strconv.Atoi(args[1])
			if err != nil {
				LogError(INVALID_INPUT)
			}

			result, err := task.Update(id, string(IN_PROGRESS), UPDATE_STATUS)
			if err != nil {
				LogError(err.Error())
			}
			fmt.Println(result)
		} else {
			LogError(MISSING_TASK_ID)
		}

	default:
		showHelpCommand()
	}
}

func showHelpCommand() {
	subcommands := []string{INIT, ADD, UPDATE, REMOVE, LIST, _DONE, _IN_PROGRESS, _TODO, HELP}
	descriptives := map[string]string{
		INIT:         "Initializing file data for storing tasks",
		ADD:          "Adding new task",
		UPDATE:       "Updating a specific task by ID",
		REMOVE:       "Removing a specific task by ID",
		LIST:         "Listing all tasks",
		_DONE:        "Set status 'Done' to a task by ID",
		_IN_PROGRESS: "Set status 'In-Progress' to a task by ID",
		_TODO:        "Set status 'To Do' to a task by ID",
		HELP:         "Show subcommand list",
	}

	columnWidth := 10
	fmt.Print("Task Tracker CLI by ishi (surtr1st)\n\n")
	fmt.Print("Standard Commands\n\n")
	for _, subcommand := range subcommands {
		if value, ok := descriptives[subcommand]; ok {
			fmt.Printf("%-*s", columnWidth, subcommand)
			fmt.Printf("%-*s\n", columnWidth, value)
		}
	}
	fmt.Println()
}

func showHelpCommandOfList() {
	subcommands := []string{string(FILTER_DONE), string(FILTER_IN_PROGRESS), string(FILTER_TODO), HELP}
	descriptives := map[string]string{
		string(FILTER_DONE):        "Show all tasks that has status 'Done'",
		string(FILTER_IN_PROGRESS): "Show all tasks that has status 'In-Progress'",
		string(FILTER_TODO):        "Show all tasks that has status 'To Do'",
		HELP:                       "Show all subcommands of list",
	}

	columnWidth := 10
	fmt.Print("Task Tracker CLI by ishi (surtr1st)\n\n")
	fmt.Print("List Subcommands\n\n")
	for _, subcommand := range subcommands {
		if value, ok := descriptives[subcommand]; ok {
			fmt.Printf("%-*s", columnWidth, subcommand)
			fmt.Printf("%-*s\n", columnWidth, value)
		}
	}
	fmt.Println()
}
