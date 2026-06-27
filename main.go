package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args

	if len(os.Args) < 2 {
		fmt.Println("Error: Please provide a command (add, update, delete, list, mark-in-progress, mark-done)")
		os.Exit(1)
	}

	switch args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide a task")
			os.Exit(1)
		}
		description := strings.Join(args[2:], " ")
		result := Add(description)
		fmt.Println(result)
	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Error: Please provide an ID and a task description")
			os.Exit(1)
		}
		idArg := args[2]
		description := strings.Join(args[3:], " ")
		result := updateTask(idArg, description)
		fmt.Println(result)
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide an ID")
			os.Exit(1)
		}
		idArg := args[2]
		deleteTask(idArg)
	case "list":
		statusFilter := ""
		if len(args) > 2 {
			statusFilter = args[2]
		}
		listTasks(statusFilter)
	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide an ID")
			os.Exit(1)
		}
		idArg := args[2]
		markTask(idArg, "in-progress")
	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Error: Please provide an ID")
			os.Exit(1)
		}
		idArg := args[2]
		markTask(idArg, "done")

	default:
		fmt.Println("Unknown command")
	}
}
