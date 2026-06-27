package main

import (
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func add(newTask string) string {
	tasks := Load()
	taskID := 0
	if len(tasks) == 0 {
		taskID = 1
	} else {
		taskID = tasks[len(tasks)-1].ID + 1
	}
}
