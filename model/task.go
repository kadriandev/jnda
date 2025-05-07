package model

import "time"

type Task struct {
	Id          int64
	Title       string
	Description string
	Status      string
	DueDate     time.Time
	UpdatedAt   time.Time
	CreatedAt   time.Time
}
