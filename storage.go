package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

func Load() []Task {
	filename := "tasks.json"
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		_, err := os.Create(filename)
		if err != nil {
			fmt.Println(err)
			return nil
		}
		return []Task{}
	}
	content, err := os.ReadFile("tasks.json")
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
	jsonBytes, err := json.MarshalIndent(task, "", "    ")
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", jsonBytes, 0o644)
	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}
	return nil
}
