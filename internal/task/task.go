package task

import "time"

const (
	ToDo       = "to-do"
	InProgress = "in-progress"
	Done       = "done"
)

var Statuses []string = []string{ToDo, InProgress, Done}
var NumStatuses = len(Statuses)

type Item struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}
