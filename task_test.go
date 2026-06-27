package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAddUpdateMarkAndDeleteTask(t *testing.T) {
	tempDir := t.TempDir()
	storageFile := filepath.Join(tempDir, "tasks.json")
	t.Setenv("TASK_TRACKER_FILE", storageFile)

	addResult := Add("Buy groceries")
	if addResult == "" {
		t.Fatal("expected add result to be non-empty")
	}

	tasks := Load()
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task after add, got %d", len(tasks))
	}
	if tasks[0].Description != "Buy groceries" {
		t.Fatalf("expected task description to be updated, got %q", tasks[0].Description)
	}

	updateResult := updateTask("1", "Buy groceries and milk")
	if updateResult == "" {
		t.Fatal("expected update result to be non-empty")
	}

	markTask("1", "done")

	tasks = Load()
	if len(tasks) != 1 {
		t.Fatalf("expected 1 task after mark, got %d", len(tasks))
	}
	if tasks[0].Status != "done" {
		t.Fatalf("expected task status to be done, got %q", tasks[0].Status)
	}

	deleteTask("1")

	tasks = Load()
	if len(tasks) != 0 {
		t.Fatalf("expected 0 tasks after delete, got %d", len(tasks))
	}

	if _, err := os.Stat(storageFile); err != nil {
		t.Fatalf("expected storage file to exist at %s: %v", storageFile, err)
	}
}
