package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func Add(task string) string {
	tasks := Load()
	taskID := 1
	if len(tasks) > 0 {
		taskID = tasks[len(tasks)-1].ID + 1
	}
	newTask := Task{
		ID:          taskID,
		Description: task,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	if err := Save(tasks); err != nil {
		return fmt.Sprintf("Failed to add task: %v", err)
	}
	return fmt.Sprintf("Task added successfully (ID: %d)", taskID)
}
