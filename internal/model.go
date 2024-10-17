package internal

import "time"

type TaskStatus string

const (
	TODO        TaskStatus = "To Do"
	IN_PROGRESS TaskStatus = "In-Progress"
	DONE        TaskStatus = "Done"
)

type Task struct {
	Id          int
	Description string
	Status      TaskStatus
	CreatedAt   time.Time
	UpdateAt 	time.Time
}