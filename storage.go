package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func getStorageFilePath() string {
	if path := os.Getenv("TASK_TRACKER_FILE"); path != "" {
		return path
	}
	return "tasks.json"
}

func Load() []Task {
	filename := getStorageFilePath()
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		file, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
		return []Task{}
	}
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file: ", err)
		return nil
	}
	if len(content) == 0 {
		return []Task{}
	}
	var tasks []Task
	err = json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return tasks
}

func Save(task []Task) error {
	filename := getStorageFilePath()
	jsonBytes, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile(filename, jsonBytes, 0o644)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}
