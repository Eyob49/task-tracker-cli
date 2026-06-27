package main

import (
	"fmt"
	"strconv"
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

func listTasks(statusFilter string) {
	tasks := Load()
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	foundTask := false
	for _, task := range tasks {
		if task.Status == statusFilter || statusFilter == "" {
			fmt.Printf("[ID: %d] %s - Status: %s\n", task.ID, task.Description, task.Status)
			foundTask = true
		}
	}
	if !foundTask {
		if statusFilter == "" {
			fmt.Println("No tasks found.")
		} else {
			fmt.Printf("No tasks found with filter: %s\n", statusFilter)
		}
	}
}

func updateTask(idStr string, newDescription string) string {
	targetID, err := strconv.Atoi(idStr)
	if err != nil {
		return fmt.Sprintf("Invalid task ID: %v", err)
	}
	tasks := Load()
	found := false
	for i, task := range tasks {
		if task.ID == targetID {
			tasks[i].Description = newDescription
			tasks[i].UpdatedAt = time.Now()
			found = true
		}
	}
	if !found {
		return "ID doesn't exist."
	}
	if err := Save(tasks); err != nil {
		return fmt.Sprintf("Failed to update task: %v", err)
	}
	return fmt.Sprintf("Task updated successfully (ID: %d)", targetID)
}

func markTask(idStr string, newStatus string) {
	targetID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID: ", err)
		return
	}
	tasks := Load()
	found := false
	for i, task := range tasks {
		if task.ID == targetID {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = time.Now()
			found = true
		}
	}
	if err := Save(tasks); err != nil {
		fmt.Printf("Failed to update task: %v", err)
	}
	if !found {
		fmt.Println("ID doesn't exist.")
	}
}

func deleteTask(idStr string) {
	targetID, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Invalid task ID: ", err)
		return
	}
	tasks := Load()
	found := false
	for i := range tasks {
		if tasks[i].ID == targetID {
			tasks = append(tasks[:i], tasks[i+1:]...)
			found = true
			break
		}
	}
	if !found {
		fmt.Println("ID doesn't exist.")
		return
	}
	if err := Save(tasks); err != nil {
		fmt.Printf("Failed to delete task: %v", err)
		return
	}
	fmt.Println("Task deleted successfully.")
}
