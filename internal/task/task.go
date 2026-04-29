package task

import "time"

type status uint8

const (
	ToDo status = iota
	InProgress
	Done
)

type Item struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Status      status    `json:"status"`
	Created     time.Time `json:"created_at"`
	Updated     time.Time `json:"updated_at"`
}
