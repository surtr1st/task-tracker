package internal

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Subcommand struct {
	args []string
}

func UseSubcommand(args []string) *Subcommand {
	return &Subcommand{args}
}

func (subcmd Subcommand) Init() {
	if err := InitData(Data()); err != nil {
		LogError(err.Error())
	}
}

func (subcmd Subcommand) ParseSubcommands() {
	task := UseTaskTracker()

	args := subcmd.args

	if len(subcmd.args) == 0 {
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
			if strings.TrimSpace(description) == "" {
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
		if len(args) > 2 {
			id, err := strconv.Atoi(args[1])
			if err != nil {
				LogError(INVALID_INPUT)
			}

			description := args[2]
			if strings.TrimSpace(description) == "" {
				LogError(MISSING_CONTENT)
			}
			result, err := task.Update(id, description, UPDATE_DESCRIPTION)
			if err != nil {
				LogError(err.Error())
			}
			fmt.Println(result)
		} else {
			LogError(MISSING_REQUIRED_VALUES)
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

	case FLUSH:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Are you sure you want to remove all available tasks? y/n: ")
		for scanner.Scan() {
			answer := strings.ToLower(scanner.Text())
			if answer == "y" {
				fmt.Println("Clearing all tasks...")
				result, err := task.RemoveAll()
				if err != nil {
					LogError(err.Error())
				}
				fmt.Println(result)
				break
			} else if answer == "n" {
				break
			}
		}

	default:
		showHelpCommand()
	}
}

func showHelpCommand() {
	subcommands := []string{INIT, ADD, UPDATE, REMOVE, LIST, _DONE, _IN_PROGRESS, _TODO, HELP, FLUSH}
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
		FLUSH:        "Flush and clear all tasks",
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
