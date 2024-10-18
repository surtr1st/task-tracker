package internal

import (
	"fmt"
	"os"
	"runtime"
	"time"
)

func UseEnv(value string) string {
	return os.Getenv(value)
}

func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func PrintTable(tasks []Task) {
	headers := []string{
		"ID",
		"STATUS",
		"DESCRIPTION",
		"CREATED AT",
		"UPDATED AT",
	}

	columnWidths := make([]int, len(headers))
	for i, header := range headers {
		columnWidths[i] = len(header)
	}

	taskTable := toTable(tasks)
	for _, row := range taskTable {
		for i, col := range row {
			if len(col) > columnWidths[i] {
				columnWidths[i] = len(col)
			}
		}
	}

	fmt.Println(formatRow(headers, columnWidths))
	for _, row := range taskTable {
		fmt.Println(formatRow(row, columnWidths))
	}
}

func formatRow(row []string, columnWidths []int) string {
	formatted := ""
	extraWidth := 3
	for i, col := range row {
		formatted += fmt.Sprintf("%-*s ", columnWidths[i]+extraWidth, col)
	}
	return formatted
}

func toTable(tasks []Task) [][]string {
	table := [][]string{}

	for _, task := range tasks {
		row := []string{
			fmt.Sprintf("%d", task.Id),
			string(task.Status),
			task.Description,
			task.CreatedAt,
			task.UpdatedAt,
		}
		table = append(table, row)
	}
	return table
}

func FormatDate(target time.Time) string {
	format := "01/02/2006 15:04" // mm/dd/yyyy
	return target.Format(format)
}

func LogError(message string) {
	fmt.Fprintf(os.Stderr, "Error: %s\n", message)
	os.Exit(1)
}
