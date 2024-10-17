package internal

import "time"

type TaskStatus string

const (
	TODO        TaskStatus = "To Do"
	IN_PROGRESS TaskStatus = "In-Progress"
	DONE        TaskStatus = "Done"
)

type Task struct {
	Id          int        `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}
